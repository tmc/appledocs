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
	// properties:
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ClickedPathItem() IPathControlItem
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	Editable() bool
	SetEditable(value bool)
	Menu() IMenu
	SetMenu(value IMenu)
	PathItems() []PathControlItem
	SetPathItems(value []PathControlItem)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
	IsEditable() bool
	SetIsEditable(value bool)
	// methods:
	SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool)
}

// A display of a file system path or virtual path information.
//
// The class uses to implement its user interface. provides cover methods for most methods—the cover method simply invokes the corresponding cell method. See also , which represents individual components of the path, and two associated protocols: and . has three styles represented by the enumeration constants , , and . The represented path can be a file system path or any other type of path leading through a sequence of nodes or components, as defined by the programmer. automatically supports drag and drop, which can be further customized via delegate methods. To accept drag and drop, calls with and . When the URL value in the object changes because of an automatic drag and drop operation or the user selecting a new path via the open panel, the action is sent. In OS X v10.5 the value returned by is , in macOS 10.6 and later, returns the clicked cell.


// A display of a file system path or virtual path information.
//
// [Full Topic]
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



// Configures the drag operation mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/setDraggingSourceOperationMask(_:forLocal:)
func (p_ PathControl) SetDraggingSourceOperationMaskForLocal(mask DragOperation, isLocal bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/allowedTypes
func (p_ PathControl) AllowedTypes() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("allowedTypes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/allowedTypes
func (p_ PathControl) SetAllowedTypes(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedTypes:"), nsArray)
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/backgroundColor
func (p_ PathControl) BackgroundColor() IColor {
	rv := objc.Send[Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/backgroundColor
func (p_ PathControl) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/clickedPathItem
func (p_ PathControl) ClickedPathItem() IPathControlItem {
	rv := objc.Send[PathControlItem](p_.ID, objc.Sel("clickedPathItem"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/delegate
func (p_ PathControl) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/delegate
func (p_ PathControl) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// The receiver’s double-click action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/doubleAction
func (p_ PathControl) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](p_.ID, objc.Sel("doubleAction"))
	return rv
}


// The receiver’s double-click action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/doubleAction
func (p_ PathControl) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/isEditable
func (p_ PathControl) Editable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("editable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/isEditable
func (p_ PathControl) SetEditable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEditable:"), value)
}


// The menu that is used for the path control’s cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/menu
func (p_ PathControl) Menu() IMenu {
	rv := objc.Send[Menu](p_.ID, objc.Sel("menu"))
	return rv
}


// The menu that is used for the path control’s cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/menu
func (p_ PathControl) SetMenu(value IMenu) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMenu:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/pathItems
func (p_ PathControl) PathItems() []PathControlItem {
	rv := objc.Send[[]PathControlItem](p_.ID, objc.Sel("pathItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/pathItems
func (p_ PathControl) SetPathItems(value []PathControlItem) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathItems:"), nsArray)
}


// The receiver’s path style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/pathStyle
func (p_ PathControl) PathStyle() PathStyle {
	rv := objc.Send[PathStyle](p_.ID, objc.Sel("pathStyle"))
	return rv
}


// The receiver’s path style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/pathStyle
func (p_ PathControl) SetPathStyle(value PathStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathStyle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/placeholderAttributedString
func (p_ PathControl) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/placeholderAttributedString
func (p_ PathControl) SetPlaceholderAttributedString(value foundation.AttributedString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/placeholderString
func (p_ PathControl) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("placeholderString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/placeholderString
func (p_ PathControl) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderString:"), value)
}


// The path value displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/url
func (p_ PathControl) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}


// The path value displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/url
func (p_ PathControl) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/iseditable
func (p_ PathControl) IsEditable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEditable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontrol/iseditable
func (p_ PathControl) SetIsEditable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEditable:"), value)
}



