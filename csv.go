package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Define a record type as a slice of strings.
type csvRecord []string

// Validate the record data.
// The record must have exactly 2 fields.
func (r csvRecord) validate() error {
	if len(r) != 2 {
		return errors.New("data format is incorrect")
	}
	return nil
}

// Get the first name of the record.
func (r csvRecord) first() string {
	return r[0]
}

// Get the last name of the record.
func (r csvRecord) last() string {
	return r[1]
}

// Create a method to read records. A few things to note:
// - The method is using is using the record type as a return value.
func readRecs(filepath string, hasHeader bool) ([]csvRecord, error) {
	// Read the whole CSV file.
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Store the content of the file (slice of bytes) in a variable as a string.
	content := string(b)

	// Split all lines by the newline character.
	lines := strings.Split(content, "\n")

	// Decare a slice of records.
	var records []csvRecord
	for i, line := range lines {
		// Skip header line.
		if hasHeader && i == 0 {
			fmt.Println("Skipping header line")
			continue
		}

		// Skip empty lines.
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Instantiate a new record and set it to the first and last name.
		var rec csvRecord = strings.Split(line, ",")

		// Validate the record.
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", i, err)
		}

		// If the line is valid, append the record to the slice of records.
		records = append(records, rec)
	}

	// Return all records without an error.
	return records, nil
}

// Same function as above, but using the bufio & bytes packages.
// This modification will allow for a more efficient way to read the data
// in regards to memory usage, as we are not converting the whole file to a string.
func readRecsBytes(filepath string, hasHeader bool) ([]csvRecord, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new scanner.
	scanner := bufio.NewScanner(file)

	// Declare a slice of records.
	var records []csvRecord
	lineNum := 0

	for scanner.Scan() {
		// Skip header line.
		if hasHeader && lineNum == 0 {
			fmt.Println("Skipping header line")
			lineNum++
			continue
		}

		// Increment the line number.
		lineNum++

		// Read the line.
		line := scanner.Text()

		// Skip empty lines.
		if strings.TrimSpace(line) == "" {
			fmt.Println("Skip empty line")
			continue
		}

		var rec csvRecord = strings.Split(line, ",")
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", lineNum, err)
		}

		records = append(records, rec)
	}

	// Return all records.
	return records, scanner.Err()
}
