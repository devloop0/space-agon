// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/pb/messages.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ClientInitialize struct {
	Cid                  int64    `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInitialize) Reset()         { *m = ClientInitialize{} }
func (m *ClientInitialize) String() string { return proto.CompactTextString(m) }
func (*ClientInitialize) ProtoMessage()    {}
func (*ClientInitialize) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{0}
}

func (m *ClientInitialize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInitialize.Unmarshal(m, b)
}
func (m *ClientInitialize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInitialize.Marshal(b, m, deterministic)
}
func (m *ClientInitialize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInitialize.Merge(m, src)
}
func (m *ClientInitialize) XXX_Size() int {
	return xxx_messageInfo_ClientInitialize.Size(m)
}
func (m *ClientInitialize) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInitialize.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInitialize proto.InternalMessageInfo

func (m *ClientInitialize) GetCid() int64 {
	if m != nil {
		return m.Cid
	}
	return 0
}

type Memos struct {
	Memos                []*Memo  `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memos) Reset()         { *m = Memos{} }
func (m *Memos) String() string { return proto.CompactTextString(m) }
func (*Memos) ProtoMessage()    {}
func (*Memos) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{1}
}

func (m *Memos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memos.Unmarshal(m, b)
}
func (m *Memos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memos.Marshal(b, m, deterministic)
}
func (m *Memos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memos.Merge(m, src)
}
func (m *Memos) XXX_Size() int {
	return xxx_messageInfo_Memos.Size(m)
}
func (m *Memos) XXX_DiscardUnknown() {
	xxx_messageInfo_Memos.DiscardUnknown(m)
}

var xxx_messageInfo_Memos proto.InternalMessageInfo

func (m *Memos) GetMemos() []*Memo {
	if m != nil {
		return m.Memos
	}
	return nil
}

