package main

import (
	"fmt"
	"strings"
)

/**
* Examples for interfaces.
*
* Interfaces are a collection of method signatures that a type can implement.
* A type implements an interface if the implementing type has declared all methods
* defined in the interface.
* Interfaces are a way to define behavior without defining the actual implementation.
* In contrast to other programming languages, Go has no "implements" keyword.
* Instead, a type implicitly implements an interface if the type has all the methods.
**/

// Example interface: Stringer
// A type that implements the Stringer interface must have a String method.
// Any varaibel that has the String() method can be stored in a variable of type Stringer.
type Stringer interface {
	String() string
}

// We will now implement 3 different types that implement the Stringer interface,
// because all 3 types will have the String() method:
// ----------------------------------------
// Person represents a record of a person. We define the String method for this type.
type Person struct {
	First, Last string
}

func (p Person) String() string {
	return fmt.Sprintf("%s,%s", p.Last, p.First)
}

// Animal represents a record of an animal. We define the String method for this type.
type Animal struct {
	Name string
}

func (a Animal) String() string {
	return a.Name
}

// StrList is a slice of strings.
type StrList []string

func (s StrList) String() string {
	return strings.Join(s, ",")
}

// Now that we have 3 types that implement the Stringer interface, we can create a function
// that implements Stringer. In this function, we know that all variables will have the String() method.
// See the implementation in the main function.
func PrintStringer(s Stringer) {
	fmt.Println(s.String())
}
