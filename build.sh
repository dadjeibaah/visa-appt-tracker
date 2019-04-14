#!/usr/bin/env bash

set -eu
declare -a arr=("darwin" "linux")
for platform in "${arr[@]}"; do
   env GOOS=${platform} GOARCH=amd64 go build -o "visa-appt-tracker-$platform" .
done
