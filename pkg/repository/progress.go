// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package repository

import (
	"fmt"
	"os"

	"github.com/cheggaaa/pb/v3"
	"golang.org/x/term"
)

// DisableProgress implement the DownloadProgress interface and disable download progress
type DisableProgress struct{}

// Start implement the DownloadProgress interface
func (d DisableProgress) Start(url string, size int64) {}

// SetCurrent implement the DownloadProgress interface
func (d DisableProgress) SetCurrent(size int64) {}

// Finish implement the DownloadProgress interface
func (d DisableProgress) Finish() {}

// ProgressBar implement the DownloadProgress interface with download progress
type ProgressBar struct {
	bar  *pb.ProgressBar
	size int64
}

// Start implement the DownloadProgress interface
func (p *ProgressBar) Start(url string, size int64) {
	p.size = size
	p.bar = pb.Start64(size)
	p.bar.Set(pb.Bytes, true)

	// Check if stdout is a TTY
	isTTY := term.IsTerminal(int(os.Stdout.Fd()))

	// Use a simple template without ANSI escape sequences when stdout is not a TTY
	if isTTY {
		p.bar.SetTemplateString(fmt.Sprintf(`download %s {{counters . }} {{percent . }} {{speed . "%%s/s" "? MiB/s"}}`, url))
	} else {
		// Simple template for non-TTY output (no progress bar, just text)
		p.bar.SetTemplateString(fmt.Sprintf(`download %s {{counters . }}`, url))
	}
}

// SetCurrent implement the DownloadProgress interface
func (p *ProgressBar) SetCurrent(size int64) {
	p.bar.SetCurrent(size)
}

// Finish implement the DownloadProgress interface
func (p *ProgressBar) Finish() {
	p.bar.Finish()
}
