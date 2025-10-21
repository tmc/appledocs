// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [WebAuthenticationSession] class.
type IWebAuthenticationSession interface {
	objectivec.IObject
	Cancel()
	Start() bool
}

// A session that an app uses to authenticate a user through a web service.
//
// Use an instance to authenticate a user through a web service, including one run by a third party. Initialize the session with a URL that points to the authentication webpage. When the user starts the authentication session, the operating system shows a modal view telling them which domain the app is authenticating with and asking whether to proceed. If the user proceeds with the authentication attempt, a browser loads and displays the page, from which the user can authenticate. In iOS, the browser is a secure, embedded web view. In macOS, the system opens the user’s default browser if it supports web authentication sessions, or Safari otherwise. On completion, the service sends a callback URL to the session with an authentication token. The session passes this URL back to the app through a completion handler. ensures that only the calling app’s session receives the authentication callback, even when more than one app registers the same callback URL scheme. For more details, see .
//
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

// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionClass) Alloc() WebAuthenticationSession {
	rv := objc.Send[WebAuthenticationSession](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/init(url:callback:completionHandler:)
func NewWebAuthenticationSessionWithURLCallbackCompletionHandler(URL unsafe.Pointer, callback unsafe.Pointer, completionHandler unsafe.Pointer) WebAuthenticationSession {
	instance := getWebAuthenticationSessionClass().Alloc()
	rv := objc.Send[WebAuthenticationSession](instance.ID, objc.Sel("initWithURL:callback:completionHandler:"), URL, callback, completionHandler)
	rv.Autorelease()
	return rv
}


// Cancels a web authentication session.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/cancel()
func (w_ WebAuthenticationSession) Cancel() {
	objc.Send[objc.ID](w_.ID, objc.Sel("cancel"))
}

// Starts a web authentication session.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/start()
func (w_ WebAuthenticationSession) Start() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("start"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/additionalHeaderFields
func (w_ WebAuthenticationSession) AdditionalHeaderFields() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("additionalHeaderFields"))
	return rv
}


// SetAdditionalHeaderFields sets the value of the additionalHeaderFields property.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/additionalHeaderFields
func (w_ WebAuthenticationSession) SetAdditionalHeaderFields(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAdditionalHeaderFields:"), value)
}

// A Boolean indicating whether the session can begin.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/canStart
func (w_ WebAuthenticationSession) CanStart() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canStart"))
	return rv
}

// A Boolean value that indicates whether the session should ask the browser for a private authentication session.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/prefersEphemeralWebBrowserSession
func (w_ WebAuthenticationSession) PrefersEphemeralWebBrowserSession() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("prefersEphemeralWebBrowserSession"))
	return rv
}


// SetPrefersEphemeralWebBrowserSession sets the value of the prefersEphemeralWebBrowserSession property.
// A Boolean value that indicates whether the session should ask the browser for a private authentication session.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/prefersEphemeralWebBrowserSession
func (w_ WebAuthenticationSession) SetPrefersEphemeralWebBrowserSession(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPrefersEphemeralWebBrowserSession:"), value)
}

// A delegate that provides a display context in which the system can present an authentication session to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/presentationContextProvider
func (w_ WebAuthenticationSession) PresentationContextProvider() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("presentationContextProvider"))
	return rv
}


// SetPresentationContextProvider sets the value of the presentationContextProvider property.
// A delegate that provides a display context in which the system can present an authentication session to the user.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/presentationContextProvider
func (w_ WebAuthenticationSession) SetPresentationContextProvider(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPresentationContextProvider:"), value)
}


