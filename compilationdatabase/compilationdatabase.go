// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compilationdatabase

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-clang/v3.9/clang"
	"github.com/pkg/errors"
)

const DefaultJsonName = "compile_commands.json"

type CompilationDatabase struct {
	projectRoot string

	cd    clang.CompilationDatabase
	found bool

	flags  map[string][]string
	flagMu sync.Mutex
}

func NewCompilationDatabase(root string) *CompilationDatabase {
	return &CompilationDatabase{
		projectRoot: root,
		flags:       make(map[string][]string),
	}
}

func (c *CompilationDatabase) findFile(filename string, pathRange []string) string {
	if pathRange == nil {
		parent := filepath.Dir(c.projectRoot)
		buildDir := filepath.Join(c.projectRoot, "build")
		pathRange = []string{c.projectRoot, parent, buildDir}
	}

	pathCh := make(chan string, len(pathRange))
	for _, d := range pathRange {
		go func(d string) {
			// log.Printf("d: %+v\n", d)
			_, err := os.Stat(filepath.Join(d, filename))
			if !os.IsNotExist(err) {
				// log.Printf("found: %+v\n", file)
				pathCh <- d
			}
		}(d)
	}

	return <-pathCh
}

func (c *CompilationDatabase) Parse(filename string, pathRange ...string) error {
	if filename == "" {
		filename = DefaultJsonName
	}

	dir := c.findFile(filename, pathRange)
	if dir == "" {
		return errors.Errorf("couldn't find the %s file", filename)
	}
	c.found = true

	err, cd := clang.FromDirectory(dir)
	if err != clang.CompilationDatabase_NoError {
		return errors.WithStack(err)
	}
	c.cd = cd

	if err := c.parseAllFlags(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *CompilationDatabase) parseAllFlags() error {
	c.flagMu.Lock()
	defer c.flagMu.Unlock()

	cmds := c.cd.AllCompileCommands()
	ncmds := cmds.Size()

	for i := uint32(0); i < ncmds; i++ {
		cmd := cmds.Command(i)
		args, err := c.parseFlags(cmd)
		if err != nil {
			return errors.WithStack(err)
		}
		c.flags[cmd.Filename()] = args
	}
	return nil
}

func (c *CompilationDatabase) Flags(filename string) ([]string, error) {
	c.flagMu.Lock()
	defer c.flagMu.Unlock()

	if c.flags[filename] != nil {
		return c.flags[filename], nil
	}

	cmds := c.cd.CompileCommands(filename)
	flags, err := c.parseFlags(cmds.Command(0))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c.flags[filename] = flags

	return flags, nil
}

func (c *CompilationDatabase) parseFlags(cmd clang.CompileCommand) ([]string, error) {
	n := cmd.NumArgs()
	flags := make([]string, 0, n)

	for i := uint32(0); i < n; i++ {
		f := cmd.Arg(i)
		switch {
		case f == "-D":
			flags = append(flags, f, cmd.Arg(i+1))
		case strings.HasPrefix(f, "-D"):
			flags = append(flags, f)
		case f == "-I":
			includeDir, err := c.absDir(cmd.Arg(i+1), cmd.Directory())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			flags = append(flags, "-I", includeDir)
		case strings.HasPrefix(f, "-I"):
			includeDir, err := c.absDir(strings.Replace(f, "-I", "", 1), cmd.Directory())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			flags = append(flags, "-I", includeDir)
		}
	}

	return flags, nil
}

func (c *CompilationDatabase) absDir(includePath, buildDir string) (string, error) {
	if filepath.IsAbs(includePath) {
		return includePath, nil
	}

	abs, err := filepath.Abs(includePath)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get absolute path: %v", err)
	}

	return filepath.Clean(abs), nil
}

func (c *CompilationDatabase) Dispose() {
	c.cd.Dispose()
}
