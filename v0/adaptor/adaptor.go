// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package adaptor implements the munbot gobot.Adaptor interface.
package adaptor

import (
	"time"

	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/internal/core"
	"github.com/munbot/master/v0/log"
)

var _ gobot.Connection = &Munbot{}

type Adaptor interface {
	gobot.Adaptor
	Interval() time.Duration
	SetInterval(time.Duration)
}

type Munbot struct {
	name     string
	interval time.Duration
}

func New() *Munbot {
	return &Munbot{
		name:     "munbot",
		interval: 300 * time.Millisecond,
	}
}

// gobot interface

func (m *Munbot) Name() string {
	return m.name
}

func (m *Munbot) SetName(name string) {
	m.name = name
}

func (m *Munbot) Connect() error {
	log.Printf("Connect %s platform.", m.name)
	log.Debug("lock core runtime")
	core.Lock()
	return nil
}

func (m *Munbot) Finalize() error {
	log.Printf("Finalize %s platform.", m.name)
	log.Debug("unlock core runtime")
	core.Unlock()
	return nil
}

// munbot interface

func (m *Munbot) Interval() time.Duration {
	return m.interval
}

func (m *Munbot) SetInterval(d time.Duration) {
	m.interval = d
}
