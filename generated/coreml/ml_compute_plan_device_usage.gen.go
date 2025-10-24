// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLComputePlanDeviceUsage */


/* debug [class_header]: Header for MLComputePlanDeviceUsage */
// The class instance for the [ComputePlanDeviceUsage] class.
var (
	ComputePlanDeviceUsageClass     _ComputePlanDeviceUsageClass
	ComputePlanDeviceUsageClassOnce sync.Once
)

func getComputePlanDeviceUsageClass() _ComputePlanDeviceUsageClass {
	ComputePlanDeviceUsageClassOnce.Do(func() {
		ComputePlanDeviceUsageClass = _ComputePlanDeviceUsageClass{objc.GetClass("MLComputePlanDeviceUsage")}
	})
	return ComputePlanDeviceUsageClass
}

type _ComputePlanDeviceUsageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePlanDeviceUsage */
// An interface definition for the [ComputePlanDeviceUsage] class.
type IComputePlanDeviceUsage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePlanDeviceUsage */
	// properties:
	PreferredComputeDevice() unsafe.Pointer
	SupportedComputeDevices() []objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePlanDeviceUsage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePlanDeviceUsage */
// Alloc allocates a new instance without initialization.
func (cc _ComputePlanDeviceUsageClass) Alloc() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComputePlanDeviceUsageClass) New() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePlanDeviceUsage) Init() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePlanDeviceUsage) Autorelease() ComputePlanDeviceUsage {
	rv := objc.Send[ComputePlanDeviceUsage](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePlanDeviceUsage creates a new ComputePlanDeviceUsage instance.
func NewComputePlanDeviceUsage() ComputePlanDeviceUsage {
	return getComputePlanDeviceUsageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePlanDeviceUsage */
// The anticipated compute devices to use for executing a layer or operation.


// The anticipated compute devices to use for executing a layer or operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage
type ComputePlanDeviceUsage struct {
	objectivec.Object
}

// ComputePlanDeviceUsageFrom constructs a [ComputePlanDeviceUsage] from an unsafe.Pointer.
//
// The anticipated compute devices to use for executing a layer or operation.
func ComputePlanDeviceUsageFrom(ptr unsafe.Pointer) ComputePlanDeviceUsage {
	return ComputePlanDeviceUsage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePlanDeviceUsage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePlanDeviceUsage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePlanDeviceUsage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePlanDeviceUsage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePlanDeviceUsage */

// The compute device that the framework prefers to execute the layer/operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/preferredComputeDevice
func (c_ ComputePlanDeviceUsage) PreferredComputeDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preferredComputeDevice"))
	return rv
}/* debug [instance_properties/getter]: preferredComputeDevice */


// The compute devices that can execute the layer/operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/supportedComputeDevices
func (c_ ComputePlanDeviceUsage) SupportedComputeDevices() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("supportedComputeDevices"))
	return rv
}/* debug [instance_properties/getter]: supportedComputeDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLComputePlanDeviceUsage */



