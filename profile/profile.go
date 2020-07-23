// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package profile handles profiled settings.
package profile

import (
	"os"
	"path/filepath"
)

var (
	homeDir      string
	homeDirErr   error
	configDir    string
	configDirErr error
)

func init() {
	homeDir, homeDirErr = os.UserHomeDir()
	configDir, configDirErr = os.UserConfigDir()
}

// Profile holds the profile settings.
type Profile struct {
	Name           string
	ConfigFilename string
	ConfigDir      string
	ConfigSysDir   string
	ConfigDistDir  string
}

func setDefaults(p *Profile) *Profile {
	if homeDir == "" || homeDirErr != nil {
		homeDir = filepath.FromSlash("./.munbot")
	} else {
		homeDir = filepath.Join(homeDir, ".munbot")
	}
	if configDir == "" || configDirErr != nil {
		configDir = filepath.Join(homeDir, "config")
	} else {
		configDir = filepath.Join(configDir, "munbot")
	}
	p.ConfigFilename = "config.json"
	p.ConfigDir = configDir
	p.ConfigSysDir = filepath.FromSlash("/usr/local/etc/munbot")
	p.ConfigDistDir = filepath.FromSlash("/etc/munbot")
	return p
}

// New creates a new object with defaults values set.
func New(name string) *Profile {
	return setDefaults(&Profile{Name: name})
}

// GetConfigFile returns the absolute filename of the configuration file for the
// current os user. This is the filename used to save configuration updates.
func (p *Profile) GetConfigFile() string {
	cfgDir := filepath.Clean(p.ConfigDir)
	return filepath.Join(cfgDir, p.Name, p.ConfigFilename)
}

// ListConfigFiles returns a list of all the profiled filenames to read the
// configuration from.
func (p *Profile) ListConfigFiles() []string {
	l := make([]string, 0)
	if p.ConfigSysDir != "" {
		sysDir := filepath.Clean(p.ConfigSysDir)
		l = append(l, filepath.Join(sysDir, p.ConfigFilename))
		l = append(l, filepath.Join(sysDir, p.Name, p.ConfigFilename))
	}
	cfgDir := filepath.Clean(p.ConfigDir)
	l = append(l, filepath.Join(cfgDir, p.ConfigFilename))
	l = append(l, filepath.Join(cfgDir, p.Name, p.ConfigFilename))
	if p.ConfigDistDir != "" {
		distDir := filepath.Clean(p.ConfigDistDir)
		l = append(l, filepath.Join(distDir, p.ConfigFilename))
		l = append(l, filepath.Join(distDir, p.Name, p.ConfigFilename))
	}
	return l
}
