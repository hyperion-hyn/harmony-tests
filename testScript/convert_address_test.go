package testScript

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperion-hyn/hyperion-tf/extension/go-sdk/pkg/address"
	"testing"
)

func TestConvertAddress(t *testing.T) {

	bech32Address := "hyn15h6ankpr645j6u6vjcvsjq8akx0svc4zxcstxf"

	fmt.Printf("after convert address:%s", address.Parse(bech32Address).Hex())

}

func TestConvertToBechAddress(t *testing.T) {

	ethAddress := "0xda07bfa929d34b352851ea835db2b6c9f5026e4c"

	fmt.Printf("after convert address:%s \n", address.ToBech32(common.HexToAddress(ethAddress)))

	fmt.Printf("after convert address:%s \n", common.MustAddressToBech32(common.HexToAddress(ethAddress)))

}
