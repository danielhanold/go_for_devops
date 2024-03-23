package main

import "fmt"

/**
* Examples for structs.
**/
// Declare a "Record" struct.
type Record struct {
	Name string
	Age  int
}

// Go does not have a built-in constructor for structs.
// Instead, you can create a function that returns a new instance of a struct.
// Constructors can be used for validation.
func newRecord(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age must be greater than 0")
	}
	return &Record{Name: name, Age: age}, nil
}

// Implement the Stringer method for the Record struct.
// String method returns a string.
func (r Record) String() string {
	return fmt.Sprintf("%s, %d", r.Name, r.Age)
}

// Change the age of a Record using a struct method.
// IncrAge method does not return a value, but instead
// modifies the value of the struct using a pointer variable.
// !! Struct arguments must not be dereferenced in a method.
func (r *Record) IncrAge() {
	r.Age++
}

// Function to modify the value of a pointer to a struct.
// Structs are automatically dereferenced and must not be
// manually dereferenced.
func changeName(r *Record) {
	r.Name = "Peter"
	fmt.Println("inside changeName: ", r.Name)
}
