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

/* debug [class.gen.go]: Generating class ASWebAuthenticationSessionRequest */


/* debug [class_header]: Header for ASWebAuthenticationSessionRequest */
// The class instance for the [WebAuthenticationSessionRequest] class.
var (
	WebAuthenticationSessionRequestClass     _WebAuthenticationSessionRequestClass
	WebAuthenticationSessionRequestClassOnce sync.Once
)

func getWebAuthenticationSessionRequestClass() _WebAuthenticationSessionRequestClass {
	WebAuthenticationSessionRequestClassOnce.Do(func() {
		WebAuthenticationSessionRequestClass = _WebAuthenticationSessionRequestClass{objc.GetClass("ASWebAuthenticationSessionRequest")}
	})
	return WebAuthenticationSessionRequestClass
}

type _WebAuthenticationSessionRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebAuthenticationSessionRequest */
// An interface definition for the [WebAuthenticationSessionRequest] class.
type IWebAuthenticationSessionRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebAuthenticationSessionRequest */
	// properties:
	AdditionalHeaderFields() foundation.IDictionary
	Callback() IASWebAuthenticationSessionCallback
	CallbackURLScheme() objc.IObject /* cross-framework: NSString */
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ShouldUseEphemeralSession() bool
	URL() objc.IObject /* cross-framework: NSURL */
	UUID() foundation.UUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebAuthenticationSessionRequest */
	// methods:
	CancelWithError(error_ objc.IObject /* cross-framework: Error */)
	CompleteWithCallbackURL(url objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebAuthenticationSessionRequest */
// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionRequestClass) Alloc() WebAuthenticationSessionRequest {
	rv := objc.Send[WebAuthenticationSessionRequest](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebAuthenticationSessionRequestClass) New() WebAuthenticationSessionRequest {
	rv := objc.Send[WebAuthenticationSessionRequest](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebAuthenticationSessionRequest) Init() WebAuthenticationSessionRequest {
	rv := objc.Send[WebAuthenticationSessionRequest](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebAuthenticationSessionRequest) Autorelease() WebAuthenticationSessionRequest {
	rv := objc.Send[WebAuthenticationSessionRequest](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebAuthenticationSessionRequest creates a new WebAuthenticationSessionRequest instance.
func NewWebAuthenticationSessionRequest() WebAuthenticationSessionRequest {
	return getWebAuthenticationSessionRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebAuthenticationSessionRequest */
// A login session request that a web browser receives from an app.


// A login session request that a web browser receives from an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest
type WebAuthenticationSessionRequest struct {
	objectivec.Object
}

// WebAuthenticationSessionRequestFrom constructs a [WebAuthenticationSessionRequest] from an unsafe.Pointer.
//
// A login session request that a web browser receives from an app.
func WebAuthenticationSessionRequestFrom(ptr unsafe.Pointer) WebAuthenticationSessionRequest {
	return WebAuthenticationSessionRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebAuthenticationSessionRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebAuthenticationSessionRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebAuthenticationSessionRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebAuthenticationSessionRequest */

// Indicates that the browser canceled the authentication attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/cancelWithError(_:)
func (w_ WebAuthenticationSessionRequest) CancelWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("cancelWithError:"), error_)
}/* debug [instance_methods/method]: CancelWithError */


// Indicates that the browser successfully completed the authentication attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/complete(withCallbackURL:)
func (w_ WebAuthenticationSessionRequest) CompleteWithCallbackURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("completeWithCallbackURL:"), url)
}/* debug [instance_methods/method]: CompleteWithCallbackURL */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebAuthenticationSessionRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/additionalHeaderFields
func (w_ WebAuthenticationSessionRequest) AdditionalHeaderFields() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("additionalHeaderFields"))
	return rv
}/* debug [instance_properties/getter]: additionalHeaderFields */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/callback
func (w_ WebAuthenticationSessionRequest) Callback() IASWebAuthenticationSessionCallback {
	rv := objc.Send[WebAuthenticationSessionCallback](w_.ID, objc.Sel("callback"))
	return rv
}/* debug [instance_properties/getter]: callback */


// The scheme the browser should use to return the result of the authentication attempt to the app requesting it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/callbackURLScheme
func (w_ WebAuthenticationSessionRequest) CallbackURLScheme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("callbackURLScheme"))
	return rv
}/* debug [instance_properties/getter]: callbackURLScheme */


// A delegate that the session request instance informs about authentication completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/delegate
func (w_ WebAuthenticationSessionRequest) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate that the session request instance informs about authentication completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/delegate
func (w_ WebAuthenticationSessionRequest) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean that indicates whether the browser should use a private browsing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/shouldUseEphemeralSession
func (w_ WebAuthenticationSessionRequest) ShouldUseEphemeralSession() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldUseEphemeralSession"))
	return rv
}/* debug [instance_properties/getter]: shouldUseEphemeralSession */


// The web address the browser should use to perform the authentication request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/url
func (w_ WebAuthenticationSessionRequest) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// A unique identifier for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionRequest/uuid
func (w_ WebAuthenticationSessionRequest) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](w_.ID, objc.Sel("UUID"))
	return rv
}/* debug [instance_properties/getter]: UUID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASWebAuthenticationSessionRequest */





