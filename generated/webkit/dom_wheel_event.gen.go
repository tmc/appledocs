// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMWheelEvent */

/* debug [class_header]: Header for DOMWheelEvent */
// The class instance for the [DOMWheelEvent] class.
var (
	DOMWheelEventClass     _DOMWheelEventClass
	DOMWheelEventClassOnce sync.Once
)

func getDOMWheelEventClass() _DOMWheelEventClass {
	DOMWheelEventClassOnce.Do(func() {
		DOMWheelEventClass = _DOMWheelEventClass{objc.GetClass("DOMWheelEvent")}
	})
	return DOMWheelEventClass
}

type _DOMWheelEventClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMWheelEvent */
// An interface definition for the [DOMWheelEvent] class.
type IDOMWheelEvent interface {
	IDOMMouseEvent

	/* debug [class_interface_properties]: Properties for DOMWheelEvent */
	// properties:
	IsHorizontal() bool
	WheelDelta() int
	WheelDeltaX() int
	WheelDeltaY() int
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMWheelEvent */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMWheelEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMWheelEventClass) Alloc() DOMWheelEvent {
	rv := objc.Send[DOMWheelEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMWheelEventClass) New() DOMWheelEvent {
	rv := objc.Send[DOMWheelEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMWheelEvent) Init() DOMWheelEvent {
	rv := objc.Send[DOMWheelEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMWheelEvent) Autorelease() DOMWheelEvent {
	rv := objc.Send[DOMWheelEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMWheelEvent creates a new DOMWheelEvent instance.
func NewDOMWheelEvent() DOMWheelEvent {
	return getDOMWheelEventClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMWheelEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMWheelEvent
type DOMWheelEvent struct {
	DOMMouseEvent
}

// DOMWheelEventFrom constructs a [DOMWheelEvent] from an unsafe.Pointer.
func DOMWheelEventFrom(ptr unsafe.Pointer) DOMWheelEvent {
	return DOMWheelEvent{
		DOMMouseEvent: DOMMouseEventFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMWheelEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMWheelEvent/initWheelEvent(_:wheelDeltaY:view:screenX:screenY:clientX:clientY:ctrlKey:altKey:shiftKey:metaKey:)
func NewDOMWheelEventWheelEventWheelDeltaYViewScreenXScreenYClientXClientYCtrlKeyAltKeyShiftKeyMetaKey(wheelDeltaX int, wheelDeltaY int, view IDOMAbstractView, screenX int, screenY int, clientX int, clientY int, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool) DOMWheelEvent {
	instance := getDOMWheelEventClass().Alloc()
	rv := objc.Send[DOMWheelEvent](instance.ID, objc.Sel("initWheelEvent:wheelDeltaY:view:screenX:screenY:clientX:clientY:ctrlKey:altKey:shiftKey:metaKey:"), wheelDeltaX, wheelDeltaY, view, screenX, screenY, clientX, clientY, ctrlKey, altKey, shiftKey, metaKey)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewDOMWheelEventWheelEventWheelDeltaYViewScreenXScreenYClientXClientYCtrlKeyAltKeyShiftKeyMetaKey */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMWheelEvent */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMWheelEvent */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMWheelEvent */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMWheelEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMWheelEvent/isHorizontal
func (d_ DOMWheelEvent) IsHorizontal() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isHorizontal"))
	return rv
} /* debug [instance_properties/getter]: isHorizontal */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMWheelEvent/wheelDelta
func (d_ DOMWheelEvent) WheelDelta() int {
	rv := objc.Send[int](d_.ID, objc.Sel("wheelDelta"))
	return rv
} /* debug [instance_properties/getter]: wheelDelta */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMWheelEvent/wheelDeltaX
func (d_ DOMWheelEvent) WheelDeltaX() int {
	rv := objc.Send[int](d_.ID, objc.Sel("wheelDeltaX"))
	return rv
} /* debug [instance_properties/getter]: wheelDeltaX */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMWheelEvent/wheelDeltaY
func (d_ DOMWheelEvent) WheelDeltaY() int {
	rv := objc.Send[int](d_.ID, objc.Sel("wheelDeltaY"))
	return rv
} /* debug [instance_properties/getter]: wheelDeltaY */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMWheelEvent */
