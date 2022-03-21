// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import "golang.org/x/exp/constraints"

// EqualsComparator returns whether the values are equals
type EqualsComparator[T any] func(a, b T) bool

// Comparator should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator[T any] func(a, b T) int

// ComparableEqualsComparator provides a basic equals comparison for comparable
func ComparableEqualsComparator[T comparable](a, b T) bool {
	return a == b
}

// OrderedComparator provides a basic comparison for constraints.Ordered
func OrderedComparator[T constraints.Ordered](a, b T) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// StringComparator provides a fast comparison on strings
func StringComparator(s1, s2 string) int {
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// AnyComparator provides a comparator that always assumes the second value is higher
func AnyComparator(s1, s2 interface{}) int {
	return 1
}
