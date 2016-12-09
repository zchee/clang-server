// go-clang-dump shows how to dump the AST of a C/C++ file via the Cursor
// visitor API.
//
// ex:
// $ go-clang-dump -fname=foo.cxx
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-clang/v3.9/clang"
)

var fname = flag.String("fname", "", "the file to analyze")

func main() {
	os.Exit(cmd(os.Args[1:]))
}

func cmd(args []string) int {
	fmt.Printf(":: go-clang-dump...\n")
	if err := flag.CommandLine.Parse(args); err != nil {
		fmt.Printf("ERROR: %s", err)

		return 1
	}

	fmt.Printf(":: fname: %s\n", *fname)
	fmt.Printf(":: args: %v\n", flag.Args())

	if *fname == "" {
		flag.Usage()
		fmt.Printf("please provide a file name to analyze\n")

		return 1
	}

	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	tuArgs := []string{}
	if len(flag.Args()) > 0 && flag.Args()[0] == "-" {
		tuArgs = make([]string, len(flag.Args()[1:]))
		copy(tuArgs, flag.Args()[1:])
	}

	tu := idx.ParseTranslationUnit(*fname, tuArgs, nil, 0)
	defer tu.Dispose()

	fmt.Printf("tu: %s\n", tu.Spelling())

	diagnostics := tu.Diagnostics()
	for _, d := range diagnostics {
		fmt.Println("PROBLEM:", d.Spelling())
	}

	cursor := tu.TranslationUnitCursor()
	fmt.Printf("cursor-isnull: %v\n", cursor.IsNull())
	fmt.Printf("cursor: %s\n", cursor.Spelling())
	fmt.Printf("cursor-kind: %s\n", cursor.Kind().Spelling())

	fmt.Printf("tu-fname: %s\n", tu.File(*fname).Name())

	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			fmt.Printf("cursor: <none>\n")

			return clang.ChildVisit_Continue
		}

		fmt.Printf("%s: %s (%s)\n", cursor.Kind().Spelling(), cursor.Spelling(), cursor.USR())

		switch cursor.Kind() {
		case clang.Cursor_ClassDecl, clang.Cursor_EnumDecl, clang.Cursor_StructDecl, clang.Cursor_Namespace:
			return clang.ChildVisit_Recurse
		}

		return clang.ChildVisit_Continue
	})

	if len(diagnostics) > 0 {
		fmt.Println("NOTE: There were problems while analyzing the given file")
	}

	fmt.Printf(":: bye.\n")

	return 0
}
