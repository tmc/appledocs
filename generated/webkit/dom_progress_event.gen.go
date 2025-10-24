// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMProgressEvent */

/* debug [class_header]: Header for DOMProgressEvent */
// The class instance for the [DOMProgressEvent] class.
var (
	DOMProgressEventClass     _DOMProgressEventClass
	DOMProgressEventClassOnce sync.Once
)

func getDOMProgressEventClass() _DOMProgressEventClass {
	DOMProgressEventClassOnce.Do(func() {
		DOMProgressEventClass = _DOMProgressEventClass{objc.GetClass("DOMProgressEvent")}
	})
	return DOMProgressEventClass
}

type _DOMProgressEventClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMProgressEvent */
// An interface definition for the [DOMProgressEvent] class.
type IDOMProgressEvent interface {
	IDOMEvent

	/* debug [class_interface_properties]: Properties for DOMProgressEvent */
	// properties:
	LengthComputable() bool
	Loaded() uint64
	Total() uint64
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMProgressEvent */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMProgressEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMProgressEventClass) Alloc() DOMProgressEvent {
	rv := objc.Send[DOMProgressEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMProgressEventClass) New() DOMProgressEvent {
	rv := objc.Send[DOMProgressEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMProgressEvent) Init() DOMProgressEvent {
	rv := objc.Send[DOMProgressEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMProgressEvent) Autorelease() DOMProgressEvent {
	rv := objc.Send[DOMProgressEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMProgressEvent creates a new DOMProgressEvent instance.
func NewDOMProgressEvent() DOMProgressEvent {
	return getDOMProgressEventClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMProgressEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProgressEvent
type DOMProgressEvent struct {
	DOMEvent
}

// DOMProgressEventFrom constructs a [DOMProgressEvent] from an unsafe.Pointer.
func DOMProgressEventFrom(ptr unsafe.Pointer) DOMProgressEvent {
	return DOMProgressEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMProgressEvent */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMProgressEvent */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMProgressEvent */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMProgressEvent */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMProgressEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProgressEvent/lengthComputable
func (d_ DOMProgressEvent) LengthComputable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("lengthComputable"))
	return rv
} /* debug [instance_properties/getter]: lengthComputable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProgressEvent/loaded
func (d_ DOMProgressEvent) Loaded() uint64 {
	rv := objc.Send[uint64](d_.ID, objc.Sel("loaded"))
	return rv
} /* debug [instance_properties/getter]: loaded */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMProgressEvent/total
func (d_ DOMProgressEvent) Total() uint64 {
	rv := objc.Send[uint64](d_.ID, objc.Sel("total"))
	return rv
} /* debug [instance_properties/getter]: total */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMProgressEvent */
