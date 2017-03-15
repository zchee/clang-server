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

	"github.com/pkgutil/osutil"
	"github.com/zchee/clang-server/internal/hashutil"
	xdgbasedir "github.com/zchee/go-xdgbasedir"
)

const dirname = "clang-server"

// CacheDir return the clang-server cache directory path.
func CacheDir() string {
	dir := filepath.Join(xdgbasedir.CacheHome(), dirname)
	if !osutil.IsExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

// ProjectCacheDir return the project path based cache directory path.
func ProjectCacheDir(root string) string {
	// assume the absolude directory path
	id := hashutil.NewHashString(root)
	dir := filepath.Join(CacheDir(), filepath.Base(root)+"."+hashutil.EncodeToString(id))
	return dir
}

// ProjectASTDir return the project path based AST cache directory path.
func ProjectASTDir(root string) string {
	cacheDir := ProjectCacheDir(root)
	astDir := filepath.Join(cacheDir, "ast")
	if osutil.IsNotExist(astDir) {
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
	".hg", ".svn", ".bzr", "_darcs", ".tup", // typical vcs directories
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
			if p := filepath.Join(path, c); osutil.IsExist(p) {
				return path, nil
			}
		}
		path = filepath.Dir(path)
	}

	return "", fmt.Errorf("couldn't find project root in %s", first)
}
