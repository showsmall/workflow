package http

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	pb "github.com/infraboard/mcube/pb/request"

	"github.com/infraboard/workflow/api/app/application"
)

func (h *handler) CreateApplication(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	tk, ok := hc.AuthInfo.(*token.Token)
	if !ok {
		response.Failed(w, fmt.Errorf("auth info is not an *token.Token"))
		return
	}

	req := application.NewCreateApplicationRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(tk)

	var header, trailer metadata.MD
	ins, err := h.service.CreateApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	response.Success(w, ins)
}

func (h *handler) QueryApplication(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	tk, ok := hc.AuthInfo.(*token.Token)
	if !ok {
		response.Failed(w, fmt.Errorf("auth info is not an *token.Token"))
		return
	}

	page := request.NewPageRequestFromHTTP(r)
	req := application.NewQueryApplicationRequest(page)
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace

	var header, trailer metadata.MD
	dommains, err := h.service.QueryApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	response.Success(w, dommains)
}

func (h *handler) DescribeApplication(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	req := application.NewDescribeApplicationRequestWithID(hc.PS.ByName("id"))
	var header, trailer metadata.MD
	ins, err := h.service.DescribeApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	ins.Desensitize()
	response.Success(w, ins)
}

func (h *handler) DeleteApplication(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	tk, ok := hc.AuthInfo.(*token.Token)
	if !ok {
		response.Failed(w, fmt.Errorf("auth info is not an *token.Token"))
		return
	}

	req := application.NewDeleteApplicationRequest(tk.Namespace, hc.PS.ByName("name"))

	var header, trailer metadata.MD
	action, err := h.service.DeleteApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	response.Success(w, action)
}

func (h *handler) PutApplication(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	tk, ok := hc.AuthInfo.(*token.Token)
	if !ok {
		response.Failed(w, fmt.Errorf("auth info is not an *token.Token"))
		return
	}

	req := application.NewUpdateApplicationRequest(hc.PS.ByName("id"))
	req.UpdateBy = tk.Account
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	ins, err := h.service.UpdateApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}

func (h *handler) PatchApplication(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	tk, ok := hc.AuthInfo.(*token.Token)
	if !ok {
		response.Failed(w, fmt.Errorf("auth info is not an *token.Token"))
		return
	}

	req := application.NewUpdateApplicationRequest(hc.PS.ByName("id"))
	req.UpdateMode = pb.UpdateMode_PATCH
	req.UpdateBy = tk.Account
	if err := request.GetDataFromRequest(r, req.Data); err != nil {
		response.Failed(w, err)
		return
	}

	var header, trailer metadata.MD
	ins, err := h.service.UpdateApplication(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	ins.Desensitize()

	response.Success(w, ins)
	return
}
