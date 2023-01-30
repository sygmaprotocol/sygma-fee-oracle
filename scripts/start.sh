#!/usr/bin/env bash
# The Licensed Work is (c) 2022 Sygma
# SPDX-License-Identifier: BUSL-1.1


dir=${PWD}

# only run this script when doing end to end testing
sygma-fee-oracle server -c $dir/e2e/config/config.yaml -d $dir/domain.json -k $dir/e2e/keystore/test_keyfile.priv -t secp256k1 >> $dir/feeOracle.log 2>&1 &

sleep 5
