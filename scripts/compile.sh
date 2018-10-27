#!/usr/bin/env bash

MAIN_DIR=$(dirname ${0})
source ${MAIN_DIR}/include/check.sh      # check_if_exists
source ${MAIN_DIR}/include/logging.sh    # fatal, info, warn

export CGO_ENABLED=0

compile_static() { # <target> <source>
    if [ -e "${2}" ]; then
        # compile for the current arch/os
        go build -ldflags "-s" -o ${1} ${2}
        # use upx for compression
        upx -1 -q ${1} 2>&1>/dev/null
        # cross-compile
        for GOARCH in amd64 arm64 arm s390x; do
            for GOOS in linux; do
                GOARCH=${GOARCH} GOARM=5 GOOS=${GOOS} go build -ldflags "-s" -o ${1}-${GOOS}-${GOARCH} ${2}
                if [ "$GOARCH" != "s390x" ]; then
                    upx -1 -q ${1}-${GOOS}-${GOARCH} 2>&1>/dev/null
                fi
            done
        done
    fi
}

compile_wasm() { # <target> <source>
    if [ -e "${2}" ]; then
        # >= go1.11.1 required to compile to Webassembly
        GOARCH=wasm GOOS=js go build -o ${1}.wasm ${2}
    fi
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

BUILD_PATH=${PWD}/build
CMD_PATH=${PWD}/cmd

check_if_exists "BUILD_PATH" ${BUILD_PATH}
check_if_exists "CMD_PATH" ${CMD_PATH}

for CMD in $(get_cmds ${CMD_PATH}); do
    compile_static ${BUILD_PATH}/package/${CMD} ${CMD_PATH}/${CMD}/main.go
    compile_wasm ${BUILD_PATH}/package/${CMD} ${CMD_PATH}/${CMD}/main_wasm.go
done
