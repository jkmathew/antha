// wtype/warning.go: Part of the Antha language
// Copyright (C) 2014 the Antha authors. All rights reserved.
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

package wtype

import "fmt"

// Warning is a representation of a non fatal error in Antha which implements the golang error interface.
type Warning string

// Error returns a string error message.
func (w Warning) Error() string {
	return string(w)
}

// NewWarning generate a new Warning using the default formats for its arguments.
// Spaces are added between operands when neither is a string.
func NewWarning(arguments ...interface{}) Warning {
	return Warning(fmt.Sprint(arguments...))
}

// NewWarningf generates a new Warning from a formatted string.
func NewWarningf(format string, arguments ...interface{}) Warning {
	return Warning(fmt.Sprintf(format, arguments...))
}
