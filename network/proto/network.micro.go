// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/go-micro/network/proto/network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/router/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Network service

type NetworkService interface {
	Graph(ctx context.Context, in *GraphRequest, opts ...client.CallOption) (*GraphResponse, error)
	Nodes(ctx context.Context, in *NodesRequest, opts ...client.CallOption) (*NodesResponse, error)
	Routes(ctx context.Context, in *RoutesRequest, opts ...client.CallOption) (*RoutesResponse, error)
}

type networkService struct {
	c    client.Client
	name string
}

func NewNetworkService(name string, c client.Client) NetworkService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.network"
	}
	return &networkService{
		c:    c,
		name: name,
	}
}

func (c *networkService) Graph(ctx context.Context, in *GraphRequest, opts ...client.CallOption) (*GraphResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Graph", in)
	out := new(GraphResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Nodes(ctx context.Context, in *NodesRequest, opts ...client.CallOption) (*NodesResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Nodes", in)
	out := new(NodesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Routes(ctx context.Context, in *RoutesRequest, opts ...client.CallOption) (*RoutesResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Routes", in)
	out := new(RoutesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Network service

type NetworkHandler interface {
	Graph(context.Context, *GraphRequest, *GraphResponse) error
	Nodes(context.Context, *NodesRequest, *NodesResponse) error
	Routes(context.Context, *RoutesRequest, *RoutesResponse) error
}

func RegisterNetworkHandler(s server.Server, hdlr NetworkHandler, opts ...server.HandlerOption) error {
	type network interface {
		Graph(ctx context.Context, in *GraphRequest, out *GraphResponse) error
		Nodes(ctx context.Context, in *NodesRequest, out *NodesResponse) error
		Routes(ctx context.Context, in *RoutesRequest, out *RoutesResponse) error
	}
	type Network struct {
		network
	}
	h := &networkHandler{hdlr}
	return s.Handle(s.NewHandler(&Network{h}, opts...))
}

type networkHandler struct {
	NetworkHandler
}

func (h *networkHandler) Graph(ctx context.Context, in *GraphRequest, out *GraphResponse) error {
	return h.NetworkHandler.Graph(ctx, in, out)
}

func (h *networkHandler) Nodes(ctx context.Context, in *NodesRequest, out *NodesResponse) error {
	return h.NetworkHandler.Nodes(ctx, in, out)
}

func (h *networkHandler) Routes(ctx context.Context, in *RoutesRequest, out *RoutesResponse) error {
	return h.NetworkHandler.Routes(ctx, in, out)
}
