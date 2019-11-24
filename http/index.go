package http

import (
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/template"
)

func Index(ctx *fasthttp.RequestCtx) {
	//b := new(bytes.Buffer)
	//template.Index(b)
	ctx.SuccessString("text/html", template.Index()) // render template
}
