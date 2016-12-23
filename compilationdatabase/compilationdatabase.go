// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compilationdatabase

import (
	"path/filepath"
	"strings"

	"github.com/go-clang/v3.9/clang"
	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/zchee/clang-server/internal/pathutil"
	"github.com/zchee/clang-server/log"
)

// DefaultJSONName default of compile_commands.json filename.
const DefaultJSONName = "compile_commands.json"

// ErrNotFound error of not found the compile_commands.json file.
var ErrNotFound = errors.New("couldn't find the compile_commands.json")

// CompilationDatabase represents a consist of an array of “command objects”,
// where each command object specifies one way a translation unit is compiled in the project.
type CompilationDatabase struct {
	projectRoot string

	cd clang.CompilationDatabase
	cc []*CompileCommand
}

// CompileCommand represents a each command object contains the translation unit’s main file,
// the working directory of the compile run and the actual compile command.
type CompileCommand struct {
	// Directory the working directory of the compilation.
	Directory string `json:"directory"`
	// File the main translation unit source processed by this compilation step.
	File string `json:"file"`
	// Command the compile command executed.
	Command string `json:"command"`
	// Arguments the compile command executed as list of strings.
	Arguments []string `json:"arguments"`
	// Output the name of the output created by this compilation step.
	Output string `json:"output"`
}

// NewCompilationDatabase return the new CompilationDatabase.
func NewCompilationDatabase(root string) *CompilationDatabase {
	return &CompilationDatabase{
		projectRoot: root,
	}
}

// findFile finds the filename on pathRange recursively.
func (c *CompilationDatabase) findFile(filename string, pathRange []string) string {
	if pathRange == nil {
		parent := filepath.Dir(c.projectRoot)             // parent of projectRoot
		buildDir := filepath.Join(c.projectRoot, "build") // projectRoot/build
		outDir := filepath.Join(c.projectRoot, "out")     // projectRoot/out
		pathRange = []string{c.projectRoot, parent, buildDir, outDir}
	}

	pathCh := make(chan string, len(pathRange))
	for _, d := range pathRange {
		go func(d string) {
			if pathutil.IsExist(filepath.Join(d, filename)) {
				log.Debug("found", zap.String("filepath", filepath.Join(d, filename)))
				pathCh <- d
			}
		}(d)
	}

	return <-pathCh
}

// Parse parses the project root directory recursively, and cache the compile
// flags to flags map.
func (c *CompilationDatabase) Parse(jsonfile string, pathRange ...string) error {
	if jsonfile == "" {
		jsonfile = DefaultJSONName
	}

	dir := c.findFile(jsonfile, pathRange)
	if dir == "" {
		return ErrNotFound
	}

	cErr, cd := clang.FromDirectory(dir)
	if cErr != clang.CompilationDatabase_NoError {
		return errors.WithStack(cErr)
	}

	c.cd = cd
	defer c.cd.Dispose()

	if err := c.parseFlags(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// parseFlags parses the all of project files compile flag.
func (c *CompilationDatabase) parseFlags() error {
	cmds := c.cd.AllCompileCommands()
	ncmds := cmds.Size()
	c.cc = make([]*CompileCommand, 0, ncmds)

	for i := uint32(0); i < ncmds; i++ {
		cmd := cmds.Command(i)
		args := c.formatFlag(cmd)
		c.cc = append(c.cc, &CompileCommand{
			Directory: cmd.Directory(),
			File:      cmd.Filename(),
			Command:   strings.Join(args, " "),
			Arguments: args,
		})
	}

	return nil
}

// formatFlag formats the compile flag for the libclang TranslationUnit arg.
func (c *CompilationDatabase) formatFlag(cmd clang.CompileCommand) []string {
	n := cmd.NumArgs()
	flags := make([]string, 0, n)

	for i := uint32(0); i < n; i++ {
		f := cmd.Arg(i)
		dir := cmd.Directory()
		switch {
		case
			f == "-D",         // <macroname>=<value>: Adds an implicit #define into the predefines buffer which is read before the source file is preprocessed
			f == "-U",         // <macroname>: Adds an implicit #undef into the predefines buffer which is read before the source file is preprocessed
			f == "-isysroot",  // <dir>: Add directory to SYSTEM include search path
			f == "-framework": // <name>: Tells the linker to search for `name.framework/name' the framework search path

			flags = append(flags, f, cmd.Arg(i+1))

		case
			strings.HasPrefix(f, "-D"),                     // <macroname>=<value>: Adds an implicit #define into the predefines buffer which is read before the source file is preprocessed
			strings.HasPrefix(f, "-U"),                     // <macroname>: Adds an implicit #undef into the predefines buffer which is read before the source file is preprocessed
			strings.HasPrefix(f, "-std"),                   // <language>: Specify the language standard
			strings.HasPrefix(f, "-stdlib"),                // <library>: Specify the C++ standard library to use
			strings.HasPrefix(f, "-mmacosx-version-min"),   // <version>: Specify the minimum version supported by your application when building for macOS
			strings.HasPrefix(f, "-miphoneos-version-min"): // <version>: Specify the minimum version supported by your application when building for iPhone OS

			flags = append(flags, f)

		case
			f == "-I",                 // <value>: Specified directory to the search path for include files
			f == "-F",                 // <directory>: Specified directory to the search path for framework include files
			f == "-idirafter",         // <value>: Add directory to AFTER include search path
			f == "-iframework",        // <value>: Add directory to SYSTEM framework search path
			f == "-imacros",           // <file>: Include macros from file before parsing
			f == "-include-pch",       // <file>: Include precompiled header file
			f == "-include",           // <file>: Adds an implicit #include into the predefines buffer which is read before the source file is preprocessed
			f == "-isystem-after",     // <directory>: Add directory to end of the SYSTEM include search path
			f == "-isystem",           // <directory>: Set the system root directory (usually /)
			f == "-ivfsoverlay",       // <value>: Overlay the virtual filesystem described by file over the real file system
			f == "-iwithprefixbefore", // <dir>: Set directory to include search path with prefix
			f == "-iwithprefix",       // <dir>: Set directory to SYSTEM include search path with prefix
			f == "-iwithsysroot":      // <directory>: Add directory to SYSTEM include search path, absolute paths are relative to -isysroot

			includeDir := c.absPath(cmd.Arg(i+1), dir)
			flags = append(flags, f, includeDir)

		case strings.HasPrefix(f, "-I"): // <value>: Specified directory to the search path for include files
			includeDir := c.absPath(strings.TrimPrefix(f, "-I"), dir)
			flags = append(flags, "-I", includeDir)

		case strings.HasPrefix(f, "-F"): // <directory>: Specified directory to the search path for framework include files
			includeDir := c.absPath(strings.TrimPrefix(f, "-F"), dir)
			flags = append(flags, "-F", includeDir)
		}
	}

	return flags
}

// absPath return the absolube directory path based by buildDir.
func (c *CompilationDatabase) absPath(includePath, buildDir string) string {
	if filepath.IsAbs(includePath) {
		return includePath
	}

	for _, d := range []string{buildDir, c.projectRoot} {
		if dir := filepath.Join(d, includePath); pathutil.IsExist(dir) {
			return dir
		}
	}

	return includePath
}

// CompileCommands return the CompileCommand struct based parse result.
func (c *CompilationDatabase) CompileCommands() []*CompileCommand {
	return c.cc
}
