// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NavigationResponse] class.
type INavigationResponse interface {
	objectivec.IObject
}

// An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
//
// Use a object to make policy decisions about whether to allow navigation within your app’s web view. You don’t create objects directly. Instead, the web view creates them and delivers them to the appropriate delegate objects. Use the methods of your delegate to analyze the response and determine whether to allow the resulting navigation to occur.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NavigationResponseClass) Alloc() NavigationResponse {
	rv := objc.Send[NavigationResponse](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/canShowMIMEType
func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("canShowMIMEType"))
	return rv
}



