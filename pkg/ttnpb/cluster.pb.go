// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go.thethings.network/lorawan-stack/api/cluster.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PeerInfo_Role int32

const (
	PeerInfo_NONE               PeerInfo_Role = 0
	PeerInfo_IDENTITY_SERVER    PeerInfo_Role = 1
	PeerInfo_GATEWAY_SERVER     PeerInfo_Role = 2
	PeerInfo_NETWORK_SERVER     PeerInfo_Role = 3
	PeerInfo_APPLICATION_SERVER PeerInfo_Role = 4
	PeerInfo_JOIN_SERVER        PeerInfo_Role = 5
)

var PeerInfo_Role_name = map[int32]string{
	0: "NONE",
	1: "IDENTITY_SERVER",
	2: "GATEWAY_SERVER",
	3: "NETWORK_SERVER",
	4: "APPLICATION_SERVER",
	5: "JOIN_SERVER",
}
var PeerInfo_Role_value = map[string]int32{
	"NONE":               0,
	"IDENTITY_SERVER":    1,
	"GATEWAY_SERVER":     2,
	"NETWORK_SERVER":     3,
	"APPLICATION_SERVER": 4,
	"JOIN_SERVER":        5,
}

func (PeerInfo_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptorCluster, []int{0, 0} }

// PeerInfo
type PeerInfo struct {
	// Port on which the gRPC server is exposed.
	GRPCPort uint32 `protobuf:"varint,1,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// Indicates whether the gRPC server uses TLS.
	TLS bool `protobuf:"varint,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// Roles of the peer ()
	Roles []PeerInfo_Role `protobuf:"varint,3,rep,packed,name=roles,enum=ttn.lorawan.v3.PeerInfo_Role" json:"roles,omitempty"`
	// Tags of the peer
	Tags []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *PeerInfo) Reset()                    { *m = PeerInfo{} }
func (m *PeerInfo) String() string            { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()               {}
func (*PeerInfo) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{0} }

func (m *PeerInfo) GetGRPCPort() uint32 {
	if m != nil {
		return m.GRPCPort
	}
	return 0
}

func (m *PeerInfo) GetTLS() bool {
	if m != nil {
		return m.TLS
	}
	return false
}

