// Copyright 2017 The pkgutil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringsutil

import (
	"sort"
	"strings"
)

// BinarySearch reports whether s is within a with binary search.
func BinarySearch(a []string, s string) bool {
	sort.Strings(a)
	i := sort.SearchStrings(a, s)
	return i < len(a) && a[i] == s
}

// MapSearch reports whether s is within a with map search.
func MapSearch(a []string, s string) bool {
	set := make(map[string]bool, len(a))
	for _, v := range a {
		set[v] = true
	}

	return set[s]
}

// IndexSlice returns the index of the first instance of s in a slice, or -1 if s is not present in a slice.
func IndexSlice(a []string, substr string) int {
	for i, s := range a {
		if s == substr {
			return i
		}
	}
	return -1
}

// IndexContainsSlice returns the index of the first instance of substr is within a slice, or -1 if s is not present in a slice.
func IndexContainsSlice(a []string, substr string) int {
	for i, s := range a {
		if strings.Contains(s, substr) {
			return i
		}
	}
	return -1
}
