// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: order.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Client API for Order service

type OrderService interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*Response, error)
	SetOrder(ctx context.Context, in *SetOrderStatusRequest, opts ...client.CallOption) (*Response, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "order"
	}
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SetOrder(ctx context.Context, in *SetOrderStatusRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.SetOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	CreateOrder(context.Context, *CreateOrderRequest, *Response) error
	SetOrder(context.Context, *SetOrderStatusRequest, *Response) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		CreateOrder(ctx context.Context, in *CreateOrderRequest, out *Response) error
		SetOrder(ctx context.Context, in *SetOrderStatusRequest, out *Response) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *CreateOrderRequest, out *Response) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}

func (h *orderHandler) SetOrder(ctx context.Context, in *SetOrderStatusRequest, out *Response) error {
	return h.OrderHandler.SetOrder(ctx, in, out)
}