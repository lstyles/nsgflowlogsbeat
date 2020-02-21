// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/agent/metrics/v1/metrics_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	v11 "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	v12 "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ExportMetricsServiceRequest struct {
	// This is required only in the first message on the stream or if the
	// previous sent ExportMetricsServiceRequest message has a different Node (e.g.
	// when the same RPC is used to send Metrics from multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of metrics that belong to the last received Node.
	Metrics []*v11.Metric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// The resource for the metrics in this message that do not have an explicit
	// resource set.
	// If unset, the most recently set resource in the RPC stream applies. It is
	// valid to never be set within a stream, e.g. when no resource info is known
	// at all or when all sent metrics have an explicit resource set.
	Resource             *v12.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportMetricsServiceRequest) Reset()         { *m = ExportMetricsServiceRequest{} }
func (m *ExportMetricsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceRequest) ProtoMessage()    {}
func (*ExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e253a956287d04, []int{0}
}

func (m *ExportMetricsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceRequest.Unmarshal(m, b)
}
func (m *ExportMetricsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceRequest.Merge(m, src)
}
func (m *ExportMetricsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceRequest.Size(m)
}
func (m *ExportMetricsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceRequest proto.InternalMessageInfo

func (m *ExportMetricsServiceRequest) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ExportMetricsServiceRequest) GetMetrics() []*v11.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ExportMetricsServiceRequest) GetResource() *v12.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ExportMetricsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportMetricsServiceResponse) Reset()         { *m = ExportMetricsServiceResponse{} }
func (m *ExportMetricsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceResponse) ProtoMessage()    {}
func (*ExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e253a956287d04, []int{1}
}

func (m *ExportMetricsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceResponse.Unmarshal(m, b)
}
func (m *ExportMetricsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceResponse.Merge(m, src)
}
func (m *ExportMetricsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceResponse.Size(m)
}
func (m *ExportMetricsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportMetricsServiceRequest)(nil), "opencensus.proto.agent.metrics.v1.ExportMetricsServiceRequest")
	proto.RegisterType((*ExportMetricsServiceResponse)(nil), "opencensus.proto.agent.metrics.v1.ExportMetricsServiceResponse")
}

func init() {
	proto.RegisterFile("opencensus/proto/agent/metrics/v1/metrics_service.proto", fileDescriptor_47e253a956287d04)
}

