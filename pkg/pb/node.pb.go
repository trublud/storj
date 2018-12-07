// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_ADMIN   NodeType = 0
	NodeType_STORAGE NodeType = 1
)

var NodeType_name = map[int32]string{
	0: "ADMIN",
	1: "STORAGE",
}
var NodeType_value = map[string]int32{
	"ADMIN":   0,
	"STORAGE": 1,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{0}
}

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP_TLS_GRPC NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP_TLS_GRPC",
}
var NodeTransport_value = map[string]int32{
	"TCP_TLS_GRPC": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{1}
}

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=free_bandwidth,json=freeBandwidth,proto3" json:"free_bandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=free_disk,json=freeDisk,proto3" json:"free_disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{0}
}
func (m *NodeRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRestrictions.Unmarshal(m, b)
}
func (m *NodeRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRestrictions.Marshal(b, m, deterministic)
}
func (dst *NodeRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRestrictions.Merge(dst, src)
}
func (m *NodeRestrictions) XXX_Size() int {
	return xxx_messageInfo_NodeRestrictions.Size(m)
}
func (m *NodeRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRestrictions proto.InternalMessageInfo

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// TODO move statdb.Update() stuff out of here
// Node represents a node in the overlay network
// Node is info for a updating a single storagenode, used in the Update rpc calls
type Node struct {
	Id                   NodeID            `protobuf:"bytes,1,opt,name=id,proto3,customtype=NodeID" json:"id"`
	Address              *NodeAddress      `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type                 NodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=node.NodeType" json:"type,omitempty"`
	Restrictions         *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions" json:"restrictions,omitempty"`
	Reputation           *NodeStats        `protobuf:"bytes,5,opt,name=reputation" json:"reputation,omitempty"`
	Metadata             *NodeMetadata     `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	LatencyList          []int64           `protobuf:"varint,7,rep,packed,name=latency_list,json=latencyList" json:"latency_list,omitempty"`
	AuditSuccess         bool              `protobuf:"varint,8,opt,name=audit_success,json=auditSuccess,proto3" json:"audit_success,omitempty"`
	IsUp                 bool              `protobuf:"varint,9,opt,name=is_up,json=isUp,proto3" json:"is_up,omitempty"`
	UpdateLatency        bool              `protobuf:"varint,10,opt,name=update_latency,json=updateLatency,proto3" json:"update_latency,omitempty"`
	UpdateAuditSuccess   bool              `protobuf:"varint,11,opt,name=update_audit_success,json=updateAuditSuccess,proto3" json:"update_audit_success,omitempty"`
	UpdateUptime         bool              `protobuf:"varint,12,opt,name=update_uptime,json=updateUptime,proto3" json:"update_uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_ADMIN
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

func (m *Node) GetReputation() *NodeStats {
	if m != nil {
		return m.Reputation
	}
	return nil
}

func (m *Node) GetMetadata() *NodeMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetLatencyList() []int64 {
	if m != nil {
		return m.LatencyList
	}
	return nil
}

func (m *Node) GetAuditSuccess() bool {
	if m != nil {
		return m.AuditSuccess
	}
	return false
}

func (m *Node) GetIsUp() bool {
	if m != nil {
		return m.IsUp
	}
	return false
}

func (m *Node) GetUpdateLatency() bool {
	if m != nil {
		return m.UpdateLatency
	}
	return false
}

func (m *Node) GetUpdateAuditSuccess() bool {
	if m != nil {
		return m.UpdateAuditSuccess
	}
	return false
}

func (m *Node) GetUpdateUptime() bool {
	if m != nil {
		return m.UpdateUptime
	}
	return false
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=node.NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{2}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (dst *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(dst, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_GRPC
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// NodeStats is the reputation characteristics of a node
type NodeStats struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	Latency_90           int64    `protobuf:"varint,2,opt,name=latency_90,json=latency90,proto3" json:"latency_90,omitempty"`
	AuditSuccessRatio    float64  `protobuf:"fixed64,3,opt,name=audit_success_ratio,json=auditSuccessRatio,proto3" json:"audit_success_ratio,omitempty"`
	UptimeRatio          float64  `protobuf:"fixed64,4,opt,name=uptime_ratio,json=uptimeRatio,proto3" json:"uptime_ratio,omitempty"`
	AuditCount           int64    `protobuf:"varint,5,opt,name=audit_count,json=auditCount,proto3" json:"audit_count,omitempty"`
	AuditSuccessCount    int64    `protobuf:"varint,6,opt,name=audit_success_count,json=auditSuccessCount,proto3" json:"audit_success_count,omitempty"`
	UptimeCount          int64    `protobuf:"varint,7,opt,name=uptime_count,json=uptimeCount,proto3" json:"uptime_count,omitempty"`
	UptimeSuccessCount   int64    `protobuf:"varint,8,opt,name=uptime_success_count,json=uptimeSuccessCount,proto3" json:"uptime_success_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{3}
}
func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (dst *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(dst, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetLatency_90() int64 {
	if m != nil {
		return m.Latency_90
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessRatio() float64 {
	if m != nil {
		return m.AuditSuccessRatio
	}
	return 0
}

func (m *NodeStats) GetUptimeRatio() float64 {
	if m != nil {
		return m.UptimeRatio
	}
	return 0
}

func (m *NodeStats) GetAuditCount() int64 {
	if m != nil {
		return m.AuditCount
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessCount() int64 {
	if m != nil {
		return m.AuditSuccessCount
	}
	return 0
}

func (m *NodeStats) GetUptimeCount() int64 {
	if m != nil {
		return m.UptimeCount
	}
	return 0
}

func (m *NodeStats) GetUptimeSuccessCount() int64 {
	if m != nil {
		return m.UptimeSuccessCount
	}
	return 0
}

type NodeMetadata struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Wallet               string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_9c604679ec4520fa, []int{4}
}
func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (dst *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(dst, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NodeMetadata) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeRestrictions)(nil), "node.NodeRestrictions")
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodeAddress)(nil), "node.NodeAddress")
	proto.RegisterType((*NodeStats)(nil), "node.NodeStats")
	proto.RegisterType((*NodeMetadata)(nil), "node.NodeMetadata")
	proto.RegisterEnum("node.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("node.NodeTransport", NodeTransport_name, NodeTransport_value)
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_9c604679ec4520fa) }

var fileDescriptor_node_9c604679ec4520fa = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x49, 0xe2, 0xfc, 0xf1, 0xb1, 0x93, 0x1b, 0x0e, 0x08, 0x59, 0xf7, 0xea, 0x5e, 0x82,
	0xd1, 0x55, 0x23, 0x2a, 0xa5, 0x94, 0xae, 0xa8, 0xba, 0x09, 0x50, 0x21, 0x24, 0xa0, 0x68, 0x12,
	0xba, 0x60, 0x63, 0x99, 0xcc, 0x94, 0x8e, 0x08, 0xb1, 0xe5, 0x19, 0x0b, 0xf1, 0x86, 0x5d, 0xf4,
	0x09, 0xba, 0xe0, 0x15, 0xfa, 0x0a, 0xd5, 0x9c, 0x71, 0x88, 0xad, 0xaa, 0xbb, 0xcc, 0xf7, 0xfd,
	0x7c, 0x8e, 0xe7, 0x7c, 0x27, 0x06, 0x58, 0x24, 0x5c, 0x8c, 0xd2, 0x2c, 0xd1, 0x09, 0x3a, 0xe6,
	0xf7, 0xdf, 0x70, 0x97, 0xdc, 0x25, 0x56, 0x09, 0x3f, 0x43, 0xff, 0x32, 0xe1, 0x82, 0x09, 0xa5,
	0x33, 0x39, 0xd3, 0x32, 0x59, 0x28, 0xfc, 0x1f, 0x7a, 0x5f, 0x32, 0x21, 0xa2, 0xdb, 0x78, 0xc1,
	0x1f, 0x25, 0xd7, 0x5f, 0x83, 0xda, 0xa0, 0x36, 0x6c, 0xb0, 0xae, 0x51, 0x8f, 0x96, 0x22, 0xfe,
	0x03, 0x2e, 0x61, 0x5c, 0xaa, 0xfb, 0xa0, 0x4e, 0x44, 0xc7, 0x08, 0x27, 0x52, 0xdd, 0x87, 0x3f,
	0x1b, 0xe0, 0x98, 0xc2, 0xf8, 0x1f, 0xd4, 0x25, 0xa7, 0x02, 0xfe, 0x51, 0xef, 0xdb, 0xf3, 0xf6,
	0xda, 0x8f, 0xe7, 0xed, 0x96, 0x71, 0xce, 0x4e, 0x58, 0x5d, 0x72, 0x7c, 0x0d, 0xed, 0x98, 0xf3,
	0x4c, 0x28, 0x45, 0x35, 0xbc, 0x83, 0xf5, 0x11, 0xbd, 0xb0, 0x41, 0xc6, 0xd6, 0x60, 0x4b, 0x02,
	0x43, 0x70, 0xf4, 0x53, 0x2a, 0x82, 0xc6, 0xa0, 0x36, 0xec, 0x1d, 0xf4, 0x56, 0xe4, 0xf4, 0x29,
	0x15, 0x8c, 0x3c, 0x7c, 0x0f, 0x7e, 0x56, 0xba, 0x4d, 0xe0, 0x50, 0xd5, 0xad, 0x15, 0x5b, 0xbe,
	0x2b, 0xab, 0xb0, 0xf8, 0x06, 0x20, 0x13, 0x69, 0xae, 0x63, 0x73, 0x0c, 0x9a, 0xf4, 0xe4, 0x5f,
	0xab, 0x27, 0x27, 0x3a, 0xd6, 0x8a, 0x95, 0x10, 0x1c, 0x41, 0xe7, 0x41, 0xe8, 0x98, 0xc7, 0x3a,
	0x0e, 0x5a, 0x84, 0xe3, 0x0a, 0xbf, 0x28, 0x1c, 0xf6, 0xc2, 0xe0, 0x0e, 0xf8, 0xf3, 0x58, 0x8b,
	0xc5, 0xec, 0x29, 0x9a, 0x4b, 0xa5, 0x83, 0xf6, 0xa0, 0x31, 0x6c, 0x30, 0xaf, 0xd0, 0xce, 0xa5,
	0xd2, 0xb8, 0x0b, 0xdd, 0x38, 0xe7, 0x52, 0x47, 0x2a, 0x9f, 0xcd, 0xcc, 0x58, 0x3a, 0x83, 0xda,
	0xb0, 0xc3, 0x7c, 0x12, 0x27, 0x56, 0xc3, 0x0d, 0x68, 0x4a, 0x15, 0xe5, 0x69, 0xe0, 0x92, 0xe9,
	0x48, 0x75, 0x9d, 0x9a, 0xdc, 0xf2, 0x94, 0xc7, 0x5a, 0x44, 0x45, 0xbd, 0x00, 0xc8, 0xed, 0x5a,
	0xf5, 0xdc, 0x8a, 0xb8, 0x0f, 0x9b, 0x05, 0x56, 0xed, 0xe3, 0x11, 0x8c, 0xd6, 0x1b, 0x97, 0xbb,
	0xed, 0x42, 0x51, 0x22, 0xca, 0x53, 0x2d, 0x1f, 0x44, 0xe0, 0xdb, 0x57, 0xb2, 0xe2, 0x35, 0x69,
	0xe1, 0x0d, 0x78, 0xa5, 0xcc, 0xf0, 0x2d, 0xb8, 0x3a, 0x8b, 0x17, 0x2a, 0x4d, 0x32, 0x4d, 0xf1,
	0xf7, 0x0e, 0x36, 0x4a, 0x79, 0x2d, 0x2d, 0xb6, 0xa2, 0x30, 0xa8, 0xae, 0x82, 0xfb, 0x92, 0x7b,
	0xf8, 0xbd, 0x0e, 0xee, 0x4b, 0x00, 0xf8, 0x0a, 0xda, 0xa6, 0x50, 0xf4, 0xc7, 0xbd, 0x6a, 0x19,
	0xfb, 0x8c, 0xe3, 0xbf, 0x00, 0xcb, 0x69, 0x1f, 0xee, 0x17, 0x2b, 0xea, 0x16, 0xca, 0xe1, 0x3e,
	0x8e, 0x60, 0xa3, 0x32, 0x81, 0x28, 0x33, 0xa1, 0xd2, 0x72, 0xd5, 0xd8, 0x7a, 0x79, 0xde, 0xcc,
	0x18, 0x26, 0x3c, 0x7b, 0xff, 0x02, 0x74, 0x08, 0xf4, 0xac, 0x66, 0x91, 0x6d, 0xf0, 0x6c, 0xc9,
	0x59, 0x92, 0x2f, 0x34, 0x6d, 0x50, 0x83, 0x01, 0x49, 0xc7, 0x46, 0xf9, 0xbd, 0xa7, 0x05, 0x5b,
	0x04, 0x56, 0x7a, 0x5a, 0x7e, 0xd5, 0xd3, 0x82, 0x6d, 0x02, 0x8b, 0x9e, 0x16, 0xa1, 0x3c, 0x09,
	0xa9, 0xd6, 0xec, 0x10, 0x8a, 0xd6, 0x2b, 0x17, 0x0d, 0x3f, 0x80, 0x5f, 0xde, 0x4f, 0xdc, 0x84,
	0xa6, 0x78, 0x88, 0xe5, 0x9c, 0xc6, 0xe9, 0x32, 0x7b, 0xc0, 0x2d, 0x68, 0x3d, 0xc6, 0xf3, 0xb9,
	0xd0, 0x45, 0x1a, 0xc5, 0x69, 0x2f, 0x84, 0xce, 0xf2, 0x2f, 0x87, 0x2e, 0x34, 0xc7, 0x27, 0x17,
	0x67, 0x97, 0xfd, 0x35, 0xf4, 0xa0, 0x3d, 0x99, 0x7e, 0x62, 0xe3, 0xd3, 0x8f, 0xfd, 0xda, 0xde,
	0x0e, 0x74, 0x2b, 0x31, 0x63, 0x1f, 0xfc, 0xe9, 0xf1, 0x55, 0x34, 0x3d, 0x9f, 0x44, 0xa7, 0xec,
	0xea, 0xb8, 0xbf, 0x76, 0xe4, 0xdc, 0xd4, 0xd3, 0xdb, 0xdb, 0x16, 0x7d, 0x86, 0xde, 0xfd, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0x35, 0xd5, 0x01, 0xa6, 0x04, 0x00, 0x00,
}