// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccountLoginClusterLoggedOutEvent */


/* debug [class_header]: Header for MTRAccountLoginClusterLoggedOutEvent */
// The class instance for the [MTRAccountLoginClusterLoggedOutEvent] class.
var (
	MTRAccountLoginClusterLoggedOutEventClass     _MTRAccountLoginClusterLoggedOutEventClass
	MTRAccountLoginClusterLoggedOutEventClassOnce sync.Once
)

func getMTRAccountLoginClusterLoggedOutEventClass() _MTRAccountLoginClusterLoggedOutEventClass {
	MTRAccountLoginClusterLoggedOutEventClassOnce.Do(func() {
		MTRAccountLoginClusterLoggedOutEventClass = _MTRAccountLoginClusterLoggedOutEventClass{objc.GetClass("MTRAccountLoginClusterLoggedOutEvent")}
	})
	return MTRAccountLoginClusterLoggedOutEventClass
}

type _MTRAccountLoginClusterLoggedOutEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccountLoginClusterLoggedOutEvent */
// An interface definition for the [MTRAccountLoginClusterLoggedOutEvent] class.
type IMTRAccountLoginClusterLoggedOutEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccountLoginClusterLoggedOutEvent */
	// properties:
	Node() objc.IObject /* cross-framework: NSNumber */
	SetNode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccountLoginClusterLoggedOutEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccountLoginClusterLoggedOutEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterLoggedOutEventClass) Alloc() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccountLoginClusterLoggedOutEventClass) New() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterLoggedOutEvent) Init() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterLoggedOutEvent) Autorelease() MTRAccountLoginClusterLoggedOutEvent {
	rv := objc.Send[MTRAccountLoginClusterLoggedOutEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterLoggedOutEvent creates a new MTRAccountLoginClusterLoggedOutEvent instance.
func NewMTRAccountLoginClusterLoggedOutEvent() MTRAccountLoginClusterLoggedOutEvent {
	return getMTRAccountLoginClusterLoggedOutEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccountLoginClusterLoggedOutEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoggedOutEvent
type MTRAccountLoginClusterLoggedOutEvent struct {
	objectivec.Object
}

// MTRAccountLoginClusterLoggedOutEventFrom constructs a [MTRAccountLoginClusterLoggedOutEvent] from an unsafe.Pointer.
func MTRAccountLoginClusterLoggedOutEventFrom(ptr unsafe.Pointer) MTRAccountLoginClusterLoggedOutEvent {
	return MTRAccountLoginClusterLoggedOutEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccountLoginClusterLoggedOutEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccountLoginClusterLoggedOutEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccountLoginClusterLoggedOutEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccountLoginClusterLoggedOutEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccountLoginClusterLoggedOutEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoggedOutEvent/node
func (m_ MTRAccountLoginClusterLoggedOutEvent) Node() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("node"))
	return rv
}/* debug [instance_properties/getter]: node */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoggedOutEvent/node
func (m_ MTRAccountLoginClusterLoggedOutEvent) SetNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}/* debug [instance_properties/setter]: node */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccountLoginClusterLoggedOutEvent */



