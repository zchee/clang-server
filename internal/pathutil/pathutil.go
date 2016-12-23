// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pathutil

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/zchee/clang-server/internal/hashutil"
)

const dirname = "clang-server"

// IsExist returns whether the filename is exists.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err) || err == nil
}

// IsNotExist returns whether the filename is exists.
func IsNotExist(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

// IsDir returns whether the filename is directory.
func IsDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

// xdgCacheHome return the XDG Base Directory Specification's cache directory path.
func xdgCacheHome() string {
	xdghome := os.Getenv("XDG_CACHE_HOME")
	if xdghome == "" {
		xdghome = filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return xdghome
}

// CacheDir return the clang-server cache directory path.
func CacheDir() string {
	dir := filepath.Join(xdgCacheHome(), dirname)
	if IsNotExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func ProjectCacheDir(root string) string {
	// assume the absolude directory path
	id := hashutil.NewHashString(root)
	dir := filepath.Join(CacheDir(), filepath.Base(root)+"."+hashutil.EncodeToString(id))
	if IsNotExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func ProjectASTDir(root string) string {
	cacheDir := ProjectCacheDir(root)
	astDir := filepath.Join(cacheDir, "ast")
	if IsNotExist(astDir) {
		if err := os.MkdirAll(astDir, 0755); err != nil {
			log.Fatal(err)
		}
	}
	return astDir
}

var rootMarks = []string{
	".git",                      // git
	"LICENSE",                   // project license
	".gitignore",                // git
	".travis.yml", "circle.yml", // CI service config files
	"CMakeLists.txt",                                                   // CMake
	"autogen.sh", "configure", "Makefile.am", "Makefile.in", "INSTALL", // GNU Autotools
	".hg", ".svn", ".bzr", "_darcs", ".tup", // not git vcs repository
}

// FindProjectRoot finds the project root directory path from path.
func FindProjectRoot(path string) (string, error) {
	if path == "" {
		return "", errors.New("project root is blank")
	}

	if !filepath.IsAbs(path) {
		abs, err := filepath.Abs(path)
		if err == nil {
			path = abs
		}
	}

	first := path
	for path != "/" {
		for _, c := range rootMarks {
			if p := filepath.Join(path, c); IsExist(p) {
				return path, nil
			}
		}
		path = filepath.Dir(path)
	}

	return "", fmt.Errorf("couldn't find project root in %s", first)
}
