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

func TestGetValidatorInformation(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	validatorAddress := common.HexToAddress("0xbF45A07ca054535Aafd703e47D4a82A5eB1f89DD")

	validator, err := client.GetValidatorInformation(context.Background(), validatorAddress, nil)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(fmt.Sprintf("got validator address :%s", validator.Validator.ValidatorAddress.Hex()))

	for _, redelegation := range validator.Redelegations {
		fmt.Println(fmt.Sprintf("user: %s,amount: %s", redelegation.DelegatorAddress.Hex(), new(big.Int).Quo(redelegation.Amount, big.NewInt(params.Ether))))
	}

}
