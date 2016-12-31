clang
=====

command line flags
------------------

```sh
clang version 4.0.0 (http://llvm.org/git/clang.git f89729f1029a1ac266ea60184cd8bd7a6ee8631a)
Target: x86_64-apple-darwin16.3.0
Thread model: posix
```

```sh
DEBUG/DEVELOPMENT OPTIONS:
  -ccc-install-dir <value> Simulate installation in the given directory
  -ccc-print-bindings Show bindings of tools to actions
  -ccc-print-phases   Dump list of actions to perform
  -via-file-asm       Write assembly to file for input to assemble jobs

DRIVER OPTIONS:
  -ccc-arcmt-check      Check for ARC migration issues that need manual handling
  -ccc-arcmt-migrate <value> Apply modifications and produces temporary files that conform to ARC
  -ccc-arcmt-modify     Apply modifications to files to conform to ARC
  -ccc-gcc-name <gcc-path> Name for native GCC compiler
  -ccc-objcmt-migrate <value>
                        Apply modifications and produces temporary files to migrate to modern ObjC syntax
  -ccc-pch-is-pch       Use lazy PCH for precompiled headers
  -ccc-pch-is-pth       Use pretokenized headers for precompiled headers
  --driver-mode=<value> Set the driver mode to either 'gcc', 'g++', 'cpp', or 'cl'
  --rsp-quoting=<value> Set the rsp quoting to either 'posix', or 'windows'

OPTIONS:
  -###                    Print (but do not run) the commands to run for this compilation
  --analyzer-output <value>
                          Static analyzer report output format (html|plist|plist-multi-file|plist-html|text).
  --analyze               Run the static analyzer
  -arcmt-migrate-emit-errors
                          Emit ARC errors even if the migrator can fix them
  -arcmt-migrate-report-output <value>
                          Output path for the plist report
  -cl-denorms-are-zero    OpenCL only. Allow denormals to be flushed to zero.
  -cl-fast-relaxed-math   OpenCL only. Sets -cl-finite-math-only and -cl-unsafe-math-optimizations, and defines __FAST_RELAXED_MATH__.
  -cl-finite-math-only    OpenCL only. Allow floating-point optimizations that assume arguments and results are not NaNs or +-Inf.
  -cl-fp32-correctly-rounded-divide-sqrt
                          OpenCL only. Specify that single precision floating-point divide and sqrt used in the program source are correctly rounded.
  -cl-kernel-arg-info     OpenCL only. Generate kernel argument metadata.
  -cl-mad-enable          OpenCL only. Allow use of less precise MAD computations in the generated binary.
  -cl-no-signed-zeros     OpenCL only. Allow use of less precise no signed zeros computations in the generated binary.
  -cl-opt-disable         OpenCL only. This option disables all optimizations. By default optimizations are enabled.
  -cl-single-precision-constant
                          OpenCL only. Treat double precision floating-point constant as single precision constant.
  -cl-std=<value>         OpenCL language standard to compile for.
  -cl-strict-aliasing     OpenCL only. This option is added for compatibility with OpenCL 1.0.
  -cl-unsafe-math-optimizations
                          OpenCL only. Allow unsafe floating-point optimizations.  Also implies -cl-no-signed-zeros and -cl-mad-enable.
  --cuda-compile-host-device
                          Compile CUDA code for both host and device (default).  Has no effect on non-CUDA compilations.
  --cuda-device-only      Compile CUDA code for device only
  --cuda-gpu-arch=<value> CUDA GPU architecture (e.g. sm_35).  May be specified more than once.
  --cuda-host-only        Compile CUDA code for host only.  Has no effect on non-CUDA compilations.
  --cuda-noopt-device-debug
                          Enable device-side debug info generation. Disables ptxas optimizations.
  --cuda-path=<value>     CUDA installation path
  -cxx-isystem <directory>
                          Add directory to the C++ SYSTEM include search path
  -c                      Only run preprocess, compile, and assemble steps
  -dD                     Print macro definitions in -E mode in addition to normal output
  -dependency-dot <value> Filename to write DOT-formatted header dependencies to
  -dependency-file <value>
                          Filename (or -) to write dependency output to
  -dI                     Print include directives in -E mode in addition to normal output
  -dM                     Print macro definitions in -E mode instead of normal output
  -emit-ast               Emit Clang AST files for source inputs
  -emit-llvm              Use the LLVM representation for assembler and object files
  -E                      Only run the preprocessor
  -faligned-allocation    Enable C++17 aligned allocation functions
  -faltivec               Enable AltiVec vector initializer syntax
  -fansi-escape-codes     Use ANSI escape codes for diagnostics
  -fapple-kext            Use Apple's kernel extensions ABI
  -fapple-pragma-pack     Enable Apple gcc-compatible #pragma pack handling
  -fapplication-extension Restrict code to those available for App Extensions
  -fblocks                Enable the 'blocks' language feature
  -fborland-extensions    Accept non-standard constructs supported by the Borland compiler
  -fbuild-session-file=<file>
                          Use the last modification time of <file> as the build session timestamp
  -fbuild-session-timestamp=<time since Epoch in seconds>
                          Time when the current build session started
  -fbuiltin-module-map    Load the clang builtins module map file.
  -fcolor-diagnostics     Use colors in diagnostics
  -fcomment-block-commands=<arg>
                          Treat each comma separated argument in <arg> as a documentation comment block command
  -fcoroutines-ts         Enable support for the C++ Coroutines TS
  -fcoverage-mapping      Generate coverage mapping to enable code coverage analysis
  -fcuda-approx-transcendentals
                          Use approximate transcendental functions
  -fcuda-flush-denormals-to-zero
                          Flush denormal floating point values to zero in CUDA device mode.
  -fcxx-exceptions        Enable C++ exceptions
  -fdata-sections         Place each data in its own section (ELF Only)
  -fdebug-prefix-map=<value>
                          remap file source paths in debug info
  -fdebug-types-section   Place debug types in their own section (ELF Only)
  -fdeclspec              Allow __declspec as a keyword
  -fdelayed-template-parsing
                          Parse templated function definitions at the end of the translation unit
  -fdiagnostics-absolute-paths
                          Print absolute paths in diagnostics
  -fdiagnostics-parseable-fixits
                          Print fix-its in machine parseable form
  -fdiagnostics-print-source-range-info
                          Print source range spans in numeric form
  -fdiagnostics-show-hotness
                          Enable profile hotness information in diagnostic line
  -fdiagnostics-show-note-include-stack
                          Display include stacks for diagnostic notes
  -fdiagnostics-show-option
                          Print option name with mappable diagnostics
  -fdiagnostics-show-template-tree
                          Print a template comparison tree for differing templates
  -fdollars-in-identifiers
                          Allow '$' in identifiers
  -fembed-bitcode-marker  Embed placeholder LLVM IR data as a marker
  -fembed-bitcode=<option>
                          Embed LLVM bitcode (option: off, all, bitcode, marker)
  -fembed-bitcode         Embed LLVM IR bitcode as data
  -femit-all-decls        Emit all declarations, even if unused
  -femulated-tls          Use emutls functions to access thread_local variables
  -fexceptions            Enable support for exception handling
  -ffast-math             Allow aggressive, lossy floating-point optimizations
  -ffixed-r9              Reserve the r9 register (ARM only)
  -ffixed-x18             Reserve the x18 register (AArch64 only)
  -ffp-contract=<value>   Form fused FP ops (e.g. FMAs): fast (everywhere) | on (according to FP_CONTRACT pragma, default) | off (never fuse)
  -ffreestanding          Assert that the compilation takes place in a freestanding environment
  -ffunction-sections     Place each function in its own section (ELF Only)
  -fgnu-keywords          Allow GNU-extension keywords regardless of language standard
  -fgnu-runtime           Generate output compatible with the standard GNU Objective-C runtime
  -fgnu89-inline          Use the gnu89 inline semantics
  -fimplicit-module-maps  Implicitly search the file system for module map files.
  -finline-functions      Inline suitable functions
  -finline-hint-functions Inline functions wich are (explicitly or implicitly) marked inline
  -finstrument-functions  Generate calls to instrument function entry and exit
  -fintegrated-as         Enable the integrated assembler
  -flto-jobs=<value>      Controls the backend parallelism of -flto=thin (default of 0 means the number of threads will be derived from the number of CPUs detected)
  -flto=<value>           Set LTO mode to either 'full' or 'thin'
  -flto                   Enable LTO in 'full' mode
  -fmath-errno            Require math functions to indicate errors by setting errno
  -fmax-type-align=<value>
                          Specify the maximum alignment to enforce on pointers lacking an explicit alignment
  -fmodule-file=<file>    Load this precompiled module file
  -fmodule-map-file=<file>
                          Load this module map file
  -fmodule-name=<name>    Specify the name of the module to build
  -fmodules-cache-path=<directory>
                          Specify the module cache path
  -fmodules-decluse       Require declaration of modules used within a module
  -fmodules-disable-diagnostic-validation
                          Disable validation of the diagnostic options when loading the module
  -fmodules-ignore-macro=<value>
                          Ignore the definition of the given macro when building and loading modules
  -fmodules-prune-after=<seconds>
                          Specify the interval (in seconds) after which a module file will be considered unused
  -fmodules-prune-interval=<seconds>
                          Specify the interval (in seconds) between attempts to prune the module cache
  -fmodules-search-all    Search even non-imported modules to resolve references
  -fmodules-strict-decluse
                          Like -fmodules-decluse but requires all headers to be in modules
  -fmodules-ts            Enable support for the C++ Modules TS
  -fmodules-user-build-path <directory>
                          Specify the module user build path
  -fmodules-validate-once-per-build-session
                          Don't verify input files for the modules if the module has been successfully validated or loaded during this build session
  -fmodules-validate-system-headers
                          Validate the system headers that a module depends on when loading the module
  -fmodules               Enable the 'modules' language feature
  -fms-compatibility-version=<value>
                          Dot-separated value representing the Microsoft compiler version number to report in _MSC_VER (0 = don't define it (default))
  -fms-compatibility      Enable full Microsoft Visual C++ compatibility
  -fms-extensions         Accept some non-standard constructs supported by the Microsoft compiler
  -fmsc-version=<value>   Microsoft compiler version number to report in _MSC_VER (0 = don't define it (default))
  -fnew-alignment=<align> Specifies the largest alignment guaranteed by '::operator new(size_t)'
  -fno-access-control     Disable C++ access control
  -fno-assume-sane-operator-new
                          Don't assume that C++'s global operator new can't alias any pointer
  -fno-autolink           Disable generation of linker directives for automatic library linking
  -fno-builtin-<value>    Disable implicit builtin knowledge of a specific function
  -fno-builtin            Disable implicit builtin knowledge of functions
  -fno-common             Compile common globals like normal definitions
  -fno-constant-cfstrings Disable creation of CodeFoundation-type constant strings
  -fno-coverage-mapping   Disable code coverage analysis
  -fno-declspec           Disallow __declspec as a keyword
  -fno-diagnostics-fixit-info
                          Do not include fixit information in diagnostics
  -fno-dollars-in-identifiers
                          Disallow '$' in identifiers
  -fno-elide-constructors Disable C++ copy constructor elision
  -fno-elide-type         Do not elide types when printing diagnostics
  -fno-gnu-inline-asm     Disable GNU style inline asm
  -fno-integrated-as      Disable the integrated assembler
  -fno-jump-tables        Do not use jump tables for lowering switches
  -fno-lax-vector-conversions
                          Disallow implicit conversions between vectors with a different number of elements or different element types
  -fno-lto                Disable LTO mode (default)
  -fno-merge-all-constants
                          Disallow merging of constants
  -fno-objc-infer-related-result-type
                          do not infer Objective-C related result type based on method family
  -fno-operator-names     Do not treat C++ operator name keywords as synonyms for operators
  -fno-preserve-as-comments
                          Do not preserve comments in inline assembly
  -fno-profile-generate   Disable generation of profile instrumentation.
  -fno-profile-instr-generate
                          Disable generation of profile instrumentation.
  -fno-profile-instr-use  Disable using instrumentation data for profile-guided optimization
  -fno-reroll-loops       Turn off loop reroller
  -fno-rtti               Disable generation of rtti information
  -fno-sanitize-address-use-after-scope
                          Disable use-after-scope detection in AddressSanitizer
  -fno-sanitize-blacklist Don't use blacklist file for sanitizers
  -fno-sanitize-cfi-cross-dso
                          Disable control flow integrity (CFI) checks for cross-DSO calls.
  -fno-sanitize-coverage=<value>
                          Disable specified features of coverage instrumentation for Sanitizers
  -fno-sanitize-memory-track-origins
                          Disable origins tracking in MemorySanitizer
  -fno-sanitize-recover=<value>
                          Disable recovery for specified sanitizers
  -fno-sanitize-stats     Disable sanitizer statistics gathering.
  -fno-sanitize-thread-atomics
                          Disable atomic operations instrumentation in ThreadSanitizer
  -fno-sanitize-thread-func-entry-exit
                          Disable function entry/exit instrumentation in ThreadSanitizer
  -fno-sanitize-thread-memory-access
                          Disable memory access instrumentation in ThreadSanitizer
  -fno-sanitize-trap=<value>
                          Disable trapping for specified sanitizers
  -fno-short-wchar        Force wchar_t to be an unsigned int
  -fno-show-column        Do not include column number on diagnostics
  -fno-show-source-location
                          Do not include source location information with diagnostics
  -fno-signed-char        Char is unsigned
  -fno-signed-zeros       Allow optimizations that ignore the sign of floating point zeros
  -fno-spell-checking     Disable spell-checking
  -fno-stack-protector    Disable the use of stack protectors
  -fno-standalone-debug   Limit debug information produced to reduce size of debug binary
  -fno-threadsafe-statics Do not emit code to make initialization of local statics thread safe
  -fno-trigraphs          Do not process trigraph sequences
  -fno-unroll-loops       Turn off loop unroller
  -fno-use-cxa-atexit     Don't use __cxa_atexit for calling destructors
  -fno-use-init-array     Don't use .init_array instead of .ctors
  -fobjc-arc-exceptions   Use EH-safe code when synthesizing retains and releases in -fobjc-arc
  -fobjc-arc              Synthesize retain and release calls for Objective-C pointers
  -fobjc-exceptions       Enable Objective-C exceptions
  -fobjc-runtime=<value>  Specify the target Objective-C runtime kind and version
  -fobjc-weak             Enable ARC-style weak references in Objective-C
  -fopenmp-targets=<value>
                          Specify comma-separated list of triples OpenMP offloading targets to be supported
  -foptimization-record-file=<value>
                          Specify the file name of any generated YAML optimization record
  -fpack-struct=<value>   Specify the default maximum struct packing alignment
  -fpascal-strings        Recognize and construct Pascal-style string literals
  -fpcc-struct-return     Override the default ABI to return all structs on the stack
  -fplugin=<dsopath>      Load the named plugin (dynamic shared object)
  -fprebuilt-module-path=<directory>
                          Specify the prebuilt module path
  -fprofile-generate=<directory>
                          Generate instrumented code to collect execution counts into <directory>/default.profraw (overridden by LLVM_PROFILE_FILE env var)
  -fprofile-generate      Generate instrumented code to collect execution counts into default.profraw (overridden by LLVM_PROFILE_FILE env var)
  -fprofile-instr-generate=<file>
                          Generate instrumented code to collect execution counts into <file> (overridden by LLVM_PROFILE_FILE env var)
  -fprofile-instr-generate
                          Generate instrumented code to collect execution counts into default.profraw file (overriden by '=' form of option or LLVM_PROFILE_FILE env var)
  -fprofile-instr-use=<value>
                          Use instrumentation data for profile-guided optimization
  -fprofile-sample-use=<value>
                          Enable sample-based profile guided optimizations
  -fprofile-use=<pathname>
                          Use instrumentation data for profile-guided optimization. If pathname is a directory, it reads from <pathname>/default.profdata. Otherwise, it reads from file <pathname>.
  -freciprocal-math       Allow division operations to be reassociated
  -freg-struct-return     Override the default ABI to return small structs in registers
  -freroll-loops          Turn on loop reroller
  -fsanitize-address-field-padding=<value>
                          Level of field padding for AddressSanitizer
  -fsanitize-address-use-after-scope
                          Enable use-after-scope detection in AddressSanitizer
  -fsanitize-blacklist=<value>
                          Path to blacklist file for sanitizers
  -fsanitize-cfi-cross-dso
                          Enable control flow integrity (CFI) checks for cross-DSO calls.
  -fsanitize-coverage=<value>
                          Specify the type of coverage instrumentation for Sanitizers
  -fsanitize-memory-track-origins=<value>
                          Enable origins tracking in MemorySanitizer
  -fsanitize-memory-track-origins
                          Enable origins tracking in MemorySanitizer
  -fsanitize-memory-use-after-dtor
                          Enable use-after-destroy detection in MemorySanitizer
  -fsanitize-recover=<value>
                          Enable recovery for specified sanitizers
  -fsanitize-stats        Enable sanitizer statistics gathering.
  -fsanitize-thread-atomics
                          Enable atomic operations instrumentation in ThreadSanitizer (default)
  -fsanitize-thread-func-entry-exit
                          Enable function entry/exit instrumentation in ThreadSanitizer (default)
  -fsanitize-thread-memory-access
                          Enable memory access instrumentation in ThreadSanitizer (default)
  -fsanitize-trap=<value> Enable trapping for specified sanitizers
  -fsanitize-undefined-strip-path-components=<number>
                          Strip (or keep only, if negative) a given number of path components when emitting check metadata.
  -fsanitize=<check>      Turn on runtime checks for various forms of undefined or suspicious behavior. See user manual for available checks
  -fsave-optimization-record
                          Generate a YAML optimization record file
  -fshort-enums           Allocate to an enum type only as many bytes as it needs for the declared range of possible values
  -fshort-wchar           Force wchar_t to be a short unsigned int
  -fshow-overloads=<value>
                          Which overload candidates to show when overload resolution fails: best|all; defaults to all
  -fsized-deallocation    Enable C++14 sized global deallocation functions
  -fsjlj-exceptions       Use SjLj style exceptions
  -fslp-vectorize-aggressive
                          Enable the BB vectorization passes
  -fslp-vectorize         Enable the superword-level parallelism vectorization passes
  -fsplit-dwarf-inlining  Place debug types in their own section (ELF Only)
  -fstack-protector-all   Force the usage of stack protectors for all functions
  -fstack-protector-strong
                          Use a strong heuristic to apply stack protectors to functions
  -fstack-protector       Enable stack protectors for functions potentially vulnerable to stack smashing
  -fstandalone-debug      Emit full debug info for all types used by the program
  -fstrict-enums          Enable optimizations based on the strict definition of an enum's value range
  -fstrict-vtable-pointers
                          Enable optimizations based on the strict rules for overwriting polymorphic C++ objects
  -fthinlto-index=<value> Perform ThinLTO importing using provided function summary index
  -ftrap-function=<value> Issue call to specified function rather than a trap instruction
  -ftrapv-handler=<function name>
                          Specify the function to be called on overflow
  -ftrapv                 Trap on integer overflow
  -ftrigraphs             Process trigraph sequences
  -funique-section-names  Use unique names for text and data sections (ELF Only)
  -funroll-loops          Turn on loop unroller
  -fuse-init-array        Use .init_array instead of .ctors
  -fveclib=<value>        Use the given vector functions library
  -fvectorize             Enable the loop vectorization passes
  -fvisibility-inlines-hidden
                          Give inline C++ member functions default visibility by default
  -fvisibility-ms-compat  Give global types 'default' visibility and global functions and variables 'hidden' visibility by default
  -fvisibility=<value>    Set the default symbol visibility for all global declarations
  -fwhole-program-vtables Enables whole-program vtable optimization. Requires -flto
  -fwrapv                 Treat signed integer overflow as two's complement
  -fwritable-strings      Store string literals as writable data
  -fxray-instruction-threshold= <value>
                          Sets the minimum function size to instrument with XRay
  -fxray-instrument       Generate XRay instrumentation sleds on function entry and exit
  -fzvector               Enable System z vector language extension
  -F <value>              Add directory to framework include search path
  --gcc-toolchain=<value> Use the gcc toolchain at the given directory
  -gcodeview              Generate CodeView debug information
  -gdwarf-2               Generate source-level debug information with dwarf version 2
  -gdwarf-3               Generate source-level debug information with dwarf version 3
  -gdwarf-4               Generate source-level debug information with dwarf version 4
  -gdwarf-5               Generate source-level debug information with dwarf version 5
  -gline-tables-only      Emit debug line number tables only
  -gmodules               Generate debug info with external references to clang modules or precompiled headers
  -g                      Generate source-level debug information
  -help                   Display available options
  -H                      Show header includes and nesting depth
  -idirafter <value>      Add directory to AFTER include search path
  -iframework <value>     Add directory to SYSTEM framework search path
  -imacros <file>         Include macros from file before parsing
  -include-pch <file>     Include precompiled header file
  -include <file>         Include file before parsing
  -index-header-map       Make the next included directory (-I or -F) an indexer header map
  -iprefix <dir>          Set the -iwithprefix/-iwithprefixbefore prefix
  -iquote <directory>     Add directory to QUOTE include search path
  -isysroot <dir>         Set the system root directory (usually /)
  -isystem-after <directory>
                          Add directory to end of the SYSTEM include search path
  -isystem <directory>    Add directory to SYSTEM include search path
  -ivfsoverlay <value>    Overlay the virtual filesystem described by file over the real file system
  -iwithprefixbefore <dir>
                          Set directory to include search path with prefix
  -iwithprefix <dir>      Set directory to SYSTEM include search path with prefix
  -iwithsysroot <directory>
                          Add directory to SYSTEM include search path, absolute paths are relative to -isysroot
  -I <value>              Add directory to include search path
  -mabicalls              Enable SVR4-style position-independent code (Mips only)
  -malign-double          Align doubles to two words in structs (x86 only)
  -mamdgpu-debugger-abi=<version>
                          Generate additional code for specified <version> of debugger ABI (AMDGPU only)
  -mbackchain             Link stack frames through backchain on System Z
  -mcrc                   Allow use of CRC instructions (ARM only)
  -MD                     Write a depfile containing user and system headers
  -meabi <value>          Set EABI type, e.g. 4, 5 or gnu (default depends on triple)
  -mfix-cortex-a53-835769 Workaround Cortex-A53 erratum 835769 (AArch64 only)
  -mfp32                  Use 32-bit floating point registers (MIPS only)
  -mfp64                  Use 64-bit floating point registers (MIPS only)
  -mfpxx                  Avoid FPU mode dependent operations when used with the O32 ABI
  -MF <file>              Write depfile output from -MMD, -MD, -MM, or -M to <file>
  -mgeneral-regs-only     Generate code which only uses the general purpose registers (AArch64 only)
  -mglobal-merge          Enable merging of globals
  -MG                     Add missing headers to depfile
  -mhvx-double            Enable Hexagon Double Vector eXtensions
  -mhvx                   Enable Hexagon Vector eXtensions
  -miamcu                 Use Intel MCU ABI
  --migrate               Run the migrator
  -mincremental-linker-compatible
                          (integrated-as) Emit an object file which can be used with an incremental linker
  -mios-version-min=<value>
                          Set iOS deployment target
  -mips1                  Equivalent to -march=mips1
  -mips2                  Equivalent to -march=mips2
  -mips32r2               Equivalent to -march=mips32r2
  -mips32r3               Equivalent to -march=mips32r3
  -mips32r5               Equivalent to -march=mips32r5
  -mips32r6               Equivalent to -march=mips32r6
  -mips32                 Equivalent to -march=mips32
  -mips3                  Equivalent to -march=mips3
  -mips4                  Equivalent to -march=mips4
  -mips5                  Equivalent to -march=mips5
  -mips64r2               Equivalent to -march=mips64r2
  -mips64r3               Equivalent to -march=mips64r3
  -mips64r5               Equivalent to -march=mips64r5
  -mips64r6               Equivalent to -march=mips64r6
  -mips64                 Equivalent to -march=mips64
  -MJ <value>             Write a compilation database entry per input
  -mllvm <value>          Additional arguments to forward to LLVM's option processing
  -mlong-calls            Generate branches with extended addressability, usually via indirect jumps.
  -mmacosx-version-min=<value>
                          Set Mac OS X deployment target
  -MMD                    Write a depfile containing user headers
  -mms-bitfields          Set the default structure layout to be compatible with the Microsoft compiler standard
  -mmsa                   Enable MSA ASE (MIPS only)
  -MM                     Like -MMD, but also implies -E and writes to stdout by default
  -mno-abicalls           Disable SVR4-style position-independent code (Mips only)
  -mno-fix-cortex-a53-835769
                          Don't workaround Cortex-A53 erratum 835769 (AArch64 only)
  -mno-global-merge       Disable merging of globals
  -mno-hvx-double         Disable Hexagon Double Vector eXtensions
  -mno-hvx                Disable Hexagon Vector eXtensions
  -mno-implicit-float     Don't generate implicit floating point instructions
  -mno-incremental-linker-compatible
                          (integrated-as) Emit an object file which cannot be used with an incremental linker
  -mno-long-calls         Restore the default behaviour of not generating long calls
  -mno-movt               Disallow use of movt/movw pairs (ARM only)
  -mno-ms-bitfields       Do not set the default structure layout to be compatible with the Microsoft compiler standard
  -mno-msa                Disable MSA ASE (MIPS only)
  -mno-odd-spreg          Disable odd single-precision floating point registers
  -mno-restrict-it        Allow generation of deprecated IT blocks for ARMv8. It is off by default for ARMv8 Thumb mode
  -mno-unaligned-access   Force all memory accesses to be aligned (AArch32/AArch64 only)
  -mnocrc                 Disallow use of CRC instructions (ARM only)
  -modd-spreg             Enable odd single-precision floating point registers
  -module-dependency-dir <value>
                          Directory to dump module dependencies to
  -module-file-info       Provide information about a particular module file
  -momit-leaf-frame-pointer
                          Omit frame pointer setup for leaf functions
  -mpie-copy-relocations  Use copy relocations support for PIE builds
  -MP                     Create phony target for each dependency (other than main file)
  -mqdsp6-compat          Enable hexagon-qdsp6 backward compatibility
  -MQ <value>             Specify name of main file output to quote in depfile
  -mrelax-all             (integrated-as) Relax all machine instructions
  -mrestrict-it           Disallow generation of deprecated IT blocks for ARMv8. It is on by default for ARMv8 Thumb mode.
  -mrtd                   Make StdCall calling convention the default
  -msoft-float            Use software floating point
  -mstack-alignment=<value>
                          Set the stack alignment
  -mstack-probe-size=<value>
                          Set the stack probe size
  -mstackrealign          Force realign the stack at entry to every function
  -mstrict-align          Force all memory accesses to be aligned (same as mno-unaligned-access)
  -mthread-model <value>  The thread model to use, e.g. posix, single (posix by default)
  -MT <value>             Specify name of main file output in depfile
  -munaligned-access      Allow memory accesses to be unaligned (AArch32/AArch64 only)
  -MV                     Use NMake/Jom format for the depfile
  -M                      Like -MD, but also implies -E and writes to stdout by default
  -no-canonical-prefixes  Use relative instead of canonical paths
  --no-cuda-gpu-arch=<value>
                          Remove GPU architecture (e.g. sm_35) from the list of GPUs to compile for. 'all' resets the list to its default value.
  --no-cuda-version-check Don't error out if the detected version of the CUDA install is too low for the requested CUDA gpu architecture.
  --no-system-header-prefix=<prefix>
                          Treat all #include paths starting with <prefix> as not including a system header.
  -nobuiltininc           Disable builtin #include directories
  -nostdinc++             Disable standard #include directories for the C++ standard library
  -ObjC++                 Treat source input files as Objective-C++ inputs
  -objcmt-atomic-property Make migration to 'atomic' properties
  -objcmt-migrate-all     Enable migration to modern ObjC
  -objcmt-migrate-annotation
                          Enable migration to property and method annotations
  -objcmt-migrate-designated-init
                          Enable migration to infer NS_DESIGNATED_INITIALIZER for initializer methods
  -objcmt-migrate-instancetype
                          Enable migration to infer instancetype for method result type
  -objcmt-migrate-literals
                          Enable migration to modern ObjC literals
  -objcmt-migrate-ns-macros
                          Enable migration to NS_ENUM/NS_OPTIONS macros
  -objcmt-migrate-property-dot-syntax
                          Enable migration of setter/getter messages to property-dot syntax
  -objcmt-migrate-property
                          Enable migration to modern ObjC property
  -objcmt-migrate-protocol-conformance
                          Enable migration to add protocol conformance on classes
  -objcmt-migrate-readonly-property
                          Enable migration to modern ObjC readonly property
  -objcmt-migrate-readwrite-property
                          Enable migration to modern ObjC readwrite property
  -objcmt-migrate-subscripting
                          Enable migration to modern ObjC subscripting
  -objcmt-ns-nonatomic-iosonly
                          Enable migration to use NS_NONATOMIC_IOSONLY macro for setting property's 'atomic' attribute
  -objcmt-returns-innerpointer-property
                          Enable migration to annotate property with NS_RETURNS_INNER_POINTER
  -objcmt-whitelist-dir-path=<value>
                          Only modify files with a filename contained in the provided directory path
  -ObjC                   Treat source input files as Objective-C inputs
  -o <file>               Write output to <file>
  -pg                     Enable mcount instrumentation
  -pipe                   Use pipes between commands, when possible
  --precompile            Only precompile the input
  -print-file-name=<file> Print the full library path of <file>
  -print-ivar-layout      Enable Objective-C Ivar layout bitmap print trace
  -print-libgcc-file-name Print the library path for the currently used compiler runtime library ("libgcc.a" or "libclang_rt.builtins.*.a")
  -print-prog-name=<name> Print the full program path of <name>
  -print-search-dirs      Print the paths used for finding libraries and programs
  -pthread                Support POSIX threads in generated code
  -P                      Disable linemarker output in -E mode
  -Qunused-arguments      Don't emit warning for unused driver arguments
  -relocatable-pch        Whether to build a relocatable precompiled header
  -resource-dir <value>   The directory which holds the compiler resource files
  -rewrite-legacy-objc    Rewrite Legacy Objective-C source to C++
  -rewrite-objc           Rewrite Objective-C source to C++
  -Rpass-analysis=<value> Report transformation analysis from optimization passes whose name matches the given POSIX regular expression
  -Rpass-missed=<value>   Report missed transformations by optimization passes whose name matches the given POSIX regular expression
  -Rpass=<value>          Report transformations performed by optimization passes whose name matches the given POSIX regular expression
  -rtlib=<value>          Compiler runtime library to use
  -R<remark>              Enable the specified remark
  -save-stats=<value>     Save llvm statistics.
  -save-stats             Save llvm statistics.
  -save-temps=<value>     Save intermediate compilation results.
  -save-temps             Save intermediate compilation results
  -serialize-diagnostics <value>
                          Serialize compiler diagnostics to a file
  -std=<value>            Language standard to compile for
  -stdlib=<value>         C++ standard library to use
  --system-header-prefix=<prefix> Treat all #include paths starting with <prefix> as including a system header.
  -S                      Only run preprocess and compilation steps
  --target=<value>        Generate code for the given target
  -time                   Time individual commands
  -traditional-cpp        Enable some traditional CPP emulation
  -trigraphs              Process trigraph sequences
  -undef                  undef all system defines
  --verify-debug-info     Verify the binary representation of debug output
  -verify-pch             Load and verify that a pre-compiled header file is not stale
  -v                      Show commands to run and use verbose output
  -Wa,<arg>               Pass the comma separated arguments in <arg> to the assembler
  -Wl,<arg>               Pass the comma separated arguments in <arg> to the linker
  -Wlarge-by-value-copy   Warn if a function definition returns or accepts an object larger in bytes than a given value
  -working-directory <value> Resolve file paths relative to the specified directory
  -Wp,<arg>               Pass the comma separated arguments in <arg> to the preprocessor
  -W<warning>             Enable the specified warning
  -w                      Suppress all warnings
  -Xanalyzer <arg>        Pass <arg> to the static analyzer
  -Xassembler <arg>       Pass <arg> to the assembler
  -Xclang <arg>           Pass <arg> to the clang compiler
  -Xcuda-fatbinary <arg>  Pass <arg> to fatbinary invocation
  -Xcuda-ptxas <arg>      Pass <arg> to the ptxas assembler
  -Xlinker <arg>          Pass <arg> to the linker
  -Xpreprocessor <arg>    Pass <arg> to the preprocessor
  -x <language>           Treat subsequent input files as having type <language>
  -z <arg>                Pass -z <arg> to the linker
```

