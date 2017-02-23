// Copyright 2017 The pkgutil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepathutil

import (
	"os"
)

// IsExist returns whether the filename is exists.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil && !os.IsNotExist(err)
}

// IsNotExist returns whether the filename is not exists.
func IsNotExist(filename string) bool {
	_, err := os.Stat(filename)
	return err != nil && os.IsNotExist(err)
}

// IsDir returns whether the filename is directory.
func IsDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

// MkdirAll checks whether the exist dir directory. and create directory to dir filepath if not exist.
func MkdirAll(dir string, perm os.FileMode) error {
	if IsNotExist(dir) {
		if err := os.MkdirAll(dir, perm); err != nil {
			return err
		}
	}
	return nil
}
