// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/provider/provider.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/ovrclk/akash/types"
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

// MsgCreateProvider defines an SDK message for creating a provider
type MsgCreateProvider struct {
	Owner      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	HostURI    string                                        `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host-uri" yaml:"host-uri"`
	Attributes Attributes                                    `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes" yaml:"attributes"`
}

func (m *MsgCreateProvider) Reset()         { *m = MsgCreateProvider{} }
func (m *MsgCreateProvider) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProvider) ProtoMessage()    {}
func (*MsgCreateProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_f678b91550e081a6, []int{0}
}
func (m *MsgCreateProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProvider.Merge(m, src)
}
func (m *MsgCreateProvider) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProvider proto.InternalMessageInfo

func (m *MsgCreateProvider) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *MsgCreateProvider) GetHostURI() string {
	if m != nil {
		return m.HostURI
	}
	return ""
}

func (m *MsgCreateProvider) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// MsgUpdateProvider defines an SDK message for updating a provider
type MsgUpdateProvider struct {
	Owner      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	HostURI    string                                        `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host-uri" yaml:"host-uri"`
	Attributes Attributes                                    `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes" yaml:"attributes"`
}

func (m *MsgUpdateProvider) Reset()         { *m = MsgUpdateProvider{} }
func (m *MsgUpdateProvider) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateProvider) ProtoMessage()    {}
func (*MsgUpdateProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_f678b91550e081a6, []int{1}
}
func (m *MsgUpdateProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateProvider.Merge(m, src)
}
func (m *MsgUpdateProvider) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateProvider proto.InternalMessageInfo

func (m *MsgUpdateProvider) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *MsgUpdateProvider) GetHostURI() string {
	if m != nil {
		return m.HostURI
	}
	return ""
}

func (m *MsgUpdateProvider) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// MsgDeleteProvider defines an SDK message for deleting a provider
type MsgDeleteProvider struct {
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
}

func (m *MsgDeleteProvider) Reset()         { *m = MsgDeleteProvider{} }
func (m *MsgDeleteProvider) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProvider) ProtoMessage()    {}
func (*MsgDeleteProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_f678b91550e081a6, []int{2}
}
func (m *MsgDeleteProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProvider.Merge(m, src)
}
func (m *MsgDeleteProvider) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProvider proto.InternalMessageInfo

func (m *MsgDeleteProvider) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

// Provider stores owner and host details
type Provider struct {
	Owner      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	HostURI    string                                        `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host-uri" yaml:"host-uri"`
	Attributes Attributes                                    `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=Attributes" json:"attributes" yaml:"attributes"`
}

func (m *Provider) Reset()      { *m = Provider{} }
func (*Provider) ProtoMessage() {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_f678b91550e081a6, []int{3}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Provider) GetHostURI() string {
	if m != nil {
		return m.HostURI
	}
	return ""
}

func (m *Provider) GetAttributes() Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateProvider)(nil), "akash.provider.MsgCreateProvider")
	proto.RegisterType((*MsgUpdateProvider)(nil), "akash.provider.MsgUpdateProvider")
	proto.RegisterType((*MsgDeleteProvider)(nil), "akash.provider.MsgDeleteProvider")
	proto.RegisterType((*Provider)(nil), "akash.provider.Provider")
}

func init() { proto.RegisterFile("akash/provider/provider.proto", fileDescriptor_f678b91550e081a6) }

