---
name: overview
description: What rbe-ci-red-go is and how the repo is laid out — read this first.
type: knowledge
scope: global
updated: '2026-09-11'
captured_sha: ff0df4660628a619c4269fc79aae172819ca1d43
sources:
  - README.md
  - Makefile
  - service/go.mod
sources_sha256:
  Makefile: 798f96ba586d243b33b406f861f04c02dfcb6644f7819ac99c262b69acf04e23
  README.md: b3adc9b5d13cfd8efaafb9e5d8b7dd0d307bbf525ea1c3e6fb39a611ad2c97f0
  service/go.mod: 5c29c7b17b9d71486bab5cabe17a557653d535009b5663aa032e11b7a277114e
---

For the fixture's purpose and the rules around fixing it, see [README.md](../../README.md) — do not restate those here, just follow them.

Structural facts not spelled out in the README:

- The Go module is `example.com/rbecired` (`service/go.mod:1`), not named after the repo. There is exactly one package in the module: `cart` (`service/internal/cart/`).
- The `Makefile` has two targets, both of which `cd service` first: `build` runs `go build ./...`, `test` runs `go test ./...` (`Makefile:1-7`). Any new package must live under `service/` to be picked up by either target.
- The whole repo is 3 non-git files: `Makefile`, `.github/workflows/ci.yml`, `README.md`, plus the `service/` module (`go.mod`, `internal/cart/cart.go`, `internal/cart/cart_test.go`). There is no `go.sum` even though CI's `setup-go` step references `cache-dependency-path: service/go.sum` (`.github/workflows/ci.yml:17`) — harmless since the module has no dependencies to lock.
