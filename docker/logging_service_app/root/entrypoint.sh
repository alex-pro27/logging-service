#!/usr/bin/env sh

GO_WORK_DIR=${GO_WORK_DIR:-$GOPATH/src}
cd ${GO_WORK_DIR}

if  [[ 1 = "$DEBUG" ]]; then
    echo "Run debug mode"
    exec dlv debug --headless --listen=:2345 --api-version=2
else
    echo "Run prod mode"
    go build -o app_server && exec ./app_server
fi