// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit_test

import (
	"github.com/tmc/appledocs/generated/webkit"
)

// Suppress unused import errors
var _ = webkit.NewUserScript

// ExampleNewUserScriptWithSourceInjectionTimeForMainFrameOnly demonstrates how to create a UserScript instance using NewUserScriptWithSourceInjectionTimeForMainFrameOnly.
// Creates a user script object that contains the specified source code and attributes.
func ExampleNewUserScriptWithSourceInjectionTimeForMainFrameOnly() {
	_ = webkit.NewUserScriptWithSourceInjectionTimeForMainFrameOnly(
		"source",                         // source string
		webkit.UserScriptInjectionTime{}, // injectionTime UserScriptInjectionTime
		false,                            // forMainFrameOnly bool
	)
	// Output:
}

// ExampleNewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld demonstrates how to create a UserScript instance using NewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld.
// Creates a user script object that is scoped to a particular content world.
func ExampleNewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld() {
	_ = webkit.NewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld(
		"source",                         // source string
		webkit.UserScriptInjectionTime{}, // injectionTime UserScriptInjectionTime
		false,                            // forMainFrameOnly bool
		webkit.WKContentWorld{},          // contentWorld WKContentWorld
	)
	// Output:
}
