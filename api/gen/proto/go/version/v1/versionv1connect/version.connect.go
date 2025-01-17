// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: version/v1/version.proto

package versionv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/version/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// VersionName is the fully-qualified name of the Version service.
	VersionName = "version.v1.Version"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// VersionVersionProcedure is the fully-qualified name of the Version's Version RPC.
	VersionVersionProcedure = "/version.v1.Version/Version"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	versionServiceDescriptor       = v1.File_version_v1_version_proto.Services().ByName("Version")
	versionVersionMethodDescriptor = versionServiceDescriptor.Methods().ByName("Version")
)

// VersionClient is a client for the version.v1.Version service.
type VersionClient interface {
	Version(context.Context, *connect.Request[v1.VersionRequest]) (*connect.Response[v1.VersionResponse], error)
}

// NewVersionClient constructs a client for the version.v1.Version service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVersionClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) VersionClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &versionClient{
		version: connect.NewClient[v1.VersionRequest, v1.VersionResponse](
			httpClient,
			baseURL+VersionVersionProcedure,
			connect.WithSchema(versionVersionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// versionClient implements VersionClient.
type versionClient struct {
	version *connect.Client[v1.VersionRequest, v1.VersionResponse]
}

// Version calls version.v1.Version.Version.
func (c *versionClient) Version(ctx context.Context, req *connect.Request[v1.VersionRequest]) (*connect.Response[v1.VersionResponse], error) {
	return c.version.CallUnary(ctx, req)
}

// VersionHandler is an implementation of the version.v1.Version service.
type VersionHandler interface {
	Version(context.Context, *connect.Request[v1.VersionRequest]) (*connect.Response[v1.VersionResponse], error)
}

// NewVersionHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVersionHandler(svc VersionHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	versionVersionHandler := connect.NewUnaryHandler(
		VersionVersionProcedure,
		svc.Version,
		connect.WithSchema(versionVersionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/version.v1.Version/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case VersionVersionProcedure:
			versionVersionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedVersionHandler returns CodeUnimplemented from all methods.
type UnimplementedVersionHandler struct{}

func (UnimplementedVersionHandler) Version(context.Context, *connect.Request[v1.VersionRequest]) (*connect.Response[v1.VersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("version.v1.Version.Version is not implemented"))
}
