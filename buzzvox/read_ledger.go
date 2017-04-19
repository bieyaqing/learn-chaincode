package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

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
		jsonResp = `{"Error":"Failed to get state for ` + key + `"}`
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
		jsonResp = `{"Error":"Failed to get state for ` + reference + `"}`
		return nil, errors.New(jsonResp)
	}

	if len(valAsbytes) == 0 {
		jsonResp = `{"Error":"Booking reference ` + reference + ` not exist"}`
		return nil, errors.New(jsonResp)
	}
	json.Unmarshal(valAsbytes, &booking)
	bookingJson, _ := json.Marshal(booking)
	return bookingJson, nil
}

func (t *SimpleChaincode) readBookingActions(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	fmt.Println("running readBookingActions()")

	var actions []Action
	var booking Booking
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	reference := args[0]

	valAsbytes, err := stub.GetState(reference)
	if err != nil {
		jsonResp = `{"Error":"Failed to get state for ` + reference + `"}`
		return nil, errors.New(jsonResp)
	}

	if len(valAsbytes) == 0 {
		jsonResp = `{"Error":"Booking reference ` + reference + ` not exist"}`
		return nil, errors.New(jsonResp)
	}
	json.Unmarshal(valAsbytes, &booking)

	count := booking.Count

	for i := 0; i <= count; i++ {
		var action Action
		actionId := `` + reference + `_` + strconv.Itoa(i) + ``
		vAsBs, err := stub.GetState(actionId)
		if err != nil {
			// NOTHING
		} else if len(vAsBs) == 0 {
			// NOTHING
		} else {
			json.Unmarshal(vAsBs, &action)
			actions = append(actions, action)
		}
	}

	actionJArr, _ := json.Marshal(actions)
	return actionJArr, nil
}
