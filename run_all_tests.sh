#!/bin/sh
go list -f '{{.Dir}}/...' -m | sort | xargs go test
