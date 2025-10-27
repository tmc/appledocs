// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	Color() IColor
	SetColor(value IColor)
	Identifier() objc.ID
	SetIdentifier(value objc.ID)
	Image() IImage
	SetImage(value IImage)
	InitialFirstResponder() IView
	SetInitialFirstResponder(value IView)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	TabState() TabState
	TabView() ITabView
	ToolTip() foundation.foundation.INSString
	SetToolTip(value foundation.foundation.INSString)
	View() IView
	SetView(value IView)
	ViewController() IViewController
	SetViewController(value IViewController)


	

	// methods:
	DrawLabelInRect(shouldTruncateLabel bool, labelRect corefoundation.CGRect)
	SizeOfLabel(computeMin bool) corefoundation.CGSize


}





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






// Performs default initialization for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(identifier:)
func NewTabViewItemWithIdentifier(identifier objectivec.IObject) TabViewItem {
	instance := getTabViewItemClass().Alloc()
	rv := objc.Send[TabViewItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func NewTabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(getTabViewItemClass().class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return rv
}












// Draws the receiver’s label in , which is the area between the curved end caps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/drawLabel(_:in:)
func (t_ TabViewItem) DrawLabelInRect(shouldTruncateLabel bool, labelRect corefoundation.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}


// Calculates the size of the receiver’s label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/sizeOfLabel(_:)
func (t_ TabViewItem) SizeOfLabel(computeMin bool) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("sizeOfLabel:"), computeMin)
	return rv
}







// Sets the background color for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) Color() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("color"))
	return rv
}


// Sets the background color for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t_ TabViewItem) SetColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColor:"), value)
}


// Sets the receiver’s optional identifier object to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t_ TabViewItem) Identifier() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("identifier"))
	return rv
}


// Sets the receiver’s optional identifier object to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t_ TabViewItem) SetIdentifier(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t_ TabViewItem) Image() IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("image"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t_ TabViewItem) SetImage(value IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImage:"), value)
}


// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t_ TabViewItem) InitialFirstResponder() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("initialFirstResponder"))
	return rv
}


// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInitialFirstResponder:"), value)
}


// Sets the label text for the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t_ TabViewItem) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("label"))
	return rv
}


// Sets the label text for the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t_ TabViewItem) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), value)
}


// Returns the current display state of the tab associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabState
func (t_ TabViewItem) TabState() TabState {
	rv := objc.Send[TabState](t_.ID, objc.Sel("tabState"))
	return rv
}


// Returns the parent tab view for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabView
func (t_ TabViewItem) TabView() ITabView {
	rv := objc.Send[TabView](t_.ID, objc.Sel("tabView"))
	return rv
}


// Sets the tooltip displayed for the tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t_ TabViewItem) ToolTip() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("toolTip"))
	return rv
}


// Sets the tooltip displayed for the tab view item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t_ TabViewItem) SetToolTip(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setToolTip:"), value)
}


// Sets the view associated with the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t_ TabViewItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}


// Sets the view associated with the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t_ TabViewItem) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t_ TabViewItem) ViewController() IViewController {
	rv := objc.Send[ViewController](t_.ID, objc.Sel("viewController"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewController:"), value)
}







