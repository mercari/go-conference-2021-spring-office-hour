// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/customer/proto/customer.proto

package protoconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/mercari/mercari-microservices-example/services/customer/proto"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CustomerServiceName is the fully-qualified name of the CustomerService service.
	CustomerServiceName = "mercari.mercari_microservices_example.customer.CustomerService"
)

// CustomerServiceClient is a client for the
// mercari.mercari_microservices_example.customer.CustomerService service.
type CustomerServiceClient interface {
	CreateCustomer(context.Context, *connect_go.Request[proto.CreateCustomerRequest]) (*connect_go.Response[proto.CreateCustomerResponse], error)
	GetCustomer(context.Context, *connect_go.Request[proto.GetCustomerRequest]) (*connect_go.Response[proto.GetCustomerResponse], error)
	GetCustomerByName(context.Context, *connect_go.Request[proto.GetCustomerByNameRequest]) (*connect_go.Response[proto.GetCustomerByNameResponse], error)
}

// NewCustomerServiceClient constructs a client for the
// mercari.mercari_microservices_example.customer.CustomerService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCustomerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CustomerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &customerServiceClient{
		createCustomer: connect_go.NewClient[proto.CreateCustomerRequest, proto.CreateCustomerResponse](
			httpClient,
			baseURL+"/mercari.mercari_microservices_example.customer.CustomerService/CreateCustomer",
			opts...,
		),
		getCustomer: connect_go.NewClient[proto.GetCustomerRequest, proto.GetCustomerResponse](
			httpClient,
			baseURL+"/mercari.mercari_microservices_example.customer.CustomerService/GetCustomer",
			opts...,
		),
		getCustomerByName: connect_go.NewClient[proto.GetCustomerByNameRequest, proto.GetCustomerByNameResponse](
			httpClient,
			baseURL+"/mercari.mercari_microservices_example.customer.CustomerService/GetCustomerByName",
			opts...,
		),
	}
}

// customerServiceClient implements CustomerServiceClient.
type customerServiceClient struct {
	createCustomer    *connect_go.Client[proto.CreateCustomerRequest, proto.CreateCustomerResponse]
	getCustomer       *connect_go.Client[proto.GetCustomerRequest, proto.GetCustomerResponse]
	getCustomerByName *connect_go.Client[proto.GetCustomerByNameRequest, proto.GetCustomerByNameResponse]
}

// CreateCustomer calls
// mercari.mercari_microservices_example.customer.CustomerService.CreateCustomer.
func (c *customerServiceClient) CreateCustomer(ctx context.Context, req *connect_go.Request[proto.CreateCustomerRequest]) (*connect_go.Response[proto.CreateCustomerResponse], error) {
	return c.createCustomer.CallUnary(ctx, req)
}

// GetCustomer calls mercari.mercari_microservices_example.customer.CustomerService.GetCustomer.
func (c *customerServiceClient) GetCustomer(ctx context.Context, req *connect_go.Request[proto.GetCustomerRequest]) (*connect_go.Response[proto.GetCustomerResponse], error) {
	return c.getCustomer.CallUnary(ctx, req)
}

// GetCustomerByName calls
// mercari.mercari_microservices_example.customer.CustomerService.GetCustomerByName.
func (c *customerServiceClient) GetCustomerByName(ctx context.Context, req *connect_go.Request[proto.GetCustomerByNameRequest]) (*connect_go.Response[proto.GetCustomerByNameResponse], error) {
	return c.getCustomerByName.CallUnary(ctx, req)
}

// CustomerServiceHandler is an implementation of the
// mercari.mercari_microservices_example.customer.CustomerService service.
type CustomerServiceHandler interface {
	CreateCustomer(context.Context, *connect_go.Request[proto.CreateCustomerRequest]) (*connect_go.Response[proto.CreateCustomerResponse], error)
	GetCustomer(context.Context, *connect_go.Request[proto.GetCustomerRequest]) (*connect_go.Response[proto.GetCustomerResponse], error)
	GetCustomerByName(context.Context, *connect_go.Request[proto.GetCustomerByNameRequest]) (*connect_go.Response[proto.GetCustomerByNameResponse], error)
}

// NewCustomerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCustomerServiceHandler(svc CustomerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/mercari.mercari_microservices_example.customer.CustomerService/CreateCustomer", connect_go.NewUnaryHandler(
		"/mercari.mercari_microservices_example.customer.CustomerService/CreateCustomer",
		svc.CreateCustomer,
		opts...,
	))
	mux.Handle("/mercari.mercari_microservices_example.customer.CustomerService/GetCustomer", connect_go.NewUnaryHandler(
		"/mercari.mercari_microservices_example.customer.CustomerService/GetCustomer",
		svc.GetCustomer,
		opts...,
	))
	mux.Handle("/mercari.mercari_microservices_example.customer.CustomerService/GetCustomerByName", connect_go.NewUnaryHandler(
		"/mercari.mercari_microservices_example.customer.CustomerService/GetCustomerByName",
		svc.GetCustomerByName,
		opts...,
	))
	return "/mercari.mercari_microservices_example.customer.CustomerService/", mux
}

// UnimplementedCustomerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCustomerServiceHandler struct{}

func (UnimplementedCustomerServiceHandler) CreateCustomer(context.Context, *connect_go.Request[proto.CreateCustomerRequest]) (*connect_go.Response[proto.CreateCustomerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mercari.mercari_microservices_example.customer.CustomerService.CreateCustomer is not implemented"))
}

func (UnimplementedCustomerServiceHandler) GetCustomer(context.Context, *connect_go.Request[proto.GetCustomerRequest]) (*connect_go.Response[proto.GetCustomerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mercari.mercari_microservices_example.customer.CustomerService.GetCustomer is not implemented"))
}

func (UnimplementedCustomerServiceHandler) GetCustomerByName(context.Context, *connect_go.Request[proto.GetCustomerByNameRequest]) (*connect_go.Response[proto.GetCustomerByNameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mercari.mercari_microservices_example.customer.CustomerService.GetCustomerByName is not implemented"))
}
