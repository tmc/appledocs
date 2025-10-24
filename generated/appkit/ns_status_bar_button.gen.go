// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSStatusBarButton */


/* debug [class_header]: Header for NSStatusBarButton */
// The class instance for the [StatusBarButton] class.
var (
	StatusBarButtonClass     _StatusBarButtonClass
	StatusBarButtonClassOnce sync.Once
)

func getStatusBarButtonClass() _StatusBarButtonClass {
	StatusBarButtonClassOnce.Do(func() {
		StatusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}
	})
	return StatusBarButtonClass
}

type _StatusBarButtonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StatusBarButton */
// An interface definition for the [StatusBarButton] class.
type IStatusBarButton interface {
	IButton
	
/* debug [class_interface_properties]: Properties for StatusBarButton */
	// properties:
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StatusBarButton */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StatusBarButton */
// Alloc allocates a new instance without initialization.
func (sc _StatusBarButtonClass) Alloc() StatusBarButton {
	rv := objc.Send[StatusBarButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StatusBarButtonClass) New() StatusBarButton {
	rv := objc.Send[StatusBarButton](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StatusBarButton) Init() StatusBarButton {
	rv := objc.Send[StatusBarButton](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StatusBarButton) Autorelease() StatusBarButton {
	rv := objc.Send[StatusBarButton](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStatusBarButton creates a new StatusBarButton instance.
func NewStatusBarButton() StatusBarButton {
	return getStatusBarButtonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StatusBarButton */
// The appearance and behavior of an item in the systemwide menu bar.


// The appearance and behavior of an item in the systemwide menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton
type StatusBarButton struct {
	Button
}

// StatusBarButtonFrom constructs a [StatusBarButton] from an unsafe.Pointer.
//
// The appearance and behavior of an item in the systemwide menu bar.
func StatusBarButtonFrom(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		Button: ButtonFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StatusBarButton *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StatusBarButton */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StatusBarButton */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StatusBarButton */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StatusBarButton */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton/appearsDisabled
func (s_ StatusBarButton) AppearsDisabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("appearsDisabled"))
	return rv
}/* debug [instance_properties/getter]: appearsDisabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBarButton/appearsDisabled
func (s_ StatusBarButton) SetAppearsDisabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAppearsDisabled:"), value)
}/* debug [instance_properties/setter]: appearsDisabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStatusBarButton */



