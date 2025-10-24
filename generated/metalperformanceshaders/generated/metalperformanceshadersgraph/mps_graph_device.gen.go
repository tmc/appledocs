// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphDevice */


/* debug [class_header]: Header for MPSGraphDevice */
// The class instance for the [GraphDevice] class.
var (
	GraphDeviceClass     _GraphDeviceClass
	GraphDeviceClassOnce sync.Once
)

func getGraphDeviceClass() _GraphDeviceClass {
	GraphDeviceClassOnce.Do(func() {
		GraphDeviceClass = _GraphDeviceClass{objc.GetClass("MPSGraphDevice")}
	})
	return GraphDeviceClass
}

type _GraphDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphDevice */
// An interface definition for the [GraphDevice] class.
type IGraphDevice interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphDevice */
	// properties:
	MetalDevice() unsafe.Pointer
	Type() GraphDeviceType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphDevice */
// Alloc allocates a new instance without initialization.
func (gc _GraphDeviceClass) Alloc() GraphDevice {
	rv := objc.Send[GraphDevice](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphDeviceClass) New() GraphDevice {
	rv := objc.Send[GraphDevice](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphDevice) Init() GraphDevice {
	rv := objc.Send[GraphDevice](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphDevice) Autorelease() GraphDevice {
	rv := objc.Send[GraphDevice](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphDevice creates a new GraphDevice instance.
func NewGraphDevice() GraphDevice {
	return getGraphDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphDevice */
// A class that describes the compute device.


// A class that describes the compute device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice
type GraphDevice struct {
	GraphObject
}

// GraphDeviceFrom constructs a [GraphDevice] from an unsafe.Pointer.
//
// A class that describes the compute device.
func GraphDeviceFrom(ptr unsafe.Pointer) GraphDevice {
	return GraphDevice{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphDevice */

// Creates a device from a given Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/init(mtlDevice:)
func NewGraphDeviceWithMTLDevice(metalDevice unsafe.Pointer) GraphDevice {
	rv := objc.Send[GraphDevice](objc.ID(getGraphDeviceClass().class), objc.Sel("deviceWithMTLDevice:"), metalDevice)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphDeviceWithMTLDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphDevice */

// Creates a device from a given Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/init(mtlDevice:)
func (gc _GraphDeviceClass) DeviceWithMTLDevice(metalDevice unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("deviceWithMTLDevice:"), metalDevice)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithMTLDevice) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphDevice */

// If device type is Metal then returns the corresponding MTLDevice else nil.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/metalDevice
func (g_ GraphDevice) MetalDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("metalDevice"))
	return rv
}/* debug [instance_properties/getter]: metalDevice */


// Device of the MPSGraphDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/type
func (g_ GraphDevice) Type() GraphDeviceType {
	rv := objc.Send[GraphDeviceType](g_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphDevice */


