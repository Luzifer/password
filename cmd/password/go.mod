module github.com/Luzifer/password/v2/cmd/password

go 1.22

replace github.com/Luzifer/password/v2 => ../../

require (
	github.com/Luzifer/go_helpers/v2 v2.25.0
	github.com/Luzifer/password/v2 v2.4.0
	github.com/gorilla/mux v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/GehirnInc/crypt v0.0.0-20230320061759-8cc1b52080c5 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
