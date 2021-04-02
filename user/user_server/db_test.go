package main

import (
	"testing"
)

func TestNumberOfUsers(t *testing.T) {
	users := getUsers()

	if len(users) != 5 {
		t.Errorf("Expectd 5 users, but got %v", len(users))
	}
}

func TestFirstUser(t *testing.T) {
	users := getUsers()
	if users[0].FirstName != "Bruce Wayne" {
		t.Errorf("Expected first user name to be Bruce Wayne, but got %v", users[0].FirstName)
	}

}

func TestLastUser(t *testing.T) {
	users := getUsers()
	if users[len(users)-1].FirstName != "Victor Stone" {
		t.Errorf("Expected last user name to be Victor Stone, but got %v", users[len(users)-1].FirstName)
	}
}
