package main

import (
	"context"
	"errors"
	pb "grpc/crud/grpc-proto"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type GrpcService struct {
	pb.UnimplementedUserCrudServer
	user_list *pb.UserList
} 

// create grpc server
func GrpcServer () *GrpcService {
	return &GrpcService {
		user_list: &pb.UserList{},
	}
} 

func (userCrud *GrpcService) Run() error {
	const address = "127.0.0.1:12345"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed listen at %v", address)
	}
	serverGrpc := grpc.NewServer()
	pb.RegisterUserCrudServer(serverGrpc, userCrud)
	log.Printf("Server listening at %v", address)
	return serverGrpc.Serve(listen)
}

/* create rpc CreateNewUser*/
func (server *GrpcService) CreateNewUser(response context.Context, request *pb.NewUser) (*pb.User, error) {
	if request.GetEmail() == "" {
		return &pb.User{}, errors.New("email cannot be emty")
	}
	if request.GetPassword() == "" {
		return &pb.User{}, errors.New("password cannot be emty")
	}
	for i := 0; i < len(server.user_list.Users); i++ {
		if server.user_list.Users[i].Email == request.GetEmail() {
			return &pb.User{}, errors.New("email already exist")
		}
	}
	log.Printf("Save data %v success", request.GetEmail())
	var user_id string = uuid.NewString()
	createUser := &pb.User{
		Id: user_id,
		Email: request.GetEmail(),
		Status: "user",
		Password: request.GetPassword(),
	}
	server.user_list.Users = append(server.user_list.Users, createUser)
	return createUser, nil
}
	
/* create rpc GetUsers*/
func (server *GrpcService) GetUsers (response context.Context, request *pb.GetAllUsers) (*pb.UserList, error) {
	if request.GetEmail() != "rikianfaisal@gmail.com" || request.GetPassword() != "r4h4514..." || request.GetToken() != "12345" {
		log.Printf("%v request all data", request.GetEmail())
		return server.user_list, errors.New("unauthorize")
	}

	log.Printf("%v request all data", request.GetEmail())
	return server.user_list, nil
}

/* create rpc GetUser*/
func (server *GrpcService) GetUser (response context.Context, request *pb.NewUser) (*pb.User, error) {
	email := request.GetEmail()
	password := request.GetPassword()
	user := &pb.User{}
	
	if email == "" || password == "" {
		log.Printf("%v request data", request.GetEmail())
		return user, errors.New("one of params cannot be emty")
	}
	
	for i := 0; i < len(server.user_list.Users); i++ {
		if server.user_list.Users[i].Email == email {
			log.Printf("%v request data", request.GetEmail())
			user.Id = server.user_list.Users[i].Id
			user.Email = server.user_list.Users[i].Email
			user.Status = server.user_list.Users[i].Status
			user.Password = server.user_list.Users[i].Password
			return user, nil
		}
	}
	log.Printf("%v request data", request.GetEmail())
	return user, errors.New("data not found")
}

/* create rpc UpdateUser*/
func (server *GrpcService) UpdateUser (response context.Context, request *pb.NewDataUser) (*pb.User, error) {
	oldEmail := request.GetOldEmail()
	newEmail := request.GetNewEmail()
	oldPassword := request.GetOldPassword()
	// newPassword := request.GetNewPassword()
	user := &pb.User{}

	if oldEmail == "" || newEmail == "" || oldPassword == "" {
		log.Printf("%v request update data", request.GetOldEmail())
		return user, errors.New("one of params cannot be emty")
	}
	
	for i := 0; i < len(server.user_list.Users); i++ {
		if server.user_list.Users[i].Email == oldEmail {
			for j := 0; j < len(server.user_list.Users); j++ {
				if server.user_list.Users[j].Email == newEmail {
					return user, errors.New("email '" + newEmail + "' already exist")
				}
			}

			log.Printf("%v change data to %v", oldEmail, newEmail)
			server.user_list.Users[i].Email = newEmail
			user.Id = server.user_list.Users[i].Id
			user.Email = server.user_list.Users[i].Email
			user.Status = server.user_list.Users[i].Status
			user.Password = server.user_list.Users[i].Password
			return user, nil
		}
	}

	log.Printf("%v request update data", request.OldEmail)
	return user, errors.New("data not found")
}

/* create rpc DeleteUser*/
func (server *GrpcService) DeleteUser (response context.Context, request *pb.NewUser) (*pb.Status, error) {
	email := request.GetEmail()
	password := request.GetPassword()
	status := &pb.Status{}
	
	if email == "" || password == "" {
		log.Printf("%v request delete data", request.GetEmail())
		return status, errors.New("one of params cannot be emty")
	}
	
	for i := 0; i < len(server.user_list.Users); i++ {
		if server.user_list.Users[i].Email == email {
			// its not delete, but we change the data to null, so we need a function for refesh list users for optimaze memory. we can do it later
			log.Printf("delete data %v", request.GetEmail())
			server.user_list.Users[i].Id = ""
			server.user_list.Users[i].Email = ""
			server.user_list.Users[i].Status = ""
			server.user_list.Users[i].Password = ""
			status.Status = true
			status.Message = "delete '" + email + "' success"
			return status, nil
		}
	}
	log.Printf("%v request delete data", request.GetEmail())
	return status, errors.New("data not found")
}

func main() {
	var grpcServer *GrpcService = GrpcServer()

	if err := grpcServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err.Error())
	}
}