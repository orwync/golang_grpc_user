package main

import (
	"context"
	"testing"
	"fmt"
	"strings"

	"../userpb"
	"google.golang.org/grpc"
)

func TestUserId1(t *testing.T) {

    // Set up a connection to the Server.
    const address = "localhost:50051"
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        t.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := userpb.NewUserServiceClient(conn)

    t.Run("Check user", func(t *testing.T) {
        id := int32(1)
        r, err := c.UserById(context.Background(), &userpb.UserRequest{Id: id})
        if err != nil {
            t.Fatalf("could not greet: %v", err)
        }
		user := r.GetUser()
        fmt.Printf("User: %v\n", user.Fname)

        
		if strings.Compare(user.Fname,"Bruce Wayne") != 0 {
			t.Fatalf("Expected fname 'Bruce Wayne', got:%v",user.Fname)
		}

		if user.City != "Gotham" {
			t.Fatalf("Expected City 'Gotham', got:%v",user.City)
		}

		if user.Phone != int64(9876523641) {
			t.Fatalf("Expected Phone 9876523641, got:%v",user.Phone)
		}

		if float32(user.Height) != float32(6.4) {
			t.Fatalf("Expected Height 6.4', got:%v",user.Height)
		}

		if user.Married != false {
			t.Fatalf("Expected Married 'false', got:%v",user.Married)
		}

		if user.Id != int32(1) {
			t.Fatalf("Expected Id '1', got:%v",user.Id)
		}

    })
}

func TestUserId6(t *testing.T) {

    // Set up a connection to the Server.
    const address = "localhost:50051"
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        t.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := userpb.NewUserServiceClient(conn)

    t.Run("Check user 6", func(t *testing.T) {
        id := int32(6)
        r, err := c.UserById(context.Background(), &userpb.UserRequest{Id: id})
        if err != nil {
            t.Fatalf("could not greet: %v", err)
        }
		user := r.GetUser()
        fmt.Printf("User: %v\n", user.Fname)

        
		if user.Fname != "" {
			t.Fatalf("Expected fname 'Bruce Wayne', got:%v",user.Fname)
		}

		if user.City != "" {
			t.Fatalf("Expected City 'Gotham', got:%v",user.City)
		}

		if user.Phone != 0 {
			t.Fatalf("Expected Phone 9876523641, got:%v",user.Phone)
		}

		if float32(user.Height) != float32(0) {
			t.Fatalf("Expected Height 6.4', got:%v",user.Height)
		}

		if user.Married != false {
			t.Fatalf("Expected Married 'false', got:%v",user.Married)
		}

		if user.Id != int32(0) {
			t.Fatalf("Expected Id '1', got:%v",user.Id)
		}

    })
}

func TestMultipleUserId12(t *testing.T) {

    // Set up a connection to the Server.
    const address = "localhost:50051"
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        t.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := userpb.NewUserServiceClient(conn)

    t.Run("Check user 1,2", func(t *testing.T) {
        ids := []int32{1,2}
        r, err := c.UserByListId(context.Background(), &userpb.UsersRequest{Ids: ids})
        if err != nil {
            t.Fatalf("could not greet: %v", err)
        }
		user := r.GetUsers()
        fmt.Printf("User: %v\n", user[0].Fname)

        
		if strings.Compare(user[0].Fname,"Bruce Wayne") != 0 {
			t.Fatalf("Expected fname 'Bruce Wayne', got:%v",user[0].Fname)
		}

		if user[0].City != "Gotham" {
			t.Fatalf("Expected City 'Gotham', got:%v",user[0].City)
		}

		if user[0].Phone != int64(9876523641) {
			t.Fatalf("Expected Phone 9876523641, got:%v",user[0].Phone)
		}

		if float32(user[0].Height) != float32(6.4) {
			t.Fatalf("Expected Height 6.4', got:%v",user[0].Height)
		}

		if user[0].Married != false {
			t.Fatalf("Expected Married 'false', got:%v",user[0].Married)
		}

		if user[0].Id != int32(1) {
			t.Fatalf("Expected Id '1', got:%v",user[0].Id)
		}

		//Id 2

		if strings.Compare(user[1].Fname,"Clark Kent") != 0 {
			t.Fatalf("Expected fname 'Bruce Wayne', got:%v",user[1].Fname)
		}

		if user[1].City != "Methropolis" {
			t.Fatalf("Expected City 'Gotham', got:%v",user[1].City)
		}

		if user[1].Phone != int64(8965475235) {
			t.Fatalf("Expected Phone 9876523641, got:%v",user[1].Phone)
		}

		if float32(user[1].Height) != float32(6.7) {
			t.Fatalf("Expected Height 6.4', got:%v",user[1].Height)
		}

		if user[1].Married != true {
			t.Fatalf("Expected Married 'false', got:%v",user[1].Married)
		}

		if user[1].Id != int32(2) {
			t.Fatalf("Expected Id '1', got:%v",user[1].Id)
		}

    })
}

func TestMultipleUserId67(t *testing.T) {

    // Set up a connection to the Server.
    const address = "localhost:50051"
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        t.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := userpb.NewUserServiceClient(conn)

    t.Run("Check user 6,7", func(t *testing.T) {
        ids := []int32{6,7}
        r, err := c.UserByListId(context.Background(), &userpb.UsersRequest{Ids: ids})
        if err != nil {
            t.Fatalf("could not greet: %v", err)
        }
		user := r.GetUsers()
        
		if len(user) !=0 {
			t.Fatalf("users length expected 0, got: %v",len(user))
		}

    })
}