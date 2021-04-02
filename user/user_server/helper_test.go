package main

import (
	"reflect"
	"testing"
)

func TestUserId3(t *testing.T) {
	user, _ := getSingleUser(3)

	testUser := User{
		Id:        3,
		FirstName: "Diana Prince",
		City:      "Amazon",
		Phone:     9875412542,
		Height:    172,
		Married:   false,
	}

	if !reflect.DeepEqual(user, testUser) {
		t.Errorf("Expect user to be Diana Prince for user id 3, but got %v", user.FirstName)
	}
}

func TestIdDosentExist(t *testing.T) {
	_, err := getSingleUser(8)

	if err.Error() != "User not found" {
		t.Errorf("Expect user not found but got %v", err.Error())
	}
}
