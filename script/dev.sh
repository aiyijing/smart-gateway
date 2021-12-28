#!/bin/bash

function dev(){
  go build -o gw main.go
  cd client && npm run build && cd ..
}

function dev_backend() {
  if [ ! -f "client/dist/index.html" ];then
    cd client && npm run build && cd ..
  fi
  go build -o gw main.go
}

function dev_frontend() {
  if [ ! -f "./gw" ];then
    go build -o gw main.go
  fi
  cd client && npm run build && cd ..
}

if [ ! -n "$1" ];then
  dev
else
  $1
fi

# run dev
./gw --address ":8080" --data ./cfg --static "client/dist" --debug debug
