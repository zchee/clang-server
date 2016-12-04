// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compilationdatabase

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/go-clang/v3.9/clang"
)

func TestNewCompilationDatabase(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want *CompilationDatabase
	}{
		{
			name: "new",
			args: args{root: "./testdata"},
			want: &CompilationDatabase{
				projectRoot: "./testdata",
				found:       false,
				flags:       make(map[string][]string),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompilationDatabase(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompilationDatabase(%v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}

func TestCompilationDatabase_findFile(t *testing.T) {
	type fields struct {
		projectRoot string
		cd          clang.CompilationDatabase
		found       bool
		flags       map[string][]string
	}
	type args struct {
		filename  string
		pathRange []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "build directory",
			fields: fields{
				projectRoot: filepath.Join("./testdata", "builddir"),
			},
			args: args{
				filename:  "",
				pathRange: nil,
			},
			want: filepath.Join("./testdata", "builddir", "build"),
		},
		{
			name: "root directory",
			fields: fields{
				projectRoot: filepath.Join("./testdata", "root"),
			},
			args: args{
				filename:  "",
				pathRange: nil,
			},
			want: filepath.Join("./testdata", "root"),
		},
		{
			name: "parent directory",
			fields: fields{
				projectRoot: filepath.Join("./testdata", "parent", "child"),
			},
			args: args{
				filename:  "",
				pathRange: nil,
			},
			want: filepath.Join("./testdata", "parent"),
		},
		{
			name: "specified filename with build directory",
			fields: fields{
				projectRoot: filepath.Join("./testdata", "specified_filename"),
			},
			args: args{
				filename:  "specified_filename.json",
				pathRange: nil,
			},
			want: filepath.Join("./testdata", "specified_filename", "build"),
		},
		{
			name: "specified pathRange",
			fields: fields{
				projectRoot: filepath.Join("./testdata", "pathRange"),
			},
			args: args{
				filename:  "",
				pathRange: []string{filepath.Join("./testdata", "pathRange", "json")},
			},
			want: filepath.Join("./testdata", "pathRange", "json"),
		},
	}
	for _, tt := range tests {
		c := &CompilationDatabase{
			projectRoot: tt.fields.projectRoot,
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := c.findFile(tt.args.filename, tt.args.pathRange); got != tt.want {
				t.Errorf("CompilationDatabase.findFile(%v, %v) = %v, want %v", tt.args.filename, tt.args.pathRange, got, tt.want)
			}
		})
	}
}
