// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMEvent */

/* debug [class_header]: Header for DOMEvent */
// The class instance for the [DOMEvent] class.
var (
	DOMEventClass     _DOMEventClass
	DOMEventClassOnce sync.Once
)

func getDOMEventClass() _DOMEventClass {
	DOMEventClassOnce.Do(func() {
		DOMEventClass = _DOMEventClass{objc.GetClass("DOMEvent")}
	})
	return DOMEventClass
}

type _DOMEventClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMEvent */
// An interface definition for the [DOMEvent] class.
type IDOMEvent interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMEvent */
	// properties:
	Bubbles() bool
	Cancelable() bool
	CancelBubble() bool
	SetCancelBubble(value bool)
	CurrentTarget() unsafe.Pointer
	EventPhase() unsafe.Pointer
	ReturnValue() bool
	SetReturnValue(value bool)
	SrcElement() unsafe.Pointer
	Target() unsafe.Pointer
	TimeStamp() DOMTimeStamp /* typedef */
	Type() objc.IObject      /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMEvent */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMEventClass) Alloc() DOMEvent {
	rv := objc.Send[DOMEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMEventClass) New() DOMEvent {
	rv := objc.Send[DOMEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMEvent) Init() DOMEvent {
	rv := objc.Send[DOMEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMEvent) Autorelease() DOMEvent {
	rv := objc.Send[DOMEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMEvent creates a new DOMEvent instance.
func NewDOMEvent() DOMEvent {
	return getDOMEventClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent
type DOMEvent struct {
	DOMObject
}

// DOMEventFrom constructs a [DOMEvent] from an unsafe.Pointer.
func DOMEventFrom(ptr unsafe.Pointer) DOMEvent {
	return DOMEvent{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/initEvent(_:canBubbleArg:cancelableArg:)
func NewDOMEventEventCanBubbleArgCancelableArg(eventTypeArg objc.IObject /* cross-framework: NSString */, canBubbleArg bool, cancelableArg bool) DOMEvent {
	instance := getDOMEventClass().Alloc()
	rv := objc.Send[DOMEvent](instance.ID, objc.Sel("initEvent:canBubbleArg:cancelableArg:"), eventTypeArg, canBubbleArg, cancelableArg)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewDOMEventEventCanBubbleArgCancelableArg */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMEvent */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMEvent */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMEvent */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/bubbles
func (d_ DOMEvent) Bubbles() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("bubbles"))
	return rv
} /* debug [instance_properties/getter]: bubbles */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/cancelable
func (d_ DOMEvent) Cancelable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("cancelable"))
	return rv
} /* debug [instance_properties/getter]: cancelable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/cancelBubble
func (d_ DOMEvent) CancelBubble() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("cancelBubble"))
	return rv
} /* debug [instance_properties/getter]: cancelBubble */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/cancelBubble
func (d_ DOMEvent) SetCancelBubble(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCancelBubble:"), value)
} /* debug [instance_properties/setter]: cancelBubble */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/currentTarget
func (d_ DOMEvent) CurrentTarget() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("currentTarget"))
	return rv
} /* debug [instance_properties/getter]: currentTarget */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/eventPhase
func (d_ DOMEvent) EventPhase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("eventPhase"))
	return rv
} /* debug [instance_properties/getter]: eventPhase */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/returnValue
func (d_ DOMEvent) ReturnValue() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("returnValue"))
	return rv
} /* debug [instance_properties/getter]: returnValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/returnValue
func (d_ DOMEvent) SetReturnValue(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setReturnValue:"), value)
} /* debug [instance_properties/setter]: returnValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/srcElement
func (d_ DOMEvent) SrcElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("srcElement"))
	return rv
} /* debug [instance_properties/getter]: srcElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/target
func (d_ DOMEvent) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("target"))
	return rv
} /* debug [instance_properties/getter]: target */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/timeStamp
func (d_ DOMEvent) TimeStamp() DOMTimeStamp /* typedef */ {
	rv := objc.Send[uint64](d_.ID, objc.Sel("timeStamp"))
	return rv
} /* debug [instance_properties/getter]: timeStamp */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEvent/type
func (d_ DOMEvent) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMEvent */
