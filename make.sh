#!/bin/bash
CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 
go build -tags static -ldflags "-s -w" -o build 
