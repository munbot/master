// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"os"

	"github.com/jrmsdev/munbot"
	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func main() {
	flags.Init("munbot-config")
	flags.Parse(os.Args[1:])
	log.Debug("start")
	cfg := munbot.Configure()
	log.Debug(cfg)
	cfg.Dump()
	cfg.Write(os.Stdout)
	//~ cfg.Update("name", "lalala")
	//~ cfg.Dump()
	//~ cfg.Write(os.Stdout)
	//~ log.Debug(cfg.Update("lalala", "name"))
	log.Printf("master.name: %s", cfg.Master.Name)
	log.Debug("end")
}
