#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/check.sh      # check_if_exists
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

build_image() { # <user> <image> <tag> <build-path>
    info "Building image [${1}/${2}:${3}] ..."
    docker build -t "${1}/${2}:${3}" -f "${4}/Dockerfile" ${4}
    docker tag "${1}/${2}:${3}" "${1}/${2}:latest"
}

get_cmds() { # <cmd-path>
    # get the command name from its corresponding path.
    for CMD in $(find ${1} -type d); do
        DIR=$(basename ${CMD})
        if [ "${DIR}" != "cmd" ]; then
            echo ${DIR}
        fi
    done
}

get_name() { # <path>
    # use the name of the current directory.
    basename ${1}
}

get_tag() { # <git-dir>
    # use the shortened Git revision as a tag.
    git rev-parse --short HEAD --git-dir=${1}
}

get_user() {
    # use the current system user.
    echo ${USER}
}

BUILD_PATH=${PWD}/build
CMD_PATH=${PWD}/cmd

check_if_exists "BUILD_PATH" ${BUILD_PATH}
check_if_exists "CMD_PATH" ${CMD_PATH}

IMAGE_NAME=$(get_name ${PWD})
IMAGE_TAG=$(get_tag ${PWD})
IMAGE_USER=$(get_user)

build_image ${IMAGE_USER} ${IMAGE_NAME} ${IMAGE_TAG} ${BUILD_PATH}
