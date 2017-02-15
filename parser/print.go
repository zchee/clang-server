// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"log"

	"github.com/go-clang/v3.9/clang"
	"github.com/zchee/clang-server/symbol"
)

// printDiagnostics prints a diagnostics information.
func printDiagnostics(diags []clang.Diagnostic) {
	for _, d := range diags {
		file, line, col, offset := d.Location().FileLocation()
		fmt.Println("Location:", file.Name(), line, col, offset)
		fmt.Println("PROBLEM:", d.Spelling())
	}
}

// printFile prints a symbol.File data information.
func printFile(file *symbol.File) {
	log.Printf("out.Name(): %+v\n", file.Name())

	for i, sym := range file.Symbols() {
		log.Printf("sym.ID: %+v\n", sym.ID())
		def := sym.Def()
		log.Printf("sym.Def(): FileName: %s, Line: %d, Col: %d, Offset: %d, USR: %s\n", def.FileName(), def.Line(), def.Col(), def.Offset(), def.USR())
		for _, decl := range sym.Decls() {
			log.Printf("sym.Decls(): FileName: %s, Line: %d, Col: %d, Offset: %d, USR: %s\n", decl.FileName(), decl.Line(), decl.Col(), decl.Offset(), decl.USR())
		}
		for _, caller := range sym.Callers() {
			loc := caller.Location()
			log.Printf("sym.Decls(): FileName: %s, Line: %d, Col: %d, Offset: %d, USR: %s\n", loc.FileName(), loc.Line(), loc.Col(), loc.Offset(), loc.USR())
			log.Printf("caller.FuncCall: %+v", caller.FuncCall())
		}
		// for _, hdr := range file.Header() {
		// 	log.Printf("hdr: FileID: %s, Mtime: %d", hdr.FileID().String(), hdr.Mtime())
		// }
		if i > 10 {
			break
		}
	}
}
