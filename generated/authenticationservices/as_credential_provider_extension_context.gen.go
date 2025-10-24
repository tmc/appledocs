// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASCredentialProviderExtensionContext */


/* debug [class_header]: Header for ASCredentialProviderExtensionContext */
// The class instance for the [CredentialProviderExtensionContext] class.
var (
	CredentialProviderExtensionContextClass     _CredentialProviderExtensionContextClass
	CredentialProviderExtensionContextClassOnce sync.Once
)

func getCredentialProviderExtensionContextClass() _CredentialProviderExtensionContextClass {
	CredentialProviderExtensionContextClassOnce.Do(func() {
		CredentialProviderExtensionContextClass = _CredentialProviderExtensionContextClass{objc.GetClass("ASCredentialProviderExtensionContext")}
	})
	return CredentialProviderExtensionContextClass
}

type _CredentialProviderExtensionContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CredentialProviderExtensionContext */
// An interface definition for the [CredentialProviderExtensionContext] class.
type ICredentialProviderExtensionContext interface {
	foundation.IExtensionContext
	
/* debug [class_interface_properties]: Properties for CredentialProviderExtensionContext */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CredentialProviderExtensionContext */
	// methods:
	CancelRequestWithError(error_ objc.IObject /* cross-framework: Error */)
	CompleteAssertionRequestWithSelectedPasskeyCredentialCompletionHandler(credential IASPasskeyAssertionCredential, completionHandler unsafe.Pointer)
	CompleteExtensionConfigurationRequest()
	CompleteOneTimeCodeRequestWithSelectedCredentialCompletionHandler(credential IASOneTimeCodeCredential, completionHandler unsafe.Pointer)
	CompleteRegistrationRequestWithSelectedPasskeyCredentialCompletionHandler(credential IASPasskeyRegistrationCredential, completionHandler unsafe.Pointer)
	CompleteRequestWithSelectedCredentialCompletionHandler(credential IASPasswordCredential, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CredentialProviderExtensionContext */
// Alloc allocates a new instance without initialization.
func (cc _CredentialProviderExtensionContextClass) Alloc() CredentialProviderExtensionContext {
	rv := objc.Send[CredentialProviderExtensionContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CredentialProviderExtensionContextClass) New() CredentialProviderExtensionContext {
	rv := objc.Send[CredentialProviderExtensionContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CredentialProviderExtensionContext) Init() CredentialProviderExtensionContext {
	rv := objc.Send[CredentialProviderExtensionContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CredentialProviderExtensionContext) Autorelease() CredentialProviderExtensionContext {
	rv := objc.Send[CredentialProviderExtensionContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCredentialProviderExtensionContext creates a new CredentialProviderExtensionContext instance.
func NewCredentialProviderExtensionContext() CredentialProviderExtensionContext {
	return getCredentialProviderExtensionContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CredentialProviderExtensionContext */
// A mechanism that credential provider extensions use to communicate with the system.


// A mechanism that credential provider extensions use to communicate with the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext
type CredentialProviderExtensionContext struct {
	foundation.ExtensionContext
}

// CredentialProviderExtensionContextFrom constructs a [CredentialProviderExtensionContext] from an unsafe.Pointer.
//
// A mechanism that credential provider extensions use to communicate with the system.
func CredentialProviderExtensionContextFrom(ptr unsafe.Pointer) CredentialProviderExtensionContext {
	return CredentialProviderExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CredentialProviderExtensionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CredentialProviderExtensionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CredentialProviderExtensionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CredentialProviderExtensionContext */

// Cancels the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/cancelRequest(withError:)
func (c_ CredentialProviderExtensionContext) CancelRequestWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancelRequestWithError:"), error_)
}/* debug [instance_methods/method]: CancelRequestWithError */


// Complete the passkey assertion request by providing the user-selected passkey credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/completeAssertionRequest(using:completionHandler:)
func (c_ CredentialProviderExtensionContext) CompleteAssertionRequestWithSelectedPasskeyCredentialCompletionHandler(credential IASPasskeyAssertionCredential, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeAssertionRequestWithSelectedPasskeyCredential:completionHandler:"), credential, completionHandler)
}/* debug [instance_methods/method]: CompleteAssertionRequestWithSelectedPasskeyCredentialCompletionHandler */


// Completes the request to configure the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/completeExtensionConfigurationRequest()
func (c_ CredentialProviderExtensionContext) CompleteExtensionConfigurationRequest() {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeExtensionConfigurationRequest"))
}/* debug [instance_methods/method]: CompleteExtensionConfigurationRequest */


// Provides the user-selected one-time passcode (OTP).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/completeOneTimeCodeRequest(using:completionHandler:)
func (c_ CredentialProviderExtensionContext) CompleteOneTimeCodeRequestWithSelectedCredentialCompletionHandler(credential IASOneTimeCodeCredential, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeOneTimeCodeRequestWithSelectedCredential:completionHandler:"), credential, completionHandler)
}/* debug [instance_methods/method]: CompleteOneTimeCodeRequestWithSelectedCredentialCompletionHandler */


// Complete the registration request by providing the newly-created passkey credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/completeRegistrationRequest(using:completionHandler:)
func (c_ CredentialProviderExtensionContext) CompleteRegistrationRequestWithSelectedPasskeyCredentialCompletionHandler(credential IASPasskeyRegistrationCredential, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeRegistrationRequestWithSelectedPasskeyCredential:completionHandler:"), credential, completionHandler)
}/* debug [instance_methods/method]: CompleteRegistrationRequestWithSelectedPasskeyCredentialCompletionHandler */


// Provides the user-selected credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialProviderExtensionContext/completeRequest(withSelectedCredential:completionHandler:)
func (c_ CredentialProviderExtensionContext) CompleteRequestWithSelectedCredentialCompletionHandler(credential IASPasswordCredential, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("completeRequestWithSelectedCredential:completionHandler:"), credential, completionHandler)
}/* debug [instance_methods/method]: CompleteRequestWithSelectedCredentialCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CredentialProviderExtensionContext */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASCredentialProviderExtensionContext */


