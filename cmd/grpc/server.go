package main

import (
	"github.com/microserv-io/oauth-credentials-server/internal"
	"log"
)

const CfgPath = "/cfg"

func main() {
	application, err := internal.NewApplication(CfgPath)
	if err != nil {
		log.Panicf("failed to create application: %v", err)
	}

	if err := application.Run(""); err != nil {
		log.Panicf("failed to run application: %v", err)
	}
}
