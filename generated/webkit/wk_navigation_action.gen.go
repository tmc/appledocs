// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKNavigationAction */

/* debug [class_header]: Header for WKNavigationAction */
// The class instance for the [NavigationAction] class.
var (
	NavigationActionClass     _NavigationActionClass
	NavigationActionClassOnce sync.Once
)

func getNavigationActionClass() _NavigationActionClass {
	NavigationActionClassOnce.Do(func() {
		NavigationActionClass = _NavigationActionClass{objc.GetClass("WKNavigationAction")}
	})
	return NavigationActionClass
}

type _NavigationActionClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for NavigationAction */
// An interface definition for the [NavigationAction] class.
type INavigationAction interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for NavigationAction */
	// properties:
	ButtonNumber() int
	IsContentRuleListRedirect() bool
	ModifierFlags() EventModifierFlags /* not a class type */
	NavigationType() NavigationType
	Request() foundation.URLRequest
	ShouldPerformDownload() bool
	SourceFrame() IWKFrameInfo
	TargetFrame() IWKFrameInfo
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for NavigationAction */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for NavigationAction */
// Alloc allocates a new instance without initialization.
func (nc _NavigationActionClass) Alloc() NavigationAction {
	rv := objc.Send[NavigationAction](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NavigationActionClass) New() NavigationAction {
	rv := objc.Send[NavigationAction](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NavigationAction) Init() NavigationAction {
	rv := objc.Send[NavigationAction](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NavigationAction) Autorelease() NavigationAction {
	rv := objc.Send[NavigationAction](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNavigationAction creates a new NavigationAction instance.
func NewNavigationAction() NavigationAction {
	return getNavigationActionClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for NavigationAction */
// An object that contains information about an action that causes navigation to occur.
//
// Use a object to make policy decisions about whether to allow navigation within your app’s web view. You don’t create objects directly. Instead, the web view creates them and delivers them to the appropriate delegate objects. Use the methods of your delegate to analyze the action and determine whether to allow the resulting navigation to occur.

// An object that contains information about an action that causes navigation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction
type NavigationAction struct {
	objectivec.Object
}

// NavigationActionFrom constructs a [NavigationAction] from an unsafe.Pointer.
//
// An object that contains information about an action that causes navigation to occur.
func NavigationActionFrom(ptr unsafe.Pointer) NavigationAction {
	return NavigationAction{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for NavigationAction */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for NavigationAction */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for NavigationAction */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for NavigationAction */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for NavigationAction */

// The number of the mouse button that caused the navigation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/buttonNumber
func (n_ NavigationAction) ButtonNumber() int {
	rv := objc.Send[int](n_.ID, objc.Sel("buttonNumber"))
	return rv
} /* debug [instance_properties/getter]: buttonNumber */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/isContentRuleListRedirect
func (n_ NavigationAction) IsContentRuleListRedirect() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isContentRuleListRedirect"))
	return rv
} /* debug [instance_properties/getter]: isContentRuleListRedirect */

// The modifier keys that were pressed at the time of the navigation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/modifierFlags
func (n_ NavigationAction) ModifierFlags() EventModifierFlags /* not a class type */ {
	rv := objc.Send[EventModifierFlags](n_.ID, objc.Sel("modifierFlags"))
	return rv
} /* debug [instance_properties/getter]: modifierFlags */

// The type of action that triggered the navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/navigationType
func (n_ NavigationAction) NavigationType() NavigationType {
	rv := objc.Send[NavigationType](n_.ID, objc.Sel("navigationType"))
	return rv
} /* debug [instance_properties/getter]: navigationType */

// The URL request object associated with the navigation action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/request
func (n_ NavigationAction) Request() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](n_.ID, objc.Sel("request"))
	return rv
} /* debug [instance_properties/getter]: request */

// A Boolean value that indicates whether the web content provided an attribute that indicates a download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/shouldPerformDownload
func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("shouldPerformDownload"))
	return rv
} /* debug [instance_properties/getter]: shouldPerformDownload */

// The frame that requested the navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/sourceFrame
func (n_ NavigationAction) SourceFrame() IWKFrameInfo {
	rv := objc.Send[FrameInfo](n_.ID, objc.Sel("sourceFrame"))
	return rv
} /* debug [instance_properties/getter]: sourceFrame */

// The frame in which to display the new content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/targetFrame
func (n_ NavigationAction) TargetFrame() IWKFrameInfo {
	rv := objc.Send[FrameInfo](n_.ID, objc.Sel("targetFrame"))
	return rv
} /* debug [instance_properties/getter]: targetFrame */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKNavigationAction */
