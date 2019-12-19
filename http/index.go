package http

import (
	"github.com/valyala/fasthttp"
)

// Index handles the / endpoint
func Index(ctx *fasthttp.RequestCtx) {
	ctx.SuccessString("text/html", IndexTemplate()) // render template
}
