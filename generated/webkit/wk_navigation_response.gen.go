// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKNavigationResponse */


/* debug [class_header]: Header for WKNavigationResponse */
// The class instance for the [NavigationResponse] class.
var (
	NavigationResponseClass     _NavigationResponseClass
	NavigationResponseClassOnce sync.Once
)

func getNavigationResponseClass() _NavigationResponseClass {
	NavigationResponseClassOnce.Do(func() {
		NavigationResponseClass = _NavigationResponseClass{objc.GetClass("WKNavigationResponse")}
	})
	return NavigationResponseClass
}

type _NavigationResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NavigationResponse */
// An interface definition for the [NavigationResponse] class.
type INavigationResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NavigationResponse */
	// properties:
	CanShowMIMEType() bool
	ForMainFrame() bool
	Response() foundation.URLResponse
	IsForMainFrame() bool
	SetIsForMainFrame(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NavigationResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NavigationResponse */
// Alloc allocates a new instance without initialization.
func (nc _NavigationResponseClass) Alloc() NavigationResponse {
	rv := objc.Send[NavigationResponse](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NavigationResponseClass) New() NavigationResponse {
	rv := objc.Send[NavigationResponse](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NavigationResponse) Init() NavigationResponse {
	rv := objc.Send[NavigationResponse](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NavigationResponse) Autorelease() NavigationResponse {
	rv := objc.Send[NavigationResponse](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNavigationResponse creates a new NavigationResponse instance.
func NewNavigationResponse() NavigationResponse {
	return getNavigationResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NavigationResponse */
// An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
//
// Use a object to make policy decisions about whether to allow navigation within your app’s web view. You don’t create objects directly. Instead, the web view creates them and delivers them to the appropriate delegate objects. Use the methods of your delegate to analyze the response and determine whether to allow the resulting navigation to occur.


// An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponse
type NavigationResponse struct {
	objectivec.Object
}

// NavigationResponseFrom constructs a [NavigationResponse] from an unsafe.Pointer.
//
// An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
func NavigationResponseFrom(ptr unsafe.Pointer) NavigationResponse {
	return NavigationResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NavigationResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NavigationResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NavigationResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NavigationResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NavigationResponse */

// A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/canShowMIMEType
func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("canShowMIMEType"))
	return rv
}/* debug [instance_properties/getter]: canShowMIMEType */


// A Boolean value that indicates whether the response targets the web view’s main frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/isForMainFrame
func (n_ NavigationResponse) ForMainFrame() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("forMainFrame"))
	return rv
}/* debug [instance_properties/getter]: forMainFrame */


// The frame’s response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/response
func (n_ NavigationResponse) Response() foundation.URLResponse {
	rv := objc.Send[foundation.URLResponse](n_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// A Boolean value that indicates whether the response targets the web view’s main frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponse/isformainframe
func (n_ NavigationResponse) IsForMainFrame() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isForMainFrame"))
	return rv
}/* debug [instance_properties/getter]: isForMainFrame */


// A Boolean value that indicates whether the response targets the web view’s main frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationresponse/isformainframe
func (n_ NavigationResponse) SetIsForMainFrame(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsForMainFrame:"), value)
}/* debug [instance_properties/setter]: isForMainFrame */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKNavigationResponse */



