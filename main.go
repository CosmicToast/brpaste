package main

import (
	"fmt"
	"flag"
	"os"

	bolt "go.etcd.io/bbolt"
	"github.com/go-redis/redis/v7"
	"github.com/valyala/fasthttp"
	"toast.cafe/x/brpaste/v2/http"
	"toast.cafe/x/brpaste/v2/storage"
)

var s settings

type settings struct {
	Bind    string
	Bolt    string
	Redis   string
	Storage string
}

func main() {
	// ---- Flags
	flag.StringVar(&s.Bind, "bind", ":8080", "address to bind to")
	flag.StringVar(&s.Bolt, "bolt", "brpaste.db", "bolt database file to use")
	flag.StringVar(&s.Redis, "redis", "redis://localhost:6379", "redis connection string")
	flag.StringVar(&s.Storage, "storage", "bolt", "type of storage to use")
	flag.Parse()

	// ---- Storage system
	var store storage.CHR

	switch s.Storage {
	case "memory":
		store = storage.NewMemory()
	case "redis":
		redisOpts, err := redis.ParseURL(s.Redis)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse redis connection string %s\n", s.Redis)
			os.Exit(1)
		}
		client := redis.NewClient(redisOpts)
		store = (*storage.Redis)(client)
	case "bolt":
		db, err := bolt.Open(s.Bolt, 0600, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open/create boltdb database at %s\n", s.Bolt)
			os.Exit(1)
		}
		store, err = storage.OpenBolt(db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to initialize boltdb database at %s: %s\n", s.Bolt, err)
			os.Exit(1)
		}
		defer db.Close()
	default:
		fmt.Fprintf(os.Stderr, "Could not figure out which storage system to use, tried %s\n", s.Storage)
		os.Exit(1)
	}

	// ---- Is storage healthy?
	if !store.Healthy() {
		fmt.Fprintf(os.Stderr, "Storage is unhealthy, cannot proceed.\n")
		os.Exit(1)
	}

	// ---- Start!
	handler := http.GenHandler(store)
	fasthttp.ListenAndServe(s.Bind, handler)
}
