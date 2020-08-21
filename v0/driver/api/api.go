// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package api implements munbot platform's api driver.
package api

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/log"
)

type Driver struct {
	gobot.Driver
	name string
	conn adaptor.Adaptor
}

func NewDriver(a adaptor.Adaptor) gobot.Driver {
	return &Driver{
		name: "munbot.api",
		conn: a,
	}
}

// gobot interface

func (a *Driver) Name() string {
	return a.name
}

func (a *Driver) SetName(name string) {
	a.name = name
}

func (a *Driver) Connection() gobot.Connection {
	return a.conn
}

func (a *Driver) Start() error {
	log.Printf("Start %s driver.", a.name)
	return nil
}

func (a *Driver) Halt() error {
	log.Printf("Halt %s driver.", a.name)
	return nil
}
