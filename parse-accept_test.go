// Copyright 2013 The Go Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.

package httputils

import (
	"net/http"
	"reflect"
	"testing"
)

var parseAcceptTests = []struct {
	s        string
	expected []acceptSpec
}{
	{"text/html", []acceptSpec{{"text/html", 1}}},
	{"text/html; q=0", []acceptSpec{{"text/html", 0}}},
	{"text/html; q=0.0", []acceptSpec{{"text/html", 0}}},
	{"text/html; q=1", []acceptSpec{{"text/html", 1}}},
	{"text/html; q=1.0", []acceptSpec{{"text/html", 1}}},
	{"text/html; q=0.1", []acceptSpec{{"text/html", 0.1}}},
	{"text/html;q=0.1", []acceptSpec{{"text/html", 0.1}}},
	{"text/html, text/plain", []acceptSpec{{"text/html", 1}, {"text/plain", 1}}},
	{"text/html; q=0.1, text/plain", []acceptSpec{{"text/html", 0.1}, {"text/plain", 1}}},
	{"iso-8859-5, unicode-1-1;q=0.8,iso-8859-1", []acceptSpec{{"iso-8859-5", 1}, {"unicode-1-1", 0.8}, {"iso-8859-1", 1}}},
	{"iso-8859-1", []acceptSpec{{"iso-8859-1", 1}}},
	{"*", []acceptSpec{{"*", 1}}},
	{"da, en-gb;q=0.8, en;q=0.7", []acceptSpec{{"da", 1}, {"en-gb", 0.8}, {"en", 0.7}}},
	{"da, q, en-gb;q=0.8", []acceptSpec{{"da", 1}, {"q", 1}, {"en-gb", 0.8}}},
	{"image/png, image/*;q=0.5", []acceptSpec{{"image/png", 1}, {"image/*", 0.5}}},
	{"text/html; Q=1", []acceptSpec{{"text/html", 1}}},
	{"text/html; q=0.123", []acceptSpec{{"text/html", 0.123}}},
	{" text/html", []acceptSpec{{"text/html", 1}}},
	{"text/html ", []acceptSpec{{"text/html", 1}}},

	// bad cases
	{"value1; q=0.1.2", []acceptSpec{{"value1", 0.1}}},
	{"da, en-gb;q=foo", []acceptSpec{{"da", 1}}},

	// failing cases
	{"text/html; q=0.1234", nil},
}

func TestParseAccept(t *testing.T) {
	for _, tt := range parseAcceptTests {
		header := http.Header{"Accept": {tt.s}}
		actual := parseAccept(header, "Accept")
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("ParseAccept(h, %q)=%v, want %v", tt.s, actual, tt.expected)
		}
	}
}
