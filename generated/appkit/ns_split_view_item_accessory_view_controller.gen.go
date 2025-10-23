// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [SplitViewItemAccessoryViewController] class.
type ISplitViewItemAccessoryViewController interface {
	IViewController
	ViewWillDisappear()
	AutomaticallyAppliesContentInsets() bool
	SetAutomaticallyAppliesContentInsets(value bool)
	BottomAlignedAccessoryViewControllers() NSSplitViewItemAccessoryViewController
	SetBottomAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController)
	TopAlignedAccessoryViewControllers() NSSplitViewItemAccessoryViewController
	SetTopAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController)
	IsHidden() bool
	SetIsHidden(value bool)
	PreferredScrollEdgeEffectStyle() NSScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value NSScrollEdgeEffectStyle)
}



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

// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemAccessoryViewControllerClass) Alloc() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/viewWillDisappear()
func (s_ SplitViewItemAccessoryViewController) ViewWillDisappear() {
	objc.Send[objc.ID](s_.ID, objc.Sel("viewWillDisappear"))
}


// Whether or not standard content insets should be applied to the view. Defaults to YES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/automaticallyAppliesContentInsets
func (s_ SplitViewItemAccessoryViewController) AutomaticallyAppliesContentInsets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyAppliesContentInsets"))
	return rv
}


// Whether or not standard content insets should be applied to the view. Defaults to YES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/automaticallyAppliesContentInsets
func (s_ SplitViewItemAccessoryViewController) SetAutomaticallyAppliesContentInsets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyAppliesContentInsets:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/bottomalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) BottomAlignedAccessoryViewControllers() NSSplitViewItemAccessoryViewController {
	rv := objc.Send[NSSplitViewItemAccessoryViewController](s_.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/bottomalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) SetBottomAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBottomAlignedAccessoryViewControllers:"), value)
}


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/topalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) TopAlignedAccessoryViewControllers() NSSplitViewItemAccessoryViewController {
	rv := objc.Send[NSSplitViewItemAccessoryViewController](s_.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return rv
}


// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitem/topalignedaccessoryviewcontrollers
func (s_ SplitViewItemAccessoryViewController) SetTopAlignedAccessoryViewControllers(value ISplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopAlignedAccessoryViewControllers:"), value)
}


// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemaccessoryviewcontroller/ishidden
func (s_ SplitViewItemAccessoryViewController) IsHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isHidden"))
	return rv
}


// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemaccessoryviewcontroller/ishidden
func (s_ SplitViewItemAccessoryViewController) SetIsHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsHidden:"), value)
}


// The split view item accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemaccessoryviewcontroller/preferredscrolledgeeffectstyle
func (s_ SplitViewItemAccessoryViewController) PreferredScrollEdgeEffectStyle() NSScrollEdgeEffectStyle {
	rv := objc.Send[NSScrollEdgeEffectStyle](s_.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return rv
}


// The split view item accessory’s preferred effect for content scrolling behind it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewitemaccessoryviewcontroller/preferredscrolledgeeffectstyle
func (s_ SplitViewItemAccessoryViewController) SetPreferredScrollEdgeEffectStyle(value NSScrollEdgeEffectStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}



