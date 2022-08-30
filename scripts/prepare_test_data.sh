#!/usr/bin/env bash
# The Licensed Work is (c) 2022 Sygma
# SPDX-License-Identifier: BUSL-1.1


set -e

rm -rf e2e/testdb

# prepare test data to testdb
go test e2e/testdata_test.go
