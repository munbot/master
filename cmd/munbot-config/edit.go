// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/munbot/master/config"
	"github.com/munbot/master/log"
)

func edit(cfg *config.Munbot, filter, args string) error {
	log.Debug("edit...")
	//~ if err := cfg.Update(filter, args); err != nil {
	//~ return log.Error(err)
	//~ }
	return save(cfg)
}
