// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"

	"gobot.io/x/gobot/sysfs"
)

type File sysfs.File

type Filesystem sysfs.Filesystem

var fs Filesystem

func init() {
	fs = new(sysfs.NativeFilesystem)
}

func SetFilesystem(newfs Filesystem) {
	fs = nil
	fs = newfs
}

func OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return fs.OpenFile(name, flag, perm)
}

func Stat(name string) (os.FileInfo, error) {
	return fs.Stat(name)
}

func Open(name string) (File, error) {
	return fs.OpenFile(name, os.O_RDONLY, 0)
}
