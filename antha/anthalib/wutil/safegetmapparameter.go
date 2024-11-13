// anthalib//wutil/safegetmapparameter.go: Part of the Antha language
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

func GetFloat64FromMap(m map[string]interface{}, k string) float64 {
	var f float64

	// it's easy mmmk
	v, ok := m[k]

	if ok {
		f = v.(float64)
	}

	return f
}

func GetIntFromMap(m map[string]interface{}, k string) int {
	var i int

	v, ok := m[k]

	if ok {
		i = v.(int)
	}

	return i
}

func GetStringFromMap(m map[string]interface{}, k string) string {
	var s string

	v, ok := m[k]

	if ok {
		s = v.(string)
	}

	return s
}

func GetMapFromMap(m map[string]interface{}, k string) map[string]interface{} {
	var m2 map[string]interface{}

	v, ok := m[k]

	if ok {
		m2 = v.(map[string]interface{})
	}

	return m2
}
