package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return the all teachers
// GET method
var GetAllTeacher = tx.Transaction{
	Tag:         "getAllTeacher",
	Label:       "Get All Teachers",
	Description: "Return all Teacher information",
	Method:      "GET",
	Callers:     []string{`$org\dMSP`, "$orgMSP"}, // Only org1 and org2 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "limit",
			Label:       "Limit",
			Description: "Limit",
			DataType:    "number",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		limit, hasLimit := req["limit"].(float64)

		// Prepare couchdb query
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType": "teacher",
			},
		}

		if hasLimit {
			query["limit"] = limit
		}

		var err error
		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error searching for teacher", 500)
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error marshaling response", 500)
		}

		return responseJSON, nil
	},
}