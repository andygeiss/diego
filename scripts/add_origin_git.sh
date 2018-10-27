#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/check.sh      # check_if_empty
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

add_origin_git() { # <git-user> <git-repo>
    ORIGIN=git@github.com:$1/$2.git
    info "Adding origin [${ORIGIN}] ..."
    git remote add origin ${ORIGIN}
}

GIT_USER=${1}
GIT_REPO=${2}

check_if_not_empty "GIT_USER" ${GIT_USER}
check_if_not_empty "GIT_REPO" ${GIT_REPO}

add_origin_git ${GIT_USER} ${GIT_REPO}
