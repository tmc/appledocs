// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PathControl] class.
var (
	PathControlClass     _PathControlClass
	PathControlClassOnce sync.Once
)

func getPathControlClass() _PathControlClass {
	PathControlClassOnce.Do(func() {
		PathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}
	})
	return PathControlClass
}

type _PathControlClass struct {
	class objc.Class
}

// An interface definition for the [PathControl] class.
type IPathControl interface {
	IControl
}

// A display of a file system path or virtual path information.
//
// The class uses to implement its user interface. provides cover methods for most methods—the cover method simply invokes the corresponding cell method. See also , which represents individual components of the path, and two associated protocols: and . has three styles represented by the enumeration constants , , and . The represented path can be a file system path or any other type of path leading through a sequence of nodes or components, as defined by the programmer. automatically supports drag and drop, which can be further customized via delegate methods. To accept drag and drop, calls with and . When the URL value in the object changes because of an automatic drag and drop operation or the user selecting a new path via the open panel, the action is sent. In OS X v10.5 the value returned by is , in macOS 10.6 and later, returns the clicked cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl
type PathControl struct {
	Control
}

// PathControlFrom constructs a [PathControl] from an unsafe.Pointer.
//
// A display of a file system path or virtual path information.
func PathControlFrom(ptr unsafe.Pointer) PathControl {
	return PathControl{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PathControlClass) Alloc() PathControl {
	rv := objc.Send[PathControl](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PathControlClass) New() PathControl {
	rv := objc.Send[PathControl](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathControl) Init() PathControl {
	rv := objc.Send[PathControl](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathControl) Autorelease() PathControl {
	rv := objc.Send[PathControl](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathControl creates a new PathControl instance.
func NewPathControl() PathControl {
	return getPathControlClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/allowedtypes
func (p_ PathControl) AllowedTypes() string {
	rv := objc.Send[string](p_.ID, objc.Sel("allowedTypes"))
	return rv
}


// SetAllowedTypes sets the value of the allowedTypes property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/allowedtypes
func (p_ PathControl) SetAllowedTypes(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedTypes:"), objc.String(value))
}

// The receiver’s background color.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/backgroundcolor
func (p_ PathControl) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The receiver’s background color.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/backgroundcolor
func (p_ PathControl) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/clickedpathitem
func (p_ PathControl) ClickedPathItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("clickedPathItem"))
	return rv
}


// SetClickedPathItem sets the value of the clickedPathItem property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/clickedpathitem
func (p_ PathControl) SetClickedPathItem(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClickedPathItem:"), value)
}

// The receiver’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/delegate
func (p_ PathControl) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The receiver’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/delegate
func (p_ PathControl) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// The receiver’s double-click action method.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/doubleaction
func (p_ PathControl) DoubleAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("doubleAction"))
	return rv
}


// SetDoubleAction sets the value of the doubleAction property.
// The receiver’s double-click action method.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/doubleaction
func (p_ PathControl) SetDoubleAction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleAction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/iseditable
func (p_ PathControl) IsEditable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/iseditable
func (p_ PathControl) SetIsEditable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEditable:"), value)
}

// The menu that is used for the path control’s cells.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/menu
func (p_ PathControl) Menu() NSMenu {
	rv := objc.Send[NSMenu](p_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The menu that is used for the path control’s cells.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/menu
func (p_ PathControl) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/pathitems
func (p_ PathControl) PathItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pathItems"))
	return rv
}


// SetPathItems sets the value of the pathItems property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/pathitems
func (p_ PathControl) SetPathItems(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathItems:"), value)
}

// The receiver’s path style.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/pathstyle
func (p_ PathControl) PathStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pathStyle"))
	return rv
}


// SetPathStyle sets the value of the pathStyle property.
// The receiver’s path style.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/pathstyle
func (p_ PathControl) SetPathStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathStyle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/placeholderattributedstring
func (p_ PathControl) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// SetPlaceholderAttributedString sets the value of the placeholderAttributedString property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/placeholderattributedstring
func (p_ PathControl) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/placeholderstring
func (p_ PathControl) PlaceholderString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("placeholderString"))
	return rv
}


// SetPlaceholderString sets the value of the placeholderString property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/placeholderstring
func (p_ PathControl) SetPlaceholderString(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}

// The path value displayed by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/url
func (p_ PathControl) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The path value displayed by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/url
func (p_ PathControl) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}



