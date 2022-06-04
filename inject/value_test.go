package inject

import (
	"reflect"
	"testing"
)

func TestAssign(t *testing.T) {
	type X struct {
		A int
		B string
	}
	type SuperX struct {
		A int
		B string
		C int
	}
	type SubX struct {
		B int
	}

	var x1, x2 X
	var ipx2 interface{} = &x2
	var ix2 interface{} = x2
	var superx SuperX
	var subx SubX

	if err := AssignableTo(&x1, &x2); err != nil {
		t.Error(err)
	}
	if err := AssignableTo(x1, &x2); err != nil {
		t.Error(err)
	}
	if err := AssignableTo(x1, x2); err == nil {
		t.Errorf("expected error assigning to value")
	}

	if err := AssignableTo(x1, ipx2); err != nil {
		t.Error(err)
	}

	if err := AssignableTo(x1, ix2); err == nil {
		t.Errorf("expected error assigning to value")
	}

	if err := AssignableTo(x1, &superx); err != nil {
		t.Error(err)
	}

	if err := AssignableTo(x1, &subx); err == nil {
		t.Errorf("expected error assigning to value")
	}

	x1.A = 3
	x1.B = "hello"
	if err := Assign(x1, &x2); err != nil {
		t.Error(err)
	} else if x2.A != 3 || x2.B != "hello" {
		t.Errorf("expecting %v got %v", x1, x2)
	}
}

func TestConcat(t *testing.T) {
	type Alpha struct {
		A  int
		AA string
	}
	type Beta struct {
		B  int
		BB string
	}

	a := MakeValue(Alpha{A: 2})
	b := MakeValue(Beta{})

	if c, err := a.Concat(b); err != nil {
		t.Error(err)
	} else if l := len(a); l != 2 {
		t.Errorf("expected original value to remain the same: %d", l)
	} else if l := len(c); l != 4 {
		t.Errorf("expected 4 fields in result of concat: %d", l)
	} else if _, ok := c["A"]; !ok {
		t.Errorf("missing field in concat result")
	} else if c["A"] = 3; a["A"] == c["A"] {
		t.Errorf("no copying")
	}
}

func TestMakeValue(t *testing.T) {
	type Alpha struct {
		A int
		B string
	}

	var a, b Alpha

	a.A = 1
	a.B = "hello"
	b = a

	if !reflect.DeepEqual(MakeValue(a), MakeValue(&b)) {
		t.Errorf("expecting %v to equal %v", a, b)
	}
}
