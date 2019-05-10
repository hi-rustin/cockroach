// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/distsqlpb/processors_base.proto

package distsqlpb

/*
	Beware! This package name must not be changed, even though it doesn't match
	the Go package name, because it defines the Protobuf message names which
	can't be changed without breaking backward compatibility.
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

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

// PostProcessSpec describes the processing required to obtain the output
// (filtering, projection). It operates on the internal schema of the processor
// (see ProcessorSpec).
type PostProcessSpec struct {
	// A filtering expression which references the internal columns of the
	// processor via ordinal references (@1, @2, etc).
	Filter Expression `protobuf:"bytes,1,opt,name=filter" json:"filter"`
	// If true, output_columns describes a projection. Used to differentiate
	// between an empty projection and no projection.
	//
	// Cannot be set at the same time with render expressions.
	Projection bool `protobuf:"varint,2,opt,name=projection" json:"projection"`
	// The output columns describe a projection on the internal set of columns;
	// only the columns in this list will be emitted.
	//
	// Can only be set if projection is true. Cannot be set at the same time with
	// render expressions.
	OutputColumns []uint32 `protobuf:"varint,3,rep,packed,name=output_columns,json=outputColumns" json:"output_columns,omitempty"`
	// If set, the output is the result of rendering these expressions. The
	// expressions reference the internal columns of the processor.
	//
	// Cannot be set at the same time with output columns.
	RenderExprs []Expression `protobuf:"bytes,4,rep,name=render_exprs,json=renderExprs" json:"render_exprs"`
	// If nonzero, the first <offset> rows will be suppressed.
	Offset uint64 `protobuf:"varint,5,opt,name=offset" json:"offset"`
	// If nonzero, the processor will stop after emitting this many rows. The rows
	// suppressed by <offset>, if any, do not count towards this limit.
	Limit                uint64   `protobuf:"varint,6,opt,name=limit" json:"limit"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostProcessSpec) Reset()         { *m = PostProcessSpec{} }
func (m *PostProcessSpec) String() string { return proto.CompactTextString(m) }
func (*PostProcessSpec) ProtoMessage()    {}
func (*PostProcessSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_processors_base_876aab0d2c832324, []int{0}
}
func (m *PostProcessSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostProcessSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *PostProcessSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostProcessSpec.Merge(dst, src)
}
func (m *PostProcessSpec) XXX_Size() int {
	return m.Size()
}
func (m *PostProcessSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PostProcessSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PostProcessSpec proto.InternalMessageInfo

type Columns struct {
	Columns              []uint32 `protobuf:"varint,1,rep,packed,name=columns" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Columns) Reset()         { *m = Columns{} }
func (m *Columns) String() string { return proto.CompactTextString(m) }
func (*Columns) ProtoMessage()    {}
func (*Columns) Descriptor() ([]byte, []int) {
	return fileDescriptor_processors_base_876aab0d2c832324, []int{1}
}
func (m *Columns) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Columns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Columns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Columns.Merge(dst, src)
}
func (m *Columns) XXX_Size() int {
	return m.Size()
}
func (m *Columns) XXX_DiscardUnknown() {
	xxx_messageInfo_Columns.DiscardUnknown(m)
}

var xxx_messageInfo_Columns proto.InternalMessageInfo

type TableReaderSpan struct {
	// TODO(radu): the dist_sql APIs should be agnostic to how we map tables to
	// KVs. The span should be described as starting and ending lists of values
	// for a prefix of the index columns, along with inclusive/exclusive flags.
	Span                 roachpb.Span `protobuf:"bytes,1,opt,name=span" json:"span"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TableReaderSpan) Reset()         { *m = TableReaderSpan{} }
func (m *TableReaderSpan) String() string { return proto.CompactTextString(m) }
func (*TableReaderSpan) ProtoMessage()    {}
func (*TableReaderSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_processors_base_876aab0d2c832324, []int{2}
}
func (m *TableReaderSpan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableReaderSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *TableReaderSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableReaderSpan.Merge(dst, src)
}
func (m *TableReaderSpan) XXX_Size() int {
	return m.Size()
}
func (m *TableReaderSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_TableReaderSpan.DiscardUnknown(m)
}

var xxx_messageInfo_TableReaderSpan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PostProcessSpec)(nil), "cockroach.sql.distsqlrun.PostProcessSpec")
	proto.RegisterType((*Columns)(nil), "cockroach.sql.distsqlrun.Columns")
	proto.RegisterType((*TableReaderSpan)(nil), "cockroach.sql.distsqlrun.TableReaderSpan")
}
func (m *PostProcessSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostProcessSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintProcessorsBase(dAtA, i, uint64(m.Filter.Size()))
	n1, err := m.Filter.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x10
	i++
	if m.Projection {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if len(m.OutputColumns) > 0 {
		dAtA3 := make([]byte, len(m.OutputColumns)*10)
		var j2 int
		for _, num := range m.OutputColumns {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProcessorsBase(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.RenderExprs) > 0 {
		for _, msg := range m.RenderExprs {
			dAtA[i] = 0x22
			i++
			i = encodeVarintProcessorsBase(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x28
	i++
	i = encodeVarintProcessorsBase(dAtA, i, uint64(m.Offset))
	dAtA[i] = 0x30
	i++
	i = encodeVarintProcessorsBase(dAtA, i, uint64(m.Limit))
	return i, nil
}

func (m *Columns) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Columns) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Columns) > 0 {
		dAtA5 := make([]byte, len(m.Columns)*10)
		var j4 int
		for _, num := range m.Columns {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintProcessorsBase(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	return i, nil
}

func (m *TableReaderSpan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableReaderSpan) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintProcessorsBase(dAtA, i, uint64(m.Span.Size()))
	n6, err := m.Span.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeVarintProcessorsBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PostProcessSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Filter.Size()
	n += 1 + l + sovProcessorsBase(uint64(l))
	n += 2
	if len(m.OutputColumns) > 0 {
		l = 0
		for _, e := range m.OutputColumns {
			l += sovProcessorsBase(uint64(e))
		}
		n += 1 + sovProcessorsBase(uint64(l)) + l
	}
	if len(m.RenderExprs) > 0 {
		for _, e := range m.RenderExprs {
			l = e.Size()
			n += 1 + l + sovProcessorsBase(uint64(l))
		}
	}
	n += 1 + sovProcessorsBase(uint64(m.Offset))
	n += 1 + sovProcessorsBase(uint64(m.Limit))
	return n
}

func (m *Columns) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Columns) > 0 {
		l = 0
		for _, e := range m.Columns {
			l += sovProcessorsBase(uint64(e))
		}
		n += 1 + sovProcessorsBase(uint64(l)) + l
	}
	return n
}

func (m *TableReaderSpan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Span.Size()
	n += 1 + l + sovProcessorsBase(uint64(l))
	return n
}

func sovProcessorsBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProcessorsBase(x uint64) (n int) {
	return sovProcessorsBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PostProcessSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsBase
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
			return fmt.Errorf("proto: PostProcessSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostProcessSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProcessorsBase
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Filter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Projection", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Projection = bool(v != 0)
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProcessorsBase
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.OutputColumns = append(m.OutputColumns, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProcessorsBase
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
					return ErrInvalidLengthProcessorsBase
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
				if elementCount != 0 && len(m.OutputColumns) == 0 {
					m.OutputColumns = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProcessorsBase
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.OutputColumns = append(m.OutputColumns, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputColumns", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenderExprs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProcessorsBase
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RenderExprs = append(m.RenderExprs, Expression{})
			if err := m.RenderExprs[len(m.RenderExprs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProcessorsBase
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
func (m *Columns) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsBase
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
			return fmt.Errorf("proto: Columns: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Columns: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProcessorsBase
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Columns = append(m.Columns, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProcessorsBase
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
					return ErrInvalidLengthProcessorsBase
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
				if elementCount != 0 && len(m.Columns) == 0 {
					m.Columns = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProcessorsBase
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Columns = append(m.Columns, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProcessorsBase
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
func (m *TableReaderSpan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsBase
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
			return fmt.Errorf("proto: TableReaderSpan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableReaderSpan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProcessorsBase
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Span.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProcessorsBase
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
func skipProcessorsBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProcessorsBase
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
					return 0, ErrIntOverflowProcessorsBase
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
					return 0, ErrIntOverflowProcessorsBase
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
				return 0, ErrInvalidLengthProcessorsBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProcessorsBase
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
				next, err := skipProcessorsBase(dAtA[start:])
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
	ErrInvalidLengthProcessorsBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProcessorsBase   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/distsqlpb/processors_base.proto", fileDescriptor_processors_base_876aab0d2c832324)
}

var fileDescriptor_processors_base_876aab0d2c832324 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x19, 0x28, 0x70, 0xef, 0x70, 0xb9, 0x98, 0x89, 0x89, 0x13, 0x42, 0x6a, 0x03, 0x24,
	0xd6, 0x4d, 0x89, 0x3e, 0x42, 0xd5, 0xa5, 0x09, 0x01, 0x57, 0x6e, 0xc8, 0xd0, 0x0e, 0x58, 0x2d,
	0x9d, 0x61, 0xce, 0x34, 0xe1, 0x31, 0x7c, 0x13, 0x5f, 0x83, 0xa5, 0x4b, 0x57, 0x46, 0xeb, 0x8b,
	0x18, 0x98, 0x96, 0xc0, 0xc2, 0x85, 0xbb, 0xc9, 0x77, 0xfe, 0x3f, 0xe7, 0xcb, 0x1c, 0xdc, 0x83,
	0x65, 0x3c, 0x08, 0x23, 0xd0, 0xb0, 0x8c, 0xe5, 0x74, 0x20, 0x95, 0x08, 0x38, 0x80, 0x50, 0x30,
	0x99, 0x32, 0xe0, 0x9e, 0x54, 0x42, 0x0b, 0x42, 0x03, 0x11, 0x3c, 0x29, 0xc1, 0x82, 0x07, 0x0f,
	0x96, 0xb1, 0x97, 0xc7, 0x55, 0x9a, 0xb4, 0xe9, 0x61, 0x3d, 0x64, 0x9a, 0x99, 0x4e, 0x9b, 0x6c,
	0xf3, 0x87, 0xec, 0x78, 0x2e, 0xe6, 0x62, 0xfb, 0x1c, 0x6c, 0x5e, 0x86, 0x76, 0x5f, 0xca, 0xb8,
	0x35, 0x14, 0xa0, 0x87, 0x66, 0xf7, 0x58, 0xf2, 0x80, 0xf8, 0xb8, 0x36, 0x8b, 0x62, 0xcd, 0x15,
	0x45, 0x0e, 0x72, 0x1b, 0x97, 0x7d, 0xef, 0x27, 0x05, 0xef, 0x66, 0x25, 0x15, 0x07, 0x88, 0x44,
	0xe2, 0x5b, 0xeb, 0xf7, 0xd3, 0xd2, 0x28, 0x6f, 0x92, 0x3e, 0xc6, 0x52, 0x89, 0x47, 0x1e, 0xe8,
	0x48, 0x24, 0xb4, 0xec, 0x20, 0xf7, 0x4f, 0x9e, 0xd8, 0xe3, 0xe4, 0x1c, 0xff, 0x17, 0xa9, 0x96,
	0xa9, 0x9e, 0x04, 0x22, 0x4e, 0x17, 0x09, 0xd0, 0x8a, 0x53, 0x71, 0x9b, 0x7e, 0xf9, 0x08, 0x8d,
	0x9a, 0x66, 0x72, 0x65, 0x06, 0xe4, 0x16, 0xff, 0x53, 0x3c, 0x09, 0xb9, 0x9a, 0xf0, 0x95, 0x54,
	0x40, 0x2d, 0xa7, 0xf2, 0x4b, 0xb5, 0x86, 0xe9, 0x6f, 0x38, 0x90, 0x0e, 0xae, 0x89, 0xd9, 0x0c,
	0xb8, 0xa6, 0x55, 0x07, 0xb9, 0x56, 0x61, 0x6f, 0x18, 0x69, 0xe3, 0x6a, 0x1c, 0x2d, 0x22, 0x4d,
	0x6b, 0x7b, 0x43, 0x83, 0xba, 0x67, 0xb8, 0x5e, 0x38, 0x75, 0x70, 0xbd, 0xf0, 0x46, 0x3b, 0xef,
	0x02, 0x75, 0xaf, 0x71, 0xeb, 0x8e, 0x4d, 0x63, 0x3e, 0xe2, 0x2c, 0xe4, 0x6a, 0x2c, 0x59, 0x42,
	0x2e, 0xb0, 0x05, 0x92, 0x25, 0xf9, 0xbf, 0x9e, 0xec, 0xc9, 0xe7, 0x07, 0xf3, 0x36, 0xb1, 0x7c,
	0xdf, 0x36, 0xea, 0xf7, 0xd6, 0x9f, 0x76, 0x69, 0x9d, 0xd9, 0xe8, 0x35, 0xb3, 0xd1, 0x5b, 0x66,
	0xa3, 0x8f, 0xcc, 0x46, 0xcf, 0x5f, 0x76, 0xe9, 0xfe, 0xef, 0xee, 0xf4, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x60, 0x58, 0x20, 0xe4, 0x49, 0x02, 0x00, 0x00,
}
