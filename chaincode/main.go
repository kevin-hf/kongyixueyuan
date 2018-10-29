/**
  author: kevin
 */

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
)

type SimpleChaincode struct {

} 

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fun == "set"{
		result, err = set(stub, args)
	}else{
		result, err = get(stub, args)
	}
	if err != nil{
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string)(string, error){

	if len(args) != 3{
		return "", fmt.Errorf("给定的参数个数不符合要求")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil{
		return "", fmt.Errorf(err.Error())
	}

	err = stub.SetEvent(args[2], []byte{})
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return string(args[0]), nil

}

func get(stub shim.ChaincodeStubInterface, args []string)(string, error){
	if len(args) != 1{
		return "", fmt.Errorf("给定的参数个数不符合要求")
	}
	result, err := stub.GetState(args[0])
	if err != nil{
		return "", fmt.Errorf("获取数据发生错误")
	}
	if result == nil{
		return "", fmt.Errorf("根据 %s 没有获取到相应的数据", args[0])
	}
	return string(result), nil

}

func main(){
	err := shim.Start(new(SimpleChaincode))
	if err != nil{
		fmt.Printf("启动SimpleChaincode时发生错误: %s", err)
	}
}

