package main

import "errors"

func getSingleUser(id int32) (User, error) {
	users := getUsers()

	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}

	return User{}, errors.New("User not found")
}

func findListOfUsers(listOfId []int32) []User {
	users := getUsers()

	var returningUsers []User

	for _, user := range users {
		for _, id := range listOfId {
			if user.Id == id {
				returningUsers = append(returningUsers, user)
			}
		}
	}

	return returningUsers
}
