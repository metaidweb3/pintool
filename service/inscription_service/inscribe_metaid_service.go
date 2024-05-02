package inscription_service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"self-tool/service/inscription_service/internal/ord"
	"self-tool/service/inscription_service/pkg/btcapi"
	"self-tool/service/inscription_service/pkg/btcapi/mempool"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func InscribeMultiMetaIdDataFromUtxo(netParams *chaincfg.Params, metaIdData InscriptionMetaIdData, fromPriKeyHex string, feeRate int64, changeAddress string, utxoList []*InscribeUtxo, outAddressType string, isOnlyCal bool, revealOutValue int64) (string, []string, []string, int64, error) {
	btcApiClient := mempool.NewClient(netParams)
	//btcApiClient := unisat.NewClient(netParams)
	utxoPrivateKeyHex := fromPriKeyHex

	localUtxo := true
	commitTxOutPointList := make([]*wire.OutPoint, 0)
	commitTxPreOutputList := make([]*wire.TxOut, 0)
	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
	commitTxUtxoAddressTypeList := make([]ord.UtxoAddressType, 0)
	metaIdDataList := make([]ord.InscriptionMetaIdData, 0)
	data := ord.InscriptionMetaIdData{
		MetaIDFlag:  metaIdData.MetaIDFlag,
		Operation:   metaIdData.Operation,
		Path:        metaIdData.Path,
		Content:     metaIdData.Content,
		Encryption:  metaIdData.Encryption,
		Version:     metaIdData.Version,
		ContentType: metaIdData.ContentType,
		Destination: metaIdData.Destination,
	}
	metaIdDataList = append(metaIdDataList, data)

	{
		utxoPrivateKeyBytes, err := hex.DecodeString(utxoPrivateKeyHex)
		if err != nil {
			return "", nil, nil, 0, err
		}
		utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)

		utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), netParams)
		if err != nil {
			return "", nil, nil, 0, err
		}
		nativeSegwitAddress, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(utxoPrivateKey.PubKey().SerializeCompressed()), netParams)
		if err != nil {
			return "", nil, nil, 0, err
		}

		unspentList := make([]*btcapi.UnspentOutput, 0)
		if utxoList != nil && len(utxoList) != 0 {
			localUtxo = true
			for _, v := range utxoList {
				txHash, err := chainhash.NewHashFromStr(v.OutTx)
				if err != nil {
					return "", nil, nil, 0, err
				}
				var addr btcutil.Address
				if ord.UtxoAddressType(outAddressType) == ord.UtxoAddressTypeSegwit {
					addr, err = btcutil.DecodeAddress(nativeSegwitAddress.EncodeAddress(), netParams)
					if err != nil {
						return "", nil, nil, 0, err
					}
				} else {
					addr, err = btcutil.DecodeAddress(utxoTaprootAddress.EncodeAddress(), netParams)
					if err != nil {
						return "", nil, nil, 0, err
					}
				}

				pkScript, err := txscript.PayToAddrScript(addr)
				if err != nil {
					return "", nil, nil, 0, err
				}
				unspentList = append(unspentList, &btcapi.UnspentOutput{
					Outpoint: &wire.OutPoint{
						Hash:  *txHash,
						Index: uint32(v.OutIndex),
					},
					Output: &wire.TxOut{
						Value:    v.OutAmount,
						PkScript: pkScript,
					},
				})
				if outAddressType != "" {
					commitTxUtxoAddressTypeList = append(commitTxUtxoAddressTypeList, ord.UtxoAddressType(outAddressType))
				}

			}
		} else {
			return "", nil, nil, 0, err
		}

		for i := range unspentList {
			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
			commitTxPreOutputList = append(commitTxPreOutputList, unspentList[i].Output)
			commitTxPrivateKeyList = append(commitTxPrivateKeyList, utxoPrivateKey)
		}
	}

	request := ord.InscriptionRequest{
		LocalUtxo:                   localUtxo,
		CommitTxOutPointList:        commitTxOutPointList,
		CommitTxPreOutputList:       commitTxPreOutputList,
		CommitTxPrivateKeyList:      commitTxPrivateKeyList,
		CommitTxUtxoAddressTypeList: commitTxUtxoAddressTypeList,
		CommitFeeRate:               feeRate,
		FeeRate:                     feeRate,
		MetaIdDataList:              metaIdDataList,
		SingleRevealTxOnly:          false,
		ChangeAddress:               changeAddress,
		RevealOutValue:              revealOutValue,
	}

	tool, err := ord.NewInscriptionToolWithBtcApiClient(netParams, btcApiClient, &request)
	if err != nil {
		return "", nil, nil, 0, errors.New(fmt.Sprintf("Failed to create inscription tool: %v", err))
	}
	if isOnlyCal {
		fees := tool.CalculateFee()
		return "", nil, nil, fees, nil
	}

	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
	if err != nil {
		return "", nil, nil, 0, errors.New(fmt.Sprintf("send tx err, %v", err))
	}
	c, _ := tool.GetCommitTxHex()
	fmt.Println(">>>>", c)
	log.Println("commitTxHash, " + commitTxHash.String())
	revealTxHashStrList := make([]string, 0)
	for i := range revealTxHashList {
		revealTxHash := revealTxHashList[i].String()
		revealTxHashStrList = append(revealTxHashStrList, revealTxHash)
		log.Println("revealTxHash, " + revealTxHash)
	}
	inscriptionList := make([]string, 0)
	for i := range inscriptions {
		inscriptionId := inscriptions[i]
		inscriptionList = append(inscriptionList, inscriptionId)
		log.Println("inscription, " + inscriptionId)
	}
	log.Println("fees: ", fees)
	return commitTxHash.String(), revealTxHashStrList, inscriptionList, fees, nil
}
