# add 'static' Go build tags for go-clang
GO_BUILD_TAGS += static

# add clang supported latest c++ std version and libc++ stdlib
CGO_CXXFLAGS += -std=c++1z -stdlib=libc++

ifeq ($(UNAME),Linux)
	# darwin ld(64) linker doesn't support "-extldflags='-static'" flag, because that is basicallly for building the xnu kernel.
	# Also will occur 'not found -crt0.o object file' error.
	# If install the 'Csu' from the opensource.apple.com, passes -crt0.o error, but needs libpthread.a static library.
	GO_LDFLAGS += -extldflags=-static
	# -Bstatic: Do not link against shared libraries.
	# for statically link the libclang libraries.
	CGO_LDFLAGS += -Wl,-Bstatic
endif

# add LLVM dependencies static libraries
CGO_LDFLAGS += $(shell llvm-config --libfiles --link-static)

LIBCLANG_STATIC_LIBS := \
	libclang \
	libclangARCMigrate \
	libclangAST \
	libclangASTMatchers \
	libclangAnalysis \
	libclangApplyReplacements \
	libclangBasic \
	libclangChangeNamespace \
	libclangCodeGen \
	libclangDriver \
	libclangDynamicASTMatchers \
	libclangEdit \
	libclangFormat \
	libclangFrontend \
	libclangFrontendTool \
	libclangIncludeFixer \
	libclangIncludeFixerPlugin \
	libclangIndex \
	libclangLex \
	libclangMove \
	libclangParse \
	libclangQuery \
	libclangRename \
	libclangReorderFields \
	libclangRewrite \
	libclangRewriteFrontend \
	libclangSema \
	libclangSerialization \
	libclangStaticAnalyzerCheckers \
	libclangStaticAnalyzerCore \
	libclangStaticAnalyzerFrontend \
	libclangTidy \
	libclangTidyBoostModule \
	libclangTidyCERTModule \
	libclangTidyCppCoreGuidelinesModule \
	libclangTidyGoogleModule \
	libclangTidyHICPPModule \
	libclangTidyLLVMModule \
	libclangTidyMPIModule \
	libclangTidyMiscModule \
	libclangTidyModernizeModule \
	libclangTidyPerformanceModule \
	libclangTidyPlugin \
	libclangTidyReadabilityModule \
	libclangTidyUtils \
	libclangTooling \
	libclangToolingCore \
	libclangToolingRefactor \
	libfindAllSymbols

# add libclang static libraries
CGO_LDFLAGS += $(foreach lib,$(LIBCLANG_STATIC_LIBS),$(LLVM_LIBDIR)/$(lib).a)

ifeq ($(UNAME),Linux)
	# -Bdynamic: Link against dynamic libraries.
	# End of libclang static libraries list.
	CGO_LDFLAGS += -Wl,-Bdynamic
endif

# add LLVM dependency system library. such as libm, ncurses and zlib.
CGO_LDFLAGS += $(shell llvm-config --system-libs --link-static)

# add '-c++' for only Darwin.
# avoid 'Undefined symbols for architecture x86_64 "std::__1::__shared_count::__add_shared()"' or etc.
CGO_LDFLAGS += $(if $(findstring Darwin,$(UNAME)),-lc++,)
