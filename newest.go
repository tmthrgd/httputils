// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import "time"

// Newest takes multiple time.Time and returns the
// newest time.
//
// It is useful for selecting the newest modification
// time for templated resources.
func Newest(t ...time.Time) time.Time {
	var newest time.Time

	for _, t := range t {
		if t.After(newest) {
			newest = t
		}
	}

	return newest
}
