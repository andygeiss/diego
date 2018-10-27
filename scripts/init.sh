#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

init_git() { # <local-git-dir>
    info "Initializing repository in [${1}] ..."
    rm -rf ${1}/.git
    cd ${1}
    git init
    git add .
    git commit -m "initial commit" .
    cd -
}

init_git ${PWD}
