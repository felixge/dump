package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	if data, err := ioutil.ReadFile(flag.Arg(0)); err != nil {
		return err
	} else if cars, err := ReadCarsCSV(bytes.NewReader(data)); err != nil {
		return err
	} else {
		return WriteCarsCSV(os.Stdout, cars)
	}
}

const carComma = ';'

func ReadCarsCSV(r io.Reader) ([]*Car, error) {
	cr := csv.NewReader(r)
	cr.Comma = carComma

	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}

	var cars []*Car
	for i, record := range records {
		if err := checkCarColumns(record); err != nil {
			return nil, fmt.Errorf("row=%d: %w", i+1, err)
		}

		switch i {
		case 0:
			for i, got := range record {
				if want := carColumns[i].Name; got != want {
					return nil, fmt.Errorf("unexpected header column %d: got=%q want=%q", i, got, want)
				}
			}
		case 1:
			for i, got := range record {
				if want := carColumns[i].Type; got != want {
					return nil, fmt.Errorf("unexpected header column %d: got=%q want=%q", i, got, want)
				}
			}
		default:
			car := &Car{}
			if err := car.UnmarshalRecord(record); err != nil {
				return nil, fmt.Errorf("row=%d: %w", i+1, err)
			}
			cars = append(cars, car)
		}
	}

	return cars, nil
}

func WriteCarsCSV(w io.Writer, cars []*Car) error {
	cw := csv.NewWriter(w)
	cw.Comma = carComma

	header := make([]string, len(carColumns))
	for i, col := range carColumns {
		header[i] = col.Name
	}
	cw.Write(header)

	types := make([]string, len(carColumns))
	for i, col := range carColumns {
		types[i] = col.Type
	}
	cw.Write(types)

	for _, car := range cars {
		record, err := car.MarshalRecord()
		if err != nil {
			return err
		}
		cw.Write(record)
	}
	cw.Flush()
	return cw.Error()
}

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

func (c *Car) UnmarshalRecord(record []string) error {
	if err := checkCarColumns(record); err != nil {
		return err
	}
	for i, val := range record {
		if err := carColumns[i].UnmarshalColumn(val, c); err != nil {
			return err
		}
	}
	return nil
}

func (c *Car) MarshalRecord() ([]string, error) {
	record := make([]string, len(carColumns))
	for i, col := range carColumns {
		val, err := col.MarshalColumn(c)
		if err != nil {
			return nil, err
		}
		record[i] = val
	}
	return record, nil
}

// checkCarColumns returns an error if record doesn't have the right number
// of columns.
func checkCarColumns(record []string) error {
	if got, want := len(record), len(carColumns); got != want {
		return fmt.Errorf("bad number of columns: got=%d want=%d", got, want)
	}
	return nil
}

type carColumn struct {
	Name            string
	Type            string
	UnmarshalColumn func(string, *Car) error
	MarshalColumn   func(*Car) (string, error)
}

var carColumns = []carColumn{
	{
		"Car",
		"STRING",
		func(val string, c *Car) error {
			c.Car = val
			return nil
		},
		func(c *Car) (string, error) {
			return c.Car, nil
		},
	},
	{
		"MPG",
		"DOUBLE",
		func(val string, c *Car) (err error) {
			c.MPG, err = strconv.ParseFloat(val, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%f", c.MPG), nil
		},
	},
	{
		"Cylinders",
		"INT",
		func(val string, c *Car) (err error) {
			c.Cylinders, err = strconv.ParseInt(val, 10, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%d", c.Cylinders), nil
		},
	},
	{
		"Displacement",
		"DOUBLE",
		func(val string, c *Car) (err error) {
			c.Displacement, err = strconv.ParseFloat(val, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%f", c.Displacement), nil
		},
	},
	{
		"Horsepower",
		"DOUBLE",
		func(val string, c *Car) (err error) {
			c.Horsepower, err = strconv.ParseFloat(val, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%f", c.Horsepower), nil
		},
	},
	{
		"Weight",
		"DOUBLE",
		func(val string, c *Car) (err error) {
			c.Weight, err = strconv.ParseFloat(val, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%f", c.Weight), nil
		},
	},
	{
		"Acceleration",
		"DOUBLE",
		func(val string, c *Car) (err error) {
			c.Acceleration, err = strconv.ParseFloat(val, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%f", c.Acceleration), nil
		},
	},
	{
		"Model",
		"INT",
		func(val string, c *Car) (err error) {
			c.Model, err = strconv.ParseInt(val, 10, 64)
			return
		},
		func(c *Car) (string, error) {
			return fmt.Sprintf("%d", c.Model), nil
		},
	},
	{
		"Origin",
		"CAT",
		func(val string, c *Car) error {
			c.Origin = val
			return nil
		},
		func(c *Car) (string, error) {
			return c.Origin, nil
		},
	},
}
