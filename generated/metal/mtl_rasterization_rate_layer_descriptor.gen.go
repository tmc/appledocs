// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RasterizationRateLayerDescriptor] class.
var (
	RasterizationRateLayerDescriptorClass     _RasterizationRateLayerDescriptorClass
	RasterizationRateLayerDescriptorClassOnce sync.Once
)

func getRasterizationRateLayerDescriptorClass() _RasterizationRateLayerDescriptorClass {
	RasterizationRateLayerDescriptorClassOnce.Do(func() {
		RasterizationRateLayerDescriptorClass = _RasterizationRateLayerDescriptorClass{objc.GetClass("MTLRasterizationRateLayerDescriptor")}
	})
	return RasterizationRateLayerDescriptorClass
}

type _RasterizationRateLayerDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [RasterizationRateLayerDescriptor] class.
type IRasterizationRateLayerDescriptor interface {
	objectivec.IObject
}

// The minimum rasterization rates to apply to sections of a layer in the render target.
//
// Use a layer map to divide the logical viewport coordinate system into a 2D grid of equal-sized rectangles, and choose different rasterization rates for each cell. Specify rasterization rates using floating-point numbers between and , inclusive. A rate of represents the normal rasterization rate, where each logical unit is equal to a physical pixel; a rate of means that two logical units equate to one physical pixel, and so on. A value of means that the GPU renders at its lowest quality level. When you create the map, the device object chooses the nearest rasterization rate supported by the GPU that meets or exceeds the rate you specified. In the layer map, you provide separate rasterization rates for the grid’s rows and columns. The horizontal rates specify a horizontal rasterization rate for each column, and the vertical rates specify a vertical rasterization rate for each row. Each cell calculates its physical size in pixels by using the logical size of cells in the map, the horizontal rate from the cell’s column, and the vertical rate from its row.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor
type RasterizationRateLayerDescriptor struct {
	objectivec.Object
}

// RasterizationRateLayerDescriptorFrom constructs a [RasterizationRateLayerDescriptor] from an unsafe.Pointer.
//
// The minimum rasterization rates to apply to sections of a layer in the render target.
func RasterizationRateLayerDescriptorFrom(ptr unsafe.Pointer) RasterizationRateLayerDescriptor {
	return RasterizationRateLayerDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateLayerDescriptorClass) Alloc() RasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RasterizationRateLayerDescriptorClass) New() RasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RasterizationRateLayerDescriptor) Init() RasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RasterizationRateLayerDescriptor) Autorelease() RasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRasterizationRateLayerDescriptor creates a new RasterizationRateLayerDescriptor instance.
func NewRasterizationRateLayerDescriptor() RasterizationRateLayerDescriptor {
	return getRasterizationRateLayerDescriptorClass().New()
}




// Initializes the layer map with an empty grid.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/init(sampleCount:)
func NewRasterizationRateLayerDescriptorWithSampleCount(sampleCount coregraphics.ISize) RasterizationRateLayerDescriptor {
	instance := getRasterizationRateLayerDescriptorClass().Alloc()
	rv := objc.Send[RasterizationRateLayerDescriptor](instance.ID, objc.Sel("initWithSampleCount:"), sampleCount)
	rv.Autorelease()
	return rv
}



// Initializes the layer map with the provided grid size and rasterization rates.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/initWithSampleCount:horizontal:vertical:
func NewRasterizationRateLayerDescriptorWithSampleCountHorizontalVertical(sampleCount coregraphics.ISize, horizontal unsafe.Pointer, vertical unsafe.Pointer) RasterizationRateLayerDescriptor {
	instance := getRasterizationRateLayerDescriptorClass().Alloc()
	rv := objc.Send[RasterizationRateLayerDescriptor](instance.ID, objc.Sel("initWithSampleCount:horizontal:vertical:"), sampleCount, horizontal, vertical)
	rv.Autorelease()
	return rv
}


// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/horizontal
func (r_ RasterizationRateLayerDescriptor) Horizontal() MTLRasterizationRateSampleArray {
	rv := objc.Send[MTLRasterizationRateSampleArray](r_.ID, objc.Sel("horizontal"))
	return rv
}

// A pointer to the storage for the layer map’s horizontal rasterization rates.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/horizontalSampleStorage
func (r_ RasterizationRateLayerDescriptor) HorizontalSampleStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("horizontalSampleStorage"))
	return rv
}

// The maximum number of rows and columns in the layer map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/maxSampleCount
func (r_ RasterizationRateLayerDescriptor) MaxSampleCount() coregraphics.Size {
	rv := objc.Send[coregraphics.Size](r_.ID, objc.Sel("maxSampleCount"))
	return rv
}

// The number of rows and columns in the layer map.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/sampleCount
func (r_ RasterizationRateLayerDescriptor) SampleCount() coregraphics.Size {
	rv := objc.Send[coregraphics.Size](r_.ID, objc.Sel("sampleCount"))
	return rv
}

// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/vertical
func (r_ RasterizationRateLayerDescriptor) Vertical() MTLRasterizationRateSampleArray {
	rv := objc.Send[MTLRasterizationRateSampleArray](r_.ID, objc.Sel("vertical"))
	return rv
}

// A pointer to the storage for the layer map’s vertical rasterization rates.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/verticalSampleStorage
func (r_ RasterizationRateLayerDescriptor) VerticalSampleStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("verticalSampleStorage"))
	return rv
}


