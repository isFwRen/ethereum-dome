package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0xb199B02A8eE9610e257574853857a0fD577fd67A")
	fmt.Println("address:", address.Hex())
	fmt.Println("address:", address.Bytes())
}
