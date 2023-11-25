#!/usr/bin/env bash

log () {
  echo "[`date +'%Y-%m-%d %H-%M-%S'`] $@"
}

mkdir -p ./build/test/

log "go vet"
go vet -C app ./...

log "go test"
go test -C app -v -coverpkg=./... -coverprofile=./../build/coverage.out ./...
