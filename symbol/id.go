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
	return hashutil.EncodeToString(id)
}

func (id ID) Bytes() []byte {
	return id[:]
}

// FileID id of filename with blake2b hash.
type FileID [blake2b.Size]byte

func (id FileID) String() string {
	return hashutil.EncodeToString(id)
}

func (id FileID) Bytes() []byte {
	return id[:]
}

var (
	NoID     ID     = [64]byte{} // empty
	NoFileID FileID = [64]byte{} // empty
)

// ToID converts the string to blake2b sum512 hash.
func ToID(s string) ID {
	return hashutil.NewHashString(s)
}

// ToFileID converts the string to blake2b sum512 hash.
func ToFileID(s string) FileID {
	return hashutil.NewHashString(s)
}
