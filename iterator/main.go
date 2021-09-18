package main

import "fmt"

func main() {
	user1 := &user{
		"John",
		24,
	}
	user2 := &user{
		"Tiffany",
		26,
	}
	userCollection := &userCollection{
		users: []*user{user1, user2},
	}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %v\n", user)
	}
}
