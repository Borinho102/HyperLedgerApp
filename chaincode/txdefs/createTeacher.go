package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
	"github.com/google/uuid"
)

// Create a new teacher on channel
// POST Method
var CreateTeacher = tx.Transaction{
	Tag:         "createTeacher",
	Label:       "Create New teacher",
	Description: "Create a New teacher",
	Method:      "POST",
	Callers:     []string{"$org1MSP", "$orgMSP"}, // Only org can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the Teacher",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "subject",
			Label:       "subject",
			Description: "Teaching Domain Type",
			DataType:    "TeacherType",
			Required:    true,
		},
		{
			Tag:         "email",
			Label:       "Email",
			Description: "Teacher's email adress",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "phone",
			Label:       "Phone",
			Description: "Teacher's phone number",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Teacher's adress",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		email, _ := req["email"].(string)
		phone, _ := req["phone"].(string)
		address, _ := req["address"].(string)
		subject, _ := req["subject"].(string)

		TeacherMap := make(map[string]interface{})
		TeacherMap["@assetType"] = "teacher"
		TeacherMap["id"] = uuid.New()
		TeacherMap["name"] = name
		TeacherMap["email"] = email
		TeacherMap["phone"] = phone
		TeacherMap["address"] = address
		TeacherMap["subject"] = subject

		teacherAsset, err := assets.NewAsset(TeacherMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new teacher asset")
		}

		// Save the new library on channel
		_, err = teacherAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving teacher asset on blockchain")
		}

		// Marshal asset back to JSON format
		teacherJSON, nerr := json.Marshal(teacherAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return teacherJSON, nil
	},
}
