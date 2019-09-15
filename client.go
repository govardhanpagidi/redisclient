package main

import (
	"fmt"

	redis "github.com/go-redis/redis"
)

type RedisCache struct {
	Address string
	Client  *redis.Client
}

//Init to  initialize Cache Cnfiguration
func (r *RedisCache) Init() {

	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := r.Client.Ping().Result()
	if err == nil {
		fmt.Println(pong)
	}
}

//Set Key value pair in cache
func (r *RedisCache) Set(key string, value interface{}) error {
	return r.Client.Set(key, value, 0).Err()
}

//Get value using key
func (r *RedisCache) Get(key string) (interface{}, error) {
	return r.Client.Get(key).Result()
}

func main() {
	r := RedisCache{Address: "localhost:6379"}
	r.Init()
	err := r.Set("govardhan", "pagidi")
	if err != nil {
		fmt.Println("set errr ", err)
	}
	val, _ := r.Get("govardhan")
	fmt.Println(val)
}
