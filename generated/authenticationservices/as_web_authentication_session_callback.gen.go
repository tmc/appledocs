// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebAuthenticationSessionCallback] class.
var (
	WebAuthenticationSessionCallbackClass     _WebAuthenticationSessionCallbackClass
	WebAuthenticationSessionCallbackClassOnce sync.Once
)

func getWebAuthenticationSessionCallbackClass() _WebAuthenticationSessionCallbackClass {
	WebAuthenticationSessionCallbackClassOnce.Do(func() {
		WebAuthenticationSessionCallbackClass = _WebAuthenticationSessionCallbackClass{objc.GetClass("ASWebAuthenticationSessionCallback")}
	})
	return WebAuthenticationSessionCallbackClass
}

type _WebAuthenticationSessionCallbackClass struct {
	class objc.Class
}

// An interface definition for the [WebAuthenticationSessionCallback] class.
type IWebAuthenticationSessionCallback interface {
	objectivec.IObject
	MatchesURL(url foundation.IURL) bool
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback
type WebAuthenticationSessionCallback struct {
	objectivec.Object
}

// WebAuthenticationSessionCallbackFrom constructs a [WebAuthenticationSessionCallback] from an unsafe.Pointer.
func WebAuthenticationSessionCallbackFrom(ptr unsafe.Pointer) WebAuthenticationSessionCallback {
	return WebAuthenticationSessionCallback{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionCallbackClass) Alloc() WebAuthenticationSessionCallback {
	rv := objc.Send[WebAuthenticationSessionCallback](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebAuthenticationSessionCallbackClass) New() WebAuthenticationSessionCallback {
	rv := objc.Send[WebAuthenticationSessionCallback](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebAuthenticationSessionCallback) Init() WebAuthenticationSessionCallback {
	rv := objc.Send[WebAuthenticationSessionCallback](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebAuthenticationSessionCallback) Autorelease() WebAuthenticationSessionCallback {
	rv := objc.Send[WebAuthenticationSessionCallback](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebAuthenticationSessionCallback creates a new WebAuthenticationSessionCallback instance.
func NewWebAuthenticationSessionCallback() WebAuthenticationSessionCallback {
	return getWebAuthenticationSessionCallbackClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback/customScheme(_:)
func (wc _WebAuthenticationSessionCallbackClass) CallbackWithCustomScheme(customScheme string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("callbackWithCustomScheme:"), objc.String(customScheme))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback/https(host:path:)
func (wc _WebAuthenticationSessionCallbackClass) CallbackWithHTTPSHostPath(host string, path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("callbackWithHTTPSHost:path:"), objc.String(host), objc.String(path))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback/matchesURL(_:)
func (w_ WebAuthenticationSessionCallback) MatchesURL(url foundation.IURL) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesURL:"), url)
	return rv
}



