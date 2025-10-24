//go:build darwin && ios

// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AccountAuthenticationModificationController


// Performs a request to upgrade the authentication credentials for an account to a strong password, or to use Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController/perform(_:)
func (a_ AccountAuthenticationModificationController) PerformRequest(request IASAccountAuthenticationModificationRequest) {
	objc.Send[objc.ID](a_.ID, objc.Sel("performRequest:"), request)
}

// iOS-only properties

// An object that receives notifications about the request’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController/delegate
func (a_ AccountAuthenticationModificationController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}
func (a_ AccountAuthenticationModificationController) SetDelegate(value unsafe.Pointer) {
	a_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// An object that provides a presentation context for the account modification request’s user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController/presentationContextProvider
func (a_ AccountAuthenticationModificationController) PresentationContextProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presentationContextProvider"))
	return rv
}
func (a_ AccountAuthenticationModificationController) SetPresentationContextProvider(value unsafe.Pointer) {
	a_.ID.Send(objc.RegisterName("setPresentationContextProvider:"), value)
}





