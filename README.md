# go-mattercloud
**go-mattercloud** is the unofficial golang implementation for the [MatterCloud API](https://developers.mattercloud.net/)

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-mattercloud?v=1)](https://golang.org/)
[![Build Status](https://travis-ci.com/mrz1836/go-mattercloud.svg?branch=master&v=1)](https://travis-ci.com/mrz1836/go-mattercloud)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-mattercloud?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-mattercloud)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-mattercloud.svg?style=flat&v=1)](https://github.com/mrz1836/go-mattercloud/releases)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-mattercloud?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-mattercloud)

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-mattercloud** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-mattercloud
```

## Documentation
You can view the generated [documentation here](https://pkg.go.dev/github.com/mrz1836/go-mattercloud).

You can also view the [MatterCloud API](https://developers.mattercloud.net/) documentation.

### Features
- [Client](client.go) is completely configurable
- Customize the network per request (`main`, `test` or `stn`)
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Current (V3) coverage for the [MatterCloud API](https://developers.mattercloud.net/) API
    - [x] Authentication
    - [x] Address
    - [x] Transaction

<details>
<summary><strong><code>Library Deployment</code></strong></summary>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>

View all `makefile` commands
```bash
$ make help
```

List of all current commands:
```text
all                            Runs test, install, clean, docs
bench                          Run all benchmarks in the Go application
clean                          Remove previous builds and any test cache data
clean-mods                     Remove all the Go mod cache
coverage                       Shows the test coverage
godocs                         Sync the latest tag with GoDocs
help                           Show all make commands available
lint                           Run the Go lint application
release                        Full production release (creates release in Github)
release-test                   Full production test release (everything except deploy)
release-snap                   Test the full release (build binaries)
tag                            Generate a new tag and push (IE: make tag version=0.0.0)
tag-remove                     Remove a tag if found (IE: make tag-remove version=0.0.0)
tag-update                     Update an existing tag to current commit (IE: make tag-update version=0.0.0)
test                           Runs vet, lint and ALL tests
test-short                     Runs vet, lint and tests (excludes integration tests)
update                         Update all project dependencies
update-releaser                Update the goreleaser application
vet                            Run the Go vet application
```
</details>

## Examples & Tests
All unit tests and [examples](mattercloud_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-mattercloud) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ make test
```

Run tests (excluding integration tests)
```bash
$ make test-short
```

## Benchmarks
Run the Go [benchmarks](mattercloud_test.go):
```bash
$ make bench
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [mattercloud examples](#examples--tests) above

Basic implementation:
```go
package main

import (
	"log"

	"github.com/mrz1836/go-mattercloud"
)

func main() {

	// Create a new client
	client, _ := mattercloud.NewClient("your-secret-api-key", mattercloud.NetworkMain, nil)

	// Get balance for an address
	balance, _ := client.AddressBalance("16ZqP5Tb22KJuvSAbjNkoiZs13mmRmexZA")

	// What's the confirmed balance?
	log.Println("confirmed balance:", balance.Confirmed)
}
```

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-mattercloud)

#### Credits

[@Attila](https://github.com/attilaaf) & [MatterCloud](https://mattercloud.net/) for their hard work on the [MatterCloud API](https://developers.mattercloud.net/)

Looking for a Javascript version? Check out the [MatterCloud JS SDK](https://github.com/MatterCloud/mattercloudjs)

## License

![License](https://img.shields.io/github/license/mrz1836/go-mattercloud.svg?style=flat&v=2)