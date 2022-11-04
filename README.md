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



## Output

```
{
  "host": "172.23.179.208",
  "allocatableResource": {
    "milliCPU": 16000,
    "memory": 12871671808,
    "storage": 269490393088
  },
  "requestedResource": {
    "milliCPU": 12,
    "memory": 618688512,
    "storage": 19994185728
  },
  "stats": {
    "cpu": {
      "total": "16 CPU",
      "used": "0%"
    },
    "os": "Ubuntu 20.04",
    "memory": {
      "total": "11 GB",
      "used": "0 GB"
    },
    "storage": {
      "total": "250 GB",
      "used": "18 GB"
    }
  }
}
```



## License

Project License can be found [here](LICENSE).



## Reference

- [gopsutil](https://github.com/shirou/gopsutil)
- [heim](https://github.com/heim-rs/heim)
- [pipego](https://github.com/pipego/plugin-fetch/blob/main/plugin/localhost.go)
- [psutil](https://github.com/giampaolo/psutil)
