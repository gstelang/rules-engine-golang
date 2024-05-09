package rules

import (
	"fmt"
	"reflect"

	"github.com/gstelang/rules-engine-golang.git/app"
)

func GetFieldValue(email app.Email, fieldName string) string {

	fieldValue := reflect.ValueOf(email).FieldByName(fieldName)
	actualVal := ""
	// Check if the field exists
	if fieldValue.IsValid() {
		// Convert the field value to the appropriate type
		strVal, ok := fieldValue.Interface().(string)
		if ok {
			actualVal = strVal
			fmt.Printf("%s: %s\n", fieldName, strVal)
		} else {
			fmt.Println("Field value is not an int")
		}
	} else {
		fmt.Println("Field not found")
	}

	return actualVal
}
