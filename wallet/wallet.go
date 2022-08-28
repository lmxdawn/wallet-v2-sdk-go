package wallet

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	appid     string
	secretKey string
	url       string
	client    *http.Client
}

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateWalletRes struct {
	Res
	Data CreateWalletDataRes `json:"data"` // 创建钱包的地址
}

type CreateWalletDataRes struct {
	Address string `json:"address"` // 钱包地址
}

type CreateWithdrawRes struct {
	Res
	Data CreateWithdrawDataRes `json:"data"` // 创建钱包的地址
}

type CreateWithdrawDataRes struct {
	From string `json:"from"` // 发起转账的地址
	Hash string `json:"hash"` // 哈希
}

// NewWalletClient 创建
func NewWalletClient(appid string, secretKey string, url string) *Client {

	return &Client{
		appid,
		secretKey,
		url,
		&http.Client{
			Timeout: time.Millisecond * time.Duration(10*1000),
		},
	}
}

// CreateWallet 创建钱包
func (w *Client) CreateWallet(memberId string, networkName string, coinSymbol string, callUrl string) (*CreateWalletDataRes, error) {

	data := make(map[string]interface{})
	data["network_name"] = networkName
	data["coin_symbol"] = coinSymbol
	data["member_id"] = memberId
	data["call_url"] = callUrl

	var res CreateWalletRes

	err := w.post("/createWallet", data, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return &res.Data, nil
}

// CreateWithdraw 生成提现
func (w *Client) CreateWithdraw(networkName string, coinSymbol string, address string, amount string, businessId, callUrl string) (*CreateWithdrawDataRes, error) {

	data := make(map[string]interface{})
	data["network_name"] = networkName
	data["coin_symbol"] = coinSymbol
	data["address"] = address
	data["amount"] = amount
	data["business_id"] = businessId
	data["call_url"] = callUrl

	var res CreateWithdrawRes

	err := w.post("/createWithdraw", data, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	return &res.Data, nil
}

// get 请求
func (w *Client) get(path string, params url.Values, res interface{}) error {

	return w.request(w.url+path, http.MethodGet, nil, params, nil, res)

}

// post 请求
func (w *Client) post(path string, data map[string]interface{}, res interface{}) error {
	return w.request(w.url+path, http.MethodPost, nil, nil, data, res)
}

// post 请求
func (w *Client) request(urlStr string, method string, header http.Header, params url.Values, data map[string]interface{}, res interface{}) error {
	var reqBody io.Reader
	if data != nil {
		data["appid"] = w.appid
		data["secret_key"] = w.secretKey
		// 签名
		sign := Sign(data)
		data["sign"] = sign
		delete(data, "secret_key")

		bytesData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(bytesData)
	}

	if params != nil {
		Url, err := url.Parse(urlStr)
		if err != nil {
			return err
		}
		//如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlStr = Url.String()
	}

	req, err := http.NewRequest(method, urlStr, reqBody)
	if header != nil {
		req.Header = header
	}
	if err != nil {
		// handle error
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := w.client.Do(req)
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return err
	}

	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}

	return nil
}
