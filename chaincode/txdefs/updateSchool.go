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
var UpdateSchool = tx.Transaction{
	Tag:         "updateSchool",
	Label:       "Update School Data",
	Description: "Change the school information",
	Method:      "PUT",
	Callers:     []string{`$org\dMSP`, "orgMSP"}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "School ID",
			Description: "ID of the School",
			DataType:    "string",
			Required:    true,
		},
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
		id, ok := req["id"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter ID must be an school ID")
		}

		// Returns School from channel
		schoolAsset, err := id.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get school asset from the ledger")
		}
		schoolMap := (map[string]interface{})(*schoolAsset)

		// Update data
		name, _ := req["name"].(string)
		email, _ := req["email"].(string)
		phone, _ := req["phone"].(string)
		address, _ := req["address"].(string)
		typeS, _ := req["type"].(string)
		schoolMap["name"] = name
		schoolMap["email"] = email
		schoolMap["phone"] = phone
		schoolMap["address"] = address
		schoolMap["type"] = typeS

		schoolMap, err = schoolAsset.Update(stub, schoolMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update school asset")
		}

		// Marshal asset back to JSON format
		schoolJSON, nerr := json.Marshal(schoolMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return schoolJSON, nil
	},
}
