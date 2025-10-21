// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that contains information about an action that causes navigation to occur.
//
// Use a object to make policy decisions about whether to allow navigation within your app’s web view. You don’t create objects directly. Instead, the web view creates them and delivers them to the appropriate delegate objects. Use the methods of your delegate to analyze the action and determine whether to allow the resulting navigation to occur.
//
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
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKNavigationAction/targetFrame
func (n_ NavigationAction) TargetFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("targetFrame"))
	return rv
}

// The number of the mouse button that caused the navigation request.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/buttonnumber
func (n_ NavigationAction) ButtonNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("buttonNumber"))
	return rv
}


// SetButtonNumber sets the value of the buttonNumber property.
// The number of the mouse button that caused the navigation request.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/buttonnumber
func (n_ NavigationAction) SetButtonNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setButtonNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/iscontentrulelistredirect
func (n_ NavigationAction) IsContentRuleListRedirect() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isContentRuleListRedirect"))
	return rv
}


// SetIsContentRuleListRedirect sets the value of the isContentRuleListRedirect property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/iscontentrulelistredirect
func (n_ NavigationAction) SetIsContentRuleListRedirect(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsContentRuleListRedirect:"), value)
}

// The modifier keys that were pressed at the time of the navigation request.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/modifierflags
func (n_ NavigationAction) ModifierFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("modifierFlags"))
	return rv
}


// SetModifierFlags sets the value of the modifierFlags property.
// The modifier keys that were pressed at the time of the navigation request.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/modifierflags
func (n_ NavigationAction) SetModifierFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setModifierFlags:"), value)
}

// The type of action that triggered the navigation.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/navigationtype
func (n_ NavigationAction) NavigationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("navigationType"))
	return rv
}


// SetNavigationType sets the value of the navigationType property.
// The type of action that triggered the navigation.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/navigationtype
func (n_ NavigationAction) SetNavigationType(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNavigationType:"), value)
}

// The URL request object associated with the navigation action.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/request
func (n_ NavigationAction) Request() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("request"))
	return rv
}


// SetRequest sets the value of the request property.
// The URL request object associated with the navigation action.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/request
func (n_ NavigationAction) SetRequest(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRequest:"), value)
}

// A Boolean value that indicates whether the web content provided an attribute that indicates a download.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/shouldperformdownload
func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("shouldPerformDownload"))
	return rv
}


// SetShouldPerformDownload sets the value of the shouldPerformDownload property.
// A Boolean value that indicates whether the web content provided an attribute that indicates a download.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/shouldperformdownload
func (n_ NavigationAction) SetShouldPerformDownload(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setShouldPerformDownload:"), value)
}

// The frame that requested the navigation.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/sourceframe
func (n_ NavigationAction) SourceFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("sourceFrame"))
	return rv
}


// SetSourceFrame sets the value of the sourceFrame property.
// The frame that requested the navigation.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationaction/sourceframe
func (n_ NavigationAction) SetSourceFrame(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceFrame:"), value)
}