type Memo struct {
	// Types that are valid to be assigned to Recipient:
	//	*Memo_To
	//	*Memo_EveryoneBut
	//	*Memo_Everyone
	Recipient isMemo_Recipient `protobuf_oneof:"recipient"`
	// Types that are valid to be assigned to Actual:
	//	*Memo_PosTracks
	//	*Memo_MomentumTracks
	//	*Memo_RotTracks
	//	*Memo_SpinTracks
	//	*Memo_ShipControlTrack
	//	*Memo_DestroyEvent
	//	*Memo_ShootMissile
	//	*Memo_SpawnMissile
	//	*Memo_SpawnExplosion
	//	*Memo_SpawnShip
	//	*Memo_RegisterPlayer
	Actual               isMemo_Actual `protobuf_oneof:"actual"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Memo) Reset()         { *m = Memo{} }
func (m *Memo) String() string { return proto.CompactTextString(m) }
func (*Memo) ProtoMessage()    {}
func (*Memo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{2}
}

func (m *Memo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memo.Unmarshal(m, b)
}
func (m *Memo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memo.Marshal(b, m, deterministic)
}
func (m *Memo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memo.Merge(m, src)
}
func (m *Memo) XXX_Size() int {
	return xxx_messageInfo_Memo.Size(m)
}
func (m *Memo) XXX_DiscardUnknown() {
	xxx_messageInfo_Memo.DiscardUnknown(m)
}

var xxx_messageInfo_Memo proto.InternalMessageInfo

type isMemo_Recipient interface {
	isMemo_Recipient()
}

type Memo_To struct {
	To int64 `protobuf:"varint,1,opt,name=to,proto3,oneof"`
}

type Memo_EveryoneBut struct {
	EveryoneBut int64 `protobuf:"varint,2,opt,name=everyone_but,json=everyoneBut,proto3,oneof"`
}

type Memo_Everyone struct {
	Everyone bool `protobuf:"varint,3,opt,name=everyone,proto3,oneof"`
}

func (*Memo_To) isMemo_Recipient() {}

func (*Memo_EveryoneBut) isMemo_Recipient() {}

func (*Memo_Everyone) isMemo_Recipient() {}

func (m *Memo) GetRecipient() isMemo_Recipient {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Memo) GetTo() int64 {
	if x, ok := m.GetRecipient().(*Memo_To); ok {
		return x.To
	}
	return 0
}

func (m *Memo) GetEveryoneBut() int64 {
	if x, ok := m.GetRecipient().(*Memo_EveryoneBut); ok {
		return x.EveryoneBut
	}
	return 0
}

func (m *Memo) GetEveryone() bool {
	if x, ok := m.GetRecipient().(*Memo_Everyone); ok {
		return x.Everyone
	}
	return false
}

type isMemo_Actual interface {
	isMemo_Actual()
}

type Memo_PosTracks struct {
	PosTracks *PosTracks `protobuf:"bytes,10,opt,name=pos_tracks,json=posTracks,proto3,oneof"`
}

type Memo_MomentumTracks struct {
	MomentumTracks *MomentumTracks `protobuf:"bytes,11,opt,name=momentum_tracks,json=momentumTracks,proto3,oneof"`
}

type Memo_RotTracks struct {
	RotTracks *RotTracks `protobuf:"bytes,12,opt,name=rot_tracks,json=rotTracks,proto3,oneof"`
}

type Memo_SpinTracks struct {
	SpinTracks *SpinTracks `protobuf:"bytes,13,opt,name=spin_tracks,json=spinTracks,proto3,oneof"`
}

type Memo_ShipControlTrack struct {
	ShipControlTrack *ShipControlTrack `protobuf:"bytes,14,opt,name=ship_control_track,json=shipControlTrack,proto3,oneof"`
}

type Memo_DestroyEvent struct {
	DestroyEvent *DestroyEvent `protobuf:"bytes,16,opt,name=destroy_event,json=destroyEvent,proto3,oneof"`
}

type Memo_ShootMissile struct {
	ShootMissile *ShootMissile `protobuf:"bytes,17,opt,name=shoot_missile,json=shootMissile,proto3,oneof"`
}

type Memo_SpawnMissile struct {
	SpawnMissile *SpawnMissile `protobuf:"bytes,18,opt,name=spawn_missile,json=spawnMissile,proto3,oneof"`
}

type Memo_SpawnExplosion struct {
	SpawnExplosion *SpawnExplosion `protobuf:"bytes,19,opt,name=spawn_explosion,json=spawnExplosion,proto3,oneof"`
}

type Memo_SpawnShip struct {
	SpawnShip *SpawnShip `protobuf:"bytes,20,opt,name=spawn_ship,json=spawnShip,proto3,oneof"`
}

type Memo_RegisterPlayer struct {
	RegisterPlayer *RegisterPlayer `protobuf:"bytes,21,opt,name=register_player,json=registerPlayer,proto3,oneof"`
}

func (*Memo_PosTracks) isMemo_Actual() {}

func (*Memo_MomentumTracks) isMemo_Actual() {}

func (*Memo_RotTracks) isMemo_Actual() {}

func (*Memo_SpinTracks) isMemo_Actual() {}

func (*Memo_ShipControlTrack) isMemo_Actual() {}

func (*Memo_DestroyEvent) isMemo_Actual() {}

func (*Memo_ShootMissile) isMemo_Actual() {}

func (*Memo_SpawnMissile) isMemo_Actual() {}

func (*Memo_SpawnExplosion) isMemo_Actual() {}

func (*Memo_SpawnShip) isMemo_Actual() {}

func (*Memo_RegisterPlayer) isMemo_Actual() {}

func (m *Memo) GetActual() isMemo_Actual {
	if m != nil {
		return m.Actual
	}
	return nil
}

func (m *Memo) GetPosTracks() *PosTracks {
	if x, ok := m.GetActual().(*Memo_PosTracks); ok {
		return x.PosTracks
	}
	return nil
}

func (m *Memo) GetMomentumTracks() *MomentumTracks {
	if x, ok := m.GetActual().(*Memo_MomentumTracks); ok {
		return x.MomentumTracks
	}
	return nil
}

func (m *Memo) GetRotTracks() *RotTracks {
	if x, ok := m.GetActual().(*Memo_RotTracks); ok {
		return x.RotTracks
	}
	return nil
}

func (m *Memo) GetSpinTracks() *SpinTracks {
	if x, ok := m.GetActual().(*Memo_SpinTracks); ok {
		return x.SpinTracks
	}
	return nil
}

func (m *Memo) GetShipControlTrack() *ShipControlTrack {
	if x, ok := m.GetActual().(*Memo_ShipControlTrack); ok {
		return x.ShipControlTrack
	}
	return nil
}

func (m *Memo) GetDestroyEvent() *DestroyEvent {
	if x, ok := m.GetActual().(*Memo_DestroyEvent); ok {
		return x.DestroyEvent
	}
	return nil
}

func (m *Memo) GetShootMissile() *ShootMissile {
	if x, ok := m.GetActual().(*Memo_ShootMissile); ok {
		return x.ShootMissile
	}
	return nil
}

func (m *Memo) GetSpawnMissile() *SpawnMissile {
	if x, ok := m.GetActual().(*Memo_SpawnMissile); ok {
		return x.SpawnMissile
	}
	return nil
}

func (m *Memo) GetSpawnExplosion() *SpawnExplosion {
	if x, ok := m.GetActual().(*Memo_SpawnExplosion); ok {
		return x.SpawnExplosion
	}
	return nil
}

func (m *Memo) GetSpawnShip() *SpawnShip {
	if x, ok := m.GetActual().(*Memo_SpawnShip); ok {
		return x.SpawnShip
	}
	return nil
}

func (m *Memo) GetRegisterPlayer() *RegisterPlayer {
	if x, ok := m.GetActual().(*Memo_RegisterPlayer); ok {
		return x.RegisterPlayer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Memo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Memo_To)(nil),
		(*Memo_EveryoneBut)(nil),
		(*Memo_Everyone)(nil),
		(*Memo_PosTracks)(nil),
		(*Memo_MomentumTracks)(nil),
		(*Memo_RotTracks)(nil),
		(*Memo_SpinTracks)(nil),
		(*Memo_ShipControlTrack)(nil),
		(*Memo_DestroyEvent)(nil),
		(*Memo_ShootMissile)(nil),
		(*Memo_SpawnMissile)(nil),
		(*Memo_SpawnExplosion)(nil),
		(*Memo_SpawnShip)(nil),
		(*Memo_RegisterPlayer)(nil),
	}
}

type PosTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	X                    []float32 `protobuf:"fixed32,2,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y                    []float32 `protobuf:"fixed32,3,rep,packed,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PosTracks) Reset()         { *m = PosTracks{} }
func (m *PosTracks) String() string { return proto.CompactTextString(m) }
func (*PosTracks) ProtoMessage()    {}
func (*PosTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{3}
}

