// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMIMETypeMatches(t *testing.T) {
	for _, tt := range []struct {
		mime  string
		types []string
	}{
		{"application/example+json; charset=utf-8", []string{"application/example+json"}},
		{"application/example+json;charset=utf-8", []string{"application/example+json"}},
		{"application/example+json", []string{"application/example+json"}},
		{"application/example", []string{"application/example"}},
		{"application/example", []string{"application/*"}},
		{"application/example", []string{"*/*"}},
	} {
		assert.True(t, MIMETypeMatches(tt.mime, tt.types), "%q", tt)
	}

	for _, tt := range []struct {
		mime  string
		types []string
	}{
		{"application/example; charset-utf-8", []string{"application/example; charset=utf-8"}},
		{"application/example ;charset-utf-8", []string{"application/example"}}, // ?
		{"application/example", []string{"application/example; charset=utf-8"}},
		{"application/example", []string{"application/example+json"}},
		{"application/example+json", []string{"application/example"}},
		{"application/example", []string{"example/*"}},
		{"", []string{"example/*"}},
		{"", []string{""}},
		{"application/example", []string{}},
		{"application/example", nil},
		{"application/example ", []string{"application/example"}}, // ?
		{" application/example", []string{"application/example"}}, // ?
		{"application/example", []string{" application/example"}},
		{"application/example", []string{"application/example "}},
	} {
		assert.False(t, MIMETypeMatches(tt.mime, tt.types), "%q", tt)
	}
}
