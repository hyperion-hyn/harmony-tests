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

func TestGetValidatorRedelegation(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	validatorAddress := common.HexToAddress("0xbF45A07ca054535Aafd703e47D4a82A5eB1f89DD")
	delegatorAddress := common.HexToAddress("0x8BB7Fa9c0Dd73C8B83A7F86144655E377F6f0776")

	delegation, err := client.GetValidatorRedelegation(context.Background(), validatorAddress, delegatorAddress, nil)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(fmt.Sprintf("user: %s,amount: %s,reward: %s",
		delegation.DelegatorAddress.Hex(),
		new(big.Int).Quo(delegation.Amount, big.NewInt(params.Ether)),
		new(big.Int).Quo(delegation.Reward, big.NewInt(params.Ether))),
	)

}
