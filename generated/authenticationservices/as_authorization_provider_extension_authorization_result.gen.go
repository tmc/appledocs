// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationProviderExtensionAuthorizationResult */


/* debug [class_header]: Header for ASAuthorizationProviderExtensionAuthorizationResult */
// The class instance for the [AuthorizationProviderExtensionAuthorizationResult] class.
var (
	AuthorizationProviderExtensionAuthorizationResultClass     _AuthorizationProviderExtensionAuthorizationResultClass
	AuthorizationProviderExtensionAuthorizationResultClassOnce sync.Once
)

func getAuthorizationProviderExtensionAuthorizationResultClass() _AuthorizationProviderExtensionAuthorizationResultClass {
	AuthorizationProviderExtensionAuthorizationResultClassOnce.Do(func() {
		AuthorizationProviderExtensionAuthorizationResultClass = _AuthorizationProviderExtensionAuthorizationResultClass{objc.GetClass("ASAuthorizationProviderExtensionAuthorizationResult")}
	})
	return AuthorizationProviderExtensionAuthorizationResultClass
}

type _AuthorizationProviderExtensionAuthorizationResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationProviderExtensionAuthorizationResult */
// An interface definition for the [AuthorizationProviderExtensionAuthorizationResult] class.
type IAuthorizationProviderExtensionAuthorizationResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationProviderExtensionAuthorizationResult */
	// properties:
	HttpAuthorizationHeaders() foundation.IDictionary
	SetHttpAuthorizationHeaders(value foundation.IDictionary)
	HttpBody() objc.IObject /* cross-framework: NSData */
	SetHttpBody(value objc.IObject /* cross-framework: NSData */)
	HttpResponse() foundation.HTTPURLResponse
	SetHttpResponse(value foundation.HTTPURLResponse)
	PrivateKeys() objc.IObject /* cross-framework: NSArray */
	SetPrivateKeys(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationProviderExtensionAuthorizationResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationProviderExtensionAuthorizationResult */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionAuthorizationResultClass) Alloc() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationProviderExtensionAuthorizationResultClass) New() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionAuthorizationResult) Init() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionAuthorizationResult) Autorelease() AuthorizationProviderExtensionAuthorizationResult {
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionAuthorizationResult creates a new AuthorizationProviderExtensionAuthorizationResult instance.
func NewAuthorizationProviderExtensionAuthorizationResult() AuthorizationProviderExtensionAuthorizationResult {
	return getAuthorizationProviderExtensionAuthorizationResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationProviderExtensionAuthorizationResult */
// The result of an authorization request.


// The result of an authorization request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult
type AuthorizationProviderExtensionAuthorizationResult struct {
	objectivec.Object
}

// AuthorizationProviderExtensionAuthorizationResultFrom constructs a [AuthorizationProviderExtensionAuthorizationResult] from an unsafe.Pointer.
//
// The result of an authorization request.
func AuthorizationProviderExtensionAuthorizationResultFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionAuthorizationResult {
	return AuthorizationProviderExtensionAuthorizationResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationProviderExtensionAuthorizationResult */

// Initializes an authorization with tokens stored in HTTP headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/init(httpAuthorizationHeaders:)
func NewAuthorizationProviderExtensionAuthorizationResultWithHTTPAuthorizationHeaders(httpAuthorizationHeaders foundation.IDictionary) AuthorizationProviderExtensionAuthorizationResult {
	instance := getAuthorizationProviderExtensionAuthorizationResultClass().Alloc()
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](instance.ID, objc.Sel("initWithHTTPAuthorizationHeaders:"), httpAuthorizationHeaders)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationProviderExtensionAuthorizationResultWithHTTPAuthorizationHeaders */


// Initializes an authorization with a HTTP response and body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/init(httpResponse:httpBody:)
func NewAuthorizationProviderExtensionAuthorizationResultWithHTTPResponseHttpBody(httpResponse foundation.HTTPURLResponse, httpBody objc.IObject /* cross-framework: NSData */) AuthorizationProviderExtensionAuthorizationResult {
	instance := getAuthorizationProviderExtensionAuthorizationResultClass().Alloc()
	rv := objc.Send[AuthorizationProviderExtensionAuthorizationResult](instance.ID, objc.Sel("initWithHTTPResponse:httpBody:"), httpResponse, httpBody)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationProviderExtensionAuthorizationResultWithHTTPResponseHttpBody */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationProviderExtensionAuthorizationResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationProviderExtensionAuthorizationResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationProviderExtensionAuthorizationResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationProviderExtensionAuthorizationResult */

// A dictionary of authorization HTTP headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/httpAuthorizationHeaders
func (a_ AuthorizationProviderExtensionAuthorizationResult) HttpAuthorizationHeaders() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("httpAuthorizationHeaders"))
	return rv
}/* debug [instance_properties/getter]: httpAuthorizationHeaders */


// A dictionary of authorization HTTP headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/httpAuthorizationHeaders
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetHttpAuthorizationHeaders(value foundation.IDictionary) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpAuthorizationHeaders:"), value)
}/* debug [instance_properties/setter]: httpAuthorizationHeaders */


// The HTTP response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/httpBody
func (a_ AuthorizationProviderExtensionAuthorizationResult) HttpBody() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("httpBody"))
	return rv
}/* debug [instance_properties/getter]: httpBody */


// The HTTP response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/httpBody
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetHttpBody(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpBody:"), value)
}/* debug [instance_properties/setter]: httpBody */


// The HTTP response for authentications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/httpResponse
func (a_ AuthorizationProviderExtensionAuthorizationResult) HttpResponse() foundation.HTTPURLResponse {
	rv := objc.Send[foundation.HTTPURLResponse](a_.ID, objc.Sel("httpResponse"))
	return rv
}/* debug [instance_properties/getter]: httpResponse */


// The HTTP response for authentications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/httpResponse
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetHttpResponse(value foundation.HTTPURLResponse) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHttpResponse:"), value)
}/* debug [instance_properties/setter]: httpResponse */


// An array of private security keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/privateKeys
func (a_ AuthorizationProviderExtensionAuthorizationResult) PrivateKeys() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("privateKeys"))
	return rv
}/* debug [instance_properties/getter]: privateKeys */


// An array of private security keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthorizationResult/privateKeys
func (a_ AuthorizationProviderExtensionAuthorizationResult) SetPrivateKeys(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrivateKeys:"), value)
}/* debug [instance_properties/setter]: privateKeys */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationProviderExtensionAuthorizationResult */


