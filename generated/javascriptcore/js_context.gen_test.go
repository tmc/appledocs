// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore_test

import (
	"github.com/tmc/appledocs/generated/javascriptcore"
)

// Suppress unused import errors
var _ = javascriptcore.NewJSContext

// ExampleNewJSContext demonstrates how to create a JSContext instance.
// Initializes a new JavaScript context.
func ExampleNewJSContext() {
	_ = javascriptcore.NewJSContext()
	// Output:
}
// ExampleNewJSContextWithJSGlobalContextRef demonstrates how to create a JSContext instance using NewJSContextWithJSGlobalContextRef.
// Creates a JavaScript context object from the equivalent C representation.
func ExampleNewJSContextWithJSGlobalContextRef() {
	_ = javascriptcore.NewJSContextWithJSGlobalContextRef(
		javascriptcore.JSGlobalContextRef /* typedef */{}, // jsGlobalContextRef JSGlobalContextRef /* typedef */
	)
	// Output:
}
