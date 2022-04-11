package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Pool *redis.Pool

func init() {
	redisHost := "120.53.241.149:6379"
	Pool = newPool(redisHost)
	close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, error := redis.Dial("tcp", server)
			checkErr(error)
			return c, error
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, error := c.Do("ping")
			return error
		},
		MaxIdle:         3,
		MaxActive:       0,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
func Get(key string) ([]byte, error) {
	coon := Pool.Get()
	defer coon.Close()
	var data []byte
	data, err := redis.Bytes(coon.Do("get", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}
func main() {
	test, err := Get("test")
	fmt.Println(test, err)
}
