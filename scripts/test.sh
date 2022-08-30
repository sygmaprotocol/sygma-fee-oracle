#!/usr/bin/env bash
# The Licensed Work is (c) 2022 Sygma
# SPDX-License-Identifier: BUSL-1.1


go test --race $(go list ./... | grep -v 'e2e')
