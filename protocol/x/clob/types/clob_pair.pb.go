// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furyaprotocol/clob/clob_pair.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Status of the CLOB.
type ClobPair_Status int32

const (
	// Default value. This value is invalid and unused.
	ClobPair_STATUS_UNSPECIFIED ClobPair_Status = 0
	// STATUS_ACTIVE represents an active clob pair.
	ClobPair_STATUS_ACTIVE ClobPair_Status = 1
	// STATUS_PAUSED behavior is unfinalized.
	// TODO(DEC-600): update this documentation.
	ClobPair_STATUS_PAUSED ClobPair_Status = 2
	// STATUS_CANCEL_ONLY behavior is unfinalized.
	// TODO(DEC-600): update this documentation.
	ClobPair_STATUS_CANCEL_ONLY ClobPair_Status = 3
	// STATUS_POST_ONLY behavior is unfinalized.
	// TODO(DEC-600): update this documentation.
	ClobPair_STATUS_POST_ONLY ClobPair_Status = 4
	// STATUS_INITIALIZING represents a newly-added clob pair.
	// Clob pairs in this state only accept orders which are
	// both short-term and post-only.
	ClobPair_STATUS_INITIALIZING ClobPair_Status = 5
	// STATUS_FINAL_SETTLEMENT represents a clob pair which is deactivated
	// and trading has ceased. All open positions will be closed by the
	// protocol. Open stateful orders will be cancelled. Open short-term
	// orders will be left to expire.
	ClobPair_STATUS_FINAL_SETTLEMENT ClobPair_Status = 6
)

var ClobPair_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ACTIVE",
	2: "STATUS_PAUSED",
	3: "STATUS_CANCEL_ONLY",
	4: "STATUS_POST_ONLY",
	5: "STATUS_INITIALIZING",
	6: "STATUS_FINAL_SETTLEMENT",
}

var ClobPair_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED":      0,
	"STATUS_ACTIVE":           1,
	"STATUS_PAUSED":           2,
	"STATUS_CANCEL_ONLY":      3,
	"STATUS_POST_ONLY":        4,
	"STATUS_INITIALIZING":     5,
	"STATUS_FINAL_SETTLEMENT": 6,
}

func (x ClobPair_Status) String() string {
	return proto.EnumName(ClobPair_Status_name, int32(x))
}

func (ClobPair_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_178b475635886947, []int{2, 0}
}

