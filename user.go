package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Define a User struct.
type User struct {
	Name string
	ID   int
	err  error
}

// Define a method to convert the user into a string.
func (u User) String() string {
	return fmt.Sprintf("%s:%d", u.Name, u.ID)
}

// Split a line from input file and return a User struct.
func getUser(s string) (User, error) {
	// Split the string by the colon.
	sp := strings.Split(s, ":")

	// If the length of the slice is not 2, user record was in in correct format.
	if len(sp) != 2 {
		return User{}, fmt.Errorf("record (%s) was not in the correct format", s)
	}

	// Convert the string to an integer using Atoi. (Alphanum to integer)
	// If this fails, we know that the second part of the record was not a number.
	id, err := strconv.Atoi(sp[1])
	if err != nil {
		return User{}, fmt.Errorf("record (%s) had a non-numeric ID", s)
	}

	// Return a new User struct.
	return User{Name: strings.TrimSpace(sp[0]), ID: id, err: nil}, nil
}

// Read in a stream andw rites the User records to a channel.
// decodeUsers reads data from the provided reader and decodes it into User structs.
// It returns a channel of User structs that are read from the reader.
// The channel has a buffer of 1 to allow non-blocking sending of User structs.
// The function runs in a separate goroutine and closes the channel when it completes.
// If the provided context is canceled, it sends a User struct with the error set to the context error.
// If there is an error while decoding a line of data, it sends a User struct with the error set to the decoding error.
func decodeUsers(ctx context.Context, r io.Reader) chan User {
	// Create a channel of User structs, with a buffer of 1.
	ch := make(chan User, 1)

	go func() {
		// Defer the closing of the channel until the function completes.
		defer close(ch)

		// Create a new Scanner using the bufio (buffer input/output) package.
		// Scanner provides a convenient interface for reading data such as a file of
		// newline-delimited lines of text.
		scanner := bufio.NewScanner(r)

		// Iterate over the lines of the input file.
		// By default, the scanner splits the input into lines using the newline character.
		// The io.Reder returns an error, io.EOF, when the end of the file is reached.
		// This will stop the loop.
		//
		// During the loop, Scanner will return the bytes read from the input file.
		// They can be retrieved with the Text() method.
		for scanner.Scan() {
			// If the provided context is canceled, it sends a User
			// struct with the error set to the context error.
			if ctx.Err() != nil {
				ch <- User{err: ctx.Err()}
				return
			}

			// Decode the line into a User struct.
			fmt.Printf("Current line being processed: %s\n", scanner.Text())

			// If this line is a comment (indicated by either `//` or `#`), skip it.
			if strings.HasPrefix(scanner.Text(), "//") || strings.HasPrefix(scanner.Text(), "#") {
				fmt.Printf("Skip the current line (%s) - it is a comment\n", scanner.Text())
				continue
			}

			u, err := getUser(scanner.Text())
			// If there is an error while decoding a line of data, it sends a User
			// struct with the error set to the decoding error.
			if err != nil {
				u.err = err
				ch <- u
				return
			}

			// Send the User struct on the channel.
			ch <- u
		}
	}()

	return ch
}

// Write a user record into a Writer (e.g. a file.)
func writeUser(ctx context.Context, w io.Writer, u User) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Write the user record to the writer.
	if _, err := w.Write([]byte(u.String())); err != nil {
		return err
	}

	// Write a newline character to separate the records.
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}
