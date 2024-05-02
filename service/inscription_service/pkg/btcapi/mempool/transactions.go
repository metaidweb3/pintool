package mempool

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"self-tool/tool"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func (c *MempoolClient) GetRawTransaction(txHash *chainhash.Hash) (*wire.MsgTx, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/tx/%s/raw", txHash.String()), nil)
	if err != nil {
		return nil, err
	}

	tx := wire.NewMsgTx(wire.TxVersion)
	if err := tx.Deserialize(bytes.NewReader(res)); err != nil {
		return nil, err
	}
	return tx, nil
}

type PostRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func postApi(txStr string) error {
	config, err := tool.ReadJSON("./setting.json")
	if err != nil {
		return err
	}
	postUrl := config["api"].(string) + "/btc/broadcastTx"
	body := []byte(`{
		"rawTx": "` + txStr + `"
	}`)
	r, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	post := &PostRes{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		return err
	}

	if post.Code != 2000 {
		fmt.Println("---------------")
		return errors.New(post.Msg)
	}
	return nil
}

func (c *MempoolClient) BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, err
	}
	txStr := hex.EncodeToString(buf.Bytes())
	fmt.Printf("tx hex: %s\n", txStr)
	txHash := tx.TxHash()
	err := postApi(txStr)
	return &txHash, err

	//res, err := c.request(http.MethodPost, "/tx", strings.NewReader(hex.EncodeToString(buf.Bytes())))
	//if err != nil {
	//	return nil, err
	//}
	//
	//txHash, err := chainhash.NewHashFromStr(string(res))
	//if err != nil {
	//	return nil, errors.Wrap(err, fmt.Sprintf("failed to parse tx hash, %s", string(res)))
	//}
	//return txHash, nil
}
