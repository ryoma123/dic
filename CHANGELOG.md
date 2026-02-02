# Changelog

## 0.2.0
- Add reverse lookup (PTR) for IP arguments via `--reverse/-r`.
- Add CNAME follow-up queries for A/AAAA via `--follow-cname/-f` and limit depth with `--cname-max/-m`.
- Support `--config/-c` and improve `go run` flag handling.
- Add unit tests for new option parsing and helpers.
- Update module dependencies to modern Go toolchain.
