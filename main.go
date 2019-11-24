package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/http"
	"toast.cafe/x/brpaste/v2/storage"
)

var S settings

type settings struct {
	Bind    string
	Redis   string
	Storage string
}

func main() {
	// ---- Flags
	flag.StringVar(&S.Bind, "bind", ":8080", "address to bind to")
	flag.StringVar(&S.Redis, "redis", "redis://localhost:6379", "redis connection string")
	flag.StringVar(&S.Storage, "storage", "redis", "type of storage to use")
	flag.Parse()

	// ---- Storage system
	var store storage.CHR

	switch S.Storage {
	case "redis":
		redisOpts, err := redis.ParseURL(S.Redis)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse redis connection string %s\n", S.Redis)
			os.Exit(1)
		}
		client := redis.NewClient(redisOpts)
		store = (*storage.Redis)(client)
	default:
		fmt.Fprintf(os.Stderr, "Could not figure out which storage system to use, tried %s\n", S.Storage)
		os.Exit(1)
	}

	// ---- Is storage healthy?
	if !store.Healthy() {
		fmt.Fprintf(os.Stderr, "Storage is unhealthy, cannot proceed.\n")
		os.Exit(1)
	}

	// ---- Start!
	handler := http.GenHandler(store)
	fasthttp.ListenAndServe(S.Bind, handler)
}
