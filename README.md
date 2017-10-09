# fsys

Local and remote Filesystem abstraction

## Badges

[![travis-ci](https://travis-ci.org/trussle/fsys.svg?branch=master)](https://travis-ci.org/trussle/fsys)
[![Coverage Status](https://coveralls.io/repos/github/trussle/fsys/badge.svg?branch=master)](https://coveralls.io/github/trussle/fsys?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/trussle/fsys)](https://goreportcard.com/report/github.com/trussle/fsys)

## Introduction

Fsys in an abstraction layer for interacting with local and remote Filesystems.

### Implementations

 - `nop`: Filesystem with sane defaults that does absolutely nothing.
 - `virtual`: Filesystem that utilises memory
 - `local`: Filesystem that utilises the local file system.
   - As a side note `mmap` can be used as an option for the `local` setup.

 