// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: inventory/v1/inventory.proto

package inventoryv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/suzushin54/actor-based-inventory/gen/inventory/v1"
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
	// InventoryServiceName is the fully-qualified name of the InventoryService service.
	InventoryServiceName = "inventory.v1.InventoryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// InventoryServiceCreateInventoryProcedure is the fully-qualified name of the InventoryService's
	// CreateInventory RPC.
	InventoryServiceCreateInventoryProcedure = "/inventory.v1.InventoryService/CreateInventory"
	// InventoryServiceGetInventoryProcedure is the fully-qualified name of the InventoryService's
	// GetInventory RPC.
	InventoryServiceGetInventoryProcedure = "/inventory.v1.InventoryService/GetInventory"
	// InventoryServiceUpdateInventoryProcedure is the fully-qualified name of the InventoryService's
	// UpdateInventory RPC.
	InventoryServiceUpdateInventoryProcedure = "/inventory.v1.InventoryService/UpdateInventory"
	// InventoryServiceDeleteInventoryProcedure is the fully-qualified name of the InventoryService's
	// DeleteInventory RPC.
	InventoryServiceDeleteInventoryProcedure = "/inventory.v1.InventoryService/DeleteInventory"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	inventoryServiceServiceDescriptor               = v1.File_inventory_v1_inventory_proto.Services().ByName("InventoryService")
	inventoryServiceCreateInventoryMethodDescriptor = inventoryServiceServiceDescriptor.Methods().ByName("CreateInventory")
	inventoryServiceGetInventoryMethodDescriptor    = inventoryServiceServiceDescriptor.Methods().ByName("GetInventory")
	inventoryServiceUpdateInventoryMethodDescriptor = inventoryServiceServiceDescriptor.Methods().ByName("UpdateInventory")
	inventoryServiceDeleteInventoryMethodDescriptor = inventoryServiceServiceDescriptor.Methods().ByName("DeleteInventory")
)

// InventoryServiceClient is a client for the inventory.v1.InventoryService service.
type InventoryServiceClient interface {
	CreateInventory(context.Context, *connect.Request[v1.CreateInventoryRequest]) (*connect.Response[v1.CreateInventoryResponse], error)
	GetInventory(context.Context, *connect.Request[v1.GetInventoryRequest]) (*connect.Response[v1.GetInventoryResponse], error)
	UpdateInventory(context.Context, *connect.Request[v1.UpdateInventoryRequest]) (*connect.Response[v1.UpdateInventoryResponse], error)
	DeleteInventory(context.Context, *connect.Request[v1.DeleteInventoryRequest]) (*connect.Response[v1.DeleteInventoryResponse], error)
}

// NewInventoryServiceClient constructs a client for the inventory.v1.InventoryService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInventoryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) InventoryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &inventoryServiceClient{
		createInventory: connect.NewClient[v1.CreateInventoryRequest, v1.CreateInventoryResponse](
			httpClient,
			baseURL+InventoryServiceCreateInventoryProcedure,
			connect.WithSchema(inventoryServiceCreateInventoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getInventory: connect.NewClient[v1.GetInventoryRequest, v1.GetInventoryResponse](
			httpClient,
			baseURL+InventoryServiceGetInventoryProcedure,
			connect.WithSchema(inventoryServiceGetInventoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateInventory: connect.NewClient[v1.UpdateInventoryRequest, v1.UpdateInventoryResponse](
			httpClient,
			baseURL+InventoryServiceUpdateInventoryProcedure,
			connect.WithSchema(inventoryServiceUpdateInventoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteInventory: connect.NewClient[v1.DeleteInventoryRequest, v1.DeleteInventoryResponse](
			httpClient,
			baseURL+InventoryServiceDeleteInventoryProcedure,
			connect.WithSchema(inventoryServiceDeleteInventoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// inventoryServiceClient implements InventoryServiceClient.
type inventoryServiceClient struct {
	createInventory *connect.Client[v1.CreateInventoryRequest, v1.CreateInventoryResponse]
	getInventory    *connect.Client[v1.GetInventoryRequest, v1.GetInventoryResponse]
	updateInventory *connect.Client[v1.UpdateInventoryRequest, v1.UpdateInventoryResponse]
	deleteInventory *connect.Client[v1.DeleteInventoryRequest, v1.DeleteInventoryResponse]
}

// CreateInventory calls inventory.v1.InventoryService.CreateInventory.
func (c *inventoryServiceClient) CreateInventory(ctx context.Context, req *connect.Request[v1.CreateInventoryRequest]) (*connect.Response[v1.CreateInventoryResponse], error) {
	return c.createInventory.CallUnary(ctx, req)
}

// GetInventory calls inventory.v1.InventoryService.GetInventory.
func (c *inventoryServiceClient) GetInventory(ctx context.Context, req *connect.Request[v1.GetInventoryRequest]) (*connect.Response[v1.GetInventoryResponse], error) {
	return c.getInventory.CallUnary(ctx, req)
}

// UpdateInventory calls inventory.v1.InventoryService.UpdateInventory.
func (c *inventoryServiceClient) UpdateInventory(ctx context.Context, req *connect.Request[v1.UpdateInventoryRequest]) (*connect.Response[v1.UpdateInventoryResponse], error) {
	return c.updateInventory.CallUnary(ctx, req)
}

// DeleteInventory calls inventory.v1.InventoryService.DeleteInventory.
func (c *inventoryServiceClient) DeleteInventory(ctx context.Context, req *connect.Request[v1.DeleteInventoryRequest]) (*connect.Response[v1.DeleteInventoryResponse], error) {
	return c.deleteInventory.CallUnary(ctx, req)
}

// InventoryServiceHandler is an implementation of the inventory.v1.InventoryService service.
type InventoryServiceHandler interface {
	CreateInventory(context.Context, *connect.Request[v1.CreateInventoryRequest]) (*connect.Response[v1.CreateInventoryResponse], error)
	GetInventory(context.Context, *connect.Request[v1.GetInventoryRequest]) (*connect.Response[v1.GetInventoryResponse], error)
	UpdateInventory(context.Context, *connect.Request[v1.UpdateInventoryRequest]) (*connect.Response[v1.UpdateInventoryResponse], error)
	DeleteInventory(context.Context, *connect.Request[v1.DeleteInventoryRequest]) (*connect.Response[v1.DeleteInventoryResponse], error)
}

// NewInventoryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInventoryServiceHandler(svc InventoryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	inventoryServiceCreateInventoryHandler := connect.NewUnaryHandler(
		InventoryServiceCreateInventoryProcedure,
		svc.CreateInventory,
		connect.WithSchema(inventoryServiceCreateInventoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceGetInventoryHandler := connect.NewUnaryHandler(
		InventoryServiceGetInventoryProcedure,
		svc.GetInventory,
		connect.WithSchema(inventoryServiceGetInventoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceUpdateInventoryHandler := connect.NewUnaryHandler(
		InventoryServiceUpdateInventoryProcedure,
		svc.UpdateInventory,
		connect.WithSchema(inventoryServiceUpdateInventoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceDeleteInventoryHandler := connect.NewUnaryHandler(
		InventoryServiceDeleteInventoryProcedure,
		svc.DeleteInventory,
		connect.WithSchema(inventoryServiceDeleteInventoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/inventory.v1.InventoryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case InventoryServiceCreateInventoryProcedure:
			inventoryServiceCreateInventoryHandler.ServeHTTP(w, r)
		case InventoryServiceGetInventoryProcedure:
			inventoryServiceGetInventoryHandler.ServeHTTP(w, r)
		case InventoryServiceUpdateInventoryProcedure:
			inventoryServiceUpdateInventoryHandler.ServeHTTP(w, r)
		case InventoryServiceDeleteInventoryProcedure:
			inventoryServiceDeleteInventoryHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedInventoryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedInventoryServiceHandler struct{}

func (UnimplementedInventoryServiceHandler) CreateInventory(context.Context, *connect.Request[v1.CreateInventoryRequest]) (*connect.Response[v1.CreateInventoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("inventory.v1.InventoryService.CreateInventory is not implemented"))
}

func (UnimplementedInventoryServiceHandler) GetInventory(context.Context, *connect.Request[v1.GetInventoryRequest]) (*connect.Response[v1.GetInventoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("inventory.v1.InventoryService.GetInventory is not implemented"))
}

func (UnimplementedInventoryServiceHandler) UpdateInventory(context.Context, *connect.Request[v1.UpdateInventoryRequest]) (*connect.Response[v1.UpdateInventoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("inventory.v1.InventoryService.UpdateInventory is not implemented"))
}

func (UnimplementedInventoryServiceHandler) DeleteInventory(context.Context, *connect.Request[v1.DeleteInventoryRequest]) (*connect.Response[v1.DeleteInventoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("inventory.v1.InventoryService.DeleteInventory is not implemented"))
}
