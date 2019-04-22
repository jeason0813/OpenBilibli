// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/Rank2018.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Rank2018GetHourRankReq struct {
	// 父分区id
	AreaV2ParentId int64 `protobuf:"varint,1,opt,name=area_v2_parent_id,json=areaV2ParentId,proto3" json:"area_v2_parent_id"`
	// 子分区id
	AreaV2Id int64 `protobuf:"varint,2,opt,name=area_v2_id,json=areaV2Id,proto3" json:"area_v2_id"`
	// 不传则为上小时
	Hour string `protobuf:"bytes,3,opt,name=hour,proto3" json:"hour"`
	// N
	Top int64 `protobuf:"varint,4,opt,name=top,proto3" json:"top"`
}

func (m *Rank2018GetHourRankReq) Reset()         { *m = Rank2018GetHourRankReq{} }
func (m *Rank2018GetHourRankReq) String() string { return proto.CompactTextString(m) }
func (*Rank2018GetHourRankReq) ProtoMessage()    {}
func (*Rank2018GetHourRankReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_Rank2018_eaa61d01276fe5ed, []int{0}
}
func (m *Rank2018GetHourRankReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rank2018GetHourRankReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rank2018GetHourRankReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Rank2018GetHourRankReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rank2018GetHourRankReq.Merge(dst, src)
}
func (m *Rank2018GetHourRankReq) XXX_Size() int {
	return m.Size()
}
func (m *Rank2018GetHourRankReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Rank2018GetHourRankReq.DiscardUnknown(m)
}

var xxx_messageInfo_Rank2018GetHourRankReq proto.InternalMessageInfo

func (m *Rank2018GetHourRankReq) GetAreaV2ParentId() int64 {
	if m != nil {
		return m.AreaV2ParentId
	}
	return 0
}

func (m *Rank2018GetHourRankReq) GetAreaV2Id() int64 {
	if m != nil {
		return m.AreaV2Id
	}
	return 0
}

func (m *Rank2018GetHourRankReq) GetHour() string {
	if m != nil {
		return m.Hour
	}
	return ""
}

func (m *Rank2018GetHourRankReq) GetTop() int64 {
	if m != nil {
		return m.Top
	}
	return 0
}

type Rank2018GetHourRankResp struct {
	//
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	//
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data []int64 `protobuf:"varint,3,rep,packed,name=data" json:"data"`
}

func (m *Rank2018GetHourRankResp) Reset()         { *m = Rank2018GetHourRankResp{} }
func (m *Rank2018GetHourRankResp) String() string { return proto.CompactTextString(m) }
func (*Rank2018GetHourRankResp) ProtoMessage()    {}
func (*Rank2018GetHourRankResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_Rank2018_eaa61d01276fe5ed, []int{1}
}
func (m *Rank2018GetHourRankResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rank2018GetHourRankResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rank2018GetHourRankResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Rank2018GetHourRankResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rank2018GetHourRankResp.Merge(dst, src)
}
func (m *Rank2018GetHourRankResp) XXX_Size() int {
	return m.Size()
}
func (m *Rank2018GetHourRankResp) XXX_DiscardUnknown() {
	xxx_messageInfo_Rank2018GetHourRankResp.DiscardUnknown(m)
}

var xxx_messageInfo_Rank2018GetHourRankResp proto.InternalMessageInfo

func (m *Rank2018GetHourRankResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Rank2018GetHourRankResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Rank2018GetHourRankResp) GetData() []int64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Rank2018GetHourRankReq)(nil), "rankdb.v1.Rank2018GetHourRankReq")
	proto.RegisterType((*Rank2018GetHourRankResp)(nil), "rankdb.v1.Rank2018GetHourRankResp")
}
func (m *Rank2018GetHourRankReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rank2018GetHourRankReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AreaV2ParentId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(m.AreaV2ParentId))
	}
	if m.AreaV2Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(m.AreaV2Id))
	}
	if len(m.Hour) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(len(m.Hour)))
		i += copy(dAtA[i:], m.Hour)
	}
	if m.Top != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(m.Top))
	}
	return i, nil
}

