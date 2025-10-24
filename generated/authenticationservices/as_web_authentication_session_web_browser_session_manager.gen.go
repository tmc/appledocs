// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASWebAuthenticationSessionWebBrowserSessionManager */


/* debug [class_header]: Header for ASWebAuthenticationSessionWebBrowserSessionManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebAuthenticationSessionWebBrowserSessionManager */
// An interface definition for the [WebAuthenticationSessionWebBrowserSessionManager] class.
type IWebAuthenticationSessionWebBrowserSessionManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebAuthenticationSessionWebBrowserSessionManager */
	// properties:
	SessionHandler() unsafe.Pointer
	SetSessionHandler(value unsafe.Pointer)
	WasLaunchedByAuthenticationServices() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebAuthenticationSessionWebBrowserSessionManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebAuthenticationSessionWebBrowserSessionManager */
// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionWebBrowserSessionManagerClass) Alloc() WebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebAuthenticationSessionWebBrowserSessionManager */
// A session manager that mediates sharing data between an app and a web browser.
//
// You don’t create a session manager directly. Instead, use the session manager to tell the system what instance within your web browser app handles authentication requests. Do this by assigning an instance of a class that adopts the protocol to the shared manager’s property. You can also use the shared managers property to determine if your web browser app was launched for the specific purpose of performing authentication.


// A session manager that mediates sharing data between an app and a web browser.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebAuthenticationSessionWebBrowserSessionManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebAuthenticationSessionWebBrowserSessionManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebAuthenticationSessionWebBrowserSessionManager */

// The shared manager for which a web browser acts as the session handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionManager/shared
func (wc _WebAuthenticationSessionWebBrowserSessionManagerClass) SharedManager() WebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](objc.ID(wc.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_properties_class/property]: sharedManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebAuthenticationSessionWebBrowserSessionManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebAuthenticationSessionWebBrowserSessionManager */

// A handler that a web browser provides to handle session requests from an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionManager/sessionHandler
func (w_ WebAuthenticationSessionWebBrowserSessionManager) SessionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("sessionHandler"))
	return rv
}/* debug [instance_properties/getter]: sessionHandler */


// A handler that a web browser provides to handle session requests from an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionManager/sessionHandler
func (w_ WebAuthenticationSessionWebBrowserSessionManager) SetSessionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSessionHandler:"), value)
}/* debug [instance_properties/setter]: sessionHandler */


// The shared manager for which a web browser acts as the session handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionManager/shared
func (w_ WebAuthenticationSessionWebBrowserSessionManager) SharedManager() IASWebAuthenticationSessionWebBrowserSessionManager {
	rv := objc.Send[WebAuthenticationSessionWebBrowserSessionManager](w_.ID, objc.Sel("sharedManager"))
	return rv
}/* debug [instance_properties/getter]: sharedManager */


// A Boolean that indicates whether the session was launched by authentication services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionWebBrowserSessionManager/wasLaunchedByAuthenticationServices
func (w_ WebAuthenticationSessionWebBrowserSessionManager) WasLaunchedByAuthenticationServices() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("wasLaunchedByAuthenticationServices"))
	return rv
}/* debug [instance_properties/getter]: wasLaunchedByAuthenticationServices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASWebAuthenticationSessionWebBrowserSessionManager */



