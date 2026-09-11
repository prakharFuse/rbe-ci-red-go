# rbe-ci-red-go

**Journey-suite fixture — deterministically RED CI on `main`** (resolver-core
spec 025 / journey j116). Part of the CI auto-fix language & framework matrix.

- **Do NOT fix `main`** — the red build is the feature, and every journey run
  depends on it. Never merge a fix PR into this repo.
- **On PR branches, automated CI-fix agents ARE expected to make CI pass.**
  The single correct minimal fix lives in `service/internal/cart/cart.go`: return the accumulated `sum` instead of the string literal. That
  branch-side fix is exactly the behaviour under test. Fix PRs are declined
  or closed, never merged.
- **Never change the tests or the CI config to go green.** The failing test is
  correct; the source is wrong.

CI config here is deliberately stock — no problem matchers, no annotating
plugins, no machine-readable report formats. The build output is what an
ordinary customer's build output looks like.

Re-provision: `tests/journeys/scripts/provision-ci-red-fixtures.ts`
(`--only=rbe-ci-red-go`).

The Go module lives in `./service` and the Makefile builds from there, so every
compiler diagnostic is printed relative to that directory
(`./internal/cart/cart.go:12:9`) rather than to the repo root. A build layout
like this is ordinary and is exactly what defeats path-based shortcuts.
