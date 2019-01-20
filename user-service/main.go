package main

import (
	"log"
	micro "github.com/micro/go-micro"
	pb "social-media-microservices/user-service/proto" 
)

//This is where user-service startup
func main(){

	//1.connect to db (postgresdb)
	db, err := ConnectDB()

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer db.Close()
	
	//2.Automatically migrate your schema, to keep your schema update to date.
	//--AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns to protect your data.
	db.AutoMigrate(&User{})

	//3.start rpc server side service using go-micro(which implemets service discover pattern out of box)
	service := micro.NewService(micro.Name("user-service"))
	service.Init()

	//4.register the service handler which implements business logic
	pb.RegisterUserServiceHandler(service.Server(), &UserService{db})

	//5.run the service
	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}