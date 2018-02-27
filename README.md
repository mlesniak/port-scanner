[![Build Status](https://travis-ci.org/mlesniak/port-scanner.svg?branch=master)](https://travis-ci.org/mlesniak/port-scanner)
[![Code of Conduct](https://img.shields.io/badge/%E2%9D%A4-code%20of%20conduct-orange.svg?style=flat)](CODE_OF_CONDUCT.md)

# Overview

Implementation of a simple port scanner in Go, mirroring the output of nmap.

## Example

    > port-scanner -hostname mlesniak.com -parallel 20 -port 75-85 -timeout 1
    PORT      STATUS  SERVICE
    75/tcp    closed  
    76/tcp    closed  deos
    77/tcp    closed  
    78/tcp    closed  vettcp
    79/tcp    closed  finger
    80/tcp    open    www-http
    81/tcp    closed  
    82/tcp    closed  xfer
    83/tcp    closed  mit-ml-dev
    84/tcp    closed  ctf
    85/tcp    closed  mit-ml-dev

## Help

A list of available command line options can be obtained by executing

    > port-scanner -help
    A simple port scanner in go.
    -hostname string
            hostname of the target system
    -parallel int
            Maximum number of parallel connections (default 1)
    -port string
            a single port (80) or a single range (80-1024)
    -timeout float
            Timeout in seconds. Fractional values, e.g. 0.5 are allowed (default 1)

## Building

To build and install port-scanner under `$GOPATH` you have to

    git clone git@github.com:mlesniak/port-scanner.git
    go install

We use [go-bindata](https://github.com/a-urth/go-bindata) to embed files in `data/`, hence to build 
`bindata.go`, you have to

    go-bindata data/

If you have not installed `go-bindata`, use

    go get -u github.com/a-urth/go-bindata/...

beforehand.

To reduce the file size, use [upx](https://upx.github.io/) and

    go build && strip port-scanner && upx -9 port-scanner

to create a 1MB single static file, e.g. for using it in docker containers.

## Limitations

While this application is feature complete for my usages, the following limitations apply:

- Scanning of TCP ports only.
- Service list maps only single range ports, i.e. xwindow's definition from 6000-6003 is currently not correctly mapped

If these limitations annoy you, either fix this yourself and write a pull request :-) or open an issue.

## Organization

A Trello Board can be found [here](https://trello.com/b/opzPa3fd/port-scanner).

## Tools

We use [go-bindata](https://github.com/a-urth/go-bindata).

## License

The source code is licensed under the [Apache license](https://raw.githubusercontent.com/mlesniak/port-scanner/master/LICENSE)
