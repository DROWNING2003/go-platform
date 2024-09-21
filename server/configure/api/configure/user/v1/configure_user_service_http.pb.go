// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/configure/user/configure_user_service.proto

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

const OperationUserLogin = "/configure.api.configure.user.v1.User/Login"
const OperationUserRefreshToken = "/configure.api.configure.user.v1.User/RefreshToken"

type UserHTTPServer interface {
	// Login 用户登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// RefreshToken RefreshToken 刷新token
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/configure/api/v1/login", _User_Login0_HTTP_Handler(srv))
	r.POST("/configure/api/v1/token/refresh", _User_RefreshToken0_HTTP_Handler(srv))
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_RefreshToken0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.RefreshToken(ctx, req.(*RefreshTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshTokenReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest, opts ...http.CallOption) (rsp *RefreshTokenReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/configure/api/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...http.CallOption) (*RefreshTokenReply, error) {
	var out RefreshTokenReply
	pattern := "/configure/api/v1/token/refresh"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserRefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
