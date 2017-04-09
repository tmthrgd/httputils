// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httputils

import (
	"net"
	"time"
)

// TCPKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections.
//
// It is an unexported net/http type that is used by
// net/http.ListenAndServe and net/http.ListenAndServeTLS so dead
// TCP connections (e.g. closing laptop mid-download) eventually
// go away.
func TCPKeepAliveListener(ln net.Listener) net.Listener {
	return &tcpKeepAliveListener{ln.(*net.TCPListener)}
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
