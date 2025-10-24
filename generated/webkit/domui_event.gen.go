// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMUIEvent */

/* debug [class_header]: Header for DOMUIEvent */
// The class instance for the [DOMUIEvent] class.
var (
	DOMUIEventClass     _DOMUIEventClass
	DOMUIEventClassOnce sync.Once
)

func getDOMUIEventClass() _DOMUIEventClass {
	DOMUIEventClassOnce.Do(func() {
		DOMUIEventClass = _DOMUIEventClass{objc.GetClass("DOMUIEvent")}
	})
	return DOMUIEventClass
}

type _DOMUIEventClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMUIEvent */
// An interface definition for the [DOMUIEvent] class.
type IDOMUIEvent interface {
	IDOMEvent

	/* debug [class_interface_properties]: Properties for DOMUIEvent */
	// properties:
	CharCode() int
	Detail() int
	KeyCode() int
	PageX() int
	PageY() int
	View() IDOMAbstractView
	Which() int
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMUIEvent */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMUIEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMUIEventClass) Alloc() DOMUIEvent {
	rv := objc.Send[DOMUIEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMUIEventClass) New() DOMUIEvent {
	rv := objc.Send[DOMUIEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMUIEvent) Init() DOMUIEvent {
	rv := objc.Send[DOMUIEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMUIEvent) Autorelease() DOMUIEvent {
	rv := objc.Send[DOMUIEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMUIEvent creates a new DOMUIEvent instance.
func NewDOMUIEvent() DOMUIEvent {
	return getDOMUIEventClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMUIEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent
type DOMUIEvent struct {
	DOMEvent
}

// DOMUIEventFrom constructs a [DOMUIEvent] from an unsafe.Pointer.
func DOMUIEventFrom(ptr unsafe.Pointer) DOMUIEvent {
	return DOMUIEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMUIEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/initUIEvent(_:canBubble:cancelable:view:detail:)
func NewDOMUIEventUIEventCanBubbleCancelableViewDetail(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, detail int) DOMUIEvent {
	instance := getDOMUIEventClass().Alloc()
	rv := objc.Send[DOMUIEvent](instance.ID, objc.Sel("initUIEvent:canBubble:cancelable:view:detail:"), type_, canBubble, cancelable, view, detail)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewDOMUIEventUIEventCanBubbleCancelableViewDetail */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMUIEvent */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMUIEvent */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMUIEvent */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMUIEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/charCode
func (d_ DOMUIEvent) CharCode() int {
	rv := objc.Send[int](d_.ID, objc.Sel("charCode"))
	return rv
} /* debug [instance_properties/getter]: charCode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/detail
func (d_ DOMUIEvent) Detail() int {
	rv := objc.Send[int](d_.ID, objc.Sel("detail"))
	return rv
} /* debug [instance_properties/getter]: detail */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/keyCode
func (d_ DOMUIEvent) KeyCode() int {
	rv := objc.Send[int](d_.ID, objc.Sel("keyCode"))
	return rv
} /* debug [instance_properties/getter]: keyCode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/pageX
func (d_ DOMUIEvent) PageX() int {
	rv := objc.Send[int](d_.ID, objc.Sel("pageX"))
	return rv
} /* debug [instance_properties/getter]: pageX */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/pageY
func (d_ DOMUIEvent) PageY() int {
	rv := objc.Send[int](d_.ID, objc.Sel("pageY"))
	return rv
} /* debug [instance_properties/getter]: pageY */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/view
func (d_ DOMUIEvent) View() IDOMAbstractView {
	rv := objc.Send[DOMAbstractView](d_.ID, objc.Sel("view"))
	return rv
} /* debug [instance_properties/getter]: view */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMUIEvent/which
func (d_ DOMUIEvent) Which() int {
	rv := objc.Send[int](d_.ID, objc.Sel("which"))
	return rv
} /* debug [instance_properties/getter]: which */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMUIEvent */
