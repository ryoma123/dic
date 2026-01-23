# dic

[![Build Status](https://travis-ci.org/ryoma123/dic.svg?branch=master)](https://travis-ci.org/ryoma123/dic)

CLI tool for collecting domain information from multiple DNS servers.

## Installation

To install, use `go get`:

```sh
$ go get github.com/ryoma123/dic/...
```

## Usage

Describe settings in `config.toml`:

```toml
[[sec]]
  name = "demo"

  [[sec.args]]
    server = ""
    qtypes = ["a", "txt"]

  [[sec.args]]
    server = "1.1.1.1"
    qtypes = ["a", "txt"]

  [[sec.args]]
    server = "ns"
    qtypes = ["a", "txt"]
```

Pass one or more domains as arguments:

```sh
$ dic example.com www.example.com
[example.com]
  -
    example.com.	40170	IN	A	93.184.216.34
    example.com.	86400	IN	TXT	"v=spf1 -all"
  @1.1.1.1
    example.com.	3200	IN	A	93.184.216.34
    example.com.	5668	IN	TXT	"v=spf1 -all"
 *@a.iana-servers.net.
    example.com.	86400	IN	A	93.184.216.34
    example.com.	86400	IN	TXT	"v=spf1 -all"

[www.example.com]
  -
    www.example.com.	66099	IN	A	93.184.216.34
    www.example.com.	86388	IN	TXT	"v=spf1 -all"
  @1.1.1.1
    www.example.com.	9276	IN	A	93.184.216.34
    www.example.com.	10788	IN	TXT	"v=spf1 -all"
 *@b.iana-servers.net.
    www.example.com.	86400	IN	A	93.184.216.34
    www.example.com.	86400	IN	TXT	"v=spf1 -all"
```

Run via `go run` (recommended without `--` so flags are parsed normally):

```sh
$ go run ./cmd/dic -r 8.8.8.8
$ go run ./cmd/dic -f -r www.example.com
```

Use a specific config file:

```sh
$ go run ./cmd/dic -c ./config.toml -r 8.8.8.8
$ dic -c /path/to/config.toml -f -r www.example.com
```

## Example

[Description example](https://github.com/ryoma123/dic/blob/master/config.toml.example) of config:

```toml
# config.toml.example
[[sec]]
  # [name]
  # Give each section a unique name.
  # It is possible to switch the section used by the option command.
  name = "example"

  [[sec.args]]
    # [server]
    # If blank, query DNS server specified in resolv.conf.
    server = ""
    qtypes = ["a"]

  [[sec.args]]
    # [server]
    # If specified, query that DNS server.
    server = "1.1.1.1"
    qtypes = ["a"]

  [[sec.args]]
    # [server]
    # If you write "ns", get DNS server from NS record and query it.
    server = "ns"
    # [qtypes]
    # Multiple query types available.
    qtypes = ["a", "ns", "cname", "soa", "ptr", "mx", "txt", "aaaa", "any"]

[[sec]]
  name = "example2"
# ...
```

## Commands

### help, h

Display a help message.

### list, l

Show default session and config description:

```sh
$ dic l
DEFAULT SECTION
 resolv

SECTION NAME		SERVER			QUERY TYPES
*resolv			-			[any]

 public			-			[a mx txt]
			1.1.1.1			[a mx txt]

 authoritative		-			[a mx txt]
			ns			[a mx txt]
```

### edit, e

Open and edit config file.

### set `<section name>`, s `<section name>`

Set a section to use by default.

## Options

### --name `<section name>`, -n `<section name>`

Pass a section for temporary use.

### --reverse, -r

Reverse lookup for IP arguments (PTR); domain args remain normal.

### --follow-cname, -f

Follow CNAMEs and query A/AAAA for the target name.

### --cname-max `<n>`, -m `<n>`

Maximum CNAME follow depth (default 5).

### --config `<path>`, -c `<path>`

Path to config file (default: `./config.toml` or GOPATH path).

### --version, -v

Display the version of dic.

## License

[MIT](https://github.com/ryoma123/dic/blob/master/LICENSE)

## Author

[ryoma123](https://github.com/ryoma123)
