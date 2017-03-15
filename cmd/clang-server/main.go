// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zchee/clang-server/parser"
)

var path = flag.String("path", "", "parse project root directory.")

func main() {
	flag.Parse()
	if *path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("clang version: %s\n", parser.ClangVersion())

	config := parser.Config{}
	p := parser.NewParser(*path, config)
	p.Parse()
}
