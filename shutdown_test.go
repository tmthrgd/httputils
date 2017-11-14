// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShutdown(t *testing.T) {
	var srv1, srv2 http.Server

	closed1 := make(chan struct{})
	srv1.RegisterOnShutdown(func() {
		close(closed1)
	})

	closed2 := make(chan struct{})
	srv2.RegisterOnShutdown(func() {
		close(closed2)
	})

	assert.NoError(t, Shutdown(context.Background(), &srv1, nil, &srv2))

	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	select {
	case <-closed1:
	case <-timer.C:
		assert.FailNow(t, "timed out waiting for close")
	}

	select {
	case <-closed2:
	case <-timer.C:
		assert.FailNow(t, "timed out waiting for close")
	}
}
