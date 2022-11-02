// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	ed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the checkpointing module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to checkpointing
	Params      Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	GenesisKeys []*GenesisKey `protobuf:"bytes,2,rep,name=genesis_keys,json=genesisKeys,proto3" json:"genesis_keys,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdd0fb065da1de51, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGenesisKeys() []*GenesisKey {
	if m != nil {
		return m.GenesisKeys
	}
	return nil
}

type GenesisKey struct {
	// validator_address is the address corresponding to a validator
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// bls_key defines the BLS key of the validator at genesis
	BlsKey *BlsKey `protobuf:"bytes,2,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	// val_pubkey defines the ed25519 public key of the validator at genesis
	ValPubkey *ed25519.PubKey `protobuf:"bytes,3,opt,name=val_pubkey,json=valPubkey,proto3" json:"val_pubkey,omitempty"`
}

func (m *GenesisKey) Reset()         { *m = GenesisKey{} }
func (m *GenesisKey) String() string { return proto.CompactTextString(m) }
func (*GenesisKey) ProtoMessage()    {}
func (*GenesisKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdd0fb065da1de51, []int{1}
}
func (m *GenesisKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisKey.Merge(m, src)
}
func (m *GenesisKey) XXX_Size() int {
	return m.Size()
}
func (m *GenesisKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisKey.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisKey proto.InternalMessageInfo

func (m *GenesisKey) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *GenesisKey) GetBlsKey() *BlsKey {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *GenesisKey) GetValPubkey() *ed25519.PubKey {
	if m != nil {
		return m.ValPubkey
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "babylon.checkpointing.v1.GenesisState")
	proto.RegisterType((*GenesisKey)(nil), "babylon.checkpointing.v1.GenesisKey")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/genesis.proto", fileDescriptor_cdd0fb065da1de51)
}

var fileDescriptor_cdd0fb065da1de51 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0xdb, 0x4d, 0x26, 0xcb, 0x76, 0xd0, 0xe2, 0xa1, 0x0e, 0xac, 0x65, 0x7a, 0x18, 0x08,
	0x09, 0x9b, 0xec, 0x30, 0x10, 0xc1, 0x5d, 0x76, 0xf0, 0xe0, 0x98, 0x37, 0x2f, 0x25, 0x69, 0x43,
	0x57, 0xd6, 0x35, 0xa5, 0xc9, 0x8a, 0x7d, 0x0b, 0x6f, 0xbe, 0x87, 0x4f, 0xb1, 0xe3, 0x8e, 0x9e,
	0x44, 0xb6, 0x17, 0x91, 0x26, 0xd9, 0x44, 0x71, 0x78, 0x6a, 0xf2, 0xe5, 0xf7, 0xff, 0xbe, 0xef,
	0xdf, 0x3f, 0xb8, 0x20, 0x98, 0x14, 0x31, 0x4b, 0x90, 0x3f, 0xa5, 0xfe, 0x2c, 0x65, 0x51, 0x22,
	0xa2, 0x24, 0x44, 0x21, 0x4d, 0x28, 0x8f, 0x38, 0x4c, 0x33, 0x26, 0x98, 0x65, 0x6b, 0x08, 0xfe,
	0x80, 0x60, 0xde, 0x6d, 0x9d, 0xfa, 0x8c, 0xcf, 0x19, 0xf7, 0x24, 0x87, 0xd4, 0x45, 0x89, 0x5a,
	0x27, 0x21, 0x0b, 0x99, 0xaa, 0x97, 0x27, 0x5d, 0x75, 0x15, 0x83, 0xfc, 0xac, 0x48, 0x05, 0x43,
	0x34, 0xe8, 0xf5, 0xfb, 0xdd, 0x01, 0x9a, 0xd1, 0x62, 0xab, 0x6b, 0xff, 0xbd, 0x51, 0x8a, 0x33,
	0x3c, 0xdf, 0x32, 0x7b, 0xb6, 0x26, 0x31, 0xf7, 0x66, 0xb4, 0x50, 0x50, 0xfb, 0xd5, 0x04, 0xcd,
	0x91, 0xf2, 0xf1, 0x28, 0xb0, 0xa0, 0xd6, 0x2d, 0xa8, 0xa9, 0x2e, 0xb6, 0xe9, 0x9a, 0x9d, 0x46,
	0xcf, 0x85, 0xfb, 0x7c, 0xc1, 0xb1, 0xe4, 0x86, 0x07, 0xcb, 0x8f, 0x73, 0x63, 0xa2, 0x55, 0xd6,
	0x08, 0x34, 0xf5, 0x7f, 0x29, 0xa7, 0x70, 0xbb, 0xe2, 0x56, 0x3b, 0x8d, 0xde, 0xe5, 0xfe, 0x2e,
	0x7a, 0xfa, 0x3d, 0x2d, 0x26, 0x8d, 0x70, 0x77, 0xe6, 0xed, 0x37, 0x13, 0x80, 0xef, 0x37, 0xeb,
	0x0a, 0x1c, 0xe7, 0x38, 0x8e, 0x02, 0x2c, 0x58, 0xe6, 0xe1, 0x20, 0xc8, 0x28, 0x57, 0x2b, 0xd6,
	0x27, 0x47, 0xbb, 0x87, 0x3b, 0x55, 0xb7, 0x06, 0xe0, 0x50, 0xdb, 0xb4, 0x2b, 0xff, 0xb9, 0x18,
	0xc6, 0x72, 0x76, 0x8d, 0xc8, 0xaf, 0x75, 0x03, 0x40, 0x8e, 0x63, 0x2f, 0x5d, 0x90, 0x52, 0x5d,
	0x95, 0xea, 0x33, 0xa8, 0x43, 0x53, 0x81, 0x40, 0x1d, 0x08, 0x1c, 0x2f, 0x48, 0x29, 0xad, 0xe7,
	0x38, 0x1e, 0x4b, 0x7e, 0xf8, 0xb0, 0x5c, 0x3b, 0xe6, 0x6a, 0xed, 0x98, 0x9f, 0x6b, 0xc7, 0x7c,
	0xd9, 0x38, 0xc6, 0x6a, 0xe3, 0x18, 0xef, 0x1b, 0xc7, 0x78, 0xea, 0x87, 0x91, 0x98, 0x2e, 0x08,
	0xf4, 0xd9, 0x1c, 0xe9, 0x5d, 0xfc, 0x29, 0x8e, 0x92, 0xed, 0x05, 0x3d, 0xff, 0xca, 0x49, 0x14,
	0x29, 0xe5, 0xa4, 0x26, 0x63, 0xba, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xd1, 0xac, 0x82,
	0x83, 0x02, 0x00, 0x00,
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
	if len(m.GenesisKeys) > 0 {
		for iNdEx := len(m.GenesisKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValPubkey != nil {
		{
			size, err := m.ValPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlsKey != nil {
		{
			size, err := m.BlsKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.GenesisKeys) > 0 {
		for _, e := range m.GenesisKeys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.BlsKey != nil {
		l = m.BlsKey.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ValPubkey != nil {
		l = m.ValPubkey.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisKeys", wireType)
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
			m.GenesisKeys = append(m.GenesisKeys, &GenesisKey{})
			if err := m.GenesisKeys[len(m.GenesisKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsKey", wireType)
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
			if m.BlsKey == nil {
				m.BlsKey = &BlsKey{}
			}
			if err := m.BlsKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValPubkey", wireType)
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
			if m.ValPubkey == nil {
				m.ValPubkey = &ed25519.PubKey{}
			}
			if err := m.ValPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
