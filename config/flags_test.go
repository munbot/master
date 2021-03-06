// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"flag"
)

func newTestFS() *flag.FlagSet {
	return flag.NewFlagSet("testing", flag.PanicOnError)
}

func (s *Suite) TestFlagsDefaults() {
	f := NewFlags(newTestFS())
	s.False(f.Debug, "default debug")
	s.False(f.Quiet, "default quiet")
	s.True(f.Verbose, "default verbose")
	s.Equal("", f.Name, "default name")
	s.Equal("", f.Profile.Name, "default profile")
}

func (s *Suite) TestFlagsParse() {
	f := NewFlags(newTestFS())
	f.Parse()
	s.False(f.Debug)
	s.False(f.Quiet)
	s.True(f.Verbose)
	s.Equal("master", f.Name)
	s.Equal("testing", f.Profile.Name)
}
