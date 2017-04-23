// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashutil

import (
	"reflect"
	"unsafe"

	blake2b "github.com/minio/blake2b-simd"
	hex "github.com/tmthrgd/go-hex"
)

// NewHash converts the b to blake2b sum512 hash.
func NewHash(b []byte) [blake2b.Size]byte {
	return blake2b.Sum512(b)
}

// NewHashString converts the s to blake2b sum512 hash.
func NewHashString(s string) [blake2b.Size]byte {
	return blake2b.Sum512(stringToByteSlice(s))
}

// Encode returns the hexadecimal encoded byte slice of blake2b hashed b.
func Encode(b [blake2b.Size]byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b[:])
	return dst
}

// EncodeToString returns the hexadecimal encoded string of blake2b hashed b.
func EncodeToString(b [blake2b.Size]byte) string {
	return hex.EncodeToString(b[:])
}

// byteSliceToString converts the []byte to string without a heap allocation.
func byteSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  len(b),
	}))
}

// stringToByteSlice converts the string to []byte without a heap allocation.
func stringToByteSlice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  len(s),
		Cap:  len(s),
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
	}))
}
