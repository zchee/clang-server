// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashutil

import (
	"reflect"
	"testing"
)

func TestNewHash(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want [Size]byte
	}{
		{
			name: "normal",
			args: args{b: []byte("test")},
			want: [Size]byte{
				167, 16, 121, 212, 40, 83, 222, 162, 110, 69, 48, 4, 51, 134, 112, 165,
				56, 20, 183, 129, 55, 255, 190, 208, 118, 3, 164, 29, 118, 164, 131, 170,
				155, 195, 59, 88, 47, 119, 211, 10, 101, 230, 242, 154, 137, 108, 4, 17,
				243, 131, 18, 225, 214, 110, 11, 241, 99, 134, 200, 106, 137, 190, 165, 114,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHash(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHash(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

func TestNewHashString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [Size]byte
	}{
		{
			name: "normal",
			args: args{s: "test"},
			want: [Size]byte{
				167, 16, 121, 212, 40, 83, 222, 162, 110, 69, 48, 4, 51, 134, 112, 165,
				56, 20, 183, 129, 55, 255, 190, 208, 118, 3, 164, 29, 118, 164, 131, 170,
				155, 195, 59, 88, 47, 119, 211, 10, 101, 230, 242, 154, 137, 108, 4, 17,
				243, 131, 18, 225, 214, 110, 11, 241, 99, 134, 200, 106, 137, 190, 165, 114,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashString(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestEncodeToString(t *testing.T) {
	type args struct {
		b [Size]byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{b: [Size]byte{
				167, 16, 121, 212, 40, 83, 222, 162,
				110, 69, 48, 4, 51, 134, 112, 165,
				56, 20, 183, 129, 55, 255, 190, 208,
				118, 3, 164, 29, 118, 164, 131, 170,
				155, 195, 59, 88, 47, 119, 211, 10,
				101, 230, 242, 154, 137, 108, 4, 17,
				243, 131, 18, 225, 214, 110, 11, 241,
				99, 134, 200, 106, 137, 190, 165, 114,
			}},
			want: "a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeToString(tt.args.b); got != tt.want {
				t.Errorf("EncodeToString(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

func Test_byteSliceToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{b: []byte("a")},
			want: "a",
		},
		{
			name: "a/b",
			args: args{b: []byte("a/b")},
			want: "a/b",
		},
		{
			name: "a/b/c",
			args: args{b: []byte("a/b/c")},
			want: "a/b/c",
		},
		{
			name: "hash byteslice",
			args: args{b: []byte("a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572")},
			want: "a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteSliceToString(tt.args.b); got != tt.want {
				t.Errorf("byteSliceToString(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

func Test_stringToByteSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "a",
			args: args{s: "a"},
			want: []byte("a"),
		},
		{
			name: "a/b",
			args: args{s: "a/b"},
			want: []byte("a/b"),
		},
		{
			name: "a/b/c",
			args: args{s: "a/b/c"},
			want: []byte("a/b/c"),
		},
		{
			name: "hash string",
			args: args{s: "a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572"},
			want: []byte("a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToByteSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToByteSlice(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
