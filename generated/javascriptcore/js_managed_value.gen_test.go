// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore_test

import (
	"github.com/tmc/appledocs/generated/javascriptcore"
)

// Suppress unused import errors
var _ = javascriptcore.NewJSManagedValue

// ExampleNewJSManagedValueWithValue demonstrates how to create a JSManagedValue instance using NewJSManagedValueWithValue.
// Initializes a managed value with the specified JavaScript value.
func ExampleNewJSManagedValueWithValue() {
	_ = javascriptcore.NewJSManagedValueWithValue(
		javascriptcore.JSValue{}, // value JSValue
	)
	// Output:
}
