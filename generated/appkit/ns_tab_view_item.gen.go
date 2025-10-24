// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTabViewItem */


/* debug [class_header]: Header for NSTabViewItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TabViewItem */
// An interface definition for the [TabViewItem] class.
type ITabViewItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TabViewItem */
	// properties:
	Color() IColor
	SetColor(value IColor)
	Identifier() objc.ID
	SetIdentifier(value objc.ID)
	Image() IImage
	SetImage(value IImage)
	InitialFirstResponder() IView
	SetInitialFirstResponder(value IView)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	TabState() TabState
	TabView() ITabView
	ToolTip() objc.IObject /* cross-framework: NSString */
	SetToolTip(value objc.IObject /* cross-framework: NSString */)
	View() IView
	SetView(value IView)
	ViewController() IViewController
	SetViewController(value IViewController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TabViewItem */
	// methods:
	DrawLabelInRect(shouldTruncateLabel bool, labelRect Rect /* not a class type */)
	SizeOfLabel(computeMin bool) Size /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TabViewItem */
// Alloc allocates a new instance without initialization.
func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TabViewItem */
// An item in a tab view.
//
// An is a convenient way for presenting information in multiple pages. A tab view is usually distinguished by a row of tabs that give the visual appearance of folder tabs. When the user clicks a tab, the tab view displays a view page provided by your application. A tab view keeps a zero-based array of tab view items, one for each tab in the view.


// An item in a tab view.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TabViewItem */

// Performs default initialization for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(identifier:)
func NewTabViewItemWithIdentifier(identifier objc.IObject) TabViewItem {
	instance := getTabViewItemClass().Alloc()
	rv := objc.Send[TabViewItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTabViewItemWithIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func NewTabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(getTabViewItemClass().class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return rv
}/* debug [class_init_methods/constructor]: NewTabViewItemWithViewController */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TabViewItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TabViewItemWithViewController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TabViewItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TabViewItem */

// Draws the receiver’s label in , which is the area between the curved end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/drawLabel(_:in:)
func (t_ TabViewItem) DrawLabelInRect(shouldTruncateLabel bool, labelRect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}/* debug [instance_methods/method]: DrawLabelInRect */


// Calculates the size of the receiver’s label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/sizeOfLabel(_:)
func (t_ TabViewItem) SizeOfLabel(computeMin bool) Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("sizeOfLabel:"), computeMin)
	return rv
}/* debug [instance_methods/method]: SizeOfLabel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TabViewItem */

// Sets the background color for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) Color() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the background color for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) SetColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Sets the receiver’s optional identifier object to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t_ TabViewItem) Identifier() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Sets the receiver’s optional identifier object to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t_ TabViewItem) SetIdentifier(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t_ TabViewItem) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t_ TabViewItem) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t_ TabViewItem) InitialFirstResponder() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("initialFirstResponder"))
	return rv
}/* debug [instance_properties/getter]: initialFirstResponder */


// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInitialFirstResponder:"), value)
}/* debug [instance_properties/setter]: initialFirstResponder */


// Sets the label text for the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t_ TabViewItem) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Sets the label text for the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t_ TabViewItem) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Returns the current display state of the tab associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabState
func (t_ TabViewItem) TabState() TabState {
	rv := objc.Send[TabState](t_.ID, objc.Sel("tabState"))
	return rv
}/* debug [instance_properties/getter]: tabState */


// Returns the parent tab view for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabView
func (t_ TabViewItem) TabView() ITabView {
	rv := objc.Send[TabView](t_.ID, objc.Sel("tabView"))
	return rv
}/* debug [instance_properties/getter]: tabView */


// Sets the tooltip displayed for the tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t_ TabViewItem) ToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("toolTip"))
	return rv
}/* debug [instance_properties/getter]: toolTip */


// Sets the tooltip displayed for the tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t_ TabViewItem) SetToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolTip:"), value)
}/* debug [instance_properties/setter]: toolTip */


// Sets the view associated with the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t_ TabViewItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// Sets the view associated with the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t_ TabViewItem) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}/* debug [instance_properties/setter]: view */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t_ TabViewItem) ViewController() IViewController {
	rv := objc.Send[ViewController](t_.ID, objc.Sel("viewController"))
	return rv
}/* debug [instance_properties/getter]: viewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewController:"), value)
}/* debug [instance_properties/setter]: viewController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTabViewItem */


