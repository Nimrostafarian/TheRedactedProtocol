// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/marketmaker/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState defines the marketmaker module's genesis state.
type GenesisState struct {
	// params defines all the parameters for the marketmaker module
	Params         Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	MarketMakers   []MarketMaker   `protobuf:"bytes,2,rep,name=market_makers,json=marketMakers,proto3" json:"market_makers" yaml:"market_makers"`
	Incentives     []Incentive     `protobuf:"bytes,3,rep,name=incentives,proto3" json:"incentives" yaml:"incentives"`
	DepositRecords []DepositRecord `protobuf:"bytes,4,rep,name=deposit_records,json=depositRecords,proto3" json:"deposit_records" yaml:"deposit_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_15caac8a2480ca91, []int{0}
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

func init() {
	proto.RegisterType((*GenesisState)(nil), "crescent.marketmaker.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("crescent/marketmaker/v1beta1/genesis.proto", fileDescriptor_15caac8a2480ca91)
}

var fileDescriptor_15caac8a2480ca91 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8a, 0xd3, 0x40,
	0x1c, 0xc7, 0x13, 0xbb, 0x2c, 0x92, 0x5d, 0x15, 0xc3, 0x22, 0xe9, 0xb2, 0x4c, 0x4a, 0x10, 0xac,
	0x8a, 0x19, 0xda, 0xde, 0xea, 0x2d, 0x08, 0xe2, 0xa1, 0x20, 0xf1, 0xa4, 0x97, 0x32, 0x49, 0x7f,
	0xc6, 0xd0, 0x4e, 0x26, 0xcc, 0x4c, 0xab, 0x7d, 0x03, 0x8f, 0x3e, 0x42, 0x0f, 0x1e, 0x7c, 0x94,
	0x1e, 0x7b, 0xf4, 0x54, 0x24, 0xbd, 0x78, 0xf6, 0x09, 0x24, 0x33, 0x69, 0x9b, 0xf6, 0x90, 0xdb,
	0xcc, 0xef, 0xf7, 0xf9, 0xfe, 0x19, 0x18, 0xeb, 0x45, 0xcc, 0x41, 0xc4, 0x90, 0x49, 0x4c, 0x09,
	0x9f, 0x82, 0xa4, 0x64, 0x0a, 0x1c, 0x2f, 0x7a, 0x11, 0x48, 0xd2, 0xc3, 0x09, 0x64, 0x20, 0x52,
	0xe1, 0xe7, 0x9c, 0x49, 0x66, 0xdf, 0xed, 0x59, 0xbf, 0xc6, 0xfa, 0x15, 0x7b, 0xdb, 0x4e, 0x18,
	0x4b, 0x66, 0x80, 0x15, 0x1b, 0xcd, 0x3f, 0x63, 0x92, 0x2d, 0xb5, 0xf0, 0xf6, 0x26, 0x61, 0x09,
	0x53, 0x47, 0x5c, 0x9e, 0xaa, 0x69, 0x3b, 0x66, 0x82, 0x32, 0x31, 0xd6, 0x0b, 0x7d, 0xa9, 0x56,
	0x48, 0xdf, 0x70, 0x44, 0x04, 0x1c, 0xca, 0xc4, 0x2c, 0xcd, 0xaa, 0xbd, 0xdf, 0xd8, 0xba, 0xde,
	0x4e, 0xf3, 0xee, 0x79, 0x37, 0x99, 0x52, 0x10, 0x92, 0xd0, 0x5c, 0x03, 0xde, 0xcf, 0x96, 0x75,
	0xfd, 0x56, 0x3f, 0xf6, 0x83, 0x24, 0x12, 0xec, 0xc0, 0xba, 0xcc, 0x09, 0x27, 0x54, 0x38, 0x66,
	0xc7, 0xec, 0x5e, 0xf5, 0x9f, 0xfa, 0x4d, 0x8f, 0xf7, 0xdf, 0x2b, 0x36, 0xb8, 0x58, 0x6f, 0x5d,
	0x23, 0xac, 0x94, 0xf6, 0xcc, 0x7a, 0xa0, 0xd9, 0xb1, 0x82, 0x85, 0x73, 0xaf, 0xd3, 0xea, 0x5e,
	0xf5, 0x9f, 0x37, 0x5b, 0x8d, 0xd4, 0x6c, 0x54, 0xce, 0x82, 0xbb, 0xd2, 0xef, 0xdf, 0xd6, 0xbd,
	0x59, 0x12, 0x3a, 0x1b, 0x7a, 0x27, 0x6e, 0x5e, 0x78, 0x4d, 0x8f, 0xa8, 0xb0, 0x23, 0xcb, 0x4a,
	0xb3, 0xd2, 0x35, 0x5d, 0x80, 0x70, 0x5a, 0x2a, 0xea, 0x59, 0x73, 0xd4, 0xbb, 0x3d, 0x1f, 0xb4,
	0xab, 0xa0, 0xc7, 0x3a, 0xe8, 0x68, 0xe4, 0x85, 0x35, 0x57, 0x5b, 0x5a, 0x8f, 0x26, 0x90, 0x33,
	0x91, 0xca, 0x31, 0x87, 0x98, 0xf1, 0x89, 0x70, 0x2e, 0x54, 0xd0, 0xcb, 0xe6, 0xa0, 0x37, 0x5a,
	0x14, 0x2a, 0x4d, 0x80, 0xaa, 0xb0, 0x27, 0x3a, 0xec, 0xcc, 0xd1, 0x0b, 0x1f, 0x4e, 0xea, 0xb8,
	0x18, 0xde, 0xff, 0xbe, 0x72, 0x8d, 0xbf, 0x2b, 0xd7, 0x08, 0x3e, 0xfe, 0x2a, 0x90, 0xb9, 0x2e,
	0x90, 0xb9, 0x29, 0x90, 0xf9, 0xa7, 0x40, 0xe6, 0x8f, 0x1d, 0x32, 0x36, 0x3b, 0x64, 0xfc, 0xde,
	0x21, 0xe3, 0xd3, 0xeb, 0x24, 0x95, 0x5f, 0xe6, 0x91, 0x1f, 0x33, 0x8a, 0xf7, 0x75, 0x5e, 0x65,
	0x20, 0xbf, 0x32, 0x3e, 0x3d, 0x0c, 0xf0, 0x62, 0x80, 0xbf, 0x9d, 0x7c, 0x1b, 0xb9, 0xcc, 0x41,
	0x44, 0x97, 0xea, 0x23, 0x0c, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x32, 0xb8, 0xa7, 0x11,
	0x03, 0x00, 0x00,
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
	if len(m.DepositRecords) > 0 {
		for iNdEx := len(m.DepositRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Incentives) > 0 {
		for iNdEx := len(m.Incentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Incentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.MarketMakers) > 0 {
		for iNdEx := len(m.MarketMakers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MarketMakers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.MarketMakers) > 0 {
		for _, e := range m.MarketMakers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Incentives) > 0 {
		for _, e := range m.Incentives {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DepositRecords) > 0 {
		for _, e := range m.DepositRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketMakers", wireType)
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
			m.MarketMakers = append(m.MarketMakers, MarketMaker{})
			if err := m.MarketMakers[len(m.MarketMakers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Incentives", wireType)
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
			m.Incentives = append(m.Incentives, Incentive{})
			if err := m.Incentives[len(m.Incentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositRecords", wireType)
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
			m.DepositRecords = append(m.DepositRecords, DepositRecord{})
			if err := m.DepositRecords[len(m.DepositRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
