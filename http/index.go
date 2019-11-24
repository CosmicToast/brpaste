package http

import (
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/template"
)

// Index handles the / endpoint
func Index(ctx *fasthttp.RequestCtx) {
	ctx.SuccessString("text/html", template.Index()) // render template
}
