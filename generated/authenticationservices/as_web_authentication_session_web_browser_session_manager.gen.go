// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebAuthenticationSessionWebBrowserSessionManager] class.
var (
	WebAuthenticationSessionWebBrowserSessionManagerClass     _WebAuthenticationSessionWebBrowserSessionManagerClass
	WebAuthenticationSessionWebBrowserSessionManagerClassOnce sync.Once
)

func getWebAuthenticationSessionWebBrowserSessionManagerClass() _WebAuthenticationSessionWebBrowserSessionManagerClass {
	WebAuthenticationSessionWebBrowserSessionManagerClassOnce.Do(func() {
		WebAuthenticationSessionWebBrowserSessionManagerClass = _WebAuthenticationSessionWebBrowserSessionManagerClass{objc.GetClass("ASWebAuthenticationSessionWebBrowserSessionManager")}
	})
	return WebAuthenticationSessionWebBrowserSessionManagerClass
}

type _WebAuthenticationSessionWebBrowserSessionManagerClass struct {
	class objc.Class
}

// An interface definition for the [WebAuthenticationSessionWebBrowserSessionManager] class.
type IWebAuthenticationSessionWebBrowserSessionManager interface {
	objectivec.IObject
}

// A session manager that mediates sharing data between an app and a web browser.
//
// You don’t create a session manager directly. Instead, use the session manager to tell the system what instance within your web browser app handles authentication requests. Do this by assigning an instance of a class that adopts the protocol to the shared manager’s property. You can also use the shared managers property to determine if your web browser app was launched for the specific purpose of performing authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionManager
type WebAuthenticationSessionWebBrowserSessionManager struct {
	objectivec.Object
}

// WebAuthenticationSessionWebBrowserSessionManagerFrom constructs a [WebAuthenticationSessionWebBrowserSessionManager] from an unsafe.Pointer.
//
// A session manager that mediates sharing data between an app and a web browser.
func WebAuthenticationSessionWebBrowserSessionManagerFrom(ptr unsafe.Pointer) WebAuthenticationSessionWebBrowserSessionManager {
	return WebAuthenticationSessionWebBrowserSessionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionWebBrowserSessionManagerClass) Alloc() WebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebAuthenticationSessionWebBrowserSessionManagerClass) New() WebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebAuthenticationSessionWebBrowserSessionManager) Init() WebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebAuthenticationSessionWebBrowserSessionManager) Autorelease() WebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebAuthenticationSessionWebBrowserSessionManager creates a new WebAuthenticationSessionWebBrowserSessionManager instance.
func NewWebAuthenticationSessionWebBrowserSessionManager() WebAuthenticationSessionWebBrowserSessionManager {
	return getWebAuthenticationSessionWebBrowserSessionManagerClass().New()
}




