package testScript

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"testing"
)

func TestGetMap3NodeInformation(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	validatorAddress := common.HexToAddress("0x699041EDf0babd30Dd206e88FBEd3f74e1A5E272")

	map3NodeWrapper, err := client.GetMap3NodeInformation(context.Background(), validatorAddress, nil)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(fmt.Sprintf("got map3NodeWrapper address :%s", map3NodeWrapper.Map3Node.Map3Address.Hex()))

	for _, redelegation := range map3NodeWrapper.Microdelegations {
		fmt.Println(fmt.Sprintf("user: %s,amount: %s", redelegation.DelegatorAddress.Hex(), new(big.Int).Quo(redelegation.Amount, big.NewInt(params.Ether))))
	}

}
