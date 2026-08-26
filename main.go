// Password-Generator Util
package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Luzifer/password/v2/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		logrus.WithError(err).Fatal("running application")
	}
}
