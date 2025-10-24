// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFAuthenticationSession] class.
var (
	SFAuthenticationSessionClass     _SFAuthenticationSessionClass
	SFAuthenticationSessionClassOnce sync.Once
)

func getSFAuthenticationSessionClass() _SFAuthenticationSessionClass {
	SFAuthenticationSessionClassOnce.Do(func() {
		SFAuthenticationSessionClass = _SFAuthenticationSessionClass{objc.GetClass("SFAuthenticationSession")}
	})
	return SFAuthenticationSessionClass
}

type _SFAuthenticationSessionClass struct {
	class objc.Class
}

// An interface definition for the [SFAuthenticationSession] class.
type ISFAuthenticationSession interface {
	objectivec.IObject
	// properties:
	SFAuthenticationErrorDomain() objc.IObject /* cross-framework: NSString */
	SFContentBlockerErrorDomain() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A class that manages sharing a one-time login between Safari and an app, which can also provide automatic login for associated apps.
//
// Along with the login, this class manages sharing cookies and website data. The two cases where you would use are: Logging in to a third party’s service using an authentication protocol (for example, OAuth). This option works well for social network applications. Providing a single sign-on (SSO) experience for applications. This option works well for enterprise companies that have many applications installed on the same device. If an application uses , users are prompted by a dialog to give explicit consent, allowing the application to access the website’s data in Safari. When the webpage is presented, it runs in a separate process, so the user and web service are guaranteed that the app has no way to gain access to the user’s credentials. Instead, the app gets a unique authentication token. Then, has a simple completion handler that’s called when the session completes. After instantiating , use the start method to show the consent dialog. If the user consents, the session will begin. If at any time you wants to stop the session, call cancel to dismiss the consent dialog or dismiss the webpage. When the session is dismissed, the completion handler is called. Then, the web service redirects to the expected URL, which contains the unique authentication token. A user can decide not to log in to the session either when they are prompted with the consent dialog or after this when they’re viewing the login page. In both cases, the completion handler will be called with the error . The dismiss button in always says Cancel. Applications can’t add their own UIActivities to the Share Sheet or exclude items from the Share Sheet. However, the Share Sheet can still be used, in case the user needs a password manager to log in; additionally, it excludes items that could prevent login.


// A class that manages sharing a one-time login between Safari and an app, which can also provide automatic login for associated apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAuthenticationSession
type SFAuthenticationSession struct {
	objectivec.Object
}

// SFAuthenticationSessionFrom constructs a [SFAuthenticationSession] from an unsafe.Pointer.
//
// A class that manages sharing a one-time login between Safari and an app, which can also provide automatic login for associated apps.
func SFAuthenticationSessionFrom(ptr unsafe.Pointer) SFAuthenticationSession {
	return SFAuthenticationSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFAuthenticationSessionClass) Alloc() SFAuthenticationSession {
	rv := objc.Send[SFAuthenticationSession](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFAuthenticationSessionClass) New() SFAuthenticationSession {
	rv := objc.Send[SFAuthenticationSession](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAuthenticationSession) Init() SFAuthenticationSession {
	rv := objc.Send[SFAuthenticationSession](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAuthenticationSession) Autorelease() SFAuthenticationSession {
	rv := objc.Send[SFAuthenticationSession](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAuthenticationSession creates a new SFAuthenticationSession instance.
func NewSFAuthenticationSession() SFAuthenticationSession {
	return getSFAuthenticationSessionClass().New()
}



// Initializes the SFAuthenticationSession in an application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAuthenticationSession/init(url:callbackURLScheme:completionHandler:)
func NewSFAuthenticationSessionWithURLCallbackURLSchemeCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, callbackURLScheme objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) SFAuthenticationSession {
	instance := getSFAuthenticationSessionClass().Alloc()
	rv := objc.Send[SFAuthenticationSession](instance.ID, objc.Sel("initWithURL:callbackURLScheme:completionHandler:"), URL, callbackURLScheme, completionHandler)
	rv.Autorelease()
	return rv
}



// The domain for authentication errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfauthenticationerrordomain
func (s_ SFAuthenticationSession) SFAuthenticationErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SFAuthenticationErrorDomain"))
	return rv
}


// The domain for content blocker errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfcontentblockererrordomain
func (s_ SFAuthenticationSession) SFContentBlockerErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SFContentBlockerErrorDomain"))
	return rv
}


