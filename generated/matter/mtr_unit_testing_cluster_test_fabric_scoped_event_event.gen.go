// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestFabricScopedEventEvent */


/* debug [class_header]: Header for MTRUnitTestingClusterTestFabricScopedEventEvent */
// The class instance for the [MTRUnitTestingClusterTestFabricScopedEventEvent] class.
var (
	MTRUnitTestingClusterTestFabricScopedEventEventClass     _MTRUnitTestingClusterTestFabricScopedEventEventClass
	MTRUnitTestingClusterTestFabricScopedEventEventClassOnce sync.Once
)

func getMTRUnitTestingClusterTestFabricScopedEventEventClass() _MTRUnitTestingClusterTestFabricScopedEventEventClass {
	MTRUnitTestingClusterTestFabricScopedEventEventClassOnce.Do(func() {
		MTRUnitTestingClusterTestFabricScopedEventEventClass = _MTRUnitTestingClusterTestFabricScopedEventEventClass{objc.GetClass("MTRUnitTestingClusterTestFabricScopedEventEvent")}
	})
	return MTRUnitTestingClusterTestFabricScopedEventEventClass
}

type _MTRUnitTestingClusterTestFabricScopedEventEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestFabricScopedEventEvent */
// An interface definition for the [MTRUnitTestingClusterTestFabricScopedEventEvent] class.
type IMTRUnitTestingClusterTestFabricScopedEventEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestFabricScopedEventEvent */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestFabricScopedEventEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestFabricScopedEventEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestFabricScopedEventEventClass) Alloc() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestFabricScopedEventEventClass) New() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestFabricScopedEventEvent) Init() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestFabricScopedEventEvent) Autorelease() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestFabricScopedEventEvent creates a new MTRUnitTestingClusterTestFabricScopedEventEvent instance.
func NewMTRUnitTestingClusterTestFabricScopedEventEvent() MTRUnitTestingClusterTestFabricScopedEventEvent {
	return getMTRUnitTestingClusterTestFabricScopedEventEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestFabricScopedEventEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScopedEventEvent
type MTRUnitTestingClusterTestFabricScopedEventEvent struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestFabricScopedEventEventFrom constructs a [MTRUnitTestingClusterTestFabricScopedEventEvent] from an unsafe.Pointer.
func MTRUnitTestingClusterTestFabricScopedEventEventFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestFabricScopedEventEvent {
	return MTRUnitTestingClusterTestFabricScopedEventEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestFabricScopedEventEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestFabricScopedEventEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestFabricScopedEventEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestFabricScopedEventEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestFabricScopedEventEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScopedEventEvent/fabricIndex
func (m_ MTRUnitTestingClusterTestFabricScopedEventEvent) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScopedEventEvent/fabricIndex
func (m_ MTRUnitTestingClusterTestFabricScopedEventEvent) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestFabricScopedEventEvent */



