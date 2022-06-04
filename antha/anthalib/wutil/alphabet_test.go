package wutil

import (
	"fmt"
	"testing"
)

func TestNumToAlpha(t *testing.T) {
	if e, f := "A", NumToAlpha(1); e != f {
		t.Errorf("expected %s found %s", e, f)
	}
	if e, f := "AA", NumToAlpha(1+1*26); e != f {
		t.Errorf("expected %s found %s", e, f)
	}
	if e, f := "BA", NumToAlpha(1+2*26); e != f {
		t.Errorf("expected %s found %s", e, f)
	}
	if e, f := "CA", NumToAlpha(1+3*26); e != f {
		t.Errorf("expected %s found %s", e, f)
	}
	if e, f := "", NumToAlpha(-1); e != f {
		t.Errorf("expected %s found %s", e, f)
	}
}

func ExampleNumToAlpha() {
	for _, v := range []int{10, 1, 2, 27, 100} {
		fmt.Printf("%d: %s\n", v, NumToAlpha(v))
	}

	// Output:
	// 10: J
	// 1: A
	// 2: B
	// 27: AA
	// 100: CV
}

func TestAlphaToNum(t *testing.T) {
	if e, f := 1, AlphaToNum("A"); e != f {
		t.Errorf("expected %d found %d", e, f)
	}
	if e, f := 1+1*26, AlphaToNum("AA"); e != f {
		t.Errorf("expected %d found %d", e, f)
	}
	if e, f := 1+2*26, AlphaToNum("BA"); e != f {
		t.Errorf("expected %d found %d", e, f)
	}
	if e, f := 1+3*26, AlphaToNum("CA"); e != f {
		t.Errorf("expected %d found %d", e, f)
	}
	if e, f := 0, AlphaToNum("someBadCharacters"); e != f {
		t.Errorf("expected %d found %d", e, f)
	}
}

func ExampleAlphaToNum() {
	for _, v := range []string{"J", "A", "B", "CV", "AA"} {
		fmt.Printf("%s: %d\n", v, AlphaToNum(v))
	}

	// Output:
	// J: 10
	// A: 1
	// B: 2
	// CV: 100
	// AA: 27
}
