// Code generated by protoc-gen-go.
// source: eventmaster.proto
// DO NOT EDIT!

/*
Package eventmaster is a generated protocol buffer package.

It is generated from these files:
	eventmaster.proto

It has these top-level messages:
	Event
	WriteResponse
	Query
	HealthcheckRequest
	HealthcheckResponse
*/
package eventmaster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	ParentEventId string   `protobuf:"bytes,1,opt,name=parent_event_id,json=parentEventId" json:"parent_event_id,omitempty"`
	EventTime     int64    `protobuf:"varint,2,opt,name=event_time,json=eventTime" json:"event_time,omitempty"`
	Dc            string   `protobuf:"bytes,3,opt,name=dc" json:"dc,omitempty"`
	TopicName     string   `protobuf:"bytes,4,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	TagSet        []string `protobuf:"bytes,5,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	Host          string   `protobuf:"bytes,6,opt,name=host" json:"host,omitempty"`
	TargetHostSet []string `protobuf:"bytes,7,rep,name=target_host_set,json=targetHostSet" json:"target_host_set,omitempty"`
	User          string   `protobuf:"bytes,8,opt,name=user" json:"user,omitempty"`
	DataJson      string   `protobuf:"bytes,9,opt,name=data_json,json=dataJson" json:"data_json,omitempty"`
	EventType     int32    `protobuf:"varint,10,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetParentEventId() string {
	if m != nil {
		return m.ParentEventId
	}
	return ""
}

func (m *Event) GetEventTime() int64 {
	if m != nil {
		return m.EventTime
	}
	return 0
}

func (m *Event) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *Event) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *Event) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *Event) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Event) GetTargetHostSet() []string {
	if m != nil {
		return m.TargetHostSet
	}
	return nil
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetDataJson() string {
	if m != nil {
		return m.DataJson
	}
	return ""
}

func (m *Event) GetEventType() int32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

type WriteResponse struct {
	Errcode int32  `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	EventId string `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WriteResponse) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *WriteResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *WriteResponse) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

type Query struct {
	Dc            []string `protobuf:"bytes,1,rep,name=dc" json:"dc,omitempty"`
	Host          []string `protobuf:"bytes,2,rep,name=host" json:"host,omitempty"`
	TargetHost    []string `protobuf:"bytes,3,rep,name=target_host,json=targetHost" json:"target_host,omitempty"`
	TopicName     []string `protobuf:"bytes,4,rep,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	TagSet        []string `protobuf:"bytes,5,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	StartTime     int64    `protobuf:"varint,6,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime       int64    `protobuf:"varint,7,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	SortField     string   `protobuf:"bytes,8,opt,name=sort_field,json=sortField" json:"sort_field,omitempty"`
	SortAscending bool     `protobuf:"varint,9,opt,name=sort_ascending,json=sortAscending" json:"sort_ascending,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Query) GetDc() []string {
	if m != nil {
		return m.Dc
	}
	return nil
}

func (m *Query) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Query) GetTargetHost() []string {
	if m != nil {
		return m.TargetHost
	}
	return nil
}

func (m *Query) GetTopicName() []string {
	if m != nil {
		return m.TopicName
	}
	return nil
}

func (m *Query) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *Query) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Query) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Query) GetSortField() string {
	if m != nil {
		return m.SortField
	}
	return ""
}

func (m *Query) GetSortAscending() bool {
	if m != nil {
		return m.SortAscending
	}
	return false
}

type HealthcheckRequest struct {
}

func (m *HealthcheckRequest) Reset()                    { *m = HealthcheckRequest{} }
func (m *HealthcheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthcheckRequest) ProtoMessage()               {}
func (*HealthcheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type HealthcheckResponse struct {
	Response string `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *HealthcheckResponse) Reset()                    { *m = HealthcheckResponse{} }
func (m *HealthcheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthcheckResponse) ProtoMessage()               {}
func (*HealthcheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HealthcheckResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "eventmaster.Event")
	proto.RegisterType((*WriteResponse)(nil), "eventmaster.WriteResponse")
	proto.RegisterType((*Query)(nil), "eventmaster.Query")
	proto.RegisterType((*HealthcheckRequest)(nil), "eventmaster.HealthcheckRequest")
	proto.RegisterType((*HealthcheckResponse)(nil), "eventmaster.HealthcheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventMaster service

type EventMasterClient interface {
	FireEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*WriteResponse, error)
	GetEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (EventMaster_GetEventsClient, error)
	Healthcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
}

