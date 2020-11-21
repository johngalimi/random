package main

import (
	"fmt"
)

// this'll allow us to access the struct across the whole package
type outsideUser struct {
	ID       int
	UserName string
}

func collections() {
	// arrays

	// verbose - see w arrays we must specify fixed length
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])

	// lets use implicit initialization
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// slices - dynamic length (used 99% of time relative to arrs)
	// this colon operator says to create a slice of this existing array from beginning to end
	slice2 := arr2[:]

	arr2[1] = 42
	slice2[2] = 27

	// these will both be [1 42 27] b/c the slice points back to elements in memory of arr2 and same vice versa
	fmt.Println(arr2, slice2)

	// however, let's instead set-up a slice WITHOUT the need for an underlying array
	// b/c we don't specify a size, the compiler figures out that the underlying array needs to be len of 3 but we dont have to manage that
	slice := []int{1, 2, 3}

	fmt.Println(slice)

	// append
	slice = append(slice, 4, 5)
	fmt.Println(slice)

	// let's create slices of slices
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)

	// maps

	// specify that keys are strs and vals are ints
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)

	// structs
	// associate disparate data types together
	// fields, however, are fixed at compile time

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	// the reason this prints {0  } is b/c it grabs the "initial" int 0, "initial" str ""
	fmt.Println(u)

	u.ID = 1
	u.FirstName = "John"
	u.LastName = "Smith"

	fmt.Println(u)
	fmt.Println(u.LastName)

	// we could also define implicitly in-line w/ kv pairs
	u2 := user{ID: 1, FirstName: "Average", LastName: "Joe"}
	fmt.Println(u2)

	ou3 := outsideUser{ID: 2, UserName: "helloWorld"}
	fmt.Println(ou3)
}
