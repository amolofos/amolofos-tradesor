#!/usr/bin/env bash

log () {
  echo "[`date +'%Y-%m-%d %H-%M-%S'`] $@"
}

log "go fmt"
go fmt -C app ./...
