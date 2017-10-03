package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Capitalize so they can be exported
type Zip struct {
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

// Can return multiple values in a function
func LoadZips(filePath string) (ZipSlice, error) {
	f, err := os.Open(filePath)
	// Good practice to check if err != nil right after possible failed call
	if err != nil {
		return nil, fmt.Errorf("Error opening file %v", err)
	}

	reader := csv.NewReader(f)
	// _ means ignore
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("Error reading header row %v", err)
	}

	// Can be more efficient by preallocating the slice
	zips := make(ZipSlice, 0, 43000)
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			return zips, nil
		}

		if err != nil {
			return nil, fmt.Errorf("Error reading record: %v", err)
		}

		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}

}
