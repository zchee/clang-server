#ifndef GO_CLANG
#define GO_CLANG

#include <stdlib.h>

#include "clang-c/Index.h"

unsigned go_clang_visit_children(CXCursor c, void *fct);

#endif
