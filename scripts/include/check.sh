#!/usr/bin/env bash

check_if_exists() { # <label> <path>
    if [ -e "${2}" ]; then
        info "Checking if [${1}] exists: [${2}] ... TRUE"
    else
        fatal "Checking if [${1}] exists: [${2}] ... FALSE!!!"
    fi
}

check_if_not_empty() { # <label> <name>
    if [ "${2}" != "" ]; then
        info "Checking if [${1}] is not empty: [${2}] ... TRUE"
    else
        fatal "Checking if [${1}] is not empty: [${2}] ... FALSE!!!"
    fi
}
