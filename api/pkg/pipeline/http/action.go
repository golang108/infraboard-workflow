package http

import (
	"net/http"

	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/pb/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/workflow/api/pkg/pipeline"
)

// Action
func (h *handler) CreateAction(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := pipeline.NewCreateActionRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.VisiableMode = resource.VisiableMode_NAMESPACE

	var header, trailer metadata.MD
	ins, err := h.service.CreateAction(
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

func (h *handler) QueryAction(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	page := request.NewPageRequestFromHTTP(r)
	req := pipeline.NewQueryActionRequest(page)

	var header, trailer metadata.MD
	actions, err := h.service.QueryAction(
		ctx.Context(),
		req,
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		response.Failed(w, gcontext.NewExceptionFromTrailer(trailer, err))
		return
	}
	response.Success(w, actions)
}

func (h *handler) DescribeAction(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	req := pipeline.NewDescribeActionRequestWithName(hc.PS.ByName("name"))

	var header, trailer metadata.MD
	ins, err := h.service.DescribeAction(
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

func (h *handler) DeleteNamespaceAction(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	req := pipeline.NewDeleteActionRequestWithName(hc.PS.ByName("name"))
	req.VisiableMode = resource.VisiableMode_NAMESPACE

	var header, trailer metadata.MD
	action, err := h.service.DeleteAction(
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

func (h *handler) DeleteGlobalAction(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	hc := context.GetContext(r)
	req := pipeline.NewDeleteActionRequestWithName(hc.PS.ByName("name"))
	req.VisiableMode = resource.VisiableMode_GLOBAL

	var header, trailer metadata.MD
	action, err := h.service.DeleteAction(
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

// Action
func (h *handler) CreateGlobalAction(w http.ResponseWriter, r *http.Request) {
	ctx, err := gcontext.NewGrpcOutCtxFromHTTPRequest(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := pipeline.NewCreateActionRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}
	req.VisiableMode = resource.VisiableMode_GLOBAL

	var header, trailer metadata.MD
	ins, err := h.service.CreateAction(
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
