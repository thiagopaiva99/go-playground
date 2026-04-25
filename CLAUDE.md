# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Structure

Go learning playground. Each top-level directory is an independent, standalone program — no shared module except `deck-cards/` which has its own `go.mod` (`module cards`).

Most directories contain a single `.go` file with `package main`. No workspace-level `go.mod` exists.

## Commands

Run a specific program:
```sh
go run <dir>/<file>.go
# e.g. go run hello-world/hello.go
```

Run a program that takes args (e.g. `hard-mode-interfaces`):
```sh
go run assignments/hard-mode-interfaces/main.go assignments/hard-mode-interfaces/file.txt
```

Run tests (only `deck-cards` has tests):
```sh
cd deck-cards && go test ./...
```

Run single test:
```sh
cd deck-cards && go test -run TestNewDeck
```

Build a binary:
```sh
go build -o out <dir>/<file>.go
```

## Conventions

- Each exercise is self-contained; add new topics as new top-level directories.
- Assignments go under `assignments/<name>/`.
- `deck-cards` is the only multi-file package; new packages with tests should follow its pattern (own `go.mod`, `package main`, `*_test.go`).

### Test
this is some testing section