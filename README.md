[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/password)](https://goreportcard.com/report/github.com/Luzifer/password)
![](https://badges.fyi/github/license/Luzifer/password)
![](https://badges.fyi/github/downloads/Luzifer/password)
![](https://badges.fyi/github/latest-release/Luzifer/password)

# Luzifer / password

The intention of this project is to provide a fast and secure way to generate one or more passwords using a CLI tool, a HTTPs page or a HTTPs API.

For the security of the passwords there are several assertions:

- The password may not contain pattern found on the keyboard or in alphabet
- The password must have 3 or 4 different character groups in it depending on whether special characters are requested
- The password may not have repeating characters
- The API generator does not transmit or store any data about the generated passwords

## Usage

### Via Web

My service [Secure Password](https://passwd.fyi/) is powered by this app and will provide you with secure passwords.

### Via CLI

1. Download the compiled binary from [Github releases](https://github.com/Luzifer/password/releases/latest)
2. Generate your password:

    ```console
    $ ./password get -h
    generate and return a secure random password
    
    Usage:
      password get [flags]
    
    Flags:
      -d, --date         prepend current date to XKCD style passwords (default true)
      -h, --help         help for get
      -j, --json         return output in JSON format
      -l, --length int   length of the generated password (default 20)
      -n, --number int   number of passwords to generate (default 1)
      -s, --special      use special characters in your password
      -x, --xkcd         use XKCD style password

    $ ./password get
    Vzupi4IaPbXmSQEX9A4e

    $ ./password get -l 32 -s
    }d.sks(4J$2G]x52=k)WAN{M68LxEg}%

    $ ./password get -l 4 -x
    20190101.SeashellSupporterTumbleweedGeneral
    ```

### Via API

- `/v1/getPassword` - Retrieve a password from the API
  - `date=true` - Set to `false` no to prepend the date to XKCD-style passwords
  - `length=20` - Specify the length of the password to generate (the API only supports values between 4 and 128 - for more characters use the CLI)
  - `special=false` - Set to `true` to enable special characters
  - `xkcd=false` - Set to `true` to enable XKCD-style passwords

#### Self-Hosted

1. Download the compiled binary from [Github releases](https://github.com/Luzifer/password/releases/latest)
2. Run the API server:

    ```console
    $ ./password serve -h
    start an API server to request passwords

    Usage:
      password serve [flags]

    Flags:
      -h, --help       help for serve
          --port int   port to listen on (default 3000)
    ```
3. Request your password using `http://localhost:3000/v1/getPassword?length=20&special=true`

#### Hosted

```console
$ curl https://passwd.fyi/v1/getPassword?length=20&special=true
0M4L-1[lT:@2&7,p,o-;
```

## Benchmark / Test

Tests and benchmark are run by Travis CI at every push to this repository:

[![Build Status](https://travis-ci.org/Luzifer/password.svg)](https://travis-ci.org/Luzifer/password)

```console
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/Luzifer/password/lib
BenchmarkGeneratePasswords8Char-8                  20000             65469 ns/op
BenchmarkGeneratePasswords8CharSpecial-8           20000             97659 ns/op
BenchmarkGeneratePasswords16Char-8                 20000             84215 ns/op
BenchmarkGeneratePasswords16CharSpecial-8          20000             92885 ns/op
BenchmarkGeneratePasswords32Char-8                 10000            152436 ns/op
BenchmarkGeneratePasswords32CharSpecial-8          10000            144352 ns/op
BenchmarkGeneratePasswords128Char-8                 1000           2199011 ns/op
BenchmarkGeneratePasswords128CharSpecial-8          2000           1089225 ns/op
BenchmarkGeneratePasswords4Words-8                200000              9472 ns/op
BenchmarkGeneratePasswords20Words-8               100000             14098 ns/op
PASS
ok      github.com/Luzifer/password/lib 21.624s
```
