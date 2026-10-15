// SPDX-License-Identifier: BSD-3-Clause
// ZXC - High-performance lossless compression
//
// Copyright (c) Bertrand Lebonnois and contributors.

//go:build !linux && !windows

package zxc

import "os"

// committedBytes has no portable source here; the memory test skips.
func committedBytes() (int64, bool) { return 0, false }

func openHandles() (int, bool) {
	e, err := os.ReadDir("/dev/fd")
	return len(e), err == nil
}