func (m *PeerInfo) GetRoles() []PeerInfo_Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *PeerInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerInfo)(nil), "ttn.lorawan.v3.PeerInfo")
	golang_proto.RegisterType((*PeerInfo)(nil), "ttn.lorawan.v3.PeerInfo")
	proto.RegisterEnum("ttn.lorawan.v3.PeerInfo_Role", PeerInfo_Role_name, PeerInfo_Role_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.PeerInfo_Role", PeerInfo_Role_name, PeerInfo_Role_value)
}
func (x PeerInfo_Role) String() string {
	s, ok := PeerInfo_Role_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *PeerInfo) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PeerInfo)
	if !ok {
		that2, ok := that.(PeerInfo)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PeerInfo")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PeerInfo but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PeerInfo but is not nil && this == nil")
	}
	if this.GRPCPort != that1.GRPCPort {
		return fmt.Errorf("GRPCPort this(%v) Not Equal that(%v)", this.GRPCPort, that1.GRPCPort)
	}
	if this.TLS != that1.TLS {
		return fmt.Errorf("TLS this(%v) Not Equal that(%v)", this.TLS, that1.TLS)
	}
	if len(this.Roles) != len(that1.Roles) {
		return fmt.Errorf("Roles this(%v) Not Equal that(%v)", len(this.Roles), len(that1.Roles))
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return fmt.Errorf("Roles this[%v](%v) Not Equal that[%v](%v)", i, this.Roles[i], i, that1.Roles[i])
		}
	}
	if len(this.Tags) != len(that1.Tags) {
		return fmt.Errorf("Tags this(%v) Not Equal that(%v)", len(this.Tags), len(that1.Tags))
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return fmt.Errorf("Tags this[%v](%v) Not Equal that[%v](%v)", i, this.Tags[i], i, that1.Tags[i])
		}
	}
	return nil
}
func (this *PeerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PeerInfo)
	if !ok {
		that2, ok := that.(PeerInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GRPCPort != that1.GRPCPort {
		return false
	}
	if this.TLS != that1.TLS {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return false
		}
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return false
		}
	}
	return true
}
func (m *PeerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeerInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GRPCPort != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCluster(dAtA, i, uint64(m.GRPCPort))
	}
	if m.TLS {
		dAtA[i] = 0x10
		i++
		if m.TLS {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Roles) > 0 {
		dAtA2 := make([]byte, len(m.Roles)*10)
		var j1 int
		for _, num := range m.Roles {
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
		i = encodeVarintCluster(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPeerInfo(r randyCluster, easy bool) *PeerInfo {
	this := &PeerInfo{}
	this.GRPCPort = r.Uint32()
	this.TLS = bool(r.Intn(2) == 0)
	v1 := r.Intn(10)
	this.Roles = make([]PeerInfo_Role, v1)
	for i := 0; i < v1; i++ {
		this.Roles[i] = PeerInfo_Role([]int32{0, 1, 2, 3, 4, 5}[r.Intn(6)])
	}
	v2 := r.Intn(10)
	this.Tags = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Tags[i] = randStringCluster(r)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyCluster interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCluster(r randyCluster) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCluster(r randyCluster) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneCluster(r)
	}
	return string(tmps)
}
func randUnrecognizedCluster(r randyCluster, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldCluster(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldCluster(dAtA []byte, r randyCluster, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateCluster(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateCluster(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *PeerInfo) Size() (n int) {
	var l int
	_ = l
	if m.GRPCPort != 0 {
		n += 1 + sovCluster(uint64(m.GRPCPort))
	}
	if m.TLS {
		n += 2
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovCluster(uint64(e))
		}
		n += 1 + sovCluster(uint64(l)) + l
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovCluster(uint64(l))
		}
	}
	return n
}

func sovCluster(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCluster(x uint64) (n int) {
	return sovCluster((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *PeerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: PeerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GRPCPort", wireType)
			}
			m.GRPCPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GRPCPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TLS", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
			m.TLS = bool(v != 0)
		case 3:
			if wireType == 0 {
				var v PeerInfo_Role
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCluster
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (PeerInfo_Role(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCluster
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
					return ErrInvalidLengthCluster
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v PeerInfo_Role
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCluster
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (PeerInfo_Role(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
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
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
				return 0, ErrInvalidLengthCluster
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCluster
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
				next, err := skipCluster(dAtA[start:])
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
	ErrInvalidLengthCluster = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("go.thethings.network/lorawan-stack/api/cluster.proto", fileDescriptorCluster)
}
func init() {
	golang_proto.RegisterFile("go.thethings.network/lorawan-stack/api/cluster.proto", fileDescriptorCluster)
}

var fileDescriptorCluster = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4f, 0xdb, 0x40,
	0x18, 0xc5, 0xef, 0x8b, 0x4d, 0x6b, 0xae, 0x6d, 0x40, 0x57, 0xa9, 0x4a, 0x2b, 0xf5, 0x23, 0x62,
	0x4a, 0xa5, 0x62, 0x4b, 0x4d, 0xff, 0x81, 0x84, 0x5a, 0xc8, 0x2d, 0x72, 0x2c, 0x63, 0x15, 0xb5,
	0x0b, 0x72, 0x22, 0xe3, 0x44, 0x71, 0x7d, 0x96, 0x7d, 0x81, 0x95, 0x91, 0xb1, 0xdd, 0x3a, 0x56,
	0x9d, 0x18, 0x19, 0x19, 0x19, 0x19, 0x19, 0x99, 0x10, 0xbe, 0x5b, 0x18, 0x19, 0x19, 0x2b, 0x9b,
	0x26, 0x52, 0xb7, 0x6e, 0xef, 0xfd, 0xee, 0x77, 0xba, 0x93, 0x1e, 0x7d, 0x1f, 0x73, 0x53, 0x8c,
	0x23, 0x31, 0x9e, 0xa4, 0x71, 0x61, 0xa6, 0x91, 0x38, 0xe4, 0xf9, 0xd4, 0x4a, 0x78, 0x1e, 0x1e,
	0x86, 0xe9, 0x46, 0x21, 0xc2, 0xd1, 0xd4, 0x0a, 0xb3, 0x89, 0x35, 0x4a, 0x66, 0x85, 0x88, 0x72,
	0x33, 0xcb, 0xb9, 0xe0, 0xac, 0x29, 0x44, 0x6a, 0xfe, 0x95, 0xcc, 0x83, 0xee, 0xab, 0x8d, 0x78,
	0x22, 0xc6, 0xb3, 0xa1, 0x39, 0xe2, 0xdf, 0xac, 0x98, 0xc7, 0xdc, 0xaa, 0xb5, 0xe1, 0x6c, 0xbf,
	0x6e, 0x75, 0xa9, 0xd3, 0xc3, 0xf5, 0xf5, 0x1f, 0x0d, 0x6a, 0x78, 0x51, 0x94, 0x3b, 0xe9, 0x3e,
	0x67, 0x6f, 0xe8, 0x72, 0x9c, 0x67, 0xa3, 0xbd, 0x8c, 0xe7, 0xa2, 0x05, 0x6d, 0xe8, 0x3c, 0xeb,
	0x3f, 0x95, 0xd7, 0x6b, 0xc6, 0x96, 0xef, 0x6d, 0x7a, 0x3c, 0x17, 0xbe, 0x51, 0x1d, 0x57, 0x89,
	0xbd, 0xa4, 0x9a, 0x48, 0x8a, 0x56, 0xa3, 0x0d, 0x1d, 0xa3, 0xff, 0x58, 0x5e, 0xaf, 0x69, 0xc1,
	0xf6, 0x8e, 0x5f, 0x31, 0xd6, 0xa5, 0x4b, 0x39, 0x4f, 0xa2, 0xa2, 0xa5, 0xb5, 0xb5, 0x4e, 0xf3,
	0xdd, 0x6b, 0xf3, 0xdf, 0x1f, 0x9a, 0xf3, 0xe7, 0x4c, 0x9f, 0x27, 0x91, 0xff, 0xe0, 0x32, 0x46,
	0x75, 0x11, 0xc6, 0x45, 0x4b, 0x6f, 0x6b, 0x9d, 0x65, 0xbf, 0xce, 0xeb, 0x07, 0x54, 0xaf, 0x14,
	0x66, 0x50, 0xdd, 0x1d, 0xb8, 0xf6, 0x2a, 0x61, 0xcf, 0xe9, 0x8a, 0xf3, 0xc1, 0x76, 0x03, 0x27,
	0xf8, 0xb2, 0xb7, 0x63, 0xfb, 0x9f, 0x6d, 0x7f, 0x15, 0x18, 0xa3, 0xcd, 0xad, 0x5e, 0x60, 0xef,
	0xf6, 0x16, 0xac, 0x51, 0x31, 0xd7, 0x0e, 0x76, 0x07, 0xfe, 0xa7, 0x39, 0xd3, 0xd8, 0x0b, 0xca,
	0x7a, 0x9e, 0xb7, 0xed, 0x6c, 0xf6, 0x02, 0x67, 0xe0, 0xce, 0xb9, 0xce, 0x56, 0xe8, 0x93, 0x8f,
	0x03, 0x67, 0x01, 0x96, 0xfa, 0xbf, 0xe1, 0xa2, 0x44, 0xb8, 0x2c, 0x11, 0xae, 0x4a, 0x84, 0x9b,
	0x12, 0xe1, 0xb6, 0x44, 0x72, 0x57, 0x22, 0xb9, 0x2f, 0x11, 0x8e, 0x24, 0x92, 0x63, 0x89, 0xe4,
	0x44, 0x22, 0x9c, 0x4a, 0x24, 0x67, 0x12, 0xe1, 0x5c, 0x22, 0x5c, 0x48, 0x84, 0x4b, 0x89, 0x70,
	0x25, 0x91, 0xdc, 0x48, 0x84, 0x5b, 0x89, 0xe4, 0x4e, 0x22, 0xdc, 0x4b, 0x24, 0x47, 0x0a, 0xc9,
	0xb1, 0x42, 0xf8, 0xae, 0x90, 0xfc, 0x54, 0x08, 0xbf, 0x14, 0x92, 0x13, 0x85, 0xe4, 0x54, 0x21,
	0x9c, 0x29, 0x84, 0x73, 0x85, 0xf0, 0xf5, 0xed, 0x7f, 0xec, 0x9f, 0x4d, 0x63, 0x4b, 0x88, 0x34,
	0x1b, 0x0e, 0x1f, 0xd5, 0xfb, 0x75, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x19, 0x90, 0x95, 0x89,
	0x36, 0x02, 0x00, 0x00,
}
