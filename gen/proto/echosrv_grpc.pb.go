// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

var echoServiceEchoStreamDesc = &grpc.StreamDesc{
	StreamName: "Echo",
}

func (c *echoServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/echo.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceService is the service API for EchoService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterEchoServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type EchoServiceService struct {
	Echo func(context.Context, *StringMessage) (*StringMessage, error)
}

func (s *EchoServiceService) echo(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/echo.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterEchoServiceService registers a service implementation with a gRPC server.
func RegisterEchoServiceService(s grpc.ServiceRegistrar, srv *EchoServiceService) {
	srvCopy := *srv
	if srvCopy.Echo == nil {
		srvCopy.Echo = func(context.Context, *StringMessage) (*StringMessage, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "echo.EchoService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Echo",
				Handler:    srvCopy.echo,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "proto/echosrv.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewEchoServiceService creates a new EchoServiceService containing the
// implemented methods of the EchoService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewEchoServiceService(s interface{}) *EchoServiceService {
	ns := &EchoServiceService{}
	if h, ok := s.(interface {
		Echo(context.Context, *StringMessage) (*StringMessage, error)
	}); ok {
		ns.Echo = h.Echo
	}
	return ns
}

// UnstableEchoServiceService is the service API for EchoService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableEchoServiceService interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
}
