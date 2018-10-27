#!/usr/bin/env bash

fatal() {
    log "ERROR:" $@
    exit 1
}

info() {
    log "INFO :" $@
}

log() {
    echo $(date "+%Y-%m-%d %H:%M:%S") $@
}

warn() {
    log "WARN :" $@
}
