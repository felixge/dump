package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}

// Car holds one row of a cars.csv file.
type Car struct {
	Car          string
	MPG          float64
	Cylinders    int64
	Displacement float64
	Horsepower   float64
	Weight       float64
	Acceleration float64
	Model        int64
	Origin       string
}

// ReadCarsCSV reads the CSV formatted cars from the given reader or returns an
// error.
func ReadCarsCSV(r io.Reader) ([]*Car, error) {
	gocsv.FailIfUnmatchedStructTags = true
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return &SkipReader{r: r}
	})

	var cars []*Car
	if err := gocsv.Unmarshal(r, &cars); err != nil {
		return nil, err
	}
	return cars, nil
}

type SkipReader struct {
	r *csv.Reader
	i int
}

func (s *SkipReader) Read() ([]string, error) {
	if s.i == 1 {
		s.r.Read()
	}
	s.i++
	return s.r.Read()
}

func (s *SkipReader) ReadAll() ([][]string, error) {
	var records [][]string
	for {
		if record, err := s.Read(); err == io.EOF {
			return records, nil
		} else if err != nil {
			return nil, err
		} else {
			records = append(records, record)
		}
	}
}

// Write writes the given cars in CSV format to the given writer or returns an
// error.
func WriteCarsCSV(w io.Writer, cars []*Car) error {
	return nil
}
