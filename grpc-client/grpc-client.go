package main

import (
	"context"
	pb "grpc/crud/grpc-proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// aktivkan method satu persatu
var method = "create new user"
// var method = "get user"
// var method = "get all user"
// var method = "update data user"

// const email = "hahaha@gmail.com"
const email = "huhuhu@gmail.com"
const password = "r4h4514..."

func main () {
	const address = "127.0.0.1:12345"
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("connection to grpc server at %v failed", address)
		conn.Close()
		return
	}
	defer cancel()
	client := pb.NewUserCrudClient(conn)
	switch method {
	case "create new user":
		response, err := client.CreateNewUser(ctx, &pb.NewUser{
			Email: email, 
			Password: password,
		})
		if err != nil {
			log.Printf("failed create new user %v. error: %v", email, err.Error())
			return
		}
		log.Println(response)
		return
	case "get user"	:
		response, err := client.GetUser(ctx, &pb.NewUser{Email: email, Password: password})
		if err != nil {
			log.Printf("user %v. data not found", email)
			return
		}
		log.Println(response)
		return
	case "get all user"	:
		params := &pb.GetAllUsers {
			Email: "rikianfaisal@gmail.com",
			Password: "r4h4514...",
			Token: "12345",
		}
		response, err := client.GetUsers(ctx, params)
		if err != nil {
			log.Printf("email %v tidak berhak melihan konten ini", email)
			return
		}
		for i := 0; i < len(response.GetUsers()); i++ {
			log.Println(response.GetUsers()[i])
		}
		return
	case "update data user":
		newDataUser := &pb.NewDataUser{}
		newDataUser.OldEmail = email
		newDataUser.OldPassword = password
		newDataUser.NewEmail = "hahaha@gmail.com"
		response, err := client.UpdateUser(ctx, newDataUser)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(response)
		return
	case "delete data user":
		response, err := client.DeleteUser(ctx, &pb.NewUser{
			Email: email,
			Password: password,
		})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(response)
		return
	default:
		log.Println("Method not allowed")	
		return
	}
}