// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

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

type BookAdminHTTPServer interface {
	CreateBeer(context.Context, *CreateBeerReq) (*CreateBeerReply, error)
	DeleteBeer(context.Context, *DeleteBeerReq) (*DeleteBeerReply, error)
	GetCustomer(context.Context, *GetCustomerReq) (*GetCustomerReply, error)
	GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error)
	ListBeer(context.Context, *ListBeerReq) (*ListBeerReply, error)
	ListCustomer(context.Context, *ListCustomerReq) (*ListCustomerReply, error)
	ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	UpdateBeer(context.Context, *UpdateBeerReq) (*UpdateBeerReply, error)
}

func RegisterBookAdminHTTPServer(s *http.Server, srv BookAdminHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login", _BookAdmin_Login0_HTTP_Handler(srv))
	r.POST("/admin/v1/logout", _BookAdmin_Logout0_HTTP_Handler(srv))
	r.GET("/admin/v1/catalog/beers", _BookAdmin_ListBeer0_HTTP_Handler(srv))
	r.POST("/admin/v1/catalog/beers", _BookAdmin_CreateBeer0_HTTP_Handler(srv))
	r.PUT("/admin/v1/catalog/beers/{id}", _BookAdmin_UpdateBeer0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/catalog/beers/{id}", _BookAdmin_DeleteBeer0_HTTP_Handler(srv))
	r.GET("/admin/v1/orders", _BookAdmin_ListOrder0_HTTP_Handler(srv))
	r.GET("/admin/v1/orders", _BookAdmin_GetOrder0_HTTP_Handler(srv))
	r.GET("/admin/v1/customers", _BookAdmin_ListCustomer0_HTTP_Handler(srv))
	r.POST("/admin/v1/customers/{id}", _BookAdmin_GetCustomer0_HTTP_Handler(srv))
}

func _BookAdmin_Login0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_Logout0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/Logout")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_ListBeer0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBeerReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/ListBeer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBeer(ctx, req.(*ListBeerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBeerReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_CreateBeer0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBeerReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/CreateBeer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBeer(ctx, req.(*CreateBeerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBeerReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_UpdateBeer0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBeerReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/UpdateBeer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBeer(ctx, req.(*UpdateBeerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateBeerReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_DeleteBeer0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBeerReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/DeleteBeer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBeer(ctx, req.(*DeleteBeerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBeerReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_ListOrder0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/ListOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_GetOrder0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/GetOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrder(ctx, req.(*GetOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOrderReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_ListCustomer0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCustomerReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/ListCustomer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCustomer(ctx, req.(*ListCustomerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCustomerReply)
		return ctx.Result(200, reply)
	}
}

func _BookAdmin_GetCustomer0_HTTP_Handler(srv BookAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCustomerReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/shop.admin.v1.BookAdmin/GetCustomer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCustomer(ctx, req.(*GetCustomerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCustomerReply)
		return ctx.Result(200, reply)
	}
}

type BookAdminHTTPClient interface {
	CreateBeer(ctx context.Context, req *CreateBeerReq, opts ...http.CallOption) (rsp *CreateBeerReply, err error)
	DeleteBeer(ctx context.Context, req *DeleteBeerReq, opts ...http.CallOption) (rsp *DeleteBeerReply, err error)
	GetCustomer(ctx context.Context, req *GetCustomerReq, opts ...http.CallOption) (rsp *GetCustomerReply, err error)
	GetOrder(ctx context.Context, req *GetOrderReq, opts ...http.CallOption) (rsp *GetOrderReply, err error)
	ListBeer(ctx context.Context, req *ListBeerReq, opts ...http.CallOption) (rsp *ListBeerReply, err error)
	ListCustomer(ctx context.Context, req *ListCustomerReq, opts ...http.CallOption) (rsp *ListCustomerReply, err error)
	ListOrder(ctx context.Context, req *ListOrderReq, opts ...http.CallOption) (rsp *ListOrderReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
	UpdateBeer(ctx context.Context, req *UpdateBeerReq, opts ...http.CallOption) (rsp *UpdateBeerReply, err error)
}

type BookAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewBookAdminHTTPClient(client *http.Client) BookAdminHTTPClient {
	return &BookAdminHTTPClientImpl{client}
}

func (c *BookAdminHTTPClientImpl) CreateBeer(ctx context.Context, in *CreateBeerReq, opts ...http.CallOption) (*CreateBeerReply, error) {
	var out CreateBeerReply
	pattern := "/admin/v1/catalog/beers"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/CreateBeer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) DeleteBeer(ctx context.Context, in *DeleteBeerReq, opts ...http.CallOption) (*DeleteBeerReply, error) {
	var out DeleteBeerReply
	pattern := "/admin/v1/catalog/beers/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/DeleteBeer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...http.CallOption) (*GetCustomerReply, error) {
	var out GetCustomerReply
	pattern := "/admin/v1/customers/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/GetCustomer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) GetOrder(ctx context.Context, in *GetOrderReq, opts ...http.CallOption) (*GetOrderReply, error) {
	var out GetOrderReply
	pattern := "/admin/v1/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/GetOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) ListBeer(ctx context.Context, in *ListBeerReq, opts ...http.CallOption) (*ListBeerReply, error) {
	var out ListBeerReply
	pattern := "/admin/v1/catalog/beers"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/ListBeer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) ListCustomer(ctx context.Context, in *ListCustomerReq, opts ...http.CallOption) (*ListCustomerReply, error) {
	var out ListCustomerReply
	pattern := "/admin/v1/customers"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/ListCustomer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderReq, opts ...http.CallOption) (*ListOrderReply, error) {
	var out ListOrderReply
	pattern := "/admin/v1/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/ListOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/admin/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/Logout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BookAdminHTTPClientImpl) UpdateBeer(ctx context.Context, in *UpdateBeerReq, opts ...http.CallOption) (*UpdateBeerReply, error) {
	var out UpdateBeerReply
	pattern := "/admin/v1/catalog/beers/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/shop.admin.v1.BookAdmin/UpdateBeer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
