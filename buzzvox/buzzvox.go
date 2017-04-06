package main

import (
	"fmt"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

type Booking struct {
	ObjectType string `json:"docType"`
	Reference string `json:"reference"`
	Actor string `json:"actor"`
	UserId string `json:"userId"`
	Stage int `json:stage`
	Station string `json:station`
	ResType string `json:resType`
	Resource string `json:resource`
	Remark string `json:remark`
	Count int `json:count`
}

type Action struct {
	ObjectType string `json:"docType"`
	ActionId string `json:"actionId"`
	Actor string `json:"actor"`
	ActionName string `json:"actionName"`
	Stage int `json:stage`
	Remark string `json:remark`
	TimeStamp int64 `json:"timeStamp"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("BuzzVox_Block_Chain", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "write_booking" {
		return t.writeBooking(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	if function == "read" {
		return t.read(stub, args)
	} else if function == "read_booking" {
		return t.readBooking(stub, args)
	} else if function == "read_booking_actions" {
		return t.readBookingActions(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}


