#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/check.sh      # check_if_empty
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

add_origin_ssh() { # <local-git-dir> <host> <port> <remote-git-dir>
    ORIGIN=ssh://${1}:${2}${3}
    info "Adding origin [${ORIGIN}] ..."
    git remote add origin ${ORIGIN}
}

REMOTE_HOST=${1}
REMOTE_PORT=${2}
REMOTE_GIT=${3}

check_if_not_empty "REMOTE_HOST" ${REMOTE_HOST}
check_if_not_empty "REMOTE_PORT" ${REMOTE_PORT}
check_if_not_empty "REMOTE_GIT" ${REMOTE_GIT}

add_origin_ssh ${REMOTE_HOST} ${REMOTE_PORT} ${REMOTE_GIT}
