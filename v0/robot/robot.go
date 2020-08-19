// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package robot implements the core munbot robot.
package robot

import (
	"gobot.io/x/gobot"
)

// Munbot implements the core worker robot.
type Munbot struct {
	*gobot.Robot
}

func New() *Munbot {
	r := &Munbot{}
	r.Robot = gobot.NewRobot(
		"Munbot",
		[]gobot.Connection{},
		[]gobot.Device{},
		r.Work,
	)
	return r
}

// Gobot returns the internal gobot reference, to be added to a master gobot.
func (r *Munbot) Gobot() *gobot.Robot {
	return r.Robot
}

func (r *Munbot) Work() {
}
