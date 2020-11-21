package main

import (
	"fmt"
)

// this is now accessible to the whole package, not just the scope of the main fxn
const piNum = 3.1415

// we are going to create a "constant block"
const (
	// every subsequent use of 'iota' is incremented
	// you can build long chains of constants w/ iota
	first  = iota
	second = iota
	// this is going to do 2 + 6
	third = iota + 6
	// this is going to do 2 * 3 (the current iota)
	// look into bit-shifts (2 << 1 == 4)
	fourth = 3 * iota
	// if nothing is specified for a const in the block, i't going to inherit the above expression w/ incremented iota
	// in this case, 3 * (4)
	// fifth
	sixth
)

const (
	// look, if we add "fifth = iota" to its own const block, it restarts the iota count and returns 0 for this block
	// iota is local to the specific const block
	// really valuable when creating families of constant blocks
	fifth = iota
)

func primitives() {

	fmt.Println(first, second, third, fourth, fifth, sixth)

	// most verbose, least used
	var i int
	i = 42
	fmt.Println(i)

	// inline typing
	var f float32 = 3.14
	fmt.Println(f)

	// implicit initialization
	firstName := "John"
	fmt.Println(firstName)

	// simple bools
	b := true
	fmt.Println(b)

	// returns (3+4i)
	c := complex(3, 4)
	fmt.Println(c)

	// this is cool b/c we can unpack real v imag numbers
	r, m := real(c), imag(c)
	fmt.Println(r, m)

	// the asterisk refers to a pointer --> myFirstName is pointing to a str
	var myFirstName *string = new(string)

	// first attempt at 'dereferencing (via *)'
	// however, this will fail b/c we're trying to assign "Arthur" to an _uninitialized_ pointer
	// go protects us from assigning to uninitalized pointers b/c it wants to make sure that there's a chunk of memory assigned before
	*myFirstName = "Arthur"

	// an empty pointer returns <nil> in go
	// after declaring myFirstName to new(string), it'll return the address in memory 0xc000..
	fmt.Println(myFirstName)

	// you need to re-apply the dereferencing operation to get the actual string output
	fmt.Println(*myFirstName)

	lastName := "Smith"

	fmt.Println(lastName)

	// let's create a pointer that is pointing to that variable (lastName)
	// we accomplish that with the "address-of" (&) operator
	myPointer := &lastName

	// this'll print 0xc000010230 Smith
	fmt.Println(myPointer, *myPointer)

	lastName = "Jones"
	// this'll print 0xc000010230 Jones --> notice that memory address did NOT change
	// b/c the pointer is pointing at the reference to lastName, the address is the same
	// however, b/c we're then dereferencing that pointer again, it'll reflect the updated lastName
	fmt.Println(myPointer, *myPointer)

	// constants
	const pi = 3.1415
	fmt.Println(pi)

	// NOTE -- the value of constant MUST be able to be evaluated @ COMPILE TIME
	// you may be tempted to try and set a const based upon return from a fxn
	// that will _not_ jive in go b/c that execution happens at runtime

	// go implicitly interprets d as an int at this line
	const d int = 3
	fmt.Println(d)

	// but then, it infers that actually d should be a float and dynamically allows that to change
	// however, to be safe, you'd rather do const d int = 3, to enforce typing and THEN explicitly convert const later on
	fmt.Println(float32(d) + 1.273)
}
