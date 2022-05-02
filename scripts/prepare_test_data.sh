#!/usr/bin/env bash

set -e

rm -rf e2e/testdb

# prepare test data to testdb
go test e2e/testdata_test.go
