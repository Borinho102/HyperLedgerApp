package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Updates the tenant of a Book
// POST Method
var UpdateStudent = tx.Transaction{
	Tag:         "updateStudent",
	Label:       "Update Student Data",
	Description: "Change the Student information",
	Method:      "PUT",
	Callers:     []string{`$org\dMSP`, "orgMSP"}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "Student ID",
			Description: "ID of the Student",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the Student",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "score",
			Label:       "Score",
			Description: "Score of the Student",
			DataType:    "float32",
			Required:    true,
		},
		{
			Tag:         "email",
			Label:       "Email",
			Description: "Email address of the student",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "phone",
			Label:       "Phone",
			Description: "Phone number of the student",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Address of the student",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, ok := req["id"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter ID must be an student ID")
		}

		// Returns Student from channel
		studentAsset, err := id.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get student asset from the ledger")
		}
		studentMap := (map[string]interface{})(*studentAsset)

		// Update data
		name, _ := req["name"].(string)
		email, _ := req["email"].(string)
		phone, _ := req["phone"].(string)
		address, _ := req["address"].(string)
		score, _ := req["score"].(float32)
		studentMap["name"] = name
		studentMap["email"] = email
		studentMap["phone"] = phone
		studentMap["address"] = address
		studentMap["type"] = score

		studentMap, err = studentAsset.Update(stub, studentMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update school asset")
		}

		// Marshal asset back to JSON format
		studentJSON, nerr := json.Marshal(studentMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return studentJSON, nil
	},
}
