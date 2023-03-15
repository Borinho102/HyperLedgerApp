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

type SchoolType string

const (
	Private = "Private"
	Public  = "Public"
)

// CheckType checks if the given value is defined as valid BookType consts
func (b SchoolType) CheckType() errors.ICCError {
	switch b {
	case Public:
		return nil
	case Private:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}

}

var schoolType = assets.DataType{
	AcceptedFormats: []string{"string"},
	DropDownValues: map[string]interface{}{
		"Private": Scientific,
		"Public":  Litterature,
	},
	Description: `School type`,
}
