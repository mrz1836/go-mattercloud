# go-mattercloud
**go-mattercloud** is the unofficial golang implementation for the [MatterCloud API](https://developers.mattercloud.net/)

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-mattercloud)](https://golang.org/)
[![Build Status](https://travis-ci.com/mrz1836/go-mattercloud.svg?branch=master&v=1)](https://travis-ci.com/mrz1836/go-mattercloud)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-mattercloud?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-mattercloud)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/dde6d46426bd4c12be65916da8cf04d2)](https://www.codacy.com/app/mrz1818/go-mattercloud?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-mattercloud&amp;utm_campaign=Badge_Grade)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-mattercloud.svg?style=flat&v=2)](https://github.com/mrz1836/go-mattercloud/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-mattercloud?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-mattercloud)

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
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-mattercloud).

You can also view the [MatterCloud API](https://developers.mattercloud.net/) documentation.

### Features
- [Client](client.go) is completely configurable
- Customize the network per request (`main`, `test` or `stn`)
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Current (V3) coverage for the [MatterCloud API](https://developers.mattercloud.net/) API
    - [x] Address
    - [ ] Transaction

## Examples & Tests
All unit tests and [examples](mattercloud_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-mattercloud) and uses [Go version 1.13.x](https://golang.org/doc/go1.13). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ cd ../go-mattercloud
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-mattercloud
$ go test ./... -v -test.short
```

## Benchmarks
Run the Go [benchmarks](mattercloud_test.go):
```bash
$ cd ../go-mattercloud
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [mattercloud examples](#examples--tests) above

Basic implementation:
```golang
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

![License](https://img.shields.io/github/license/mrz1836/go-mattercloud.svg?style=flat&v=1)