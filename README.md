# Luzifer / password

This project is a rewrite of my former password generator written in Python. The intention is to provide a fast and secure way to generate one or more passwords using a CLI tool, a HTTPs page or a HTTPs API.

## Usage

### Via Web

My service [Secure Password](https://pwd.luzifer.io/) is powered by this app and will provide you with secure passwords.

### Via CLI

1. Download the compiled binary from [Github releases](https://github.com/Luzifer/password/releases/latest)
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

    # ./password get -l 32 -s
    }d.sks(4J$2G]x52=k)WAN{M68LxEg}%
    ```

### Via API

- `/v1/getPassword` - Retrieve a password from the API
  - `length=20` - Specify the length of the password to generate
  - `special=false` - Set to `true` to enable special characters

#### Self-Hosted

1. Download the compiled binary from [Github releases](https://github.com/Luzifer/password/releases/latest)
2. Run the API server:

    ```bash
    # ./password serve -h
    NAME:
      serve - start an API server to request passwords

    USAGE:
      command serve [command options] [arguments...]

    OPTIONS:
      --port "3000"	port to listen on
    ```
3. Request your password using `http://localhost:3000/v1/getPassword?length=20&special=true`

#### Hosted

```bash
# curl https://pwd.luzifer.io/v1/getPassword?length=20&special=true
0M4L-1[lT:@2&7,p,o-;
```

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
