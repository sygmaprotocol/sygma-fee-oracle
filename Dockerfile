# Copyright 2021 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM  golang:1.17-stretch AS builder
WORKDIR /app
RUN mkdir -p /feeOracle
WORKDIR /app/feeOracle
COPY . .
RUN go mod download
RUN go build -o feeOracle ./
EXPOSE 8091 8091
RUN ["./feeOracle", "key-generate"]
CMD ["./feeOracle", "server", "-c", "./config.yaml", "-d", "./domain.json", "-k", "./keyfile.priv", "-t", "secp256k1"]
