package main

import (
	
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
type Company struct {

}


type COMPANYJSON struct {

CompanyId string `json: "CompanyId"` 
CompanyName string `json: "CompanyName"`
CompanyStatus string `json: "CompanyStatus"`
CompanyRegTime string `json: "CompanyRegTime"`

}


func (t *Company) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Check if table already exists
	_, err := stub.GetTable("CompanyRegistration")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}


	// Create BOL Table
	err = stub.CreateTable("CompanyRegistration", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "CompanyId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "CompanyName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CompanyStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CompanyRegTime", Type: shim.ColumnDefinition_STRING, Key: false}
	})
	if err != nil {
		return nil, errors.New("Failed creating Company Registration Table.")
	}


	return nil, nil

	}


	func (t *Company) SubmitDoc(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		
		if len(args) != 4 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 4. Got: %d.", len(args))
		}

		CompanyId := args[0]
		CompanyName := args[1]
		CompanyStatus := args[2]
		CompanyRegTime := args[3]

		// Insert a row
	ok, err := stub.InsertRow("CompanyRegistration", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "CompanyRegistration"}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyId}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyName}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyRegTime}}

	})

	if !ok && err == nil {
		return nil, errors.New("Document already exists.")
	}

	return nil, err
}


func (t *Company) GetCD (stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
		}


	CompanyId := args[0]

	// Get the row pertaining to this UID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: "CompanyRegistration"}}
	columns = append(columns, col1)
	col2 := shim.Column{Value: &shim.Column_String_{String_: CompanyId}}
	columns = append(columns, col2)

	row, err := stub.GetRow("CompanyRegistration", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving document with Company ID %s. Error %s", CompanyId, err.Error())
	}

	var companyJSON COMPANYJSON

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		
	companyJSON.CompanyId = ""
	companyJSON.CompanyName = ""
	companyJSON.CompanyStatus = ""
	companyJSON.CompanyRegTime = ""

	} else {


	companyJSON.CompanyId = row.Columns[2].GetString_()
	companyJSON.CompanyName = row.Columns[3].GetString_()
	companyJSON.CompanyStatus = row.Columns[4].GetString_()
	companyJSON.CompanyRegTime = row.Columns[5].GetString_())

	}

	jsonCR, err := json.Marshal(companyJSON)

	if err != nil {

		return nil, err
	}

	fmt.Println(jsonCR)

 	return jsonCR, nil

	}

