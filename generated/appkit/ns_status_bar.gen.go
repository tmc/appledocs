// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [StatusBar] class.
var (
	StatusBarClass     _StatusBarClass
	StatusBarClassOnce sync.Once
)

func getStatusBarClass() _StatusBarClass {
	StatusBarClassOnce.Do(func() {
		StatusBarClass = _StatusBarClass{objc.GetClass("NSStatusBar")}
	})
	return StatusBarClass
}

type _StatusBarClass struct {
	class objc.Class
}





// An interface definition for the [StatusBar] class.
type IStatusBar interface {
	objectivec.IObject
	

	// properties:
	IsVertical() bool
	SetIsVertical(value bool)
	Thickness() float64
	SetThickness(value float64)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _StatusBarClass) Alloc() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StatusBarClass) New() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatusBar) Init() StatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatusBar) Autorelease() StatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatusBar creates a new StatusBar instance.
func NewStatusBar() StatusBar {
	return getStatusBarClass().New()
}





// An object that manages a collection of status items displayed within the system-wide menu bar.
//
// A status item (an instance of ) can be displayed with text or an icon, can provide a menu and a target-action message when clicked, or can be a fully customized view that you create. Use status items sparingly and only if the alternatives (such as a Dock menu, preference pane, or status window) are not suitable. Because there is limited space in which to display status items, status items are not guaranteed to be available at all times. For this reason, do not rely on them being available and always provide a user preference for hiding your application’s status items to free up space in the menu bar.


// An object that manages a collection of status items displayed within the system-wide menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar
type StatusBar struct {
	objectivec.Object
}

// StatusBarFrom constructs a [StatusBar] from an unsafe.Pointer.
//
// An object that manages a collection of status items displayed within the system-wide menu bar.
func StatusBarFrom(ptr unsafe.Pointer) StatusBar {
	return StatusBar{objectivec.Object{objc.ID(ptr)}}
}















// Returns the system-wide status bar located in the menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar/system
func (sc _StatusBarClass) SystemStatusBar() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.class), objc.Sel("systemStatusBar"))
	return rv
}











// Returns the system-wide status bar located in the menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar/system
func (s_ StatusBar) SystemStatusBar() IStatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("systemStatusBar"))
	return rv
}


// A Boolean value indicating whether the status bar has a vertical orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/isvertical
func (s_ StatusBar) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}


// A Boolean value indicating whether the status bar has a vertical orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/isvertical
func (s_ StatusBar) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}


// The thickness of the status bar, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/thickness
func (s_ StatusBar) Thickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("thickness"))
	return rv
}


// The thickness of the status bar, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/thickness
func (s_ StatusBar) SetThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setThickness:"), value)
}








