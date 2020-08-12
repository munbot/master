#!/bin/sh
if test "X${1}" = 'X--all'; then
	exec gofmt -w -l -s .
fi
exec gofmt -w -l -s \
	api \
	auth \
	cmd \
	config \
	console \
	env \
	internal \
	log \
	mb \
	platform \
	robot \
	testing \
	utils \
	version \
	vfs
