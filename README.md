# certcheck

[![ci](https://github.com/14ten/certcheck/actions/workflows/ci.yml/badge.svg)](https://github.com/14ten/certcheck/actions/workflows/ci.yml)
[![go report](https://goreportcard.com/badge/github.com/14ten/certcheck)](https://goreportcard.com/report/github.com/14ten/certcheck)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A small Go CLI that checks TLS certificate expiry across a list of hosts.
Useful for daily cron jobs, monitoring pipelines, or a quick sanity check.

## Install

```sh
# from source
go install github.com/14ten/certcheck@latest

# docker
docker run --rm ghcr.io/14ten/certcheck:latest example.com
```

Pre-built binaries (linux / macOS / windows on amd64 + arm64) are attached to
each [release](https://github.com/14ten/certcheck/releases).

## Usage

```sh
# Arguments
certcheck example.com github.com:443 cloudflare.com

# Stdin (one host per line, # comments allowed)
cat hosts.txt | certcheck --json

# Tighter thresholds
certcheck --warn-days 14 --crit-days 3 example.com
```

## Output

```
HOST          EXPIRES     DAYS  STATUS  ISSUER
example.com   2026-08-12  82    OK      DigiCert TLS RSA SHA256 2020 CA1
api.test      2026-06-04  13    WARN    Let's Encrypt R3
expired.test  -           -     ERR     tls handshake failed
```

## Flags

| Flag           | Default | Description                                     |
|----------------|---------|-------------------------------------------------|
| `--warn-days`  | `30`    | Warn when cert expires within N days            |
| `--crit-days`  | `7`     | Critical when cert expires within N days        |
| `--json`       | `false` | Emit JSON instead of a table                    |
| `--workers`    | `8`     | Concurrent checks                               |
| `--timeout`    | `5s`    | Per-host TLS dial timeout                       |

## Exit codes

| Code | Meaning             |
|------|---------------------|
| 0    | all certs OK        |
| 1    | at least one WARN   |
| 2    | at least one CRIT   |
| 3    | a check errored     |

Exit codes are designed to plug straight into Nagios-style monitoring.

## License

MIT — see [LICENSE](LICENSE).
