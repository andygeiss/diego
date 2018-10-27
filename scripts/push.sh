#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

push_git() { # <local-git-dir>
    info "Pushing repository ..."
    cd ${1}
    git push -f -u origin master
    cd -
}

push_git ${PWD}
