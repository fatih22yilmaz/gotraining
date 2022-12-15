// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to understand method sets.
package main

import (
	"fmt"
	"strings"
)

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

//instead of notifier interface lets create uppercase interface

type uppercase interface {
	uppercase() // if we declare its behavior as this only way to uppercase sth can be done via pointer semantic implementations.
}

func (u *user) uppercase() {
	u.name = strings.ToUpper(u.name)
	u.email = strings.ToUpper(u.email)
}

func (u user) valueRecFunc() {

}

func main() { //https://stackoverflow.com/a/33591156

	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	u.valueRecFunc()
	(&u).valueRecFunc()

	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	u.notify()

	sendNotification(u)

	// ./example1.go:36: cannot use u (type user) as type notifier in argument to sendNotification:
	//   user does not implement notifier (notify method has pointer receiver)

	callUpperCase(u) //if we could pass the copy of u its value semantic received function uppercase could uppercase our data but it would be on a copy.
	//compiler doesn't allow us to do that bc. when our function accepts an interface we can pass a concrete data (value or pointer) on contrary if we define a function
	//that accepts a struct value or pointer we could not have pass interchangeably. if method accepts a struct pointer we should have pass a struct pointer not a value
	// or if it accepts a struct value we should have pass a struct value not a struct pointer.
	// continuing to interface accepting functions as we said they can accept both pointers and values which implements its methods. To avoid breaking the pointer receiver functions obligation
	//compiler does not allow us to provide a value of that struct.

	sth(u)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

func sth(u user) {
	u.uppercase()
}

func callUpperCase(u uppercase) {
	u.uppercase()
}
