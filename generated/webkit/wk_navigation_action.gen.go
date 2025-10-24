// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NavigationAction] class.
type INavigationAction interface {
	objectivec.IObject
	// properties:
	TargetFrame() IWKFrameInfo
	ButtonNumber() unsafe.Pointer
	SetButtonNumber(value unsafe.Pointer)
	IsContentRuleListRedirect() bool
	SetIsContentRuleListRedirect(value bool)
	ModifierFlags() KeyModifierFlags /* not a class type */
	SetModifierFlags(value KeyModifierFlags /* not a class type */)
	NavigationType() NavigationType /* not a class type */
	SetNavigationType(value NavigationType /* not a class type */)
	Request() objc.IObject /* cross-framework: URLRequest */
	SetRequest(value objc.IObject /* cross-framework: URLRequest */)
	ShouldPerformDownload() bool
	SetShouldPerformDownload(value bool)
	SourceFrame() IWKFrameInfo
	SetSourceFrame(value IWKFrameInfo)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (nc _NavigationActionClass) Alloc() NavigationAction {
	rv := objc.Send[NavigationAction](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The frame in which to display the new content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/targetFrame
func (n_ NavigationAction) TargetFrame() IWKFrameInfo {
	rv := objc.Send[FrameInfo](n_.ID, objc.Sel("targetFrame"))
	return rv
}


// The number of the mouse button that caused the navigation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/buttonnumber
func (n_ NavigationAction) ButtonNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("buttonNumber"))
	return rv
}


// The number of the mouse button that caused the navigation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/buttonnumber
func (n_ NavigationAction) SetButtonNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setButtonNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/iscontentrulelistredirect
func (n_ NavigationAction) IsContentRuleListRedirect() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isContentRuleListRedirect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/iscontentrulelistredirect
func (n_ NavigationAction) SetIsContentRuleListRedirect(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsContentRuleListRedirect:"), value)
}


// The modifier keys that were pressed at the time of the navigation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/modifierflags
func (n_ NavigationAction) ModifierFlags() KeyModifierFlags /* not a class type */ {
	rv := objc.Send[KeyModifierFlags](n_.ID, objc.Sel("modifierFlags"))
	return rv
}


// The modifier keys that were pressed at the time of the navigation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/modifierflags
func (n_ NavigationAction) SetModifierFlags(value KeyModifierFlags /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setModifierFlags:"), value)
}


// The type of action that triggered the navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/navigationtype
func (n_ NavigationAction) NavigationType() NavigationType /* not a class type */ {
	rv := objc.Send[NavigationType](n_.ID, objc.Sel("navigationType"))
	return rv
}


// The type of action that triggered the navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/navigationtype
func (n_ NavigationAction) SetNavigationType(value NavigationType /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNavigationType:"), value)
}


// The URL request object associated with the navigation action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/request
func (n_ NavigationAction) Request() objc.IObject /* cross-framework: URLRequest */ {
	rv := objc.Send[foundation.URLRequest](n_.ID, objc.Sel("request"))
	return rv
}


// The URL request object associated with the navigation action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/request
func (n_ NavigationAction) SetRequest(value objc.IObject /* cross-framework: URLRequest */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRequest:"), value)
}


// A Boolean value that indicates whether the web content provided an attribute that indicates a download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/shouldperformdownload
func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("shouldPerformDownload"))
	return rv
}


// A Boolean value that indicates whether the web content provided an attribute that indicates a download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/shouldperformdownload
func (n_ NavigationAction) SetShouldPerformDownload(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setShouldPerformDownload:"), value)
}


// The frame that requested the navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/sourceframe
func (n_ NavigationAction) SourceFrame() IWKFrameInfo {
	rv := objc.Send[FrameInfo](n_.ID, objc.Sel("sourceFrame"))
	return rv
}


// The frame that requested the navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/sourceframe
func (n_ NavigationAction) SetSourceFrame(value IWKFrameInfo) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceFrame:"), value)
}



