---
name: architecture
description: How CI, the Makefile, and the single Go package fit together.
type: knowledge
scope: global
updated: '2026-09-11'
captured_sha: ff0df4660628a619c4269fc79aae172819ca1d43
sources:
  - .github/workflows/ci.yml
  - Makefile
  - service/go.mod
  - service/internal/cart/cart.go
  - service/internal/cart/cart_test.go
sources_sha256:
  .github/workflows/ci.yml: de0737eb46042dfb879afdd6266c817b5a76e8c64ee081ec99b5779af3d49699
  Makefile: 798f96ba586d243b33b406f861f04c02dfcb6644f7819ac99c262b69acf04e23
  service/go.mod: 5c29c7b17b9d71486bab5cabe17a557653d535009b5663aa032e11b7a277114e
  service/internal/cart/cart.go: 5e5166701864bee356635ba944fc502a9bb0a317a9631ff5e22166cba43a8a29
  service/internal/cart/cart_test.go: b5f9e9bc65d2e5108f9959517198b1e7b5c1b81f5eefcbeb95b3053f02173b93
---

```mermaid
flowchart LR
    A[push / pull_request] --> B["ci.yml: build job"]
    B --> C["make build"]
    C --> D["cd service && go build ./..."]
    D --> E["package cart\ncart.go: Total()"]
    F["make test"] -.not invoked by CI.-> G["cd service && go test ./..."]
    G --> H["cart_test.go: TestTotal"]
    H --> E
```

Everything lives in one Go module (`service/`, module path `example.com/rbecired`) with a single package `cart`. There's no service boundary, network call, or database here — the "architecture" is just: GitHub Actions → Makefile → `go build`/`go test` → the `cart` package.

The `test` Makefile target exists but the CI workflow only ever runs `make build` (`.github/workflows/ci.yml:18`) — see [gotchas](gotchas.md) for why that matters.
