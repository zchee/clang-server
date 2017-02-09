// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package indexdb

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zchee/clang-server/internal/hashutil"
	"github.com/zchee/clang-server/internal/pathutil"
)

const (
	index = "index.db"
)

type IndexDB struct {
	root string
	db   *leveldb.DB
}

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

func (i *IndexDB) Close() error {
	return i.db.Close()
}

func (i *IndexDB) Put(filename string, data []byte) error {
	log.Debug("put")
	fileID := hashutil.NewHashString(filename)
	return i.db.Put(hashutil.Encode(fileID), data, nil)
}

func (i *IndexDB) Get(filename string) ([]byte, error) {
	log.Debug("get")
	fileID := hashutil.NewHashString(filename)
	return i.db.Get(hashutil.Encode(fileID), nil)
}

func (i *IndexDB) Has(filename string) bool {
	fileID := hashutil.NewHashString(filename)
	has, err := i.db.Has(hashutil.Encode(fileID), nil)
	if err != nil {
		return false
	}
	return has
}
