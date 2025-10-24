// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSStatusBar */


/* debug [class_header]: Header for NSStatusBar */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StatusBar */
// An interface definition for the [StatusBar] class.
type IStatusBar interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StatusBar */
	// properties:
	IsVertical() bool
	SetIsVertical(value bool)
	Thickness() float64
	SetThickness(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StatusBar */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StatusBar */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StatusBar */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StatusBar *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StatusBar */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StatusBar */

// Returns the system-wide status bar located in the menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar/system
func (sc _StatusBarClass) SystemStatusBar() StatusBar {
	rv := objc.Send[StatusBar](objc.ID(sc.class), objc.Sel("systemStatusBar"))
	return rv
}/* debug [class_properties_class/property]: systemStatusBar */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StatusBar */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StatusBar */

// Returns the system-wide status bar located in the menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar/system
func (s_ StatusBar) SystemStatusBar() IStatusBar {
	rv := objc.Send[StatusBar](s_.ID, objc.Sel("systemStatusBar"))
	return rv
}/* debug [instance_properties/getter]: systemStatusBar */


// A Boolean value indicating whether the status bar has a vertical orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/isvertical
func (s_ StatusBar) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}/* debug [instance_properties/getter]: isVertical */


// A Boolean value indicating whether the status bar has a vertical orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/isvertical
func (s_ StatusBar) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}/* debug [instance_properties/setter]: isVertical */


// The thickness of the status bar, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/thickness
func (s_ StatusBar) Thickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("thickness"))
	return rv
}/* debug [instance_properties/getter]: thickness */


// The thickness of the status bar, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbar/thickness
func (s_ StatusBar) SetThickness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setThickness:"), value)
}/* debug [instance_properties/setter]: thickness */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStatusBar */



