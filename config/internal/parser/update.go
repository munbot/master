// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"fmt"

	"github.com/munbot/master/config/value"
)

func Update(c *Config, option, newval string) error {
	sect, opt := c.getSectOpt(option)
	if opt == "" {
		return fmt.Errorf("update invalid format: %s %s", option, newval)
	}
	if !c.HasSection(sect) {
		return fmt.Errorf("update invalid section: %s", sect)
	}
	if !c.HasOption(sect, opt) {
		return fmt.Errorf("update invalid option: %s.%s", sect, opt)
	}
	c.db[sect][opt] = newval
	return nil
}

func Set(c *Config, option, val string) error {
	sect, opt := c.getSectOpt(option)
	if opt == "" {
		return fmt.Errorf("set invalid format: %s %s", option, val)
	}
	if !c.HasSection(sect) {
		c.db[sect] = value.Map{}
	} else if c.HasOption(sect, opt) {
		return fmt.Errorf("set option already exists: %s.%s", sect, opt)
	}
	c.db[sect][opt] = val
	return nil
}

func Unset(c *Config, option string) error {
	sect, opt := c.getSectOpt(option)
	if opt == "" {
		return fmt.Errorf("unset invalid option: %s", option)
	}
	if c.HasOption(sect, opt) {
		delete(c.db[sect], opt)
	}
	return nil
}
