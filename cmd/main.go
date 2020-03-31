package main

import (
	"os"

	"code.cloudfoundry.org/quarks-container-run/cmd/containerrun"
)

func main() {
	if err := containerrun.NewDefaultContainerRunCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
