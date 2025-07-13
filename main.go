package main

import (
	"github.com/Luzifer/password/v2/pkg/cli"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cli.Execute(); err != nil {
		logrus.WithError(err).Fatal("running application")
	}
}