func (m *PosTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PosTracks.Unmarshal(m, b)
}
func (m *PosTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PosTracks.Marshal(b, m, deterministic)
}
func (m *PosTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosTracks.Merge(m, src)
}
func (m *PosTracks) XXX_Size() int {
	return xxx_messageInfo_PosTracks.Size(m)
}
func (m *PosTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_PosTracks.DiscardUnknown(m)
}

var xxx_messageInfo_PosTracks proto.InternalMessageInfo

func (m *PosTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *PosTracks) GetX() []float32 {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *PosTracks) GetY() []float32 {
	if m != nil {
		return m.Y
	}
	return nil
}

type MomentumTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	X                    []float32 `protobuf:"fixed32,2,rep,packed,name=x,proto3" json:"x,omitempty"`
	Y                    []float32 `protobuf:"fixed32,3,rep,packed,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MomentumTracks) Reset()         { *m = MomentumTracks{} }
func (m *MomentumTracks) String() string { return proto.CompactTextString(m) }
func (*MomentumTracks) ProtoMessage()    {}
func (*MomentumTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{4}
}

func (m *MomentumTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MomentumTracks.Unmarshal(m, b)
}
func (m *MomentumTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MomentumTracks.Marshal(b, m, deterministic)
}
func (m *MomentumTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MomentumTracks.Merge(m, src)
}
func (m *MomentumTracks) XXX_Size() int {
	return xxx_messageInfo_MomentumTracks.Size(m)
}
func (m *MomentumTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_MomentumTracks.DiscardUnknown(m)
}

var xxx_messageInfo_MomentumTracks proto.InternalMessageInfo

func (m *MomentumTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *MomentumTracks) GetX() []float32 {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *MomentumTracks) GetY() []float32 {
	if m != nil {
		return m.Y
	}
	return nil
}

type RotTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	R                    []float32 `protobuf:"fixed32,2,rep,packed,name=r,proto3" json:"r,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RotTracks) Reset()         { *m = RotTracks{} }
func (m *RotTracks) String() string { return proto.CompactTextString(m) }
func (*RotTracks) ProtoMessage()    {}
func (*RotTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{5}
}

func (m *RotTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RotTracks.Unmarshal(m, b)
}
func (m *RotTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RotTracks.Marshal(b, m, deterministic)
}
func (m *RotTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RotTracks.Merge(m, src)
}
func (m *RotTracks) XXX_Size() int {
	return xxx_messageInfo_RotTracks.Size(m)
}
func (m *RotTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_RotTracks.DiscardUnknown(m)
}

var xxx_messageInfo_RotTracks proto.InternalMessageInfo

func (m *RotTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *RotTracks) GetR() []float32 {
	if m != nil {
		return m.R
	}
	return nil
}

