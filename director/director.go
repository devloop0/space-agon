// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"open-match.dev/open-match/pkg/pb"

	allocationpb "agones.dev/agones/pkg/allocation/go"
	agonesv1 "agones.dev/agones/pkg/apis/agones/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

const (
	gcgsRealmLabelName = "gameservices.googleapis.com/realm"
	gameServerDeploymentName = "<DEPLOYMENT-NAME>"
	allocatorIP = "<IP>"
	allocatorPort = "443"
)

var (
	realmNames = [2]string{"space-agon-realm1", "space-agon-realm2"}
	realmIndex = 0
)

func main() {
	log.Println("Starting Director")

	allocatorEndpoint := allocatorIP + ":" + allocatorPort
	log.Printf("Trying to connect to endpoint %v", allocatorEndpoint)
	dialOpts := grpc.WithInsecure()
	conn, err := grpc.Dial(allocatorEndpoint, dialOpts)
	if err != nil {
		log.Fatalf("Could not connect to allocator(%v): %v", allocatorEndpoint, err)
	}
	grpcClient := allocationpb.NewAllocationServiceClient(conn)

	for range time.Tick(time.Second) {
		if err := run(grpcClient); err != nil {
			log.Println("Error running director:", err.Error())
		}
	}
}

func createOMBackendClient() (pb.BackendClient, func() error) {
	conn, err := grpc.Dial("om-backend.open-match.svc.cluster.local:50505", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewBackendClient(conn), conn.Close
}

func createAgonesClient() *versioned.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	agonesClient, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return agonesClient
}

// Customize the backend.FetchMatches request, the default one will return all tickets in the statestore
func createOMFetchMatchesRequest() *pb.FetchMatchesRequest {
	return &pb.FetchMatchesRequest{
		// om-function:50502 -> the internal hostname & port number of the MMF service in our Kubernetes cluster
		Config: &pb.FunctionConfig{
			Host: "mmf.default.svc.cluster.local",
			Port: 50502,
			Type: pb.FunctionConfig_GRPC,
		},
		Profiles: []*pb.MatchProfile{
			{
				Name: "1v1",
				Pools: []*pb.Pool{
					{
						Name: "everyone",
					},
				},
			},
		},
	}
}

func createAgonesAllocationRequest() *allocationpb.AllocationRequest {
	currRealm := realmNames[realmIndex]
	realmIndex = (realmIndex + 1) % len(realmNames)
	log.Printf("Chose realm %v to create AllocationRequest for", currRealm)
	return &allocationpb.AllocationRequest{
		Namespace: "default",
		MultiClusterSetting: &allocationpb.MultiClusterSetting{
			Enabled: true,
			PolicySelector: &allocationpb.LabelSelector{
				MatchLabels: map[string]string{gcgsRealmLabelName: currRealm},
			},
		},
		RequiredGameServerSelector: &allocationpb.LabelSelector{
			MatchLabels: map[string]string{agonesv1.FleetNameLabel: gameServerDeploymentName},
		},
	}
}

func createOMAssignTicketRequestFromAllocationResponse(match *pb.Match, ar *allocationpb.AllocationResponse) *pb.AssignTicketsRequest {
	tids := []string{}
	for _, t := range match.GetTickets() {
		tids = append(tids, t.GetId())
	}

	return &pb.AssignTicketsRequest{
		TicketIds: tids,
		Assignment: &pb.Assignment{
			Connection: fmt.Sprintf("%s:%d", ar.Address, ar.Ports[0].Port),
		},
	}
}

func run(grpcClient allocationpb.AllocationServiceClient) error {
	bc, closer := createOMBackendClient()
	defer closer()

	// agonesClient := createAgonesClient()

	stream, err := bc.FetchMatches(context.Background(), createOMFetchMatchesRequest())
	if err != nil {
		return fmt.Errorf("fail to get response stream from backend.FetchMatches call: %w", err)
	}

	totalMatches := 0

	// Read the FetchMatches response. Each loop fetches an available game match that satisfies the match profiles.
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error streaming response from backend.FetchMatches call: %w", err)
		}

		ar, err := grpcClient.Allocate(context.Background(), createAgonesAllocationRequest())
		if err != nil {
			return fmt.Errorf("allocation error: %w", err)
		}

		log.Printf("AllocationResponse: %+v", *ar)

		// TODO: This will leak gameservers from people who left the queue early (before their game started)
		bc.AssignTickets(context.Background(), createOMAssignTicketRequestFromAllocationResponse(resp.GetMatch(), ar))

		totalMatches++
	}

	log.Printf("Created and assigned %d matches", totalMatches)

	return nil
}
