// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hex is an efficient hexadecimal implementation for Golang.
package hex

import (
	"errors"
	"fmt"
)

// ErrLength results from decoding an odd length slice.
var ErrLength = errors.New("go-hex: odd length hex string")

var (
	lower = []byte("0123456789abcdef")
	upper = []byte("0123456789ABCDEF")
)

// InvalidByteError values describe errors resulting from an invalid byte in a hex string.
type InvalidByteError byte

func (e InvalidByteError) Error() string {
	return fmt.Sprintf("go-hex: invalid byte: %#U", rune(e))
}

// EncodedLen returns the length of an encoding of n source bytes.
func EncodedLen(n int) int {
	return n * 2
}

// DecodedLen returns the length of a decoding of n source bytes.
func DecodedLen(n int) int {
	return n / 2
}

// Encode encodes src into EncodedLen(len(src))
// bytes of dst.  As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements lowercase hexadecimal encoding.
func Encode(dst, src []byte) int {
	return RawEncode(dst, src, lower)
}

// EncodeUC encodes src into EncodedLen(len(src))
// bytes of dst.  As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// EncodeUC implements uppercase hexadecimal encoding.
func EncodeUC(dst, src []byte) int {
	return RawEncode(dst, src, upper)
}

// EncodeToString returns the lowercase hexadecimal encoding of src.
func EncodeToString(src []byte) string {
	return RawEncodeToString(src, lower)
}

// EncodeUCToString returns the uppercase hexadecimal encoding of src.
func EncodeUCToString(src []byte) string {
	return RawEncodeToString(src, upper)
}

// RawEncodeToString returns the hexadecimal encoding of src for a given
// alphabet.
func RawEncodeToString(src, alpha []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	RawEncode(dst, src, alpha)
	return string(dst)
}

// DecodeString returns the bytes represented by the hexadecimal string s.
func DecodeString(s string) ([]byte, error) {
	src := []byte(s)
	dst := make([]byte, DecodedLen(len(src)))

	if _, err := Decode(dst, src); err != nil {
		return nil, err
	}

	return dst, nil
}
