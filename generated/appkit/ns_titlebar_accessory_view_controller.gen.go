// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [TitlebarAccessoryViewController] class.
type ITitlebarAccessoryViewController interface {
	IViewController
	// properties:
	FullScreenMinHeight() float64 /* primitive/slice/pointer. */
	SetFullScreenMinHeight(value float64 /* primitive/slice/pointer. */)
	AutomaticallyAdjustsSize() bool /* primitive/slice/pointer. */
	SetAutomaticallyAdjustsSize(value bool /* primitive/slice/pointer. */)
	IsHidden() bool /* primitive/slice/pointer. */
	SetIsHidden(value bool /* primitive/slice/pointer. */)
	LayoutAttribute() unsafe.Pointer
	SetLayoutAttribute(value unsafe.Pointer)
	PreferredScrollEdgeEffectStyle() IScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value IScrollEdgeEffectStyle)
	FullScreenAccessoryView() IView
	SetFullScreenAccessoryView(value IView)
	View() IView
	SetView(value IView)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/fullScreenMinHeight
func (t_ TitlebarAccessoryViewController) FullScreenMinHeight() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("fullScreenMinHeight"))
	return rv
}


// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/fullScreenMinHeight
func (t_ TitlebarAccessoryViewController) SetFullScreenMinHeight(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenMinHeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/automaticallyadjustssize
func (t_ TitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallyAdjustsSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/automaticallyadjustssize
func (t_ TitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallyAdjustsSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/ishidden
func (t_ TitlebarAccessoryViewController) IsHidden() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/ishidden
func (t_ TitlebarAccessoryViewController) SetIsHidden(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHidden:"), value)
}


// The location of the accessory view, in relation to the window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/layoutattribute
func (t_ TitlebarAccessoryViewController) LayoutAttribute() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("layoutAttribute"))
	return rv
}


// The location of the accessory view, in relation to the window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/layoutattribute
func (t_ TitlebarAccessoryViewController) SetLayoutAttribute(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutAttribute:"), value)
}


// The titlebar accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/preferredscrolledgeeffectstyle
func (t_ TitlebarAccessoryViewController) PreferredScrollEdgeEffectStyle() IScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](t_.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return rv
}


// The titlebar accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/preferredscrolledgeeffectstyle
func (t_ TitlebarAccessoryViewController) SetPreferredScrollEdgeEffectStyle(value IScrollEdgeEffectStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}


// The toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/fullscreenaccessoryview
func (t_ TitlebarAccessoryViewController) FullScreenAccessoryView() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("fullScreenAccessoryView"))
	return rv
}


// The toolbar’s full screen accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbar/fullscreenaccessoryview
func (t_ TitlebarAccessoryViewController) SetFullScreenAccessoryView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFullScreenAccessoryView:"), value)
}


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (t_ TitlebarAccessoryViewController) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}


// The view controller’s primary view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/view
func (t_ TitlebarAccessoryViewController) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}



