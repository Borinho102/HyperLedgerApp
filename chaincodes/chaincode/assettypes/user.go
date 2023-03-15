package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

// Org1 Admin, School
// Org2 Teacher
// Org3 Student

var User = assets.AssetType{
	Tag:         "user",
	Label:       "User",
	Description: "Credentials & Personal data of User",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "User ID",
			DataType: "int",                         // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the user",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "email",
			Label:    "Email of the user",
			DataType: "string",
			// Validate funcion
			Validate: func(email interface{}) error {
				emailStr := email.(string)
				if emailStr == "" {
					return fmt.Errorf("email must be non-empty")
				}
				return nil
			},
		},
		{
			// Property with default value
			Required: true,
			Tag:      "password",
			Label:    "User password",
			DataType: "->secret",
			Validate: func(pass interface{}) error {
				passWord := pass.(string)
				if passWord == "" {
					return fmt.Errorf("password must be non-empty")
				}
				return nil
			},
		},
	},
}
