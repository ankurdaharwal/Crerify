package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Employee struct {
}

type EMPLOYEEJSON struct {
	EmployeeId      string `json: "EmployeeId"`
	EmployeeName    string `json: "EmployeeName"`
	CompanyName     string `json: "CompanyName"`
	JoiningDate     string `json: "JoiningDate"`
	RelievingDate   string `json: "RelievingDate"`
	Designation     string `json: "Designation"`
	EmployeeStatus  string `json: "EmployeeStatus"`
	EmployeeRegTime string `json: "EmployeeRegTime"`
}

func (t *Employee) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Check if table already exists
	_, err := stub.GetTable("EmployeeRegistration")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create BOL Table
	err = stub.CreateTable("EmployeeRegistration", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "EmployeeId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "EmployeeName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CompanyName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "JoiningDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "RelievingDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Designation", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "EmployeeStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "EmployeeRegTime", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating Employee Registration Table.")
	}

	return nil, nil

}

func (t *Employee) SubmitDoc(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, fmt.Errorf("Incorrect number of arguments. Expecting 8. Got: %d.", len(args))
	}

	EmployeeId := args[0]
	EmployeeName := args[1]
	CompanyName := args[2]
	JoiningDate := args[3]
	RelievingDate := args[4]
	Designation := args[5]
	EmployeeStatus := args[6]
	EmployeeRegTime := args[7]

	// Insert a row
	ok, err := stub.InsertRow("EmployeeRegistration", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "EmployeeRegistration"}},
			&shim.Column{Value: &shim.Column_String_{String_: EmployeeId}},
			&shim.Column{Value: &shim.Column_String_{String_: EmployeeName}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyName}},
			&shim.Column{Value: &shim.Column_String_{String_: JoiningDate}},
			&shim.Column{Value: &shim.Column_String_{String_: RelievingDate}},
			&shim.Column{Value: &shim.Column_String_{String_: Designation}},
			&shim.Column{Value: &shim.Column_String_{String_: EmployeeStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: EmployeeRegTime}}},
	})

	if !ok && err == nil {
		return nil, errors.New("Document already exists.")
	}

	return nil, err
}

func (t *Employee) GetED(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	EmployeeId := args[0]

	// Get the row pertaining to this UID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: "EmployeeRegistration"}}
	columns = append(columns, col1)
	col2 := shim.Column{Value: &shim.Column_String_{String_: EmployeeId}}
	columns = append(columns, col2)

	row, err := stub.GetRow("EmployeeRegistration", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving document with Employee ID %s. Error %s", EmployeeId, err.Error())
	}

	var EmployeeJSON EMPLOYEEJSON

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {

		EmployeeJSON.EmployeeId = ""
		EmployeeJSON.EmployeeName = ""
		EmployeeJSON.CompanyName = ""
		EmployeeJSON.JoiningDate = ""
		EmployeeJSON.RelievingDate = ""
		EmployeeJSON.Designation = ""
		EmployeeJSON.EmployeeStatus = ""
		EmployeeJSON.EmployeeRegTime = ""

	} else {

		EmployeeJSON.EmployeeId = row.Columns[2].GetString_()
		EmployeeJSON.EmployeeName = row.Columns[3].GetString_()
		EmployeeJSON.CompanyName = row.Columns[4].GetString_()
		EmployeeJSON.JoiningDate = row.Columns[5].GetString_()
		EmployeeJSON.RelievingDate = row.Columns[6].GetString_()
		EmployeeJSON.Designation = row.Columns[7].GetString_()
		EmployeeJSON.EmployeeStatus = row.Columns[8].GetString_()
		EmployeeJSON.EmployeeRegTime = row.Columns[9].GetString_()

	}

	jsonER, err := json.Marshal(EmployeeJSON)

	if err != nil {

		return nil, err
	}

	fmt.Println(jsonER)

	return jsonER, nil

}
