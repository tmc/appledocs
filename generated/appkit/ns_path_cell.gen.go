// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// The user interface of a path control object.
//
// maintains a collection of objects that represent a particular path to be displayed to the user. The path shown can be set with the method. Doing so removes all displayed objects and automatically fills the control with objects set to have the appropriate icons, display titles, and values for the particular path component they represent. Alternatively, you can fill the control manually by setting the cell array or directly modifying existing cells. Both an action and double-click action can be set for the path control. To find out what path component cell was clicked in the action, you can read the value of . When the style is set to , the action is still sent, and the value for the represented menu item is correctly set. The value is valid only when the action is being sent. It is also valid when the keyboard is used to invoke the action. Automatic animated expansion of partially hidden objects happens if you correctly call and for each in the object. This is not required if the is set to , or if you wish to not have the animation. supports several path display styles. has a light blue background with arrows indicating the path. has more defined arrows (chevrons) and looks a little like a segmented button. looks and works like an object to display the full path, or, if the cell is editable, select a new path. If the cell’s method returns (the default), you can drag and drop into the cell to change the value. You can constrain what can be dropped using UTIs (Uniform Type Identifiers) with or the appropriate delegate methods on . If the cell’s method returns (the default), the cell’s contents can automatically be dragged out. The proper UTI, filename, and URL are placed on the pasteboard. You can further control or limit this by using the appropriate delegate methods on . If the cell is editable and has the path style set to , an additional item in the pop-up menu allows selecting another location. By default, an object is configured based on the allowed types. The object can be customized with a delegate method.
//
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
