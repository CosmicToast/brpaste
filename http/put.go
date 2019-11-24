package http

import (
	"encoding/base64"
	"fmt"

	"github.com/twmb/murmur3"
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/storage"
)

func Put(store storage.CHR, put bool) handler {
	return func(ctx *fasthttp.RequestCtx) {
		data := ctx.FormValue("data")
		if len(data) == 0 { // works with nil
			ctx.Error("Missing data field", fasthttp.StatusBadRequest)
			return
		}

		ukey := ctx.UserValue("key")
		var key string
		if ukey != nil {
			key = ukey.(string)
		} else {
			hasher := murmur3.New32()
			hasher.Write(data)
			keybuf := hasher.Sum(nil)
			key = base64.RawURLEncoding.EncodeToString(keybuf)
		}
		val := string(data)

		err := store.Create(key, val, put)

		switch err {
		case storage.Collision:
			ctx.Error("Collision detected when undesired", fasthttp.StatusConflict)
		case storage.Unhealthy:
			ctx.Error("Backend did not respond", fasthttp.StatusInternalServerError)
		case nil: // everything succeeded
			if isBrowser(string(ctx.UserAgent())) {
				ctx.Redirect(fmt.Sprintf("/%s", key), fasthttp.StatusSeeOther)
			} else {
				ctx.SetStatusCode(fasthttp.StatusCreated)
				ctx.SetBodyString(key)
			}
		default:
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	}
}
