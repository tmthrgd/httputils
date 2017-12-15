// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewest(t *testing.T) {
	assert.Equal(t,
		Newest(time.Unix(0, 2), time.Unix(0, 1), time.Unix(0, 5),
			time.Unix(0, 4), time.Unix(0, 3), time.Unix(0, 3)),
		time.Unix(0, 5))
	assert.Equal(t, Newest(), time.Time{})
	assert.Equal(t, Newest(time.Unix(0, 0)), time.Unix(0, 0))
}

func TestNewestModTime(t *testing.T) {
	require.NotPanics(t, func() {
		NewestModTime(modTimeInfo(time.Unix(0, 0)))
	})

	assert.Equal(t,
		NewestModTime(
			modTimeInfo(time.Unix(0, 2)),
			modTimeInfo(time.Unix(0, 1)),
			modTimeInfo(time.Unix(0, 5)),
			modTimeInfo(time.Unix(0, 4)),
			modTimeInfo(time.Unix(0, 3)),
			modTimeInfo(time.Unix(0, 3))),
		time.Unix(0, 5))
	assert.Equal(t, NewestModTime(), time.Time{})
	assert.Equal(t, NewestModTime(modTimeInfo(time.Unix(0, 0))), time.Unix(0, 0))
}

type modTimeInfo time.Time

func (mt modTimeInfo) ModTime() time.Time {
	return time.Time(mt)
}

func (modTimeInfo) Name() string      { panic("should not be called") }
func (modTimeInfo) Size() int64       { panic("should not be called") }
func (modTimeInfo) Mode() os.FileMode { panic("should not be called") }
func (modTimeInfo) IsDir() bool       { panic("should not be called") }
func (modTimeInfo) Sys() interface{}  { panic("should not be called") }
