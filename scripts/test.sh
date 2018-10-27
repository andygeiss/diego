#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

test() { # <package>
    info "Testing [${1}] ..."
    go test -v ${1}
}

test ./internal/...
