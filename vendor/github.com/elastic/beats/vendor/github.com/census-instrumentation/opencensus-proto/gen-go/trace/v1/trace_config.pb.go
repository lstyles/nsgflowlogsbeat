// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/trace/v1/trace_config.proto

package v1

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

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

var ConstantSampler_ConstantDecision_name = map[int32]string{
	0: "ALWAYS_OFF",
	1: "ALWAYS_ON",
	2: "ALWAYS_PARENT",
}

var ConstantSampler_ConstantDecision_value = map[string]int32{
	"ALWAYS_OFF":    0,
	"ALWAYS_ON":     1,
	"ALWAYS_PARENT": 2,
}

func (x ConstantSampler_ConstantDecision) String() string {
	return proto.EnumName(ConstantSampler_ConstantDecision_name, int32(x))
}

func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{2, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ProbabilitySampler
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfAnnotations int64 `protobuf:"varint,5,opt,name=max_number_of_annotations,json=maxNumberOfAnnotations,proto3" json:"max_number_of_annotations,omitempty"`
	// The global default max number of message events per span.
	MaxNumberOfMessageEvents int64 `protobuf:"varint,6,opt,name=max_number_of_message_events,json=maxNumberOfMessageEvents,proto3" json:"max_number_of_message_events,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks     int64    `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{0}
}

func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceConfig.Unmarshal(m, b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceConfig.Size(m)
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
}

type TraceConfig_ProbabilitySampler struct {
	ProbabilitySampler *ProbabilitySampler `protobuf:"bytes,1,opt,name=probability_sampler,json=probabilitySampler,proto3,oneof"`
}

type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,2,opt,name=constant_sampler,json=constantSampler,proto3,oneof"`
}

type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof"`
}

func (*TraceConfig_ProbabilitySampler) isTraceConfig_Sampler() {}

func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler() {}

