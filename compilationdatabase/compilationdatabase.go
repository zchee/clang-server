// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compilationdatabase

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-clang/v3.9/clang"
	"github.com/pkg/errors"
	"github.com/pkgutil/osutil"
)

// DefaultJSONName default of compile_commands.json filename.
const DefaultJSONName = "compile_commands.json"

// ErrNotFound error of not found the compile_commands.json file.
var ErrNotFound = errors.New("couldn't find the compile_commands.json")

// CompilationDatabase represents a consist of an array of “command objects”,
// where each command object specifies one way a translation unit is compiled in the project.
type CompilationDatabase struct {
	root string

	cd             clang.CompilationDatabase
	cmds           []*CompileCommand
	CompilerConfig *CompilerConfig

	mu sync.Mutex
}

type CompilerConfig struct {
	Target              string
	ThreadModel         string
	InstalledDir        string
	DefaultFlag         []string
	Version             string
	SystemCIncludeDir   []string
	SystemCXXIncludeDir []string
	SystemFrameworkDir  []string
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
		root: root,
	}
}

// Parse parses the project root directory recursively, and cache the compile
// flags to flags map.
func (c *CompilationDatabase) Parse(jsonfile string, pathRange []string) error {
	ch := make(chan *CompilerConfig, 1)
	go func() { ch <- c.DefaultCompilerConfig() }()

	if jsonfile == "" {
		jsonfile = DefaultJSONName
	}
	dir := c.findJSONFile(jsonfile, pathRange)
	if dir == "" {
		return ErrNotFound
	}

	cErr, cd := clang.FromDirectory(dir)
	if cErr != clang.CompilationDatabase_NoError {
		return errors.WithStack(clang.CompilationDatabase_Error(cErr))
	}
	c.cd = cd
	defer c.cd.Dispose()

	c.CompilerConfig = <-ch

	if err := c.parseFlags(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// CompileCommands return the CompileCommand struct based parse result.
func (c *CompilationDatabase) CompileCommands() []*CompileCommand {
	return c.cmds
}

// DefaultCompilerConfig gets the compiler defalut configs.
// The parse target sample:
//  clang -v -x c++ /dev/null -fsyntax-only
//
//  Apple LLVM version 8.1.0 (clang-802.0.27.2)
//  Target: x86_64-apple-darwin16.5.0
//  Thread model: posix
//  InstalledDir: /Applications/Xcode-beta.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin
//   "/Applications/Xcode-beta.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang" -cc1 -triple x86_64-apple-macosx10.12.0 ...
//  clang -cc1 version 8.1.0 (clang-802.0.27.2) default target x86_64-apple-darwin16.5.0
//  ignoring nonexistent directory "/usr/include/c++/v1"
//  #include "..." search starts here:
//  #include <...> search starts here:
//   /Applications/Xcode-beta.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/../include/c++/v1
//   /usr/local/include
//   /Applications/Xcode-beta.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/../lib/clang/8.1.0/include
//   /Applications/Xcode-beta.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include
//   /usr/include
//   /System/Library/Frameworks (framework directory)
//   /Library/Frameworks (framework directory)
//  End of search list.
func (c *CompilationDatabase) DefaultCompilerConfig() *CompilerConfig {
	cc := "clang" // default is clang
	if envCC := os.Getenv("CC"); envCC != "" {
		cc = envCC
	}
	ccPath, err := exec.LookPath(cc)
	if err != nil {
		log.Fatalf("couldn't find %s", cc)
	}

	// Get C++ include directory together with -x flag
	cmd := exec.Command(ccPath, "-v", "-x", "c++", "/dev/null", "-fsyntax-only")

	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cfg := new(CompilerConfig)
	var includeSection bool
	scan := bufio.NewScanner(&b)
	for scan.Scan() {
		s := scan.Text()
		if includeSection {
			if strings.HasPrefix(s, "End") {
				includeSection = false
				break
			}
			path := strings.TrimSpace(s)
			l := len(path)
			switch {
			case path[l-6:l-3] == "c++":
				cfg.SystemCXXIncludeDir = append(cfg.SystemCXXIncludeDir, "-I"+path)
			case strings.Contains(s, "framework directory"):
				cfg.SystemFrameworkDir = append(cfg.SystemFrameworkDir, "-F"+strings.TrimSuffix(path, " (framework directory)"))
			default:
				cfg.SystemCIncludeDir = append(cfg.SystemCIncludeDir, "-I"+path)
			}
		}
		switch {
		case strings.HasPrefix(s, "Target"):
			cfg.Target = strings.TrimPrefix(s, "Target: ")
		case strings.HasPrefix(s, "Thread model"):
			cfg.ThreadModel = strings.TrimPrefix(s, "Thread model: ")
		case strings.HasPrefix(s, "InstalledDir"):
			cfg.InstalledDir = strings.TrimPrefix(s, "InstalledDir: ")
		case (strings.Contains(s, cfg.InstalledDir) && strings.Contains(s, "-cc1")):
			flags := strings.Fields(s)
			cfg.DefaultFlag = flags[1 : len(flags)-1]
		case strings.HasPrefix(s, cc):
			sep := strings.Split(s, "version ")
			version := strings.Split(sep[1], " ")[0]
			cfg.Version = version
		case strings.HasPrefix(s, "#include <...>"):
			includeSection = true
		}
	}

	return cfg
}

// findFile finds the filename on pathRange recursively.
func (c *CompilationDatabase) findJSONFile(filename string, pathRange []string) string {
	if pathRange == nil {
		parent := filepath.Dir(c.root)             // parent of projectRoot
		buildDir := filepath.Join(c.root, "build") // projectRoot/build
		outDir := filepath.Join(c.root, "out")     // projectRoot/out
		pathRange = []string{c.root, parent, buildDir, outDir}
	}

	pathCh := make(chan string, 1)
	for _, d := range pathRange {
		go func(d string) {
			if osutil.IsExist(filepath.Join(d, filename)) {
				log.Printf("found filepath: %s", filepath.Join(d, filename))
				pathCh <- d
			}
		}(d)
	}

	return <-pathCh
}

// parseFlags parses the all of project files compile flag.
func (c *CompilationDatabase) parseFlags() error {
	cmds := c.cd.AllCompileCommands()
	ncmds := cmds.Size()
	c.cmds = make([]*CompileCommand, 0, ncmds)

	var wg sync.WaitGroup
	for i := uint32(0); i < ncmds; i++ {
		wg.Add(1)
		i := i

		go func(i uint32) {
			c.mu.Lock()
			defer func() {
				wg.Done()
				c.mu.Unlock()
			}()

			cmd := cmds.Command(i)
			args := c.formatFlag(cmd)
			c.cmds = append(c.cmds, &CompileCommand{
				Directory: cmd.Directory(),
				File:      cmd.Filename(),
				Command:   strings.Join(args, " "),
				Arguments: args,
			})
		}(i)
	}

	wg.Wait()

	return nil
}

// formatFlag formats the compile flag for the libclang TranslationUnit arg.
func (c *CompilationDatabase) formatFlag(cmd clang.CompileCommand) []string {
	n := cmd.NumArgs()
	flags := make([]string, 0, n)

	for i := uint32(0); i < n; i++ {
		dir := cmd.Directory()
		f := cmd.Arg(i)

		switch {
		case strings.HasPrefix(f, "-I"): // <value>: Specified directory to the search path for include files
			includeDir := c.fixArg(strings.TrimPrefix(f, "-I"), dir)
			flags = append(flags, "-I", includeDir)

		case strings.HasPrefix(f, "-F"): // <directory>: Specified directory to the search path for framework include files
			frameworkDir := c.fixArg(strings.TrimPrefix(f, "-F"), dir)
			flags = append(flags, "-F", frameworkDir)

		case f == "-I", // <value>: Specified directory to the search path for include files
			f == "-F",                 // <directory>: Specified directory to the search path for framework include files
			f == "-D",                 // <macroname>=<value>: Adds an implicit #define into the predefines buffer which is read before the source file is preprocessed
			f == "-U",                 // <macroname>: Adds an implicit #undef into the predefines buffer which is read before the source file is preprocessed
			f == "-framework",         // <name>: Tells the linker to search for `name.framework/name' the framework search path
			f == "-x",                 // <language> Treat subsequent input files as having type language.
			f == "-arch",              // <architecture> Specify the architecture to build for.
			f == "-include",           // <file>: Adds an implicit #include into the predefines buffer which is read before the source file is preprocessed
			f == "-isysroot",          // <dir>: Add directory to SYSTEM include search path
			f == "-isystem",           // <directory>: Set the system root directory (usually /)
			f == "-iframework",        // <value>: Add directory to SYSTEM framework search path
			f == "-include-pch",       // <file>: Include precompiled header file
			f == "-isystem-after",     // <directory>: Add directory to end of the SYSTEM include search path
			f == "-idirafter",         // <value>: Add directory to AFTER include search path
			f == "-imacros",           // <file>: Include macros from file before parsing
			f == "-ivfsoverlay",       // <value>: Overlay the virtual filesystem described by file over the real file system
			f == "-iwithprefix",       // <dir>: Set directory to SYSTEM include search path with prefix
			f == "-iwithprefixbefore", // <dir>: Set directory to include search path with prefix
			f == "-iwithsysroot":      // <directory>: Add directory to SYSTEM include search path, absolute paths are relative to -isysroot
			flags = append(flags, f, c.fixArg(cmd.Arg(i+1), dir))

		case strings.HasPrefix(f, "-D"), // <macroname>=<value>: Adds an implicit #define into the predefines buffer which is read before the source file is preprocessed
			strings.HasPrefix(f, "-U"),                     // <macroname>: Adds an implicit #undef into the predefines buffer which is read before the source file is preprocessed
			strings.HasPrefix(f, "-std"),                   // <language>: Specify the language standard
			strings.HasPrefix(f, "-stdlib"),                // <library>: Specify the C++ standard library to use
			strings.HasPrefix(f, "-x"),                     // <language> Treat subsequent input files as having type language.
			strings.HasPrefix(f, "-arch"),                  // <architecture> Specify the architecture to build for.
			strings.HasPrefix(f, "-mmacosx-version-min"),   // <version>: Specify the minimum version supported by your application when building for macOS
			strings.HasPrefix(f, "-miphoneos-version-min"), // <version>: Specify the minimum version supported by your application when building for iPhone OS
			strings.HasPrefix(f, "-include"),               // <file>: Adds an implicit #include into the predefines buffer which is read before the source file is preprocessed
			strings.HasPrefix(f, "-isysroot"),              // <dir>: Add directory to SYSTEM include search path
			strings.HasPrefix(f, "-isystem"),               // <directory>: Set the system root directory (usually /)
			strings.HasPrefix(f, "-iframework"),            // <value>: Add directory to SYSTEM framework search path
			strings.HasPrefix(f, "-include-pch"),           // <file>: Include precompiled header file
			strings.HasPrefix(f, "-isystem-after"),         // <directory>: Add directory to end of the SYSTEM include search path
			strings.HasPrefix(f, "-G"),
			strings.HasPrefix(f, "-T"),
			strings.HasPrefix(f, "-V"),
			strings.HasPrefix(f, "-target"),
			strings.HasPrefix(f, "-Xanalyzer"),
			strings.HasPrefix(f, "-Xassembler"),
			strings.HasPrefix(f, "-Xclang"),
			strings.HasPrefix(f, "-Xlinker"),
			strings.HasPrefix(f, "-Xpreprocessor"),
			strings.HasPrefix(f, "-b"),
			strings.HasPrefix(f, "-gcc-toolchain"),
			strings.HasPrefix(f, "-idirafter"), // <value>: Add directory to AFTER include search path
			strings.HasPrefix(f, "-imacros"),   // <file>: Include macros from file before parsing
			strings.HasPrefix(f, "-imultilib"),
			strings.HasPrefix(f, "-iprefix"),
			strings.HasPrefix(f, "-ivfsoverlay"),       // <value>: Overlay the virtual filesystem described by file over the real file system
			strings.HasPrefix(f, "-iwithprefix"),       // <dir>: Set directory to SYSTEM include search path with prefix
			strings.HasPrefix(f, "-iwithprefixbefore"), // <dir>: Set directory to include search path with prefix
			strings.HasPrefix(f, "-iwithsysroot"):      // <directory>: Add directory to SYSTEM include search path, absolute paths are relative to -isysroot

			flags = append(flags, f)
		}
	}

	return flags
}

// fixArg return the absolube directory path based by buildDir if contains filepath.Separator.
func (c *CompilationDatabase) fixArg(arg, buildDir string) string {
	if filepath.IsAbs(arg) {
		return arg
	}

	for _, d := range []string{buildDir, c.root} {
		if dir := filepath.Join(d, arg); osutil.IsExist(dir) {
			return dir
		}
	}

	return arg
}
