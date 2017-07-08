// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	blake2b "github.com/minio/blake2b-simd"
	"github.com/zchee/clang-server/internal/hashutil"
)

// ID id of cursor.USR with blake2b hash.
type ID [blake2b.Size]byte

func (id ID) String() string {
	return hashutil.EncodeToString(id[:])
}

// Bytes reterun the ID byte slice.
func (id ID) Bytes() []byte {
	return id[:]
}

// IsEmpty reports whether id is empty.
func (id ID) IsEmpty() bool {
	return len(id) == 0
}

// FileID id of filename with blake2b hash.
type FileID [blake2b.Size]byte

func (id FileID) String() string {
	return hashutil.EncodeToString(id[:])
}

// Bytes reterun the FileID byte slice.
func (id FileID) Bytes() []byte {
	return id[:]
}

// IsEmpty reports whether id is empty.
func (id FileID) IsEmpty() bool {
	return len(id) == 0
}

// ToID converts the string to blake2b sum512 hash.
func ToID(s string) ID {
	return hashutil.NewHashString(s)
}

// ToFileID converts the string to blake2b sum512 hash.
func ToFileID(s string) FileID {
	return hashutil.NewHashString(s)
}
