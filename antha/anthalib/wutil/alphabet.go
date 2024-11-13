// anthalib//wutil/displaymap.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
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

// Utility package
package wutil

const (
	numLetters = 26
)

// Convert n to a string: 1 to "A", 27 to "AA", 53 to "BA", ...
//
// Returns "" when n < 1.
func NumToAlpha(n int) string {
	var s []byte

	if n < 1 {
		return ""
	}

	for {
		n -= 1
		r := (n % numLetters)
		s = append(s, byte('A'+r))
		n /= numLetters
		if n <= 0 {
			break
		}
	}

	var t []byte

	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}

	return string(t)
}

// Convert string to n: "A" to 1, 53 to "BA", ... Returns 0 when s contains [^A-Z]
//
// Inverse of NumToAlpha.
func AlphaToNum(s string) int {
	v := 0

	for _, b := range s {
		off := int(b - 'A')
		if off < 0 || off >= numLetters {
			return 0
		}
		v *= numLetters
		v += off + 1
	}

	return v
}
