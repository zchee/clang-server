// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/zchee/clang-server/indexdb"
	"github.com/zchee/clang-server/internal/log"
	"github.com/zchee/clang-server/internal/pathutil"
	"github.com/zchee/clang-server/symbol/internal/symbol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":50051"

func NewClangClient(cc *grpc.ClientConn) symbol.ClangClient {
	return symbol.NewClangClient(cc)
}

type server struct {
	db *indexdb.IndexDB
}

func (s *server) Completion(ctx context.Context, loc *symbol.Location) (*flatbuffers.Builder, error) {
	f := string(loc.FileName())
	dir, _ := pathutil.FindProjectRoot(f)
	db, err := indexdb.NewIndexDB(dir)
	if err != nil {
		return nil, err
	}
	s.db = db
	defer db.Close()

	buf, err := db.Get(f)
	if err != nil {
		return nil, err
	}

	file := GetRootAsFile(buf, 0)

	return file.Serialize(), nil
}

func Serve() {
	println("Serve")
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	symbol.RegisterClangServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
