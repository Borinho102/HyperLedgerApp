package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/google/uuid"
)

// Create a new Student on channel
// POST Method
var CreateStudent = tx.Transaction{
	Tag:         "createStudent",
	Label:       "Create New Student",
	Description: "Create a New Student",
	Method:      "POST",
	Callers:     []string{"$org1MSP", "$orgMSP"}, // Only org can call this transaction

	Args: []tx.Argument{
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
			Description: "Student score",
			DataType:    "float32",
			Required:    true,
		},
		{
			Tag:         "email",
			Label:       "Email",
			Description: "Email address of the Student",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "phone",
			Label:       "Phone",
			Description: "Phone number of the Student",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Address of the Student",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		email, _ := req["email"].(string)
		phone, _ := req["phone"].(string)
		address, _ := req["address"].(string)
		score, _ := req["score"].(float32)

		studentMap := make(map[string]interface{})
		studentMap["@assetType"] = "student"
		studentMap["id"] = uuid.New()
		studentMap["name"] = name
		studentMap["email"] = email
		studentMap["phone"] = phone
		studentMap["address"] = address
		studentMap["score"] = score

		studentAsset, err := assets.NewAsset(studentMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new student asset")
		}

		// Save the new library on channel
		_, err = studentAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving student asset on blockchain")
		}

		// Marshal asset back to JSON format
		studentJSON, nerr := json.Marshal(studentAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return studentJSON, nil
	},
}
