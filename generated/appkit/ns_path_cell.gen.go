// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSPathCell */


/* debug [class_header]: Header for NSPathCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PathCell */
// An interface definition for the [PathCell] class.
type IPathCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for PathCell */
	// properties:
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ClickedPathComponentCell() IPathComponentCell
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	PathComponentCells() []PathComponentCell
	SetPathComponentCells(value []PathComponentCell)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() objc.IObject /* cross-framework: NSString */
	SetPlaceholderString(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	IsEditable() bool
	SetIsEditable(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PathCell */
	// methods:
	MouseEnteredWithFrameInView(event IEvent, frame Rect /* not a class type */, view IView)
	MouseExitedWithFrameInView(event IEvent, frame Rect /* not a class type */, view IView)
	PathComponentCellAtPointWithFrameInView(point vision.Point, frame Rect /* not a class type */, view IView) IPathComponentCell
	RectOfPathComponentCellWithFrameInView(cell IPathComponentCell, frame Rect /* not a class type */, view IView) Rect /* not a class type */
	SetObjectValue(obj unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PathCell */
// Alloc allocates a new instance without initialization.
func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PathCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PathCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PathCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PathCell */

// Returns the class used to create objects when automatically filling up the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCellClass
func (pc _PathCellClass) PathComponentCellClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(pc.class), objc.Sel("pathComponentCellClass"))
	return rv
}/* debug [class_properties_class/property]: pathComponentCellClass */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PathCell */

// Displays the cell component over which the mouse is hovering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/mouseEntered(with:frame:in:)
func (p_ PathCell) MouseEnteredWithFrameInView(event IEvent, frame Rect /* not a class type */, view IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("mouseEntered:withFrame:inView:"), event, frame, view)
}/* debug [instance_methods/method]: MouseEnteredWithFrameInView */


// Hides the cell component over which the mouse is hovering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/mouseExited(with:frame:in:)
func (p_ PathCell) MouseExitedWithFrameInView(event IEvent, frame Rect /* not a class type */, view IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("mouseExited:withFrame:inView:"), event, frame, view)
}/* debug [instance_methods/method]: MouseExitedWithFrameInView */


// Returns the cell located at the given point within the given frame of the given view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCell(at:withFrame:in:)
func (p_ PathCell) PathComponentCellAtPointWithFrameInView(point vision.Point, frame Rect /* not a class type */, view IView) IPathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("pathComponentCellAtPoint:withFrame:inView:"), point, frame, view)
	return rv
}/* debug [instance_methods/method]: PathComponentCellAtPointWithFrameInView */


// Returns the current rectangle being displayed for a given path component cell, with respect to a given frame in a given view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/rect(of:withFrame:in:)
func (p_ PathCell) RectOfPathComponentCellWithFrameInView(cell IPathComponentCell, frame Rect /* not a class type */, view IView) Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("rectOfPathComponentCell:withFrame:inView:"), cell, frame, view)
	return rv
}/* debug [instance_methods/method]: RectOfPathComponentCellWithFrameInView */


// Sets the receiver’s object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/setObjectValue(_:)
func (p_ PathCell) SetObjectValue(obj unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObjectValue:"), obj)
}/* debug [instance_methods/method]: SetObjectValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PathCell */

// Sets the component types allowed in the path when the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/allowedTypes
func (p_ PathCell) AllowedTypes() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("allowedTypes"))
	return rv
}/* debug [instance_properties/getter]: allowedTypes */


// Sets the component types allowed in the path when the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/allowedTypes
func (p_ PathCell) SetAllowedTypes(value []string) {
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
}/* debug [instance_properties/setter]: allowedTypes */


// Returns the current background color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/backgroundColor
func (p_ PathCell) BackgroundColor() IColor {
	rv := objc.Send[Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// Returns the current background color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/backgroundColor
func (p_ PathCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// Sets the value of the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/clickedPathComponentCell
func (p_ PathCell) ClickedPathComponentCell() IPathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("clickedPathComponentCell"))
	return rv
}/* debug [instance_properties/getter]: clickedPathComponentCell */


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/delegate
func (p_ PathCell) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/delegate
func (p_ PathCell) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Sets the receiver’s double-click action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/doubleAction
func (p_ PathCell) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](p_.ID, objc.Sel("doubleAction"))
	return rv
}/* debug [instance_properties/getter]: doubleAction */


// Sets the receiver’s double-click action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/doubleAction
func (p_ PathCell) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDoubleAction:"), value)
}/* debug [instance_properties/setter]: doubleAction */


// Returns the class used to create objects when automatically filling up the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCellClass
func (p_ PathCell) PathComponentCellClass() objc.Class {
	rv := objc.Send[objc.Class](p_.ID, objc.Sel("pathComponentCellClass"))
	return rv
}/* debug [instance_properties/getter]: pathComponentCellClass */


// Sets the array of objects currently being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCells
func (p_ PathCell) PathComponentCells() []PathComponentCell {
	rv := objc.Send[[]PathComponentCell](p_.ID, objc.Sel("pathComponentCells"))
	return rv
}/* debug [instance_properties/getter]: pathComponentCells */


// Sets the array of objects currently being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCells
func (p_ PathCell) SetPathComponentCells(value []PathComponentCell) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathComponentCells:"), nsArray)
}/* debug [instance_properties/setter]: pathComponentCells */


// Sets the receiver’s path style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathStyle
func (p_ PathCell) PathStyle() PathStyle {
	rv := objc.Send[PathStyle](p_.ID, objc.Sel("pathStyle"))
	return rv
}/* debug [instance_properties/getter]: pathStyle */


// Sets the receiver’s path style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/pathStyle
func (p_ PathCell) SetPathStyle(value PathStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPathStyle:"), value)
}/* debug [instance_properties/setter]: pathStyle */


// Sets the value of the placeholder attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/placeholderAttributedString
func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](p_.ID, objc.Sel("placeholderAttributedString"))
	return rv
}/* debug [instance_properties/getter]: placeholderAttributedString */


// Sets the value of the placeholder attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/placeholderAttributedString
func (p_ PathCell) SetPlaceholderAttributedString(value foundation.AttributedString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}/* debug [instance_properties/setter]: placeholderAttributedString */


// Returns the placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/placeholderString
func (p_ PathCell) PlaceholderString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("placeholderString"))
	return rv
}/* debug [instance_properties/getter]: placeholderString */


// Returns the placeholder string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/placeholderString
func (p_ PathCell) SetPlaceholderString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderString:"), value)
}/* debug [instance_properties/setter]: placeholderString */


// Returns the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/url
func (p_ PathCell) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// Returns the path displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell/url
func (p_ PathCell) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/controlsize
func (p_ PathCell) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](p_.ID, objc.Sel("controlSize"))
	return rv
}/* debug [instance_properties/getter]: controlSize */


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/controlsize
func (p_ PathCell) SetControlSize(value ControlSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlSize:"), value)
}/* debug [instance_properties/setter]: controlSize */


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (p_ PathCell) IsEditable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (p_ PathCell) SetIsEditable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (p_ PathCell) IsSelectable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSelectable"))
	return rv
}/* debug [instance_properties/getter]: isSelectable */


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (p_ PathCell) SetIsSelectable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSelectable:"), value)
}/* debug [instance_properties/setter]: isSelectable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPathCell */



