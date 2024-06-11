package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

// 合约地址和 ABI（请替换为实际的 ABI 和合约地址）
const contractAddress = "0xD06CBe0ec2138c7aAFA8eAB031EA164f5c1C6bC1"

var contractABI = `[{"inputs":[{"components":[{"internalType":"address","name":"feeRecipient","type":"address"},{"internalType":"uint256","name":"fastTradeFeeBps","type":"uint256"},{"internalType":"uint256","name":"sniperFeeBps","type":"uint256"},{"internalType":"uint256","name":"limitFeeBps","type":"uint256"},{"internalType":"uint256","name":"feeBaseBps","type":"uint256"},{"internalType":"address","name":"permit2","type":"address"},{"internalType":"address","name":"weth9","type":"address"},{"internalType":"address","name":"v2Factory","type":"address"},{"internalType":"address","name":"v3Factory","type":"address"},{"internalType":"bytes32","name":"pairInitCodeHash","type":"bytes32"},{"internalType":"bytes32","name":"poolInitCodeHash","type":"bytes32"}],"internalType":"struct RouterParameters","name":"params","type":"tuple"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"BalanceTooLow","type":"error"},{"inputs":[],"name":"BuyPunkFailed","type":"error"},{"inputs":[],"name":"ContractLocked","type":"error"},{"inputs":[],"name":"ETHNotAccepted","type":"error"},{"inputs":[{"internalType":"uint256","name":"commandIndex","type":"uint256"},{"internalType":"bytes","name":"message","type":"bytes"}],"name":"ExecutionFailed","type":"error"},{"inputs":[],"name":"FeeRecipientAddressCannotBeZeroAddress","type":"error"},{"inputs":[],"name":"FeeRecipientAddressCannotBeZeroAddress1","type":"error"},{"inputs":[],"name":"FromAddressIsNotOwner","type":"error"},{"inputs":[],"name":"InsufficientETH","type":"error"},{"inputs":[],"name":"InsufficientToken","type":"error"},{"inputs":[],"name":"InvalidBips","type":"error"},{"inputs":[{"internalType":"uint256","name":"commandType","type":"uint256"}],"name":"InvalidCommandType","type":"error"},{"inputs":[{"internalType":"uint256","name":"fastTradeFeeBps","type":"uint256"}],"name":"InvalidFastTradeFeeBps","type":"error"},{"inputs":[{"internalType":"uint256","name":"feeBase","type":"uint256"}],"name":"InvalidFeeBase","type":"error"},{"inputs":[{"internalType":"uint256","name":"feeType","type":"uint256"}],"name":"InvalidFeeType","type":"error"},{"inputs":[{"internalType":"uint256","name":"limitFeeBps","type":"uint256"}],"name":"InvalidLimitFeeBps","type":"error"},{"inputs":[],"name":"InvalidPath","type":"error"},{"inputs":[],"name":"InvalidReserves","type":"error"},{"inputs":[{"internalType":"uint256","name":"sniperFeeBps","type":"uint256"}],"name":"InvalidSniperFeeBps","type":"error"},{"inputs":[],"name":"InvalidSpender","type":"error"},{"inputs":[],"name":"LengthMismatch","type":"error"},{"inputs":[],"name":"SliceOutOfBounds","type":"error"},{"inputs":[],"name":"TransactionDeadlinePassed","type":"error"},{"inputs":[],"name":"UnsafeCast","type":"error"},{"inputs":[],"name":"V2InvalidPath","type":"error"},{"inputs":[],"name":"V2TooLittleReceived","type":"error"},{"inputs":[],"name":"V2TooMuchRequested","type":"error"},{"inputs":[],"name":"V3InvalidAmountOut","type":"error"},{"inputs":[],"name":"V3InvalidCaller","type":"error"},{"inputs":[],"name":"V3InvalidSwap","type":"error"},{"inputs":[],"name":"V3TooLittleReceived","type":"error"},{"inputs":[],"name":"V3TooMuchRequested","type":"error"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"msgSender","type":"address"},{"indexed":false,"internalType":"uint256","name":"fastTradeFeeBps","type":"uint256"}],"name":"FastTradeFeeBpsUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"msgSender","type":"address"},{"indexed":false,"internalType":"uint256","name":"feeBaseBps","type":"uint256"}],"name":"FeeBaseBpsUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"msgSender","type":"address"},{"indexed":false,"internalType":"address","name":"feeRecipient","type":"address"}],"name":"FeeRecipientUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"msgSender","type":"address"},{"indexed":false,"internalType":"uint256","name":"limitFeeBps","type":"uint256"}],"name":"LimitFeeBpsUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"msgSender","type":"address"},{"indexed":false,"internalType":"uint256","name":"sniperFeeBps","type":"uint256"}],"name":"SniperFeeBpsUpdated","type":"event"},{"inputs":[{"internalType":"bytes","name":"commands","type":"bytes"},{"internalType":"bytes[]","name":"inputs","type":"bytes[]"}],"name":"execute","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"bytes","name":"commands","type":"bytes"},{"internalType":"bytes[]","name":"inputs","type":"bytes[]"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"execute","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"fastTradeFeeBps","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"feeBaseBps","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"feeRecipient","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"limitFeeBps","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"feeBps","type":"uint256"}],"name":"setFastTradeFeeBps","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"feeBaseBps","type":"uint256"}],"name":"setFeeBaseBps","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"feeRecipient","type":"address"}],"name":"setFeeRecipient","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"feeBps","type":"uint256"}],"name":"setLimitFeeBps","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"feeBps","type":"uint256"}],"name":"setSniperFeeBps","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"sniperFeeBps","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"int256","name":"amount0Delta","type":"int256"},{"internalType":"int256","name":"amount1Delta","type":"int256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"uniswapV3SwapCallback","outputs":[],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]`

func main() {
	// 连接到本地以太坊节点
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 创建合约实例
	contractAddress := common.HexToAddress(contractAddress)
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}
	contract := bind.NewBoundContract(contractAddress, parsedABI, client, client, client)

	// 调用 get 方法读取数据
	var result *[]interface{}
	err = contract.Call(nil, result, "feeBaseBps")
	if err != nil {
		log.Fatalf("Failed to call contract method: %v", err)
	}
	fmt.Printf("Stored value: %v\n", result)
}