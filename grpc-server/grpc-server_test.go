package main_test

import (
	"context"
	pb "grpc/crud/grpc-proto"
	"grpc/crud/grpc-server"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var method = "create new user"
// var method = "get user"
// var method = "get all user"
// var method = "update data user"

// const email = "hahaha@gmail.com"
const email = "huhuhu@gmail.com"
const password = "r4h4514..."

func init() {
	var grpcServer *main.GrpcService = main.GrpcServer()
	go func() {
		if err := grpcServer.Run(); err != nil {
			log.Fatalf("failed to serve: %v", err.Error())
		}
	}()
}

func TestGrpcClient(t *testing.T) {
	const address = "127.0.0.1:12345"
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
		conn.Close()
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
			t.Fatalf("failed create new user %v. error: %v", email, err.Error())
		}
		// test output here
		require.Equal(t, email, response.Email, "result failed" )
		require.Equal(t, password, response.Password, "result failed" )
		log.Println("Success test email and password")
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