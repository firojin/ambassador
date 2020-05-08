// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/v3/hash_policy.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Specifies the hash policy
type HashPolicy struct {
	// Types that are valid to be assigned to PolicySpecifier:
	//	*HashPolicy_SourceIp_
	PolicySpecifier      isHashPolicy_PolicySpecifier `protobuf_oneof:"policy_specifier"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *HashPolicy) Reset()         { *m = HashPolicy{} }
func (m *HashPolicy) String() string { return proto.CompactTextString(m) }
func (*HashPolicy) ProtoMessage()    {}
func (*HashPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc16c1c24253da6, []int{0}
}
func (m *HashPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HashPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HashPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashPolicy.Merge(m, src)
}
func (m *HashPolicy) XXX_Size() int {
	return m.Size()
}
func (m *HashPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_HashPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_HashPolicy proto.InternalMessageInfo

type isHashPolicy_PolicySpecifier interface {
	isHashPolicy_PolicySpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HashPolicy_SourceIp_ struct {
	SourceIp *HashPolicy_SourceIp `protobuf:"bytes,1,opt,name=source_ip,json=sourceIp,proto3,oneof" json:"source_ip,omitempty"`
}

func (*HashPolicy_SourceIp_) isHashPolicy_PolicySpecifier() {}

func (m *HashPolicy) GetPolicySpecifier() isHashPolicy_PolicySpecifier {
	if m != nil {
		return m.PolicySpecifier
	}
	return nil
}

func (m *HashPolicy) GetSourceIp() *HashPolicy_SourceIp {
	if x, ok := m.GetPolicySpecifier().(*HashPolicy_SourceIp_); ok {
		return x.SourceIp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HashPolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HashPolicy_SourceIp_)(nil),
	}
}

// The source IP will be used to compute the hash used by hash-based load balancing
// algorithms.
type HashPolicy_SourceIp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashPolicy_SourceIp) Reset()         { *m = HashPolicy_SourceIp{} }
func (m *HashPolicy_SourceIp) String() string { return proto.CompactTextString(m) }
func (*HashPolicy_SourceIp) ProtoMessage()    {}
func (*HashPolicy_SourceIp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cc16c1c24253da6, []int{0, 0}
}
func (m *HashPolicy_SourceIp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HashPolicy_SourceIp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashPolicy_SourceIp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HashPolicy_SourceIp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashPolicy_SourceIp.Merge(m, src)
}
func (m *HashPolicy_SourceIp) XXX_Size() int {
	return m.Size()
}
func (m *HashPolicy_SourceIp) XXX_DiscardUnknown() {
	xxx_messageInfo_HashPolicy_SourceIp.DiscardUnknown(m)
}

var xxx_messageInfo_HashPolicy_SourceIp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HashPolicy)(nil), "envoy.type.v3.HashPolicy")
	proto.RegisterType((*HashPolicy_SourceIp)(nil), "envoy.type.v3.HashPolicy.SourceIp")
}

func init() { proto.RegisterFile("envoy/type/v3/hash_policy.proto", fileDescriptor_6cc16c1c24253da6) }

var fileDescriptor_6cc16c1c24253da6 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x33, 0xd6, 0xcf, 0x48, 0x2c, 0xce, 0x88, 0x2f,
	0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0x2b, 0xd0,
	0x03, 0x29, 0xd0, 0x2b, 0x33, 0x96, 0x92, 0x2d, 0x4d, 0x29, 0x48, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x86, 0xa8,
	0x96, 0x52, 0xc4, 0x90, 0x2e, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xcb, 0xcc, 0x4b, 0x87, 0x2a,
	0x11, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x12, 0x4a, 0x3b,
	0x19, 0xb9, 0xb8, 0x3c, 0x12, 0x8b, 0x33, 0x02, 0xc0, 0xd6, 0x0b, 0x39, 0x72, 0x71, 0x16, 0xe7,
	0x97, 0x16, 0x25, 0xa7, 0xc6, 0x67, 0x16, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0xe9,
	0xa1, 0x38, 0x46, 0x0f, 0xa1, 0x5a, 0x2f, 0x18, 0xac, 0xd4, 0xb3, 0xc0, 0x83, 0x21, 0x88, 0xa3,
	0x18, 0xca, 0x96, 0x32, 0xe4, 0xe2, 0x80, 0x89, 0x5b, 0xa9, 0xce, 0x3a, 0xda, 0x21, 0xa7, 0xc0,
	0x25, 0x87, 0x64, 0x02, 0x16, 0xed, 0x56, 0x32, 0x20, 0x65, 0xe2, 0x5c, 0xa2, 0x58, 0x95, 0x39,
	0x89, 0x73, 0x09, 0x40, 0x02, 0x27, 0xbe, 0xb8, 0x20, 0x35, 0x39, 0x33, 0x2d, 0x33, 0xb5, 0x48,
	0x88, 0xf9, 0x87, 0x13, 0xa3, 0x93, 0xcb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0xb8, 0xab, 0xe1, 0xc4, 0x45, 0x36, 0x26, 0x01, 0x46, 0x2e, 0xe9, 0xcc, 0x7c,
	0x88, 0x8b, 0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0x51, 0x1d, 0xef, 0xc4, 0x8f, 0x30, 0x37, 0x00, 0xe4,
	0xff, 0x00, 0xc6, 0x24, 0x36, 0x70, 0x40, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x5b,
	0x1f, 0x16, 0x95, 0x01, 0x00, 0x00,
}

func (m *HashPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PolicySpecifier != nil {
		{
			size := m.PolicySpecifier.Size()
			i -= size
			if _, err := m.PolicySpecifier.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *HashPolicy_SourceIp_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashPolicy_SourceIp_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourceIp != nil {
		{
			size, err := m.SourceIp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHashPolicy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *HashPolicy_SourceIp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashPolicy_SourceIp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashPolicy_SourceIp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintHashPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovHashPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HashPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PolicySpecifier != nil {
		n += m.PolicySpecifier.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HashPolicy_SourceIp_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceIp != nil {
		l = m.SourceIp.Size()
		n += 1 + l + sovHashPolicy(uint64(l))
	}
	return n
}
func (m *HashPolicy_SourceIp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHashPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHashPolicy(x uint64) (n int) {
	return sovHashPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HashPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashPolicy
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
			return fmt.Errorf("proto: HashPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceIp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashPolicy
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
				return ErrInvalidLengthHashPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HashPolicy_SourceIp{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.PolicySpecifier = &HashPolicy_SourceIp_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashPolicy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHashPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashPolicy_SourceIp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashPolicy
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
			return fmt.Errorf("proto: SourceIp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SourceIp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHashPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashPolicy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHashPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHashPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHashPolicy
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
					return 0, ErrIntOverflowHashPolicy
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
					return 0, ErrIntOverflowHashPolicy
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
				return 0, ErrInvalidLengthHashPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHashPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHashPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHashPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHashPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHashPolicy = fmt.Errorf("proto: unexpected end of group")
)