// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GraphDevice] class.
type IGraphDevice interface {
	IGraphObject
	MetalDevice() objc.ID
	Type() GraphDeviceType
}

// A class that describes the compute device.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphDeviceClass) Alloc() GraphDevice {
	rv := objc.Send[GraphDevice](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a device from a given Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/init(mtlDevice:)
func NewGraphDeviceWithMTLDevice(metalDevice objectivec.IObject) GraphDevice {
	rv := objc.Send[GraphDevice](objc.ID(getGraphDeviceClass().class), objc.Sel("deviceWithMTLDevice:"), metalDevice)
	return rv
}


// Creates a device from a given Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/init(mtlDevice:)
func (gc _GraphDeviceClass) DeviceWithMTLDevice(metalDevice objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("deviceWithMTLDevice:"), metalDevice)
	return rv
}

// If device type is Metal then returns the corresponding MTLDevice else nil.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/metalDevice
func (g_ GraphDevice) MetalDevice() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("metalDevice"))
	return rv
}

// Device of the MPSGraphDevice.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDevice/type
func (g_ GraphDevice) Type() GraphDeviceType {
	rv := objc.Send[GraphDeviceType](g_.ID, objc.Sel("type"))
	return rv
}


