package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCSVReadWriteCycle(t *testing.T) {
	in := strings.TrimSpace(`
Car;MPG;Cylinders;Displacement;Horsepower;Weight;Acceleration;Model;Origin
STRING;DOUBLE;INT;DOUBLE;DOUBLE;DOUBLE;DOUBLE;INT;CAT
Chevrolet Chevelle Malibu;18.0;8;307.0;130.0;3504.;12.0;70;US
Buick Skylark 320;15.0;8;350.0;165.0;3693.;11.5;70;US
Plymouth Satellite;18.0;8;318.0;150.0;3436.;11.0;70;US
`)
	wantOut := strings.TrimSpace(`
Car;MPG;Cylinders;Displacement;Horsepower;Weight;Acceleration;Model;Origin
STRING;DOUBLE;INT;DOUBLE;DOUBLE;DOUBLE;DOUBLE;INT;CAT
Chevrolet Chevelle Malibu;18.000000;8;307.000000;130.000000;3504.000000;12.000000;70;US
Buick Skylark 320;15.000000;8;350.000000;165.000000;3693.000000;11.500000;70;US
Plymouth Satellite;18.000000;8;318.000000;150.000000;3436.000000;11.000000;70;US
`)

	for i := 0; i < 2; i++ {
		cars, err := ReadCarsCSV(strings.NewReader(in))
		if err != nil {
			t.Fatal(err)
		}

		buf := &bytes.Buffer{}
		if err := WriteCarsCSV(buf, cars); err != nil {
			t.Fatal(err)
		}
		gotOut := strings.TrimSpace(buf.String())
		if gotOut != wantOut {
			t.Fatalf("\ngot:\n%s\nwant:\n%s", gotOut, wantOut)
		}
		in = gotOut
	}
}
