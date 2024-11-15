// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: metastore/v1/tenant.proto

package metastorev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"
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
	// TenantServiceName is the fully-qualified name of the TenantService service.
	TenantServiceName = "metastore.v1.TenantService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TenantServiceGetTenantProcedure is the fully-qualified name of the TenantService's GetTenant RPC.
	TenantServiceGetTenantProcedure = "/metastore.v1.TenantService/GetTenant"
	// TenantServiceDeleteTenantProcedure is the fully-qualified name of the TenantService's
	// DeleteTenant RPC.
	TenantServiceDeleteTenantProcedure = "/metastore.v1.TenantService/DeleteTenant"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tenantServiceServiceDescriptor            = v1.File_metastore_v1_tenant_proto.Services().ByName("TenantService")
	tenantServiceGetTenantMethodDescriptor    = tenantServiceServiceDescriptor.Methods().ByName("GetTenant")
	tenantServiceDeleteTenantMethodDescriptor = tenantServiceServiceDescriptor.Methods().ByName("DeleteTenant")
)

// TenantServiceClient is a client for the metastore.v1.TenantService service.
type TenantServiceClient interface {
	GetTenant(context.Context, *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error)
	DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error)
}

// NewTenantServiceClient constructs a client for the metastore.v1.TenantService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTenantServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TenantServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tenantServiceClient{
		getTenant: connect.NewClient[v1.GetTenantRequest, v1.GetTenantResponse](
			httpClient,
			baseURL+TenantServiceGetTenantProcedure,
			connect.WithSchema(tenantServiceGetTenantMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTenant: connect.NewClient[v1.DeleteTenantRequest, v1.DeleteTenantResponse](
			httpClient,
			baseURL+TenantServiceDeleteTenantProcedure,
			connect.WithSchema(tenantServiceDeleteTenantMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tenantServiceClient implements TenantServiceClient.
type tenantServiceClient struct {
	getTenant    *connect.Client[v1.GetTenantRequest, v1.GetTenantResponse]
	deleteTenant *connect.Client[v1.DeleteTenantRequest, v1.DeleteTenantResponse]
}

// GetTenant calls metastore.v1.TenantService.GetTenant.
func (c *tenantServiceClient) GetTenant(ctx context.Context, req *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error) {
	return c.getTenant.CallUnary(ctx, req)
}

// DeleteTenant calls metastore.v1.TenantService.DeleteTenant.
func (c *tenantServiceClient) DeleteTenant(ctx context.Context, req *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return c.deleteTenant.CallUnary(ctx, req)
}

// TenantServiceHandler is an implementation of the metastore.v1.TenantService service.
type TenantServiceHandler interface {
	GetTenant(context.Context, *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error)
	DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error)
}

// NewTenantServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTenantServiceHandler(svc TenantServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tenantServiceGetTenantHandler := connect.NewUnaryHandler(
		TenantServiceGetTenantProcedure,
		svc.GetTenant,
		connect.WithSchema(tenantServiceGetTenantMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tenantServiceDeleteTenantHandler := connect.NewUnaryHandler(
		TenantServiceDeleteTenantProcedure,
		svc.DeleteTenant,
		connect.WithSchema(tenantServiceDeleteTenantMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/metastore.v1.TenantService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TenantServiceGetTenantProcedure:
			tenantServiceGetTenantHandler.ServeHTTP(w, r)
		case TenantServiceDeleteTenantProcedure:
			tenantServiceDeleteTenantHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTenantServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTenantServiceHandler struct{}

func (UnimplementedTenantServiceHandler) GetTenant(context.Context, *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metastore.v1.TenantService.GetTenant is not implemented"))
}

func (UnimplementedTenantServiceHandler) DeleteTenant(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("metastore.v1.TenantService.DeleteTenant is not implemented"))
}
