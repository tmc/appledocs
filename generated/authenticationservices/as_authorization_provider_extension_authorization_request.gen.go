// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationProviderExtensionAuthorizationRequest */


/* debug [class_header]: Header for ASAuthorizationProviderExtensionAuthorizationRequest */
// The class instance for the [AuthorizationProviderExtensionAuthorizationRequest] class.
var (
	AuthorizationProviderExtensionAuthorizationRequestClass     _AuthorizationProviderExtensionAuthorizationRequestClass
	AuthorizationProviderExtensionAuthorizationRequestClassOnce sync.Once
)

func getAuthorizationProviderExtensionAuthorizationRequestClass() _AuthorizationProviderExtensionAuthorizationRequestClass {
	AuthorizationProviderExtensionAuthorizationRequestClassOnce.Do(func() {
		AuthorizationProviderExtensionAuthorizationRequestClass = _AuthorizationProviderExtensionAuthorizationRequestClass{objc.GetClass("ASAuthorizationProviderExtensionAuthorizationRequest")}
	})
	return AuthorizationProviderExtensionAuthorizationRequestClass
}

type _AuthorizationProviderExtensionAuthorizationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationProviderExtensionAuthorizationRequest */
// An interface definition for the [AuthorizationProviderExtensionAuthorizationRequest] class.
type IAuthorizationProviderExtensionAuthorizationRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationProviderExtensionAuthorizationRequest */
	// properties:
	AuthorizationOptions() objc.IObject /* cross-framework: NSDictionary */
	CallerAuditToken() objc.IObject /* cross-framework: NSData */
	CallerBundleIdentifier() objc.IObject /* cross-framework: NSString */
	CallerTeamIdentifier() objc.IObject /* cross-framework: NSString */
	ExtensionData() objc.IObject /* cross-framework: NSDictionary */
	HttpBody() objc.IObject /* cross-framework: NSData */
	HttpHeaders() foundation.IDictionary
	CallerManaged() bool
	UserInterfaceEnabled() bool
	LocalizedCallerDisplayName() objc.IObject /* cross-framework: NSString */
	LoginManager() IASAuthorizationProviderExtensionLoginManager
	Realm() objc.IObject /* cross-framework: NSString */
	RequestedOperation() AuthorizationProviderAuthorizationOperation /* typedef */
	Url() objc.IObject /* cross-framework: NSURL */
	IsCallerManaged() bool
	SetIsCallerManaged(value bool)
	IsUserInterfaceEnabled() bool
	SetIsUserInterfaceEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationProviderExtensionAuthorizationRequest */
	// methods:
	Cancel()
	Complete()
	CompleteWithAuthorizationResult(authorizationResult IASAuthorizationProviderExtensionAuthorizationResult)
	CompleteWithError(error_ objc.IObject /* cross-framework: Error */)
	CompleteWithHTTPAuthorizationHeaders(httpAuthorizationHeaders foundation.IDictionary)
	CompleteWithHTTPResponseHttpBody(httpResponse foundation.HTTPURLResponse, httpBody objc.IObject /* cross-framework: NSData */)
	DoNotHandle()
	PresentAuthorizationViewControllerWithCompletion(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationProviderExtensionAuthorizationRequest */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionAuthorizationRequestClass) Alloc() AuthorizationProviderExtensionAuthorizationRequest {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationProviderExtensionAuthorizationRequestClass) New() AuthorizationProviderExtensionAuthorizationRequest {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionAuthorizationRequest) Init() AuthorizationProviderExtensionAuthorizationRequest {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionAuthorizationRequest) Autorelease() AuthorizationProviderExtensionAuthorizationRequest {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionAuthorizationRequest creates a new AuthorizationProviderExtensionAuthorizationRequest instance.
func NewAuthorizationProviderExtensionAuthorizationRequest() AuthorizationProviderExtensionAuthorizationRequest {
	return getAuthorizationProviderExtensionAuthorizationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationProviderExtensionAuthorizationRequest */
// An authorization request that your provider extension handles.


// An authorization request that your provider extension handles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest
type AuthorizationProviderExtensionAuthorizationRequest struct {
	objectivec.Object
}

// AuthorizationProviderExtensionAuthorizationRequestFrom constructs a [AuthorizationProviderExtensionAuthorizationRequest] from an unsafe.Pointer.
//
// An authorization request that your provider extension handles.
func AuthorizationProviderExtensionAuthorizationRequestFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionAuthorizationRequest {
	return AuthorizationProviderExtensionAuthorizationRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationProviderExtensionAuthorizationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationProviderExtensionAuthorizationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationProviderExtensionAuthorizationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationProviderExtensionAuthorizationRequest */

// Cancels the request, for example, because the user taps a cancel button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/cancel()
func (a_ AuthorizationProviderExtensionAuthorizationRequest) Cancel() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Indicates the requested authorization completed with no output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/complete()
func (a_ AuthorizationProviderExtensionAuthorizationRequest) Complete() {
	objc.Send[objc.ID](a_.ID, objc.Sel("complete"))
}/* debug [instance_methods/method]: Complete */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/complete(authorizationResult:)
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CompleteWithAuthorizationResult(authorizationResult IASAuthorizationProviderExtensionAuthorizationResult) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeWithAuthorizationResult:"), authorizationResult)
}/* debug [instance_methods/method]: CompleteWithAuthorizationResult */


// Indicates the requested authorization failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/complete(error:)
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CompleteWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeWithError:"), error_)
}/* debug [instance_methods/method]: CompleteWithError */


// Indicates the requested authorization succeeded with tokens in the HTTP headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/complete(httpAuthorizationHeaders:)
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CompleteWithHTTPAuthorizationHeaders(httpAuthorizationHeaders foundation.IDictionary) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeWithHTTPAuthorizationHeaders:"), httpAuthorizationHeaders)
}/* debug [instance_methods/method]: CompleteWithHTTPAuthorizationHeaders */


