package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {

	fmt.Println("Creating a redis client connection")
	client := newClient()

	fmt.Println("Inserting.... [key1, valueOfKey1]")
	saveKey(client, "key1", "valueOfKey1")

	fmt.Println("Getting value of key1")
	getKey(client, "key1")
}

var ctx = context.Background()

func newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func getKey(client *redis.Client, key string) {

	value, err := client.Get(client.Context(), key).Result()
	if err != nil {
		fmt.Errorf("We have an error, verify this \n: %v", err)
		return
	}

	fmt.Println("Value is: ", value)

}

func saveKey(client *redis.Client, key string, value string) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Errorf("Setting %s, %s throw an error: %v")
	}
}
