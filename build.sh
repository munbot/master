#!/bin/sh
set -eu
SRC=${1:-'munbot'}
# https://github.com/golang/go/issues/26492#issuecomment-635563222
# STATIC="-tags 'osusergo netgo'"
TAGS='munbot'
if test 'static' = "${2:-'default'}"; then
	TAGS='munbot,static,osusergo,netgo'
fi
imp="github.com/jrmsdev/munbot/version"
BUILD_DATE="-X ${imp}.buildDate=$(date -u '+%Y%m%d.%H%M%S')"
BUILD_INFO="-X ${imp}.buildOS=$(go env GOOS)"
BUILD_INFO="${BUILD_INFO} -X ${imp}.buildArch=$(go env GOARCH)"
BUILD_INFO="${BUILD_INFO} -X ${imp}.buildTags=${TAGS}"
build_cmds=${SRC}
if test 'all' = "${build_cmds}"; then
	build_cmds='munbot munbot-config'
fi
for cmd in ${build_cmds}; do
	go build -v -mod vendor -i -o ./_build/cmd/${cmd}.bin -tags "${TAGS}" \
		-ldflags "${BUILD_DATE} ${BUILD_INFO}" ./cmd/${cmd} || exit 1
	echo "--- build $(ls ./_build/cmd/${cmd}.bin) done!"
done
exit 0
