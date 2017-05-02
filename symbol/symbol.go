// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package symbol

import (
	"github.com/go-clang/v3.9/clang"
	"github.com/zchee/clang-server/internal/symbol"
	"google.golang.org/grpc"
)

// FromCursor return the location of symbol from cursor.
func FromCursor(cursor clang.Cursor) Location {
	if cursor.IsNull() {
		return Location{}
	}

	usr := cursor.USR()
	if usr == "" && cursor.Kind() == clang.Cursor_MacroExpansion {
		usr = cursor.DisplayName()
	}

	file, line, col, offset := cursor.Location().FileLocation()

	return Location{
		fileName: file.Name(),
		line:     line,
		col:      col,
		offset:   offset,
		usr:      usr,
	}
}

// NewClangClient retern the new symbol.ClangClient.
func NewClangClient(cc *grpc.ClientConn) symbol.ClangClient {
	return symbol.NewClangClient(cc)
}

// RegisterClangServer register a service and its implementation to the gRPC server.
func RegisterClangServer(s *grpc.Server, srv symbol.ClangServer) {
	symbol.RegisterClangServer(s, srv)
}
