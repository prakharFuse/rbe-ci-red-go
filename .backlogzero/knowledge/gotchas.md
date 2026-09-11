---
name: gotchas
description: Non-obvious facts about why CI is red here — read before touching cart.go or ci.yml.
type: knowledge
scope: global
updated: '2026-09-11'
captured_sha: ff0df4660628a619c4269fc79aae172819ca1d43
sources:
  - service/internal/cart/cart.go
  - service/internal/cart/cart_test.go
  - .github/workflows/ci.yml
  - README.md
sources_sha256:
  .github/workflows/ci.yml: de0737eb46042dfb879afdd6266c817b5a76e8c64ee081ec99b5779af3d49699
  README.md: b3adc9b5d13cfd8efaafb9e5d8b7dd0d307bbf525ea1c3e6fb39a611ad2c97f0
  service/internal/cart/cart.go: 5e5166701864bee356635ba944fc502a9bb0a317a9631ff5e22166cba43a8a29
  service/internal/cart/cart_test.go: b5f9e9bc65d2e5108f9959517198b1e7b5c1b81f5eefcbeb95b3053f02173b93
---

`Total` is declared to return `int` but its final statement is `return "sum"` (`service/internal/cart/cart.go:12,17`). Because Go is statically typed, this is a **compile-time type error** ("cannot use \"sum\" (untyped string constant) as int value"), not a runtime bug and not a test assertion failure. `go build ./...` fails before any test binary can even be produced — `go test ./...` would fail for the identical compile reason, never reaching `TestTotal`'s assertion in `cart_test.go:6`.

Diverges from README.md: the README (and the comment block in `cart.go:4-9`) both describe the fixture in terms of "the failing test is correct; the source is wrong." In reality the red signal on `main` is a **failing build** (compile error), not a failing test — and CI here (`.github/workflows/ci.yml:18`) only runs `make build`, never `make test`, so `TestTotal` never actually executes in CI. The behavioral rule is the same either way (fix `cart.go:17` to `return sum`, never touch the test or CI config), but agents reasoning about "why did the test fail" should know no test run actually happened — the build itself never got that far.

The compiler diagnostic for this error is emitted relative to `service/` (e.g. `./internal/cart/cart.go:12:9`), not the repo root, because `Makefile` does `cd service` first — see `README.md:22-25`. Don't assume repo-root-relative paths when parsing build output.
