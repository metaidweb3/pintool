package inscription_service

import (
	"encoding/hex"
	"self-tool/service/inscription_service/pkg/btcapi"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// import (
//
//	"encoding/hex"
//	"errors"
//	"fmt"
//	"github.com/btcsuite/btcd/btcec/v2"
//	"github.com/btcsuite/btcd/btcec/v2/schnorr"
//	"github.com/btcsuite/btcd/btcutil"
//	"github.com/btcsuite/btcd/chaincfg"
//	"github.com/btcsuite/btcd/chaincfg/chainhash"
//	"github.com/btcsuite/btcd/txscript"
//	"github.com/btcsuite/btcd/wire"
//	"log"
//	"self-tool/service/create_key"
//	"self-tool/service/inscription_service/internal/ord"
//	"self-tool/service/inscription_service/pkg/btcapi"
//	"self-tool/service/inscription_service/pkg/btcapi/mempool"
//
// )
type InscriptionMetaIdData struct {
	MetaIDFlag  string
	Operation   string
	Path        string
	Content     []byte
	Encryption  string
	Version     string
	ContentType string
	Destination string
}

var (
	testnetFakePriKey         string = "5457216ea9134624eb667c68de54ddca9dcb626c9a978a4bb52ba616d1d1285f"
	testnetFakeTaprootAddress string = "tb1plc2nakpmxjkp3uqzva3we6hv0txjwv3wfcadxf5ydjslkn86c4hqpvm9y6"
	fakerHash, _                     = chainhash.NewHashFromStr("f3125da7f0e0894ae51b1b7b25996026aac45617fa518113f2d28b8277f5a9da")
	fakerPkScript, _                 = hex.DecodeString("fe153ed83b34ac18f0026762eceaec7acd27322e4e3ad326846ca1fb4cfac56e")
	unspentFakerList                 = []*btcapi.UnspentOutput{
		&btcapi.UnspentOutput{
			Outpoint: &wire.OutPoint{
				Hash:  *fakerHash,
				Index: 3,
			},
			Output: &wire.TxOut{
				Value:    1000000,
				PkScript: fakerPkScript,
			},
		},
	}
)

//func CreateKeyAndCalculateInscribe(netParams *chaincfg.Params, toAddress, content string, feeRate int64) (string, string, int64, error) {
//	fromPriKeyHex, fromTaprootAddress, err := create_key.CreateTaprootKey(netParams)
//	if err != nil {
//		return "", "", 0, err
//	}
//
//	testnetNetParams := &chaincfg.SigNetParams
//	btcApiClient := mempool.NewClient(testnetNetParams)
//	contentType := "text/plain;charset=utf-8"
//	//dataMap := make(map[string]interface{})
//
//	utxoPrivateKeyHex := testnetFakePriKey
//	destination := testnetFakeTaprootAddress
//
//	commitTxOutPointList := make([]*wire.OutPoint, 0)
//	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
//
//	{
//		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
//		if err != nil {
//			return "", "", 0, err
//		}
//		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
//
//		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), testnetNetParams)
//		if err != nil {
//			return "", "", 0, err
//		}
//		_ = utxoTaprootAddress
//
//		unspentList, err := btcApiClient.ListUnspent(utxoTaprootAddress)
//		if err != nil {
//			return "", "", 0, errors.New(fmt.Sprintf("list unspent err %v", err))
//		}
//		if unspentList == nil || len(unspentList) == 0 {
//			return "", "", 0, errors.New(fmt.Sprintf("list unspent is empty"))
//		}
//
//		//unspentList := unspentFakerList
//		for i := range unspentList {
//			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
//			commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
//		}
//	}
//
//	request := ord.InscriptionRequest{
//		CommitTxOutPointList:   commitTxOutPointList,
//		CommitTxPrivateKeyList: commitTxPrivateKeyList,
//		CommitFeeRate:          feeRate,
//		FeeRate:                feeRate,
//		DataList: []ord.InscriptionData{
//			{
//				ContentType: contentType,
//				Body:        []byte(content),
//				Destination: destination,
//			},
//		},
//		SingleRevealTxOnly: false,
//	}
//
//	tool, err := ord.NewInscriptionToolWithBtcApiClient(testnetNetParams, btcApiClient, &request)
//	if err != nil {
//		//log.Fatalf("Failed to create inscription tool: %v", err)
//		return "", "", 0, err
//	}
//	fee := tool.CalculateFee()
//	return fromPriKeyHex, fromTaprootAddress, fee, nil
//}
//
//func InscribeOneData(netParams *chaincfg.Params, fromPriKeyHex, toAddress, content string, feeRate int64, changeAddress string) (string, string, string, error) {
//	btcApiClient := mempool.NewClient(netParams)
//	contentType := "text/plain;charset=utf-8"
//
//	utxoPrivateKeyHex := fromPriKeyHex
//	destination := toAddress
//
//	commitTxOutPointList := make([]*wire.OutPoint, 0)
//	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
//
//	{
//		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
//		if err != nil {
//			return "", "", "", err
//		}
//		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
//
//		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), netParams)
//		if err != nil {
//			return "", "", "", err
//		}
//
//		unspentList, err := btcApiClient.ListUnspent(utxoTaprootAddress)
//
//		if err != nil {
//			return "", "", "", errors.New(fmt.Sprintf("list unspent err %v", err))
//		}
//		if unspentList == nil || len(unspentList) == 0 {
//			return "", "", "", errors.New(fmt.Sprintf("list unspent is empty"))
//		}
//
//		for i := range unspentList {
//			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
//			commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
//		}
//	}
//
//	request := ord.InscriptionRequest{
//		CommitTxOutPointList:   commitTxOutPointList,
//		CommitTxPrivateKeyList: commitTxPrivateKeyList,
//		//CommitFeeRate:          2,
//		//FeeRate:                1,
//		CommitFeeRate: feeRate,
//		FeeRate:       feeRate,
//		DataList: []ord.InscriptionData{
//			{
//				ContentType: contentType,
//				Body:        []byte(content),
//				Destination: destination,
//			},
//		},
//		SingleRevealTxOnly: false,
//		ChangeAddress:      changeAddress,
//	}
//
//	tool, err := ord.NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
//	if err != nil {
//		return "", "", "", errors.New(fmt.Sprintf("Failed to create inscription tool: %v", err))
//	}
//	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
//	if err != nil {
//		return "", "", "", errors.New(fmt.Sprintf("send tx errr, %v", err))
//	}
//	log.Println("commitTxHash, " + commitTxHash.String())
//	revealTxHash := ""
//	for i := range revealTxHashList {
//		revealTxHash = revealTxHashList[i].String()
//		log.Println("revealTxHash, " + revealTxHashList[i].String())
//	}
//	inscription := ""
//	for i := range inscriptions {
//		inscription = inscriptions[i]
//		log.Println("inscription, " + inscriptions[i])
//	}
//	log.Println("fees: ", fees)
//	return commitTxHash.String(), revealTxHash, inscription, nil
//}

type InscribeUtxo struct {
	OutTx     string
	OutIndex  int64
	OutAmount int64
}

//
//func InscribeOneDataFromUtxo(netParams *chaincfg.Params, fromPriKeyHex, toAddress, content string, feeRate int64, changeAddress string, utxoList []*InscribeUtxo) (string, string, string, error) {
//	btcApiClient := mempool.NewClient(netParams)
//	contentType := "text/plain;charset=utf-8"
//
//	utxoPrivateKeyHex := fromPriKeyHex
//	destination := toAddress
//
//	commitTxOutPointList := make([]*wire.OutPoint, 0)
//	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
//
//	{
//		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
//		if err != nil {
//			return "", "", "", err
//		}
//		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
//
//		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), netParams)
//		if err != nil {
//			return "", "", "", err
//		}
//
//		unspentList := make([]*btcapi.UnspentOutput, 0)
//		if utxoList != nil && len(utxoList) != 0 {
//			for _, v := range utxoList {
//				txHash, err := chainhash.NewHashFromStr(v.OutTx)
//				if err != nil {
//					return "", "", "", err
//				}
//				addr, err := btcutil.DecodeAddress(utxoTaprootAddress.EncodeAddress(), netParams)
//				if err != nil {
//					return "", "", "", err
//				}
//				pkScript, err := txscript.PayToAddrScript(addr)
//				if err != nil {
//					return "", "", "", err
//				}
//				unspentList = append(unspentList, &btcapi.UnspentOutput{
//					Outpoint: &wire.OutPoint{
//						Hash:  *txHash,
//						Index: uint32(v.OutIndex),
//					},
//					Output: &wire.TxOut{
//						Value:    v.OutAmount,
//						PkScript: pkScript,
//					},
//				})
//			}
//		} else {
//			return "", "", "", err
//		}
//
//		for i := range unspentList {
//			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
//			commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
//		}
//	}
//
//	request := ord.InscriptionRequest{
//		CommitTxOutPointList:   commitTxOutPointList,
//		CommitTxPrivateKeyList: commitTxPrivateKeyList,
//		//CommitFeeRate:          2,
//		//FeeRate:                1,
//		CommitFeeRate: feeRate,
//		FeeRate:       feeRate,
//		DataList: []ord.InscriptionData{
//			{
//				ContentType: contentType,
//				Body:        []byte(content),
//				Destination: destination,
//			},
//		},
//		SingleRevealTxOnly: false,
//		ChangeAddress:      changeAddress,
//	}
//
//	tool, err := ord.NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
//	if err != nil {
//		return "", "", "", errors.New(fmt.Sprintf("Failed to create inscription tool: %v", err))
//	}
//	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
//	if err != nil {
//		return "", "", "", errors.New(fmt.Sprintf("send tx errr, %v", err))
//	}
//	log.Println("commitTxHash, " + commitTxHash.String())
//	revealTxHash := ""
//	for i := range revealTxHashList {
//		revealTxHash = revealTxHashList[i].String()
//		log.Println("revealTxHash, " + revealTxHashList[i].String())
//	}
//	inscription := ""
//	for i := range inscriptions {
//		inscription = inscriptions[i]
//		log.Println("inscription, " + inscriptions[i])
//	}
//	log.Println("fees: ", fees)
//	return commitTxHash.String(), revealTxHash, inscription, nil
//}
//
//func InscribeMultiDataFromUtxo(netParams *chaincfg.Params, fromPriKeyHex, toAddress, content string, feeRate int64, changeAddress string, count int64, utxoList []*InscribeUtxo, isLocalUtxo bool, outAddressType string, isOnlyCal bool, revealOutValue int64) (string, []string, []string, int64, error) {
//	btcApiClient := mempool.NewClient(netParams)
//	//btcApiClient := unisat.NewClient(netParams)
//	contentType := "text/plain;charset=utf-8"
//
//	utxoPrivateKeyHex := fromPriKeyHex
//	destination := toAddress
//
//	localUtxo := isLocalUtxo
//	commitTxOutPointList := make([]*wire.OutPoint, 0)
//	commitTxPreOutputList := make([]*wire.TxOut, 0)
//	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
//	commitTxUtxoAddressTypeList := make([]ord.UtxoAddressType, 0)
//	dataList := make([]ord.InscriptionData, 0)
//	for i := int64(0); i < count; i++ {
//		dataList = append(dataList, ord.InscriptionData{
//			ContentType: contentType,
//			Body:        []byte(content),
//			Destination: destination,
//		})
//	}
//
//	{
//		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
//		if err != nil {
//			return "", nil, nil, 0, err
//		}
//		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
//
//		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), netParams)
//		if err != nil {
//			return "", nil, nil, 0, err
//		}
//		nativeSegwitAddress, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(utxoPrivateKey.PubKey().SerializeCompressed()), netParams)
//		if err != nil {
//			return "", nil, nil, 0, err
//		}
//
//		unspentList := make([]*btcapi.UnspentOutput, 0)
//		if utxoList != nil && len(utxoList) != 0 {
//			localUtxo = true
//			for _, v := range utxoList {
//				txHash, err := chainhash.NewHashFromStr(v.OutTx)
//				if err != nil {
//					return "", nil, nil, 0, err
//				}
//				var addr btcutil.Address
//				if ord.UtxoAddressType(outAddressType) == ord.UtxoAddressTypeSegwit {
//					addr, err = btcutil.DecodeAddress(nativeSegwitAddress.EncodeAddress(), netParams)
//					if err != nil {
//						return "", nil, nil, 0, err
//					}
//				} else {
//					addr, err = btcutil.DecodeAddress(utxoTaprootAddress.EncodeAddress(), netParams)
//					if err != nil {
//						return "", nil, nil, 0, err
//					}
//				}
//
//				pkScript, err := txscript.PayToAddrScript(addr)
//				if err != nil {
//					return "", nil, nil, 0, err
//				}
//				unspentList = append(unspentList, &btcapi.UnspentOutput{
//					Outpoint: &wire.OutPoint{
//						Hash:  *txHash,
//						Index: uint32(v.OutIndex),
//					},
//					Output: &wire.TxOut{
//						Value:    v.OutAmount,
//						PkScript: pkScript,
//					},
//				})
//				if outAddressType != "" {
//					commitTxUtxoAddressTypeList = append(commitTxUtxoAddressTypeList, ord.UtxoAddressType(outAddressType))
//				}
//
//			}
//		} else {
//			return "", nil, nil, 0, err
//		}
//
//		for i := range unspentList {
//			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
//			commitTxPreOutputList = append(commitTxPreOutputList, unspentList[i].Output)
//			commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
//		}
//	}
//
//	request := ord.InscriptionRequest{
//		LocalUtxo:                   localUtxo,
//		CommitTxOutPointList:        commitTxOutPointList,
//		CommitTxPreOutputList:       commitTxPreOutputList,
//		CommitTxPrivateKeyList:      commitTxPrivateKeyList,
//		CommitTxUtxoAddressTypeList: commitTxUtxoAddressTypeList,
//		CommitFeeRate:               feeRate,
//		FeeRate:                     feeRate,
//		DataList:                    dataList,
//		SingleRevealTxOnly:          false,
//		ChangeAddress:               changeAddress,
//		RevealOutValue:              revealOutValue,
//	}
//
//	tool, err := ord.NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
//	if err != nil {
//		return "", nil, nil, 0, errors.New(fmt.Sprintf("Failed to create inscription tool: %v", err))
//	}
//	if isOnlyCal {
//		fees := tool.CalculateFee()
//		return "", nil, nil, fees, nil
//	}
//
//	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
//	if err != nil {
//		return "", nil, nil, 0, errors.New(fmt.Sprintf("send tx err, %v", err))
//	}
//	log.Println("commitTxHash, " + commitTxHash.String())
//	revealTxHashStrList := make([]string, 0)
//	for i := range revealTxHashList {
//		revealTxHash := revealTxHashList[i].String()
//		revealTxHashStrList = append(revealTxHashStrList, revealTxHash)
//		log.Println("revealTxHash, " + revealTxHash)
//	}
//	inscriptionList := make([]string, 0)
//	for i := range inscriptions {
//		inscriptionId := inscriptions[i]
//		inscriptionList = append(inscriptionList, inscriptionId)
//		log.Println("inscription, " + inscriptionId)
//	}
//	log.Println("fees: ", fees)
//	return commitTxHash.String(), revealTxHashStrList, inscriptionList, fees, nil
//}
//
//func InscribeMultiContentDataFromUtxo(netParams *chaincfg.Params, fromPriKeyHex, toAddress string, contentList []string, feeRate int64, changeAddress string, utxoList []*InscribeUtxo, outAddressType string, isOnlyCal bool, revealOutValue int64) (string, []string, []string, int64, error) {
//	btcApiClient := mempool.NewClient(netParams)
//	//btcApiClient := unisat.NewClient(netParams)
//	contentType := "text/plain;charset=utf-8"
//
//	utxoPrivateKeyHex := fromPriKeyHex
//	destination := toAddress
//
//	commitTxOutPointList := make([]*wire.OutPoint, 0)
//	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
//	commitTxUtxoAddressTypeList := make([]ord.UtxoAddressType, 0)
//	dataList := make([]ord.InscriptionData, 0)
//	for _, content := range contentList {
//		dataList = append(dataList, ord.InscriptionData{
//			ContentType: contentType,
//			Body:        []byte(content),
//			Destination: destination,
//		})
//	}
//
//	{
//		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
//		if err != nil {
//			return "", nil, nil, 0, err
//		}
//		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
//
//		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), netParams)
//		if err != nil {
//			return "", nil, nil, 0, err
//		}
//		nativeSegwitAddress, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(utxoPrivateKey.PubKey().SerializeCompressed()), netParams)
//		if err != nil {
//			return "", nil, nil, 0, err
//		}
//
//		unspentList := make([]*btcapi.UnspentOutput, 0)
//		if utxoList != nil && len(utxoList) != 0 {
//			for _, v := range utxoList {
//				txHash, err := chainhash.NewHashFromStr(v.OutTx)
//				if err != nil {
//					return "", nil, nil, 0, err
//				}
//				var addr btcutil.Address
//				if ord.UtxoAddressType(outAddressType) == ord.UtxoAddressTypeSegwit {
//					addr, err = btcutil.DecodeAddress(nativeSegwitAddress.EncodeAddress(), netParams)
//					if err != nil {
//						return "", nil, nil, 0, err
//					}
//				} else {
//					addr, err = btcutil.DecodeAddress(utxoTaprootAddress.EncodeAddress(), netParams)
//					if err != nil {
//						return "", nil, nil, 0, err
//					}
//				}
//
//				pkScript, err := txscript.PayToAddrScript(addr)
//				if err != nil {
//					return "", nil, nil, 0, err
//				}
//				unspentList = append(unspentList, &btcapi.UnspentOutput{
//					Outpoint: &wire.OutPoint{
//						Hash:  *txHash,
//						Index: uint32(v.OutIndex),
//					},
//					Output: &wire.TxOut{
//						Value:    v.OutAmount,
//						PkScript: pkScript,
//					},
//				})
//				if outAddressType != "" {
//					commitTxUtxoAddressTypeList = append(commitTxUtxoAddressTypeList, ord.UtxoAddressType(outAddressType))
//				}
//
//			}
//		} else {
//			return "", nil, nil, 0, err
//		}
//
//		for i := range unspentList {
//			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
//			commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
//		}
//	}
//
//	request := ord.InscriptionRequest{
//		CommitTxOutPointList:        commitTxOutPointList,
//		CommitTxPrivateKeyList:      commitTxPrivateKeyList,
//		CommitTxUtxoAddressTypeList: commitTxUtxoAddressTypeList,
//		CommitFeeRate:               feeRate,
//		FeeRate:                     feeRate,
//		DataList:                    dataList,
//		SingleRevealTxOnly:          false,
//		ChangeAddress:               changeAddress,
//		RevealOutValue:              revealOutValue,
//	}
//
//	tool, err := ord.NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
//	if err != nil {
//		return "", nil, nil, 0, errors.New(fmt.Sprintf("Failed to create inscription tool: %v", err))
//	}
//	if isOnlyCal {
//		fees := tool.CalculateFee()
//		return "", nil, nil, fees, nil
//	}
//
//	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
//	if err != nil {
//		return "", nil, nil, 0, errors.New(fmt.Sprintf("send tx err, %v", err))
//	}
//	log.Println("commitTxHash, " + commitTxHash.String())
//	revealTxHashStrList := make([]string, 0)
//	for i := range revealTxHashList {
//		revealTxHash := revealTxHashList[i].String()
//		revealTxHashStrList = append(revealTxHashStrList, revealTxHash)
//		log.Println("revealTxHash, " + revealTxHash)
//	}
//	inscriptionList := make([]string, 0)
//	for i := range inscriptions {
//		inscriptionId := inscriptions[i]
//		inscriptionList = append(inscriptionList, inscriptionId)
//		log.Println("inscription, " + inscriptionId)
//	}
//	log.Println("fees: ", fees)
//	return commitTxHash.String(), revealTxHashStrList, inscriptionList, fees, nil
//}
