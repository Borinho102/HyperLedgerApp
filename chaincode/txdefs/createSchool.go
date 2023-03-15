package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/google/uuid"
)

// Create a new School on channel
// POST Method
var CreateSchool = tx.Transaction{
	Tag:         "createSchool",
	Label:       "Create New School",
	Description: "Create a New School",
	Method:      "POST",
	Callers:     []string{"$org1MSP", "$orgMSP"}, // Only org can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the School",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "type",
			Label:       "Type",
			Description: "Type of the School",
			DataType:    "schoolType",
			Required:    true,
		},
		{
			Tag:         "email",
			Label:       "Email",
			Description: "Email address of the School",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "phone",
			Label:       "Phone",
			Description: "Phone number of the School",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Address of the School",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		email, _ := req["email"].(string)
		phone, _ := req["phone"].(string)
		address, _ := req["address"].(string)
		typeS, _ := req["type"].(string)

		schoolMap := make(map[string]interface{})
		schoolMap["@assetType"] = "school"
		schoolMap["id"] = uuid.New()
		schoolMap["name"] = name
		schoolMap["email"] = email
		schoolMap["phone"] = phone
		schoolMap["address"] = address
		schoolMap["type"] = typeS

		schoolAsset, err := assets.NewAsset(schoolMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new school asset")
		}

		// Save the new library on channel
		_, err = schoolAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving school asset on blockchain")
		}

		// Marshal asset back to JSON format
		schoolJSON, nerr := json.Marshal(schoolAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return schoolJSON, nil
	},
}