var fileDescriptor_f678b91550e081a6 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0x82, 0x33, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xf8, 0xc0, 0xd2, 0x7a, 0x30, 0x51, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0,
	0x94, 0x3e, 0x88, 0x05, 0x51, 0x25, 0x25, 0x0a, 0x31, 0x24, 0xb1, 0xa4, 0xa4, 0x28, 0x33, 0xa9,
	0xb4, 0x24, 0x15, 0x22, 0xac, 0xb4, 0x9a, 0x89, 0x4b, 0xd0, 0xb7, 0x38, 0xdd, 0xb9, 0x28, 0x35,
	0xb1, 0x24, 0x35, 0x00, 0x6a, 0x84, 0x50, 0x02, 0x17, 0x6b, 0x7e, 0x79, 0x5e, 0x6a, 0x91, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xd7, 0xab, 0x7b, 0xf2, 0x10, 0x81, 0x4f, 0xf7, 0xe4, 0x79,
	0x2a, 0x13, 0x73, 0x73, 0xac, 0x94, 0xc0, 0x5c, 0xa5, 0x5f, 0xf7, 0xe4, 0x75, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xa1, 0x94,
	0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0xb1, 0x9e, 0x63, 0x72, 0xb2, 0x63, 0x4a,
	0x4a, 0x51, 0x6a, 0x71, 0x71, 0x10, 0xc4, 0x1c, 0x21, 0x57, 0x2e, 0x8e, 0x8c, 0xfc, 0xe2, 0x92,
	0xf8, 0xd2, 0xa2, 0x4c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xad, 0x47, 0xf7, 0xe4, 0xd9,
	0x3d, 0xf2, 0x8b, 0x4b, 0x42, 0x83, 0x3c, 0x5f, 0xdd, 0x93, 0x07, 0x4b, 0xeb, 0x96, 0x16, 0x65,
	0x7e, 0xba, 0x27, 0xcf, 0x0f, 0xb1, 0x12, 0x26, 0xa2, 0x14, 0xc4, 0x0e, 0x62, 0x86, 0x16, 0x65,
	0x0a, 0xa5, 0x73, 0x71, 0xc1, 0x7d, 0x54, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa0,
	0x07, 0x09, 0x10, 0x47, 0x98, 0x84, 0x93, 0xe9, 0x89, 0x7b, 0xf2, 0x0c, 0xaf, 0xee, 0xc9, 0x23,
	0xa9, 0xfd, 0x74, 0x4f, 0x5e, 0x10, 0x62, 0x2a, 0x42, 0x4c, 0x69, 0xd5, 0x7d, 0x79, 0x2e, 0xb8,
	0xae, 0xe2, 0x20, 0x24, 0xe5, 0x56, 0x2c, 0x2f, 0x16, 0xc8, 0x33, 0xc0, 0x42, 0x2b, 0xb4, 0x20,
	0x65, 0x34, 0xb4, 0x08, 0x87, 0x56, 0x35, 0x38, 0xb0, 0x5c, 0x52, 0x73, 0x52, 0xe9, 0x19, 0x58,
	0x50, 0xcb, 0x97, 0x31, 0x71, 0x71, 0x8c, 0xc6, 0x10, 0xce, 0x18, 0xe2, 0x98, 0xb1, 0x40, 0x9e,
	0x01, 0x14, 0x50, 0x4e, 0x4e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5,
	0x81, 0x14, 0x12, 0xf9, 0x65, 0x45, 0xc9, 0x39, 0xd9, 0xfa, 0x90, 0x42, 0xa4, 0x02, 0x51, 0x16,
	0x81, 0xc3, 0x23, 0x89, 0x0d, 0x5c, 0x98, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0xa7,
	0x90, 0x99, 0xaa, 0x04, 0x00, 0x00,
}

func (m *MsgCreateProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProvider(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HostURI) > 0 {
		i -= len(m.HostURI)
		copy(dAtA[i:], m.HostURI)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.HostURI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProvider(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HostURI) > 0 {
		i -= len(m.HostURI)
		copy(dAtA[i:], m.HostURI)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.HostURI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Provider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProvider(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HostURI) > 0 {
		i -= len(m.HostURI)
		copy(dAtA[i:], m.HostURI)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.HostURI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.HostURI)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovProvider(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdateProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.HostURI)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovProvider(uint64(l))
		}
	}
	return n
}

func (m *MsgDeleteProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	return n
}

func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.HostURI)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovProvider(uint64(l))
		}
	}
	return n
}

func sovProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvider(x uint64) (n int) {
	return sovProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: MsgCreateProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, types.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProvider
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *MsgUpdateProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: MsgUpdateProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, types.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProvider
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *MsgDeleteProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: MsgDeleteProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProvider
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, types.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProvider
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProvider
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
func skipProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
				return 0, ErrInvalidLengthProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvider = fmt.Errorf("proto: unexpected end of group")
)