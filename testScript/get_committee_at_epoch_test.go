package testScript

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGetCommitteeAtEpoch(t *testing.T) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	address, err := client.GetCommitteeAtEpoch(context.Background(), 6)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("got address length :%d", len(address))

	for _, addressTemp := range address {
		fmt.Printf("got address :%s", addressTemp.Hex())
	}

}
