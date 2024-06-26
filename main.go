/**
* @file
*
* Go Basics.
* See for more details: https://gobyexample.com/
*
* More crash courses:
* https://devhints.io/go
*
* Local Go module & package:
* https://chat.openai.com/share/616d6cbe-5032-408d-b8b4-cc272ebe31d9
**/

package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go_for_devops/say"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	// Declare and instantiate a variable.
	myvar := "Just a test"
	_ = myvar

	// Print a variable.
	hello := "Hello World!"
	fmt.Println(hello)

	// Print and format a string using constants.
	fmt.Printf("The body temperature is %d degress Celsius.\n", bodyTemp)
	fmt.Printf("constA: %d, constB: %d, constC: %d\n", constA, constB, constC)
	fmt.Printf("autoConstA: %d, autoConstB: %d, autoConstC: %d\n", autoConstA, autoConstB, autoConstC)
	fmt.Printf("My name is %s\n", myName)

	/**
	*
	*
	*
	* Loops
	 */

	// Basic for loop.
	for i := 0; i < 10; i++ {
		fmt.Println("i is", i)
	}
	fmt.Println("i once the loop has completed is not defined")

	// j declared outside the loop.
	var j int
	for ; j < 10; j += 1 {
		fmt.Println("j is", j)
	}
	fmt.Println("j once the loop has completed is:", j)

	// while loop with for statement: remove the "post statement"
	var wi int
	for wi < 10 {
		fmt.Println("wi is", wi)
		wi++
	}

	// while loop that runs until a certain event happens.
	var wi2 bool = true
	var wi2_count int = 0
	for wi2 {
		wi2_count++
		fmt.Println("wi2_count is", wi2_count)
		if wi2_count > 12 {
			wi2 = false
		}
	}

	// Loop that uses a separate function to break.
	break_loop := 0
	for {
		break_loop++
		if err := doSomething(break_loop); err != nil {
			break
		}
		fmt.Println("keep going")
	}

	// Loop that skips iterations.
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("i is 5 - always display this")
		} else {
			fmt.Println("i is not 5, it is", i)
		}

		if i%2 == 0 {
			continue
		}

		fmt.Println("Odd number: ", i)
	}

	// Another loop with a switch statement.
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Println("i is 3")
		case 4, 5:
			fmt.Println("i is 4 or 5")
		default:
			fmt.Println("i is something else")
		}
	}

	// Loop with switch statement that uses conditional evaluation.
	for i := -10; i < 10; i++ {
		switch {
		case i < 0:
			fmt.Println("is is negative:", i)
		case i == 0:
			fmt.Println("i is zero")
		case i > 0:
			fmt.Println("i is positive:", i)
		}
	}

	/**
	*
	*
	*
	* Examples of function implementations.
	 */

	// Print the result of a function.
	result := add(5, 15)
	fmt.Println("result is", result)

	// Print the result of a function with multiple values and named results.
	res, rem := divide(10, 3)
	fmt.Printf("Divide 10 by 3. Result: %d, Reaminder: %d\n", res, rem)

	// Print the sum of multiple integers using a variadic function.
	sumVariadic := addVariadic(1, 2, 3, 4, 5)
	fmt.Println("sumVariadic is", sumVariadic)

	// Anonymous function.
	// This is a function that is declared and called immediately.
	// This function does not have a name.
	combinedWord := func(word1, word2 string) string {
		return word1 + " " + word2
	}("Hello", "World!")
	fmt.Println("combinedWord is", combinedWord)

	// Examples for a public function from another package.
	fmt.Println()
	say.PrintHello()
	say.PrintHelloWorld()

	/**
	*
	*
	*
	* Sequential data: Arrays & Slices
	 */
	// Create an array. Almost never used, as arrays have a fixed size.
	var array_a [5]int // Alternative: var array_a := [5]int{}
	array_a[2] = 7
	fmt.Println("array_a is", array_a)

	// Arrays are not pointer warpper types, but instead passed as a copy.
	array_b := [2]int{}
	changeValueAtZeroIndex(array_b)
	fmt.Println("array_b[0] outside the function is", array_b[0])

	// Slices are a wrapper around arrays:
	// Slices are not statically sized.
	// Slices can grow to accommodate new values.

	// Create a new empty slice.
	var slice_a = []int{}
	fmt.Println("slice is", slice_a)

	// Create a populated slice.
	var slice_b = []int{1, 2, 3, 4, 5}
	fmt.Println("slice_b is", slice_b)

	// Get the length of a slice.
	fmt.Println("length of sclice_a is", len(slice_a))
	fmt.Println("length of sclice_b is", len(slice_b))

	// Change a value in a slice.
	slice_b[3] = 100
	fmt.Println("slice_b is now", slice_b)

	// Append a value to a slice.
	slice_b = append(slice_b, 200, 200)
	fmt.Println("slice_b (just appended) is now", slice_b)

	// (Uncommon) Create a limited view of a slice.
	// Slicing a slice: Inclusive first, exclusive second.
	array_c := slice_b[1:3]
	fmt.Println("array_c is", array_c)

	// Slices are pointer-wrapped types, but a slice's view will not change.
	// When a slice is passed to a function and should be changed in the caller,
	// you have to either (a) return the new slice or (b) pass a pointer to the slice.
	slice_c := []int{1, 2, 3}
	slice_c = doAppend(slice_c)
	fmt.Println("slice_c is", slice_c)

	// Loop over values in a slice, using older C-style loop.
	for i := 0; i < len(slice_c); i++ {
		fmt.Printf("slice entry %d: %d\n", i, slice_c[i])
	}

	// Loop over values in a slice, using Go range.
	for index, val := range slice_c {
		fmt.Printf("slice entry %d: %d\n", index, val)
	}

	// Loop over values in a slice using range, but not using the index.
	for _, val := range slice_c {
		fmt.Printf("slice value only: %d\n", val)
	}

	/**
	*
	*
	*
	* Collections: Maps
	 */
	// Declare a map using the make function.
	// Map keys are strings, and values are integers.
	var countersMap = make(map[string]int)
	countersMap["a"] = 1
	countersMap["b"] = 2
	fmt.Println("countersMap is", countersMap)

	// Declare a map using a composite literal.
	// Map keys are strings, values are strings.
	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevrolet",
	}
	fmt.Println("compositeLiteralMap is", modelToMake)

	// Accessing a value in a map.
	carMake := modelToMake["chevelle"]
	fmt.Println("carMake is", carMake)

	// Add a value to a map.
	modelToMake["outback"] = "subaru"

	// Accessing a value by key that does not exist returns
	// the zero value for the value type.
	carMakeNonExisting := modelToMake["does_not_exist"]
	fmt.Println("carMake is", carMakeNonExisting)

	// React differently based on key existing in map.
	// When you retrieve a value from a map, you can optionally
	// receive a second return value. This second value is a
	// boolean that indicates whether the key was found in the map.
	// If the key was found, the boolean is true, and if not, it's false.
	searchForModel := "golf"
	if carMake, ok := modelToMake[searchForModel]; ok {
		fmt.Printf("car model \"%s\" has make %s\n", searchForModel, carMake)
	} else {
		fmt.Printf("car model \"%s\" does not exist\n", searchForModel)
	}

	// Loop over maps.
	for model, make := range modelToMake {
		fmt.Printf("Car model %q has make %q\n", model, make)
	}

	/**
	*
	*
	*
	* Go pointers.
	*
	* The core piece about go pointers:
	* When we call a function and pass a variable as a function
	* argument, inside the function you are working with a copy of
	* the variable. If you want to modify the original variable,
	* you have to pass a pointer to the variable.
	**/

	var intNonPointer int = 10
	// Declare a variable that is a pointer type.
	//
	// A pointer to an int stores the memory address of an int.
	//
	// You can receive the memory address of any variable by
	// using the & operator.
	//
	// You would say: "intPointer is a pointer to an int."
	//
	var intPointer *int = &intNonPointer
	// Alternative:
	// intPointer := &intNonPointer
	fmt.Println("intNonPointer is", intPointer)

	// Pointers are helpful when passing variable to a function,
	// as they allow a function to modify the original variable.
	fmt.Println("intNonPointer before function call is", intNonPointer)
	fmt.Println("intPointer holds the memory address", intPointer)
	fmt.Println("*intPointer is the value at the memory address", *intPointer)
	// Now change the value of the variable at memory address intPointer.
	*intPointer = 20
	fmt.Println("intNonPointer is now", intNonPointer)

	// A much better example is using a function to modify the original variable
	// without requiring the function to return the modified variable.
	sayWord := "Hello"
	changeStringPointer(&sayWord)
	fmt.Println("sayWord is", sayWord)

	/**
	*
	*
	*
	* Structs: A collection of variables.
	*
	* Used to group related data together.
	* An example for a struct would be a vaccine card:
	* - Name
	* - Date of birth
	* - Vaccine name
	**/

	// Uncommon method: Declare a struct, primarily used in tests.
	// This type of declaration does not allow us to reuse the struct's definition.
	var record = struct {
		Name string
		Age  int
	}{
		Name: "john Doak",
		Age:  80,
	}
	fmt.Printf("%s is %d years old.\n", record.Name, record.Age)

	// Declare a custom type.
	// Custom types are their own type, but they are based on an existing type.
	// Custom types and the types they are based on are not interchangeable.
	//
	// Define a custom type for a car model based on the string type.
	type CarModel string
	// Declare a myCar CarModel variable and assigne a value.
	var myCar CarModel = "Chevelle"
	myOtherCar := CarModel("Outback")
	fmt.Printf("My to cards are: %s and %s\n", myCar, myOtherCar)
	// Convert a CarModel back to a string through type conversion.
	myCarString := string(myCar)
	fmt.Println("My car is now a string:", myCarString)

	// Create two Record instances.
	// The Recort struct was declared outside of the main function body.
	david := Record{
		Name: "David Justice",
		Age:  28,
	}
	sarah := Record{
		Name: "Sarah Murhpy",
		Age:  28,
	}
	fmt.Printf("%+v\n", david)
	fmt.Printf("%+v\n", sarah)

	// Create a Record variable.
	john := Record{
		Name: "John Doe",
		Age:  30,
	}
	fmt.Println(john.String())

	// Change a field value in a struct.
	john.Age = 35
	fmt.Println("John is now slightly older:", john.String())

	// A struct is not a reference type, i.e. structs that are passed to a function
	// are passed as a copy. If you want to modify the original struct, you have to
	// pass a pointer to the struct.
	rec := Record{Name: "John Doe", Age: 30}
	fmt.Println("Record name is:", rec.Name)
	changeName(&rec)
	fmt.Println("Changed name is now:", rec.Name)
	// Now increase the age.
	rec.IncrAge()
	fmt.Println("Increased age is now:", rec.Age)

	// Create a new Record instance using a constructor.
	// A constructor is a function that returns a new instance of a struct,
	// usually it is a function that contains additional validation.
	// The constructor returns a pointer to the new instance.
	constructorRecord, err := newRecord("John May", 45)
	if err != nil {
		fmt.Println("Error creating record:", err)
	} else {
		fmt.Println("New record is:", constructorRecord)
	}

	/**
	*
	*
	*
	* Interfaces: A collection of method signatures that a type can implement.
	*
	* Our example interface is created outside of the main function body,
	* we will simply use it here.
	**/
	tom := Person{First: "Tom", Last: "Jones"}
	myPet := Animal{Name: "Fido"}
	nameList := StrList{"Tom", "Mary", "Michael"}
	PrintStringer(tom)
	PrintStringer(myPet)
	PrintStringer(nameList)

	/**
	*
	*
	*
	* Error handling, defer, panic, recover.
	* Go does not have exceptions, but instead uses errors.
	*
	* 2 methods for creating errors:
	* err := errors.New("error message")
	* err := fmt.Errorf("error message with %s", "formatting")
	**/
	divideBy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, div := range divideBy {
		res, err := Divide(100, div)
		if err != nil {
			fmt.Printf("100 divided by %d error: %s\n", div, err)
			continue
		}
		fmt.Printf("100 divided by %d = %d\n", div, res)
	}

	// Defer is used to schedule a function call to be run once the function completes.
	deferredText := printStuff()
	fmt.Println("deferredText is", deferredText)

	// panic will stop the program and exit while displaying text.
	// panic should only be used in the main package
	// In most cases, you should return an error instead of panicking.
	// panic("This is a panic, the program will stop here.")

	// recover allows the program to recover from a misbehaving package.
	// recover is only useful in deferred functions.

	/**
		*
		*
		*
		* Goroutines (used for concurrency in Go)
		*
		* Goroutines provide an easy and cheap way to run code concurrently,
		* i.e. to run code at the same time on a system that supports multiple cores.
		*
		* By applying the "go" keyword before a function call, the function
		* will be executed concurrently with the rest of the code.
	  *
		* Once a function or method is executed concurrently, the order of execution
		* is random. The main() function is also a goroutine. This means in the following
		* example, the "hello goroutines" may be printed before the goroutines are executed,
		* in the middle of it, or at the end.
		*
		* The concept of synchronization is important when working with goroutines.
		* Synchronization is the coordination of concurrent code execution, in detail
		* described here: https://chat.openai.com/share/d02c5ab3-483e-4954-b20d-416549236c3a
		**/
	for i := 0; i < 10; i++ {
		// go fmt.Println(i) // This will be executed concurrently.
	}
	fmt.Println("hello goroutines")
	// Following `select {}` statement is required to make goroutines above run,
	// as it will prevent the program from existing. However, it will crash this program.
	// select {}

	// WaitGroups.
	// WaitGroups are ideal for waiting on multiple goroutines to complete
	// A WaitGroup is used to wait for a collection of goroutines to finish executing.
	// Think of a WaitGroup as a counter that counts how many goroutines you've started.
	// After launching each goroutine, you tell the WaitGroup (using Add(1)).
	// When a goroutine finishes its work, it notifies the WaitGroup (using Done()),
	// which decrements the counter. The main goroutine can wait for all other
	// goroutines to finish by calling Wait() on the WaitGroup.
	// It's like having a checklist for tasks; you check off each task when
	// it's done and only proceed when all tasks are completed.

	// Create a new WaitGroup.
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		// Add a new goroutine to the WaitGroup.
		wg.Add(1)
		// Create an anonymous function that takes an integer as an argument.
		go func(n int) {
			// Use a defer statement to notify the WaitGroup that the goroutine is done.
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	// Wait for all goroutines to finish, i.e. block execution until WaitGroup counter is 0.
	wg.Wait()
	fmt.Println("All waitgroups are done")

	// Channels.
	// Channels are a synchronization mechanism that allows you to send values between
	// goroutines.
	//
	// Channels can be:
	// * buffered: holds a certain amount of data before blocking
	// * unbuffered: blocks until the data is received
	//
	// Channels are created using the make function, and are pointer-scoped types
	// (i.e. they are passed by reference). Channels are typed, so only data of that
	// specific type can be sent through the channel.
	//
	// How to get data into / out of a channel (e.g. channel = variable =)
	// Send data to a channel: ch <- "some string"
	// Receive data from a channel: myVar := <-ch
	// Pull all data out of a channel using "for range" syntax:
	// for val := range ch {   // (this acts like a <- ch)
	//   fmt.Println(val)
	// }
	//
	// Close off a channel using the close function. Once a channel is closed, it cannot
	// be used to send data, but it can still be used to receive data.
	//
	// Example: Passing strings from one channel to another.

	// Create a buffered channel with a capacity of 1.
	ch := make(chan string, 1)

	// Execute an anonymous function in a goroutine that passes words
	// from a slice to the channel ch.
	go func() {
		for _, word := range []string{"hello", "world"} {
			ch <- word
		}
		// Close the channel once all words have been sent.
		close(ch)
	}()

	// Receive the words from the channel ch and print them.
	for word := range ch {
		fmt.Println(word)
	}

	// Select statement in channels.
	// The select sattement is similar to a switch statement, but it is used for
	// listening to multiple channels. The select statement in Go allows a goroutine
	// to wait on multiple communication operations (channel sends and receives).
	// A select blocks until one of its cases can proceed, which means it waits for
	// at least one of the channels involved to be ready for the operation (send or receive).
	//
	// Blocking Explained
	// Send on a Channel: If a channel is unbuffered or its buffer is full,
	// a send operation on that channel will block the goroutine. The goroutine stays
	// suspended until another goroutine reads from the channel, making space in
	// the buffer (or directly accepting the sent value in case of an unbuffered channel).
	//
	// Receive from a Channel: If a channel is empty, a receive operation on that channel
	// will block the goroutine. The goroutine is suspended until there's data available
	// in the channel to be received.
	//
	// @see https://chat.openai.com/share/48dc5fd8-bee7-4394-83f6-cd31a775c40c

	// Example
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		// time.Sleep(2 * time.Second)
		time.Sleep(2 * time.Millisecond)
		channel1 <- "data from channel1"
	}()

	go func() {
		// time.Sleep(1 * time.Second)
		time.Sleep(1 * time.Millisecond)
		channel2 <- "data from channel2"
	}()

	for i := 0; i < 2; i++ {
		// The "select" statement blocks - i.e. waits - until one of the channels is ready.
		select {
		case msg1 := <-channel1:
			fmt.Println("Received", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received", msg2)
		}
	}

	// Mutexes (stands for mutual exclusion lock)
	// A mutex is a lock that allows only one goroutine at a time to access a shared resource.
	// If another goroutine tries to access the shared resource while it's locked, it will
	// be blocked until the mutex is unlocked.
	//
	// The purpose of a mutex is to protect a shared resource (i.e. a variable) from being accessed
	// by multiple goroutines. If a  goroutines tries to write to a varaible at the same time
	// another goroutine wants to read or write to that variable, that variable must be protected
	// by a synchronization mechanism, such as a mutex.

	// Example: Spin off 10 goroutines to add a number to a sum value.
	// The sum value must be protected as we are reading and writing from multiple goroutines.
	mySum := &sum{}

	// sumWG lets us know when all goroutines are done.
	sumWg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		sumWg.Add(1)
		go func(x int) {
			defer sumWg.Done()
			fmt.Println("Adding to sum:", x)
			mySum.add(x)
		}(i)
	}
	// This also blocks until all WaitGroups are finished.
	sumWg.Wait()
	fmt.Println("Final sum: ", mySum.get())

	/**
	*
	*
	*
	* Context type
	*
	* The context package is usd for two main purposes:
	* 1. Cancel a chain of function calls after some event (e.g. a timeout for a request)
	* 2. Pass information through a chain of function calls (e.g. user information, or request-scoped values)
	*
	* A conext is object is created from a parent context. The first context is the background context.
	* A new context in the `main` function is created from a background Context object, e.g.
	* ctx := context.Background()
	*
	* Here is a ChatGPT discussion about Contexts Do's & Don'ts:
	* https://chat.openai.com/share/cc6b7a94-fd25-434b-b5af-c9352cadfb30
	**/

	// Example: Use a context to signal a timeout.
	//
	// Create a new Context that:
	// * Context will automatically be cancelled after 50 milliseconds.
	// * Returns a cancel function that can be used to cancel the context.
	//   The cancel function will cancel the context if a timeout does NOT occur.
	//
	// A context can be canceled by calling the cancel function or with a timeout.
	// When a context is canceled, all child contexts are also canceled.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	gatherDataResponse, gatherDataErr := GatherData(ctx, "https://www.devdungeon.com/content/web-scraping-go")
	cancel()
	if gatherDataErr != nil {
		fmt.Println("Error gathering data:", gatherDataErr)
	}
	// Display the first 800 characters of the data.
	fmt.Printf("Data gathered (first 100 chars): %s... \n", gatherDataResponse[0:800])

	// Exampole: Use Context to pass a value through a chain of function calls ("call chain")
	//
	// Context should only be used to pass values that are useful on a "per-call" basis, e.g.
	// * Security information about the user making a call
	// * Telemetry information related to the call
	//
	// Values in a context are handled the following way:
	// * Values are stored in key-value paris
	// * All packages can add keys to a context
	// * Key collisions can be avoided by creating a custom key type that
	//   is only implemented by the current package.
	//
	// Create a new key for the context and a dummy value.
	var ctxKey contextKey = "myKey"
	var ctxValue = "Text that I added into a context."

	// Store a value in the context.
	ctxWithValue := context.WithValue(context.Background(), ctxKey, ctxValue)

	// Retrieve the value from the context.
	// Even though this looks like a simple function call, the value is actually
	// retrieved from the context that was created in the previous line.
	valueFromContext := extraValueFromContext(ctxWithValue, ctxKey)
	fmt.Println("Value from context is:", valueFromContext)

	/**
	*
	*
	*
	* Reading from / writing to files.
	**/
	// Read the contents of a file.
	fileContents, err := os.ReadFile("random.yaml")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	// Return value from ReadFile is a slice of bytes ([]byte), and can be
	// converted to a string using the string() function.
	fmt.Println("File contents in bytes:", fileContents)
	// for _, b := range fileContents {
	// 	fmt.Println(b)
	// }
	// for _, b := range fileContents {
	// 	fmt.Println(string(b))
	// }
	fmt.Println("File contents as a string:")
	fmt.Println(string(fileContents))

	// Writing local files (using os.Writefile())
	// * Writes a complete file to disk
	// * Will create the file if necessary
	// * Will truncate the file if it exists
	//
	// Existing strings can be converted to a slice of bytes as following
	// myString = "Hello, World!"
	// byteSlice := []byte(myString)
	if err := os.WriteFile("newrandom.yaml", fileContents, 0644); err != nil {
		fmt.Println("Error writing file:", err)
	}

	// Reading remote files.
	remoteClient := &http.Client{}
	remoteReq, err := http.NewRequest("GET", "https://www.devdungeon.com/content/web-scraping-go", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	remoteResp, err := remoteClient.Do(remoteReq)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	// remoteResp contains an io.ReaderCloser that we can read as a file.
	// Let's use io.ReadAll to read the entier content to data.
	remoteData, err := io.ReadAll(remoteResp.Body)
	if err != nil {
		fmt.Println("Error reading remote data:", err)
	}
	remoteResp.Body.Close()
	fmt.Printf("Remote data: %s...\n", string(remoteData[0:850]))

	// Writing remote content to local file using os.OpenFile.
	// Define several flags for how to interact with a file.
	// Full list can be found here: https://pkg.go.dev/os#pkg-constants
	// os.O_CREATE: Create the file if it does not exist.
	// os.O_WRONLY: Open the file for writing only.
	// os.O_TRUNC: Truncate the file to 0 bytes.
	localFileFlags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	localFile, err := os.OpenFile("remoteData.html", localFileFlags, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer localFile.Close()
	// In Go, when you read the contents of an HTTP response body
	// using io.ReadAll (or any reader function that consumes the body),
	// you're reading from a stream. Once you read the stream to its end,
	// there's no more data to read, and the stream does not automatically
	// reset to the beginning. This means if you try to read from the same
	// response body a second time with io.ReadAll, you won't get the data again;
	// instead, you'll get an empty result because the stream is already at the end.
	//
	// Workaround: If you need to read the response body more than once,
	// you will need to read it into a buffer and then work with that buffer
	// multiple times, like so:
	newReader := bytes.NewReader(remoteData)
	// Write the remote data to the local file.
	if _, err := io.Copy(localFile, newReader); err != nil {
		fmt.Println("Error writing remote data to loca file: ", err)
	}

	// Using stdin/stdou/sterr: They are just files!
	// os.Stdin, os.Stdout, and os.Stderr are all of type *os.File.
	// This means you can use them with the same functions you would use with a file.
	// For example, you can use io.Copy to copy data from os.Stdin to os.Stdout.
	soFile, err := os.Open("random.yaml")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	if _, err := io.Copy(os.Stdout, soFile); err != nil {
		fmt.Println("Error copying file to stdout:", err)
	}

	// Reading data out of a stream: User records.
	userFile, err := os.Open("users_source.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer userFile.Close()
	// The return value of `decodeUsers` is a channel.
	// We can loop over the channel to get the decoded users
	// using a simple `range` function. The range loop will
	// automatically exit once the channel is closed.
	// @see https://chat.openai.com/share/0d7ccf2c-29c0-4e7c-bfcc-1f72ffe54116
	fmt.Println("Decoding users from file:")
	for user := range decodeUsers(context.Background(), userFile) {
		if user.err != nil {
			fmt.Println("Error decoding user:", user.err)
			continue
		}
		fmt.Println(user)
	}

	// We can write data into a stream (a file), using the user list as a source.
	userTargetFile, err := os.OpenFile("users_processed.txt", localFileFlags, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer userTargetFile.Close()

	// Loop over the contents of the file again and write the detected
	// users to another file.
	// The source file is a stream in Go, and because we have already read all the
	// contents of the file, we need to reset the file pointer to the beginning of the file.
	// We can do this by using the Seek method on the file object.
	// @see https://chat.openai.com/share/bae4fda7-314b-4aa2-b98e-233b87b0f3be for details.
	_, err = userFile.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
	}
	for u := range decodeUsers(context.Background(), userFile) {
		if err := writeUser(context.Background(), userTargetFile, u); err != nil {
			fmt.Println("Error writing user:", err)
		}
	}

	// Os-agnostic pathing.
	// Print out all OS types and hardware architecture supported by go:
	// go tool dist list
	//
	// The filepath package provides functions for manipulating file paths.
	// The runtime package provides information about the runtime environment.
	fmt.Println("Current runtime:", runtime.GOOS)
	fmt.Println("Current runtime architecture:", runtime.GOARCH)
	fmt.Println("Current number of CPU cores:", runtime.NumCPU())
	fmt.Println("Current Go version:", runtime.Version())

	// Joining a filepath.
	wd, err := os.Getwd() // Get the current working directory.
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	}
	fmt.Println("Current working directory:", wd)
	content, err := os.ReadFile(filepath.Join(wd, "json_data", "config.json"))
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Printf("Config file content: %s...\n", string(content)[0:100])

	// We can use the filepath package to get additional details about the file location.
	sourceFilePath := filepath.Join(wd, "json_data", "config.json")
	fmt.Println("Current working directory:", sourceFilePath)
	fmt.Println("Base of the path:", filepath.Base(sourceFilePath))
	fmt.Println("File extension:", filepath.Ext(sourceFilePath))
	fpDir, fpFile := filepath.Split(sourceFilePath)
	fmt.Println("Path directory:", fpDir)
	fmt.Println("Filename:", fpFile)
	// Convert our absolute path into a relative path (relative to the current workdir).
	relPath, _ := filepath.Rel(wd, sourceFilePath)
	fmt.Printf("Relative path (to %s): %s\n", wd, relPath)
	// Convert a relative path to an absolute path.
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
	fmt.Printf("Absolute path (for relative path %s, from %s): %s\n", relPath, wd, absPath)
	// Determine if a path is absolute.
	if filepath.IsAbs(relPath) {
		fmt.Println("Path is absolute:", relPath)
	} else {
		fmt.Println("Path is relative:", relPath)
	}

	// Copy a file from its location to our OS' temporary directory (TMPDIR)
	// Get the filename for the source file.
	filename := filepath.Base(sourceFilePath)
	if filename == "." {
		// Path is empty
		panic("Path is empty - cannot proceed")
	}
	// Open the source file.
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("Could not open source file", err)
	}
	defer sourceFile.Close()

	// Create a new file in the temporary directory.
	tmpPath := filepath.Join(os.TempDir(), filename)
	// Open the temporary file.
	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open a new temporary file", err)
	}
	defer tmpFile.Close()
	// Copy the file to the temporary file.
	_, err = io.Copy(tmpFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file", err)
	} else {
		fmt.Printf("File %s has been copied to %s\n", sourceFilePath, tmpPath)
	}
	// Confirm that the file has been created, and print out the size.
	tmpFileInfo, err := os.Stat(tmpPath)
	if err != nil {
		fmt.Println("Coudl you get stats for file:", err)
	}
	fmt.Printf("Temporary file - size: %d bytes\n", tmpFileInfo.Size())
	fmt.Printf("Temporary file - modified time: %s\n", tmpFileInfo.ModTime())

	// Using files embedded in the binary through the `embed` package, see `embed.go`
	fmt.Printf("Embedded users source file:\n %s\n", userSource)
	embeddedConfig, err := modulesFs.ReadFile("json_data/config.json")
	if err != nil {
		fmt.Println("Error opening embedded file:", err)
	}
	fmt.Printf("Embedded file from list of files (config.json):\n%.80s...\n", string(embeddedConfig))

	// Walking any filesystem using the io/fs package (in this case it's an embedded one)
	// The io/fs package provides an interface for reading files from a filesystem.
	err = fs.WalkDir(
		modulesFs,
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".json" {
				fmt.Println("JSON file:", path)
			}

			return nil
		},
	)
	if err != nil {
		fmt.Println("Error walking directory:", err)
	}

	// Walk the filesystem of this module.
	fmt.Println("Module files, excluding subdirectories")
	err = fs.WalkDir(
		os.DirFS("."),
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// Exclude hidden files, directories, and files in subdirectories.
			if !strings.HasPrefix(path, ".") && !d.IsDir() && !strings.Contains(path, string(os.PathSeparator)) {
				fmt.Println("Regular file (not hidden, directory, file in subdir):", path)
			}
			return nil
		},
	)
	if err != nil {
		fmt.Println("Error walking directory:", err)
	}

	/**
	*
	*
	*
	* Working with common data types (CSV, YAML).
	**/
	csvInfile := "csv_data/names.csv"
	csvOutfile := "csv_data/names_sorted.csv"
	csvRecs, err := readRecs(csvInfile, true)
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
	}
	fmt.Println("CSV records (last, first):")
	for _, rec := range csvRecs {
		fmt.Printf("%s, %s\n", rec.last(), rec.first())
	}

	// Read byte records.
	csvByteRecs, err := readRecsBytes(csvInfile, true)
	if err != nil {
		fmt.Println("Error reading CSV byte records:", err)
	}
	fmt.Println("CVS records again, but read by a bufio scanner (first, last):")
	for _, rec := range csvByteRecs {
		fmt.Printf("%s %s\n", rec.first(), rec.last())
	}

	// Write the slice of records sorted to a new outfile.
	writeRecs(csvOutfile, csvByteRecs)

	// Read records using the encoding/csv package.
	csvRecsEncoding, err := readRecsCSV(csvInfile)
	if err != nil {
		fmt.Println("Error reading CSV records with encoding/csv:", err)
	}
	fmt.Println("CSV records using encoding/csv (first, last):")
	for _, rec := range csvRecsEncoding {
		fmt.Printf("%s %s\n", rec.first(), rec.last())
	}

	// Write the slice of recrds to a new outfile.
	csvOutfileWriter := "csv_data/names_writer.csv"
	writeCSVWriter(csvOutfileWriter, csvRecsEncoding)
	if err != nil {
		fmt.Println("Error writing CSV records with encoding/csv:", err)
	}
}

