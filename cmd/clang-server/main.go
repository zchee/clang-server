// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zchee/clang-server/parser"
)

var path = flag.String("path", "", "parse project root directory.")

func main() {
	flag.Parse()

	fmt.Printf("clang version: %s\n", parser.ClangVersion())

	config := parser.Config{}
	if *path != "" {
		p := parser.NewParser(*path, config)
		if err := p.Parse(); err != nil {
			log.Fatal(err)
		}
	}
}
