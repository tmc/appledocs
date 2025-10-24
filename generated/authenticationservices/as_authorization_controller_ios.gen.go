//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AuthorizationController


// iOS-only properties

// An array of custom authorization methods for the user to choose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/customAuthorizationMethods
func (a_ AuthorizationController) CustomAuthorizationMethods() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("customAuthorizationMethods"))
	return rv
}
func (a_ AuthorizationController) SetCustomAuthorizationMethods(value []string) {
	a_.ID.Send(objc.RegisterName("setCustomAuthorizationMethods:"), value)
}