type SpinTracks struct {
	Nid                  []uint64  `protobuf:"varint,1,rep,packed,name=nid,proto3" json:"nid,omitempty"`
	S                    []float32 `protobuf:"fixed32,2,rep,packed,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SpinTracks) Reset()         { *m = SpinTracks{} }
func (m *SpinTracks) String() string { return proto.CompactTextString(m) }
func (*SpinTracks) ProtoMessage()    {}
func (*SpinTracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{6}
}

func (m *SpinTracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpinTracks.Unmarshal(m, b)
}
func (m *SpinTracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpinTracks.Marshal(b, m, deterministic)
}
func (m *SpinTracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpinTracks.Merge(m, src)
}
func (m *SpinTracks) XXX_Size() int {
	return xxx_messageInfo_SpinTracks.Size(m)
}
func (m *SpinTracks) XXX_DiscardUnknown() {
	xxx_messageInfo_SpinTracks.DiscardUnknown(m)
}

var xxx_messageInfo_SpinTracks proto.InternalMessageInfo

func (m *SpinTracks) GetNid() []uint64 {
	if m != nil {
		return m.Nid
	}
	return nil
}

func (m *SpinTracks) GetS() []float32 {
	if m != nil {
		return m.S
	}
	return nil
}

type ShipControlTrack struct {
	Nid                  uint64   `protobuf:"varint,1,opt,name=nid,proto3" json:"nid,omitempty"`
	Up                   bool     `protobuf:"varint,2,opt,name=up,proto3" json:"up,omitempty"`
	Left                 bool     `protobuf:"varint,3,opt,name=left,proto3" json:"left,omitempty"`
	Right                bool     `protobuf:"varint,4,opt,name=right,proto3" json:"right,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipControlTrack) Reset()         { *m = ShipControlTrack{} }
func (m *ShipControlTrack) String() string { return proto.CompactTextString(m) }
func (*ShipControlTrack) ProtoMessage()    {}
func (*ShipControlTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{7}
}

func (m *ShipControlTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipControlTrack.Unmarshal(m, b)
}
func (m *ShipControlTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipControlTrack.Marshal(b, m, deterministic)
}
func (m *ShipControlTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipControlTrack.Merge(m, src)
}
func (m *ShipControlTrack) XXX_Size() int {
	return xxx_messageInfo_ShipControlTrack.Size(m)
}
func (m *ShipControlTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipControlTrack.DiscardUnknown(m)
}

var xxx_messageInfo_ShipControlTrack proto.InternalMessageInfo

func (m *ShipControlTrack) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func (m *ShipControlTrack) GetUp() bool {
	if m != nil {
		return m.Up
	}
	return false
}

func (m *ShipControlTrack) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *ShipControlTrack) GetRight() bool {
	if m != nil {
		return m.Right
	}
	return false
}