type eventMasterClient struct {
	cc *grpc.ClientConn
}

func NewEventMasterClient(cc *grpc.ClientConn) EventMasterClient {
	return &eventMasterClient{cc}
}

func (c *eventMasterClient) FireEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/FireEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) GetEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (EventMaster_GetEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventMaster_serviceDesc.Streams[0], c.cc, "/eventmaster.EventMaster/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventMasterGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventMaster_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventMasterGetEventsClient struct {
	grpc.ClientStream
}

func (x *eventMasterGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventMasterClient) Healthcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/Healthcheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventMaster service

type EventMasterServer interface {
	FireEvent(context.Context, *Event) (*WriteResponse, error)
	GetEvents(*Query, EventMaster_GetEventsServer) error
	Healthcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
}

func RegisterEventMasterServer(s *grpc.Server, srv EventMasterServer) {
	s.RegisterService(&_EventMaster_serviceDesc, srv)
}

func _EventMaster_FireEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).FireEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/FireEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).FireEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventMasterServer).GetEvents(m, &eventMasterGetEventsServer{stream})
}

type EventMaster_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventMasterGetEventsServer struct {
	grpc.ServerStream
}

func (x *eventMasterGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventMaster_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).Healthcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/Healthcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).Healthcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventmaster.EventMaster",
	HandlerType: (*EventMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FireEvent",
			Handler:    _EventMaster_FireEvent_Handler,
		},
		{
			MethodName: "Healthcheck",
			Handler:    _EventMaster_Healthcheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _EventMaster_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eventmaster.proto",
}

