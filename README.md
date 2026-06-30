# yze-go-ptrrecv

A [`yze`](https://github.com/gomatic/yze) analyzer (category `immutability`) enforcing the gomatic Go immutability standard: methods use **value receivers**, never pointer receivers, unless the receiver type transitively contains a field that cannot be copied (a `sync` primitive, `sync/atomic` type, `bytes.Buffer`, or `strings.Builder`).

Pointer receivers imply mutability — and on exported types that leaks into interface semantics and concurrency expectations. This analyzer flags every pointer-receiver method whose type holds no no-copy field.

The rule is intentionally narrow: only a transitively-uncopyable field justifies a pointer receiver. Mutation and interface-satisfaction are deliberately **not** carve-outs — that narrowness reflects the gomatic immutability preference. The no-copy walk recurses through nested structs and through array elements (an array stores its elements inline, so a no-copy element makes the array uncopyable), but not through slices, maps, channels, or pointers, which are references that leave the enclosing struct copyable.

- **Rule:** `yze/ptrrecv`
- **No-copy types:** 16 baked-in standard-library types, extensible at runtime via the `-allow` flag (comma-separated fully-qualified `pkgpath.Name` entries).
- **Binary:** `cmd/yze-go-ptrrecv` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
