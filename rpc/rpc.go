// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"net"
	"time"

	"github.com/go-clang/v3.9/clang"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/zchee/clang-server/indexdb"
	"github.com/zchee/clang-server/internal/log"
	"github.com/zchee/clang-server/internal/pathutil"
	"github.com/zchee/clang-server/symbol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":50051"

// GRPCServer represents a clang-server gRPC server.
type server struct {
	db       *indexdb.IndexDB
	filename string
	idx      clang.Index
	tu       clang.TranslationUnit
}

type GRPCServer struct {
	server
}

// NewGRPCServer return the new Server with initialize idx.
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		server{
			idx: clang.NewIndex(0, 0),
		},
	}
}

// Serve serve clang-server server with flatbuffers gRPC custom codec.
func (s *server) Serve() {
	println("Serve")
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	symbol.RegisterClangServer(grpcServer, s)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// Completion implements symbol.ClangServer Completion interface.
func (s *server) Completion(ctx context.Context, loc *symbol.SymbolLocation) (*flatbuffers.Builder, error) {
	defer profile(time.Now(), "Completion")
	f := string(loc.FileName())

	if s.filename != f {
		s.filename = f
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

		file := symbol.GetRootAsFile(buf, 0)

		if cErr := s.idx.ParseTranslationUnit2(file.Name(), file.Flags(), nil, clang.DefaultEditingTranslationUnitOptions()|clang.DefaultCodeCompleteOptions()|uint32(clang.TranslationUnit_KeepGoing), &s.tu); clang.ErrorCode(cErr) != clang.Error_Success {
			log.Fatal(cErr)
		}
	}

	codeCompleteResults := new(symbol.CodeCompleteResults)
	result := codeCompleteResults.Marshal(s.tu.CodeCompleteAt(f, loc.Line(), loc.Col(), nil, clang.DefaultCodeCompleteOptions()))

	return result, nil
}

func profile(start time.Time, name string) {
	elapsed := time.Since(start).Seconds()
	log.Debugf("%s: %fsec\n", name, elapsed)
}
