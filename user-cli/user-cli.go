package main

import (
	pb "social-media-microservices/user-service/proto"
	"log"
	"context"
	micro "github.com/micro/go-micro"
)


func main() {

	service := micro.NewService(micro.Name("user.client"))
	service.Init()

	
	client := pb.NewUserService("user-service", service.Client())

	user := pb.User{Id:1}


	resp, err := client.Get(context.Background(), &user)
	if err != nil {
		log.Fatalf("get user error: %v", err)
	} else {
		log.Printf("user: %+v", resp.User)
	}

}