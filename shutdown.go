// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package httputils

import (
	"context"
	"net/http"
	"sync"
)

// Shutdown gracefully shutsdown several http.Server's.
func Shutdown(ctx context.Context, srvs ...*http.Server) error {
	var wg sync.WaitGroup
	errs := make([]error, len(srvs))

	for i, srv := range srvs {
		if srv == nil {
			continue
		}

		wg.Add(1)

		go func(srv *http.Server, err *error) {
			defer wg.Done()
			*err = srv.Shutdown(ctx)
		}(srv, &errs[i])
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
