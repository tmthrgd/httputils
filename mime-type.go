// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import "strings"

// MIMETypeMatches returns whether a MIME type is found
// within a given list.
//
// MIME types are compared in a case-sensitive manner.
//
// A MIME type without any subtype will match any more
// precise MIME type, i.e. image/* will match image/png,
// image/svg, image/gif and any other image types.
//
// Any directives that may be present in the MIME type will
// be skipped when comparing, i.e. text/html will match
// 'text/html; charset=utf-8'.
func MIMETypeMatches(mime string, types []string) bool {
	// 'text/plain; charset=utf-8' is a valid MIME
	// type which would not match the string
	// 'text/plain'.
	//
	// We only test the actual type and skip any
	// directives that may be present.
	if i := strings.IndexByte(mime, ';'); i != -1 {
		mime = mime[:i]
	}

	// Should we trim whitespace from mime or expect it
	// to be well formed?
	//	mime = strings.TrimSpace(mime)

	// An empty string can't match anything.
	if mime == "" {
		return false
	}

	for _, typ := range types {
		// An exact match.
		if typ == mime {
			return true
		}

		// Any MIME type.
		if typ == "*/*" {
			return true
		}

		// A MIME type, but without any subtype.
		if strings.HasSuffix(typ, "/*") &&
			strings.HasPrefix(mime, typ[:len(typ)-1]) {
			return true
		}
	}

	return false
}
