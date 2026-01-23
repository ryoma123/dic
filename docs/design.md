# Design

## Flags
- `--reverse, -r`: For IP args, query PTR only. Domain args remain normal.
- `--follow-cname, -f`: When a CNAME is returned without A/AAAA, follow to target and query A/AAAA.
- `--cname-max, -m`: Max follow depth (default 5).

## Flow
- Parse options in CLI, normalize defaults.
- If input is IP and `--reverse` is set, query PTR for each configured server.
- For non-IP inputs, run configured qtypes as before.
- If `--follow-cname` is set and CNAME appears without A/AAAA, query target A/AAAA.
- If `--reverse` is also set, query PTR for A/AAAA results from CNAME follow.

## Ordering
- Added records are assigned higher qIndex values so they appear after configured qtypes.
