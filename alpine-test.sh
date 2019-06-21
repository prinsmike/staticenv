#!/usr/bin/env bash

APPNAME="staticenv"
IMG="alpine-go-builder"
PREFIX=""
TAG=""

main() {
	while getopts ":hp:" opt; do
		case ${opt} in
			h )
				usage
				exit 2
				;;
			t )
				TAG=$OPTARG
				;;
			p )
				PREFIX=$OPTARG
				;;
			\? )
				echo "Invalid option: -$OPTARG" 1>&2
				echo
				usage
				exit 1
				;;
			: )
				echo "Invalid option: -$OPTARG requires an argument" 1>&2
				echo
				usage
				exit 1
				;;
		esac
	done
	shift $((OPTIND -1))

	run
}

usage () {
	echo "Usage: $0 [OPTIONS]"
	echo "      Run tests for the $APPNAME package."
	echo "  -t"
	echo "      The image release tag."
	echo "      (default: latest)"
	echo "  -p"
	echo "      The prefix for the image name. (i.e. PREFIX/IMAGENAME)"
	echo " -h"
	echo "      Display this help information."
}

run() {
	if [[ ! -z "$PREFIX" ]]; then
		PREFIX="${PREFIX}/"
	fi
	if [[ -z "$TAG" ]]; then
		TAG=latest
	fi

	docker run -it --rm -v $PWD:/usr/src/myapp \
		-e GO111MODULE=on \
		-w /usr/src/myapp \
		"${PREFIX}${IMG}:${TAG}" \
		go test -v -coverprofile cover.out
}

main $@
