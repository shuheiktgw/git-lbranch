git-lbranch
=====
[![GitHub release](http://img.shields.io/github/release/shuheiktgw/git-lbranch.svg?style=flat-square)](https://github.com/shuheiktgw/git-lbranch/releases/latest)
[![Build Status](https://travis-ci.org/shuheiktgw/git-lbranch.svg?branch=master)](https://travis-ci.org/shuheiktgw/git-lbranch)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/shuheiktgw/git-lbranch?status.svg)](https://godoc.org/github.com/shuheiktgw/git-lbranch)

`git-lbranch` is a git subcommand to show a list of recently committed branches. 

## Usage

```
git lbranch [--days] [--through]
```

```
OPTIONS:
  --days value, -d value  specifies the number of days branches last committed (default 5)
  --through, -t           print detailed explanation of branches, adding last commit hash and date
  --version, -v           print the current version
  --help, -h              print help
```


## Install

If you are MacOS user, you can use [Homebrew](http://brew.sh/):

```bash
$ brew tap shuheiktgw/git-lbranch
$ brew install git-lbranch
```

If you use another OS, you can download a suitable binary from [release page](https://github.com/shuheiktgw/git-lbranch/releases) and place it in `$PATH` directory.

Alternatively, if you are Golang programmer, you can use `go get`:

```bash
$ go get -u github.com/shuheiktgw/git-lbranch
```


## Author

[Shuhei Kitagawa](https://github.com/shuheiktgw)






