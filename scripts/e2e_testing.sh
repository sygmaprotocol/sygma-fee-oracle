#!/usr/bin/env bash

set -e

# fee handler solidity code branch name
solidity_branch=$1

E2E_DIR=${PWD}/e2e
SCRIPT_DIR=${PWD}/scripts
CONTRACT_DIR=$SCRIPT_DIR/e2e_test/chainbridge-solidity/contracts/handlers/fee
E2E_TEST_DIR=$SCRIPT_DIR/e2e_test
SOLIDITY_DIR=$SCRIPT_DIR/e2e_test/chainbridge-solidity

mkdir -p $E2E_TEST_DIR
mkdir -p $E2E_TEST_DIR/bridge
mkdir -p $E2E_TEST_DIR/erc20Handler
mkdir -p $E2E_TEST_DIR/erc20PresetMinterPauser
mkdir -p $E2E_TEST_DIR/feeHandler
mkdir -p $E2E_TEST_DIR/basicFeeHandler
mkdir -p $E2E_TEST_DIR/centrifugeAsset
mkdir -p $E2E_TEST_DIR/erc721Handler
mkdir -p $E2E_TEST_DIR/erc721MinterBurnerPauser
mkdir -p $E2E_TEST_DIR/genericHandler

# clone solidity code and checkout target branch
git clone --branch $solidity_branch https://github.com/ChainSafe/chainbridge-solidity.git $E2E_TEST_DIR/chainbridge-solidity

# install chainbridge-solidity
npm i --prefix $SOLIDITY_DIR

# install solc with correct version
npm i -g solc@0.8.11

# generate abi
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $CONTRACT_DIR/FeeHandlerWithOracle.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/contracts/Bridge.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/contracts/handlers/ERC20Handler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $CONTRACT_DIR/BasicFeeHandler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/contracts/handlers/ERC721Handler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/contracts/handlers/GenericHandler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/contracts/ERC721MinterBurnerPauser.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --abi $SOLIDITY_DIR/contracts/CentrifugeAsset.sol -o $SOLIDITY_DIR/tmp

# generate bin
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $CONTRACT_DIR/FeeHandlerWithOracle.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/contracts/Bridge.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/contracts/handlers/ERC20Handler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $CONTRACT_DIR/BasicFeeHandler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/contracts/handlers/ERC721Handler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/contracts/handlers/GenericHandler.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/contracts/ERC721MinterBurnerPauser.sol -o $SOLIDITY_DIR/tmp
solcjs --include-path $SOLIDITY_DIR/node_modules/ --base-path $SOLIDITY_DIR --bin $SOLIDITY_DIR/contracts/CentrifugeAsset.sol -o $SOLIDITY_DIR/tmp

# download and extract abigen
ABIGEN_EXISTS=false
if [ -x "$(command -v abigen)" ]; then
  ABIGEN_EXISTS=true
  echo "abigen found, skipping install, make sure you have v1.10.17 installed"
else
  unameOut="$(uname -s)"
  case "${unameOut}" in
  Linux*)
    echo "Found linux machine, downloading..."
    curl -o $E2E_TEST_DIR/geth-tool-pack.tar.gz https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-1.10.17-25c9b49f.tar.gz
    tar -zxvf $E2E_TEST_DIR/geth-tool-pack.tar.gz -C $E2E_TEST_DIR --strip-components=1 geth-alltools-linux-amd64-1.10.17-25c9b49f/abigen
    ;;
  Darwin*)
    echo "Found macOS machine, downloading..."
    curl -o $E2E_TEST_DIR/geth-tool-pack.tar.gz https://gethstore.blob.core.windows.net/builds/geth-alltools-darwin-amd64-1.10.17-25c9b49f.tar.gz
    tar -zxvf $E2E_TEST_DIR/geth-tool-pack.tar.gz -C $E2E_TEST_DIR --strip-components=1 ./geth-alltools-darwin-amd64-1.10.17-25c9b49f/abigen
    ;;
  *)
    echo "Operating system not supported, please manually install: https://geth.ethereum.org/downloads"
    exit 1
    ;;
  esac
fi

# generate go binding files
if [ $ABIGEN_EXISTS = false ]; then
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_fee_FeeHandlerWithOracle_sol_FeeHandlerWithOracle.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_fee_FeeHandlerWithOracle_sol_FeeHandlerWithOracle.abi \
    --pkg=feeHandler --out=$E2E_TEST_DIR/feeHandler/feeHandler.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_fee_BasicFeeHandler_sol_BasicFeeHandler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_fee_BasicFeeHandler_sol_BasicFeeHandler.abi \
    --pkg=basicFeeHandler --out=$E2E_TEST_DIR/basicFeeHandler/basicFeeHandler.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_Bridge_sol_Bridge.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_Bridge_sol_Bridge.abi \
    --pkg=bridge --out=$E2E_TEST_DIR/bridge/bridge.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/@openzeppelin_contracts_token_ERC20_presets_ERC20PresetMinterPauser_sol_ERC20PresetMinterPauser.bin \
    --abi=$SOLIDITY_DIR/tmp/@openzeppelin_contracts_token_ERC20_presets_ERC20PresetMinterPauser_sol_ERC20PresetMinterPauser.abi \
    --pkg=ERC20PresetMinterPauser --out=$E2E_TEST_DIR/erc20PresetMinterPauser/ERC20PresetMinterPauser.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_ERC20Handler_sol_ERC20Handler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_ERC20Handler_sol_ERC20Handler.abi \
    --pkg=ERC20Handler --out=$E2E_TEST_DIR/erc20Handler/ERC20Handler.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_ERC721Handler_sol_ERC721Handler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_ERC721Handler_sol_ERC721Handler.abi \
    --pkg=ERC721Handler --out=$E2E_TEST_DIR/erc721Handler/ERC721Handler.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_GenericHandler_sol_GenericHandler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_GenericHandler_sol_GenericHandler.abi \
    --pkg=GenericHandler --out=$E2E_TEST_DIR/genericHandler/GenericHandler.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_ERC721MinterBurnerPauser_sol_ERC721MinterBurnerPauser.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_ERC721MinterBurnerPauser_sol_ERC721MinterBurnerPauser.abi \
    --pkg=ERC721MinterBurnerPauser --out=$E2E_TEST_DIR/erc721MinterBurnerPauser/ERC721721MinterBurnerPauser.go
  $E2E_TEST_DIR/abigen --bin=$SOLIDITY_DIR/tmp/contracts_CentrifugeAsset_sol_CentrifugeAsset.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_CentrifugeAsset_sol_CentrifugeAsset.abi \
    --pkg=CentrifugeAsset --out=$E2E_TEST_DIR/centrifugeAsset/CentrifugeAsset.go
else
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_fee_FeeHandlerWithOracle_sol_FeeHandlerWithOracle.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_fee_FeeHandlerWithOracle_sol_FeeHandlerWithOracle.abi \
    --pkg=feeHandler --out=$E2E_TEST_DIR/feeHandler/feeHandler.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_fee_BasicFeeHandler_sol_BasicFeeHandler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_fee_BasicFeeHandler_sol_BasicFeeHandler.abi \
    --pkg=basicFeeHandler --out=$E2E_TEST_DIR/basicFeeHandler/basicFeeHandler.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_Bridge_sol_Bridge.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_Bridge_sol_Bridge.abi \
    --pkg=bridge --out=$E2E_TEST_DIR/bridge/bridge.go
  abigen --bin=$SOLIDITY_DIR/tmp/@openzeppelin_contracts_token_ERC20_presets_ERC20PresetMinterPauser_sol_ERC20PresetMinterPauser.bin \
    --abi=$SOLIDITY_DIR/tmp/@openzeppelin_contracts_token_ERC20_presets_ERC20PresetMinterPauser_sol_ERC20PresetMinterPauser.abi \
    --pkg=ERC20PresetMinterPauser --out=$E2E_TEST_DIR/erc20PresetMinterPauser/ERC20PresetMinterPauser.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_ERC20Handler_sol_ERC20Handler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_ERC20Handler_sol_ERC20Handler.abi \
    --pkg=ERC20Handler --out=$E2E_TEST_DIR/erc20Handler/ERC20Handler.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_ERC721Handler_sol_ERC721Handler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_ERC721Handler_sol_ERC721Handler.abi \
    --pkg=ERC721Handler --out=$E2E_TEST_DIR/erc721Handler/ERC721Handler.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_handlers_GenericHandler_sol_GenericHandler.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_handlers_GenericHandler_sol_GenericHandler.abi \
    --pkg=GenericHandler --out=$E2E_TEST_DIR/genericHandler/GenericHandler.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_ERC721MinterBurnerPauser_sol_ERC721MinterBurnerPauser.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_ERC721MinterBurnerPauser_sol_ERC721MinterBurnerPauser.abi \
    --pkg=ERC721MinterBurnerPauser --out=$E2E_TEST_DIR/erc721MinterBurnerPauser/ERC721MinterBurnerPauser.go
  abigen --bin=$SOLIDITY_DIR/tmp/contracts_CentrifugeAsset_sol_CentrifugeAsset.bin \
    --abi=$SOLIDITY_DIR/tmp/contracts_CentrifugeAsset_sol_CentrifugeAsset.abi \
    --pkg=CentrifugeAsset --out=$E2E_TEST_DIR/centrifugeAsset/CentrifugeAsset.go
fi

# run e2e test
go test e2e/endpoint_test.go

# clean up
echo "testing is done, cleaning up..."

gPid=`pgrep -f "ganache-cli"`
if [ "$gPid" ]
then
  echo "terminating ganache-cli"
  kill $gPid
fi

foPid=`pgrep -f "chainbridge-fee-oracle"`
if [ "$foPid" ]
then
  echo "terminating fee oracle service"
  kill $foPid
fi

rm -rf $SOLIDITY_DIR
rm -rf $E2E_DIR/testdb
rm -f ./feeOracle.log
rm -f $E2E_TEST_DIR/abigen
rm -f $E2E_TEST_DIR/geth-tool-pack.tar.gz
