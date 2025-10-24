// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSplitViewItemAccessoryViewController */


/* debug [class_header]: Header for NSSplitViewItemAccessoryViewController */
// The class instance for the [SplitViewItemAccessoryViewController] class.
var (
	SplitViewItemAccessoryViewControllerClass     _SplitViewItemAccessoryViewControllerClass
	SplitViewItemAccessoryViewControllerClassOnce sync.Once
)

func getSplitViewItemAccessoryViewControllerClass() _SplitViewItemAccessoryViewControllerClass {
	SplitViewItemAccessoryViewControllerClassOnce.Do(func() {
		SplitViewItemAccessoryViewControllerClass = _SplitViewItemAccessoryViewControllerClass{objc.GetClass("NSSplitViewItemAccessoryViewController")}
	})
	return SplitViewItemAccessoryViewControllerClass
}

type _SplitViewItemAccessoryViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SplitViewItemAccessoryViewController */
// An interface definition for the [SplitViewItemAccessoryViewController] class.
type ISplitViewItemAccessoryViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for SplitViewItemAccessoryViewController */
	// properties:
	AutomaticallyAppliesContentInsets() bool
	SetAutomaticallyAppliesContentInsets(value bool)
	Hidden() bool
	SetHidden(value bool)
	PreferredScrollEdgeEffectStyle() IScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value IScrollEdgeEffectStyle)
	BottomAlignedAccessoryViewControllers() ISplitViewItemAccessoryViewController
	SetBottomAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController)
	TopAlignedAccessoryViewControllers() ISplitViewItemAccessoryViewController
	SetTopAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController)
	IsHidden() bool
	SetIsHidden(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SplitViewItemAccessoryViewController */
	// methods:
	ViewDidAppear()
	ViewDidDisappear()
	ViewWillAppear()
	ViewWillDisappear()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SplitViewItemAccessoryViewController */
// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemAccessoryViewControllerClass) Alloc() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SplitViewItemAccessoryViewControllerClass) New() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitViewItemAccessoryViewController) Init() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitViewItemAccessoryViewController) Autorelease() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitViewItemAccessoryViewController creates a new SplitViewItemAccessoryViewController instance.
func NewSplitViewItemAccessoryViewController() SplitViewItemAccessoryViewController {
	return getSplitViewItemAccessoryViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SplitViewItemAccessoryViewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController
type SplitViewItemAccessoryViewController struct {
	ViewController
}

// SplitViewItemAccessoryViewControllerFrom constructs a [SplitViewItemAccessoryViewController] from an unsafe.Pointer.
func SplitViewItemAccessoryViewControllerFrom(ptr unsafe.Pointer) SplitViewItemAccessoryViewController {
	return SplitViewItemAccessoryViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SplitViewItemAccessoryViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SplitViewItemAccessoryViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SplitViewItemAccessoryViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SplitViewItemAccessoryViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/viewDidAppear()
func (s_ SplitViewItemAccessoryViewController) ViewDidAppear() {
	objc.Send[objc.ID](s_.ID, objc.Sel("viewDidAppear"))
}/* debug [instance_methods/method]: ViewDidAppear */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/viewDidDisappear()
func (s_ SplitViewItemAccessoryViewController) ViewDidDisappear() {
	objc.Send[objc.ID](s_.ID, objc.Sel("viewDidDisappear"))
}/* debug [instance_methods/method]: ViewDidDisappear */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/viewWillAppear()
func (s_ SplitViewItemAccessoryViewController) ViewWillAppear() {
	objc.Send[objc.ID](s_.ID, objc.Sel("viewWillAppear"))
}/* debug [instance_methods/method]: ViewWillAppear */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/viewWillDisappear()
func (s_ SplitViewItemAccessoryViewController) ViewWillDisappear() {
	objc.Send[objc.ID](s_.ID, objc.Sel("viewWillDisappear"))
}/* debug [instance_methods/method]: ViewWillDisappear */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SplitViewItemAccessoryViewController */

// Whether or not standard content insets should be applied to the view. Defaults to YES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/automaticallyAppliesContentInsets
func (s_ SplitViewItemAccessoryViewController) AutomaticallyAppliesContentInsets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyAppliesContentInsets"))
	return rv
}/* debug [instance_properties/getter]: automaticallyAppliesContentInsets */


// Whether or not standard content insets should be applied to the view. Defaults to YES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/automaticallyAppliesContentInsets
func (s_ SplitViewItemAccessoryViewController) SetAutomaticallyAppliesContentInsets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyAppliesContentInsets:"), value)
}/* debug [instance_properties/setter]: automaticallyAppliesContentInsets */


// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window. Set through the animator object to animate it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/isHidden
func (s_ SplitViewItemAccessoryViewController) Hidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window. Set through the animator object to animate it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/isHidden
func (s_ SplitViewItemAccessoryViewController) SetHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// The split view item accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/preferredScrollEdgeEffectStyle
func (s_ SplitViewItemAccessoryViewController) PreferredScrollEdgeEffectStyle() IScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return rv
}/* debug [instance_properties/getter]: preferredScrollEdgeEffectStyle */


// The split view item accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/preferredScrollEdgeEffectStyle
func (s_ SplitViewItemAccessoryViewController) SetPreferredScrollEdgeEffectStyle(value IScrollEdgeEffectStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}/* debug [instance_properties/setter]: preferredScrollEdgeEffectStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/bottomalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) BottomAlignedAccessoryViewControllers() ISplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return rv
}/* debug [instance_properties/getter]: bottomAlignedAccessoryViewControllers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/bottomalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) SetBottomAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBottomAlignedAccessoryViewControllers:"), value)
}/* debug [instance_properties/setter]: bottomAlignedAccessoryViewControllers */


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/topalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) TopAlignedAccessoryViewControllers() ISplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return rv
}/* debug [instance_properties/getter]: topAlignedAccessoryViewControllers */


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/topalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) SetTopAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopAlignedAccessoryViewControllers:"), value)
}/* debug [instance_properties/setter]: topAlignedAccessoryViewControllers */


// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemaccessoryviewcontroller/ishidden
func (s_ SplitViewItemAccessoryViewController) IsHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemaccessoryviewcontroller/ishidden
func (s_ SplitViewItemAccessoryViewController) SetIsHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSplitViewItemAccessoryViewController */



