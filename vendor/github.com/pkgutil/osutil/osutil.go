// Copyright 2017 The pkgutil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osutil

import (
	"os"
)

// IsExist reports whether the filename is exists.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// IsNotExist reports whether the filename is not exists.
func IsNotExist(filename string) bool {
	_, err := os.Stat(filename)
	return err != nil && os.IsNotExist(err)
}

// IsDirExist reports whether dir exists and which is directory.
func IsDirExist(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.IsDir()
}

// MkdirAll checks whether the exist dir directory and create directory to dir filepath if not exist.
func MkdirAll(dir string, perm os.FileMode) error {
	if !IsDirExist(dir) {
		if err := os.MkdirAll(dir, perm); err != nil {
			return err
		}
	}
	return nil
}
