// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"context"
	"log"
	"net/http"
	"sync"
)

// Shutdown gracefully shutsdown several http.Server's.
func Shutdown(ctx context.Context, srvs ...*http.Server) {
	var wg sync.WaitGroup

	for _, srv := range srvs {
		if srv == nil {
			continue
		}

		wg.Add(1)

		go func(srv *http.Server) {
			defer wg.Done()

			if err := srv.Shutdown(ctx); err != nil {
				log.Printf("error shutting down: %s", err)
			}
		}(srv)
	}

	wg.Wait()
}
