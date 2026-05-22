# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added
- ANSI status colors with `NO_COLOR` and `--no-color` support
- `--sni` flag to override server name in TLS handshake
- Result sorting by urgency (errors and crit first)
- Distroless Dockerfile
- `examples/hosts.txt` sample
- Makefile with `build`, `test`, `vet`, `install` targets

## [0.3.0] - 2026-05-13

### Added
- `--quiet` flag for exit-code-only mode
- `--workers` and `--timeout` flags
- Read hosts from stdin (skip blanks and `#` comments)
- Exit codes: 0 (ok) / 1 (warn) / 2 (crit) / 3 (error)

## [0.2.0] - 2026-05-10

### Added
- `--warn-days` / `--crit-days` thresholds and STATUS column
- `--json` output mode

## [0.1.0] - 2026-05-04

### Added
- Initial CLI scaffolding
- TLS dial and leaf cert extraction (NotAfter, Issuer, Subject)
- Tabwriter aligned table output
