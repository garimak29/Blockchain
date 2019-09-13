package main

import(

	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type product struct{
	Id string `json:"Id"`
	Name string `json:"Name"`
	Color string `json:"Color"`
	Length string `json:"Length"`
	Width string `json:"Width"`
}

type productPrice struct{
	Id string `json:"Id"`
	BuyPrice string `json:"BuyPrice"`
	SellPrice string `json:"SellPrice"`
	
}

type MediumChaincode struct{

}
func (t *MediumChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{

	fmt.Println("Successfully Invoked init function")
	return shim.Success(nil)	
}

func (t *MediumChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Println("Start Invoke")
	defer fmt.Println("Stop Invoke")

	function
}