package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
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

// csv converts the csvRecord to a byte slice in CSV format.
// Each field in the csvRecord is separated by a comma.
// The resulting byte slice includes a newline character at the end.
func (r csvRecord) csv() []byte {
	// Create a new bytes.Buffer interface, which acts similar to an in-memory file.
	b := bytes.Buffer{}

	// Loop over each field in the record and write it to the buffer, separate by a comma
	for _, field := range r {
		b.WriteString(field + ",")
	}

	// Add a carriage return after the name.
	b.WriteString("\n")

	// Returns the buffer as a slice of bytes.
	return b.Bytes()
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

// Write CSV records sorted to a CSV outfile.
func writeRecs(filepath string, recs []csvRecord) error {
	// Create a new file and open it for writing.
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Sort passed slice of records by last name.
	sort.Slice(
		recs,
		func(i, j int) bool {
			return recs[i].last() < recs[j].last()
		},
	)

	for _, rec := range recs {
		_, err := file.Write(rec.csv())
		if err != nil {
			return err
		}
	}

	return nil
}

// The encoding/csv package provides two types for handling CSV files: Reader and Writer.
// It confirms to the RFC 4180 standard, used to define CSV files.
func readRecsCSV(filepath string) ([]csvRecord, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new reader.
	reader := csv.NewReader(file)
	// Define the number of expected fields per record.
	reader.FieldsPerRecord = 2
	// Ignore any leading whitespace for each record, including leading commas.
	reader.TrimLeadingSpace = true

	// Create a slice of CVS records.
	var recs []csvRecord

	// Loop over al records and read line by line.
	for {
		// Read the next line in the CSV file.
		data, err := reader.Read()

		// Handle any errors, including the end of the file.
		if err != nil {
			// If the error is that there are no more lines (i.e. we have reached the end of the file),
			// stop the loop execution.
			if err == io.EOF {
				break
			}
			// Otherwise, actually return the error.
			return nil, err
		}

		// Skip lines that would be a comment.
		if strings.HasPrefix(data[0], "#") || strings.HasPrefix(data[0], ";") {
			fmt.Println("Skipping comment line:", strings.Join(data, " "))
			continue
		}

		// Append the record to the slice of records.
		// Validation is handled as part of the CSV reader.
		rec := csvRecord(data)
		recs = append(recs, rec)
	}

	// Return all records.
	return recs, nil
}

// Function to write CSV records to a CSV file using the encoding/csv package.
func writeCSVWriter(filepath string, recs []csvRecord) error {
	// Create a new file and open it for writing.
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer.
	w := csv.NewWriter(file)
	defer w.Flush()

	// Loop over slice of csvRecords and write each record to the file.
	for _, rec := range recs {
		if err := w.Write(rec); err != nil {
			return err
		}
	}

	return nil
}
