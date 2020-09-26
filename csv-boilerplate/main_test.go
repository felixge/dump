package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCSVReadWriteCycle(t *testing.T) {
	data, err := ioutil.ReadFile("cars.csv")
	if err != nil {
		t.Fatal(err)
	}

	var prevWrite []byte
	for i := 0; i < 2; i++ {
		cars, err := ReadCarsCSV(bytes.NewReader(data))
		if err != nil {
			t.Fatal()
		} else if got, want := len(cars), 406; got != want {
			t.Fatalf("got=%d want=%d", got, want)
		} else if got, want := cars[0].Car, "Chevrolet Chevelle Malibu"; got != want {
			t.Fatalf("got=%q want=%q", got, want)
		}

		write := &bytes.Buffer{}
		if err := WriteCarsCSV(write, cars); err != nil {
			t.Fatal(err)
		}

		if i == 0 {
			prevWrite = write.Bytes()
		} else if !bytes.Equal(write.Bytes(), prevWrite) {
			t.Fatal("read write cycle failed")
		}
	}

}
