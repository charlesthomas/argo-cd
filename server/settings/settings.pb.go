// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/settings/settings.proto

package settings // import "github.com/argoproj/argo-cd/server/settings"

/*
	Settings Service

	Settings Service API retrieves Argo CD settings
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

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

// SettingsQuery is a query for Argo CD settings
type SettingsQuery struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingsQuery) Reset()         { *m = SettingsQuery{} }
func (m *SettingsQuery) String() string { return proto.CompactTextString(m) }
func (*SettingsQuery) ProtoMessage()    {}
func (*SettingsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_settings_90d4947f5a5e2583, []int{0}
}
func (m *SettingsQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettingsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettingsQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SettingsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsQuery.Merge(dst, src)
}
func (m *SettingsQuery) XXX_Size() int {
	return m.Size()
}
func (m *SettingsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsQuery proto.InternalMessageInfo

type Settings struct {
	URL                  string      `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	DexConfig            *DexConfig  `protobuf:"bytes,2,opt,name=dexConfig" json:"dexConfig,omitempty"`
	OIDCConfig           *OIDCConfig `protobuf:"bytes,3,opt,name=oidcConfig" json:"oidcConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_settings_90d4947f5a5e2583, []int{1}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(dst, src)
}
func (m *Settings) XXX_Size() int {
	return m.Size()
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Settings) GetDexConfig() *DexConfig {
	if m != nil {
		return m.DexConfig
	}
	return nil
}

func (m *Settings) GetOIDCConfig() *OIDCConfig {
	if m != nil {
		return m.OIDCConfig
	}
	return nil
}

type DexConfig struct {
	Connectors           []*Connector `protobuf:"bytes,1,rep,name=connectors" json:"connectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DexConfig) Reset()         { *m = DexConfig{} }
func (m *DexConfig) String() string { return proto.CompactTextString(m) }
func (*DexConfig) ProtoMessage()    {}
func (*DexConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_settings_90d4947f5a5e2583, []int{2}
}
func (m *DexConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DexConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DexConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DexConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DexConfig.Merge(dst, src)
}
func (m *DexConfig) XXX_Size() int {
	return m.Size()
}
func (m *DexConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DexConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DexConfig proto.InternalMessageInfo

func (m *DexConfig) GetConnectors() []*Connector {
	if m != nil {
		return m.Connectors
	}
	return nil
}

type Connector struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connector) Reset()         { *m = Connector{} }
func (m *Connector) String() string { return proto.CompactTextString(m) }
func (*Connector) ProtoMessage()    {}
func (*Connector) Descriptor() ([]byte, []int) {
	return fileDescriptor_settings_90d4947f5a5e2583, []int{3}
}
func (m *Connector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Connector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Connector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Connector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connector.Merge(dst, src)
}
func (m *Connector) XXX_Size() int {
	return m.Size()
}
func (m *Connector) XXX_DiscardUnknown() {
	xxx_messageInfo_Connector.DiscardUnknown(m)
}

var xxx_messageInfo_Connector proto.InternalMessageInfo

func (m *Connector) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Connector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type OIDCConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Issuer               string   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	ClientID             string   `protobuf:"bytes,3,opt,name=clientID,proto3" json:"clientID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OIDCConfig) Reset()         { *m = OIDCConfig{} }
func (m *OIDCConfig) String() string { return proto.CompactTextString(m) }
func (*OIDCConfig) ProtoMessage()    {}
func (*OIDCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_settings_90d4947f5a5e2583, []int{4}
}
func (m *OIDCConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OIDCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OIDCConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *OIDCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OIDCConfig.Merge(dst, src)
}
func (m *OIDCConfig) XXX_Size() int {
	return m.Size()
}
func (m *OIDCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OIDCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OIDCConfig proto.InternalMessageInfo

func (m *OIDCConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OIDCConfig) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *OIDCConfig) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func init() {
	proto.RegisterType((*SettingsQuery)(nil), "cluster.SettingsQuery")
	proto.RegisterType((*Settings)(nil), "cluster.Settings")
	proto.RegisterType((*DexConfig)(nil), "cluster.DexConfig")
	proto.RegisterType((*Connector)(nil), "cluster.Connector")
	proto.RegisterType((*OIDCConfig)(nil), "cluster.OIDCConfig")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SettingsService service

type SettingsServiceClient interface {
	// Get returns Argo CD settings
	Get(ctx context.Context, in *SettingsQuery, opts ...grpc.CallOption) (*Settings, error)
}

type settingsServiceClient struct {
	cc *grpc.ClientConn
}

func NewSettingsServiceClient(cc *grpc.ClientConn) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) Get(ctx context.Context, in *SettingsQuery, opts ...grpc.CallOption) (*Settings, error) {
	out := new(Settings)
	err := c.cc.Invoke(ctx, "/cluster.SettingsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SettingsService service

type SettingsServiceServer interface {
	// Get returns Argo CD settings
	Get(context.Context, *SettingsQuery) (*Settings, error)
}

func RegisterSettingsServiceServer(s *grpc.Server, srv SettingsServiceServer) {
	s.RegisterService(&_SettingsService_serviceDesc, srv)
}

func _SettingsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.SettingsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).Get(ctx, req.(*SettingsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _SettingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SettingsService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/settings/settings.proto",
}

func (m *SettingsQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettingsQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Settings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Settings) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.URL) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSettings(dAtA, i, uint64(len(m.URL)))
		i += copy(dAtA[i:], m.URL)
	}
	if m.DexConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSettings(dAtA, i, uint64(m.DexConfig.Size()))
		n1, err := m.DexConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.OIDCConfig != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSettings(dAtA, i, uint64(m.OIDCConfig.Size()))
		n2, err := m.OIDCConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DexConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DexConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Connectors) > 0 {
		for _, msg := range m.Connectors {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSettings(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Connector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Connector) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSettings(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSettings(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *OIDCConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OIDCConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSettings(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Issuer) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSettings(dAtA, i, uint64(len(m.Issuer)))
		i += copy(dAtA[i:], m.Issuer)
	}
	if len(m.ClientID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSettings(dAtA, i, uint64(len(m.ClientID)))
		i += copy(dAtA[i:], m.ClientID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSettings(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SettingsQuery) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Settings) Size() (n int) {
	var l int
	_ = l
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.DexConfig != nil {
		l = m.DexConfig.Size()
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.OIDCConfig != nil {
		l = m.OIDCConfig.Size()
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DexConfig) Size() (n int) {
	var l int
	_ = l
	if len(m.Connectors) > 0 {
		for _, e := range m.Connectors {
			l = e.Size()
			n += 1 + l + sovSettings(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Connector) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OIDCConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovSettings(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSettings(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSettings(x uint64) (n int) {
	return sovSettings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SettingsQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: SettingsQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettingsQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
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
func (m *Settings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: Settings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Settings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DexConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DexConfig == nil {
				m.DexConfig = &DexConfig{}
			}
			if err := m.DexConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OIDCConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OIDCConfig == nil {
				m.OIDCConfig = &OIDCConfig{}
			}
			if err := m.OIDCConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
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
func (m *DexConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: DexConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DexConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connectors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Connectors = append(m.Connectors, &Connector{})
			if err := m.Connectors[len(m.Connectors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
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
func (m *Connector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: Connector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Connector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
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
func (m *OIDCConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettings
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
			return fmt.Errorf("proto: OIDCConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OIDCConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettings
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
				return ErrInvalidLengthSettings
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSettings
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
func skipSettings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSettings
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
					return 0, ErrIntOverflowSettings
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
					return 0, ErrIntOverflowSettings
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
				return 0, ErrInvalidLengthSettings
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSettings
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
				next, err := skipSettings(dAtA[start:])
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
	ErrInvalidLengthSettings = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSettings   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("server/settings/settings.proto", fileDescriptor_settings_90d4947f5a5e2583)
}

var fileDescriptor_settings_90d4947f5a5e2583 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x8a, 0xdb, 0x30,
	0x18, 0x44, 0x71, 0x49, 0xe2, 0xaf, 0x3f, 0x69, 0xd5, 0x12, 0xdc, 0x50, 0x9c, 0xe0, 0x53, 0xa0,
	0xd4, 0x6e, 0x93, 0x53, 0x4f, 0x05, 0x3b, 0x50, 0x52, 0x0a, 0xa5, 0x0a, 0xbd, 0x14, 0x7a, 0x70,
	0x14, 0xd5, 0x55, 0x71, 0xa4, 0x20, 0xcb, 0xa1, 0xb9, 0xf6, 0x15, 0xf6, 0xba, 0x0f, 0xb4, 0xc7,
	0x85, 0xbd, 0x87, 0xc5, 0xec, 0x83, 0x2c, 0x51, 0x6c, 0x27, 0xd9, 0xdd, 0xdb, 0x68, 0x46, 0x23,
	0xbe, 0xd1, 0x37, 0xe0, 0x66, 0x4c, 0xad, 0x99, 0x0a, 0x32, 0xa6, 0x35, 0x17, 0x49, 0x56, 0x03,
	0x7f, 0xa5, 0xa4, 0x96, 0xb8, 0x45, 0xd3, 0x3c, 0xd3, 0x4c, 0xf5, 0x5e, 0x25, 0x32, 0x91, 0x86,
	0x0b, 0x76, 0x68, 0x2f, 0xf7, 0xde, 0x24, 0x52, 0x26, 0x29, 0x0b, 0xe2, 0x15, 0x0f, 0x62, 0x21,
	0xa4, 0x8e, 0x35, 0x97, 0xa2, 0x34, 0x7b, 0x1d, 0x78, 0x3a, 0x2b, 0x9f, 0xfb, 0x9e, 0x33, 0xb5,
	0xf1, 0xce, 0x11, 0xb4, 0x2b, 0x06, 0xbf, 0x06, 0x2b, 0x57, 0xa9, 0x83, 0x06, 0x68, 0x68, 0x87,
	0xad, 0x62, 0xdb, 0xb7, 0x7e, 0x90, 0xaf, 0x64, 0xc7, 0xe1, 0xf7, 0x60, 0x2f, 0xd8, 0xbf, 0x48,
	0x8a, 0xdf, 0x3c, 0x71, 0x1a, 0x03, 0x34, 0x7c, 0x3c, 0xc2, 0x7e, 0x39, 0x89, 0x3f, 0xa9, 0x14,
	0x72, 0xb8, 0x84, 0x23, 0x00, 0xc9, 0x17, 0xb4, 0xb4, 0x58, 0xc6, 0xf2, 0xb2, 0xb6, 0x7c, 0x9b,
	0x4e, 0xa2, 0xbd, 0x14, 0x3e, 0x2b, 0xb6, 0x7d, 0x38, 0x9c, 0xc9, 0x91, 0xcd, 0xfb, 0x04, 0x76,
	0xfd, 0x38, 0x1e, 0x01, 0x50, 0x29, 0x04, 0xa3, 0x5a, 0xaa, 0xcc, 0x41, 0x03, 0xeb, 0x64, 0x88,
	0xa8, 0x92, 0xc8, 0xd1, 0x2d, 0x6f, 0x0c, 0x76, 0x2d, 0x60, 0x0c, 0x8f, 0x44, 0xbc, 0x64, 0xfb,
	0x80, 0xc4, 0xe0, 0x1d, 0xa7, 0x37, 0x2b, 0x66, 0x32, 0xd9, 0xc4, 0x60, 0x6f, 0x0e, 0x47, 0xf3,
	0x3c, 0xe8, 0xea, 0x42, 0x93, 0x67, 0x59, 0xce, 0x54, 0xe9, 0x2b, 0x4f, 0x78, 0x08, 0x6d, 0x9a,
	0x72, 0x26, 0xf4, 0x74, 0x62, 0x22, 0xdb, 0xe1, 0x93, 0x62, 0xdb, 0x6f, 0x47, 0x25, 0x47, 0x6a,
	0x75, 0xf4, 0x0b, 0x3a, 0xd5, 0xbf, 0xcf, 0x98, 0x5a, 0x73, 0xca, 0xf0, 0x17, 0xb0, 0x3e, 0x33,
	0x8d, 0xbb, 0x75, 0xa4, 0x93, 0x55, 0xf5, 0x5e, 0xdc, 0xe3, 0x3d, 0xe7, 0xff, 0xd5, 0xcd, 0x59,
	0x03, 0xe3, 0xe7, 0x66, 0xdd, 0xeb, 0x0f, 0x75, 0x57, 0xc2, 0x8f, 0x17, 0x85, 0x8b, 0x2e, 0x0b,
	0x17, 0x5d, 0x17, 0x2e, 0xfa, 0xf9, 0x36, 0xe1, 0xfa, 0x4f, 0x3e, 0xf7, 0xa9, 0x5c, 0x06, 0xb1,
	0x32, 0xad, 0xf9, 0x6b, 0xc0, 0x3b, 0xba, 0x08, 0xee, 0xf4, 0x6d, 0xde, 0x34, 0x55, 0x19, 0xdf,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xa2, 0x52, 0xc6, 0x89, 0x02, 0x00, 0x00,
}
