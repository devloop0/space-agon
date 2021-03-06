module github.com/laremere/space-agon

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

go 1.13

require (
	agones.dev/agones v1.7.0
	github.com/golang/protobuf v1.3.2
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0
	google.golang.org/grpc v1.25.0
	k8s.io/apimachinery v0.15.11
	k8s.io/client-go v9.0.0+incompatible
	open-match.dev/open-match v0.4.1-0.20191113214301-7a1dcbdf9324
)

replace (
	agones.dev/agones => agones.dev/agones v1.7.0
	k8s.io/api => k8s.io/api v0.15.11
	k8s.io/apimachinery => k8s.io/apimachinery v0.15.11
	k8s.io/client-go => k8s.io/client-go v0.15.11
)
