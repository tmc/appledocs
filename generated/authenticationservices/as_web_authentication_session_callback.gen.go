// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASWebAuthenticationSessionCallback */


/* debug [class_header]: Header for ASWebAuthenticationSessionCallback */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebAuthenticationSessionCallback */
// An interface definition for the [WebAuthenticationSessionCallback] class.
type IWebAuthenticationSessionCallback interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebAuthenticationSessionCallback */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebAuthenticationSessionCallback */
	// methods:
	MatchesURL(url objc.IObject /* cross-framework: NSURL */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebAuthenticationSessionCallback */
// Alloc allocates a new instance without initialization.
func (wc _WebAuthenticationSessionCallbackClass) Alloc() WebAuthenticationSessionCallback {
	rv := objc.Send[WebAuthenticationSessionCallback](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebAuthenticationSessionCallback */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback
type WebAuthenticationSessionCallback struct {
	objectivec.Object
}

// WebAuthenticationSessionCallbackFrom constructs a [WebAuthenticationSessionCallback] from an unsafe.Pointer.
func WebAuthenticationSessionCallbackFrom(ptr unsafe.Pointer) WebAuthenticationSessionCallback {
	return WebAuthenticationSessionCallback{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebAuthenticationSessionCallback *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebAuthenticationSessionCallback */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback/customScheme(_:)
func (wc _WebAuthenticationSessionCallbackClass) CallbackWithCustomScheme(customScheme objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("callbackWithCustomScheme:"), customScheme)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CallbackWithCustomScheme) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback/https(host:path:)
func (wc _WebAuthenticationSessionCallbackClass) CallbackWithHTTPSHostPath(host objc.IObject /* cross-framework: NSString */, path objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("callbackWithHTTPSHost:path:"), host, path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CallbackWithHTTPSHostPath) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebAuthenticationSessionCallback */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebAuthenticationSessionCallback */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSession/Callback/matchesURL(_:)
func (w_ WebAuthenticationSessionCallback) MatchesURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("matchesURL:"), url)
	return rv
}/* debug [instance_methods/method]: MatchesURL */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebAuthenticationSessionCallback */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASWebAuthenticationSessionCallback */



