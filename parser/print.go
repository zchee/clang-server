// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/go-clang/v3.9/clang"
	"github.com/zchee/clang-server/internal/log"
	"github.com/zchee/clang-server/symbol"
)

// printDiagnostics prints a diagnostics information.
func printDiagnostics(diags []clang.Diagnostic) {
	for _, d := range diags {
		file, line, col, offset := d.Location().FileLocation()
		log.Debugf("Location:", file.Name(), line, col, offset)
		log.Debugf("PROBLEM:", d.Spelling())
	}
}

// printFile prints a symbol.File data information.
func printFile(out *symbol.File) {
	log.Debugf("out.Name(): %+v\n", out.Name())

	for i, sym := range out.Symbols() {
		log.Debugf("sym.ID: %+v\n", sym.ID())
		def := sym.Def()
		log.Debugf("sym.Def(): FileName: %s, Line: %d, Col: %d, Offset: %d, USR: %s\n", def.FileName(), def.Line(), def.Col(), def.Offset(), def.USR())
		for _, decl := range sym.Decls() {
			log.Debugf("sym.Decls(): FileName: %s, Line: %d, Col: %d, Offset: %d, USR: %s\n", decl.FileName(), decl.Line(), decl.Col(), decl.Offset(), decl.USR())
		}
		for _, caller := range sym.Callers() {
			loc := caller.Location()
			log.Debugf("sym.Decls(): FileName: %s, Line: %d, Col: %d, Offset: %d, USR: %s\n", loc.FileName(), loc.Line(), loc.Col(), loc.Offset(), loc.USR())
			log.Debugf("caller.FuncCall: %+v", caller.FuncCall())
		}
		// for _, hdr := range file.Header() {
		// 	log.Printf("hdr: FileID: %s, Mtime: %d", hdr.FileID().String(), hdr.Mtime())
		// }
		if i > 10 {
			break
		}
	}
}
