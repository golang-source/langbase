
package main
import (
"fmt"
redis "github.com/go-redis/redis"
)


func ExampleNewClient(redisSvrIp string) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     redisSvrIp,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	// Output: PONG <nil>
	return  client
}

func ExampleClient(client *redis.Client) {
	err := client.Set("key_wenwp", "hello world", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		//panic(err)
		fmt.Printf("not exist %s in reids\n","key2")
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main()  {
	c:=ExampleNewClient("192.168.1.131:6379")
    ExampleClient(c)
    c.Close()
}