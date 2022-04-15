package main

import (
	"exacple/cache7GeeCache/day2-single-node/geecache"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{"tom": "1", "jak": "2", "ben": "3"}

func main() {
	geecache.NewGroup("source", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := geecache.NewHTTPPool(addr)
	log.Println("geecache server is running", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
