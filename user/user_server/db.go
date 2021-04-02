package main

func getUsers() []User {
	users := []User{
		User{
			Id:        1,
			FirstName: "Bruce Wayne",
			City:      "Gotham",
			Phone:     9876523641,
			Height:    6.4,
			Married:   false,
		},
		User{
			Id:        2,
			FirstName: "Clark Kent",
			City:      "Methropolis",
			Phone:     8965475235,
			Height:    6.7,
			Married:   true,
		},
		User{
			Id:        3,
			FirstName: "Diana Prince",
			City:      "Amazon",
			Phone:     9875412542,
			Height:    5.9,
			Married:   false,
		},
		User{
			Id:        4,
			FirstName: "Barry Allen",
			City:      "Star City",
			Phone:     9876523641,
			Height:    5.6,
			Married:   true,
		},
		User{
			Id:        5,
			FirstName: "Victor Stone",
			City:      "Gotham",
			Phone:     9856324515,
			Height:    6.7,
			Married:   false,
		},
	}

	return users
}
