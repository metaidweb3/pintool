package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	service "self-tool/service/inscription_service"
	"self-tool/tool"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

type txRes struct {
	TxId    string `json:"txId"`
	Vout    int    `json:"vout"`
	Satoshi int64  `json:"satoshi"`
}
type utxoRes struct {
	Message string  `json:"message"`
	Result  []txRes `json:"result"`
	Status  string  `json:"status"`
}

func GetUtxoList(address string, net string) (data []string) {
	config, err := getConfig()
	if err != nil {
		return
	}
	apiUrl := config["api"].(string) + "/wallet-v1/address/btc-utxo"
	params := make(map[string]string)
	params["address"] = config["address"].(string)
	params["order"] = "desc"
	params["unconfirmed"] = "0"
	res, err := httpGet(apiUrl, params)
	if err != nil {
		return
	}
	var resData utxoRes
	err = json.Unmarshal(res, &resData)
	if err != nil {
		return
	}
	for _, v := range resData.Result {
		s := fmt.Sprintf("%s:%d,%d", v.TxId, v.Vout, v.Satoshi)
		data = append(data, s)
	}
	return
}

func TxHashformat(txHash string) string {
	if len(txHash) < 20 {
		return txHash
	}
	arr := strings.Split(txHash, ":")
	return txHash[0:5] + "..." + arr[0][len(arr[0])-5:] + ":" + arr[1]
}
func getConfig() (map[string]interface{}, error) {
	return tool.ReadJSON("./setting.json")
}
func httpGet(apiUrl string, params map[string]string) (body []byte, err error) {
	paramsValue := url.Values{}
	Url, err := url.Parse(apiUrl)
	if err != nil {
		return
	}
	for k, v := range params {
		paramsValue.Set(k, v)
	}
	Url.RawQuery = paramsValue.Encode()
	urlPath := Url.String()
	//fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	return
}

func SendPin(metaIdData service.InscriptionMetaIdData, utxoString string, shovel string, feeRate int64) error {
	if utxoString == "" {
		return errors.New("At least one utxo is required")
	}
	config, err := getConfig()
	if err != nil {
		return err
	}
	flag := "metaid"
	var netParams *chaincfg.Params
	netParams = &chaincfg.MainNetParams
	if config["net"].(string) == "Btc TestNet" {
		flag = "testid"
		netParams = &chaincfg.TestNet3Params
	}
	if config["net"].(string) == "Btc Regtest" {
		flag = "testid"
		netParams = &chaincfg.RegressionNetParams
	}
	metaIdData.MetaIDFlag = flag
	address := config["address"].(string)
	metaIdData.Destination = address
	fromPriKeyHex := config["pk"].(string)
	//feeRate := int64(2)
	outAddressType := "segwit"
	isOnlyCal := false
	revealOutValue := int64(546)
	var utxoList []*service.InscribeUtxo
	arr1 := strings.Split(utxoString, ",")
	arr2 := strings.Split(arr1[0], ":")
	idx, _ := strconv.ParseInt(arr2[1], 10, 64)
	value, _ := strconv.ParseInt(arr1[1], 10, 64)
	utxoList = append(utxoList, &service.InscribeUtxo{
		OutTx:     arr2[0],
		OutIndex:  idx,
		OutAmount: value,
	})
	if shovel != "" {
		index := strings.LastIndex(shovel, "i")
		v, _ := strconv.ParseInt(shovel[index+1:], 10, 64)
		utxoList = append(utxoList, &service.InscribeUtxo{
			OutTx:     shovel[:index],
			OutIndex:  v,
			OutAmount: 546,
		})
	}
	fmt.Println("utxoList:", utxoList)
	commitTxHash, revealTxHashStrList, inscriptionList, fees, err := service.InscribeMultiMetaIdDataFromUtxo(netParams, metaIdData, fromPriKeyHex, feeRate, address, utxoList, outAddressType, isOnlyCal, revealOutValue)
	if err != nil {
		fmt.Printf("InscribeMultiMetaIdDataFromUtxo() error = %v", err)
		return err
	}
	fmt.Printf("commitTxHash = %v", commitTxHash)
	fmt.Printf("revealTxHashStrList = %v", revealTxHashStrList)
	fmt.Printf("inscriptionList = %v", inscriptionList)
	fmt.Printf("fees = %v", fees)
	return nil
}

func BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, err
	}

	//fmt.Printf("tx1111 hex: %s\n", hex.EncodeToString(buf.Bytes()))
	txHash := tx.TxHash()
	return &txHash, nil
}
