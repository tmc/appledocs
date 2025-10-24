// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterManagementClusterBoostStartedEvent */


/* debug [class_header]: Header for MTRWaterHeaterManagementClusterBoostStartedEvent */
// The class instance for the [MTRWaterHeaterManagementClusterBoostStartedEvent] class.
var (
	MTRWaterHeaterManagementClusterBoostStartedEventClass     _MTRWaterHeaterManagementClusterBoostStartedEventClass
	MTRWaterHeaterManagementClusterBoostStartedEventClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterBoostStartedEventClass() _MTRWaterHeaterManagementClusterBoostStartedEventClass {
	MTRWaterHeaterManagementClusterBoostStartedEventClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterBoostStartedEventClass = _MTRWaterHeaterManagementClusterBoostStartedEventClass{objc.GetClass("MTRWaterHeaterManagementClusterBoostStartedEvent")}
	})
	return MTRWaterHeaterManagementClusterBoostStartedEventClass
}

type _MTRWaterHeaterManagementClusterBoostStartedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterManagementClusterBoostStartedEvent */
// An interface definition for the [MTRWaterHeaterManagementClusterBoostStartedEvent] class.
type IMTRWaterHeaterManagementClusterBoostStartedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterManagementClusterBoostStartedEvent */
	// properties:
	BoostInfo() IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct
	SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterManagementClusterBoostStartedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterManagementClusterBoostStartedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterBoostStartedEventClass) Alloc() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterManagementClusterBoostStartedEventClass) New() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) Init() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) Autorelease() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterBoostStartedEvent creates a new MTRWaterHeaterManagementClusterBoostStartedEvent instance.
func NewMTRWaterHeaterManagementClusterBoostStartedEvent() MTRWaterHeaterManagementClusterBoostStartedEvent {
	return getMTRWaterHeaterManagementClusterBoostStartedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterManagementClusterBoostStartedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostStartedEvent
type MTRWaterHeaterManagementClusterBoostStartedEvent struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterBoostStartedEventFrom constructs a [MTRWaterHeaterManagementClusterBoostStartedEvent] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterBoostStartedEventFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterBoostStartedEvent {
	return MTRWaterHeaterManagementClusterBoostStartedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterManagementClusterBoostStartedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterManagementClusterBoostStartedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterManagementClusterBoostStartedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterManagementClusterBoostStartedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterManagementClusterBoostStartedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostStartedEvent/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) BoostInfo() IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("boostInfo"))
	return rv
}/* debug [instance_properties/getter]: boostInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostStartedEvent/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoostInfo:"), value)
}/* debug [instance_properties/setter]: boostInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterManagementClusterBoostStartedEvent */



