// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSTitlebarAccessoryViewController */


/* debug [class_header]: Header for NSTitlebarAccessoryViewController */
// The class instance for the [TitlebarAccessoryViewController] class.
var (
	TitlebarAccessoryViewControllerClass     _TitlebarAccessoryViewControllerClass
	TitlebarAccessoryViewControllerClassOnce sync.Once
)

func getTitlebarAccessoryViewControllerClass() _TitlebarAccessoryViewControllerClass {
	TitlebarAccessoryViewControllerClassOnce.Do(func() {
		TitlebarAccessoryViewControllerClass = _TitlebarAccessoryViewControllerClass{objc.GetClass("NSTitlebarAccessoryViewController")}
	})
	return TitlebarAccessoryViewControllerClass
}

type _TitlebarAccessoryViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TitlebarAccessoryViewController */
// An interface definition for the [TitlebarAccessoryViewController] class.
type ITitlebarAccessoryViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for TitlebarAccessoryViewController */
	// properties:
	AutomaticallyAdjustsSize() bool
	SetAutomaticallyAdjustsSize(value bool)
	FullScreenMinHeight() float64
	SetFullScreenMinHeight(value float64)
	Hidden() bool
	SetHidden(value bool)
	LayoutAttribute() LayoutAttribute
	SetLayoutAttribute(value LayoutAttribute)
	PreferredScrollEdgeEffectStyle() IScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value IScrollEdgeEffectStyle)
	IsHidden() bool
	SetIsHidden(value bool)
	FullScreenAccessoryView() IView
	SetFullScreenAccessoryView(value IView)
	View() IView
	SetView(value IView)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TitlebarAccessoryViewController */
	// methods:
	ViewDidAppear()
	ViewDidDisappear()
	ViewWillAppear()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TitlebarAccessoryViewController */
// Alloc allocates a new instance without initialization.
func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TitlebarAccessoryViewControllerClass) New() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TitlebarAccessoryViewController) Autorelease() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTitlebarAccessoryViewController creates a new TitlebarAccessoryViewController instance.
func NewTitlebarAccessoryViewController() TitlebarAccessoryViewController {
	return getTitlebarAccessoryViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TitlebarAccessoryViewController */
// An object that manages a custom view—known as an accessory view—in the title bar–toolbar area of a window.
//
// Because a title bar accessory view controller is contained in a visual effect view (that is, ), it automatically handles the blur behind the accessory view and the size and location changes for the content of the view when a window goes in and out of full screen mode. If you’re currently using fullscreen accessory APIs, such as , you should use APIs instead. Typically, you create an object, give it your custom view, set the property to ensure that it displays correctly in relation to the title bar, and add the view controller to your window. For more information about methods you can use to add and remove a title bar accessory view controller, see Managing Title Bars. Don’t override the property in your subclass. Instead, you can override , and set the property in that method.


// An object that manages a custom view—known as an accessory view—in the title bar–toolbar area of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController
type TitlebarAccessoryViewController struct {
	ViewController
}

// TitlebarAccessoryViewControllerFrom constructs a [TitlebarAccessoryViewController] from an unsafe.Pointer.
//
// An object that manages a custom view—known as an accessory view—in the title bar–toolbar area of a window.
func TitlebarAccessoryViewControllerFrom(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TitlebarAccessoryViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TitlebarAccessoryViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TitlebarAccessoryViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TitlebarAccessoryViewController */

// Called when the title bar accessory view controller’s view is fully transitioned onto the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/viewDidAppear()
func (t_ TitlebarAccessoryViewController) ViewDidAppear() {
	objc.Send[objc.ID](t_.ID, objc.Sel("viewDidAppear"))
}/* debug [instance_methods/method]: ViewDidAppear */


// Called after the title bar accessory view controller’s view is removed from the window’s view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/viewDidDisappear()
func (t_ TitlebarAccessoryViewController) ViewDidDisappear() {
	objc.Send[objc.ID](t_.ID, objc.Sel("viewDidDisappear"))
}/* debug [instance_methods/method]: ViewDidDisappear */


// Called after the title bar accessory view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/viewWillAppear()
func (t_ TitlebarAccessoryViewController) ViewWillAppear() {
	objc.Send[objc.ID](t_.ID, objc.Sel("viewWillAppear"))
}/* debug [instance_methods/method]: ViewWillAppear */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TitlebarAccessoryViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/automaticallyAdjustsSize
func (t_ TitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallyAdjustsSize"))
	return rv
}/* debug [instance_properties/getter]: automaticallyAdjustsSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/automaticallyAdjustsSize
func (t_ TitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallyAdjustsSize:"), value)
}/* debug [instance_properties/setter]: automaticallyAdjustsSize */


// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/fullScreenMinHeight
func (t_ TitlebarAccessoryViewController) FullScreenMinHeight() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("fullScreenMinHeight"))
	return rv
}/* debug [instance_properties/getter]: fullScreenMinHeight */


// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/fullScreenMinHeight
func (t_ TitlebarAccessoryViewController) SetFullScreenMinHeight(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenMinHeight:"), value)
}/* debug [instance_properties/setter]: fullScreenMinHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/isHidden
func (t_ TitlebarAccessoryViewController) Hidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/isHidden
func (t_ TitlebarAccessoryViewController) SetHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// The location of the accessory view, in relation to the window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/layoutAttribute
func (t_ TitlebarAccessoryViewController) LayoutAttribute() LayoutAttribute {
	rv := objc.Send[LayoutAttribute](t_.ID, objc.Sel("layoutAttribute"))
	return rv
}/* debug [instance_properties/getter]: layoutAttribute */


// The location of the accessory view, in relation to the window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/layoutAttribute
func (t_ TitlebarAccessoryViewController) SetLayoutAttribute(value LayoutAttribute) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutAttribute:"), value)
}/* debug [instance_properties/setter]: layoutAttribute */


// The titlebar accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/preferredScrollEdgeEffectStyle
func (t_ TitlebarAccessoryViewController) PreferredScrollEdgeEffectStyle() IScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](t_.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return rv
}/* debug [instance_properties/getter]: preferredScrollEdgeEffectStyle */


// The titlebar accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/preferredScrollEdgeEffectStyle
func (t_ TitlebarAccessoryViewController) SetPreferredScrollEdgeEffectStyle(value IScrollEdgeEffectStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}/* debug [instance_properties/setter]: preferredScrollEdgeEffectStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/ishidden
func (t_ TitlebarAccessoryViewController) IsHidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/ishidden
func (t_ TitlebarAccessoryViewController) SetIsHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */


// The toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/fullscreenaccessoryview
func (t_ TitlebarAccessoryViewController) FullScreenAccessoryView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("fullScreenAccessoryView"))
	return rv
}/* debug [instance_properties/getter]: fullScreenAccessoryView */


// The toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/fullscreenaccessoryview
func (t_ TitlebarAccessoryViewController) SetFullScreenAccessoryView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryView:"), value)
}/* debug [instance_properties/setter]: fullScreenAccessoryView */


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (t_ TitlebarAccessoryViewController) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (t_ TitlebarAccessoryViewController) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}/* debug [instance_properties/setter]: view */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTitlebarAccessoryViewController */



