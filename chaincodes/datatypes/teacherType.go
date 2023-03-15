package datatypes

import (
	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

// Example of a custom data type using enum-like structure (iota)
// This allows the use of verification by const values instead of float64, improving readability
// Example:
// 		if assetMap["bookType"].(float64) == (float64)(BookTypeHardcover)
// 			...

type TeacherType string

const (
	Scientific  = "Scientific"
	Litterature = "Litterature"
	Other       = "Other"
)

// CheckType checks if the given value is defined as valid BookType consts
func (b TeacherType) CheckType() errors.ICCError {
	switch b {
	case Scientific:
		return nil
	case Litterature:
		return nil
	case Other:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}

}

var teacherType = assets.DataType{
	AcceptedFormats: []string{"string"},
	DropDownValues: map[string]interface{}{
		"Scientific":  Scientific,
		"Litterature": Litterature,
		"Other":       Other,
	},
	Description: `Teacher domain type`,
}
