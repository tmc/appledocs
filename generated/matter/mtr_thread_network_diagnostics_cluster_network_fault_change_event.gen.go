// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
// The class instance for the [MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent] class.
var (
	MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass     _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass
	MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass() _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass {
	MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass = _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent")}
	})
	return MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass
}

type _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent] class.
type IMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass) Alloc() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass) New() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) Init() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) Autorelease() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent creates a new MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent instance.
func NewMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	return getMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent
type MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventFrom constructs a [MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	return MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent/current
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent/current
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent/previous
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent/previous
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent */



