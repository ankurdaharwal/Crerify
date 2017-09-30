package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)

// CRERIFY is a Blockchain-based Smart Contract for Employment Background Verification Checks

type CRERIFY struct {
	Company
	Employee
}

// Inits on Chaincode
func (t *CRERIFY) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	t.Company.Init(stub, function, args)
	t.Employee.Init(stub, function, args)

	return nil, nil
}



// Invokes on chaincode
func (t *CRERIFY) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerCompany" {
		
		return t.Company.SubmitDoc(stub, args)

	}  else if function == "registerCandidate"{

		return t.Employee.SubmitDoc(stub, args)

	} 

	return nil, errors.New("Invalid invoke function name.")
}



// Query callback representing the query of a chaincode
func (t *CRERIFY) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
 	
 	if function == "getCompany"{

		return t.Company.GetCD(stub, args)

 	} else if function == "getCandidate"{

		return t.sn.GetED(stub, args)
	}
	
	return nil, errors.New("Invalid query function name.")
}


func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(CRERIFY))
	if err != nil {
		fmt.Printf("Error starting CRERIFY: %s", err)
	}
}