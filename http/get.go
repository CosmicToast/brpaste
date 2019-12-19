package http

import (
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/storage"
)

// Get generates a handler for the /:key[/:lang] endpoint
func Get(store storage.CHR) handler {
	return func(ctx *fasthttp.RequestCtx) {
		ukey := ctx.UserValue("key")
		ulang := ctx.UserValue("lang")

		var key, lang string
		key = ukey.(string) // there's no recovering otherwise
		if ulang != nil {
			lang = ulang.(string)
		}

		res, err := store.Read(key)
		switch err {
		case storage.Unhealthy:
			ctx.Error("Backend did not respond", fasthttp.StatusInternalServerError)
		case nil: // all good
			if lang == "raw" {
				ctx.SuccessString("text/plain", res)
			} else {
				ctx.SuccessString("text/html", CodeTemplate(lang, res)) // render template
			}
		default:
			ctx.NotFound()
		}
	}
}
