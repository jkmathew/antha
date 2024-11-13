// antha/ast/print_test.go: Part of the Antha language
// Copyright (C) 2014 The Antha authors. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
//
// For more information relating to the software or licensing issues please
// contact license@jkmathew.org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"bytes"
	"strings"
	"testing"
)

var tests = []struct {
	x interface{} // x is printed as s
	s string
}{
	// basic types
	{nil, "0  nil"},
	{true, "0  true"},
	{42, "0  42"},
	{3.14, "0  3.14"},
	{1 + 2.718i, "0  (1+2.718i)"},
	{"foobar", "0  \"foobar\""},

	// maps
	{map[Expr]string{}, `0  map[ast.Expr]string (len = 0) {}`},
	{map[string]int{"a": 1},
		`0  map[string]int (len = 1) {
		1  .  "a": 1
		2  }`},

	// pointers
	{new(int), "0  *0"},

	// arrays
	{[0]int{}, `0  [0]int {}`},
	{[3]int{1, 2, 3},
		`0  [3]int {
		1  .  0: 1
		2  .  1: 2
		3  .  2: 3
		4  }`},
	{[...]int{42},
		`0  [1]int {
		1  .  0: 42
		2  }`},

	// slices
	{[]int{}, `0  []int (len = 0) {}`},
	{[]int{1, 2, 3},
		`0  []int (len = 3) {
		1  .  0: 1
		2  .  1: 2
		3  .  2: 3
		4  }`},

	// structs
	{struct{}{}, `0  struct {} {}`},
	{struct{ x int }{007}, `0  struct { x int } {}`},
	{struct{ X, y int }{42, 991},
		`0  struct { X int; y int } {
		1  .  X: 42
		2  }`},
	{struct{ X, Y int }{42, 991},
		`0  struct { X int; Y int } {
		1  .  X: 42
		2  .  Y: 991
		3  }`},
}

// Split s into lines, trim whitespace from all lines, and return
// the concatenated non-empty lines.
func trim(s string) string {
	lines := strings.Split(s, "\n")
	i := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			lines[i] = line
			i++
		}
	}
	return strings.Join(lines[0:i], "\n")
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	for _, test := range tests {
		buf.Reset()
		if err := Fprint(&buf, nil, test.x, nil); err != nil {
			t.Errorf("Fprint failed: %s", err)
		}
		if s, ts := trim(buf.String()), trim(test.s); s != ts {
			t.Errorf("got:\n%s\nexpected:\n%s\n", s, ts)
		}
	}
}
