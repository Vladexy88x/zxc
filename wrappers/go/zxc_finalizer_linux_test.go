// SPDX-License-Identifier: BSD-3-Clause
// ZXC - High-performance lossless compression
//
// Copyright (c) Bertrand Lebonnois and contributors.

package zxc

import (
	"os"
	"strconv"
	"strings"
)

// committedBytes returns the process's data segment (VmData). Resident memory
// would miss a leak: the library reserves buffers it may never touch.
func committedBytes() (int64, bool) {
	b, err := os.ReadFile("/proc/self/status")
	if err != nil {
		return 0, false
	}
	for _, line := range strings.Split(string(b), "\n") {
		if f := strings.Fields(line); len(f) >= 2 && f[0] == "VmData:" {
			kb, err := strconv.ParseInt(f[1], 10, 64)
			return kb << 10, err == nil
		}
	}
	return 0, false
}

func openHandles() (int, bool) {
	e, err := os.ReadDir("/proc/self/fd")
	return len(e), err == nil
}
