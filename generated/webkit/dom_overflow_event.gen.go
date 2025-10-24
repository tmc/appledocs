// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMOverflowEvent */


/* debug [class_header]: Header for DOMOverflowEvent */
// The class instance for the [DOMOverflowEvent] class.
var (
	DOMOverflowEventClass     _DOMOverflowEventClass
	DOMOverflowEventClassOnce sync.Once
)

func getDOMOverflowEventClass() _DOMOverflowEventClass {
	DOMOverflowEventClassOnce.Do(func() {
		DOMOverflowEventClass = _DOMOverflowEventClass{objc.GetClass("DOMOverflowEvent")}
	})
	return DOMOverflowEventClass
}

type _DOMOverflowEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMOverflowEvent */
// An interface definition for the [DOMOverflowEvent] class.
type IDOMOverflowEvent interface {
	IDOMEvent
	
/* debug [class_interface_properties]: Properties for DOMOverflowEvent */
	// properties:
	HorizontalOverflow() bool
	Orient() objectivec.IObject
	VerticalOverflow() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMOverflowEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMOverflowEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMOverflowEventClass) Alloc() DOMOverflowEvent {
	rv := objc.Send[DOMOverflowEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMOverflowEventClass) New() DOMOverflowEvent {
	rv := objc.Send[DOMOverflowEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMOverflowEvent) Init() DOMOverflowEvent {
	rv := objc.Send[DOMOverflowEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMOverflowEvent) Autorelease() DOMOverflowEvent {
	rv := objc.Send[DOMOverflowEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMOverflowEvent creates a new DOMOverflowEvent instance.
func NewDOMOverflowEvent() DOMOverflowEvent {
	return getDOMOverflowEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMOverflowEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMOverflowEvent
type DOMOverflowEvent struct {
	DOMEvent
}

// DOMOverflowEventFrom constructs a [DOMOverflowEvent] from an unsafe.Pointer.
func DOMOverflowEventFrom(ptr unsafe.Pointer) DOMOverflowEvent {
	return DOMOverflowEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMOverflowEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMOverflowEvent/initOverflowEvent(_:horizontalOverflow:verticalOverflow:)
func NewDOMOverflowEventOverflowEventHorizontalOverflowVerticalOverflow(orient objectivec.IObject, horizontalOverflow bool, verticalOverflow bool) DOMOverflowEvent {
	instance := getDOMOverflowEventClass().Alloc()
	rv := objc.Send[DOMOverflowEvent](instance.ID, objc.Sel("initOverflowEvent:horizontalOverflow:verticalOverflow:"), orient, horizontalOverflow, verticalOverflow)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMOverflowEventOverflowEventHorizontalOverflowVerticalOverflow */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMOverflowEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMOverflowEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMOverflowEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMOverflowEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMOverflowEvent/horizontalOverflow
func (d_ DOMOverflowEvent) HorizontalOverflow() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("horizontalOverflow"))
	return rv
}/* debug [instance_properties/getter]: horizontalOverflow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMOverflowEvent/orient
func (d_ DOMOverflowEvent) Orient() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("orient"))
	return rv
}/* debug [instance_properties/getter]: orient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMOverflowEvent/verticalOverflow
func (d_ DOMOverflowEvent) VerticalOverflow() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("verticalOverflow"))
	return rv
}/* debug [instance_properties/getter]: verticalOverflow */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMOverflowEvent */


