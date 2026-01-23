// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !plan9 && !js && !wasip1 && !wasm

package neterror

import (
	"errors"
	"io"
	"io/fs"
	"runtime"
	"syscall"
)

// Reports whether err resulted from reading or writing to a closed or broken pipe.
func IsClosedPipeError(err error) bool {
	expect := syscall.EPIPE
	if runtime.GOOS == "windows" {
		// 232 is Windows error code ERROR_NO_DATA, "The pipe is being closed".
		expect = syscall.Errno(232)
	}

	// EPIPE/ENOTCONN are common errors when a send fails due to a closed
	// socket. There is some platform and version inconsistency in which
	// error is returned, but the meaning is the same.
	// Libraries may also return root errors like fs.ErrClosed/io.ErrClosedPipe
	// due to a closed socket.
	return errors.Is(err, expect) ||
		errors.Is(err, syscall.ENOTCONN) ||
		errors.Is(err, fs.ErrClosed) ||
		errors.Is(err, io.ErrClosedPipe)
}
