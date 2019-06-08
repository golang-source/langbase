package main

import (
	"github.com/redigo/redis"
)

func getReidsClient() redis.Conn {
	c, err := redis.Dial("tcp", "192.168.1.131:6379")
	if err == nil {
		return c
	}
	return c
}

type RedisLock struct {
	Lock string
}

func main() {
	c := getReidsClient()

	if c != nil {
		value, err := c.Do("get", "redis:lock")
		if err == nil {

		}
	}

}
