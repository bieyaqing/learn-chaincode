package main

import (
	// "encoding/json"
	"fmt"
	"errors"
	"strconv"
	// "strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *SimpleChaincode) writeBooking(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	fmt.Println("running writeBooking()")

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8.")
	}

	reference := args[0]
	actor := args[1]
	userId := args[2]
	stage, err := strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("4th argument must be a numeric string")
	}
	station := args[4]
	resType := args[5]
	resource := args[6]
	remark := args[7]

	str := `{
		"docType": "booking",
		"reference": "` + reference + `",
		"actor": "` + actor + `",
		"userId": "` + userId + `",
		"stage": ` + strconv.Itoa(stage) + `,
		"station": "` + station + `",
		"resType": "` + resType + `",
		"resource": "` + resource + `",
		"remark": "` + remark + `"
	}`

	err = stub.PutState(reference, []byte(str))
	if err != nil {
		return nil, err
	}
	return nil, nil
}



