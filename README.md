# Luzifer / password

This project is a rewrite of my former password generator written in Python. The intention is to provide a fast and secure way to generate one or more passwords using a CLI tool or a HTTP(s) API.

## Usage

### Via Web

My service [pwd.knut.me](http://pwd.knut.me/) is powered by this API and will provide you with secure passwords.

### Via CLI

1. Download the compiled binary from GoBuilder.me and unzip the package
2. Generate your password:

    ```bash
    # ./password get -h
    NAME:
      get - generate and return a secure random password

    USAGE:
      command get [command options] [arguments...]

    OPTIONS:
      --length, -l "20"	length of the generated password
      --special, -s	use special characters in your password

    # ./password get
    Vzupi4IaPbXmSQEX9A4e
    ```

### Via API

TBD

## Benchmark / Test

Tests and benchmark are run by Travis CI at every push to this repository:

[![Build Status](https://travis-ci.org/Luzifer/password.svg)](https://travis-ci.org/Luzifer/password)

```bash
# go test -bench .
PASS
BenchmarkGeneratePasswords8Char           20000       74875 ns/op
BenchmarkGeneratePasswords8CharSpecial    10000      108451 ns/op
BenchmarkGeneratePasswords16Char          10000      135059 ns/op
BenchmarkGeneratePasswords16CharSpecial   10000      142958 ns/op
BenchmarkGeneratePasswords32Char           5000      307994 ns/op
BenchmarkGeneratePasswords32CharSpecial    5000      284031 ns/op
BenchmarkGeneratePasswords128Char           300     6721034 ns/op
BenchmarkGeneratePasswords128CharSpecial    500     3244446 ns/op
ok  	github.com/Luzifer/password/lib	13.764s
```
