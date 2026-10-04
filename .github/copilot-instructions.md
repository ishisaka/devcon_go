# Copilot instructions

## Build, run, test, and lint

- The module requires Go 1.27 (`go.mod`); `mise.toml` and `_Dockerfile` pin the toolchain to Go 1.27.1.
- Run the application from the repository root with `go run ./cmd`.
- Build the executable package with `go build ./cmd`, or verify every package builds with `go build ./...`.
- Run all tests with `go test ./...`.
- Run one named test with `go test ./pkg/hello -run '^TestName$' -count=1`; replace the package and anchored test name as needed. There are currently no `_test.go` files.
- Format all packages with `go fmt ./...`, run standard static checks with `go vet ./...`, and run the configured linter with `golangci-lint run`.

## Architecture

This is a single-module Go command-line application:

- `cmd/main.go` is the executable entry point. It should remain thin: compose package behavior and handle terminal output here.
- Reusable behavior lives under `pkg`. Currently `pkg/hello` owns construction of the greeting, while `cmd` decides how to present it.
- The module path is `devcon_go`, so internal imports use paths such as `devcon_go/pkg/hello`.

The VS Code dev container (`.devcontainer/devcontainer.json`) installs the `mise` feature, which reads the Go 1.27.1 toolchain from `mise.toml`; it does not reference `_Dockerfile`. Treat `_Dockerfile` as a separate Go-based development image unless the container configuration is deliberately changed.

## Repository conventions

- Keep user-facing output at the command boundary: package functions return values, and `cmd/main.go` prints them.
- Use one package directory per reusable concern beneath `pkg`, and import it from `cmd` rather than putting application logic directly in `main`.
- Keep Go source formatted with `go fmt`; use standard Go package and test naming (`package hello`, `hello_test.go`, `TestXxx`).
