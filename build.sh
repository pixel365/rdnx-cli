#!/bin/bash

if [ -z "$1" ]
  then
    echo "Empty argument"
    exit 1
fi

BIN_NAME=rdnx-cli
DARWIN=("amd64" "arm64")
LINUX=("amd64" "arm64")
WINDOWS=("amd64")

case $1 in
    "darwin")
        for arch in "${DARWIN[@]}"
        do
            FN="${BIN_NAME}_${1}_${arch}.tar.gz"
            GOOS=$1 GOARCH=$arch go build -o "./release/${BIN_NAME}"
            cd ./release
            tar -czf $FN $BIN_NAME && shasum $FN
            rm $BIN_NAME
            cd ../
        done
    ;;
    "linux")
        for arch in "${LINUX[@]}"
        do
            FN="${BIN_NAME}_${1}_${arch}.tar.gz"
            GOOS=$1 GOARCH=$arch go build -o "./release/${BIN_NAME}"
            cd ./release
            tar -czf $FN $BIN_NAME && shasum $FN
            rm $BIN_NAME
            cd ../
        done
    ;;
    "windows")
        for arch in "${WINDOWS[@]}"
        do
            FN="${BIN_NAME}_${1}_${arch}.tar.gz"
            GOOS=$1 GOARCH=$arch go build -o "./release/${BIN_NAME}.exe"
            cd ./release
            tar -czf $FN "${BIN_NAME}.exe" && shasum $FN
            rm "${BIN_NAME}.exe"
            cd ../
        done
    ;;
esac
