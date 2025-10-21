// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A single action to present when the user swipes horizontally on a table row.
//
// In an editable table, performing a horizontal swipe on a row reveals a button to delete the row by default. This class lets you define one or more custom actions to display for a given row in your table. Each instance of this class represents a single action to perform and includes the text, formatting information, and behavior for the corresponding button. To add custom actions to your table view’s rows, implement the method in your table view’s delegate object. In that method, create and return an array of actions for the specified row. The table handles the remaining work of displaying the action buttons and executing the appropriate handler block when the user clicks the button.
//
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


// The background color of the action button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/backgroundcolor
func (t_ TableViewRowAction) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The background color of the action button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/backgroundcolor
func (t_ TableViewRowAction) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/image
func (t_ TableViewRowAction) Image() Image {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/image
func (t_ TableViewRowAction) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}

// The style applied to the action button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/style-swift.property
func (t_ TableViewRowAction) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// The style applied to the action button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/style-swift.property
func (t_ TableViewRowAction) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStyle:"), value)
}

// The title of the action button.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/title
func (t_ TableViewRowAction) Title() string {
	rv := objc.Send[string](t_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the action button.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/title
func (t_ TableViewRowAction) SetTitle(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), objc.String(value))
}



