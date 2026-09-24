blastradius is a Go CLI that takes a leaked credential, verifies it is live with read-only checks, enriches it with impact context (repo reach, package dependents, production signals), and outputs a severity rating with evidence. Deterministic rules are the baseline; the AI classifier is an optional overlay and the tool must work offline without it.

<!-- BEGIN:agent-rules -->
## Working in this repo

- Provider clients are read-only: never create, modify, delete, or publish; do not expose write-capable SDK methods
- The secret never appears in logs, errors, reports, caches, test fixtures, or anything sent to the AI classifier; only derived evidence leaves the process
- Make the smallest correct change; match the structure of neighbouring files; look for an existing helper before writing one
- Favor the simplest mechanism that works: no new abstraction, interface, option, or dependency without a concrete second use
- Keep the diff to one concern; leave unrelated code, comments, and formatting alone
- Comments say why, not what, in one plain sentence; a paragraph-long comment means the code needs simplifying
- Doc comments are one or two sentences stating the contract; no narration of the implementation, no restating the signature
- Fix the cause rather than disabling a linter or weakening a test; if unavoidable, narrowest scope plus a one-line reason
- Bug fixes include a regression test that fails without the fix; extend an existing table before adding a new test
- Wait on a condition, never sleep, in tests

## Go

- `gofmt`, `go vet ./...`, and `golangci-lint run` must pass; one test: `go test -run '^TestName$' ./internal/...`
- Stdlib first (`slices`, `maps`, `cmp`, `log/slog`); official provider SDKs are fine
- Every network call takes a `context.Context` with a timeout
- Wrap errors with `fmt.Errorf("doing x: %w", err)` and handle each once; never log and return
- Every goroutine has a known exit; prefer synchronous APIs and let the caller go concurrent
- Table-driven tests with fake credentials of the correct shape
<!-- END:agent-rules -->