type DestroyEvent struct {
	Nid                  uint64   `protobuf:"varint,1,opt,name=nid,proto3" json:"nid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyEvent) Reset()         { *m = DestroyEvent{} }
func (m *DestroyEvent) String() string { return proto.CompactTextString(m) }
func (*DestroyEvent) ProtoMessage()    {}
func (*DestroyEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{8}
}

func (m *DestroyEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyEvent.Unmarshal(m, b)
}
func (m *DestroyEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyEvent.Marshal(b, m, deterministic)
}
func (m *DestroyEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyEvent.Merge(m, src)
}
func (m *DestroyEvent) XXX_Size() int {
	return xxx_messageInfo_DestroyEvent.Size(m)
}
func (m *DestroyEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyEvent proto.InternalMessageInfo

func (m *DestroyEvent) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

type ShootMissile struct {
	Owner                uint64   `protobuf:"varint,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShootMissile) Reset()         { *m = ShootMissile{} }
func (m *ShootMissile) String() string { return proto.CompactTextString(m) }
func (*ShootMissile) ProtoMessage()    {}
func (*ShootMissile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{9}
}

func (m *ShootMissile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShootMissile.Unmarshal(m, b)
}
func (m *ShootMissile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShootMissile.Marshal(b, m, deterministic)
}
func (m *ShootMissile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShootMissile.Merge(m, src)
}
func (m *ShootMissile) XXX_Size() int {
	return xxx_messageInfo_ShootMissile.Size(m)
}
func (m *ShootMissile) XXX_DiscardUnknown() {
	xxx_messageInfo_ShootMissile.DiscardUnknown(m)
}

var xxx_messageInfo_ShootMissile proto.InternalMessageInfo

func (m *ShootMissile) GetOwner() uint64 {
	if m != nil {
		return m.Owner
	}
	return 0
}

// Server is always authority
type SpawnMissile struct {
	Nid                  uint64   `protobuf:"varint,6,opt,name=nid,proto3" json:"nid,omitempty"`
	Owner                uint64   `protobuf:"varint,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Pos                  *Vec2    `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	Momentum             *Vec2    `protobuf:"bytes,3,opt,name=momentum,proto3" json:"momentum,omitempty"`
	Rot                  float32  `protobuf:"fixed32,4,opt,name=rot,proto3" json:"rot,omitempty"`
	Spin                 float32  `protobuf:"fixed32,5,opt,name=spin,proto3" json:"spin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnMissile) Reset()         { *m = SpawnMissile{} }
func (m *SpawnMissile) String() string { return proto.CompactTextString(m) }
func (*SpawnMissile) ProtoMessage()    {}
func (*SpawnMissile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{10}
}

func (m *SpawnMissile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnMissile.Unmarshal(m, b)
}
func (m *SpawnMissile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnMissile.Marshal(b, m, deterministic)
}
func (m *SpawnMissile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnMissile.Merge(m, src)
}
func (m *SpawnMissile) XXX_Size() int {
	return xxx_messageInfo_SpawnMissile.Size(m)
}
func (m *SpawnMissile) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnMissile.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnMissile proto.InternalMessageInfo

func (m *SpawnMissile) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func (m *SpawnMissile) GetOwner() uint64 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *SpawnMissile) GetPos() *Vec2 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SpawnMissile) GetMomentum() *Vec2 {
	if m != nil {
		return m.Momentum
	}
	return nil
}

func (m *SpawnMissile) GetRot() float32 {
	if m != nil {
		return m.Rot
	}
	return 0
}

func (m *SpawnMissile) GetSpin() float32 {
	if m != nil {
		return m.Spin
	}
	return 0
}

type SpawnExplosion struct {
	Pos                  *Vec2    `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Momentum             *Vec2    `protobuf:"bytes,2,opt,name=momentum,proto3" json:"momentum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnExplosion) Reset()         { *m = SpawnExplosion{} }
func (m *SpawnExplosion) String() string { return proto.CompactTextString(m) }
func (*SpawnExplosion) ProtoMessage()    {}
func (*SpawnExplosion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{11}
}

func (m *SpawnExplosion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnExplosion.Unmarshal(m, b)
}
func (m *SpawnExplosion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnExplosion.Marshal(b, m, deterministic)
}
func (m *SpawnExplosion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnExplosion.Merge(m, src)
}
func (m *SpawnExplosion) XXX_Size() int {
	return xxx_messageInfo_SpawnExplosion.Size(m)
}
func (m *SpawnExplosion) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnExplosion.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnExplosion proto.InternalMessageInfo

func (m *SpawnExplosion) GetPos() *Vec2 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SpawnExplosion) GetMomentum() *Vec2 {
	if m != nil {
		return m.Momentum
	}
	return nil
}

type SpawnShip struct {
	Nid                  uint64   `protobuf:"varint,1,opt,name=nid,proto3" json:"nid,omitempty"`
	Authority            int64    `protobuf:"varint,2,opt,name=authority,proto3" json:"authority,omitempty"`
	Pos                  *Vec2    `protobuf:"bytes,3,opt,name=pos,proto3" json:"pos,omitempty"`
	Momentum             *Vec2    `protobuf:"bytes,4,opt,name=momentum,proto3" json:"momentum,omitempty"`
	Rot                  float32  `protobuf:"fixed32,5,opt,name=rot,proto3" json:"rot,omitempty"`
	Spin                 float32  `protobuf:"fixed32,6,opt,name=spin,proto3" json:"spin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnShip) Reset()         { *m = SpawnShip{} }
func (m *SpawnShip) String() string { return proto.CompactTextString(m) }
func (*SpawnShip) ProtoMessage()    {}
func (*SpawnShip) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{12}
}

func (m *SpawnShip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnShip.Unmarshal(m, b)
}
func (m *SpawnShip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnShip.Marshal(b, m, deterministic)
}
func (m *SpawnShip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnShip.Merge(m, src)
}
func (m *SpawnShip) XXX_Size() int {
	return xxx_messageInfo_SpawnShip.Size(m)
}
func (m *SpawnShip) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnShip.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnShip proto.InternalMessageInfo

func (m *SpawnShip) GetNid() uint64 {
	if m != nil {
		return m.Nid
	}
	return 0
}

func (m *SpawnShip) GetAuthority() int64 {
	if m != nil {
		return m.Authority
	}
	return 0
}

func (m *SpawnShip) GetPos() *Vec2 {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SpawnShip) GetMomentum() *Vec2 {
	if m != nil {
		return m.Momentum
	}
	return nil
}

func (m *SpawnShip) GetRot() float32 {
	if m != nil {
		return m.Rot
	}
	return 0
}

func (m *SpawnShip) GetSpin() float32 {
	if m != nil {
		return m.Spin
	}
	return 0
}

type RegisterPlayer struct {
	Cid                  int64    `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPlayer) Reset()         { *m = RegisterPlayer{} }
func (m *RegisterPlayer) String() string { return proto.CompactTextString(m) }
func (*RegisterPlayer) ProtoMessage()    {}
func (*RegisterPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{13}
}

func (m *RegisterPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPlayer.Unmarshal(m, b)
}
func (m *RegisterPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPlayer.Marshal(b, m, deterministic)
}
func (m *RegisterPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPlayer.Merge(m, src)
}
func (m *RegisterPlayer) XXX_Size() int {
	return xxx_messageInfo_RegisterPlayer.Size(m)
}
func (m *RegisterPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPlayer proto.InternalMessageInfo

func (m *RegisterPlayer) GetCid() int64 {
	if m != nil {
		return m.Cid
	}
	return 0
}

type Vec2 struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vec2) Reset()         { *m = Vec2{} }
func (m *Vec2) String() string { return proto.CompactTextString(m) }
func (*Vec2) ProtoMessage()    {}
func (*Vec2) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8bea4e98c5fae7, []int{14}
}

func (m *Vec2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vec2.Unmarshal(m, b)
}
func (m *Vec2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vec2.Marshal(b, m, deterministic)
}
func (m *Vec2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vec2.Merge(m, src)
}
func (m *Vec2) XXX_Size() int {
	return xxx_messageInfo_Vec2.Size(m)
}
func (m *Vec2) XXX_DiscardUnknown() {
	xxx_messageInfo_Vec2.DiscardUnknown(m)
}

var xxx_messageInfo_Vec2 proto.InternalMessageInfo

func (m *Vec2) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vec2) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientInitialize)(nil), "spaceagon.ClientInitialize")
	proto.RegisterType((*Memos)(nil), "spaceagon.Memos")
	proto.RegisterType((*Memo)(nil), "spaceagon.Memo")
	proto.RegisterType((*PosTracks)(nil), "spaceagon.PosTracks")
	proto.RegisterType((*MomentumTracks)(nil), "spaceagon.MomentumTracks")
	proto.RegisterType((*RotTracks)(nil), "spaceagon.RotTracks")
	proto.RegisterType((*SpinTracks)(nil), "spaceagon.SpinTracks")
	proto.RegisterType((*ShipControlTrack)(nil), "spaceagon.ShipControlTrack")
	proto.RegisterType((*DestroyEvent)(nil), "spaceagon.DestroyEvent")
	proto.RegisterType((*ShootMissile)(nil), "spaceagon.ShootMissile")
	proto.RegisterType((*SpawnMissile)(nil), "spaceagon.SpawnMissile")
	proto.RegisterType((*SpawnExplosion)(nil), "spaceagon.SpawnExplosion")
	proto.RegisterType((*SpawnShip)(nil), "spaceagon.SpawnShip")
	proto.RegisterType((*RegisterPlayer)(nil), "spaceagon.RegisterPlayer")
	proto.RegisterType((*Vec2)(nil), "spaceagon.vec2")
}

func init() { proto.RegisterFile("game/pb/messages.proto", fileDescriptor_ae8bea4e98c5fae7) }

var fileDescriptor_ae8bea4e98c5fae7 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xc1, 0x6e, 0xe3, 0x36,
	0x10, 0x8d, 0x24, 0xdb, 0xb0, 0xc7, 0x8e, 0xe3, 0xb2, 0x49, 0xab, 0xa2, 0x39, 0xb8, 0x6a, 0x5a,
	0x18, 0x48, 0x6b, 0x03, 0x29, 0x0a, 0xb4, 0x97, 0x1c, 0x9c, 0x04, 0x70, 0x51, 0x04, 0x08, 0x98,
	0x9e, 0x7a, 0xa8, 0x2b, 0x2b, 0x5c, 0x9b, 0x58, 0x49, 0x24, 0x48, 0x2a, 0x89, 0xf7, 0x9b, 0xf6,
	0xb0, 0xfb, 0x87, 0x0b, 0x92, 0xa6, 0x2c, 0x47, 0x06, 0x36, 0x7b, 0xe3, 0x9b, 0x99, 0xf7, 0x66,
	0xc4, 0x99, 0xa1, 0xe0, 0x9b, 0x65, 0x9c, 0x91, 0x09, 0x5f, 0x4c, 0x32, 0x22, 0x65, 0xbc, 0x24,
	0x72, 0xcc, 0x05, 0x53, 0x0c, 0x75, 0x24, 0x8f, 0x13, 0x12, 0x2f, 0x59, 0x1e, 0x9d, 0xc1, 0xe0,
	0x2a, 0xa5, 0x24, 0x57, 0x7f, 0xe5, 0x54, 0xd1, 0x38, 0xa5, 0xef, 0x08, 0x1a, 0x40, 0x90, 0xd0,
	0x87, 0xd0, 0x1b, 0x7a, 0xa3, 0x00, 0xeb, 0x63, 0x34, 0x86, 0xe6, 0x2d, 0xc9, 0x98, 0x44, 0x3f,
	0x41, 0x33, 0xd3, 0x87, 0xd0, 0x1b, 0x06, 0xa3, 0xee, 0xc5, 0xd1, 0xb8, 0x54, 0x1a, 0xeb, 0x00,
	0x6c, 0xbd, 0xd1, 0x87, 0x16, 0x34, 0x34, 0x46, 0x03, 0xf0, 0x15, 0xb3, 0x4a, 0xb3, 0x03, 0xec,
	0x2b, 0x86, 0x7e, 0x84, 0x1e, 0x79, 0x24, 0x62, 0xcd, 0x72, 0x32, 0x5f, 0x14, 0x2a, 0xf4, 0x37,
	0xbe, 0xae, 0xb3, 0x4e, 0x0b, 0x85, 0x4e, 0xa1, 0xed, 0x60, 0x18, 0x0c, 0xbd, 0x51, 0x7b, 0x76,
	0x80, 0x4b, 0x0b, 0xfa, 0x1d, 0x80, 0x33, 0x39, 0x57, 0x22, 0x4e, 0xde, 0xca, 0x10, 0x86, 0xde,
	0xa8, 0x7b, 0x71, 0x5c, 0xa9, 0xe4, 0x8e, 0xc9, 0x7f, 0x8c, 0x6f, 0xe6, 0xe1, 0x0e, 0x77, 0x00,
	0x5d, 0xc3, 0x51, 0xc6, 0x32, 0x92, 0xab, 0x22, 0x73, 0xdc, 0xae, 0xe1, 0x7e, 0x57, 0xfd, 0x8a,
	0x4d, 0x44, 0x29, 0xd0, 0xcf, 0x76, 0x2c, 0x3a, 0xb9, 0x60, 0xca, 0x09, 0xf4, 0x6a, 0xc9, 0x31,
	0x53, 0xdb, 0xe4, 0xc2, 0x01, 0xf4, 0x07, 0x74, 0x25, 0xa7, 0xb9, 0xe3, 0x1d, 0x1a, 0xde, 0x49,
	0x85, 0x77, 0xcf, 0x69, 0x5e, 0x12, 0x41, 0x96, 0x08, 0xfd, 0x0d, 0x48, 0xae, 0x28, 0x9f, 0x27,
	0x2c, 0x57, 0x82, 0xa5, 0x56, 0x21, 0xec, 0x1b, 0x81, 0xef, 0xab, 0x02, 0x2b, 0xca, 0xaf, 0x6c,
	0x8c, 0x61, 0xce, 0x3c, 0x3c, 0x90, 0x2f, 0x6c, 0xe8, 0x12, 0x0e, 0x1f, 0x88, 0x54, 0x82, 0xad,
	0xe7, 0xe4, 0x91, 0xe4, 0x2a, 0x1c, 0x18, 0x9d, 0x6f, 0x2b, 0x3a, 0xd7, 0xd6, 0x7f, 0xa3, 0xdd,
	0x33, 0x0f, 0xf7, 0x1e, 0x2a, 0x58, 0xf3, 0xe5, 0x8a, 0x31, 0x35, 0xcf, 0xa8, 0x94, 0x34, 0x25,
	0xe1, 0x57, 0x35, 0xfe, 0xbd, 0xf6, 0xdf, 0x5a, 0xb7, 0xe6, 0xcb, 0x0a, 0x36, 0x7c, 0x1e, 0x3f,
	0xe5, 0x25, 0x1f, 0xd5, 0xf9, 0xda, 0x5f, 0xe5, 0x57, 0xb0, 0xee, 0xa1, 0xe5, 0x93, 0x67, 0x9e,
	0x32, 0x49, 0x59, 0x1e, 0x7e, 0x5d, 0xeb, 0xa1, 0x51, 0xb8, 0x71, 0x01, 0xba, 0x87, 0x72, 0xc7,
	0xa2, 0x7b, 0x68, 0x55, 0xf4, 0xfd, 0x84, 0xc7, 0xb5, 0x1e, 0x1a, 0x01, 0x7d, 0x9f, 0xba, 0x87,
	0xd2, 0x01, 0x9d, 0x5c, 0x90, 0x25, 0x95, 0x8a, 0x88, 0x39, 0x4f, 0xe3, 0x35, 0x11, 0xe1, 0x49,
	0x2d, 0x39, 0xde, 0x44, 0xdc, 0x99, 0x00, 0x9d, 0x5c, 0xec, 0x58, 0xa6, 0x5d, 0xe8, 0x08, 0x92,
	0x50, 0xae, 0x97, 0x6e, 0xda, 0x86, 0x56, 0x9c, 0xa8, 0x22, 0x4e, 0xa3, 0x3f, 0xa1, 0x53, 0xce,
	0xad, 0xde, 0xc0, 0xdc, 0x6c, 0x60, 0x30, 0x6a, 0x60, 0x7d, 0x44, 0x3d, 0xf0, 0x9e, 0x43, 0x7f,
	0x18, 0x8c, 0x7c, 0xec, 0x3d, 0x6b, 0xb4, 0x0e, 0x03, 0x8b, 0xd6, 0xd1, 0x25, 0xf4, 0x77, 0xc7,
	0xf6, 0x0b, 0xf9, 0xe7, 0xd0, 0x29, 0xa7, 0x76, 0x3f, 0x55, 0x38, 0xaa, 0x88, 0x7e, 0x01, 0xd8,
	0x8e, 0xea, 0xfe, 0x68, 0xe9, 0xa2, 0x65, 0xf4, 0x1f, 0x0c, 0x5e, 0xce, 0xe5, 0x96, 0xe3, 0x39,
	0x4e, 0x1f, 0xfc, 0x82, 0x9b, 0x97, 0xa0, 0x8d, 0xfd, 0x82, 0x23, 0x04, 0x8d, 0x94, 0xbc, 0x51,
	0x76, 0xf5, 0xb1, 0x39, 0xa3, 0x63, 0x68, 0x0a, 0xba, 0x5c, 0xa9, 0xb0, 0x61, 0x8c, 0x16, 0x44,
	0x43, 0xe8, 0x55, 0xe7, 0xb5, 0xae, 0x1d, 0x9d, 0x41, 0xaf, 0x3a, 0x91, 0x5a, 0x87, 0x3d, 0xe5,
	0x44, 0x6c, 0x62, 0x2c, 0x88, 0xde, 0x7b, 0xd0, 0xab, 0x0e, 0x9e, 0x13, 0x6a, 0x6d, 0x8b, 0xdc,
	0x4b, 0x44, 0x3f, 0x40, 0xc0, 0x99, 0x34, 0xb5, 0xef, 0x3e, 0x87, 0x8f, 0x24, 0xb9, 0xc0, 0xda,
	0x87, 0xce, 0xa1, 0xed, 0xde, 0x10, 0xf3, 0x45, 0x7b, 0xe2, 0xca, 0x00, 0x9d, 0x57, 0x30, 0xfb,
	0x91, 0x3e, 0xd6, 0x47, 0x7d, 0x19, 0xfa, 0x35, 0x08, 0x9b, 0xc6, 0x64, 0xce, 0xd1, 0xff, 0xd0,
	0xdf, 0x1d, 0x72, 0x57, 0x87, 0xf7, 0xca, 0x3a, 0xfc, 0xcf, 0xd4, 0x11, 0x7d, 0xf4, 0xa0, 0x53,
	0xae, 0xc1, 0x9e, 0x96, 0x9d, 0x42, 0x27, 0x2e, 0xd4, 0x8a, 0x09, 0xaa, 0xd6, 0xf6, 0x0d, 0xc7,
	0x5b, 0x83, 0xab, 0x26, 0x78, 0x65, 0x35, 0x8d, 0x57, 0xde, 0x4a, 0xb3, 0x7e, 0x2b, 0xad, 0xca,
	0xad, 0x44, 0xd0, 0xdf, 0xdd, 0xbe, 0x3d, 0x7f, 0xb2, 0x08, 0x1a, 0x5a, 0xdb, 0xee, 0x83, 0x67,
	0xc8, 0x6e, 0x1f, 0x7c, 0x8b, 0xd6, 0xd3, 0xd1, 0xbf, 0x3f, 0x2f, 0xa9, 0x5a, 0x15, 0x8b, 0x71,
	0xc2, 0xb2, 0x49, 0x1a, 0x0b, 0x92, 0x11, 0x41, 0x26, 0xa6, 0xba, 0x5f, 0x75, 0x79, 0x93, 0xcd,
	0x7f, 0x75, 0xd1, 0x32, 0xff, 0xd3, 0xdf, 0x3e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x41, 0x26, 0x30,
	0x2a, 0x69, 0x07, 0x00, 0x00,
}
