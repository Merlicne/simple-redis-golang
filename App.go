package main

import (
	"SimpleRedis/EnvFactory"
	"SimpleRedis/Model"
	"SimpleRedis/RedisService/RedisImplementation"
	"context"
	"fmt"
	"log"
	"strconv"
)
func init() {
	EnvFactory.NewEnvFactory("config.yaml")
}
func main() {


	db, _ := strconv.Atoi(EnvFactory.GetStringValue("redis.db"))
	redisService, redisConnection := RedisImplementation.NewRedisService(
						EnvFactory.GetStringValue("redis.host") + ":" + EnvFactory.GetStringValue("redis.port"),
						EnvFactory.GetStringValue("redis.password"),
						db,
					)
	
	log.Println("Redis Connection Established :" + redisConnection.String())

	user := &Model.User{
		ID: 1,
		FirstName: "John",
		LastName: "Doe",
	}

	fmt.Println("Raw User Data : ", user)

	ctx := context.Background()

	err := redisService.SetUser(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	user, err = redisService.GetUser(ctx, 1)
	if err != nil || user == nil {
		log.Fatal(err)
	}

	fmt.Println("User Data from Redis : ", user)

	err = redisService.DeleteUser(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	user, err = redisService.GetUser(ctx, 1)
	if err != nil || user == nil  {
		log.Println(err)
	} else {
		fmt.Println("User Data from Redis : ", user)
	}


	defer func(){
		redisConnection.Close()
		log.Println("Redis Connection Closed")
	}()
}