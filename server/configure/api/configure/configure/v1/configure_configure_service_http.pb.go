// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/configure/configure/configure_configure_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationConfigureCompareConfigure = "/configure.api.configure.configure.v1.Configure/CompareConfigure"
const OperationConfigureGetConfigure = "/configure.api.configure.configure.v1.Configure/GetConfigure"
const OperationConfigureUpdateConfigure = "/configure.api.configure.configure.v1.Configure/UpdateConfigure"

type ConfigureHTTPServer interface {
	CompareConfigure(context.Context, *CompareConfigureRequest) (*CompareConfigureReply, error)
	GetConfigure(context.Context, *GetConfigureRequest) (*GetConfigureReply, error)
	UpdateConfigure(context.Context, *UpdateConfigureRequest) (*UpdateConfigureReply, error)
}

func RegisterConfigureHTTPServer(s *http.Server, srv ConfigureHTTPServer) {
	r := s.Route("/")
	r.GET("/configure/api/v1/configure", _Configure_GetConfigure0_HTTP_Handler(srv))
	r.PUT("/configure/api/v1/configure", _Configure_UpdateConfigure0_HTTP_Handler(srv))
	r.POST("/configure/api/v1/configure/compare", _Configure_CompareConfigure0_HTTP_Handler(srv))
}

func _Configure_GetConfigure0_HTTP_Handler(srv ConfigureHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetConfigureRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationConfigureGetConfigure)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetConfigure(ctx, req.(*GetConfigureRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetConfigureReply)
		return ctx.Result(200, reply)
	}
}

func _Configure_UpdateConfigure0_HTTP_Handler(srv ConfigureHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateConfigureRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationConfigureUpdateConfigure)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateConfigure(ctx, req.(*UpdateConfigureRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateConfigureReply)
		return ctx.Result(200, reply)
	}
}

func _Configure_CompareConfigure0_HTTP_Handler(srv ConfigureHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompareConfigureRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationConfigureCompareConfigure)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CompareConfigure(ctx, req.(*CompareConfigureRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompareConfigureReply)
		return ctx.Result(200, reply)
	}
}

type ConfigureHTTPClient interface {
	CompareConfigure(ctx context.Context, req *CompareConfigureRequest, opts ...http.CallOption) (rsp *CompareConfigureReply, err error)
	GetConfigure(ctx context.Context, req *GetConfigureRequest, opts ...http.CallOption) (rsp *GetConfigureReply, err error)
	UpdateConfigure(ctx context.Context, req *UpdateConfigureRequest, opts ...http.CallOption) (rsp *UpdateConfigureReply, err error)
}

type ConfigureHTTPClientImpl struct {
	cc *http.Client
}

func NewConfigureHTTPClient(client *http.Client) ConfigureHTTPClient {
	return &ConfigureHTTPClientImpl{client}
}

func (c *ConfigureHTTPClientImpl) CompareConfigure(ctx context.Context, in *CompareConfigureRequest, opts ...http.CallOption) (*CompareConfigureReply, error) {
	var out CompareConfigureReply
	pattern := "/configure/api/v1/configure/compare"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationConfigureCompareConfigure))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ConfigureHTTPClientImpl) GetConfigure(ctx context.Context, in *GetConfigureRequest, opts ...http.CallOption) (*GetConfigureReply, error) {
	var out GetConfigureReply
	pattern := "/configure/api/v1/configure"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationConfigureGetConfigure))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ConfigureHTTPClientImpl) UpdateConfigure(ctx context.Context, in *UpdateConfigureRequest, opts ...http.CallOption) (*UpdateConfigureReply, error) {
	var out UpdateConfigureReply
	pattern := "/configure/api/v1/configure"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationConfigureUpdateConfigure))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
