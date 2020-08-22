// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package master implements the master robot.
package master

import (
	"gobot.io/x/gobot"

	"github.com/munbot/master/v0/adaptor"
	"github.com/munbot/master/v0/log"
	"github.com/munbot/master/v0/robot"
)

// Robot works around a gobot.Master.
type Robot struct {
	*gobot.Master
	AutoRun bool
}

// New creates a new master robot.
func New() *Robot {
	return &Robot{Master: gobot.NewMaster(), AutoRun: true}
}

// Start starts the core runtime and then the gobot master robot.
func (m *Robot) Start() error {
	log.Print("Start master robot.")
	m.Master.AutoRun = m.AutoRun
	return m.Master.Start()
}

// Stop stops the gobot master robot and then the core runtime.
func (m *Robot) Stop() error {
	log.Print("Stop master robot.")
	return m.Master.Stop()
}

// Run runs the robot's main loop.
//
// Add core munbot robot.
// Start.
// Watch for failures and abort if any.
func (m *Robot) Run() error {
	log.Debug("run...")
	log.Debug("add core munbot")
	m.AddRobot(robot.New(adaptor.New(m.Master)).Gobot())
	if err := m.Start(); err != nil {
		m.Stop()
		return err
	}
	return nil
}
