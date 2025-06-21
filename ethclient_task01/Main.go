package main

import (
	"context"
	"crypto/ecdsa"
	"ethclient_task01/counter"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(sepolia_url)
	if err != nil {
		log.Fatal(err)
	}
	// Task01()
	Task02(client)
}

const (
	sepolia_url = "https://eth-sepolia.g.alchemy.com/v2/Z71Xlftwd4OCqQ68eV6z8xwakHQ89PS2"
)

// 任务一：区块链读写（查询区块、发送交易）
func Task01(client *ethclient.Client) {
	QueryBlock(client)
	SendTransaction(client)

}
func Task02(client *ethclient.Client) {
	// 先通过abigen 工具部署计数器合约
	DeployCounterContract(client)
	InvokeCounterContract(client)

}

// 调用下面部署的合约
func InvokeCounterContract(client *ethclient.Client) {
	// 合约地址
	contrractAddress := "0x8c5D902B1d903C5407bE59017B52941a5AAbD90a"
	// 加载合约
	counterContract, err := counter.NewCounter(common.HexToAddress(contrractAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// 调用增加函数
	tx, err := counterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Increment transaction hash:", tx.Hash().Hex())

	// 查询count 当前值
	countVal, err := counterContract.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count current value :", countVal.Uint64())

	// 调用减少 函数
	tx, err = counterContract.Decrement(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrement transaction hash:", tx.Hash().Hex())
	// 查询count 当前值
	countVal, err = counterContract.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count current value :", countVal.Uint64())
	// 调用 reset 函数
	tx, err = counterContract.Reset(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reset transaction hash:", tx.Hash().Hex())
	// 查询count 当前值
	countVal, err = counterContract.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count current value :", countVal.Uint64())
}

// 部署合约
func DeployCounterContract(client *ethclient.Client) {
	// 加载私钥
	privateKey, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥
	publicKey := privateKey.Public()
	// 断言
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot asset type of publicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 创建交易
	txOpt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	txOpt.Nonce = big.NewInt(int64(nonce))
	txOpt.GasPrice = gasPrice
	txOpt.GasLimit = uint64(300000)
	txOpt.Value = big.NewInt(0)

	// 部署合约
	contracAddress, tx, instance, err := counter.DeployCounter(txOpt, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("counter contract address :", contracAddress.Hex())
	fmt.Println("transaction hash：", tx.Hash().Hex())

	countValue, err := instance.GetCount(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count default value: ", countValue)
}

// 发送交易
func SendTransaction(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot asset type of publicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x491729AbA2a8A12d90D0Bf7B82d06638203012AF")
	ammount := big.NewInt(10000000000) // 单位 wei
	gasLimit := uint64(300000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte

	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, ammount, gasLimit, gasPrice, data)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 加签
	singedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 发送交易
	err = client.SendTransaction(context.Background(), singedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transaction hash = ", singedTx.Hash().Hex())
}

// 查询区块信息
func QueryBlock(client *ethclient.Client) {

	// 1、根据区块 hash 查询
	// bolckHash := common.HexToHash("0x41695f202e2484825231bf8e8ef34fe09e42689fda81a1311927f4fed8dce065")
	// block, err := client.BlockByHash(context.Background(), bolckHash)

	// 2、根据blockNumber 查询
	blockNumber := big.NewInt(8596091)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blockHash: ", block.Hash().Hex())
	fmt.Println("blockNumber: ", block.Number().Uint64())
	fmt.Println("time: ", block.Time())
	fmt.Println("gassLimit: ", block.GasLimit())
	fmt.Println("gassUsed: ", block.GasUsed())
	fmt.Println("transations count: ", len(block.Transactions()))

	txCount, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transations count: ", txCount)
}
