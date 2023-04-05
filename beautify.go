package main

import (
	"bytes"
	"encoding/json"
)

// PrettyString Formats the JSON Response from one line to original type
func PrettyString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return ""
	}
	return prettyJSON.String()
}