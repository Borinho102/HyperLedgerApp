package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

// Org1 Admin, School
// Org2 Teacher
// Org3 Student

var School = assets.AssetType{
	Tag:         "school",
	Label:       "School",
	Description: "School data",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Teacher ID",
			DataType: "string",                      // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 and org2 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required:     true,
			Tag:          "typ",
			Label:        "School Type",
			DataType:     "schoolType",
			DefaultValue: "public",
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the school",
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
			Label:    "Email address of the school",
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
			// Mandatory property
			Required: true,
			Tag:      "phone",
			Label:    "Phone number of the school",
			DataType: "string",
			// Validate funcion
			Validate: func(phone interface{}) error {
				phoneStr := phone.(string)
				if phoneStr == "" {
					return fmt.Errorf("phone must be non-empty")
				}
				return nil
			},
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "address",
			Label:    "Address of the school",
			DataType: "string",
			// Validate funcion
			Validate: func(addr interface{}) error {
				addressStr := addr.(string)
				if addressStr == "" {
					return fmt.Errorf("phone must be non-empty")
				}
				return nil
			},
		},
		{
			// Student list
			Tag:      "students",
			Label:    "Students",
			DataType: "[]->student",
		},
		{
			// Teacher list
			Tag:      "teachers",
			Label:    "Teachers",
			DataType: "[]->teacher",
		},
	},
}
