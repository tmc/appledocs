// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSCustomTouchBarItem */


/* debug [class_header]: Header for NSCustomTouchBarItem */
// The class instance for the [CustomTouchBarItem] class.
var (
	CustomTouchBarItemClass     _CustomTouchBarItemClass
	CustomTouchBarItemClassOnce sync.Once
)

func getCustomTouchBarItemClass() _CustomTouchBarItemClass {
	CustomTouchBarItemClassOnce.Do(func() {
		CustomTouchBarItemClass = _CustomTouchBarItemClass{objc.GetClass("NSCustomTouchBarItem")}
	})
	return CustomTouchBarItemClass
}

type _CustomTouchBarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomTouchBarItem */
// An interface definition for the [CustomTouchBarItem] class.
type ICustomTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for CustomTouchBarItem */
	// properties:
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	View() IView
	SetView(value IView)
	ViewController() IViewController
	SetViewController(value IViewController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomTouchBarItem */
// Alloc allocates a new instance without initialization.
func (cc _CustomTouchBarItemClass) Alloc() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CustomTouchBarItemClass) New() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomTouchBarItem) Init() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomTouchBarItem) Autorelease() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomTouchBarItem creates a new CustomTouchBarItem instance.
func NewCustomTouchBarItem() CustomTouchBarItem {
	return getCustomTouchBarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomTouchBarItem */
// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.


// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem
type CustomTouchBarItem struct {
	TouchBarItem
}

// CustomTouchBarItemFrom constructs a [CustomTouchBarItem] from an unsafe.Pointer.
//
// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.
func CustomTouchBarItemFrom(ptr unsafe.Pointer) CustomTouchBarItem {
	return CustomTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomTouchBarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomTouchBarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomTouchBarItem */

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/customizationLabel
func (c_ CustomTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/customizationLabel
func (c_ CustomTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// The view displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/view
func (c_ CustomTouchBarItem) View() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// The view displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/view
func (c_ CustomTouchBarItem) SetView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setView:"), value)
}/* debug [instance_properties/setter]: view */


// A view controller whose view is displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/viewController
func (c_ CustomTouchBarItem) ViewController() IViewController {
	rv := objc.Send[ViewController](c_.ID, objc.Sel("viewController"))
	return rv
}/* debug [instance_properties/getter]: viewController */


// A view controller whose view is displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/viewController
func (c_ CustomTouchBarItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setViewController:"), value)
}/* debug [instance_properties/setter]: viewController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCustomTouchBarItem */



