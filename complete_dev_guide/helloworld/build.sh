#! /bin/bash

Filename="$1"

#go get $Filename
go fmt $Filename
go build $Filename
go run $Filename
