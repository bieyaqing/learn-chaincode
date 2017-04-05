package main

import (
	"encoding/json"
	"fmt"
	"errors"
	// "strconv"
	// "strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	fmt.Println("running read()")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func (t *SimpleChaincode) readBooking(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	fmt.Println("running readBooking()")

	var booking Booking

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	reference := args[0]

	valAsbytes, err := stub.GetState(reference)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + reference + "\"}"
		return nil, errors.New(jsonResp)
	}

	fmt.Println("HELLO WORLD =====================")
	json.Unmarshal(valAsbytes, &booking)
	fmt.Println(valAsbytes)
	fmt.Println(booking)
	bookingJson, _ := json.Marshal(booking)
	fmt.Println(bookingJson)
	fmt.Println("HELLO WORLD =====================")

	return valAsbytes, nil
}



