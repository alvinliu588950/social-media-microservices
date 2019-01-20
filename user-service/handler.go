package main

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"context"
	pb "social-media-microservices/user-service/proto" 
)

type UserService struct{
	db *gorm.DB
}

func (s *UserService) Create(ctx context.Context, req *pb.User, resp *pb.UserResponse) error{

	//hashed the password using bcrypt for sercurity
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(passwordHashed)

	if err != nil {
		return err
	}

	user := &pb.User{Name:req.Name, Email:req.Email, Password:req.Password}
	
	if err := s.db.FirstOrCreate(&user).Error;err != nil {
		return err
	}else {
		resp.User = user
	}

	return nil
} 

func (s *UserService) Get(ctx context.Context, req *pb.User, resp *pb.UserResponse) error{
	
	user := &pb.User{}
	if err := s.db.Where("id = ?", req.Id).First(&user).Error;err != nil {
		return err
	}

	resp.User = user
	return nil
}









