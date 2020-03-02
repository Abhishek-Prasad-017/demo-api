package server

import (
	"github.com/magiconair/properties/assert"
	"github.com/valyala/fasthttp"
	"testing"
)


func TestController1(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	s := &Server{}
	s.GetSoftwareDetails(ctx)
	res := ctx.Response.StatusCode()
	assert.Equal(t, res,200)
}

