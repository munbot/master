// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config2

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/jrmsdev/munbot/flags"
	"github.com/jrmsdev/munbot/log"
)

func New() *Munbot {
	return &Munbot{}
}

func SetDefaults(c *Munbot) {
	c.Master = &Master{
		Name: "munbot",
		Api: &Api{
			Enable: true,
			Addr: "0.0.0.0",
			Port: 3000,
			Cert: filepath.FromSlash("ssl/api/cert.pem"),
			Key: filepath.FromSlash("ssl/api/key.pem"),
			Path: "/api",
		},
		Robot: &Robot{
			Enable: true,
			Name:   "munbot",
			AutoRun: true,
		},
	}
	c.User = make(map[string]*User)
}

func Read(c *Munbot, fh io.ReadCloser) error {
	defer func() {
		if err := fh.Close(); err != nil {
			log.Error(err)
		}
	}()
	blob, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}
	return json.Unmarshal(blob, c)
}

func Bytes(c *Munbot) ([]byte, error) {
	return json.MarshalIndent(c, "", "\t")
}

func Write(c *Munbot, fh io.Writer) error {
	log.Debug("write...")
	blob, err := Bytes(c)
	if err != nil {
		return log.Error(err)
	}
	if _, err := fh.Write(blob); err != nil {
		return log.Error(err)
	}
	return nil
}

func Save(c *Munbot) error {
	log.Debug("save...")
	fn := filepath.Join(flags.ConfigDir, flags.ConfigFile + ".cfg2")
	blob, err := Bytes(c)
	if err != nil {
		return log.Error(err)
	}
	if err := ioutil.WriteFile(fn, blob, 0600); err != nil {
		return log.Error(err)
	}
	log.Printf("%s saved", fn)
	return nil
}

func NewUser(c *Munbot, name string) (*User, error) {
	// TODO: check that name does not exists already
	u := &User{Enable: true, Name: name}
	c.User[name] = u
	return c.User[name], nil
}
