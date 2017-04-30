// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"time"

	"github.com/zchee/clang-server/internal/log"
)

func profile(start time.Time, name string) {
	elapsed := time.Since(start).Seconds()
	log.Debugf("%s: %fsec\n", name, elapsed)
}
