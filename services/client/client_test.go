package client

import (
	"fmt"
	"testing"
)

func TestSupportChain(t *testing.T) {
	client := NewWalletClient("http://127.0.0.1:8970")
	result, err := client.GetSupportCoins("Bitcoin", "MainNet")
	if err != nil {
		fmt.Println("get support chain fail")
		return
	}
	fmt.Println("SupportChain result : ", result)
}

func TestWalletAddress(t *testing.T) {
	client := NewWalletClient("http://127.0.0.1:8970")
	addressInfo, err := client.GetWalletAddress("Bitcoin", "MainNet")
	if err != nil {
		fmt.Println("get wallet address fail")
		return
	}
	fmt.Println("Wallet Address Result:", addressInfo.Address, addressInfo.PublicKey)

}
