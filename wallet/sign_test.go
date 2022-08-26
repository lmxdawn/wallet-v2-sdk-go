package wallet

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	// map[appid:123123 call_url:http://127.0.1.1:10002/recharge/call coin_symbol:USDT member_id:1 network_name:ERC20 secret_key:123123]
	// map[appid:123123 call_url:http://127.0.1.1:10002/recharge/call coin_symbol:USDT member_id:1 network_name:ERC20 secret_key:123123 sign:f389a273840e46f2a64a07cb0c9beb44]

	data := make(map[string]string)
	data["appid"] = "123123"
	data["call_url"] = "http://127.0.1.1:10002/recharge/call"
	data["network_name"] = "ERC20"
	data["coin_symbol"] = "USDT"
	data["member_id"] = "1"
	data["secret_key"] = "123123"

	str := Sign(data)

	fmt.Println(str)
}

func TestVerifySign(t *testing.T) {

	data := make(map[string]string)
	data["appid"] = "123123"
	data["call_url"] = "http://127.0.1.1:10002/recharge/call"
	data["network_name"] = "ERC20"
	data["coin_symbol"] = "USDT"
	data["member_id"] = "1"
	data["secret_key"] = "123123"
	data["sign"] = "f389a273840e46f2a64a07cb0c9beb44"
	b := VerifySign(data)

	fmt.Println(b)

}
