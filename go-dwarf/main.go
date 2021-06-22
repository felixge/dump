package main

import (
	"debug/dwarf"
	"debug/macho"
	"flag"
	"fmt"
	"os"
)

type Foo struct {
	A string
	B int
}

func main() {
	f := Foo{A: "hello", B: 42}
	fmt.Printf("%#v\n", f)
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		return err
	}
	mFile, err := macho.NewFile(file)
	if err != nil {
		return err
	}
	//for _, s := range mFile.Sections {
	//if strings.HasPrefix(s.Name, "__debug_") {
	//continue
	//}
	//b, err := s.Data()
	//if err != nil {
	//return err
	//}
	//if len(b) >= 12 && string(b[:4]) == "ZLIB" {
	//dlen := binary.BigEndian.Uint64(b[4:12])
	//dbuf := make([]byte, dlen)
	//r, err := zlib.NewReader(bytes.NewBuffer(b[12:]))
	//if err != nil {
	//return err
	//}
	//if _, err := io.ReadFull(r, dbuf); err != nil {
	//return err
	//}
	//if strings.Contains(string(dbuf), "main.Foo") {
	//fmt.Printf("%#v\n", s.Name)
	//}
	////fmt.Printf("%d\n", len(dbuf))
	//}
	//}
	d, err := mFile.DWARF()
	if err != nil {
		return err
	}
	r := d.Reader()
	if err != nil {
		return err
	}
	for {
		e, err := r.Next()
		if err != nil {
			return err
		} else if e == nil {
			break
		}
		if e.Tag == dwarf.TagStructType {
			tt, err := d.Type(e.Offset)
			if err != nil {
				return err
			}
			st := tt.(*dwarf.StructType)
			if st.StructName == "runtime.g" {
				fmt.Printf("%s\n", st.StructName)
				for _, sf := range st.Field {
					fmt.Printf(" %s: %d\n", sf.Name, sf.Type.Size())
				}
			}
		}
		//r.Seek()
		//t := e.Val(dwarf.AttrType)
		//if t != nil {
		//tt, err := d.Type(e.Offset)
		//if err != nil {
		//return err
		//}
		//if tt.String() == "g" {
		//fmt.Printf("%#v\n", tt)
		//}
		//}
		//fmt.Printf("%#v\n", t)
		//if e.Tag == dwarf.TagStructType {
		//for _, f := range e.Field {
		//fmt.Printf("%#v\n", f)
		//}
		//}
	}
	return nil
}
