// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PathCell] class.
var (
	PathCellClass     _PathCellClass
	PathCellClassOnce sync.Once
)

func getPathCellClass() _PathCellClass {
	PathCellClassOnce.Do(func() {
		PathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}
	})
	return PathCellClass
}

type _PathCellClass struct {
	class objc.Class
}

// An interface definition for the [PathCell] class.
type IPathCell interface {
	IActionCell
	ControlSize() ControlSize
	SetControlSize(value IControlSize)
	IsEditable() bool
	SetIsEditable(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
	AllowedTypes() string
	SetAllowedTypes(value string)
	BackgroundColor() NSColor
	SetBackgroundColor(value IColor)
	ClickedPathComponentCell() NSPathComponentCell
	SetClickedPathComponentCell(value IPathComponentCell)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DoubleAction() unsafe.Pointer
	SetDoubleAction(value unsafe.Pointer)
	PathComponentCells() NSPathComponentCell
	SetPathComponentCells(value IPathComponentCell)
	PathStyle() unsafe.Pointer
	SetPathStyle(value unsafe.Pointer)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
	Url() foundation.URL
	SetUrl(value foundation.IURL)
}

// The user interface of a path control object.
//
// maintains a collection of objects that represent a particular path to be displayed to the user. The path shown can be set with the method. Doing so removes all displayed objects and automatically fills the control with objects set to have the appropriate icons, display titles, and values for the particular path component they represent. Alternatively, you can fill the control manually by setting the cell array or directly modifying existing cells. Both an action and double-click action can be set for the path control. To find out what path component cell was clicked in the action, you can read the value of . When the style is set to , the action is still sent, and the value for the represented menu item is correctly set. The value is valid only when the action is being sent. It is also valid when the keyboard is used to invoke the action. Automatic animated expansion of partially hidden objects happens if you correctly call and for each in the object. This is not required if the is set to , or if you wish to not have the animation. supports several path display styles. has a light blue background with arrows indicating the path. has more defined arrows (chevrons) and looks a little like a segmented button. looks and works like an object to display the full path, or, if the cell is editable, select a new path. If the cell’s method returns (the default), you can drag and drop into the cell to change the value. You can constrain what can be dropped using UTIs (Uniform Type Identifiers) with or the appropriate delegate methods on . If the cell’s method returns (the default), the cell’s contents can automatically be dragged out. The proper UTI, filename, and URL are placed on the pasteboard. You can further control or limit this by using the appropriate delegate methods on . If the cell is editable and has the path style set to , an additional item in the pop-up menu allows selecting another location. By default, an object is configured based on the allowed types. The object can be customized with a delegate method.


// The user interface of a path control object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell

type PathCell struct {
	ActionCell
}

// PathCellFrom constructs a [PathCell] from an unsafe.Pointer.
//
// The user interface of a path control object.
func PathCellFrom(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PathCellClass) New() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathCell) Init() PathCell {
	rv := objc.Send[PathCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathCell) Autorelease() PathCell {
	rv := objc.Send[PathCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathCell creates a new PathCell instance.
func NewPathCell() PathCell {
	return getPathCellClass().New()
}



// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/controlsize

func (p_ PathCell) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](p_.ID, objc.Sel("controlSize"))
	return rv
}


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/controlsize

func (p_ PathCell) SetControlSize(value IControlSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlSize:"), value)
}


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable

func (p_ PathCell) IsEditable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable

func (p_ PathCell) SetIsEditable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable

func (p_ PathCell) IsSelectable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable

func (p_ PathCell) SetIsSelectable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSelectable:"), value)
}


// Sets the component types allowed in the path when the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/allowedtypes

func (p_ PathCell) AllowedTypes() string {
	rv := objc.Send[string](p_.ID, objc.Sel("allowedTypes"))
	return rv
}


// Sets the component types allowed in the path when the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/allowedtypes

func (p_ PathCell) SetAllowedTypes(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowedTypes:"), objc.String(value))
}


// Returns the current background color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/backgroundcolor

func (p_ PathCell) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// Returns the current background color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/backgroundcolor

func (p_ PathCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}


// Sets the value of the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/clickedpathcomponentcell

func (p_ PathCell) ClickedPathComponentCell() NSPathComponentCell {
	rv := objc.Send[NSPathComponentCell](p_.ID, objc.Sel("clickedPathComponentCell"))
	return rv
}


// Sets the value of the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/clickedpathcomponentcell

func (p_ PathCell) SetClickedPathComponentCell(value IPathComponentCell) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClickedPathComponentCell:"), value)
}


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/delegate

func (p_ PathCell) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/delegate

func (p_ PathCell) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// Sets the receiver’s double-click action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/doubleaction

func (p_ PathCell) DoubleAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("doubleAction"))
	return rv
}


// Sets the receiver’s double-click action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/doubleaction

func (p_ PathCell) SetDoubleAction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleAction:"), value)
}


// Sets the array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/pathcomponentcells

func (p_ PathCell) PathComponentCells() NSPathComponentCell {
	rv := objc.Send[NSPathComponentCell](p_.ID, objc.Sel("pathComponentCells"))
	return rv
}


// Sets the array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/pathcomponentcells

func (p_ PathCell) SetPathComponentCells(value IPathComponentCell) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathComponentCells:"), value)
}


// Sets the receiver’s path style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/pathstyle

func (p_ PathCell) PathStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pathStyle"))
	return rv
}


// Sets the receiver’s path style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/pathstyle

func (p_ PathCell) SetPathStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathStyle:"), value)
}


// Sets the value of the placeholder attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/placeholderattributedstring

func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}


// Sets the value of the placeholder attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/placeholderattributedstring

func (p_ PathCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}


// Returns the placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/placeholderstring

func (p_ PathCell) PlaceholderString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("placeholderString"))
	return rv
}


// Returns the placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/placeholderstring

func (p_ PathCell) SetPlaceholderString(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}


// Returns the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/url

func (p_ PathCell) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// Returns the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/url

func (p_ PathCell) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}



