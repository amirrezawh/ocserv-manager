package main

import (
	"fmt"

	"github.com/amirrezawh/ocserv-manager/config"
	db "github.com/amirrezawh/ocserv-manager/pkg/db"
	gen "github.com/amirrezawh/ocserv-manager/pkg/generator"
)

func main() {
	fmt.Println("Starting ocserv manager")

	// set environment variables
	cfg := &config.GeneralConfig{}
	config.SetConfigs(cfg)
	// start migration
	db.Migrate(cfg)
	gen.Interval(cfg)

}
