//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AccountAuthenticationModificationController


// iOS-only properties

// An object that receives notifications about the request’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController/delegate
func (a_ AccountAuthenticationModificationController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}
func (a_ AccountAuthenticationModificationController) SetDelegate(value objc.ID) {
	a_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





