package main

import (
	"context"
	"fmt"

	"../userpb"
	"google.golang.org/grpc"
)

func main() {
	const address = "localhost:50051"
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        fmt.Printf("did not connect: %v", err)
    }
    defer conn.Close()

    c := userpb.NewUserServiceClient(conn)

	getUserById(c)
	getUsersByMultipleId(c)
}

func getUserById(c userpb.UserServiceClient) {

	id := 1
	req := &userpb.UserRequest{
		Id: int32(id),
	}

	res, err := c.UserById(context.Background(),req)
	fmt.Println("Get single User")
	fmt.Println("User id:",id)

	if err!=nil {
		fmt.Printf("error while calling UserById RPC: %v",err)
	}

	fmt.Println(res)	
}

func getUsersByMultipleId(c userpb.UserServiceClient) {
	ids := []int32{1,2,3}
	req := &userpb.UsersRequest{
		Ids: ids,
	}

	res, err := c.UserByListId(context.Background(),req)
	fmt.Println("Get multiple users")
	fmt.Println("User ids:",ids)

	if err!=nil {
		fmt.Printf("error while calling UserById RPC: %v",err)
	}

	fmt.Println(res)
}