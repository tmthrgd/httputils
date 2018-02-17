// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegotiate(t *testing.T) {
	for i, tt := range []struct {
		header   []string
		offers   []string
		expected string
	}{
		// Accept
		{
			[]string{"text/html, application/xhtml+xml, application/xml;q=0.9, */*; q=0.8"},
			[]string{"application/example", "application/xml", "text/html", "application/xhtml+xml"},
			"text/html",
		},
		{
			[]string{"text/html, application/xhtml+xml, application/xml;q=0.9, */*; q=0.8 "},
			[]string{"application/example", "application/xml", "application/xhtml+xml", "text/html"},
			"application/xhtml+xml",
		},
		{
			[]string{"TeXt/HtMl, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8"},
			[]string{"application/example", "application/xml", "text/html", "application/xhtml+xml"},
			"text/html",
		},
		{
			[]string{"text/html, application/xhtml+xml, application/xml; q=0.9, */*;q=0.8"},
			[]string{"application/example", "*/*"},
			"*/*",
		},
		{
			[]string{"text/html, application/xhtml+xml, application/xml;q=0.9,*/*;q=0.8"},
			[]string{"application/example"},
			"",
		},
		{
			[]string{"text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8"},
			nil,
			"",
		},
		{
			[]string{"TeXt/HtMl,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
			[]string{"application/example", "application/xml", "TEXT/HTML", "application/xhtml+xml"},
			"TEXT/HTML",
		},
		{
			[]string{"text/html", "application/xhtml+xml", "application/xml;q=0.9", "*/*;q=0.8"},
			[]string{"application/example", "application/xml", "text/html", "application/xhtml+xml"},
			"text/html",
		},
		{
			[]string{"text/html", "application/xhtml+xml", "application/xml;q=0.9", "*/*;q=0.8"},
			[]string{"application/example", "application/xml", "application/xhtml+xml"},
			"application/xhtml+xml",
		},
		{
			[]string{"text/html", "application/xhtml+xml", "application/xml;   q=0.9  ", " */*;  q=0.8 "},
			[]string{"application/example", "application/xml"},
			"application/xml",
		},
		{
			[]string{" text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8"},
			[]string{"application/example", "application/xml", "text/html", "application/xhtml+xml"},
			"text/html",
		},
		{
			nil,
			[]string{"application/example"},
			"",
		},

		// Accept-Charset
		{
			[]string{"utf-8, iso-8859-1;q=0.5"},
			[]string{"invalid", "iso-8859-1", "utf-8"},
			"utf-8",
		},
		{
			[]string{"utf-8, iso-8859-1;q=0.5"},
			[]string{"invalid", "iso-8859-1"},
			"iso-8859-1",
		},
		{
			[]string{"utf-8", "iso-8859-1;q=0.5"},
			[]string{"invalid", "iso-8859-1"},
			"iso-8859-1",
		},
		{
			[]string{"utf-8, iso-8859-1;q=0.5"},
			[]string{"invalid"},
			"",
		},
		{
			[]string{"utf-8, iso-8859-1;q=0.5"},
			nil,
			"",
		},
		{
			nil,
			[]string{"invalid"},
			"",
		},

		// Accept-Encoding
		{
			[]string{"deflate, gzip;q=1.0, *;q=0.5"},
			[]string{"invalid", "gzip", "deflate"},
			"gzip",
		},
		{
			[]string{"deflate, gzip;q=1.0, *;q=0.5"},
			[]string{"invalid", "*", "deflate"},
			"deflate",
		},
		{
			[]string{"deflate, gzip;q=1.0, *;q=0.5"},
			[]string{"invalid", "*"},
			"*",
		},
		{
			[]string{"deflate, gzip;q=1.0, *;q=0.5, identity;q=0.1"},
			[]string{"invalid", "identity"},
			"identity",
		},
		{
			[]string{"deflate;q=0.7, gzip;q=0.9, *;q=0.5, identity;q=0.1", "br"},
			[]string{"invalid", "gzip", "identity", "br"},
			"br",
		},
		{
			[]string{"deflate, gzip;q=1.0, *;q=0.5"},
			[]string{"invalid"},
			"",
		},
		{
			[]string{"deflate, gzip;q=1.0, *;q=0.5"},
			nil,
			"",
		},
		{
			nil,
			[]string{"invalid"},
			"",
		},

		// Accept-Language
		{
			[]string{"fr-CH, fr;q=0.9, en;q=0.8", "de;q=0.7, *;q=0.5"},
			[]string{"invalid", "en", "de", "fr"},
			"fr",
		},
		{
			[]string{"fr-CH, fr;q=0.9, en;q=0.8", "de;q=0.7, *;q=0.5"},
			[]string{"invalid", "en", "de", "fr-CH"},
			"fr-CH",
		},
		{
			[]string{"fr-CH, fr;q=0.9, en;q=0.8", "de;q=0.7, *;q=0.5"},
			[]string{"invalid", "de", "en-US", "en-GB", "en"},
			"en",
		},
		{
			[]string{"fr-CH, fr;q=0.9, en;q=0.8", "de;q=0.7, *;q=0.5"},
			[]string{"invalid"},
			"",
		},
		{
			[]string{"fr-CH, fr;q=0.9, en;q=0.8", "de;q=0.7, *;q=0.5"},
			nil,
			"",
		},
		{
			nil,
			[]string{"invalid"},
			"",
		},
	} {
		assert.Equal(t, tt.expected, Negotiate(http.Header{
			"X-Accept": tt.header,
		}, "X-Accept", tt.offers...), "#%d %q", i, tt)
	}
}
