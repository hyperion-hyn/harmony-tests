package testScript

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGetAllMap3NodeAddresses(t *testing.T) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	address, err := client.GetAllMap3NodeAddresses(context.Background(), nil)

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("got address :%s", address)

}
