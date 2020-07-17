// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/munbot/master/config/parser"
	"github.com/munbot/master/log"
	"github.com/munbot/master/profile"
	"github.com/munbot/master/vfs"
)

type Munbot struct {
	Master *Master `json:"master,omitempty"`
}

type Config struct {
	Munbot *Munbot `json:"munbot,omitempty"`
	handler *parser.Config
}

func New() *Config {
	return &Config{
		handler: parser.New(),
		Munbot: &Munbot{
			Master: &Master{},
		},
	}
}

func (c *Config) SetDefaults() {
	parser.SetDefaults(c.handler)
	c.loadConfig(c.handler)
}

func (c *Config) loadConfig(h *parser.Config) {
	c.Munbot.Master.load(h.Section("master"))
}

func (c *Config) Load(p *profile.Profile) error {
	for _, fn := range p.ListConfigFiles() {
		if err := c.readFile(fn); err != nil {
			return fmt.Errorf("%s: %s", fn, err)
		}
	}
	return nil
}

func (c *Config) readFile(name string) error {
	fh, err := vfs.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug(err)
			return nil
		} else {
			return err
		}
	}
	defer fh.Close()
	return c.Read(fh)
}

func (c *Config) Read(r io.Reader) error {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := c.handler.Load(blob); err != nil {
		return err
	}
	c.loadConfig(c.handler)
	return nil
}

func (c *Config) Save(p *profile.Profile) error {
	fn := p.GetConfigFile()
	fh, err := vfs.Create(fn)
	if err != nil {
		return err
	}
	defer fh.Close()
	return c.Write(fh)
}

func (c *Config) Write(w io.Writer) error {
	blob, err := c.handler.Dump()
	if err != nil {
		return err
	}
	if _, err := w.Write(blob); err != nil {
		return err
	}
	return nil
}
