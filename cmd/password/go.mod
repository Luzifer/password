module github.com/Luzifer/password/cmd/password

go 1.12

replace github.com/Luzifer/password/v2 => ../../

require (
	github.com/Luzifer/go_helpers/v2 v2.9.1
	github.com/Luzifer/password/v2 v2.1.0
	github.com/gorilla/mux v1.7.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20190228124157-a34e9553db1e // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