man page
--------

```roff
CLANG(1)                                                                                                                           Clang                                                                                                                          CLANG(1)



NAME
       clang - the Clang C, C++, and Objective-C compiler

SYNOPSIS
       clang [options] filename ...

DESCRIPTION
       clang is a C, C++, and Objective-C compiler which encompasses preprocessing, parsing, optimization, code generation, assembly, and linking.  Depending on which high-level mode setting is passed, Clang will stop before doing a full link.  While Clang is highly
       integrated, it is important to understand the stages of compilation, to understand how to invoke it.  These stages are:

       Driver The clang executable is actually a small driver which controls the overall execution of other tools such as the compiler, assembler and linker.  Typically you do not need to interact with the driver, but you transparently use it to run the other tools.

       Preprocessing
              This  stage  handles  tokenization  of  the input source file, macro expansion, #include expansion and handling of other preprocessor directives.  The output of this stage is typically called a ".i" (for C), ".ii" (for C++), ".mi" (for Objective-C), or
              ".mii" (for Objective-C++) file.

       Parsing and Semantic Analysis
              This stage parses the input file, translating preprocessor tokens into a parse tree.  Once in the form of a parse tree, it applies semantic analysis to compute types for expressions as well and determine whether the code is well formed. This  stage  is
              responsible for generating most of the compiler warnings as well as parse errors. The output of this stage is an "Abstract Syntax Tree" (AST).

       Code Generation and Optimization
              This  stage  translates an AST into low-level intermediate code (known as "LLVM IR") and ultimately to machine code.  This phase is responsible for optimizing the generated code and handling target-specific code generation.  The output of this stage is
              typically called a ".s" file or "assembly" file.

              Clang also supports the use of an integrated assembler, in which the code generator produces object files directly. This avoids the overhead of generating the ".s" file and of calling the target assembler.

       Assembler
              This stage runs the target assembler to translate the output of the compiler into a target object file. The output of this stage is typically called a ".o" file or "object" file.

       Linker This stage runs the target linker to merge multiple object files into an executable or dynamic library. The output of this stage is typically called an "a.out", ".dylib" or ".so" file.

       Clang Static Analyzer

       The Clang Static Analyzer is a tool that scans source code to try to find bugs through code analysis.  This tool uses many parts of Clang and is built into the same driver.  Please see <http://clang-analyzer.llvm.org> for more details on how to use the static
       analyzer.

OPTIONS
   Stage Selection Options
       -E     Run the preprocessor stage.

       -fsyntax-only
              Run the preprocessor, parser and type checking stages.

       -S     Run the previous stages as well as LLVM generation and optimization stages and target-specific code generation, producing an assembly file.

       -c     Run all of the above, plus the assembler, generating a target ".o" object file.

       no stage selection option
              If no stage selection option is specified, all stages above are run, and the linker is run to combine the results into an executable or shared library.

   Language Selection and Mode Options
       -x <language>
              Treat subsequent input files as having type language.

       -std=<language>
              Specify the language standard to compile for.

       -stdlib=<library>
              Specify the C++ standard library to use; supported options are libstdc++ and libc++. If not specified, platform default will be used.

       -rtlib=<library>
              Specify the compiler runtime library to use; supported options are libgcc and compiler-rt. If not specified, platform default will be used.

       -ansi  Same as -std=c89.

       -ObjC  Treat source input files as Objective-C and Object-C++ inputs respectively.

       -trigraphs
              Enable trigraphs.

       -ffreestanding
              Indicate that the file should be compiled for a freestanding, not a hosted, environment.

       -fno-builtin
              Disable special handling and optimizations of builtin functions like strlen() and malloc().

       -fmath-errno
              Indicate that math functions should be treated as updating errno.

       -fpascal-strings
              Enable support for Pascal-style strings with "\pfoo".

       -fms-extensions
              Enable support for Microsoft extensions.

       -fmsc-version=
              Set _MSC_VER. Defaults to 1300 on Windows. Not set otherwise.

       -fborland-extensions
              Enable support for Borland extensions.

       -fwritable-strings
              Make all string literals default to writable.  This disables uniquing of strings and other optimizations.

       -flax-vector-conversions
              Allow loose type checking rules for implicit vector conversions.

       -fblocks
              Enable the "Blocks" language feature.

       -fobjc-gc-only
              Indicate that Objective-C code should be compiled in GC-only mode, which only works when Objective-C Garbage Collection is enabled.

       -fobjc-gc
              Indicate that Objective-C code should be compiled in hybrid-GC mode, which works with both GC and non-GC mode.

       -fobjc-abi-version=version
              Select the Objective-C ABI version to use. Available versions are 1 (legacy "fragile" ABI), 2 (non-fragile ABI 1), and 3 (non-fragile ABI 2).

       -fobjc-nonfragile-abi-version=<version>
              Select the Objective-C non-fragile ABI version to use by default. This will only be used as the Objective-C ABI when the non-fragile ABI is enabled (either via -fobjc-nonfragile-abi, or because it is the platform default).

       -fobjc-nonfragile-abi, -fno-objc-nonfragile-abi
              Enable use of the Objective-C non-fragile ABI. On platforms for which this is the default ABI, it can be disabled with -fno-objc-nonfragile-abi.

   Target Selection Options
       Clang fully supports cross compilation as an inherent part of its design.  Depending on how your version of Clang is configured, it may have support for a number of cross compilers, or may only support a native target.

       -arch <architecture>
              Specify the architecture to build for.

       -mmacosx-version-min=<version>
              When building for Mac OS X, specify the minimum version supported by your application.

       -miphoneos-version-min
              When building for iPhone OS, specify the minimum version supported by your application.

       -march=<cpu>
              Specify that Clang should generate code for a specific processor family member and later.  For example, if you specify -march=i486, the compiler is allowed to generate instructions that are valid on i486 and later processors, but which may not exist on
              earlier ones.

   Code Generation Options
       -O0, -O1, -O2, -O3, -Ofast, -Os, -Oz, -Og, -O, -O4
              Specify which optimization level to use:
                 -O0 Means "no optimization": this level compiles the fastest and generates the most debuggable code.

                 -O1 Somewhere between -O0 and -O2.

                 -O2 Moderate level of optimization which enables most optimizations.

                 -O3 Like -O2, except that it enables optimizations that take longer to perform or that may generate larger code (in an attempt to make the program run faster).

                 -Ofast Enables all the optimizations from -O3 along with other aggressive optimizations that may violate strict compliance with language standards.

                 -Os Like -O2 with extra optimizations to reduce code size.

                 -Oz Like -Os (and thus -O2), but reduces code size further.

                 -Og Like -O1. In future versions, this option might disable different optimizations in order to improve debuggability.

                 -O Equivalent to -O2.

                 -O4 and higher
                     Currently equivalent to -O3

       -g, -gline-tables-only, -gmodules
              Control debug information output.  Note that Clang debug information works best at -O0.  When more than one option starting with -g is specified, the last one wins:
                 -g Generate debug information.

                 -gline-tables-only Generate only line table debug information. This allows for symbolicated backtraces with inlining information, but does not include any information about variables, their locations or types.

                 -gmodules Generate debug information that contains external references to types defined in Clang modules or precompiled headers instead of emitting redundant debug type information into every object file.   This  option  transparently  switches  the
                 Clang  module format to object file containers that hold the Clang module together with the debug information.  When compiling a program that uses Clang modules or precompiled headers, this option produces complete debug information with faster com-
                 pile times and much smaller object files.

                 This option should not be used when building static libraries for distribution to other machines because the debug info will contain references to the module cache on the machine the object files in the library were built on.

       -fstandalone-debug -fno-standalone-debug
              Clang supports a number of optimizations to reduce the size of debug information in the binary. They work based on the assumption that the debug type information can be spread out over multiple compilation units.  For instance, Clang will not emit type
              definitions for types that are not needed by a module and could be replaced with a forward declaration.  Further, Clang will only emit type info for a dynamic C++ class in the module that contains the vtable for the class.

              The  -fstandalone-debug option turns off these optimizations.  This is useful when working with 3rd-party libraries that don't come with debug information.  This is the default on Darwin.  Note that Clang will never emit type information for types that
              are not referenced at all by the program.

       -fexceptions
              Enable generation of unwind information. This allows exceptions to be thrown through Clang compiled stack frames.  This is on by default in x86-64.

       -ftrapv
              Generate code to catch integer overflow errors.  Signed integer overflow is undefined in C. With this flag, extra code is generated to detect this and abort when it happens.

       -fvisibility
              This flag sets the default visibility level.

       -fcommon, -fno-common
              This flag specifies that variables without initializers get common linkage.  It can be disabled with -fno-common.

       -ftls-model=<model>
              Set the default thread-local storage (TLS) model to use for thread-local variables. Valid values are: "global-dynamic", "local-dynamic", "initial-exec" and "local-exec". The default is "global-dynamic". The default model  can  be  overridden  with  the
              tls_model attribute. The compiler will try to choose a more efficient model if possible.

       -flto, -flto=full, -flto=thin, -emit-llvm
              Generate  output  files  in  LLVM  formats,  suitable  for link time optimization.  When used with -S this generates LLVM intermediate language assembly files, otherwise this generates LLVM bitcode format object files (which may be passed to the linker
              depending on the stage selection options).

              The default for -flto is "full", in which the LLVM bitcode is suitable for monolithic Link Time Optimization (LTO), where the linker merges all such modules into a single combined module for optimization. With "thin",  ThinLTO  compilation  is  invoked
              instead.

   Driver Options
       -###   Print (but do not run) the commands to run for this compilation.

       --help Display available options.

       -Qunused-arguments
              Do not emit any warnings for unused driver arguments.

       -Wa,<args>
              Pass the comma separated arguments in args to the assembler.

       -Wl,<args>
              Pass the comma separated arguments in args to the linker.

       -Wp,<args>
              Pass the comma separated arguments in args to the preprocessor.

       -Xanalyzer <arg>
              Pass arg to the static analyzer.

       -Xassembler <arg>
              Pass arg to the assembler.

       -Xlinker <arg>
              Pass arg to the linker.

       -Xpreprocessor <arg>
              Pass arg to the preprocessor.

       -o <file>
              Write output to file.

       -print-file-name=<file>
              Print the full library path of file.

       -print-libgcc-file-name
              Print the library path for the currently used compiler runtime library ("libgcc.a" or "libclang_rt.builtins.*.a").

       -print-prog-name=<name>
              Print the full program path of name.

       -print-search-dirs
              Print the paths used for finding libraries and programs.

       -save-temps
              Save intermediate compilation results.

       -save-stats, -save-stats=cwd, -save-stats=obj
              Save internal code generation (LLVM) statistics to a file in the current directory (-save-stats/"-save-stats=cwd") or the directory of the output file ("-save-state=obj").

       -integrated-as, -no-integrated-as
              Used to enable and disable, respectively, the use of the integrated assembler. Whether the integrated assembler is on by default is target dependent.

       -time  Time individual commands.

       -ftime-report
              Print timing summary of each stage of compilation.

       -v     Show commands to run and use verbose output.

   Diagnostics Options
       -fshow-column, -fshow-source-location, -fcaret-diagnostics, -fdiagnostics-fixit-info, -fdiagnostics-parseable-fixits, -fdiagnostics-print-source-range-info, -fprint-source-range-info, -fdiagnostics-show-option, -fmessage-length
              These options control how Clang prints out information about diagnostics (errors and warnings). Please see the Clang User's Manual for more information.

   Preprocessor Options
       -D<macroname>=<value>
              Adds an implicit #define into the predefines buffer which is read before the source file is preprocessed.

       -U<macroname>
              Adds an implicit #undef into the predefines buffer which is read before the source file is preprocessed.

       -include <filename>
              Adds an implicit #include into the predefines buffer which is read before the source file is preprocessed.

       -I<directory>
              Add the specified directory to the search path for include files.

       -F<directory>
              Add the specified directory to the search path for framework include files.

       -nostdinc
              Do not search the standard system directories or compiler builtin directories for include files.

       -nostdlibinc
              Do not search the standard system directories for include files, but do search compiler builtin include directories.

       -nobuiltininc
              Do not search clang's builtin directory for include files.

ENVIRONMENT
       TMPDIR, TEMP, TMP
              These environment variables are checked, in order, for the location to write temporary files used during the compilation process.

       CPATH  If this environment variable is present, it is treated as a delimited list of paths to be added to the default system include path list. The delimiter is the platform dependent delimiter, as used in the PATH environment variable.

              Empty components in the environment variable are ignored.

       C_INCLUDE_PATH, OBJC_INCLUDE_PATH, CPLUS_INCLUDE_PATH, OBJCPLUS_INCLUDE_PATH
              These environment variables specify additional paths, as for CPATH, which are only used when processing the appropriate language.

       MACOSX_DEPLOYMENT_TARGET
              If -mmacosx-version-min is unspecified, the default deployment target is read from this environment variable. This option only affects Darwin targets.

BUGS
       To report bugs, please visit <http://llvm.org/bugs/>.  Most bug reports should include preprocessed source files (use the -E option) and the full output of the compiler, along with information to reproduce.

SEE ALSO
       as(1), ld(1)

AUTHOR
       Maintained by the Clang / LLVM Team (<http://clang.llvm.org>)

COPYRIGHT
       2007-2016, The Clang Team



4.0                                                                                                                            Dec 10, 2016                                                                                                                       CLANG(1)
```
