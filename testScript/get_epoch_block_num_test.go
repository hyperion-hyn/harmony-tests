package testScript

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGetEpochFirstBlockNum(t *testing.T) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	blockNum, err := client.GetEpochFirstBlockNum(context.Background(), 2)

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("got epoch first blockNum :%d", blockNum)

}

func TestGetEpochLastBlockNum(t *testing.T) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	blockNum, err := client.GetEpochLastBlockNum(context.Background(), 2)

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("got epoch last blockNum :%d", blockNum)

}
