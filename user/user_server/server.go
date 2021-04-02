package main

import (
	"context"
	"fmt"
	"net"

	"../userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {}

//Return User by Id
func (*server) UserById(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	fmt.Printf("UserById function was invoked with %v", req)
	id := req.GetId()

	fmt.Println("req id====>",id)

	user,_ := getSingleUser(id)

	res := &userpb.UserResponse{
		User: &userpb.User{
			Id: user.Id,
			Fname: user.FirstName,
			City: user.City,
			Phone: user.Phone,
			Height: user.Height,
			Married: user.Married,
		},
	}

	return res, nil
}

//Return users by the list of ids
func (*server) UserByListId(ctx context.Context, req *userpb.UsersRequest) (*userpb.UsersResponse, error) {
	fmt.Printf("UserByListId function was invoked with %v", req)
	
	var listOfId []int32

	ids := req.GetIds();

	fmt.Println("request ids===>",ids);

	listOfId = append(listOfId,ids...)

	users := findListOfUsers(listOfId)

	var returnUsers []*userpb.User

	for _,user := range users {
		appendUser := &userpb.User{
			Id: user.Id,
			Fname: user.FirstName,
			City: user.City,
			Phone: user.Phone,
			Height: user.Height,
			Married: user.Married,
		}

		returnUsers = append(returnUsers,appendUser)
	}

	res := &userpb.UsersResponse {
		Users: returnUsers,
	}

	return res,nil
}	

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		fmt.Printf("Failed to listen %v", err)
	}
	fmt.Println("working: server running on port: 50051")

	s := grpc.NewServer()

	userpb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve %v", err)
	}
}