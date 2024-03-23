package main

import "sync"

/**
* sum struct that is used in the example of mutexes in goroutines.
**/
// Define a struct with 2 fields: a mutex and an integer.
type sum struct {
	mu  sync.Mutex
	sum int
}

// Method to add a number to the sum value.
func (s *sum) get() int {
	// This locks so that only one get() or add() call
	// can acces sum.sum at a time. This gives us the
	// required syncronization.
	s.mu.Lock()
	// Defer the unlocking of the mutex until the function completes.
	defer s.mu.Unlock()
	// Return the sum value.
	return s.sum
}

// Method to add a number to the sum value.
func (s *sum) add(n int) {
	// Lock the mutex to protect the sum value.
	s.mu.Lock()
	// Defer the unlocking of the mutex until the function completes.
	defer s.mu.Unlock()
	// Add the number to the sum value.
	s.sum += n
}
