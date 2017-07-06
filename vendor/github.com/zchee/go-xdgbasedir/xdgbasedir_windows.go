// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package xdgbasedir

import (
	"path/filepath"

	"github.com/zchee/go-xdgbasedir/home"
)

var (
	defaultPath       = filepath.Join(home.Dir(), "AppData", "Local")
	defaultDataHome   = defaultPath
	defaultConfigHome = defaultPath
	defaultDataDirs   = defaultPath
	defaultConfigDirs = defaultPath
	defaultCacheHome  = filepath.Join(defaultPath, "cache")
	defaultRuntimeDir = home.Dir()
)

func dataHome() string {
	return defaultDataHome
}

func configHome() string {
	return defaultConfigHome
}

func dataDirs() string {
	return defaultDataDirs
}

func configDirs() string {
	return defaultConfigDirs
}

func cacheHome() string {
	return defaultCacheHome
}

func runtimeDir() string {
	return defaultRuntimeDir
}
