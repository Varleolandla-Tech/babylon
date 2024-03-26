// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/finality.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
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

// IndexedBlock is the necessary metadata and finalization status of a block
type IndexedBlock struct {
	// height is the height of the block
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// app_hash is the AppHash of the block
	AppHash []byte `protobuf:"bytes,2,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	// finalized indicates whether the IndexedBlock is finalised by 2/3
	// finality providers or not
	Finalized bool `protobuf:"varint,3,opt,name=finalized,proto3" json:"finalized,omitempty"`
}

func (m *IndexedBlock) Reset()         { *m = IndexedBlock{} }
func (m *IndexedBlock) String() string { return proto.CompactTextString(m) }
func (*IndexedBlock) ProtoMessage()    {}
func (*IndexedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{0}
}
func (m *IndexedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedBlock.Merge(m, src)
}
func (m *IndexedBlock) XXX_Size() int {
	return m.Size()
}
func (m *IndexedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedBlock proto.InternalMessageInfo

func (m *IndexedBlock) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *IndexedBlock) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *IndexedBlock) GetFinalized() bool {
	if m != nil {
		return m.Finalized
	}
	return false
}

// Evidence is the evidence that a finality provider has signed finality
// signatures with correct public randomness on two conflicting Babylon headers
type Evidence struct {
	// fp_btc_pk is the BTC PK of the finality provider that casts this vote
	FpBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=fp_btc_pk,json=fpBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"fp_btc_pk,omitempty"`
	// block_height is the height of the conflicting blocks
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// master_pub_rand is the master public randomness the finality provider has committed to
	MasterPubRand string `protobuf:"bytes,3,opt,name=master_pub_rand,json=masterPubRand,proto3" json:"master_pub_rand,omitempty"`
	// canonical_app_hash is the AppHash of the canonical block
	CanonicalAppHash []byte `protobuf:"bytes,4,opt,name=canonical_app_hash,json=canonicalAppHash,proto3" json:"canonical_app_hash,omitempty"`
	// fork_app_hash is the AppHash of the fork block
	ForkAppHash []byte `protobuf:"bytes,5,opt,name=fork_app_hash,json=forkAppHash,proto3" json:"fork_app_hash,omitempty"`
	// canonical_finality_sig is the finality signature to the canonical block
	// where finality signature is an EOTS signature, i.e.,
	// the `s` in a Schnorr signature `(r, s)`
	// `r` is the public randomness that is already committed by the finality provider
	CanonicalFinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,6,opt,name=canonical_finality_sig,json=canonicalFinalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"canonical_finality_sig,omitempty"`
	// fork_finality_sig is the finality signature to the fork block
	// where finality signature is an EOTS signature
	ForkFinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,7,opt,name=fork_finality_sig,json=forkFinalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"fork_finality_sig,omitempty"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca5b87e52e3e6d02, []int{1}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

func (m *Evidence) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Evidence) GetMasterPubRand() string {
	if m != nil {
		return m.MasterPubRand
	}
	return ""
}

func (m *Evidence) GetCanonicalAppHash() []byte {
	if m != nil {
		return m.CanonicalAppHash
	}
	return nil
}

func (m *Evidence) GetForkAppHash() []byte {
	if m != nil {
		return m.ForkAppHash
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexedBlock)(nil), "babylon.finality.v1.IndexedBlock")
	proto.RegisterType((*Evidence)(nil), "babylon.finality.v1.Evidence")
}

func init() {
	proto.RegisterFile("babylon/finality/v1/finality.proto", fileDescriptor_ca5b87e52e3e6d02)
}

var fileDescriptor_ca5b87e52e3e6d02 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x6e, 0xb6, 0xd1, 0x0f, 0xaf, 0xd3, 0xc0, 0x4c, 0x53, 0x41, 0x28, 0x2b, 0x3d, 0xa0, 0x1e,
	0x50, 0xb3, 0xb1, 0x09, 0x71, 0x25, 0xd2, 0xd0, 0x06, 0x07, 0x2a, 0x97, 0x13, 0x17, 0xcb, 0x76,
	0xdc, 0xd8, 0x6a, 0x67, 0x5b, 0x89, 0x5b, 0x2d, 0xfc, 0x0a, 0x7e, 0x16, 0xc7, 0x71, 0x43, 0x3b,
	0x54, 0xa8, 0xfd, 0x23, 0x28, 0x4e, 0x9a, 0xc0, 0x09, 0x24, 0x6e, 0x7e, 0x9f, 0xf7, 0xd1, 0xf3,
	0x21, 0xbf, 0x60, 0x40, 0x09, 0xcd, 0xe6, 0x5a, 0x05, 0x53, 0xa9, 0xc8, 0x5c, 0xda, 0x2c, 0x58,
	0x9e, 0x55, 0xef, 0x91, 0x49, 0xb4, 0xd5, 0xf0, 0x71, 0xc9, 0x19, 0x55, 0xf8, 0xf2, 0xec, 0xe9,
	0x51, 0xac, 0x63, 0xed, 0xf6, 0x41, 0xfe, 0x2a, 0xa8, 0x03, 0x0c, 0xba, 0xd7, 0x2a, 0xe2, 0xb7,
	0x3c, 0x0a, 0xe7, 0x9a, 0xcd, 0xe0, 0x31, 0x68, 0x0a, 0x2e, 0x63, 0x61, 0x7b, 0x5e, 0xdf, 0x1b,
	0xee, 0xa1, 0x72, 0x82, 0x4f, 0x40, 0x9b, 0x18, 0x83, 0x05, 0x49, 0x45, 0x6f, 0xa7, 0xef, 0x0d,
	0xbb, 0xa8, 0x45, 0x8c, 0xb9, 0x22, 0xa9, 0x80, 0xcf, 0x40, 0xa7, 0xf0, 0xf9, 0xc2, 0xa3, 0xde,
	0x6e, 0xdf, 0x1b, 0xb6, 0x51, 0x0d, 0x0c, 0xbe, 0xef, 0x82, 0xf6, 0xe5, 0x52, 0x46, 0x5c, 0x31,
	0x0e, 0x11, 0xe8, 0x4c, 0x0d, 0xa6, 0x96, 0x61, 0x33, 0x73, 0x06, 0xdd, 0xf0, 0xf5, 0xfd, 0xea,
	0xe4, 0x55, 0x2c, 0xad, 0x58, 0xd0, 0x11, 0xd3, 0x37, 0x41, 0x19, 0x9d, 0x09, 0x22, 0xd5, 0x76,
	0x08, 0x6c, 0x66, 0x78, 0x3a, 0x0a, 0xaf, 0xc7, 0xe7, 0x17, 0xa7, 0xe3, 0x05, 0xfd, 0xc0, 0x33,
	0xd4, 0x9a, 0x9a, 0xd0, 0xb2, 0xf1, 0x0c, 0x3e, 0x07, 0x5d, 0x9a, 0x47, 0xc7, 0x65, 0xee, 0x1d,
	0x97, 0x7b, 0xdf, 0x61, 0x57, 0x45, 0xf8, 0x17, 0xe0, 0xf0, 0x86, 0xa4, 0x96, 0x27, 0xd8, 0x2c,
	0x28, 0x4e, 0x88, 0x2a, 0x72, 0x76, 0xd0, 0x41, 0x01, 0x8f, 0x17, 0x14, 0x11, 0x15, 0xc1, 0x97,
	0x00, 0x32, 0xa2, 0xb4, 0x92, 0x8c, 0xcc, 0x71, 0x55, 0x77, 0xcf, 0xd5, 0x7d, 0x58, 0x6d, 0xde,
	0x96, 0xbd, 0x07, 0xe0, 0x60, 0xaa, 0x93, 0x59, 0x4d, 0x7c, 0xe0, 0x88, 0xfb, 0x39, 0xb8, 0xe5,
	0x28, 0x70, 0x5c, 0x2b, 0x6e, 0x7f, 0x03, 0xa7, 0x32, 0xee, 0x35, 0x5d, 0xfb, 0x37, 0xf7, 0xab,
	0x93, 0x8b, 0x7f, 0x6b, 0x3f, 0x61, 0x42, 0xe9, 0x24, 0xb9, 0xfc, 0xf8, 0x69, 0x32, 0x91, 0x31,
	0x3a, 0xaa, 0x74, 0xdf, 0x95, 0xb2, 0x13, 0x19, 0xc3, 0x08, 0x3c, 0x72, 0x99, 0xfe, 0xb0, 0x6a,
	0xfd, 0xa7, 0xd5, 0x61, 0x2e, 0xf9, 0x9b, 0x4b, 0xf8, 0xfe, 0xdb, 0xda, 0xf7, 0xee, 0xd6, 0xbe,
	0xf7, 0x73, 0xed, 0x7b, 0x5f, 0x37, 0x7e, 0xe3, 0x6e, 0xe3, 0x37, 0x7e, 0x6c, 0xfc, 0xc6, 0xe7,
	0xd3, 0xbf, 0x19, 0xdc, 0xd6, 0x77, 0xeb, 0xbc, 0x68, 0xd3, 0xdd, 0xe1, 0xf9, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x12, 0x77, 0xc2, 0xaa, 0xd8, 0x02, 0x00, 0x00,
}

func (m *IndexedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Finalized {
		i--
		if m.Finalized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ForkFinalitySig != nil {
		{
			size := m.ForkFinalitySig.Size()
			i -= size
			if _, err := m.ForkFinalitySig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.CanonicalFinalitySig != nil {
		{
			size := m.CanonicalFinalitySig.Size()
			i -= size
			if _, err := m.CanonicalFinalitySig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ForkAppHash) > 0 {
		i -= len(m.ForkAppHash)
		copy(dAtA[i:], m.ForkAppHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.ForkAppHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CanonicalAppHash) > 0 {
		i -= len(m.CanonicalAppHash)
		copy(dAtA[i:], m.CanonicalAppHash)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.CanonicalAppHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MasterPubRand) > 0 {
		i -= len(m.MasterPubRand)
		copy(dAtA[i:], m.MasterPubRand)
		i = encodeVarintFinality(dAtA, i, uint64(len(m.MasterPubRand)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintFinality(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.FpBtcPk != nil {
		{
			size := m.FpBtcPk.Size()
			i -= size
			if _, err := m.FpBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintFinality(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFinality(dAtA []byte, offset int, v uint64) int {
	offset -= sovFinality(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovFinality(uint64(m.Height))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.Finalized {
		n += 2
	}
	return n
}

func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FpBtcPk != nil {
		l = m.FpBtcPk.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovFinality(uint64(m.BlockHeight))
	}
	l = len(m.MasterPubRand)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	l = len(m.CanonicalAppHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	l = len(m.ForkAppHash)
	if l > 0 {
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.CanonicalFinalitySig != nil {
		l = m.CanonicalFinalitySig.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	if m.ForkFinalitySig != nil {
		l = m.ForkFinalitySig.Size()
		n += 1 + l + sovFinality(uint64(l))
	}
	return n
}

func sovFinality(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFinality(x uint64) (n int) {
	return sovFinality(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: IndexedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Finalized = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinality
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FpBtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.FpBtcPk = &v
			if err := m.FpBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterPubRand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
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
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterPubRand = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CanonicalAppHash = append(m.CanonicalAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CanonicalAppHash == nil {
				m.CanonicalAppHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForkAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForkAppHash = append(m.ForkAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ForkAppHash == nil {
				m.ForkAppHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalFinalitySig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.CanonicalFinalitySig = &v
			if err := m.CanonicalFinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForkFinalitySig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinality
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFinality
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFinality
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.ForkFinalitySig = &v
			if err := m.ForkFinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFinality(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinality
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
func skipFinality(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFinality
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
					return 0, ErrIntOverflowFinality
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
					return 0, ErrIntOverflowFinality
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
				return 0, ErrInvalidLengthFinality
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFinality
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFinality
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFinality        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFinality          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFinality = fmt.Errorf("proto: unexpected end of group")
)
