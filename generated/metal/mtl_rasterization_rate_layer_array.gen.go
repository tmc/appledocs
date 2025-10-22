// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RasterizationRateLayerArray] class.
var (
	RasterizationRateLayerArrayClass     _RasterizationRateLayerArrayClass
	RasterizationRateLayerArrayClassOnce sync.Once
)

func getRasterizationRateLayerArrayClass() _RasterizationRateLayerArrayClass {
	RasterizationRateLayerArrayClassOnce.Do(func() {
		RasterizationRateLayerArrayClass = _RasterizationRateLayerArrayClass{objc.GetClass("MTLRasterizationRateLayerArray")}
	})
	return RasterizationRateLayerArrayClass
}

type _RasterizationRateLayerArrayClass struct {
	class objc.Class
}

// An interface definition for the [RasterizationRateLayerArray] class.
type IRasterizationRateLayerArray interface {
	objectivec.IObject
	SetObjectAtIndexedSubscript(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint)
	ObjectAtIndexedSubscript(layerIndex uint) RasterizationRateLayerDescriptor
	LayerCount() int
	SetLayerCount(value int)
	Layers() MTLRasterizationRateLayerArray
	SetLayers(value IMTLRasterizationRateLayerArray)
}

// Descriptions for the rasterization rates to apply to the set of layers in a rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray
type RasterizationRateLayerArray struct {
	objectivec.Object
}

// RasterizationRateLayerArrayFrom constructs a [RasterizationRateLayerArray] from an unsafe.Pointer.
//
// Descriptions for the rasterization rates to apply to the set of layers in a rate map.
func RasterizationRateLayerArrayFrom(ptr unsafe.Pointer) RasterizationRateLayerArray {
	return RasterizationRateLayerArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateLayerArrayClass) Alloc() RasterizationRateLayerArray {
	rv := objc.Send[RasterizationRateLayerArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RasterizationRateLayerArrayClass) New() RasterizationRateLayerArray {
	rv := objc.Send[RasterizationRateLayerArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RasterizationRateLayerArray) Init() RasterizationRateLayerArray {
	rv := objc.Send[RasterizationRateLayerArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RasterizationRateLayerArray) Autorelease() RasterizationRateLayerArray {
	rv := objc.Send[RasterizationRateLayerArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRasterizationRateLayerArray creates a new RasterizationRateLayerArray instance.
func NewRasterizationRateLayerArray() RasterizationRateLayerArray {
	return getRasterizationRateLayerArrayClass().New()
}


// Stores a sample value at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray/setObject:atIndexedSubscript:
func (r_ RasterizationRateLayerArray) SetObjectAtIndexedSubscript(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), layer, layerIndex)
}

// Retrieves the sample value at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerArray/subscript(_:)
func (r_ RasterizationRateLayerArray) ObjectAtIndexedSubscript(layerIndex uint) RasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](r_.ID, objc.Sel("objectAtIndexedSubscript:"), layerIndex)
	return rv
}

// The number of layers in the rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/layercount
func (r_ RasterizationRateLayerArray) LayerCount() int {
	rv := objc.Send[int](r_.ID, objc.Sel("layerCount"))
	return rv
}


// SetLayerCount sets the value of the layerCount property.
// The number of layers in the rate map.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/layercount
func (r_ RasterizationRateLayerArray) SetLayerCount(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayerCount:"), value)
}

// The rasterization rates for one or more layers in the rate map.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/layers
func (r_ RasterizationRateLayerArray) Layers() MTLRasterizationRateLayerArray {
	rv := objc.Send[MTLRasterizationRateLayerArray](r_.ID, objc.Sel("layers"))
	return rv
}


// SetLayers sets the value of the layers property.
// The rasterization rates for one or more layers in the rate map.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratemapdescriptor/layers
func (r_ RasterizationRateLayerArray) SetLayers(value IMTLRasterizationRateLayerArray) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayers:"), value)
}