func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetProbabilitySampler() *ProbabilitySampler {
	if x, ok := m.GetSampler().(*TraceConfig_ProbabilitySampler); ok {
		return x.ProbabilitySampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAnnotations() int64 {
	if m != nil {
		return m.MaxNumberOfAnnotations
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfMessageEvents() int64 {
	if m != nil {
		return m.MaxNumberOfMessageEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TraceConfig_ProbabilitySampler)(nil),
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

// Sampler that tries to uniformly sample traces with a given probability.
// The probability of sampling a trace is equal to that of the specified probability.
type ProbabilitySampler struct {
	// The desired probability of sampling. Must be within [0.0, 1.0].
	SamplingProbability  float64  `protobuf:"fixed64,1,opt,name=samplingProbability,proto3" json:"samplingProbability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbabilitySampler) Reset()         { *m = ProbabilitySampler{} }
func (m *ProbabilitySampler) String() string { return proto.CompactTextString(m) }
func (*ProbabilitySampler) ProtoMessage()    {}
func (*ProbabilitySampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{1}
}

func (m *ProbabilitySampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbabilitySampler.Unmarshal(m, b)
}
func (m *ProbabilitySampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbabilitySampler.Marshal(b, m, deterministic)
}
func (m *ProbabilitySampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbabilitySampler.Merge(m, src)
}
func (m *ProbabilitySampler) XXX_Size() int {
	return xxx_messageInfo_ProbabilitySampler.Size(m)
}
func (m *ProbabilitySampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbabilitySampler.DiscardUnknown(m)
}

var xxx_messageInfo_ProbabilitySampler proto.InternalMessageInfo

func (m *ProbabilitySampler) GetSamplingProbability() float64 {
	if m != nil {
		return m.SamplingProbability
	}
	return 0
}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	Decision             ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opencensus.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{2}
}

func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConstantSampler.Unmarshal(m, b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
}
func (m *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(m, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return xxx_messageInfo_ConstantSampler.Size(m)
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if m != nil {
		return m.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps                  int64    `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{3}
}

func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitingSampler.Unmarshal(m, b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
}
func (m *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(m, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return xxx_messageInfo_RateLimitingSampler.Size(m)
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterEnum("opencensus.proto.trace.v1.ConstantSampler_ConstantDecision", ConstantSampler_ConstantDecision_name, ConstantSampler_ConstantDecision_value)
	proto.RegisterType((*TraceConfig)(nil), "opencensus.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ProbabilitySampler)(nil), "opencensus.proto.trace.v1.ProbabilitySampler")
	proto.RegisterType((*ConstantSampler)(nil), "opencensus.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*RateLimitingSampler)(nil), "opencensus.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opencensus/proto/trace/v1/trace_config.proto", fileDescriptor_5359209b41ff50c5)
}

var fileDescriptor_5359209b41ff50c5 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x97, 0x76, 0x6c, 0xec, 0x9b, 0xb6, 0x05, 0x57, 0x43, 0xa9, 0xb4, 0xc3, 0x94, 0x0b,
	0x13, 0x22, 0x09, 0x1d, 0x07, 0x84, 0x90, 0x90, 0xda, 0x6e, 0x15, 0x87, 0xd2, 0x56, 0xd9, 0x44,
	0x05, 0x97, 0xe0, 0x64, 0x6e, 0xb0, 0x68, 0xec, 0x60, 0x3b, 0xd5, 0x78, 0x0d, 0xce, 0x3c, 0x04,
	0xcf, 0xc5, 0x53, 0xa0, 0x3a, 0x21, 0x49, 0xdb, 0x6d, 0xe2, 0x96, 0xef, 0xfb, 0x7f, 0xbf, 0x9f,
	0xad, 0xd8, 0x86, 0x17, 0x3c, 0x25, 0x2c, 0x22, 0x4c, 0x66, 0xd2, 0x4b, 0x05, 0x57, 0xdc, 0x53,
	0x02, 0x47, 0xc4, 0x5b, 0x74, 0xf2, 0x8f, 0x20, 0xe2, 0x6c, 0x46, 0x63, 0x57, 0x67, 0xa8, 0x5d,
	0x4d, 0xe7, 0x1d, 0x57, 0x0f, 0xb9, 0x8b, 0x8e, 0xfd, 0x6b, 0x1b, 0xf6, 0xaf, 0x97, 0x45, 0x5f,
	0x03, 0xe8, 0x0b, 0xb4, 0x52, 0xc1, 0x43, 0x1c, 0xd2, 0x39, 0x55, 0x3f, 0x02, 0x89, 0x93, 0x74,
	0x4e, 0x84, 0x65, 0x9c, 0x1a, 0x67, 0xfb, 0xe7, 0x8e, 0x7b, 0xaf, 0xc8, 0x9d, 0x54, 0xd4, 0x55,
	0x0e, 0xbd, 0xdf, 0xf2, 0x51, 0xba, 0xd1, 0x45, 0x53, 0x30, 0x23, 0xce, 0xa4, 0xc2, 0x4c, 0x95,
	0xfa, 0x86, 0xd6, 0x3f, 0x7f, 0x40, 0xdf, 0x2f, 0x90, 0xca, 0x7d, 0x14, 0xad, 0xb6, 0xd0, 0x0d,
	0x1c, 0x0b, 0xac, 0x48, 0x30, 0xa7, 0x09, 0x55, 0x94, 0xc5, 0xa5, 0xbd, 0xa9, 0xed, 0xee, 0x03,
	0x76, 0x1f, 0x2b, 0x32, 0x2c, 0xb0, 0x6a, 0x85, 0x96, 0xd8, 0x6c, 0xa3, 0xd7, 0x60, 0x25, 0xf8,
	0x36, 0x60, 0x59, 0x12, 0x12, 0x11, 0xf0, 0x59, 0x80, 0x95, 0x12, 0x34, 0xcc, 0x14, 0x91, 0xd6,
	0xf6, 0xa9, 0x71, 0xd6, 0xf4, 0x8f, 0x13, 0x7c, 0x3b, 0xd2, 0xf1, 0x78, 0xd6, 0x2d, 0x43, 0xf4,
	0x06, 0xda, 0x6b, 0x20, 0x63, 0x5c, 0x61, 0x45, 0x39, 0x93, 0xd6, 0x23, 0x4d, 0x3e, 0xad, 0x93,
	0x55, 0x8a, 0xde, 0xc1, 0xc9, 0x2a, 0x9a, 0x10, 0x29, 0x71, 0x4c, 0x02, 0xb2, 0x20, 0x4c, 0x49,
	0x6b, 0x47, 0xd3, 0x56, 0x8d, 0xfe, 0x90, 0x0f, 0x5c, 0xea, 0x1c, 0x39, 0xd0, 0x5a, 0xe5, 0xe7,
	0x94, 0x7d, 0x93, 0xd6, 0xae, 0xc6, 0xcc, 0x1a, 0x36, 0x5c, 0xf6, 0x7b, 0x7b, 0xb0, 0x5b, 0xfc,
	0x3a, 0x7b, 0x00, 0x68, 0xf3, 0x60, 0xd1, 0x4b, 0x68, 0xe9, 0x01, 0xca, 0xe2, 0x5a, 0xaa, 0x2f,
	0x89, 0xe1, 0xdf, 0x15, 0xd9, 0xbf, 0x0d, 0x38, 0x5a, 0x3b, 0x42, 0x34, 0x85, 0xc7, 0x37, 0x24,
	0xa2, 0x92, 0x72, 0xa6, 0xd1, 0xc3, 0xf3, 0xb7, 0xff, 0x7f, 0x01, 0xca, 0xfa, 0xa2, 0x50, 0xf8,
	0xa5, 0xcc, 0xbe, 0x00, 0x73, 0x3d, 0x45, 0x87, 0x00, 0xdd, 0xe1, 0xb4, 0xfb, 0xe9, 0x2a, 0x18,
	0x0f, 0x06, 0xe6, 0x16, 0x3a, 0x80, 0xbd, 0x7f, 0xf5, 0xc8, 0x34, 0xd0, 0x13, 0x38, 0x28, 0xca,
	0x49, 0xd7, 0xbf, 0x1c, 0x5d, 0x9b, 0x0d, 0xfb, 0x19, 0xb4, 0xee, 0xb8, 0x16, 0xc8, 0x84, 0xe6,
	0xf7, 0x54, 0xea, 0x0d, 0x37, 0xfd, 0xe5, 0x67, 0xef, 0xa7, 0x01, 0x27, 0x94, 0xdf, 0xbf, 0xf5,
	0x9e, 0x59, 0x7b, 0x60, 0x93, 0x65, 0x34, 0x31, 0x3e, 0xf7, 0x62, 0xaa, 0xbe, 0x66, 0xa1, 0x1b,
	0xf1, 0xc4, 0xcb, 0x29, 0x87, 0x32, 0xa9, 0x44, 0x96, 0x10, 0x96, 0x1f, 0xbb, 0x57, 0x09, 0x9d,
	0xfc, 0x89, 0xc7, 0x84, 0x39, 0x71, 0xf5, 0xd2, 0xff, 0x34, 0xda, 0xe3, 0x94, 0xb0, 0x7e, 0xbe,
	0xa6, 0x16, 0xbb, 0x7a, 0x25, 0xf7, 0x63, 0x27, 0xdc, 0xd1, 0xc8, 0xab, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x50, 0x0c, 0xfe, 0x32, 0x29, 0x04, 0x00, 0x00,
}
