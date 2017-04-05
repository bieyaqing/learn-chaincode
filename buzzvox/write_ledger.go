package main

import (
	"fmt"
	"errors"
	"strconv"
	"time"
	"encoding/json"

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
	var jsonResp string
	var booking Booking
	var action Action
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

	// check booking
	valAsbytes, err := stub.GetState(reference)
	if err != nil {
		jsonResp = `{"Error":"Failed to get state for `+reference+`"}`
		return nil, errors.New(jsonResp)
	}

	if len(valAsbytes) == 0 {
		// bookingStr := `{
		// 	"docType": "booking",
		// 	"reference": "` + reference + `",
		// 	"actor": "` + actor + `",
		// 	"userId": "` + userId + `",
		// 	"stage": ` + strconv.Itoa(stage) + `,
		// 	"station": "` + station + `",
		// 	"resType": "` + resType + `",
		// 	"resource": "` + resource + `",
		// 	"remark": "` + remark + `",
		// 	"count": ` + strconv.Itoa(0) + `
		// }`
		booking = Booking{"booking", reference, actor, userId, stage, station, resType, resource, remark, 0}
		bookingJson, _ := json.Marshal(booking)

		err = stub.PutState(reference, bookingJson)
		if err != nil {
			return nil, err
		}
		t := time.Now()
		actionId := `` + reference + `_` + strconv.Itoa(0) + ``
		// actionStr := `{
		// 	"docType": "action",
		// 	"actionId": "` + actionId + `",
		// 	"actor": "` + actor + `",
		// 	"actionName": "create",
		// 	"stage": ` + strconv.Itoa(stage) + `,
		// 	"remark": "` + remark + `",
		// 	"timeStamp": ` + strconv.FormatInt(t.UnixNano(), 10) + `
		// }`
		action = Action{"action", actionId, actor, "create", stage, remark, t.UnixNano()}
		actionJson, _ := json.Marshal(action)
		err = stub.PutState(actionId, actionJson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		json.Unmarshal(valAsbytes, &booking)
		booking.Actor = actor
		booking.Stage = stage
		booking.Remark = remark
		booking.Count = booking.Count + 1
		bookingJson, _ := json.Marshal(booking)
		err = stub.PutState(reference, bookingJson)
		if err != nil {
			return nil, err
		}
		t := time.Now()
		actionId := `` + reference + `_` + strconv.Itoa(booking.Count) + ``
		// actionStr := `{
		// 	"docType": "action",
		// 	"actionId": "` + actionId + `",
		// 	"actor": "` + actor + `",
		// 	"actionName": "update",
		// 	"stage": ` + strconv.Itoa(stage) + `,
		// 	"remark": "` + remark + `",
		// 	"timeStamp": ` + strconv.FormatInt(t.UnixNano(), 10) + `
		// }`
		action = Action{"action", actionId, actor, "update", stage, remark, t.UnixNano()}
		actionJson, _ := json.Marshal(action)
		err = stub.PutState(actionId, actionJson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}



