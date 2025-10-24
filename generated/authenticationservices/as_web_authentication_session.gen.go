// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASWebAuthenticationSession */


/* debug [class_header]: Header for ASWebAuthenticationSession */
// The class instance for the [WebAuthenticationSession] class.
var (
	WebAuthenticationSessionClass     _WebAuthenticationSessionClass
	WebAuthenticationSessionClassOnce sync.Once
)

func getWebAuthenticationSessionClass() _WebAuthenticationSessionClass {
	WebAuthenticationSessionClassOnce.Do(func() {
		WebAuthenticationSessionClass = _WebAuthenticationSessionClass{objc.GetClass("ASWebAuthenticationSession")}
	})
	return WebAuthenticationSessionClass
}

type _WebAuthenticationSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebAuthenticationSession */
// An interface definition for the [WebAuthenticationSession] class.
type IWebAuthenticationSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebAuthenticationSession */
	// properties:
	AdditionalHeaderFields() foundation.IDictionary
	SetAdditionalHeaderFields(value foundation.IDictionary)
	CanStart() bool
	PrefersEphemeralWebBrowserSession() bool
	SetPrefersEphemeralWebBrowserSession(value bool)
	PresentationContextProvider() unsafe.Pointer
	SetPresentationContextProvider(value unsafe.Pointer)
	ASWebAuthenticationSessionErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebAuthenticationSession */
	// methods:
	Cancel()
	Start() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebAuthenticationSession */
// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionClass) Alloc() WebAuthenticationSession {
	rv := objc.Send[WebAuthenticationSession](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebAuthenticationSessionClass) New() WebAuthenticationSession {
	rv := objc.Send[WebAuthenticationSession](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebAuthenticationSession) Init() WebAuthenticationSession {
	rv := objc.Send[WebAuthenticationSession](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebAuthenticationSession) Autorelease() WebAuthenticationSession {
	rv := objc.Send[WebAuthenticationSession](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebAuthenticationSession creates a new WebAuthenticationSession instance.
func NewWebAuthenticationSession() WebAuthenticationSession {
	return getWebAuthenticationSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebAuthenticationSession */
// A session that an app uses to authenticate a user through a web service.
//
// Use an instance to authenticate a user through a web service, including one run by a third party. Initialize the session with a URL that points to the authentication webpage. When the user starts the authentication session, the operating system shows a modal view telling them which domain the app is authenticating with and asking whether to proceed. If the user proceeds with the authentication attempt, a browser loads and displays the page, from which the user can authenticate. In iOS, the browser is a secure, embedded web view. In macOS, the system opens the user’s default browser if it supports web authentication sessions, or Safari otherwise. On completion, the service sends a callback URL to the session with an authentication token. The session passes this URL back to the app through a completion handler. ensures that only the calling app’s session receives the authentication callback, even when more than one app registers the same callback URL scheme. For more details, see .


// A session that an app uses to authenticate a user through a web service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession
type WebAuthenticationSession struct {
	objectivec.Object
}

// WebAuthenticationSessionFrom constructs a [WebAuthenticationSession] from an unsafe.Pointer.
//
// A session that an app uses to authenticate a user through a web service.
func WebAuthenticationSessionFrom(ptr unsafe.Pointer) WebAuthenticationSession {
	return WebAuthenticationSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebAuthenticationSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/init(url:callback:completionHandler:)
func NewWebAuthenticationSessionWithURLCallbackCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, callback IASWebAuthenticationSessionCallback, completionHandler WebAuthenticationSessionCompletionHandler /* not a class type */) WebAuthenticationSession {
	instance := getWebAuthenticationSessionClass().Alloc()
	rv := objc.Send[WebAuthenticationSession](instance.ID, objc.Sel("initWithURL:callback:completionHandler:"), URL, callback, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebAuthenticationSessionWithURLCallbackCompletionHandler */


// Creates a web authentication session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/init(url:callbackURLScheme:completionHandler:)
func NewWebAuthenticationSessionWithURLCallbackURLSchemeCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, callbackURLScheme objc.IObject /* cross-framework: NSString */, completionHandler WebAuthenticationSessionCompletionHandler /* not a class type */) WebAuthenticationSession {
	instance := getWebAuthenticationSessionClass().Alloc()
	rv := objc.Send[WebAuthenticationSession](instance.ID, objc.Sel("initWithURL:callbackURLScheme:completionHandler:"), URL, callbackURLScheme, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebAuthenticationSessionWithURLCallbackURLSchemeCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebAuthenticationSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebAuthenticationSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebAuthenticationSession */

// Cancels a web authentication session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/cancel()
func (w_ WebAuthenticationSession) Cancel() {
	objc.Send[objc.ID](w_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Starts a web authentication session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/start()
func (w_ WebAuthenticationSession) Start() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("start"))
	return rv
}/* debug [instance_methods/method]: Start */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebAuthenticationSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/additionalHeaderFields
func (w_ WebAuthenticationSession) AdditionalHeaderFields() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("additionalHeaderFields"))
	return rv
}/* debug [instance_properties/getter]: additionalHeaderFields */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/additionalHeaderFields
func (w_ WebAuthenticationSession) SetAdditionalHeaderFields(value foundation.IDictionary) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAdditionalHeaderFields:"), value)
}/* debug [instance_properties/setter]: additionalHeaderFields */


// A Boolean indicating whether the session can begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/canStart
func (w_ WebAuthenticationSession) CanStart() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canStart"))
	return rv
}/* debug [instance_properties/getter]: canStart */


// A Boolean value that indicates whether the session should ask the browser for a private authentication session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/prefersEphemeralWebBrowserSession
func (w_ WebAuthenticationSession) PrefersEphemeralWebBrowserSession() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("prefersEphemeralWebBrowserSession"))
	return rv
}/* debug [instance_properties/getter]: prefersEphemeralWebBrowserSession */


// A Boolean value that indicates whether the session should ask the browser for a private authentication session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/prefersEphemeralWebBrowserSession
func (w_ WebAuthenticationSession) SetPrefersEphemeralWebBrowserSession(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPrefersEphemeralWebBrowserSession:"), value)
}/* debug [instance_properties/setter]: prefersEphemeralWebBrowserSession */


// A delegate that provides a display context in which the system can present an authentication session to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/presentationContextProvider
func (w_ WebAuthenticationSession) PresentationContextProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("presentationContextProvider"))
	return rv
}/* debug [instance_properties/getter]: presentationContextProvider */


// A delegate that provides a display context in which the system can present an authentication session to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/presentationContextProvider
func (w_ WebAuthenticationSession) SetPresentationContextProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPresentationContextProvider:"), value)
}/* debug [instance_properties/setter]: presentationContextProvider */


// The error domain for a web authentication session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aswebauthenticationsessionerrordomain
func (w_ WebAuthenticationSession) ASWebAuthenticationSessionErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("ASWebAuthenticationSessionErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: ASWebAuthenticationSessionErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASWebAuthenticationSession */


