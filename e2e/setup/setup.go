// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package setup

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/scripts/e2e_test/bridge"
	ERC20Handler "github.com/ChainSafe/chainbridge-fee-oracle/scripts/e2e_test/erc20Handler"
	ERC20PresetMinterPauser "github.com/ChainSafe/chainbridge-fee-oracle/scripts/e2e_test/erc20PresetMinterPauser"
	"github.com/ChainSafe/chainbridge-fee-oracle/scripts/e2e_test/feeHandler"
	"github.com/ChainSafe/chainbridge-fee-oracle/store/db"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	FromDomainId       = 1
	ToDomainID         = 2
	gasUsed            = 100000
	feePercent         = 500
	FeeOracleHexPriKey = "6937d1d0b52f2fa7f4e071c7e64934ad988a8f21c6bf4f323fc19af4c77e3c5e" // 0x74d2946319bEEe4A140068eb83F9ee3a90B06F4f
	FeeOracleAddress   = common.HexToAddress("0x74d2946319bEEe4A140068eb83F9ee3a90B06F4f")
	SenderAddress      = common.HexToAddress("0x74d2946319bEEe4A140068eb83F9ee3a90B06F4f")
)

type ContractsSetupResp struct {
	Auth                            *bind.TransactOpts
	BridgeInstance                  *bridge.Bridge
	BridgeAddress                   common.Address
	FeeHandlerInstance              *feeHandler.FeeHandler
	FeeHandlerAddress               common.Address
	ERC20PresetMinterPauserInstance *ERC20PresetMinterPauser.ERC20PresetMinterPauser
	ERC20PresetMinterPauserAddress  common.Address
	ERC20HandlerInstance            *ERC20Handler.ERC20Handler
	ERC20HandlerAddress             common.Address
}

func ContractsSetup() *ContractsSetupResp {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	auth := getAccountAuth(client, FeeOracleHexPriKey)

	// preparation - deploying contracts
	bridgeAddress, _, bridgeInstance, err := bridge.DeployBridge(auth, client, uint8(FromDomainId))
	if err != nil {
		panic(err)
	}
	feeHandlerAddress, _, feeHandlerInstance, err := feeHandler.DeployFeeHandler(IncreaseNonce(auth), client, bridgeAddress)
	if err != nil {
		panic(err)
	}
	_, err = feeHandlerInstance.SetFeeOracle(IncreaseNonce(auth), FeeOracleAddress)
	if err != nil {
		panic(err)
	}
	registeredFeeOracle, err := feeHandlerInstance.OracleAddress(&bind.CallOpts{From: auth.From})
	if err != nil {
		panic(err)
	}
	if registeredFeeOracle != FeeOracleAddress {
		panic("fee oracle address does not match with contract")
	}
	_, err = feeHandlerInstance.SetFeeProperties(IncreaseNonce(auth), uint32(gasUsed), uint16(feePercent))
	if err != nil {
		panic(err)
	}
	resourceTokenAddress, _, _, err := ERC20PresetMinterPauser.DeployERC20PresetMinterPauser(IncreaseNonce(auth), client, "token", "TOK")
	if err != nil {
		panic(err)
	}
	erc20HandlerAddress, _, _, err := ERC20Handler.DeployERC20Handler(IncreaseNonce(auth), client, bridgeAddress)
	if err != nil {
		panic(err)
	}

	fmt.Println("bridgeAddress:", bridgeAddress.String())
	fmt.Println("feeHandlerAddress:", feeHandlerAddress.String())
	fmt.Println("resourceTokenAddress:", resourceTokenAddress.String())
	fmt.Println("erc20HandlerAddress:", erc20HandlerAddress.String())

	return &ContractsSetupResp{
		Auth:                           auth,
		BridgeInstance:                 bridgeInstance,
		BridgeAddress:                  bridgeAddress,
		FeeHandlerInstance:             feeHandlerInstance,
		FeeHandlerAddress:              feeHandlerAddress,
		ERC20PresetMinterPauserAddress: resourceTokenAddress,
		ERC20HandlerAddress:            erc20HandlerAddress,
	}
}

func getAccountAuth(client *ethclient.Client, hexKey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = big.NewInt(20000000000)

	return auth
}

func IncreaseNonce(auth *bind.TransactOpts) *bind.TransactOpts {
	auth.Nonce = big.NewInt(0).Add(auth.Nonce, big.NewInt(1))
	return auth
}

func DataPrepare(path string) error {
	database, err := db.NewLvlDB(path)
	if err != nil {
		return err
	}

	dataKey1 := bytes.Buffer{}
	dataKey1.WriteString(fmt.Sprintf("%s%s:%s:%s", "conversionrate:", "coinmarketcap", "matic", "eth"))
	dataValue1 := &types.ConversionRate{
		Base:       "matic",
		Foreign:    "eth",
		Rate:       0.000445,
		OracleName: "coinmarketcap",
		Time:       time.Now().Unix(),
	}
	data1, err := json.Marshal(dataValue1)
	if err != nil {
		return err
	}
	err = database.Set(dataKey1.Bytes(), data1)
	if err != nil {
		return err
	}

	dataKey2 := bytes.Buffer{}
	dataKey2.WriteString(fmt.Sprintf("%s%s:%s:%s", "conversionrate:", "coinmarketcap", "matic", "doge"))
	dataValue2 := &types.ConversionRate{
		Base:       "matic",
		Foreign:    "doge",
		Rate:       8.948864,
		OracleName: "coinmarketcap",
		Time:       time.Now().Unix(),
	}
	data2, err := json.Marshal(dataValue2)
	if err != nil {
		return err
	}
	err = database.Set(dataKey2.Bytes(), data2)
	if err != nil {
		return err
	}

	dataKey3 := bytes.Buffer{}
	dataKey3.WriteString(fmt.Sprintf("%s%s:%s", "gasprice:", "polygonscan", "polygon"))
	dataValue3 := &types.GasPrices{
		SafeGasPrice:    "9000000000",
		ProposeGasPrice: "0",
		FastGasPrice:    "0",
		OracleName:      "polygonscan",
		DomainName:      "polygon",
		Time:            time.Now().Unix(),
	}
	data3, err := json.Marshal(dataValue3)
	if err != nil {
		return err
	}
	err = database.Set(dataKey3.Bytes(), data3)
	if err != nil {
		return err
	}

	return nil
}