func init() { proto.RegisterFile("eventmaster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0xc7, 0x37, 0x33, 0x3b, 0xf9, 0xa8, 0x61, 0x56, 0x2c, 0x45, 0x63, 0x44, 0x76, 0x08, 0x28,
	0x9e, 0x16, 0x3f, 0x0e, 0x9e, 0x3c, 0x78, 0x70, 0x5d, 0x05, 0x05, 0xa3, 0xe0, 0x45, 0x08, 0x31,
	0x29, 0x67, 0xa2, 0x3b, 0x49, 0xec, 0xee, 0x11, 0xf6, 0xee, 0x93, 0xf8, 0x56, 0xbe, 0x8d, 0xd5,
	0xd5, 0xc9, 0x9a, 0x71, 0x15, 0xf6, 0x56, 0xf5, 0xab, 0xfa, 0x87, 0xae, 0x7f, 0x55, 0xe0, 0x2a,
	0x7d, 0xa7, 0xc6, 0x6c, 0x0a, 0x6d, 0x48, 0x1d, 0x75, 0xaa, 0x35, 0x2d, 0xce, 0x47, 0x28, 0xfd,
	0x39, 0x81, 0xd9, 0x73, 0x9b, 0xe3, 0x3d, 0xb8, 0xd2, 0x15, 0x8a, 0xa3, 0x5c, 0xea, 0x79, 0x5d,
	0xc5, 0xde, 0xd2, 0xbb, 0x1f, 0x65, 0x0b, 0x87, 0xa5, 0xeb, 0x65, 0x85, 0x77, 0x00, 0x5c, 0x83,
	0xa9, 0x37, 0x14, 0x4f, 0xb8, 0x65, 0x9a, 0x45, 0x42, 0xde, 0x33, 0xc0, 0x03, 0x98, 0x54, 0x65,
	0x3c, 0x15, 0x25, 0x47, 0xb6, 0xdd, 0xb4, 0x5d, 0x5d, 0xe6, 0x4d, 0xc1, 0xed, 0xfb, 0xc2, 0x23,
	0x21, 0x6f, 0x18, 0xe0, 0x4d, 0x08, 0x4c, 0xb1, 0xca, 0x35, 0x99, 0x78, 0xb6, 0x9c, 0x72, 0xcd,
	0xe7, 0xf4, 0x1d, 0x19, 0x44, 0xd8, 0x5f, 0xb7, 0xda, 0xc4, 0xbe, 0x28, 0x24, 0xb6, 0x4f, 0x34,
	0x85, 0x5a, 0x91, 0xc9, 0x6d, 0x2a, 0xa2, 0x40, 0x44, 0x0b, 0x87, 0x4f, 0x98, 0xf6, 0xda, 0xad,
	0x26, 0x15, 0x87, 0x4e, 0x6b, 0x63, 0xbc, 0x0d, 0x51, 0x55, 0x98, 0x22, 0xff, 0xa2, 0xdb, 0x26,
	0x8e, 0xa4, 0x10, 0x5a, 0xf0, 0x8a, 0xf3, 0xd1, 0x4c, 0x67, 0x1d, 0xc5, 0xc0, 0xd5, 0xd9, 0x30,
	0x13, 0x83, 0xf4, 0x23, 0x2c, 0x3e, 0xa8, 0xda, 0x50, 0x46, 0xba, 0x6b, 0x1b, 0x4d, 0x18, 0x43,
	0x40, 0x4a, 0x95, 0x6d, 0x45, 0xe2, 0xd1, 0x2c, 0x1b, 0x52, 0xbc, 0x01, 0x3e, 0x87, 0x1b, 0xbd,
	0x12, 0x67, 0x78, 0x1c, 0x97, 0xe1, 0x2d, 0x08, 0xcf, 0x6d, 0x75, 0xe6, 0x04, 0xe4, 0x0c, 0x4d,
	0x7f, 0xf0, 0x0a, 0xde, 0x6e, 0x49, 0x9d, 0xf5, 0xde, 0x79, 0x32, 0x92, 0xf5, 0x6e, 0xf0, 0x60,
	0x22, 0xc4, 0x79, 0x70, 0x08, 0xf3, 0x91, 0x07, 0xfc, 0x2d, 0x5b, 0x82, 0x3f, 0xf3, 0x5f, 0x30,
	0x7c, 0x7a, 0x49, 0xc3, 0x59, 0xa7, 0xf9, 0x33, 0xfd, 0x5e, 0x7d, 0xb7, 0x57, 0x21, 0xb2, 0x57,
	0x3b, 0x40, 0x53, 0xb9, 0x62, 0x20, 0xc5, 0x80, 0x73, 0x29, 0x59, 0x65, 0xcb, 0xc2, 0xcf, 0x35,
	0x9d, 0x56, 0xbd, 0xe9, 0x91, 0x25, 0xc7, 0x16, 0xe0, 0x5d, 0x38, 0x90, 0x72, 0xa1, 0x4b, 0x56,
	0xd4, 0xcd, 0x4a, 0xec, 0x0f, 0xb3, 0x85, 0xa5, 0xcf, 0x06, 0x98, 0x5e, 0x07, 0x3c, 0xa1, 0xe2,
	0xd4, 0xac, 0xcb, 0x35, 0x95, 0x5f, 0x33, 0xfa, 0xb6, 0x25, 0x6d, 0xd2, 0x87, 0x70, 0x6d, 0x87,
	0xf6, 0x0b, 0x48, 0x20, 0x54, 0x7d, 0xdc, 0x5f, 0xe9, 0x79, 0xfe, 0xe8, 0x97, 0x07, 0x73, 0x39,
	0xd6, 0xd7, 0x72, 0xe2, 0xf8, 0x14, 0xa2, 0xe3, 0x5a, 0x91, 0xbb, 0x72, 0x3c, 0x1a, 0xff, 0x10,
	0xc2, 0x92, 0x64, 0x87, 0xed, 0x6c, 0x3a, 0xdd, 0xc3, 0x27, 0x10, 0xbd, 0x20, 0x77, 0xfd, 0xfa,
	0x2f, 0xb9, 0x6c, 0x2d, 0xf9, 0xc7, 0x27, 0xd3, 0xbd, 0x07, 0x1e, 0x66, 0x30, 0x1f, 0x3d, 0x1d,
	0x0f, 0x77, 0xda, 0x2e, 0x8e, 0x9a, 0x2c, 0xff, 0xdf, 0x30, 0x3c, 0xe6, 0x93, 0x2f, 0xbf, 0xf0,
	0xe3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x22, 0x42, 0xe2, 0xd7, 0x03, 0x00, 0x00,
}
