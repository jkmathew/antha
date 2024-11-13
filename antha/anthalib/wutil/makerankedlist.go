// anthalib//wutil/makerankedlist.go: Part of the Antha language
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

package wutil

import (
	"sort"
)

//
// makerankedlist.go
//
// given a slice containing integers the function
// "tidies up" by converting to ranks starting from 0
//

func MakeRankedList(arr []int) []int {
	ret := make([]int, len(arr))

	a2 := make([]int, len(arr))
	copy(a2, arr)

	// first sort ascending

	sort.Ints(a2)
	tmp := make(map[int]int, len(arr))

	// we don't follow the .5 convention, instead if there are
	// non-distinct integers we give them the lowest available rank

	x := 0
	y := 1
	l := 0

	for i, v := range a2 {
		if i != 0 && v != l {
			x += y
			y = 1
		} else if i != 0 && v == l {
			y += 1
		}

		tmp[v] = x
		l = v
	}

	for i, v := range arr {
		ret[i] = tmp[v]
	}

	return ret
}
