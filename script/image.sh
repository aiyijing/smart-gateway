#!/bin/bash
set -x

rm -fr gw
rm -fr client/dist

cd client && npm run build && cd ../
go build -o gw main.go

docker build . -t aiyijing/gw:latest