// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterRFIDEvent */


/* debug [class_header]: Header for MTREnergyEVSEClusterRFIDEvent */
// The class instance for the [MTREnergyEVSEClusterRFIDEvent] class.
var (
	MTREnergyEVSEClusterRFIDEventClass     _MTREnergyEVSEClusterRFIDEventClass
	MTREnergyEVSEClusterRFIDEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterRFIDEventClass() _MTREnergyEVSEClusterRFIDEventClass {
	MTREnergyEVSEClusterRFIDEventClassOnce.Do(func() {
		MTREnergyEVSEClusterRFIDEventClass = _MTREnergyEVSEClusterRFIDEventClass{objc.GetClass("MTREnergyEVSEClusterRFIDEvent")}
	})
	return MTREnergyEVSEClusterRFIDEventClass
}

type _MTREnergyEVSEClusterRFIDEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterRFIDEvent */
// An interface definition for the [MTREnergyEVSEClusterRFIDEvent] class.
type IMTREnergyEVSEClusterRFIDEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterRFIDEvent */
	// properties:
	Uid() objc.IObject /* cross-framework: NSData */
	SetUid(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterRFIDEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterRFIDEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterRFIDEventClass) Alloc() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterRFIDEventClass) New() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterRFIDEvent) Init() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterRFIDEvent) Autorelease() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterRFIDEvent creates a new MTREnergyEVSEClusterRFIDEvent instance.
func NewMTREnergyEVSEClusterRFIDEvent() MTREnergyEVSEClusterRFIDEvent {
	return getMTREnergyEVSEClusterRFIDEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterRFIDEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterRFIDEvent
type MTREnergyEVSEClusterRFIDEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterRFIDEventFrom constructs a [MTREnergyEVSEClusterRFIDEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterRFIDEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterRFIDEvent {
	return MTREnergyEVSEClusterRFIDEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterRFIDEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterRFIDEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterRFIDEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterRFIDEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterRFIDEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterRFIDEvent/uid
func (m_ MTREnergyEVSEClusterRFIDEvent) Uid() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("uid"))
	return rv
}/* debug [instance_properties/getter]: uid */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterRFIDEvent/uid
func (m_ MTREnergyEVSEClusterRFIDEvent) SetUid(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUid:"), value)
}/* debug [instance_properties/setter]: uid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterRFIDEvent */



