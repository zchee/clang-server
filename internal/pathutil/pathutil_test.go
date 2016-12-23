// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pathutil

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	testCwd, _     = os.Getwd()
	projectRoot, _ = filepath.Abs(filepath.Join(testCwd, "../../../"))
	testdata       = filepath.Join(projectRoot, "src", "nvim-go", "testdata")
	testGoPath     = filepath.Join(testdata, "go")
	testGbPath     = filepath.Join(testdata, "gb")

	astdump     = filepath.Join(testGoPath, "src", "astdump")
	astdumpMain = filepath.Join(astdump, "astdump.go")
)

func TestIsExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exist (own file)",
			args: args{filename: "./pathutil_test.go"},
			want: true,
		},
		{
			name: "not exist",
			args: args{filename: "./not_exist.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsExist(tt.args.filename); got != tt.want {
				t.Errorf("IsExist(%v) = %v, want %v", tt.args.filename, got, tt.want)
			}
		})
	}
}

func TestIsNotExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exist (own file)",
			args: args{filename: "./pathutil_test.go"},
			want: false,
		},
		{
			name: "not exist",
			args: args{filename: "./not_exist.go"},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsNotExist(tt.args.filename); got != tt.want {
				t.Errorf("IsExist(%v) = %v, want %v", tt.args.filename, got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true (own parent directory)",
			args: args{filename: testCwd},
			want: true,
		},
		{
			name: "false (own file path)",
			args: args{filename: filepath.Join(testCwd, "pathutil_test.go")},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.args.filename); got != tt.want {
				t.Errorf("IsDir(%v) = %v, want %v", tt.args.filename, got, tt.want)
			}
		})
	}
}