func (m *Rank2018GetHourRankResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rank2018GetHourRankResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Data) > 0 {
		dAtA2 := make([]byte, len(m.Data)*10)
		var j1 int
		for _, num1 := range m.Data {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRank2018(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func encodeVarintRank2018(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Rank2018GetHourRankReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AreaV2ParentId != 0 {
		n += 1 + sovRank2018(uint64(m.AreaV2ParentId))
	}
	if m.AreaV2Id != 0 {
		n += 1 + sovRank2018(uint64(m.AreaV2Id))
	}
	l = len(m.Hour)
	if l > 0 {
		n += 1 + l + sovRank2018(uint64(l))
	}
	if m.Top != 0 {
		n += 1 + sovRank2018(uint64(m.Top))
	}
	return n
}

func (m *Rank2018GetHourRankResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovRank2018(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovRank2018(uint64(l))
	}
	if len(m.Data) > 0 {
		l = 0
		for _, e := range m.Data {
			l += sovRank2018(uint64(e))
		}
		n += 1 + sovRank2018(uint64(l)) + l
	}
	return n
}

func sovRank2018(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRank2018(x uint64) (n int) {
	return sovRank2018(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Rank2018GetHourRankReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRank2018
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rank2018GetHourRankReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rank2018GetHourRankReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AreaV2ParentId", wireType)
			}
			m.AreaV2ParentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AreaV2ParentId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AreaV2Id", wireType)
			}
			m.AreaV2Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AreaV2Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hour", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRank2018
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hour = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Top", wireType)
			}
			m.Top = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Top |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRank2018(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRank2018
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
func (m *Rank2018GetHourRankResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRank2018
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rank2018GetHourRankResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rank2018GetHourRankResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRank2018
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRank2018
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Data = append(m.Data, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRank2018
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRank2018
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Data) == 0 {
					m.Data = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRank2018
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Data = append(m.Data, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRank2018(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRank2018
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
func skipRank2018(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRank2018
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
					return 0, ErrIntOverflowRank2018
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRank2018
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRank2018
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRank2018
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRank2018(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRank2018 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRank2018   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/Rank2018.proto", fileDescriptor_Rank2018_eaa61d01276fe5ed) }

var fileDescriptor_Rank2018_eaa61d01276fe5ed = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x59, 0x4a, 0xfe, 0x7f, 0x58, 0x13, 0x12, 0x36, 0x51, 0x2b, 0x21, 0x2d, 0x72, 0xe2,
	0xa0, 0xc5, 0xd6, 0x8b, 0x47, 0xc3, 0x45, 0xb9, 0x99, 0x8d, 0xf1, 0xe0, 0x05, 0xb7, 0x6c, 0x2d,
	0x84, 0xc0, 0xae, 0xdb, 0x6d, 0x9f, 0xc3, 0x57, 0xf2, 0xe6, 0x91, 0xa3, 0xa7, 0xc6, 0xc0, 0xad,
	0x4f, 0x61, 0x76, 0xb0, 0x6a, 0x22, 0x89, 0x97, 0x2f, 0xfd, 0x4d, 0xe7, 0xfb, 0x66, 0xda, 0xc1,
	0xad, 0xcc, 0x1f, 0x50, 0xb6, 0x9c, 0x07, 0x67, 0xfe, 0x85, 0x27, 0x95, 0xd0, 0x82, 0x34, 0x14,
	0x5b, 0xce, 0x79, 0xe8, 0x65, 0x7e, 0xfb, 0x34, 0x9e, 0xe9, 0x69, 0x1a, 0x7a, 0x13, 0xb1, 0x18,
	0xc4, 0x22, 0x16, 0x03, 0xe8, 0x08, 0xd3, 0x47, 0x20, 0x00, 0x78, 0xda, 0x3a, 0x7b, 0x2f, 0x08,
	0x1f, 0x94, 0x61, 0x57, 0x91, 0xbe, 0x16, 0xa9, 0x32, 0x48, 0xa3, 0x27, 0x72, 0x89, 0x5b, 0x4c,
	0x45, 0x6c, 0x9c, 0x05, 0x63, 0xc9, 0x54, 0xb4, 0xd4, 0xe3, 0x19, 0xb7, 0x51, 0x17, 0xf5, 0xad,
	0xe1, 0x7e, 0x91, 0xbb, 0xbf, 0x5f, 0xd2, 0xa6, 0x29, 0xdd, 0x05, 0x37, 0x50, 0x18, 0x71, 0x72,
	0x82, 0x71, 0xd9, 0x34, 0xe3, 0x76, 0x15, 0xac, 0xcd, 0x22, 0x77, 0x7f, 0x54, 0x69, 0x7d, 0xeb,
	0x19, 0x71, 0xd2, 0xc1, 0xb5, 0xa9, 0x48, 0x95, 0x6d, 0x75, 0x51, 0xbf, 0x31, 0xac, 0x17, 0xb9,
	0x0b, 0x4c, 0x41, 0xc9, 0x11, 0xb6, 0xb4, 0x90, 0x76, 0x0d, 0x42, 0xfe, 0x17, 0xb9, 0x6b, 0x90,
	0x1a, 0xe9, 0x49, 0x7c, 0xb8, 0xf3, 0x13, 0x12, 0x69, 0x32, 0x27, 0x82, 0x47, 0x9f, 0x6b, 0x43,
	0xa6, 0x61, 0x0a, 0x6a, 0x32, 0x17, 0x49, 0x0c, 0x8b, 0x35, 0xb6, 0x99, 0x8b, 0x24, 0xa6, 0x46,
	0x8c, 0x91, 0x33, 0xcd, 0x6c, 0xab, 0x6b, 0x95, 0x46, 0xc3, 0x14, 0x34, 0x78, 0xc0, 0xf5, 0x72,
	0x22, 0xb9, 0xc5, 0x7b, 0xf1, 0xf7, 0x54, 0x72, 0xec, 0x7d, 0xdd, 0xc2, 0xdb, 0xfd, 0x63, 0xdb,
	0xbd, 0xbf, 0x5a, 0x12, 0x39, 0xec, 0xbc, 0xae, 0x1d, 0xb4, 0x5a, 0x3b, 0xe8, 0x7d, 0xed, 0xa0,
	0xe7, 0x8d, 0x53, 0x59, 0x6d, 0x9c, 0xca, 0xdb, 0xc6, 0xa9, 0xdc, 0x57, 0x33, 0x3f, 0xfc, 0x07,
	0xc7, 0x3b, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8a, 0x7b, 0xf7, 0x0b, 0x02, 0x00, 0x00,
}