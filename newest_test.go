// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewest(t *testing.T) {
	assert.Equal(t,
		Newest(time.Unix(0, 2), time.Unix(0, 1), time.Unix(0, 5),
			time.Unix(0, 4), time.Unix(0, 3), time.Unix(0, 3)),
		time.Unix(0, 5))
	assert.Equal(t, Newest(), time.Time{})
	assert.Equal(t, Newest(time.Unix(0, 0)), time.Unix(0, 0))
}
