#!/bin/bash -eu

# Extract value of Version const from version.go
VERSION=`gobump show -r`

# Path to built files
FILES=./pkg/dist/v${VERSION}

goxz -pv=v${VERSION} -arch=386,amd64 -d=${FILES}
ghr --soft v${VERSION} ${FILES}
ghbr release --merge