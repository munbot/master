// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package vfs

import (
	"os"
)

// NativeFilesystem it's just a wrapper for os package functions.
type NativeFilesystem struct{}

// OpenFile calls os.OpenFile.
func (fs *NativeFilesystem) OpenFile(name string, flag int) (File, error) {
	perm := filePerm
	if flag == os.O_RDONLY {
		perm = 0
	}
	return os.OpenFile(name, flag, perm)
}

// Stat calls os.Stat.
func (fs *NativeFilesystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Mkdir calls os.Mkdir.
func (fs *NativeFilesystem) Mkdir(path string) error {
	return os.Mkdir(path, dirPerm)
}

// MkdirAll calls os.MkdirAll.
func (fs *NativeFilesystem) MkdirAll(path string) error {
	return os.MkdirAll(path, dirPerm)
}
