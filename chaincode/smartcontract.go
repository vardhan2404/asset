package chaincode

import (
	//"crypto/sha256"
	//"os"
	//"log"
	"fmt"
	//"io/ioutil"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}


func (s *SmartContract) hw(ctx contractapi.TransactionContextInterface) (string, error) {
	str :=  "hello world"
	fmt.Printf("%s", str)
	return readString, nil
}
