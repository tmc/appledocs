// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableViewRowAction] class.
var (
	TableViewRowActionClass     _TableViewRowActionClass
	TableViewRowActionClassOnce sync.Once
)

func getTableViewRowActionClass() _TableViewRowActionClass {
	TableViewRowActionClassOnce.Do(func() {
		TableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}
	})
	return TableViewRowActionClass
}

type _TableViewRowActionClass struct {
	class objc.Class
}

// An interface definition for the [TableViewRowAction] class.
type ITableViewRowAction interface {
	objectivec.IObject
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	Image() IImage
	SetImage(value IImage)
	Style() TableViewRowActionStyle
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A single action to present when the user swipes horizontally on a table row.
//
// In an editable table, performing a horizontal swipe on a row reveals a button to delete the row by default. This class lets you define one or more custom actions to display for a given row in your table. Each instance of this class represents a single action to perform and includes the text, formatting information, and behavior for the corresponding button. To add custom actions to your table view’s rows, implement the method in your table view’s delegate object. In that method, create and return an array of actions for the specified row. The table handles the remaining work of displaying the action buttons and executing the appropriate handler block when the user clicks the button.


// A single action to present when the user swipes horizontally on a table row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction
type TableViewRowAction struct {
	objectivec.Object
}

// TableViewRowActionFrom constructs a [TableViewRowAction] from an unsafe.Pointer.
//
// A single action to present when the user swipes horizontally on a table row.
func TableViewRowActionFrom(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableViewRowAction) Autorelease() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableViewRowAction creates a new TableViewRowAction instance.
func NewTableViewRowAction() TableViewRowAction {
	return getTableViewRowActionClass().New()
}



// Creates and returns a new table view row action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/init(style:title:handler:)
func NewTableViewRowActionWithStyleTitleHandler(style TableViewRowActionStyle, title objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(getTableViewRowActionClass().class), objc.Sel("rowActionWithStyle:title:handler:"), style, title, handler)
	return rv
}



// Creates and returns a new table view row action object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/init(style:title:handler:)
func (tc _TableViewRowActionClass) RowActionWithStyleTitleHandler(style TableViewRowActionStyle, title objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("rowActionWithStyle:title:handler:"), style, title, handler)
	return rv
}


// The background color of the action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/backgroundColor
func (t_ TableViewRowAction) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/backgroundColor
func (t_ TableViewRowAction) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/image
func (t_ TableViewRowAction) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/image
func (t_ TableViewRowAction) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}


// The style applied to the action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/style-swift.property
func (t_ TableViewRowAction) Style() TableViewRowActionStyle {
	rv := objc.Send[TableViewRowActionStyle](t_.ID, objc.Sel("style"))
	return rv
}


// The title of the action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/title
func (t_ TableViewRowAction) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("title"))
	return rv
}


// The title of the action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/title
func (t_ TableViewRowAction) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), value)
}


