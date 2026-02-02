# PRD

## Goal
Add reverse lookup and CNAME-following capabilities to the dic CLI for faster DNS investigation.

## Scope
- Reverse lookup (PTR) for IP arguments via `--reverse`/`-r`.
- Follow CNAME to fetch A/AAAA via `--follow-cname`/`-f`.
- Limit CNAME follow depth with `--cname-max`/`-m`.
- Add unit tests for new helper logic.

## Non-Goals
- GUI or interactive mode changes.
- New config file formats.

## Success Metrics
- New flags appear in help and function as expected.
- Tests pass without network dependency.
