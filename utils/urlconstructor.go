package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// StructToURLParams converts a struct into URL query parameters
func StructToURLParams(data interface{}) string {
	var queryParams []string

	// Get the type and value of the struct
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	// Iterate over the struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// Get the JSON tag of the field (or use field name if no JSON tag)
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		} else {
			// Strip ",omitempty" from the JSON tag
			tag = strings.Split(tag, ",")[0]
		}

		// Check if the field value is non-empty
		if !isEmpty(value) {
			// Format the query parameter and add it to the slice
			queryParams = append(queryParams, fmt.Sprintf("%s=%v", tag, value))
		}
	}

	// Join all query parameters with "&" to form the query string
	queryString := strings.Join(queryParams, "&")
	return queryString
}

// isEmpty checks if a value is nil or an empty string
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return v == ""
	default:
		return false
	}
}
