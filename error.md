```sh
bzlmod-llvm-go-macos % bazel build //go/app
INFO: Analyzed target //go/app:app (0 packages loaded, 0 targets configured).
ERROR: /private/var/tmp/_bazel_aran/efc359ce909c32eae874b66e23a9926d/external/zlib~1.3/BUILD.bazel:72:11: Linking external/zlib~1.3/libzlib.a [for tool] failed: (Exit 1): llvm-ar failed: error executing CppArchive command (from target @@zlib~1.3//:zlib) external/toolchains_llvm~0.10.3~llvm~llvm_toolchain/bin/llvm-ar @bazel-out/darwin_arm64-opt-exec-ST-13d3ddad9198/bin/external/zlib~1.3/libzlib.a-2.params

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
external/toolchains_llvm~0.10.3~llvm~llvm_toolchain/bin/llvm-ar: error: only one operation may be specified
OVERVIEW: LLVM Archiver

USAGE: llvm-ar [options] [-]<operation>[modifiers] [relpos] [count] <archive> [files]
       llvm-ar -M [<mri-script]

OPTIONS:
  --format              - archive format to create
    =default            -   default
    =gnu                -   gnu
    =darwin             -   darwin
    =bsd                -   bsd
    =bigarchive         -   big archive (AIX OS)
  --plugin=<string>     - ignored for compatibility
  -h --help             - display this help and exit
  --output              - the directory to extract archive members to
  --rsp-quoting         - quoting style for response files
    =posix              -   posix
    =windows            -   windows
  --thin                - create a thin archive
  --version             - print the version and exit
  -X{32|64|32_64|any}   - object mode (only for AIX OS)
  @<file>               - read options from <file>

OPERATIONS:
  d - delete [files] from the archive
  m - move [files] in the archive
  p - print contents of [files] found in the archive
  q - quick append [files] to the archive
  r - replace or insert [files] into the archive
  s - act as ranlib
  t - display list of files in archive
  x - extract [files] from the archive

MODIFIERS:
  [a] - put [files] after [relpos]
  [b] - put [files] before [relpos] (same as [i])
  [c] - do not warn if archive had to be created
  [D] - use zero for timestamps and uids/gids (default)
  [h] - display this help and exit
  [i] - put [files] before [relpos] (same as [b])
  [l] - ignored for compatibility
  [L] - add archive's contents
  [N] - use instance [count] of name
  [o] - preserve original dates
  [O] - display member offsets
  [P] - use full names when matching (implied for thin archives)
  [s] - create an archive index (cf. ranlib)
  [S] - do not build a symbol table
  [T] - deprecated, use --thin instead
  [u] - update only [files] newer than archive contents
  [U] - use actual timestamps and uids/gids
  [v] - be verbose about actions taken
  [V] - display the version and exit
Target //go/app:app failed to build
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 0.253s, Critical Path: 0.05s
INFO: 8 processes: 8 internal.
ERROR: Build did NOT complete successfully
```