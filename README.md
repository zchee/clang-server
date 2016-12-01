# clang-server

A C/C++ AST index server using libclang over the msgpack-rpc written in Go.

## Concept

* Fast indexing of C/C++ AST database onto the NoSQL
 * Now using the [leveldb][leveldb] key-value storage
 * Without C bindings using the [syndtr/goleveldb][goleveldb], which is natively implemented leveldb in Go
* Support cross-platform and multi-architecture AST indexing
 * Linux, macOS, BSD and Windows
 * arm, arm64 m68k, mips, sparc and x86_(16|32|64)
* Server/Client architecture over the msgpack-rpc
* Built-in `compile_commands.json` generator using [google/kati][kati] and [ninja][ninja] for `Makefile`
 * No need `make` for the generating `compile_commands.json`


[leveldb]: https://github.com/google/leveldb
[goleveldb]: https://github.com/syndtr/goleveldb
[kati]: https://github.com/google/kati
[ninja]: https://github.com/ninja-build/ninja
