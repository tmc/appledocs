// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASSettingsHelper */


/* debug [class_header]: Header for ASSettingsHelper */
// The class instance for the [SettingsHelper] class.
var (
	SettingsHelperClass     _SettingsHelperClass
	SettingsHelperClassOnce sync.Once
)

func getSettingsHelperClass() _SettingsHelperClass {
	SettingsHelperClassOnce.Do(func() {
		SettingsHelperClass = _SettingsHelperClass{objc.GetClass("ASSettingsHelper")}
	})
	return SettingsHelperClass
}

type _SettingsHelperClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SettingsHelper */
// An interface definition for the [SettingsHelper] class.
type ISettingsHelper interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SettingsHelper */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SettingsHelper */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SettingsHelper */
// Alloc allocates a new instance without initialization.
func (sc _SettingsHelperClass) Alloc() SettingsHelper {
	rv := objc.Send[SettingsHelper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SettingsHelperClass) New() SettingsHelper {
	rv := objc.Send[SettingsHelper](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SettingsHelper) Init() SettingsHelper {
	rv := objc.Send[SettingsHelper](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SettingsHelper) Autorelease() SettingsHelper {
	rv := objc.Send[SettingsHelper](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSettingsHelper creates a new SettingsHelper instance.
func NewSettingsHelper() SettingsHelper {
	return getSettingsHelperClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SettingsHelper */
// A class that opens Settings and navigates to the settings for configuring credential providers.


// A class that opens Settings and navigates to the settings for configuring credential providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASSettingsHelper
type SettingsHelper struct {
	objectivec.Object
}

// SettingsHelperFrom constructs a [SettingsHelper] from an unsafe.Pointer.
//
// A class that opens Settings and navigates to the settings for configuring credential providers.
func SettingsHelperFrom(ptr unsafe.Pointer) SettingsHelper {
	return SettingsHelper{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SettingsHelper *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SettingsHelper */

// Open the Settings app and navigate to the AutoFill provider settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASSettingsHelper/openCredentialProviderAppSettings(completionHandler:)
func (sc _SettingsHelperClass) OpenCredentialProviderAppSettingsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("openCredentialProviderAppSettingsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenCredentialProviderAppSettingsWithCompletionHandler) */


// Open the Settings app and navigate to the verification code provider settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASSettingsHelper/openVerificationCodeAppSettings(completionHandler:)
func (sc _SettingsHelperClass) OpenVerificationCodeAppSettingsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("openVerificationCodeAppSettingsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenVerificationCodeAppSettingsWithCompletionHandler) */


// Call this method from your containing app to request to turn on a contained Credential Provider Extension. If the extension is not currently enabled, a prompt will be shown to allow it to be turned on. The completion handler is called with YES or NO depending on whether the credential provider is enabled. You need to wait 10 seconds in order to make additional request to this API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASSettingsHelper/requestToTurnOnCredentialProviderExtension(completionHandler:)
func (sc _SettingsHelperClass) RequestToTurnOnCredentialProviderExtensionWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestToTurnOnCredentialProviderExtensionWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestToTurnOnCredentialProviderExtensionWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SettingsHelper */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SettingsHelper */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SettingsHelper */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASSettingsHelper */



