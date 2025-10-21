// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RasterizationRateMapDescriptor] class.
var (
	RasterizationRateMapDescriptorClass     _RasterizationRateMapDescriptorClass
	RasterizationRateMapDescriptorClassOnce sync.Once
)

func getRasterizationRateMapDescriptorClass() _RasterizationRateMapDescriptorClass {
	RasterizationRateMapDescriptorClassOnce.Do(func() {
		RasterizationRateMapDescriptorClass = _RasterizationRateMapDescriptorClass{objc.GetClass("MTLRasterizationRateMapDescriptor")}
	})
	return RasterizationRateMapDescriptorClass
}

type _RasterizationRateMapDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [RasterizationRateMapDescriptor] class.
type IRasterizationRateMapDescriptor interface {
	objectivec.IObject
	LayerAtIndex(layerIndex uint) unsafe.Pointer
	SetLayerAtIndex(layer unsafe.Pointer, layerIndex uint)
}

// An object that you use to configure new rasterization rate maps.
//
// To create a new rate map, first create an instance and set its property values. Then, create a new rasterization rate-map by calling an instance’s method. When creating a rate map, Metal copies into it property values from the descriptor. You can reuse a descrptor by modifying its property values, which doesn’t affect the other rate-map instances that already exist.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor
type RasterizationRateMapDescriptor struct {
	objectivec.Object
}

// RasterizationRateMapDescriptorFrom constructs a [RasterizationRateMapDescriptor] from an unsafe.Pointer.
//
// An object that you use to configure new rasterization rate maps.
func RasterizationRateMapDescriptorFrom(ptr unsafe.Pointer) RasterizationRateMapDescriptor {
	return RasterizationRateMapDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateMapDescriptorClass) Alloc() RasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RasterizationRateMapDescriptorClass) New() RasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RasterizationRateMapDescriptor) Init() RasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RasterizationRateMapDescriptor) Autorelease() RasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRasterizationRateMapDescriptor creates a new RasterizationRateMapDescriptor instance.
func NewRasterizationRateMapDescriptor() RasterizationRateMapDescriptor {
	return getRasterizationRateMapDescriptorClass().New()
}


// Creates a rate map descriptor with a given size and identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSize(screenSize unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:"), screenSize)
	return rv
}

// Creates a rate map descriptor with a single rate layer.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:layer:
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayer(screenSize unsafe.Pointer, layer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layer:"), screenSize, layer)
	return rv
}

// Creates a rate map descriptor with a set of layer descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers(screenSize unsafe.Pointer, layerCount uint, layers unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:"), screenSize, layerCount, layers)
	return rv
}

// Returns the layer description for a layer in the rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layer(at:)
func (r_ RasterizationRateMapDescriptor) LayerAtIndex(layerIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("layerAtIndex:"), layerIndex)
	return rv
}

// Sets a configuration for a layer rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/setLayer(_:at:)
func (r_ RasterizationRateMapDescriptor) SetLayerAtIndex(layer unsafe.Pointer, layerIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayer:atIndex:"), layer, layerIndex)
}

// A string used to identify the rate map you create with the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/label
func (r_ RasterizationRateMapDescriptor) Label() string {
	rv := objc.Send[string](r_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string used to identify the rate map you create with the descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/label
func (r_ RasterizationRateMapDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), objc.String(value))
}
// The number of layers in the rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layerCount
func (r_ RasterizationRateMapDescriptor) LayerCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("layerCount"))
	return rv
}

// The rasterization rates for one or more layers in the rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layers
func (r_ RasterizationRateMapDescriptor) Layers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("layers"))
	return rv
}

// The size of the viewport coordinate system, in logical pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/screenSize
func (r_ RasterizationRateMapDescriptor) ScreenSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("screenSize"))
	return rv
}


// SetScreenSize sets the value of the screenSize property.
// The size of the viewport coordinate system, in logical pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/screenSize
func (r_ RasterizationRateMapDescriptor) SetScreenSize(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScreenSize:"), value)
}