/**
* Examples for functions
**/

// Basic function that adds two numbers together.
func add(x int, y int) int {
	return x + y
}

// Basic function that shows a simplified way to declare parameters.
func addSimplified(x, y int) int {
	return x + y
}

// Function that returns multiple values and named results.
// Named return values result and remainder are automatically
// declared and ready for use in the function.
//
// Parentheses around return values are required when:
// - Returning multiple values
// - Named return values
func divide(num, div int) (result, remainder int) {
	result = num / div
	remainder = num % div
	return result, remainder
}

// Function with variadic arguments, i.e. the number of arguments is not fixed.
// An example is the following function of adding values.
func addVariadic(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Function that returns an error if a number is larger than 10.
func doSomething(num int) error {
	if num > 10 {
		return fmt.Errorf("num is greater than 10")
	}
	return nil
}

// Function to modify the value of a passed array.
func changeValueAtZeroIndex(array [2]int) {
	array[0] = 100
	fmt.Println("array[0] inside the function is", array[0])
}

// Function to append a value to a slice and return.
func doAppend(sl []int) []int {
	return append(sl, 1000)
}

// Function to modify the value of a pointer variable,
// i.e. modify the value of a variable at a specific
// memory address.
func changeStringPointer(word *string) {
	// Add "world" but do not return the variable.
	*word = *word + " world"
}

// Function to divide a number by a visior, including
// returning an error if the divisor is 0.
func Divide(num int, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return num / div, nil
}

// The defer keyword allows you to schedule a function call to be run
// once the function completes.
func printStuff() (value string) {
	defer fmt.Println("printStuff function is exiting")
	defer func() {
		value = "we returned this"
	}()
	fmt.Println("I am printing stuff (this will print first, as the other Println statements are deferred)")
	return ""
}

/**
 * Example function to make a web request and return the response.
 */
func GatherData(ctx context.Context, url string) (string, error) {
	// Make an HTTP GET request.
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Close the response body once the function completes.
	defer response.Body.Close()

	// Copy data from the response body to a byte slice.
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

/**
 * Define a custom string type for a context.
 */
type contextKey string

/**
 * Retrieve an example value from a context
 */
func extraValueFromContext(ctx context.Context, key contextKey) string {
	// Retrieve the value from the context using the provided key.
	value := ctx.Value(key)
	return value.(string)
}
