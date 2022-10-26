// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contractmanager/genesis.proto

package types

import (
	fmt "fmt"
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

type Failure struct {
	// Address of the failed contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Offset used to add more failures under one address
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// ACK id to restore
	AckId string `protobuf:"bytes,3,opt,name=ackId,proto3" json:"ackId,omitempty"`
	// Ackonowledgement type
	AckType string `protobuf:"bytes,4,opt,name=ackType,proto3" json:"ackType,omitempty"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23af9b9805fb076, []int{0}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return m.Size()
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Failure) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Failure) GetAckId() string {
	if m != nil {
		return m.AckId
	}
	return ""
}

func (m *Failure) GetAckType() string {
	if m != nil {
		return m.AckType
	}
	return ""
}

type NextFailureId struct {
	// Address of the failed contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Next offset for the failed contracts mapping
	NextOffset uint64 `protobuf:"varint,2,opt,name=nextOffset,proto3" json:"nextOffset,omitempty"`
}

func (m *NextFailureId) Reset()         { *m = NextFailureId{} }
func (m *NextFailureId) String() string { return proto.CompactTextString(m) }
func (*NextFailureId) ProtoMessage()    {}
func (*NextFailureId) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23af9b9805fb076, []int{1}
}
func (m *NextFailureId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NextFailureId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NextFailureId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NextFailureId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextFailureId.Merge(m, src)
}
func (m *NextFailureId) XXX_Size() int {
	return m.Size()
}
func (m *NextFailureId) XXX_DiscardUnknown() {
	xxx_messageInfo_NextFailureId.DiscardUnknown(m)
}

var xxx_messageInfo_NextFailureId proto.InternalMessageInfo

func (m *NextFailureId) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NextFailureId) GetNextOffset() uint64 {
	if m != nil {
		return m.NextOffset
	}
	return 0
}

// GenesisState defines the contractmanager module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// List of the contract failures
	FailureList []Failure `protobuf:"bytes,2,rep,name=failureList,proto3" json:"failureList"`
	// List of the contract failures
	NextFailureId []NextFailureId `protobuf:"bytes,3,rep,name=nextFailureId,proto3" json:"nextFailureId"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23af9b9805fb076, []int{2}
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

func (m *GenesisState) GetFailureList() []Failure {
	if m != nil {
		return m.FailureList
	}
	return nil
}

func (m *GenesisState) GetNextFailureId() []NextFailureId {
	if m != nil {
		return m.NextFailureId
	}
	return nil
}

func init() {
	proto.RegisterType((*Failure)(nil), "neutronorg.neutron.contractmanager.Failure")
	proto.RegisterType((*NextFailureId)(nil), "neutronorg.neutron.contractmanager.NextFailureId")
	proto.RegisterType((*GenesisState)(nil), "neutronorg.neutron.contractmanager.GenesisState")
}

func init() { proto.RegisterFile("contractmanager/genesis.proto", fileDescriptor_c23af9b9805fb076) }

var fileDescriptor_c23af9b9805fb076 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x77, 0xd5, 0x94, 0xc6, 0xbc, 0x0c, 0x12, 0x8b, 0xd4, 0x24, 0x7b, 0x92, 0xa2, 0x5d,
	0x32, 0xe8, 0x03, 0x78, 0xa8, 0x84, 0x28, 0xd1, 0x4e, 0x41, 0x87, 0x71, 0x7d, 0x4e, 0x8b, 0x39,
	0xb3, 0xcc, 0x8c, 0xa0, 0xdf, 0xa1, 0x43, 0x1f, 0xcb, 0xa3, 0xc7, 0x4e, 0x11, 0xfa, 0x45, 0xc2,
	0xd9, 0x11, 0x56, 0x09, 0xf2, 0xf6, 0xfe, 0xc3, 0xfb, 0xff, 0xfe, 0xf3, 0x1e, 0x0f, 0x9d, 0x46,
	0x82, 0x6b, 0x49, 0x23, 0x3d, 0xa6, 0x9c, 0x32, 0x90, 0x21, 0x03, 0x0e, 0x2a, 0x56, 0x41, 0x22,
	0x85, 0x16, 0xd8, 0xe7, 0x30, 0xd1, 0x52, 0x70, 0x21, 0x59, 0x60, 0xcb, 0x60, 0xc7, 0x51, 0xab,
	0x32, 0xc1, 0x84, 0x69, 0x0f, 0xd7, 0x55, 0xea, 0xac, 0x9d, 0xec, 0x82, 0x13, 0x2a, 0xe9, 0xd8,
	0x72, 0xfd, 0x11, 0x2a, 0xdd, 0xd2, 0xf8, 0x7d, 0x22, 0x01, 0x7b, 0xa8, 0x44, 0x07, 0x03, 0x09,
	0x4a, 0x79, 0x6e, 0xdd, 0x6d, 0x1c, 0x76, 0x37, 0x12, 0x1f, 0xa3, 0xa2, 0x18, 0x0e, 0x15, 0x68,
	0x2f, 0x57, 0x77, 0x1b, 0x85, 0xae, 0x55, 0xb8, 0x8a, 0x0e, 0x68, 0x34, 0x6a, 0x0f, 0xbc, 0xbc,
	0xe9, 0x4f, 0x85, 0xe1, 0x44, 0xa3, 0xe7, 0x59, 0x02, 0x5e, 0xc1, 0x72, 0x52, 0xe9, 0xb7, 0x51,
	0xe5, 0x11, 0xa6, 0xda, 0x06, 0xda, 0xd6, 0xbf, 0x23, 0x09, 0x42, 0x1c, 0xa6, 0xfa, 0x29, 0x1b,
	0x9b, 0x79, 0xf1, 0x3f, 0x72, 0xe8, 0xe8, 0x2e, 0xdd, 0x50, 0x4f, 0x53, 0x0d, 0xf8, 0x1e, 0x15,
	0xd3, 0xc1, 0x0c, 0xa9, 0xdc, 0x3c, 0x0f, 0xfe, 0xdf, 0x58, 0xd0, 0x31, 0x8e, 0x56, 0x61, 0xfe,
	0x7d, 0xe6, 0x74, 0xad, 0x1f, 0xf7, 0x50, 0x79, 0x98, 0xfe, 0xf0, 0x21, 0x56, 0xeb, 0xec, 0x7c,
	0xa3, 0xdc, 0xbc, 0xd8, 0x07, 0x67, 0x07, 0xb3, 0xbc, 0x2c, 0x05, 0xbf, 0xa2, 0x0a, 0xcf, 0x8e,
	0xee, 0xe5, 0x0d, 0xf6, 0x6a, 0x1f, 0xec, 0xd6, 0xce, 0x2c, 0x7c, 0x9b, 0xd6, 0xea, 0xcc, 0x97,
	0xc4, 0x5d, 0x2c, 0x89, 0xfb, 0xb3, 0x24, 0xee, 0xe7, 0x8a, 0x38, 0x8b, 0x15, 0x71, 0xbe, 0x56,
	0xc4, 0x79, 0xb9, 0x61, 0xb1, 0x7e, 0x9b, 0xf4, 0x83, 0x48, 0x8c, 0x43, 0x1b, 0x70, 0x29, 0x24,
	0xdb, 0xd4, 0xe1, 0x34, 0xdc, 0xbd, 0x0f, 0x3d, 0x4b, 0x40, 0xf5, 0x8b, 0xe6, 0x3e, 0xae, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x9d, 0x2e, 0xb5, 0x98, 0x02, 0x00, 0x00,
}

func (m *Failure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Failure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Failure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AckType) > 0 {
		i -= len(m.AckType)
		copy(dAtA[i:], m.AckType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AckType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AckId) > 0 {
		i -= len(m.AckId)
		copy(dAtA[i:], m.AckId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AckId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Offset != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NextFailureId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NextFailureId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NextFailureId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextOffset != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextOffset))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.NextFailureId) > 0 {
		for iNdEx := len(m.NextFailureId) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NextFailureId[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.FailureList) > 0 {
		for iNdEx := len(m.FailureList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FailureList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *Failure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovGenesis(uint64(m.Offset))
	}
	l = len(m.AckId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.AckType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *NextFailureId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NextOffset != 0 {
		n += 1 + sovGenesis(uint64(m.NextOffset))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FailureList) > 0 {
		for _, e := range m.FailureList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NextFailureId) > 0 {
		for _, e := range m.NextFailureId {
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
func (m *Failure) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Failure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckId", wireType)
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
			m.AckId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
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
			m.AckType = string(dAtA[iNdEx:postIndex])
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
func (m *NextFailureId) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NextFailureId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NextFailureId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextOffset", wireType)
			}
			m.NextOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
				return fmt.Errorf("proto: wrong wireType = %d for field FailureList", wireType)
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
			m.FailureList = append(m.FailureList, Failure{})
			if err := m.FailureList[len(m.FailureList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextFailureId", wireType)
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
			m.NextFailureId = append(m.NextFailureId, NextFailureId{})
			if err := m.NextFailureId[len(m.NextFailureId)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
