# exporter

[![Build Status](https://github.com/pipego/exporter/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/pipego/exporter/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/pipego/exporter/branch/main/graph/badge.svg?token=61G1TNDUS6)](https://codecov.io/gh/pipego/exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/pipego/exporter)](https://goreportcard.com/report/github.com/pipego/exporter)
[![License](https://img.shields.io/github/license/pipego/exporter.svg)](https://github.com/pipego/exporter/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/pipego/exporter.svg)](https://github.com/pipego/exporter/tags)



## Introduction

*exporter* is the exporter of [pipego](https://github.com/pipego) written in Go.



## Prerequisites

- Go >= 1.18.0



## Run

```bash
version=latest make build
./bin/exporter
```



## Docker

```bash
version=latest make docker
docker run ghcr.io/pipego/exporter:latest
```



## Usage

```
usage: exporter [<flags>]

pipego exporter

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.
```



## License

Project License can be found [here](LICENSE).



## Reference
