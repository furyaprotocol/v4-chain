// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furyaprotocol/clob/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the clob module's genesis state.
type GenesisState struct {
	ClobPairs             []ClobPair                   `protobuf:"bytes,1,rep,name=clob_pairs,json=clobPairs,proto3" json:"clob_pairs"`
	LiquidationsConfig    LiquidationsConfig           `protobuf:"bytes,2,opt,name=liquidations_config,json=liquidationsConfig,proto3" json:"liquidations_config"`
	BlockRateLimitConfig  BlockRateLimitConfiguration  `protobuf:"bytes,3,opt,name=block_rate_limit_config,json=blockRateLimitConfig,proto3" json:"block_rate_limit_config"`
	EquityTierLimitConfig EquityTierLimitConfiguration `protobuf:"bytes,4,opt,name=equity_tier_limit_config,json=equityTierLimitConfig,proto3" json:"equity_tier_limit_config"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2de77065a6fbee92, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClobPairs() []ClobPair {
	if m != nil {
		return m.ClobPairs
	}
	return nil
}

func (m *GenesisState) GetLiquidationsConfig() LiquidationsConfig {
	if m != nil {
		return m.LiquidationsConfig
	}
	return LiquidationsConfig{}
}

func (m *GenesisState) GetBlockRateLimitConfig() BlockRateLimitConfiguration {
	if m != nil {
		return m.BlockRateLimitConfig
	}
	return BlockRateLimitConfiguration{}
}

func (m *GenesisState) GetEquityTierLimitConfig() EquityTierLimitConfiguration {
	if m != nil {
		return m.EquityTierLimitConfig
	}
	return EquityTierLimitConfiguration{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "furyaprotocol.clob.GenesisState")
}

func init() { proto.RegisterFile("furyaprotocol/clob/genesis.proto", fileDescriptor_2de77065a6fbee92) }

var fileDescriptor_2de77065a6fbee92 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdb, 0x77, 0xe3, 0x05, 0x33, 0x2f, 0xd6, 0x89, 0x65, 0x42, 0x37, 0x05, 0x61, 0x20,
	0xb6, 0x32, 0xc5, 0xb3, 0x6c, 0x88, 0x97, 0x1d, 0xc6, 0xf4, 0x24, 0x42, 0x49, 0xb3, 0xd8, 0xfd,
	0x59, 0xd6, 0x6c, 0x69, 0x2a, 0xdb, 0x67, 0xf0, 0xe2, 0xc7, 0xda, 0x71, 0x47, 0x4f, 0x22, 0xdb,
	0x17, 0x91, 0xa6, 0x65, 0x6c, 0x34, 0xbb, 0x94, 0xf6, 0xe9, 0xef, 0x79, 0x9e, 0xe4, 0x9f, 0xa0,
	0xfa, 0x60, 0x3e, 0x98, 0x4d, 0x04, 0x97, 0x9c, 0x70, 0xe6, 0x11, 0xc6, 0x03, 0x2f, 0xa4, 0x11,
	0x8d, 0x21, 0x76, 0x95, 0x6a, 0x1d, 0x6d, 0x03, 0x6e, 0x0a, 0xd4, 0xaa, 0x21, 0x0f, 0xb9, 0x92,
	0xbc, 0xf4, 0x2d, 0x03, 0x6b, 0x5e, 0x31, 0x29, 0x60, 0x9c, 0x8c, 0x7c, 0x81, 0x25, 0xf5, 0x19,
	0x8c, 0x41, 0xfa, 0x84, 0x47, 0xef, 0x10, 0xe6, 0x86, 0xf3, 0xa2, 0x21, 0x7d, 0xf8, 0x13, 0x0c,
	0x22, 0x47, 0x6e, 0x8a, 0x08, 0x9d, 0x26, 0x20, 0xe7, 0xbe, 0x04, 0x2a, 0x74, 0xa1, 0x57, 0x45,
	0x07, 0x83, 0x69, 0x02, 0x03, 0x2c, 0x81, 0x47, 0xf1, 0x0e, 0x7c, 0xf1, 0x59, 0x42, 0x87, 0x4f,
	0xd9, 0x6e, 0x9f, 0x25, 0x96, 0xd4, 0x7a, 0x40, 0x68, 0xb3, 0x84, 0xd8, 0x36, 0x1b, 0xa5, 0x66,
	0xa5, 0x75, 0xe6, 0x16, 0x26, 0xe0, 0x76, 0x18, 0x0f, 0x7a, 0x18, 0x44, 0xbb, 0xbc, 0xf8, 0xa9,
	0x1b, 0xfd, 0x03, 0x92, 0x7f, 0xc7, 0xd6, 0x1b, 0x3a, 0xd6, 0xf4, 0xd9, 0xff, 0x1a, 0x66, 0xb3,
	0xd2, 0xba, 0xd4, 0x44, 0x75, 0xb7, 0xe8, 0x8e, 0x82, 0xf3, 0x50, 0x8b, 0x15, 0xfe, 0x58, 0x23,
	0x74, 0xba, 0x67, 0xa6, 0x76, 0x49, 0x35, 0xb8, 0x9a, 0x86, 0x76, 0xea, 0xe8, 0x63, 0x49, 0xbb,
	0x29, 0x9f, 0x25, 0x25, 0x42, 0xe5, 0xe6, 0x55, 0xd5, 0x40, 0x83, 0x58, 0x11, 0xb2, 0xf7, 0x0d,
	0xdb, 0x2e, 0xab, 0x36, 0x4f, 0xd3, 0xf6, 0xa8, 0x2c, 0x2f, 0x40, 0xc5, 0xde, 0xba, 0x13, 0xaa,
	0x63, 0xda, 0xbd, 0xc5, 0xca, 0x31, 0x97, 0x2b, 0xc7, 0xfc, 0x5d, 0x39, 0xe6, 0xd7, 0xda, 0x31,
	0x96, 0x6b, 0xc7, 0xf8, 0x5e, 0x3b, 0xc6, 0xeb, 0x7d, 0x08, 0x72, 0x98, 0x04, 0x2e, 0xe1, 0xe3,
	0xdd, 0x5b, 0xf6, 0x71, 0x77, 0x4d, 0x86, 0x18, 0x22, 0x6f, 0xa3, 0xcc, 0xb2, 0x33, 0x97, 0xf3,
	0x09, 0x8d, 0x83, 0xff, 0x4a, 0xbe, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x44, 0xfb, 0x91, 0xc2,
	0xe5, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EquityTierLimitConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.BlockRateLimitConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.LiquidationsConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ClobPairs) > 0 {
		for iNdEx := len(m.ClobPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClobPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClobPairs) > 0 {
		for _, e := range m.ClobPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.LiquidationsConfig.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.BlockRateLimitConfig.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.EquityTierLimitConfig.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClobPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClobPairs = append(m.ClobPairs, ClobPair{})
			if err := m.ClobPairs[len(m.ClobPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationsConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationsConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockRateLimitConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockRateLimitConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EquityTierLimitConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EquityTierLimitConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
