#!/usr/bin/bash

export PATH=$PATH:~/tmp/go/bin

go build -o triangulate.out ./cmd/triangulate/main.go
go build -o integrate.out ./cmd/integrate/main.go