// PerpetualClobMetadata contains metadata for a `ClobPair`
// representing a Perpetual product.
type PerpetualClobMetadata struct {
	// Id of the Perpetual the CLOB allows trading of.
	PerpetualId uint32 `protobuf:"varint,1,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
}

func (m *PerpetualClobMetadata) Reset()         { *m = PerpetualClobMetadata{} }
func (m *PerpetualClobMetadata) String() string { return proto.CompactTextString(m) }
func (*PerpetualClobMetadata) ProtoMessage()    {}
func (*PerpetualClobMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_178b475635886947, []int{0}
}
func (m *PerpetualClobMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpetualClobMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpetualClobMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpetualClobMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpetualClobMetadata.Merge(m, src)
}
func (m *PerpetualClobMetadata) XXX_Size() int {
	return m.Size()
}
func (m *PerpetualClobMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpetualClobMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PerpetualClobMetadata proto.InternalMessageInfo

func (m *PerpetualClobMetadata) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

// PerpetualClobMetadata contains metadata for a `ClobPair`
// representing a Spot product.
type SpotClobMetadata struct {
	// Id of the base Asset in the trading pair.
	BaseAssetId uint32 `protobuf:"varint,1,opt,name=base_asset_id,json=baseAssetId,proto3" json:"base_asset_id,omitempty"`
	// Id of the quote Asset in the trading pair.
	QuoteAssetId uint32 `protobuf:"varint,2,opt,name=quote_asset_id,json=quoteAssetId,proto3" json:"quote_asset_id,omitempty"`
}

func (m *SpotClobMetadata) Reset()         { *m = SpotClobMetadata{} }
func (m *SpotClobMetadata) String() string { return proto.CompactTextString(m) }
func (*SpotClobMetadata) ProtoMessage()    {}
func (*SpotClobMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_178b475635886947, []int{1}
}
func (m *SpotClobMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpotClobMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpotClobMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpotClobMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpotClobMetadata.Merge(m, src)
}
func (m *SpotClobMetadata) XXX_Size() int {
	return m.Size()
}
func (m *SpotClobMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SpotClobMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SpotClobMetadata proto.InternalMessageInfo

func (m *SpotClobMetadata) GetBaseAssetId() uint32 {
	if m != nil {
		return m.BaseAssetId
	}
	return 0
}

func (m *SpotClobMetadata) GetQuoteAssetId() uint32 {
	if m != nil {
		return m.QuoteAssetId
	}
	return 0
}

// ClobPair represents a single CLOB pair for a given product
// in state.
type ClobPair struct {
	// ID of the orderbook that stores all resting liquidity for this CLOB.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Product-specific metadata. Perpetual CLOBs will have
	// PerpetualClobMetadata, and Spot CLOBs will have SpotClobMetadata.
	//
	// Types that are valid to be assigned to Metadata:
	//
	//	*ClobPair_PerpetualClobMetadata
	//	*ClobPair_SpotClobMetadata
	Metadata isClobPair_Metadata `protobuf_oneof:"metadata"`
	// Minimum increment in the size of orders on the CLOB, in base quantums.
	StepBaseQuantums uint64 `protobuf:"varint,4,opt,name=step_base_quantums,json=stepBaseQuantums,proto3" json:"step_base_quantums,omitempty"`
	// Defines the tick size of the orderbook by defining how many subticks
	// are in one tick. That is, the subticks of any valid order must be a
	// multiple of this value. Generally this value should start `>= 100`to
	// allow room for decreasing it.
	SubticksPerTick uint32 `protobuf:"varint,5,opt,name=subticks_per_tick,json=subticksPerTick,proto3" json:"subticks_per_tick,omitempty"`
	// `10^Exponent` gives the number of QuoteQuantums traded per BaseQuantum
	// per Subtick.
	QuantumConversionExponent int32           `protobuf:"zigzag32,6,opt,name=quantum_conversion_exponent,json=quantumConversionExponent,proto3" json:"quantum_conversion_exponent,omitempty"`
	Status                    ClobPair_Status `protobuf:"varint,7,opt,name=status,proto3,enum=furyaprotocol.clob.ClobPair_Status" json:"status,omitempty"`
}

func (m *ClobPair) Reset()         { *m = ClobPair{} }
func (m *ClobPair) String() string { return proto.CompactTextString(m) }
func (*ClobPair) ProtoMessage()    {}
func (*ClobPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_178b475635886947, []int{2}
}
func (m *ClobPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClobPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClobPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClobPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClobPair.Merge(m, src)
}
func (m *ClobPair) XXX_Size() int {
	return m.Size()
}
func (m *ClobPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ClobPair.DiscardUnknown(m)
}

var xxx_messageInfo_ClobPair proto.InternalMessageInfo

type isClobPair_Metadata interface {
	isClobPair_Metadata()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ClobPair_PerpetualClobMetadata struct {
	PerpetualClobMetadata *PerpetualClobMetadata `protobuf:"bytes,2,opt,name=perpetual_clob_metadata,json=perpetualClobMetadata,proto3,oneof" json:"perpetual_clob_metadata,omitempty"`
}
type ClobPair_SpotClobMetadata struct {
	SpotClobMetadata *SpotClobMetadata `protobuf:"bytes,3,opt,name=spot_clob_metadata,json=spotClobMetadata,proto3,oneof" json:"spot_clob_metadata,omitempty"`
}

func (*ClobPair_PerpetualClobMetadata) isClobPair_Metadata() {}
func (*ClobPair_SpotClobMetadata) isClobPair_Metadata()      {}

func (m *ClobPair) GetMetadata() isClobPair_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ClobPair) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClobPair) GetPerpetualClobMetadata() *PerpetualClobMetadata {
	if x, ok := m.GetMetadata().(*ClobPair_PerpetualClobMetadata); ok {
		return x.PerpetualClobMetadata
	}
	return nil
}

func (m *ClobPair) GetSpotClobMetadata() *SpotClobMetadata {
	if x, ok := m.GetMetadata().(*ClobPair_SpotClobMetadata); ok {
		return x.SpotClobMetadata
	}
	return nil
}

func (m *ClobPair) GetStepBaseQuantums() uint64 {
	if m != nil {
		return m.StepBaseQuantums
	}
	return 0
}

func (m *ClobPair) GetSubticksPerTick() uint32 {
	if m != nil {
		return m.SubticksPerTick
	}
	return 0
}

func (m *ClobPair) GetQuantumConversionExponent() int32 {
	if m != nil {
		return m.QuantumConversionExponent
	}
	return 0
}

func (m *ClobPair) GetStatus() ClobPair_Status {
	if m != nil {
		return m.Status
	}
	return ClobPair_STATUS_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClobPair) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClobPair_PerpetualClobMetadata)(nil),
		(*ClobPair_SpotClobMetadata)(nil),
	}
}

func init() {
	proto.RegisterEnum("furyaprotocol.clob.ClobPair_Status", ClobPair_Status_name, ClobPair_Status_value)
	proto.RegisterType((*PerpetualClobMetadata)(nil), "furyaprotocol.clob.PerpetualClobMetadata")
	proto.RegisterType((*SpotClobMetadata)(nil), "furyaprotocol.clob.SpotClobMetadata")
	proto.RegisterType((*ClobPair)(nil), "furyaprotocol.clob.ClobPair")
}

func init() { proto.RegisterFile("furyaprotocol/clob/clob_pair.proto", fileDescriptor_178b475635886947) }

var fileDescriptor_178b475635886947 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x6a, 0xdb, 0x30,
	0x1c, 0xc6, 0xed, 0x34, 0xcd, 0x8a, 0xda, 0x64, 0x8e, 0xb6, 0x2e, 0x19, 0x05, 0x93, 0x7a, 0x3b,
	0x84, 0xb1, 0x39, 0xd0, 0x8d, 0x1d, 0x7a, 0x18, 0x38, 0xae, 0xbb, 0x1a, 0x52, 0xd7, 0xb3, 0x9d,
	0xc1, 0xca, 0xc0, 0xc8, 0xb6, 0x58, 0x4d, 0x13, 0xcb, 0xb5, 0xe4, 0x92, 0xbe, 0xc5, 0x9e, 0x63,
	0x4f, 0xb2, 0x63, 0x8f, 0x3b, 0x8e, 0xe4, 0x45, 0x86, 0x15, 0x27, 0x4b, 0xd3, 0x5c, 0x8c, 0xf9,
	0x7d, 0xdf, 0xff, 0x43, 0xff, 0x4f, 0x08, 0x1c, 0x46, 0x77, 0xd1, 0x24, 0xcd, 0x08, 0x23, 0x21,
	0x19, 0xf5, 0xc2, 0x11, 0x09, 0xf8, 0xc7, 0x4f, 0x51, 0x9c, 0xa9, 0x9c, 0xc3, 0xe6, 0xaa, 0x45,
	0x2d, 0x54, 0xe5, 0x18, 0xec, 0xdb, 0x38, 0x4b, 0x31, 0xcb, 0xd1, 0x48, 0x1f, 0x91, 0xe0, 0x1c,
	0x33, 0x14, 0x21, 0x86, 0xe0, 0x21, 0xd8, 0x4b, 0x17, 0x82, 0x1f, 0x47, 0x6d, 0xb1, 0x23, 0x76,
	0xeb, 0xce, 0xee, 0x92, 0x99, 0x91, 0xf2, 0x1d, 0x48, 0x6e, 0x4a, 0xd8, 0x83, 0x31, 0x05, 0xd4,
	0x03, 0x44, 0xb1, 0x8f, 0x28, 0xc5, 0x6c, 0x65, 0xae, 0x80, 0x5a, 0xc1, 0xcc, 0x08, 0xbe, 0x06,
	0x8d, 0x9b, 0x9c, 0xb0, 0x15, 0x53, 0x85, 0x9b, 0xf6, 0x38, 0x2d, 0x5d, 0xca, 0xb4, 0x0a, 0x76,
	0x8a, 0x68, 0x1b, 0xc5, 0x19, 0x6c, 0x80, 0xca, 0x32, 0xab, 0x12, 0x47, 0x30, 0x00, 0xad, 0xff,
	0xa7, 0xe3, 0x6b, 0x8e, 0xcb, 0x13, 0xf0, 0xac, 0xdd, 0xa3, 0xae, 0xfa, 0x68, 0x57, 0x75, 0xe3,
	0xa2, 0x67, 0x82, 0xb3, 0x9f, 0x6e, 0x6c, 0xc0, 0x05, 0x90, 0xa6, 0x84, 0xad, 0xc5, 0x6f, 0xf1,
	0xf8, 0x57, 0x1b, 0xe2, 0xd7, 0xbb, 0x38, 0x13, 0x1c, 0x89, 0xae, 0xf7, 0xf3, 0x16, 0x40, 0xca,
	0x70, 0xea, 0xf3, 0x92, 0x6e, 0x72, 0x94, 0xb0, 0x7c, 0x4c, 0xdb, 0xd5, 0x8e, 0xd8, 0xad, 0x3a,
	0x52, 0xa1, 0xf4, 0x11, 0xc5, 0x5f, 0x4a, 0x0e, 0xdf, 0x80, 0x26, 0xcd, 0x03, 0x16, 0x87, 0xd7,
	0xd4, 0x4f, 0x71, 0xe6, 0x17, 0x7f, 0xed, 0x6d, 0xde, 0xc2, 0xd3, 0x85, 0x60, 0xe3, 0xcc, 0x8b,
	0xc3, 0x6b, 0xf8, 0x09, 0x1c, 0x94, 0x79, 0x7e, 0x48, 0x92, 0x5b, 0x9c, 0xd1, 0x98, 0x24, 0x3e,
	0x9e, 0xa4, 0x24, 0xc1, 0x09, 0x6b, 0xd7, 0x3a, 0x62, 0xb7, 0xe9, 0xbc, 0x2c, 0x2d, 0xfa, 0xd2,
	0x61, 0x94, 0x06, 0x78, 0x0c, 0x6a, 0x94, 0x21, 0x96, 0xd3, 0xf6, 0x93, 0x8e, 0xd8, 0x6d, 0x1c,
	0x29, 0x1b, 0x56, 0x5c, 0xdc, 0x87, 0xea, 0x72, 0xa7, 0x53, 0x4e, 0x28, 0xbf, 0x44, 0x50, 0x9b,
	0x23, 0xf8, 0x02, 0x40, 0xd7, 0xd3, 0xbc, 0xa1, 0xeb, 0x0f, 0x2d, 0xd7, 0x36, 0x74, 0xf3, 0xd4,
	0x34, 0x4e, 0x24, 0x01, 0x36, 0x41, 0xbd, 0xe4, 0x9a, 0xee, 0x99, 0x5f, 0x0d, 0x49, 0x5c, 0x41,
	0xb6, 0x36, 0x74, 0x8d, 0x13, 0xa9, 0xb2, 0x32, 0xad, 0x6b, 0x96, 0x6e, 0x0c, 0xfc, 0x0b, 0x6b,
	0xf0, 0x4d, 0xda, 0x82, 0xcf, 0x81, 0xb4, 0xb0, 0x5e, 0xb8, 0xde, 0x9c, 0x56, 0x61, 0x0b, 0x3c,
	0x2b, 0xa9, 0x69, 0x99, 0x9e, 0xa9, 0x0d, 0xcc, 0x4b, 0xd3, 0xfa, 0x2c, 0x6d, 0xc3, 0x03, 0xd0,
	0x2a, 0x85, 0x53, 0xd3, 0xd2, 0x06, 0xbe, 0x6b, 0x78, 0xde, 0xc0, 0x38, 0x37, 0x2c, 0x4f, 0xaa,
	0xf5, 0x01, 0xd8, 0x59, 0xdc, 0x66, 0xdf, 0xfe, 0x3d, 0x95, 0xc5, 0xfb, 0xa9, 0x2c, 0xfe, 0x9d,
	0xca, 0xe2, 0xcf, 0x99, 0x2c, 0xdc, 0xcf, 0x64, 0xe1, 0xcf, 0x4c, 0x16, 0x2e, 0x3f, 0xfe, 0x88,
	0xd9, 0x55, 0x1e, 0xa8, 0x21, 0x19, 0xf7, 0x1e, 0xbc, 0xac, 0xdb, 0x0f, 0xef, 0xc2, 0x2b, 0x14,
	0x27, 0xbd, 0x25, 0x99, 0xcc, 0x5f, 0x1b, 0xbb, 0x4b, 0x31, 0x0d, 0x6a, 0x1c, 0xbf, 0xff, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x96, 0xa7, 0xc6, 0x50, 0x8f, 0x03, 0x00, 0x00,
}

func (m *PerpetualClobMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpetualClobMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpetualClobMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PerpetualId != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SpotClobMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpotClobMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpotClobMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QuoteAssetId != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.QuoteAssetId))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseAssetId != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.BaseAssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClobPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClobPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClobPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if m.QuantumConversionExponent != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64((uint32(m.QuantumConversionExponent)<<1)^uint32((m.QuantumConversionExponent>>31))))
		i--
		dAtA[i] = 0x30
	}
	if m.SubticksPerTick != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.SubticksPerTick))
		i--
		dAtA[i] = 0x28
	}
	if m.StepBaseQuantums != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.StepBaseQuantums))
		i--
		dAtA[i] = 0x20
	}
	if m.Metadata != nil {
		{
			size := m.Metadata.Size()
			i -= size
			if _, err := m.Metadata.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Id != 0 {
		i = encodeVarintClobPair(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClobPair_PerpetualClobMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClobPair_PerpetualClobMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PerpetualClobMetadata != nil {
		{
			size, err := m.PerpetualClobMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClobPair(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ClobPair_SpotClobMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClobPair_SpotClobMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SpotClobMetadata != nil {
		{
			size, err := m.SpotClobMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClobPair(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintClobPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovClobPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PerpetualClobMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualId != 0 {
		n += 1 + sovClobPair(uint64(m.PerpetualId))
	}
	return n
}

func (m *SpotClobMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAssetId != 0 {
		n += 1 + sovClobPair(uint64(m.BaseAssetId))
	}
	if m.QuoteAssetId != 0 {
		n += 1 + sovClobPair(uint64(m.QuoteAssetId))
	}
	return n
}

func (m *ClobPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovClobPair(uint64(m.Id))
	}
	if m.Metadata != nil {
		n += m.Metadata.Size()
	}
	if m.StepBaseQuantums != 0 {
		n += 1 + sovClobPair(uint64(m.StepBaseQuantums))
	}
	if m.SubticksPerTick != 0 {
		n += 1 + sovClobPair(uint64(m.SubticksPerTick))
	}
	if m.QuantumConversionExponent != 0 {
		n += 1 + sozClobPair(uint64(m.QuantumConversionExponent))
	}
	if m.Status != 0 {
		n += 1 + sovClobPair(uint64(m.Status))
	}
	return n
}

func (m *ClobPair_PerpetualClobMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualClobMetadata != nil {
		l = m.PerpetualClobMetadata.Size()
		n += 1 + l + sovClobPair(uint64(l))
	}
	return n
}
func (m *ClobPair_SpotClobMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpotClobMetadata != nil {
		l = m.SpotClobMetadata.Size()
		n += 1 + l + sovClobPair(uint64(l))
	}
	return n
}

func sovClobPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClobPair(x uint64) (n int) {
	return sovClobPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerpetualClobMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClobPair
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PerpetualClobMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpetualClobMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PerpetualId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClobPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClobPair
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpotClobMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClobPair
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SpotClobMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpotClobMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAssetId", wireType)
			}
			m.BaseAssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseAssetId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetId", wireType)
			}
			m.QuoteAssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuoteAssetId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClobPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClobPair
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClobPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClobPair
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClobPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClobPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualClobMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClobPair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClobPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PerpetualClobMetadata{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Metadata = &ClobPair_PerpetualClobMetadata{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpotClobMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClobPair
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClobPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SpotClobMetadata{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Metadata = &ClobPair_SpotClobMetadata{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepBaseQuantums", wireType)
			}
			m.StepBaseQuantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepBaseQuantums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubticksPerTick", wireType)
			}
			m.SubticksPerTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubticksPerTick |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuantumConversionExponent", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.QuantumConversionExponent = v
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ClobPair_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClobPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClobPair
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClobPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClobPair
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClobPair
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClobPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClobPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClobPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClobPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClobPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClobPair = fmt.Errorf("proto: unexpected end of group")
)
