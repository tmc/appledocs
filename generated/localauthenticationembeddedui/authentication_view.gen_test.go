// Code generated from Apple documentation for LocalAuthenticationEmbeddedUI. DO NOT EDIT.

package localauthenticationembeddedui_test

import (
	"github.com/tmc/appledocs/generated/localauthenticationembeddedui"
)

// Suppress unused import errors
var _ = localauthenticationembeddedui.NewAuthenticationView

// ExampleNewAuthenticationViewWithContext demonstrates how to create a AuthenticationView instance using NewAuthenticationViewWithContext.
// Creates a new authentication icon that reflects the current authentication state.
func ExampleNewAuthenticationViewWithContext() {
	_ = localauthenticationembeddedui.NewAuthenticationViewWithContext(
		localauthenticationembeddedui.Context{}, // context Context
	)
	// Output:
}

// ExampleNewAuthenticationViewWithContextControlSize demonstrates how to create a AuthenticationView instance using NewAuthenticationViewWithContextControlSize.
// Creates a new authentication icon that reflects the current authentication state,   using a specified size.
func ExampleNewAuthenticationViewWithContextControlSize() {
	_ = localauthenticationembeddedui.NewAuthenticationViewWithContextControlSize(
		localauthenticationembeddedui.Context{},     // context Context
		localauthenticationembeddedui.ControlSize{}, // controlSize ControlSize
	)
	// Output:
}
