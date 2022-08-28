package wallet

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	// map[appid:123123 call_url:http://127.0.1.1:10002/recharge/call coin_symbol:USDT member_id:1 network_name:ERC20 secret_key:123123]
	// map[appid:123123 call_url:http://127.0.1.1:10002/recharge/call coin_symbol:USDT member_id:1 network_name:ERC20 secret_key:123123 sign:f389a273840e46f2a64a07cb0c9beb44]

	data := make(map[string]interface{})

	data["address"] = "0xBfcb4849c722BA0277F75887A14F905dDd94af28"
	data["amount"] = "10000000000000000"
	data["appid"] = "123123"
	data["block_hash"] = "0xf390e8adc7bdfda08b4ede531bcbef8b15b9a8243197323198f3b2444dfa7d3f"
	data["block_high"] = "22288318"
	data["business_id"] = "0xd4d3c0cac4b685839bd465ddcbde3fb4be3093532cd9d117c97e601bc997c67b37"
	data["coin_symbol"] = "USDT"
	data["decimals"] = 0
	data["network_name"] = "ERC20"
	data["secret_key"] = "123123"
	data["status"] = 2
	data["txid"] = "0xd4d3c0cac4b685839bd465ddcbde3fb4be3093532cd9d117c97e601bc997c67b"

	str := Sign(data)

	fmt.Println(str)
}

func TestVerifySign(t *testing.T) {

	data := make(map[string]interface{})

	data["address"] = "0xBfcb4849c722BA0277F75887A14F905dDd94af28"
	data["amount"] = "10000000000000000"
	data["appid"] = "123123"
	data["secret_key"] = "123123"
	data["block_hash"] = "0xf390e8adc7bdfda08b4ede531bcbef8b15b9a8243197323198f3b2444dfa7d3f"
	data["block_high"] = "22288318"
	data["business_id"] = "0xd4d3c0cac4b685839bd465ddcbde3fb4be3093532cd9d117c97e601bc997c67b37"
	data["coin_symbol"] = "USDT"
	data["decimals"] = 0
	data["network_name"] = "ERC20"
	data["status"] = 2
	data["txid"] = "0xd4d3c0cac4b685839bd465ddcbde3fb4be3093532cd9d117c97e601bc997c67b"
	data["sign"] = "2be0c9a0ef80d2fc63b3e66fcd54c6a9"

	b := VerifySign(data)

	fmt.Println(b)

}
