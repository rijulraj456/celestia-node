// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fraud/pb/proof.proto

package fraud_pb

import (
	fmt "fmt"
	pb "github.com/celestiaorg/celestia-node/ipld/pb"
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

type Axis int32

const (
	Axis_ROW Axis = 0
	Axis_COL Axis = 1
)

var Axis_name = map[int32]string{
	0: "ROW",
	1: "COL",
}

var Axis_value = map[string]int32{
	"ROW": 0,
	"COL": 1,
}

func (x Axis) String() string {
	return proto.EnumName(Axis_name, int32(x))
}

func (Axis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{0}
}

type BadEncoding struct {
	HeaderHash []byte      `protobuf:"bytes,1,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	Height     uint64      `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Shares     []*pb.Share `protobuf:"bytes,3,rep,name=Shares,proto3" json:"Shares,omitempty"`
	Index      uint32      `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	Axis       Axis        `protobuf:"varint,5,opt,name=Axis,proto3,enum=fraud.pb.Axis" json:"Axis,omitempty"`
}

func (m *BadEncoding) Reset()         { *m = BadEncoding{} }
func (m *BadEncoding) String() string { return proto.CompactTextString(m) }
func (*BadEncoding) ProtoMessage()    {}
func (*BadEncoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{0}
}
func (m *BadEncoding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadEncoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadEncoding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadEncoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadEncoding.Merge(m, src)
}
func (m *BadEncoding) XXX_Size() int {
	return m.Size()
}
func (m *BadEncoding) XXX_DiscardUnknown() {
	xxx_messageInfo_BadEncoding.DiscardUnknown(m)
}

var xxx_messageInfo_BadEncoding proto.InternalMessageInfo

func (m *BadEncoding) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *BadEncoding) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BadEncoding) GetShares() []*pb.Share {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *BadEncoding) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BadEncoding) GetAxis() Axis {
	if m != nil {
		return m.Axis
	}
	return Axis_ROW
}

type FraudMessageRequest struct {
	RequestedProofType []string `protobuf:"bytes,1,rep,name=RequestedProofType,proto3" json:"RequestedProofType,omitempty"`
}

func (m *FraudMessageRequest) Reset()         { *m = FraudMessageRequest{} }
func (m *FraudMessageRequest) String() string { return proto.CompactTextString(m) }
func (*FraudMessageRequest) ProtoMessage()    {}
func (*FraudMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{1}
}
func (m *FraudMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FraudMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FraudMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FraudMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FraudMessageRequest.Merge(m, src)
}
func (m *FraudMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *FraudMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FraudMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FraudMessageRequest proto.InternalMessageInfo

func (m *FraudMessageRequest) GetRequestedProofType() []string {
	if m != nil {
		return m.RequestedProofType
	}
	return nil
}

type ProofResponse struct {
	Type  string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value [][]byte `protobuf:"bytes,2,rep,name=Value,proto3" json:"Value,omitempty"`
}

func (m *ProofResponse) Reset()         { *m = ProofResponse{} }
func (m *ProofResponse) String() string { return proto.CompactTextString(m) }
func (*ProofResponse) ProtoMessage()    {}
func (*ProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{2}
}
func (m *ProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofResponse.Merge(m, src)
}
func (m *ProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProofResponse proto.InternalMessageInfo

func (m *ProofResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ProofResponse) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FraudMessageResponse struct {
	Proofs []*ProofResponse `protobuf:"bytes,1,rep,name=Proofs,proto3" json:"Proofs,omitempty"`
}

func (m *FraudMessageResponse) Reset()         { *m = FraudMessageResponse{} }
func (m *FraudMessageResponse) String() string { return proto.CompactTextString(m) }
func (*FraudMessageResponse) ProtoMessage()    {}
func (*FraudMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{3}
}
func (m *FraudMessageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FraudMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FraudMessageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FraudMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FraudMessageResponse.Merge(m, src)
}
func (m *FraudMessageResponse) XXX_Size() int {
	return m.Size()
}
func (m *FraudMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FraudMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FraudMessageResponse proto.InternalMessageInfo

func (m *FraudMessageResponse) GetProofs() []*ProofResponse {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func init() {
	proto.RegisterEnum("fraud.pb.Axis", Axis_name, Axis_value)
	proto.RegisterType((*BadEncoding)(nil), "fraud.pb.BadEncoding")
	proto.RegisterType((*FraudMessageRequest)(nil), "fraud.pb.FraudMessageRequest")
	proto.RegisterType((*ProofResponse)(nil), "fraud.pb.ProofResponse")
	proto.RegisterType((*FraudMessageResponse)(nil), "fraud.pb.FraudMessageResponse")
}

func init() { proto.RegisterFile("fraud/pb/proof.proto", fileDescriptor_318cb87a8bb2d394) }

var fileDescriptor_318cb87a8bb2d394 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4a, 0xeb, 0x40,
	0x18, 0xcd, 0x34, 0x69, 0xee, 0xed, 0xd7, 0x1f, 0xca, 0xb4, 0xdc, 0x3b, 0xb8, 0x08, 0x21, 0x0b,
	0x19, 0x5c, 0x4c, 0xa1, 0xae, 0x5c, 0x5a, 0xa9, 0x56, 0x50, 0x2a, 0xa3, 0xe8, 0x7a, 0x62, 0xa6,
	0x6d, 0xa0, 0x24, 0x63, 0xa6, 0x85, 0xfa, 0x16, 0xbe, 0x85, 0xaf, 0xe2, 0xb2, 0x4b, 0x97, 0xd2,
	0xbe, 0x88, 0xcc, 0x34, 0xa2, 0x05, 0x77, 0xe7, 0x27, 0x39, 0xe7, 0x7c, 0x0c, 0x74, 0x27, 0x85,
	0x58, 0x26, 0x3d, 0x15, 0xf7, 0x54, 0x91, 0xe7, 0x13, 0xa6, 0x8a, 0x7c, 0x91, 0xe3, 0xbf, 0x56,
	0x65, 0x2a, 0x3e, 0xe8, 0xa4, 0x6a, 0x6e, 0x6d, 0x3d, 0x13, 0x85, 0xdc, 0xd9, 0xd1, 0x2b, 0x82,
	0xfa, 0x40, 0x24, 0xc3, 0xec, 0x31, 0x4f, 0xd2, 0x6c, 0x8a, 0x03, 0x80, 0x91, 0x14, 0x89, 0x2c,
	0x46, 0x42, 0xcf, 0x08, 0x0a, 0x11, 0x6d, 0xf0, 0x1f, 0x0a, 0xfe, 0x07, 0xfe, 0x48, 0xa6, 0xd3,
	0xd9, 0x82, 0x54, 0x42, 0x44, 0x3d, 0x5e, 0x32, 0x7c, 0x08, 0xfe, 0xad, 0x89, 0xd5, 0xc4, 0x0d,
	0x5d, 0x5a, 0xef, 0xb7, 0x98, 0x69, 0x63, 0x2a, 0x66, 0x56, 0xe6, 0xa5, 0x8b, 0xbb, 0x50, 0xbd,
	0xcc, 0x12, 0xb9, 0x22, 0x5e, 0x88, 0x68, 0x93, 0xef, 0x08, 0x8e, 0xc0, 0x3b, 0x5d, 0xa5, 0x9a,
	0x54, 0x43, 0x44, 0x5b, 0xfd, 0x16, 0xfb, 0xda, 0xcc, 0xc4, 0x2a, 0xd5, 0xdc, 0x7a, 0xd1, 0x10,
	0x3a, 0xe7, 0x46, 0xbe, 0x96, 0x5a, 0x8b, 0xa9, 0xe4, 0xf2, 0x69, 0x29, 0xf5, 0x02, 0x33, 0xc0,
	0x25, 0x94, 0xc9, 0x8d, 0xb9, 0xfb, 0xee, 0x59, 0x49, 0x82, 0x42, 0x97, 0xd6, 0xf8, 0x2f, 0x4e,
	0x74, 0x02, 0x4d, 0x4b, 0xb8, 0xd4, 0x2a, 0xcf, 0xb4, 0xc4, 0x18, 0xbc, 0xf2, 0x17, 0x44, 0x6b,
	0xdc, 0x62, 0xb3, 0xf2, 0x5e, 0xcc, 0x97, 0x92, 0x54, 0x42, 0x97, 0x36, 0xf8, 0x8e, 0x44, 0x17,
	0xd0, 0xdd, 0x5f, 0x50, 0x26, 0xf4, 0xc0, 0xb7, 0x91, 0xda, 0xd6, 0xd6, 0xfb, 0xff, 0xbf, 0xf7,
	0xef, 0x55, 0xf1, 0xf2, 0xb3, 0x23, 0x02, 0x9e, 0x39, 0x0c, 0xff, 0x01, 0x97, 0x8f, 0x1f, 0xda,
	0x8e, 0x01, 0x67, 0xe3, 0xab, 0x36, 0x1a, 0x90, 0xb7, 0x4d, 0x80, 0xd6, 0x9b, 0x00, 0x7d, 0x6c,
	0x02, 0xf4, 0xb2, 0x0d, 0x9c, 0xf5, 0x36, 0x70, 0xde, 0xb7, 0x81, 0x13, 0xfb, 0xf6, 0xbd, 0x8e,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x6e, 0x3b, 0x05, 0xe6, 0x01, 0x00, 0x00,
}

func (m *BadEncoding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadEncoding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadEncoding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Axis != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Axis))
		i--
		dAtA[i] = 0x28
	}
	if m.Index != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProof(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintProof(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FraudMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FraudMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FraudMessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestedProofType) > 0 {
		for iNdEx := len(m.RequestedProofType) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RequestedProofType[iNdEx])
			copy(dAtA[i:], m.RequestedProofType[iNdEx])
			i = encodeVarintProof(dAtA, i, uint64(len(m.RequestedProofType[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		for iNdEx := len(m.Value) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Value[iNdEx])
			copy(dAtA[i:], m.Value[iNdEx])
			i = encodeVarintProof(dAtA, i, uint64(len(m.Value[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintProof(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FraudMessageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FraudMessageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FraudMessageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for iNdEx := len(m.Proofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProof(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BadEncoding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovProof(uint64(m.Height))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	if m.Index != 0 {
		n += 1 + sovProof(uint64(m.Index))
	}
	if m.Axis != 0 {
		n += 1 + sovProof(uint64(m.Axis))
	}
	return n
}

func (m *FraudMessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RequestedProofType) > 0 {
		for _, s := range m.RequestedProofType {
			l = len(s)
			n += 1 + l + sovProof(uint64(l))
		}
	}
	return n
}

func (m *ProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	if len(m.Value) > 0 {
		for _, b := range m.Value {
			l = len(b)
			n += 1 + l + sovProof(uint64(l))
		}
	}
	return n
}

func (m *FraudMessageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for _, e := range m.Proofs {
			l = e.Size()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BadEncoding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: BadEncoding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadEncoding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, &pb.Share{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Axis", wireType)
			}
			m.Axis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Axis |= Axis(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func (m *FraudMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: FraudMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FraudMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedProofType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestedProofType = append(m.RequestedProofType, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func (m *ProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: ProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value, make([]byte, postIndex-iNdEx))
			copy(m.Value[len(m.Value)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func (m *FraudMessageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: FraudMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FraudMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proofs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proofs = append(m.Proofs, &ProofResponse{})
			if err := m.Proofs[len(m.Proofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)
