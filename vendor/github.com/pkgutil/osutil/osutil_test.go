// Copyright 2017 The pkgutil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osutil

import (
	"os"
	"testing"
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
			name: "Exist file",
			args: args{filename: "testdata/exist.go"},
			want: true,
		},
		{
			name: "Not exist file",
			args: args{filename: "testdata/not_exist.go"},
			want: false,
		},
		{
			name: "OWn file",
			args: args{filename: "osutil_test.go"},
			want: true,
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
			name: "Exist file",
			args: args{filename: "testdata/exist.go"},
			want: false,
		},
		{
			name: "Not exist file",
			args: args{filename: "testdata/not_exist.go"},
			want: true,
		},
		{
			name: "Own file",
			args: args{filename: "osutil_test.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsNotExist(tt.args.filename); got != tt.want {
				t.Errorf("IsNotExist(%v) = %v, want %v", tt.args.filename, got, tt.want)
			}
		})
	}
}

func TestIsDirExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testdata dir",
			args: args{filename: "testdata"},
			want: true,
		},
		{
			name: "Own file",
			args: args{filename: "osutil_test.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsDirExist(tt.args.filename); got != tt.want {
				t.Errorf("IsDir(%v) = %v, want %v", tt.args.filename, got, tt.want)
			}
		})
	}
}

func TestMkdirAll(t *testing.T) {
	type args struct {
		dir  string
		perm os.FileMode
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		deferRemove bool
	}{
		{
			name: "Exist testdata directory",
			args: args{
				dir:  "testdata",
				perm: 0700,
			},
			wantErr:     false,
			deferRemove: false,
		},
		{
			name: "Not exist testdata/test directory",
			args: args{
				dir:  "testdata/test",
				perm: 0700,
			},
			wantErr:     false,
			deferRemove: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := MkdirAll(tt.args.dir, tt.args.perm); (err != nil) != tt.wantErr {
				t.Errorf("CreateDir(%v, %v) error = %v, wantErr %v", tt.args.dir, tt.args.perm, err, tt.wantErr)
			}
			if fi, err := os.Stat(tt.args.dir); err != nil || !fi.IsDir() {
				t.Errorf("CreateDir(%v, %v) error = %v, wantErr %v", tt.args.dir, tt.args.perm, err, tt.wantErr)
			}
			if tt.deferRemove {
				defer os.RemoveAll(tt.args.dir)
			}
		})
	}
}
