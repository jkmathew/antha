package wtype

import (
	"fmt"
	"testing"
)

func TestLHError(t *testing.T) {
	e := LHError(LH_OK, "")

	if e.Error() != "0 (LH_OK) : no problem : " {
		t.Fatal("LH Error format changed... this may break stuff")
	}

	e = LHError(LH_ERR_NO_DECK_SPACE, "can't fit tip box in")

	if e.Error() != "1 (LH_ERR_NO_DECK_SPACE) : insufficient deck space to fit all required items; this may be due to constraints : can't fit tip box in" {
		t.Fatal("LH Error format changed... this may break stuff")
	}
}

func TestLHErrorCodeFromErr(t *testing.T) {
	e := LHError(LH_ERR_DIRE, "yes")
	c := LHErrorCodeFromErr(e)

	if c != LH_ERR_DIRE {
		t.Fatal(fmt.Sprintf("Unexpected errorcode mismatch: got %d expected %d", c, LH_ERR_DIRE))
	}
}

func TestLHErrorIsInternal(t *testing.T) {
	e := LHError(LH_OK, "")

	if LHErrorIsInternal(e) {
		t.Fatal(fmt.Sprint("Error ", e, " is not an internal error but is reported as one"))
	}

	e = LHError(LH_ERR_DIRE, "YES")

	if !LHErrorIsInternal(e) {
		t.Fatal(fmt.Sprint("ERROR ", e, " is an internal error but is not reported as one"))
	}

}
