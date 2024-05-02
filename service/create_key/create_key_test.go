package create_key

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

//	func TestCreateBlackHoleAddress(t *testing.T) {
//		_, address, err := CreateBlackHoleAddress(&chaincfg.MainNetParams)
//		if err != nil {
//			fmt.Printf("CreateBlackHoleAddress() error = %v\n", err)
//			return
//		}
//		fmt.Printf("address = %v\n", address)
//	}
func TestKey(t *testing.T) {
	//a, b, err := CreateSegwitKey(&chaincfg.TestNet3Params)
	a, b, err := CreateSegwitKey(&chaincfg.RegressionNetParams)
	fmt.Println(err)
	fmt.Println(a)
	fmt.Println(b)
	//key:46601a6f08ec68453fcef9beefb6b8d8a41816ac709ee4b50ce76273b5616b82
	//addr:tb1qfkr3t4v55kypvgs8lvpqnhduz9gy79xwqzgpnt
}
