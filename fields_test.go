package struc

import (
	"bytes"
	"reflect"
	"testing"
)

var refVal = reflect.ValueOf(reference)

func TestFieldsParse(t *testing.T) {
	if _, err := parseFields(refVal); err != nil {
		t.Fatal(err)
	}
}

func TestFieldsString(t *testing.T) {
	fields, _ := parseFields(refVal)
	fields.String()
}

type sizefromStruct struct {
	Size1 uint `struc:"sizeof=Var1"`
	Var1  []byte
	Size2 int `struc:"sizeof=Var2"`
	Var2  []byte
}

func TestFieldsSizefrom(t *testing.T) {
	var test = sizefromStruct{
		Var1: []byte{1, 2, 3},
		Var2: []byte{4, 5, 6},
	}
	var buf bytes.Buffer
	err := Pack(&buf, &test)
	if err != nil {
		t.Fatal(err)
	}
	err = Unpack(&buf, &test)
	if err != nil {
		t.Fatal(err)
	}
}

type sizefromStructBad struct {
	Size1 string `struc:"sizeof=Var1"`
	Var1  []byte
}

func TestFieldsSizefromBad(t *testing.T) {
	var test = &sizefromStructBad{Var1: []byte{1, 2, 3}}
	var buf bytes.Buffer
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("failed to panic on bad sizeof type")
		}
	}()
	Pack(&buf, &test)
}

type offsetfromStruct struct {
	Size1   uint `struc:"sizeof=Var1"`
	Offset2 uint `struc:"offsetof=Var2"`
	Size2   uint `struc:"sizeof=Var2"`
	Var1    []byte
	Var2    []byte
}

func TestFieldsOffsetfrom(t *testing.T) {
	var test = offsetfromStruct{
		Var1: []byte{1, 2, 3},
		Var2: []byte{4, 5, 6},
	}
	var buf bytes.Buffer
	err := Pack(&buf, &test)
	if err != nil {
		t.Fatal(err)
	}
	err = Unpack(&buf, &test)
	if err != nil {
		t.Fatal(err)
	}
}

type offsetfromStructBad struct {
	Size1   uint   `struc:"sizeof=Var1"`
	Offset1 string `struc:"offsetof=Var1"`
	Var1    []byte
}

func TestFieldsOffsetfromBad(t *testing.T) {
	var test = &offsetfromStructBad{Var1: []byte{1, 2, 3}}
	var buf bytes.Buffer
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("failed to panic on bad offset type")
		}
	}()
	Pack(&buf, &test)
}

type StructWithinArray struct {
	a uint32
}

type StructHavingArray struct {
	Props [1]StructWithinArray `struc:"[1]StructWithinArray"`
}

func TestStrucArray(t *testing.T) {
	var buf bytes.Buffer
	a := &StructHavingArray{[1]StructWithinArray{}}
	err := Pack(&buf, a)
	if err != nil {
		t.Fatal(err)
	}
	b := &StructHavingArray{}
	err = Unpack(&buf, b)
	if err != nil {
		t.Fatal(err)
	}
}
