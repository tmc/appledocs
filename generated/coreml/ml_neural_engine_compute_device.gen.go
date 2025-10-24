// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLNeuralEngineComputeDevice */


/* debug [class_header]: Header for MLNeuralEngineComputeDevice */
// The class instance for the [NeuralEngineComputeDevice] class.
var (
	NeuralEngineComputeDeviceClass     _NeuralEngineComputeDeviceClass
	NeuralEngineComputeDeviceClassOnce sync.Once
)

func getNeuralEngineComputeDeviceClass() _NeuralEngineComputeDeviceClass {
	NeuralEngineComputeDeviceClassOnce.Do(func() {
		NeuralEngineComputeDeviceClass = _NeuralEngineComputeDeviceClass{objc.GetClass("MLNeuralEngineComputeDevice")}
	})
	return NeuralEngineComputeDeviceClass
}

type _NeuralEngineComputeDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NeuralEngineComputeDevice */
// An interface definition for the [NeuralEngineComputeDevice] class.
type INeuralEngineComputeDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NeuralEngineComputeDevice */
	// properties:
	TotalCoreCount() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NeuralEngineComputeDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NeuralEngineComputeDevice */
// Alloc allocates a new instance without initialization.
func (nc _NeuralEngineComputeDeviceClass) Alloc() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NeuralEngineComputeDeviceClass) New() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NeuralEngineComputeDevice) Init() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NeuralEngineComputeDevice) Autorelease() NeuralEngineComputeDevice {
	rv := objc.Send[NeuralEngineComputeDevice](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralEngineComputeDevice creates a new NeuralEngineComputeDevice instance.
func NewNeuralEngineComputeDevice() NeuralEngineComputeDevice {
	return getNeuralEngineComputeDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NeuralEngineComputeDevice */
// An object that represents a Neural Engine compute device.


// An object that represents a Neural Engine compute device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice
type NeuralEngineComputeDevice struct {
	objectivec.Object
}

// NeuralEngineComputeDeviceFrom constructs a [NeuralEngineComputeDevice] from an unsafe.Pointer.
//
// An object that represents a Neural Engine compute device.
func NeuralEngineComputeDeviceFrom(ptr unsafe.Pointer) NeuralEngineComputeDevice {
	return NeuralEngineComputeDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NeuralEngineComputeDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NeuralEngineComputeDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NeuralEngineComputeDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NeuralEngineComputeDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NeuralEngineComputeDevice */

// The total number of cores in the Neural Engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDevice/totalCoreCount
func (n_ NeuralEngineComputeDevice) TotalCoreCount() int {
	rv := objc.Send[int](n_.ID, objc.Sel("totalCoreCount"))
	return rv
}/* debug [instance_properties/getter]: totalCoreCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLNeuralEngineComputeDevice */



