#!/bin/sh
set -eu
SRC=${1:-''}
if test '' = "${SRC}"; then
	SRC='mb'
else
	shift
fi
./build.sh ${SRC}
export MBENV='devel'
export MBENV_CONFIG=${PWD}/env
export MB_CONFIG=${PWD}/_devel/etc
exec ./_build/cmd/${SRC}.bin $@
