<div class="info" align="center">
  <h1 class="name">📬<br>filateli</h1>
  convert postman collection to markdown files
  <br>
  <br>

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]

</div>


## Installation

Download latest binary release from release page.

You can use this install script to download the latest version:

```sh
# install latest release to /usr/local/bin/
curl https://i.jpillora.com/avrebarra/filateli! | *remove_this* bash
```

```sh
# install specific version
curl https://i.jpillora.com/avrebarra/filateli@{version} | *remove_this* bash
```

## Usage
### CLI

```bash
$ filateli -help

filateli v0 - postman collections as local documentation server

Available commands:

   build   build documentation folder 

Flags:

  -help
        Get help on the 'filateli' command.
  -quiet
        perform quiet operation
```

[godoc-image]: https://godoc.org/github.com/avrebarra/filateli?status.svg
[godoc-url]: https://godoc.org/github.com/avrebarra/filateli
[report-image]: https://goreportcard.com/badge/github.com/avrebarra/filateli
[report-url]: https://goreportcard.com/report/github.com/avrebarra/filateli
[tests-image]: https://cloud.drone.io/api/badges/avrebarra/filateli/status.svg
[tests-url]: https://cloud.drone.io/avrebarra/filateli
[coverage-image]: https://codecov.io/gh/avrebarra/filateli/graph/badge.svg
[coverage-url]: https://codecov.io/gh/avrebarra/filateli
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg