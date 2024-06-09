package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	Server(
		endpoint Endpoint,
		decode func(ctx context.Context, r *http.Request) (interface{}, error),
		encode func(ctx context.Context, w http.ResponseWriter, resp interface{}) error,
		encodeError func(ctx context.Context, err error, w http.ResponseWriter) error,
	)
}
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type transport struct {
	w   http.ResponseWriter
	r   *http.Request
	ctx context.Context
}

func New(w http.ResponseWriter, r *http.Request, ctx context.Context) Transport {
	return &transport{
		w:   w,
		r:   r,
		ctx: ctx,
	}
}
func (t *transport) Server(
	endpoint Endpoint,
	decode func(ctx context.Context, r *http.Request) (interface{}, error),
	encode func(ctx context.Context, w http.ResponseWriter, resp interface{}) error,
	encodeError func(ctx context.Context, err error, w http.ResponseWriter) error,
) {
	data, err2 := decode(t.ctx, t.r)
	if err2 != nil {
		encodeError(t.ctx, err2, t.w)
		return
	}
	resp, err := endpoint(t.ctx, data)
	if err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}
	if err := encode(t.ctx, t.w, resp); err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}
}
