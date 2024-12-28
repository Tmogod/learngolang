package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func string(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "lock"
	err := client.Set(ctx, key, value, 1*time.Second).Err() // 0 代表永不过期
	checkErr(err)

	// 单独调用函数设置过期时间
	client.Expire(ctx, key, 3*time.Second)

	time.Sleep(2 * time.Second)

	val, err := client.Get(ctx, key).Result()
	checkErr(err)
	fmt.Println(val)

	client.Del(ctx, key)
}

func list(ctx context.Context, client *redis.Client) {
	key := "ids"
	value := []interface{}{1, "22", "44", 2}
	_, err := client.RPush(ctx, key, value...).Result()
	checkErr(err)

	result, err := client.LRange(ctx, key, 0, -2).Result()
	checkErr(err)
	fmt.Println(result)

	client.Del(ctx, key)
}

func hashTable(ctx context.Context, client *redis.Client) {
	err := client.HSet(ctx, "学生1", "Name", "张三", "年龄", 18, "Height", 188.0).Err()
	checkErr(err)
	err = client.HSet(ctx, "学生2", "Name", "李斯", "年龄", 19, "Height", 180.0).Err()
	checkErr(err)

	result, err := client.HGet(ctx, "学生2", "年龄").Result()
	checkErr(err)
	fmt.Println(result)

	m, err := client.HGetAll(ctx, "学生1").Result()

	for k, v := range m {
		fmt.Println(k, v)
	}

	client.Del(ctx, "学生1")
	client.Del(ctx, "学生2")
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.164.132:6379",
		Password: "ll1995ll",
		DB:       0,
	})
	ctx := context.TODO()

	// string(ctx, client)

	// list(ctx, client)

	hashTable(ctx, client)
}

func checkErr(err error) {
	if err != nil {
		if !(redis.Nil == err) {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("key不存在")
			return
		}

	}
}
