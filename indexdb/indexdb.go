// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package indexdb

import (
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zchee/clang-server/internal/hashutil"
	"github.com/zchee/clang-server/internal/pathutil"
)

// index filename of leveldb which indexed C/C++ symbols.
const index = "index.db"

// IndexDB represets a C/C++ file database using leveldb.
type IndexDB struct {
	root string
	db   *leveldb.DB
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
		db:   db,
	}, nil
}

// Close closes the leveldb.
func (i *IndexDB) Close() error {
	return i.db.Close()
}

// Put puts the selialized C/C++ files symbol data to leveldb.
// The key is using blake2b hashed filename.
func (i *IndexDB) Put(filename string, data []byte) error {
	fileID := hashutil.NewHashString(filename)
	return i.db.Put(hashutil.Encode(fileID), data, nil)
}

// Get gets the selialized C/C++ files symbol data from leveldb.
// The key is using blake2b hashed filename.
func (i *IndexDB) Get(filename string) ([]byte, error) {
	fileID := hashutil.NewHashString(filename)
	return i.db.Get(hashutil.Encode(fileID), nil)
}

// Has reports whether filename symbol data on leveldb.
// The key is using blake2b hashed filename.
func (i *IndexDB) Has(filename string) bool {
	fileID := hashutil.NewHashString(filename)
	has, err := i.db.Has(hashutil.Encode(fileID), nil)
	if err != nil {
		return false
	}
	return has
}
