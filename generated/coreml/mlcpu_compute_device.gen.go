// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCPUComputeDevice */


/* debug [class_header]: Header for MLCPUComputeDevice */
// The class instance for the [CPUComputeDevice] class.
var (
	CPUComputeDeviceClass     _CPUComputeDeviceClass
	CPUComputeDeviceClassOnce sync.Once
)

func getCPUComputeDeviceClass() _CPUComputeDeviceClass {
	CPUComputeDeviceClassOnce.Do(func() {
		CPUComputeDeviceClass = _CPUComputeDeviceClass{objc.GetClass("MLCPUComputeDevice")}
	})
	return CPUComputeDeviceClass
}

type _CPUComputeDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CPUComputeDevice */
// An interface definition for the [CPUComputeDevice] class.
type ICPUComputeDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CPUComputeDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CPUComputeDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CPUComputeDevice */
// Alloc allocates a new instance without initialization.
func (cc _CPUComputeDeviceClass) Alloc() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CPUComputeDeviceClass) New() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPUComputeDevice) Init() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPUComputeDevice) Autorelease() CPUComputeDevice {
	rv := objc.Send[CPUComputeDevice](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPUComputeDevice creates a new CPUComputeDevice instance.
func NewCPUComputeDevice() CPUComputeDevice {
	return getCPUComputeDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CPUComputeDevice */
// An object that represents a CPU compute device.


// An object that represents a CPU compute device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLCPUComputeDevice
type CPUComputeDevice struct {
	objectivec.Object
}

// CPUComputeDeviceFrom constructs a [CPUComputeDevice] from an unsafe.Pointer.
//
// An object that represents a CPU compute device.
func CPUComputeDeviceFrom(ptr unsafe.Pointer) CPUComputeDevice {
	return CPUComputeDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CPUComputeDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CPUComputeDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CPUComputeDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CPUComputeDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CPUComputeDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCPUComputeDevice */



