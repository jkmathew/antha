// antha/parser/short_test.go: Part of the Antha language
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

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains test cases for short valid and invalid programs.

package parser

import "testing"

var valids = []string{
	"package p\n",
	`package p;`,
	`package p; import "fmt"; func f() { fmt.Println("Hello, World!") };`,
	`package p; func f() { if f(T{}) {} };`,
	`package p; func f() { _ = <-chan int(nil) };`,
	`package p; func f() { _ = (<-chan int)(nil) };`,
	`package p; func f() { _ = (<-chan <-chan int)(nil) };`,
	`package p; func f() { _ = <-chan <-chan <-chan <-chan <-int(nil) };`,
	`package p; func f(func() func() func());`,
	`package p; func f(...T);`,
	`package p; func f(float, ...int);`,
	`package p; func f(x int, a ...int) { f(0, a...); f(1, a...,) };`,
	`package p; func f(int,) {};`,
	`package p; func f(...int,) {};`,
	`package p; func f(x ...int,) {};`,
	`package p; type T []int; var a []bool; func f() { if a[T{42}[0]] {} };`,
	`package p; type T []int; func g(int) bool { return true }; func f() { if g(T{42}[0]) {} };`,
	`package p; type T []int; func f() { for _ = range []int{T{42}[0]} {} };`,
	`package p; var a = T{{1, 2}, {3, 4}}`,
	`package p; func f() { select { case <- c: case c <- d: case c <- <- d: case <-c <- d: } };`,
	`package p; func f() { select { case x := (<-c): } };`,
	`package p; func f() { if ; true {} };`,
	`package p; func f() { switch ; {} };`,
	`package p; func f() { for _ = range "foo" + "bar" {} };`,
	`package p; func f() { var s []int; g(s[:], s[i:], s[:j], s[i:j], s[i:j:k], s[:j:k]) };`,
	`package p; var ( _ = (struct {*T}).m; _ = (interface {T}).m )`,
}

func TestValid(t *testing.T) {
	t.Skip("external files")
	for _, src := range valids {
		checkErrors(t, src, src)
	}
}

var invalids = []string{
	`foo /* ERROR "expected 'package'" */ !`,
	`package p; func f() { if { /* ERROR "expected operand" */ } };`,
	`package p; func f() { if ; { /* ERROR "expected operand" */ } };`,
	`package p; func f() { if f(); { /* ERROR "expected operand" */ } };`,
	`package p; func f() { if _ /* ERROR "expected boolean expression" */ = range x; true {} };`,
	`package p; func f() { switch _ /* ERROR "expected switch expression" */ = range x; true {} };`,
	`package p; func f() { for _ = range x ; /* ERROR "expected '{'" */ ; {} };`,
	`package p; func f() { for ; ; _ = range /* ERROR "expected operand" */ x {} };`,
	`package p; func f() { for ; _ /* ERROR "expected boolean or range expression" */ = range x ; {} };`,
	`package p; func f() { switch t /* ERROR "expected switch expression" */ = t.(type) {} };`,
	`package p; func f() { switch t /* ERROR "expected switch expression" */ , t = t.(type) {} };`,
	`package p; func f() { switch t /* ERROR "expected switch expression" */ = t.(type), t {} };`,
	`package p; var a = [ /* ERROR "expected expression" */ 1]int;`,
	`package p; var a = [ /* ERROR "expected expression" */ ...]int;`,
	`package p; var a = struct /* ERROR "expected expression" */ {}`,
	`package p; var a = func /* ERROR "expected expression" */ ();`,
	`package p; var a = interface /* ERROR "expected expression" */ {}`,
	`package p; var a = [ /* ERROR "expected expression" */ ]int`,
	`package p; var a = map /* ERROR "expected expression" */ [int]int`,
	`package p; var a = chan /* ERROR "expected expression" */ int;`,
	`package p; var a = []int{[ /* ERROR "expected expression" */ ]int};`,
	`package p; var a = ( /* ERROR "expected expression" */ []int);`,
	`package p; var a = a[[ /* ERROR "expected expression" */ ]int:[]int];`,
	`package p; var a = <- /* ERROR "expected expression" */ chan int;`,
	`package p; func f() { select { case _ <- chan /* ERROR "expected expression" */ int: } };`,
	`package p; func f() { _ = (<-<- /* ERROR "expected 'chan'" */ chan int)(nil) };`,
	`package p; func f() { _ = (<-chan<-chan<-chan<-chan<-chan<- /* ERROR "expected channel type" */ int)(nil) };`,
	`package p; func f() { var t []int; t /* ERROR "expected identifier on left side of :=" */ [0] := 0 };`,
	`package p; func f() { if x := g(); x = /* ERROR "expected '=='" */ 0 {}};`,
	`package p; func f() { _ = x = /* ERROR "expected '=='" */ 0 {}};`,
	`package p; func f() { _ = 1 == func()int { var x bool; x = x = /* ERROR "expected '=='" */ true; return x }() };`,
	`package p; func f() { var s []int; _ = s[] /* ERROR "expected operand" */ };`,
	`package p; func f() { var s []int; _ = s[i:j: /* ERROR "3rd index required" */ ] };`,
	`package p; func f() { var s []int; _ = s[i: /* ERROR "2nd index required" */ :k] };`,
	`package p; func f() { var s []int; _ = s[i: /* ERROR "2nd index required" */ :] };`,
	`package p; func f() { var s []int; _ = s[: /* ERROR "2nd index required" */ :] };`,
	`package p; func f() { var s []int; _ = s[: /* ERROR "2nd index required" */ ::] };`,
	`package p; func f() { var s []int; _ = s[i:j:k: /* ERROR "expected ']'" */ l] };`,
	`package p; func f() { for x /* ERROR "boolean or range expression" */ = []string {} }`,
	`package p; func f() { for x /* ERROR "boolean or range expression" */ := []string {} }`,
	`package p; func f() { for i /* ERROR "boolean or range expression" */ , x = []string {} }`,
	`package p; func f() { for i /* ERROR "boolean or range expression" */ , x := []string {} }`,
	`package p; func f() { go f /* ERROR HERE "function must be invoked" */ }`,
	`package p; func f() { defer func() {} /* ERROR HERE "function must be invoked" */ }`,
	`package p; func f() { go func() { func() { f(x func /* ERROR "expected '\)'" */ (){}) } } }`,
}

func TestInvalid(t *testing.T) {
	t.Skip("external files")
	for _, src := range invalids {
		checkErrors(t, src, src)
	}
}
