package graph

import (
	"testing"
)

func TestDisjointSet(t *testing.T) {
	ds := NewDisjointSet()

	ds.Union("a", "b")
	ds.Union("c", "d")

	if l := ds.NumNodes(); l != 4 {
		t.Errorf("expected 4 nodes found %d", l)
	} else if ar, br := ds.Find("a"), ds.Find("b"); ar != br {
		t.Errorf("%q not equal to %q", ar, br)
	} else if cr, dr := ds.Find("c"), ds.Find("d"); cr != dr {
		t.Errorf("%q not equal to %q", cr, dr)
	} else if ar == cr {
		t.Errorf("%q equal to %q", ar, cr)
	}

	ds.Union("b", "c")
	if ar, br, cr, dr := ds.Find("a"), ds.Find("b"), ds.Find("c"), ds.Find("d"); ar != br || br != cr || cr != dr {
		t.Errorf("not equal all equal %q %q %q %q", ar, br, cr, dr)
	}
}
