// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	"net"

	"github.com/go-clang/v3.9/clang"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/zchee/clang-server/indexdb"
	"github.com/zchee/clang-server/internal/log"
	"github.com/zchee/clang-server/internal/pathutil"
	"github.com/zchee/clang-server/internal/symbol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":50051"

// NewClangClient retern the new symbol.ClangClient.
func NewClangClient(cc *grpc.ClientConn) symbol.ClangClient {
	return symbol.NewClangClient(cc)
}

// Server represents a clang-server gRPC server.
type Server struct {
	db       *indexdb.IndexDB
	filename string
	idx      clang.Index
	tu       clang.TranslationUnit
}

// NewServer return the new Server with initialize idx.
func NewServer() *Server {
	return &Server{
		idx: clang.NewIndex(0, 0),
	}
}

// Serve serve clang-server server with the flatbuffers gRPC custom codec.
func (s *Server) Serve() {
	println("Serve")
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	gprcServer := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	symbol.RegisterClangServer(gprcServer, s)
	if err := gprcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

// Completion implements symbol.ClangServer Completion interface.
func (s *Server) Completion(ctx context.Context, loc *symbol.Location) (*flatbuffers.Builder, error) {
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

		file := GetRootAsFile(buf, 0)

		if cErr := s.idx.ParseTranslationUnit2(file.Name(), file.Flags(), nil, clang.DefaultEditingTranslationUnitOptions()|uint32(clang.TranslationUnit_KeepGoing), &s.tu); clang.ErrorCode(cErr) != clang.Error_Success {
			log.Fatal(cErr)
		}
	}

	codeCompleteResults := new(CodeCompleteResults)
	result := codeCompleteResults.Marshal(s.tu.CodeCompleteAt(f, loc.Line(), loc.Col(), nil, clang.DefaultCodeCompleteOptions()))

	return result, nil
}
