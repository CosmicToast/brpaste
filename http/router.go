package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/storage"
)

// GenHandler generates the brpaste handler
func GenHandler(store storage.CHR) func(ctx *fasthttp.RequestCtx) {
	get := Get(store)
	post := Put(store, false)
	put := Put(store, true)

	r := router.New()
	r.GET("/", Index)
	r.GET("/:key", get)
	r.GET("/:key/:lang", get)
	r.POST("/", post)
	r.PUT("/:key", put)

	return r.Handler
}
