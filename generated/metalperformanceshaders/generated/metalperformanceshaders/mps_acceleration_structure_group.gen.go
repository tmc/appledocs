// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSAccelerationStructureGroup */


/* debug [class_header]: Header for MPSAccelerationStructureGroup */
// The class instance for the [AccelerationStructureGroup] class.
var (
	AccelerationStructureGroupClass     _AccelerationStructureGroupClass
	AccelerationStructureGroupClassOnce sync.Once
)

func getAccelerationStructureGroupClass() _AccelerationStructureGroupClass {
	AccelerationStructureGroupClassOnce.Do(func() {
		AccelerationStructureGroupClass = _AccelerationStructureGroupClass{objc.GetClass("MPSAccelerationStructureGroup")}
	})
	return AccelerationStructureGroupClass
}

type _AccelerationStructureGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructureGroup */
// An interface definition for the [AccelerationStructureGroup] class.
type IAccelerationStructureGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccelerationStructureGroup */
	// properties:
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructureGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructureGroup */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureGroupClass) Alloc() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructureGroupClass) New() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureGroup) Init() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureGroup) Autorelease() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureGroup creates a new AccelerationStructureGroup instance.
func NewAccelerationStructureGroup() AccelerationStructureGroup {
	return getAccelerationStructureGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructureGroup */
// A group of acceleration structures.


// A group of acceleration structures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup
type AccelerationStructureGroup struct {
	objectivec.Object
}

// AccelerationStructureGroupFrom constructs a [AccelerationStructureGroup] from an unsafe.Pointer.
//
// A group of acceleration structures.
func AccelerationStructureGroupFrom(ptr unsafe.Pointer) AccelerationStructureGroup {
	return AccelerationStructureGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructureGroup */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructuregroup/2980784-initwithdevice
func NewAccelerationStructureGroupWithDevice(device unsafe.Pointer) AccelerationStructureGroup {
	instance := getAccelerationStructureGroupClass().Alloc()
	rv := objc.Send[AccelerationStructureGroup](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccelerationStructureGroupWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructureGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructureGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructureGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructureGroup */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructuregroup/2980783-device
func (a_ AccelerationStructureGroup) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructuregroup/2980783-device
func (a_ AccelerationStructureGroup) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSAccelerationStructureGroup */


