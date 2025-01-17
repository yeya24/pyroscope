// Code generated by protoc-gen-connect-go-mux. DO NOT EDIT.
//
// Source: metastore/v1/tenant.proto

package metastorev1connect

import (
	connect "connectrpc.com/connect"
	mux "github.com/gorilla/mux"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

// RegisterTenantServiceHandler register an HTTP handler to a mux.Router from the service
// implementation.
func RegisterTenantServiceHandler(mux *mux.Router, svc TenantServiceHandler, opts ...connect.HandlerOption) {
	mux.Handle("/metastore.v1.TenantService/GetTenant", connect.NewUnaryHandler(
		"/metastore.v1.TenantService/GetTenant",
		svc.GetTenant,
		opts...,
	))
	mux.Handle("/metastore.v1.TenantService/DeleteTenant", connect.NewUnaryHandler(
		"/metastore.v1.TenantService/DeleteTenant",
		svc.DeleteTenant,
		opts...,
	))
}
