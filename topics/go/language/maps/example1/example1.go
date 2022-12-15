// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to initialize a map, write to
// it, then read and delete from it.
package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and make a map that stores values
	// of type user with a key of type string.
	users := make(map[string]user)

	// Add key/value pairs to the map.
	firstUser := user{"Rob", "Roy"}
	users["Roy"] = firstUser
	users["Ford"] = user{"Henry", "Ford"}
	thirdUser := user{"Mickey", "Mouse"}
	users["Mouse"] = thirdUser
	users["Jackson"] = user{"Michael", "Jackson"}

	firstUser.name = "changed"
	thirdUser.name = "changed3"

	// Read the value at a specific key.
	mouse := users["Mouse"]

	fmt.Printf("%+v\n", mouse)

	// Replace the value at the Mouse key.
	users["Mouse"] = user{"Jerry", "Mouse"}

	// Read the Mouse key again.
	fmt.Printf("%+v\n", users["Mouse"])

	// Delete the value at a specific key.
	delete(users, "Roy")

	// Check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "Roy")

	fmt.Println("Goodbye.")
}
