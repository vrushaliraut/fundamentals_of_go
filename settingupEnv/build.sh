#!/bin/bash
# A simple CI script to build a Go project

# Exit immediately if any command fails

set -e

echo "=== Formatting ==="

# Note: In CI, you might use 'gofmt -l .' to check for unformatted
# files and fail, but 'go fmt' is the simplest form.
go fmt ./...

echo "=== Vetting Code ==="
# This will fail the script if vet finds any issues

go vet ./...

echo "=== Building Binary ==="
go build -o settingup_env

echo "=== Build Successful ==="
