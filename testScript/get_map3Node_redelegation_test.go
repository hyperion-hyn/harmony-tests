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

func TestGetMap3NodeDelegation(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Printf("%v", err)
	}

	validatorAddress := common.HexToAddress("0x3D62D5692AE15Bbcc70a0C47A0cEEc1C19Cd410d")
	delegatorAddress := common.HexToAddress("0x07FF12833cC5CCb2514cF73AdfF64Bd1dAB12cd5")

	delegation, err := client.GetMap3NodeDelegation(context.Background(), validatorAddress, delegatorAddress, nil)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(fmt.Sprintf("user: %s,amount: %s,reward: %s",
		delegation.DelegatorAddress.Hex(),
		new(big.Int).Quo(delegation.Amount, big.NewInt(params.Ether)),
		new(big.Int).Quo(delegation.Reward, big.NewInt(params.Ether))),
	)

}
