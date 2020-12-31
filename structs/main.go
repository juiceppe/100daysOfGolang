package main

import "fmt"

type userInfo struct {
	email   string
	zipCode int
}

type user struct {
	firstName string
	lastName  string
	userInfo  //we can omit the field name! Go will assume that the fieldname will be the same as the embedded struct, in this case userInfo
}

func main() {
	//define all the properties for the struct
	//user1 := user{"User", "One", {"ciao", 1234}} //Depends on the order of the fields in the struct definition
	user2 := user{
		firstName: "User",
		lastName:  "Two",
		userInfo: userInfo{
			email:   "user2@user2.com",
			zipCode: 91025,
		},
	}
	//This is more "order" proof, by listing the property name and assigning the value
	var user3 user
	user3.firstName = "User"
	user3.lastName = "Three"
	//fmt.Println(user3)       can't print an empty struct, showing just { }
	//fmt.Printf("%+v", user2) //%+v will print all the fields and print them for user3, if user3 is empty it will show {firstName: lastName:}
	//user2Pointer := &user2 //& is an operator that we need to use when we want to access memory address for that var, we transform user2 into a pointer and assign it to user2Pointer
	user2.updateUser("updatedName")
	user2.print()
	//fmt.Println(user2, user3)
}

//receiver with struct
func (u user) print() {
	fmt.Printf("%+v", u)
}

//Update struct
func (pointerToUser *user) updateUser(newFirstName string) {
	// '*'in this case is *user it means we're working with a pointer to an user type
	(*pointerToUser).firstName = newFirstName //the '*' operator is used to get value that is existing at that memory address - a pointer that point at that user
}