// Indicates the requested authorization succeeded with an HTTP response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/complete(httpResponse:httpBody:)
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CompleteWithHTTPResponseHttpBody(httpResponse foundation.HTTPURLResponse, httpBody objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeWithHTTPResponse:httpBody:"), httpResponse, httpBody)
}/* debug [instance_methods/method]: CompleteWithHTTPResponseHttpBody */


// Indicates the request wasn’t handled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/doNotHandle()
func (a_ AuthorizationProviderExtensionAuthorizationRequest) DoNotHandle() {
	objc.Send[objc.ID](a_.ID, objc.Sel("doNotHandle"))
}/* debug [instance_methods/method]: DoNotHandle */


// Asks the authorization service to show the extension’s view controller to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/presentAuthorizationViewController(completion:)
func (a_ AuthorizationProviderExtensionAuthorizationRequest) PresentAuthorizationViewControllerWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("presentAuthorizationViewControllerWithCompletion:"), completion)
}/* debug [instance_methods/method]: PresentAuthorizationViewControllerWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationProviderExtensionAuthorizationRequest */

// A collection of options associated with the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/authorizationOptions
func (a_ AuthorizationProviderExtensionAuthorizationRequest) AuthorizationOptions() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](a_.ID, objc.Sel("authorizationOptions"))
	return rv
}/* debug [instance_properties/getter]: authorizationOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/callerAuditToken
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CallerAuditToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("callerAuditToken"))
	return rv
}/* debug [instance_properties/getter]: callerAuditToken */


// The bundle ID of the app making the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/callerBundleIdentifier
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CallerBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("callerBundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: callerBundleIdentifier */


// The team identifier of the app making the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/callerTeamIdentifier
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CallerTeamIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("callerTeamIdentifier"))
	return rv
}/* debug [instance_properties/getter]: callerTeamIdentifier */


// Extension data from the Mobile Device Management (MDM) configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/extensionData
func (a_ AuthorizationProviderExtensionAuthorizationRequest) ExtensionData() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](a_.ID, objc.Sel("extensionData"))
	return rv
}/* debug [instance_properties/getter]: extensionData */


// The HTTP body of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/httpBody
func (a_ AuthorizationProviderExtensionAuthorizationRequest) HttpBody() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("httpBody"))
	return rv
}/* debug [instance_properties/getter]: httpBody */


// The HTTP headers of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/httpHeaders
func (a_ AuthorizationProviderExtensionAuthorizationRequest) HttpHeaders() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("httpHeaders"))
	return rv
}/* debug [instance_properties/getter]: httpHeaders */


// A Boolean value that indicates whether the app making the request is managed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/isCallerManaged
func (a_ AuthorizationProviderExtensionAuthorizationRequest) CallerManaged() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("callerManaged"))
	return rv
}/* debug [instance_properties/getter]: callerManaged */


// Determines if user interface is available for the current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/isUserInterfaceEnabled
func (a_ AuthorizationProviderExtensionAuthorizationRequest) UserInterfaceEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("userInterfaceEnabled"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceEnabled */


// The localized display name of the app making the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/localizedCallerDisplayName
func (a_ AuthorizationProviderExtensionAuthorizationRequest) LocalizedCallerDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("localizedCallerDisplayName"))
	return rv
}/* debug [instance_properties/getter]: localizedCallerDisplayName */


// The manager that interacts with Platform SSO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/loginManager
func (a_ AuthorizationProviderExtensionAuthorizationRequest) LoginManager() IASAuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](a_.ID, objc.Sel("loginManager"))
	return rv
}/* debug [instance_properties/getter]: loginManager */


// The realm to which the request applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/realm
func (a_ AuthorizationProviderExtensionAuthorizationRequest) Realm() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("realm"))
	return rv
}/* debug [instance_properties/getter]: realm */


// The operation for the extension to execute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/requestedOperation
func (a_ AuthorizationProviderExtensionAuthorizationRequest) RequestedOperation() AuthorizationProviderAuthorizationOperation /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("requestedOperation"))
	return rv
}/* debug [instance_properties/getter]: requestedOperation */


// The complete URL of the request, including all components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationRequest/url
func (a_ AuthorizationProviderExtensionAuthorizationRequest) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A Boolean value that indicates whether the app making the request is managed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationrequest/iscallermanaged
func (a_ AuthorizationProviderExtensionAuthorizationRequest) IsCallerManaged() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCallerManaged"))
	return rv
}/* debug [instance_properties/getter]: isCallerManaged */


// A Boolean value that indicates whether the app making the request is managed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationrequest/iscallermanaged
func (a_ AuthorizationProviderExtensionAuthorizationRequest) SetIsCallerManaged(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCallerManaged:"), value)
}/* debug [instance_properties/setter]: isCallerManaged */


// Determines if user interface is available for the current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationrequest/isuserinterfaceenabled
func (a_ AuthorizationProviderExtensionAuthorizationRequest) IsUserInterfaceEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isUserInterfaceEnabled"))
	return rv
}/* debug [instance_properties/getter]: isUserInterfaceEnabled */


// Determines if user interface is available for the current request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asauthorizationproviderextensionauthorizationrequest/isuserinterfaceenabled
func (a_ AuthorizationProviderExtensionAuthorizationRequest) SetIsUserInterfaceEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsUserInterfaceEnabled:"), value)
}/* debug [instance_properties/setter]: isUserInterfaceEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationProviderExtensionAuthorizationRequest */



