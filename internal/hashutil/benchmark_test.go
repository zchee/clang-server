// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashutil

import "testing"

func BenchmarkNewHash(b *testing.B) {
	data := make([]byte, 8)
	b.SetBytes(8)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewHash(data[:])
	}
}

func BenchmarkNewHashString(b *testing.B) {
	data := string("test")
	b.SetBytes(int64(len([]byte(data))))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewHashString(data)
	}
}

func BenchmarkEncodeToString(b *testing.B) {
	data := []byte{
		167, 16, 121, 212, 40, 83, 222, 162, 110, 69, 48, 4, 51, 134, 112, 165,
		56, 20, 183, 129, 55, 255, 190, 208, 118, 3, 164, 29, 118, 164, 131, 170,
		155, 195, 59, 88, 47, 119, 211, 10, 101, 230, 242, 154, 137, 108, 4, 17,
		243, 131, 18, 225, 214, 110, 11, 241, 99, 134, 200, 106, 137, 190, 165, 114,
	}
	b.SetBytes(Size)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		EncodeToString(data)
	}
}

func Benchmark_byteSliceToString(b *testing.B) {
	data := []byte("test")
	b.SetBytes(int64(len(data)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		byteSliceToString(data[:])
	}
}

func Benchmark_stringToByteSlice(b *testing.B) {
	data := "test"
	b.SetBytes(int64(len([]byte(data))))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stringToByteSlice(data)
	}
}
