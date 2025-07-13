[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/password)](https://goreportcard.com/report/github.com/Luzifer/password)
![GitHub License](https://img.shields.io/github/license/Luzifer/password)
![GitHub Release](https://img.shields.io/github/v/release/Luzifer/password)


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

#### As library in your own code

```go
package main

import pwd "github.com/Luzifer/password/lib/v2"

func getPassword() (string, error) {
	return pwd.NewSecurePassword().GeneratePassword(16, false)
}
```

## Benchmark / Test

Tests and benchmark are run by Github Actions at every push to this repository:

```console
# go test -bench=. ./...    
goos: linux
goarch: amd64
pkg: github.com/Luzifer/password/lib/v2
cpu: AMD Ryzen 9 7950X 16-Core Processor            
BenchmarkGeneratePasswords8Char-32                113706              9946 ns/op
BenchmarkGeneratePasswords8CharSpecial-32          78422             14886 ns/op
BenchmarkGeneratePasswords16Char-32                68236             17104 ns/op
BenchmarkGeneratePasswords16CharSpecial-32         64616             20329 ns/op
BenchmarkGeneratePasswords32Char-32                26655             46330 ns/op
BenchmarkGeneratePasswords32CharSpecial-32         30051             40588 ns/op
BenchmarkGeneratePasswords128Char-32                1119           1144492 ns/op
BenchmarkGeneratePasswords128CharSpecial-32         2860            494610 ns/op
BenchmarkGeneratePasswords4Words-32               421377              2435 ns/op
BenchmarkGeneratePasswords20Words-32               97959             11978 ns/op
PASS
ok      github.com/Luzifer/password/lib/v2      13.962s
```
