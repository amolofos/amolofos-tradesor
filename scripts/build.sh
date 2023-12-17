#!/usr/bin/env bash

log () {
  echo "[`date +'%Y-%m-%d %H-%M-%S'`] $@"
}

go build \
  -C ./app/ \
  -o ../build/bin/ \
  -race

log "Build can be found in ../build/bin/"
ls -h ./build/bin/*
