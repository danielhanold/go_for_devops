package main

// Define a constant: Constants cannot be changed.
const bodyTemp = 37

// Enumeration via constants.
// iota is a special constant that is used to create a sequence of related values.
const (
	constA = iota
	constB = iota
	constC = iota
)

// iota is reset to 0 for each const block.
// Iota also works when it's just set on the first value.
const (
	autoConstA = iota
	autoConstB
	autoConstC
)

// Constant can, of course, also be strings.
const myName = "Danny"
