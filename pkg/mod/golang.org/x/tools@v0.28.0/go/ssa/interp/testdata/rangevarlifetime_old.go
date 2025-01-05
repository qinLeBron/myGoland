// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21

// goversion can be pinned to anything strictly before 1.22.

package main

func main() {
	test_init()

	// Clones from cmd/compile/internal/loopvar/testdata .
	range_esc_address()
	range_esc_closure()
	range_esc_method()
}

// pre-go1.22 all of i will have the same address.
var same = func(a [3]int) []*int {
	var r []*int
	for i := range a {
		r = append(r, &i)
	}
	return r
}([3]int{})

func test_init() {
	if len(same) != 3 {
		panic(same)
	}
	for i := range same {
		for j := range same {
			if !(same[i] == same[j]) {
				panic(same)
			}
		}
	}
	for i := range same {
		if *(same[i]) != 2 {
			panic(same)
		}
	}
}

func range_esc_address() {
	// Clone of range_esc_address.go
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0
	var is []*int
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, &i)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw ", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, pi := range is {
		sum += *pi
	}
	if sum != 9+9+9+9+9 {
		println("wrong sum, expected", 45, ", saw", sum)
		bug = true
	}
	if bug {
		panic("range_esc_address")
	}
}

func range_esc_closure() {
	// Clone of range_esc_closure.go
	var ints = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var is []func() int

	sum := 0
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, func() int {
				if i%17 == 15 {
					i++
				}
				return i
			})
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, f := range is {
		sum += f()
	}
	if sum != 9+9+9+9+9 {
		println("wrong sum, expected ", 45, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("range_esc_closure")
	}
}

type I int

func (x *I) method() int {
	return int(*x)
}

func range_esc_method() {
	// Clone of range_esc_method.go
	var ints = []I{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0
	var is []func() int
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if int(i) == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, i.method)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, m := range is {
		sum += m()
	}
	if sum != 9+9+9+9+9 {
		println("wrong sum, expected ", 45, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("range_esc_method")
	}
}
