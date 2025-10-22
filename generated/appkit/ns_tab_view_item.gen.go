// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TabViewItem] class.
var (
	TabViewItemClass     _TabViewItemClass
	TabViewItemClassOnce sync.Once
)

func getTabViewItemClass() _TabViewItemClass {
	TabViewItemClassOnce.Do(func() {
		TabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}
	})
	return TabViewItemClass
}

type _TabViewItemClass struct {
	class objc.Class
}

// An interface definition for the [TabViewItem] class.
type ITabViewItem interface {
	objectivec.IObject
	DrawLabelInRect(shouldTruncateLabel bool, labelRect coregraphics.CGRect)
	SizeOfLabel(computeMin bool) coregraphics.CGSize
	Color() NSColor
	SetColor(value IColor)
	Identifier() objc.ID
	SetIdentifier(value objc.ID)
	Image() Image
	SetImage(value IImage)
	InitialFirstResponder() NSView
	SetInitialFirstResponder(value IView)
	Label() string
	SetLabel(value string)
	TabState() TabState
	TabView() NSTabView
	ToolTip() string
	SetToolTip(value string)
	View() NSView
	SetView(value IView)
	ViewController() NSViewController
	SetViewController(value IViewController)
}

// An item in a tab view.
//
// An is a convenient way for presenting information in multiple pages. A tab view is usually distinguished by a row of tabs that give the visual appearance of folder tabs. When the user clicks a tab, the tab view displays a view page provided by your application. A tab view keeps a zero-based array of tab view items, one for each tab in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem
type TabViewItem struct {
	objectivec.Object
}

// TabViewItemFrom constructs a [TabViewItem] from an unsafe.Pointer.
//
// An item in a tab view.
func TabViewItemFrom(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TabViewItem) Autorelease() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTabViewItem creates a new TabViewItem instance.
func NewTabViewItem() TabViewItem {
	return getTabViewItemClass().New()
}




// Performs default initialization for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(identifier:)
func NewTabViewItemWithIdentifier(identifier objectivec.IObject) TabViewItem {
	instance := getTabViewItemClass().Alloc()
	rv := objc.Send[TabViewItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func NewTabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(getTabViewItemClass().class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return rv
}

// Draws the receiver’s label in , which is the area between the curved end caps.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/drawLabel(_:in:)
func (t_ TabViewItem) DrawLabelInRect(shouldTruncateLabel bool, labelRect coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}

// Calculates the size of the receiver’s label.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/sizeOfLabel(_:)
func (t_ TabViewItem) SizeOfLabel(computeMin bool) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("sizeOfLabel:"), computeMin)
	return rv
}

// Sets the background color for content in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) Color() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the background color for content in the view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) SetColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}

// Sets the receiver’s optional identifier object to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t_ TabViewItem) Identifier() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// Sets the receiver’s optional identifier object to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t_ TabViewItem) SetIdentifier(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t_ TabViewItem) Image() Image {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t_ TabViewItem) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}

// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t_ TabViewItem) InitialFirstResponder() NSView {
	rv := objc.Send[NSView](t_.ID, objc.Sel("initialFirstResponder"))
	return rv
}


// SetInitialFirstResponder sets the value of the initialFirstResponder property.
// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInitialFirstResponder:"), value)
}

// Sets the label text for the receiver to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t_ TabViewItem) Label() string {
	rv := objc.Send[string](t_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Sets the label text for the receiver to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t_ TabViewItem) SetLabel(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Returns the current display state of the tab associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabState
func (t_ TabViewItem) TabState() TabState {
	rv := objc.Send[TabState](t_.ID, objc.Sel("tabState"))
	return rv
}

// Returns the parent tab view for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabView
func (t_ TabViewItem) TabView() NSTabView {
	rv := objc.Send[NSTabView](t_.ID, objc.Sel("tabView"))
	return rv
}

// Sets the tooltip displayed for the tab view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t_ TabViewItem) ToolTip() string {
	rv := objc.Send[string](t_.ID, objc.Sel("toolTip"))
	return rv
}


// SetToolTip sets the value of the toolTip property.
// Sets the tooltip displayed for the tab view item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t_ TabViewItem) SetToolTip(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolTip:"), objc.String(value))
}

// Sets the view associated with the receiver to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t_ TabViewItem) View() NSView {
	rv := objc.Send[NSView](t_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// Sets the view associated with the receiver to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t_ TabViewItem) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t_ TabViewItem) ViewController() NSViewController {
	rv := objc.Send[NSViewController](t_.ID, objc.Sel("viewController"))
	return rv
}


// SetViewController sets the value of the viewController property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewController:"), value)
}


