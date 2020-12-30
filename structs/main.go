package main

import "fmt"

type userInfo struct {
	email   string
	zipCode int
}

type user struct {
	firstName string
	lastName  string
	info      userInfo
}

func main() {
	//define all the properties for the struct
	//user1 := user{"User", "One", {"ciao", 1234}} //Depends on the order of the fields in the struct definition
	user2 := user{
		firstName: "User",
		lastName:  "Two",
		info: userInfo{
			email:   "user2@user2.com",
			zipCode: 91025,
		},
	}
	//This is more "order" proof, by listing the property name and assigning the value
	var user3 user
	user3.firstName = "User"
	user3.lastName = "Three"
	//fmt.Println(user3)       can't print an empty struct, showing just { }
	fmt.Printf("%+v", user2) //%+v will print all the fields and print them for user3, if user3 is empty it will show {firstName: lastName:}
	//fmt.Println(user2, user3)
}
