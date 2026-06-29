# yze-go-ptrrecv

A [`yze`](https://github.com/gomatic/yze) analyzer (group `go`, category `immutability`) enforcing the gomatic Go immutability standard: methods use **value receivers**, never pointer receivers, unless the receiver type transitively contains a field that cannot be copied (a `sync` primitive, `sync/atomic` type, `bytes.Buffer`, or `strings.Builder`).

Pointer receivers imply mutability — and on exported types that leaks into interface semantics and concurrency expectations. This analyzer flags every pointer-receiver method whose type holds no no-copy field.

- **Rule:** `yze/go/ptrrecv`
- **Binary:** `cmd/yze-go-ptrrecv` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework. The no-copy allow-list is baked in for v1; a configurable `-allow` flag is planned.
