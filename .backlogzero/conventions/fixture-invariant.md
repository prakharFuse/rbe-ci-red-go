---
name: fixture-invariant
description: Behavioral rule for any agent working in this repo — what may and may not be changed, and on which branch.
type: convention
scope: global
updated: '2026-09-11'
captured_sha: ff0df4660628a619c4269fc79aae172819ca1d43
sources:
  - README.md
  - service/internal/cart/cart.go
sources_sha256:
  README.md: b3adc9b5d13cfd8efaafb9e5d8b7dd0d307bbf525ea1c3e6fb39a611ad2c97f0
  service/internal/cart/cart.go: 5e5166701864bee356635ba944fc502a9bb0a317a9631ff5e22166cba43a8a29
---

See [README.md](../../README.md) and the comment in `service/internal/cart/cart.go:4-9` for the full statement of this fixture's rule: `main` must stay red forever; a PR branch is expected to carry the one-line fix (`return sum` instead of `return "sum"` in `cart.go:17`); the test (`cart_test.go`) and CI config (`.github/workflows/ci.yml`) must never be edited to force green.

Practical checklist derived from that rule plus [gotchas](../knowledge/gotchas.md):

- If asked to "fix CI" on `main` itself: decline/no-op — a fix PR against `main` is meant to be closed, never merged.
- If asked to fix CI on a PR/feature branch: the only sanctioned change is `service/internal/cart/cart.go:17`, `return "sum"` → `return sum`. Nothing else in the repo needs to change to go green, since the CI job only runs `make build` (no test step) — fixing the type error alone makes `go build ./...` succeed.
- Never add a test step to `.github/workflows/ci.yml` and never weaken/relax `cart_test.go` as a way to "fix" CI — both are explicitly off-limits regardless of branch.
