package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Data for IndexerCallbacks#importedASTFile.
type IdxImportedASTFileInfo struct {
	c C.CXIdxImportedASTFileInfo
}

// Top level AST file containing the imported PCH, module or submodule.
func (iiastfi IdxImportedASTFileInfo) File() File {
	return File{iiastfi.c.file}
}

// The imported module or NULL if the AST file is a PCH.
func (iiastfi IdxImportedASTFileInfo) Module() Module {
	return Module{iiastfi.c.module}
}

// Location where the file is imported. Applicable only for modules.
func (iiastfi IdxImportedASTFileInfo) Loc() IdxLoc {
	return IdxLoc{iiastfi.c.loc}
}

// Non-zero if an inclusion directive was automatically turned into a module import. Applicable only for modules.
func (iiastfi IdxImportedASTFileInfo) IsImplicit() bool {
	o := iiastfi.c.isImplicit

	return o != C.int(0)
}