var fileDescriptor_47e253a956287d04 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x9d, 0x56, 0xaa, 0x4c, 0xc1, 0x45, 0xdc, 0x94, 0x2a, 0xd2, 0x56, 0x91, 0x8a, 0x64,
	0x62, 0xea, 0x42, 0x10, 0x54, 0xac, 0xb8, 0x11, 0xd4, 0x12, 0xc1, 0x85, 0x1b, 0x69, 0xd3, 0x47,
	0xcc, 0x22, 0x33, 0x71, 0x66, 0x12, 0xbc, 0x85, 0x77, 0x70, 0xef, 0x8d, 0x3c, 0x81, 0xa7, 0x90,
	0xe4, 0x4d, 0x5a, 0x4a, 0x8c, 0x05, 0x77, 0x8f, 0xe4, 0xff, 0xfe, 0xf7, 0xff, 0x33, 0x43, 0x4f,
	0x44, 0x0c, 0xdc, 0x07, 0xae, 0x12, 0xe5, 0xc4, 0x52, 0x68, 0xe1, 0x8c, 0x03, 0xe0, 0xda, 0x89,
	0x40, 0xcb, 0xd0, 0x57, 0x4e, 0xea, 0x16, 0xe3, 0xb3, 0x02, 0x99, 0x86, 0x3e, 0xb0, 0x5c, 0x66,
	0x75, 0xe7, 0x20, 0x7e, 0x61, 0x39, 0xc8, 0x8c, 0x9a, 0xa5, 0x6e, 0xdb, 0xae, 0xf0, 0xf6, 0x45,
	0x14, 0x09, 0x9e, 0x59, 0xe3, 0x84, 0x7c, 0xfb, 0xa0, 0x24, 0x2f, 0x87, 0x30, 0xd2, 0xc3, 0x92,
	0x54, 0x82, 0x12, 0x89, 0xf4, 0x21, 0xd3, 0x16, 0x33, 0x8a, 0x7b, 0x5f, 0x84, 0x6e, 0x5d, 0xbf,
	0xc5, 0x42, 0xea, 0x5b, 0x34, 0x79, 0xc0, 0x22, 0x1e, 0xbc, 0x26, 0xa0, 0xb4, 0x75, 0x4a, 0x57,
	0xb9, 0x98, 0x42, 0x8b, 0x74, 0x48, 0xbf, 0x39, 0xd8, 0x67, 0x15, 0xc5, 0x4c, 0xd6, 0xd4, 0x65,
	0x77, 0x62, 0x0a, 0x5e, 0xce, 0x58, 0x67, 0x74, 0xcd, 0x24, 0x6b, 0xd5, 0x3a, 0xf5, 0x7e, 0x73,
	0xb0, 0x5b, 0xc6, 0xe7, 0x27, 0xc2, 0x30, 0x80, 0x57, 0x30, 0xd6, 0x90, 0xae, 0x17, 0x61, 0x5b,
	0xf5, 0xaa, 0xf5, 0xb3, 0x3a, 0xa9, 0xcb, 0x3c, 0x33, 0x7b, 0x33, 0xae, 0xb7, 0x43, 0xb7, 0x7f,
	0x6f, 0xa7, 0x62, 0xc1, 0x15, 0x0c, 0x3e, 0x08, 0xdd, 0x58, 0xfc, 0x65, 0xbd, 0x13, 0xda, 0x40,
	0xc6, 0x3a, 0x67, 0x4b, 0xef, 0x91, 0xfd, 0x71, 0x78, 0xed, 0x8b, 0x7f, 0xf3, 0x18, 0xaf, 0xb7,
	0xd2, 0x27, 0x47, 0x64, 0xf8, 0x49, 0xe8, 0x5e, 0x28, 0x96, 0x7b, 0x0d, 0x37, 0x17, 0x6d, 0x46,
	0x99, 0x6a, 0x44, 0x9e, 0x6e, 0x82, 0x50, 0xbf, 0x24, 0x93, 0xec, 0x92, 0x1c, 0x34, 0xb0, 0x43,
	0xae, 0xb4, 0x4c, 0x22, 0xe0, 0x7a, 0xac, 0x43, 0xc1, 0x9d, 0xb9, 0xb7, 0x8d, 0x4f, 0x26, 0x00,
	0x6e, 0x07, 0xe5, 0xf7, 0xfe, 0x5d, 0xeb, 0xde, 0xc7, 0xc0, 0xaf, 0x30, 0x46, 0xbe, 0x80, 0x5d,
	0xe6, 0x31, 0xcc, 0x6a, 0xf6, 0xe8, 0x4e, 0x1a, 0xb9, 0xc5, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x19, 0x28, 0xa4, 0x50, 0x3f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ExportClient, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) Export(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetricsService_serviceDesc.Streams[0], "/opencensus.proto.agent.metrics.v1.MetricsService/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceExportClient{stream}
	return x, nil
}

type MetricsService_ExportClient interface {
	Send(*ExportMetricsServiceRequest) error
	Recv() (*ExportMetricsServiceResponse, error)
	grpc.ClientStream
}

type metricsServiceExportClient struct {
	grpc.ClientStream
}

func (x *metricsServiceExportClient) Send(m *ExportMetricsServiceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceExportClient) Recv() (*ExportMetricsServiceResponse, error) {
	m := new(ExportMetricsServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(MetricsService_ExportServer) error
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) Export(srv MetricsService_ExportServer) error {
	return status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).Export(&metricsServiceExportServer{stream})
}

type MetricsService_ExportServer interface {
	Send(*ExportMetricsServiceResponse) error
	Recv() (*ExportMetricsServiceRequest, error)
	grpc.ServerStream
}

type metricsServiceExportServer struct {
	grpc.ServerStream
}

func (x *metricsServiceExportServer) Send(m *ExportMetricsServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceExportServer) Recv() (*ExportMetricsServiceRequest, error) {
	m := new(ExportMetricsServiceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opencensus.proto.agent.metrics.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Export",
			Handler:       _MetricsService_Export_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "opencensus/proto/agent/metrics/v1/metrics_service.proto",
}
