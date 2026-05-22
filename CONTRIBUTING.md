# Contributing

Thanks for considering a contribution. This project is small on purpose — please
keep PRs focused.

## Local setup

```sh
git clone https://github.com/14ten/certcheck.git
cd certcheck
make test
make build
```

Go 1.22 or newer.

## Conventions

- **Commit messages**: `area: short imperative summary` — e.g. `checker: add --sni`.
- **Branches**: name them after the change (`add-makefile`, `sort-by-expiry`).
- **One change per PR**. Refactors that touch a lot of files should land
  separately from features.
- **Add a test** when fixing a bug or adding behavior.
- **Update `CHANGELOG.md`** under `[Unreleased]` for user-visible changes.

## Before opening a PR

```sh
make vet
make test
```

CI runs the same plus `golangci-lint`. If you can run it locally, even better.

## Reporting bugs

Open an issue with the bug-report template. Include the `certcheck --version`
output, your OS, and the exact command you ran.
