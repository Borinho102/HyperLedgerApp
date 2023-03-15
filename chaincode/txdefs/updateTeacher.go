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
var UpdateTeacher = tx.Transaction{
	Tag:         "updateTeacher",
	Label:       "Update Teacher Data",
	Description: "Change the Teacher information",
	Method:      "PUT",
	Callers:     []string{`$org\dMSP`, "orgMSP"}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "Teacher ID",
			Description: "Teacher's ID",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Teacher's name",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "subject",
			Label:       "subject",
			Description: "Teacher's subject",
			DataType:    "TeacherType",
			Required:    true,
		},
		{
			Tag:         "email",
			Label:       "Email",
			Description: "Teacher's Email address",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "phone",
			Label:       "Phone",
			Description: "Teacher's Phone number ",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Teacher's Address",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, ok := req["id"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter ID must be an teacher ID")
		}

		// Returns teacher from channel
		TeacherAsset, err := id.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get teacher asset from the ledger")
		}
		teacherMap := (map[string]interface{})(*TeacherAsset)

		// Update data
		name, _ := req["name"].(string)
		email, _ := req["email"].(string)
		phone, _ := req["phone"].(string)
		address, _ := req["address"].(string)
		subject, _ := req["subject"].(string)
		teacherMap["name"] = name
		teacherMap["email"] = email
		teacherMap["phone"] = phone
		teacherMap["address"] = address
		teacherMap["subject"] = subject

		teacherMap, err = TeacherAsset.Update(stub, teacherMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update teacher asset")
		}

		// Marshal asset back to JSON format
		teacherJSON, nerr := json.Marshal(teacherMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return teacherJSON, nil
	},
}
