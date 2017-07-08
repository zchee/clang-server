// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package indexdb

import (
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zchee/clang-server/internal/pathutil"
)

// index filename of leveldb which indexed C/C++ symbols.
const index = "index.db"

// IndexDB represets a C/C++ file database using leveldb.
type IndexDB struct {
	root string
	ldb  *leveldb.DB
}

// NewIndexDB creates the project cache directory, open the leveldb db file
// on same location and return the new IndexDB.
func NewIndexDB(root string) (*IndexDB, error) {
	dir := pathutil.ProjectCacheDir(root)
	db, err := leveldb.OpenFile(filepath.Join(dir, index), nil)
	if err != nil {
		return nil, err
	}

	return &IndexDB{
		root: root,
		ldb:  db,
	}, nil
}

// Close closes the leveldb.
func (i *IndexDB) Close() error {
	return i.ldb.Close()
}

// Put puts the selialized C/C++ files symbol data to leveldb.
// The key is using blake2b hashed filename.
func (i *IndexDB) Put(key []byte, value []byte) error {
	return i.ldb.Put(key[:], value, nil)
}

// Get gets the selialized C/C++ files symbol data from leveldb.
// The key is using blake2b hashed filename.
func (i *IndexDB) Get(key []byte) ([]byte, error) {
	return i.ldb.Get(key[:], nil)
}

// Has reports whether filename symbol data on leveldb.
// The key is using blake2b hashed filename.
func (i *IndexDB) Has(key []byte) bool {
	has, err := i.ldb.Has(key[:], nil)
	if err != nil {
		return false
	}
	return has
}
