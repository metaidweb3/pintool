package inscription_service

import (
	"self-tool/tool"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

/*
2024/01/05 14:47:23 new priviate key c84a895632c43e86cf13f70e2e4b1e6a99aed5fbf4b8ffdd45653971e8894944
2024/01/05 14:47:23 new public key 03add89c45d3c9bde8d86d33a532981f2368d8021b86c3c6fee41e82ef82dd4c38
2024/01/05 14:47:23 new legacy address mgoKYHgb463YZ91rsyLJeHkoHNiKuWKV3z
2024/01/05 14:47:23 new native segwit address tb1qpcggg9dl043nde4qle5dawvzrcwss4ekfx0awz
2024/01/05 14:47:23 new taproot address tb1pc2h33h2gfqp4s6p6fqnx5zxrkh22tnsc34gx9jwswtpe8nrp3shqvzrcmd
2024/01/05 14:47:23 restore taproot address tb1pc2h33h2gfqp4s6p6fqnx5zxrkh22tnsc34gx9jwswtpe8nrp3shqvzrcmd
*/

/*
2024/01/06 09:39:41 new priviate key 0e05df2a58d50cc8494dbefe183fadfeca0363ce073ca13ad346eb413ac0124d
2024/01/06 09:39:41 new public key 03d324e7b8db8bdad0d0883f805a0aa469195272be08bbcb85c366986299f87c55
2024/01/06 09:39:41 new legacy address mu8fk9dzAw63qpugfm8AgQHVAptf6dwhFh
2024/01/06 09:39:41 new native segwit address tb1qj4dvkex66vwddf3tvvu79vevv2v7u6prnh80yq
2024/01/06 09:39:41 new taproot address tb1phcc2al0plg9rrysftk2ngpjvfgtz5d0aagvah7lc6nnml3vx9vns2tl8w5
2024/01/06 09:39:41 restore taproot address tb1phcc2al0plg9rrysftk2ngpjvfgtz5d0aagvah7lc6nnml3vx9vns2tl8w5
0.01180469
*/

/*
2024/01/15 16:40:13 new priviate key b6e79ae63d00258bde4c4f1cd360c09181ac8c3663f407bc1627fc87cf14c3d2
2024/01/15 16:40:13 new public key 0385d15f94f0c7606890263d3ba790c174dbae2389b4cca80124c5dc5148cd9e45
2024/01/15 16:40:13 new legacy address n4TJ9Qbb9Va9L2EhzachG82WRZxLESHFBr
2024/01/15 16:40:13 new native segwit address tb1qlwvue3swm044hqf7s3ww8um2tuh0ncx65a6yme
2024/01/15 16:40:13 new taproot address tb1pk3ll8qd28wthj9mpcl45ydh5ge9zql5vt7a8vm65dv7t2hamymms0l8t57
2024/01/15 16:40:13 restore taproot address tb1pk3ll8qd28wthj9mpcl45ydh5ge9zql5vt7a8vm65dv7t2hamymms0l8t57

*/

/*
2024/01/18 11:36:20 new priviate key 0d28e203aa9096cfeef914eb0777621fe706e1c362a25d9711000d45cbe527e6
2024/01/18 11:36:20 new public key 03bcacaac637a63449aeac47136edbc93db7af9b772a511be84edd320b53218793
2024/01/18 11:36:20 new legacy address mi2dHPF6JgoMhAhNC2UeZC7mjjq9PMkUYi
2024/01/18 11:36:20 new native segwit address tb1qrwxdcs0npvn02rm8uzrp5ry8zutkfg2vqydc29
2024/01/18 11:36:20 new taproot address tb1p3snt0cfxn6fuej2kfr8e32mx6nyjgs5rkuz908av8kerplywrvjqqz3rj5
2024/01/18 11:36:20 restore taproot address tb1p3snt0cfxn6fuej2kfr8e32mx6nyjgs5rkuz908av8kerplywrvjqqz3rj5
*/

//https://coinfaucet.eu/en/btc-testnet/

func TestInscribeMultiMetaIdDataFromUtxo(t *testing.T) {
	var (
		netParams      *chaincfg.Params = &chaincfg.TestNet3Params
		metaIdData     InscriptionMetaIdData
		fromPriKeyHex  string = "46601a6f08ec68453fcef9beefb6b8d8a41816ac709ee4b50ce76273b5616b82"
		feeRate        int64  = 2
		toAddress      string = "tb1qfkr3t4v55kypvgs8lvpqnhduz9gy79xwqzgpnt" // 0.01794841
		changeAddress  string = toAddress
		utxoList       []*InscribeUtxo
		outAddressType string = "segwit"
		isOnlyCal      bool   = false
		revealOutValue int64  = 546

		//filePath string = "../../res/file/my-pfp.jpg"
		filePath string = "../../res/file/my-bitcoin.png"
		fileData []byte
		err      error
	)

	fileData, err = tool.ReadFile(filePath)
	if err != nil {
		t.Errorf("tool.ReadFile() error = %v", err)
		return
	}

	buzz := map[string]interface{}{
		"createTime": tool.MakeTimestamp(),
		//"content":     "",
		"content":     "a test #4 ",
		"contentType": "text/plain",
		"quoteTx":     "",
		//"attachments":[],
		//"mention":["{metaID1}","{metaID2}"]
	}

	buzz2 := map[string]interface{}{
		"createTime": tool.MakeTimestamp(),
		//"content":     "",
		"content":     "Hello bitcoin world!",
		"contentType": "text/plain",
		"quoteTx":     "",
		//"attachments":[],
		//"mention":["{metaID1}","{metaID2}"]
	}
	paylike := map[string]interface{}{
		"isLike": 1,
		"likeTo": "a2",
	}

	_ = buzz
	_ = buzz2
	_ = fileData
	_ = paylike

	metaIdData = InscriptionMetaIdData{
		//root
		// Destination: toAddress,
		// MetaIDFlag:  "testid",
		// Operation:   "init",

		////info-name
		MetaIDFlag: "testid",
		//Operation:   "create",
		Operation:   "modify",
		Path:        "@43d99745ee1e78c93aa3c9ccc6ab44da867b1c828a3dbbb4adf85f0da620bfc2i0",
		Content:     []byte("momo-e"),
		Encryption:  "",
		Version:     "",
		ContentType: "",
		Destination: toAddress,

		//protocols-SimpleBuzz
		// MetaIDFlag:  "testid",
		// Operation:   "create",
		// Path:        "/protocols/simplebuzz/01",
		// Content:     []byte(tool.AnyToStr(buzz)),
		// Encryption:  "",
		// Version:     "1",
		// ContentType: "application/json",
		// Destination: toAddress,

		//protocols-paylike
		// MetaIDFlag:  "testid",
		// Operation:   "create",
		// Path:        "/protocols/paylike/01",
		// Content:     []byte(tool.AnyToStr(paylike)),
		// Encryption:  "",
		// Version:     "1",
		// ContentType: "application/json",
		// Destination: toAddress,

		//protocols-SimpleBuzz2
		//MetaIDFlag:  "testid",
		// Operation:   "create",
		// Path:        "/protocols/simplebuzz/2",
		// Content:     []byte(tool.AnyToStr(buzz2)),
		// Encryption:  "0",
		// Version:     "",
		// ContentType: "",
		// Destination: toAddress,

		//file
		// MetaIDFlag:  "testid",
		// Operation:   "create",
		// Path:        "/file/my-bitcoin.png",
		// Content:     fileData,
		// Encryption:  "",
		// Version:     "2",
		// ContentType: "image/jpeg",
		// Destination: toAddress,
	}

	utxoList = []*InscribeUtxo{
		{
			OutTx:     "06a41225b9c4b0281a32b26f78ea8f10556adaf0f7cefd6fc5abc5b8875464c8",
			OutIndex:  1,
			OutAmount: 35793,
		},
	}

	commitTxHash, revealTxHashStrList, inscriptionList, fees, err := InscribeMultiMetaIdDataFromUtxo(netParams, metaIdData,
		fromPriKeyHex, feeRate, changeAddress, utxoList, outAddressType, isOnlyCal, revealOutValue)
	if err != nil {
		t.Errorf("InscribeMultiMetaIdDataFromUtxo() error = %v", err)
		return
	}
	t.Logf("commitTxHash = %v", commitTxHash)
	t.Logf("revealTxHashStrList = %v", revealTxHashStrList)
	t.Logf("inscriptionList = %v", inscriptionList)
	t.Logf("fees = %v", fees)

}
