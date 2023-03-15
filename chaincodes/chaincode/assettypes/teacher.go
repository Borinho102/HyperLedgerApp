package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

// Org1 Admin, School
// Org2 Teacher
// Org3 Student

var Teacher = assets.AssetType{
	Tag:         "teacher",
	Label:       "Teacher",
	Description: "Teacher Personal data",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Teacher ID",
			DataType: "int",                                    // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, `org2MSP`, "orgMSP"}, // This means only org1 and org2 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required:     true,
			Tag:          "subject",
			Label:        "Teaching Domain Type",
			DataType:     "teacherType",
			DefaultValue: "other",
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the teacher",
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
			Label:    "Email address of the teacher",
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
			Label:    "Phone number of the teacher",
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
			Label:    "Address of the teacher",
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
	},
}
