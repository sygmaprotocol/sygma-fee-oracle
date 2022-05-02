#!/usr/bin/env bash

go test --race $(go list ./... | grep -v 'e2e')
