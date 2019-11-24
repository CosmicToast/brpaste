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
	Bind  string
	Redis string
}

func main() {
	// ---- Flags
	flag.StringVar(&S.Bind, "bind", ":8080", "address to bind to")
	flag.StringVar(&S.Redis, "redis", "redis://localhost:6379", "redis connection string")
	flag.Parse()

	// ---- Storage system
	redisOpts, err := redis.ParseURL(S.Redis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse redis connection string %s\n", S.Redis)
		os.Exit(1)
	}
	client := redis.NewClient(redisOpts)
	storage := (*storage.Redis)(client)

	// ---- Is storage healthy?
	if !storage.Healthy() {
		fmt.Fprintf(os.Stderr, "Storage is unhealthy, cannot proceed.\n")
		os.Exit(1)
	}

	// ---- Start!
	handler := http.GenHandler(storage)
	fasthttp.ListenAndServe(S.Bind, handler)
}
