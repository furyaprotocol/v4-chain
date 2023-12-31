// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furyaprotocol/perpetuals/params.proto

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

// Params defines the parameters for x/perpetuals module.
type Params struct {
	// Funding rate clamp factor in parts-per-million, used for clamping 8-hour
	// funding rates according to equation: |R| <= funding_rate_clamp_factor *
	// (initial margin - maintenance margin).
	FundingRateClampFactorPpm uint32 `protobuf:"varint,1,opt,name=funding_rate_clamp_factor_ppm,json=fundingRateClampFactorPpm,proto3" json:"funding_rate_clamp_factor_ppm,omitempty"`
	// Premium vote clamp factor in parts-per-million, used for clamping premium
	// votes according to equation: |V| <= premium_vote_clamp_factor *
	// (initial margin - maintenance margin).
	PremiumVoteClampFactorPpm uint32 `protobuf:"varint,2,opt,name=premium_vote_clamp_factor_ppm,json=premiumVoteClampFactorPpm,proto3" json:"premium_vote_clamp_factor_ppm,omitempty"`
	// Minimum number of premium votes per premium sample. If number of premium
	// votes is smaller than this number, pad with zeros up to this number.
	MinNumVotesPerSample uint32 `protobuf:"varint,3,opt,name=min_num_votes_per_sample,json=minNumVotesPerSample,proto3" json:"min_num_votes_per_sample,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b16af88c7880f7e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetFundingRateClampFactorPpm() uint32 {
	if m != nil {
		return m.FundingRateClampFactorPpm
	}
	return 0
}

func (m *Params) GetPremiumVoteClampFactorPpm() uint32 {
	if m != nil {
		return m.PremiumVoteClampFactorPpm
	}
	return 0
}

func (m *Params) GetMinNumVotesPerSample() uint32 {
	if m != nil {
		return m.MinNumVotesPerSample
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "furyaprotocol.perpetuals.Params")
}

func init() {
	proto.RegisterFile("furyaprotocol/perpetuals/params.proto", fileDescriptor_8b16af88c7880f7e)
}

var fileDescriptor_8b16af88c7880f7e = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x73, 0x0a, 0x1d, 0x02, 0x2e, 0x45, 0x30, 0x0e, 0x1e, 0x22, 0x0e, 0x2e, 0x26, 0x83,
	0xe2, 0xe4, 0x20, 0x0a, 0x8e, 0x12, 0x2a, 0x74, 0x70, 0x39, 0xae, 0x97, 0xd7, 0xf6, 0x20, 0xef,
	0xee, 0x71, 0x77, 0x29, 0xed, 0xb7, 0xf0, 0x43, 0x39, 0x38, 0x76, 0x74, 0x94, 0xe4, 0x8b, 0x48,
	0xaf, 0x41, 0x2b, 0xba, 0xbe, 0xdf, 0xef, 0xfd, 0x86, 0x7f, 0x7a, 0x5e, 0xad, 0xaa, 0x25, 0x39,
	0x1b, 0xac, 0xb2, 0x75, 0x41, 0xe0, 0x08, 0x42, 0x23, 0x6b, 0x5f, 0x90, 0x74, 0x12, 0x7d, 0x1e,
	0xd1, 0xf0, 0x68, 0xd7, 0xca, 0x7f, 0xac, 0xb3, 0x37, 0x96, 0x0e, 0xca, 0x68, 0x0e, 0xef, 0xd2,
	0x93, 0x69, 0x63, 0x2a, 0x6d, 0x66, 0xc2, 0xc9, 0x00, 0x42, 0xd5, 0x12, 0x49, 0x4c, 0xa5, 0x0a,
	0xd6, 0x09, 0x22, 0xcc, 0xd8, 0x29, 0xbb, 0x38, 0x18, 0x1d, 0xf7, 0xd2, 0x48, 0x06, 0x78, 0xd8,
	0x28, 0x8f, 0xd1, 0x28, 0x09, 0x37, 0x05, 0x72, 0x80, 0xba, 0x41, 0xb1, 0xb0, 0xff, 0x15, 0xf6,
	0xb6, 0x85, 0x5e, 0x1a, 0xdb, 0x3f, 0x85, 0x9b, 0x34, 0x43, 0x6d, 0x84, 0xe9, 0x0b, 0x5e, 0x10,
	0x38, 0xe1, 0x25, 0x52, 0x0d, 0xd9, 0x7e, 0x7c, 0x3e, 0x44, 0x6d, 0x9e, 0xb6, 0xbf, 0xbe, 0x04,
	0xf7, 0x1c, 0xd9, 0xfd, 0xf8, 0xbd, 0xe5, 0x6c, 0xdd, 0x72, 0xf6, 0xd9, 0x72, 0xf6, 0xda, 0xf1,
	0x64, 0xdd, 0xf1, 0xe4, 0xa3, 0xe3, 0xc9, 0xcb, 0xed, 0x4c, 0x87, 0x79, 0x33, 0xc9, 0x95, 0xc5,
	0xe2, 0xd7, 0x54, 0x8b, 0xeb, 0x4b, 0x35, 0x97, 0xda, 0x14, 0xdf, 0x97, 0xe5, 0xee, 0x7c, 0x61,
	0x45, 0xe0, 0x27, 0x83, 0x08, 0xaf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x45, 0x37, 0xf4, 0x5b,
	0x66, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinNumVotesPerSample != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinNumVotesPerSample))
		i--
		dAtA[i] = 0x18
	}
	if m.PremiumVoteClampFactorPpm != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PremiumVoteClampFactorPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.FundingRateClampFactorPpm != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FundingRateClampFactorPpm))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FundingRateClampFactorPpm != 0 {
		n += 1 + sovParams(uint64(m.FundingRateClampFactorPpm))
	}
	if m.PremiumVoteClampFactorPpm != 0 {
		n += 1 + sovParams(uint64(m.PremiumVoteClampFactorPpm))
	}
	if m.MinNumVotesPerSample != 0 {
		n += 1 + sovParams(uint64(m.MinNumVotesPerSample))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingRateClampFactorPpm", wireType)
			}
			m.FundingRateClampFactorPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FundingRateClampFactorPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PremiumVoteClampFactorPpm", wireType)
			}
			m.PremiumVoteClampFactorPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PremiumVoteClampFactorPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinNumVotesPerSample", wireType)
			}
			m.MinNumVotesPerSample = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinNumVotesPerSample |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
