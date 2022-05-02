#!/usr/bin/env bash

dir=${PWD}

# only run this script when doing end to end testing
chainbridge-fee-oracle server -c $dir/e2e/config/config.yaml -d $dir/domain.json -r $dir/resource.json -k $dir/e2e/keystore/test_keyfile.priv -t secp256k1 >> $dir/feeOracle.log 2>&1 &

sleep 5
