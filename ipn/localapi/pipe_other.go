// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !plan9

package localapi

import (
	"errors"
	"io"
	"runtime"
	"syscall"
)

// IsClosedPipeError reports if err resulted from reading or writing to a closed or broken pipe.
func IsClosedPipeError(err error) bool {
	expect := syscall.EPIPE
	if runtime.GOOS == "windows" {
		// 232 is Windows error code ERROR_NO_DATA, "The pipe is being closed".
		expect = syscall.Errno(232)
	}
	if errors.Is(err, expect) {
		return true
	}

	return errors.Is(err, io.ErrClosedPipe)
}
