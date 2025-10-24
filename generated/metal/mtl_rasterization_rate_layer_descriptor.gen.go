// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRasterizationRateLayerDescriptor */


/* debug [class_header]: Header for MTLRasterizationRateLayerDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RasterizationRateLayerDescriptor */
// An interface definition for the [RasterizationRateLayerDescriptor] class.
type IRasterizationRateLayerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RasterizationRateLayerDescriptor */
	// properties:
	Horizontal() IMTLRasterizationRateSampleArray
	HorizontalSampleStorage() objectivec.IObject
	MaxSampleCount() objc.IObject /* cross-framework: MTLSize */
	SampleCount() objc.IObject /* cross-framework: MTLSize */
	Vertical() IMTLRasterizationRateSampleArray
	VerticalSampleStorage() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RasterizationRateLayerDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RasterizationRateLayerDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateLayerDescriptorClass) Alloc() RasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RasterizationRateLayerDescriptor */
// The minimum rasterization rates to apply to sections of a layer in the render target.
//
// Use a layer map to divide the logical viewport coordinate system into a 2D grid of equal-sized rectangles, and choose different rasterization rates for each cell. Specify rasterization rates using floating-point numbers between and , inclusive. A rate of represents the normal rasterization rate, where each logical unit is equal to a physical pixel; a rate of means that two logical units equate to one physical pixel, and so on. A value of means that the GPU renders at its lowest quality level. When you create the map, the device object chooses the nearest rasterization rate supported by the GPU that meets or exceeds the rate you specified. In the layer map, you provide separate rasterization rates for the grid’s rows and columns. The horizontal rates specify a horizontal rasterization rate for each column, and the vertical rates specify a vertical rasterization rate for each row. Each cell calculates its physical size in pixels by using the logical size of cells in the map, the horizontal rate from the cell’s column, and the vertical rate from its row.


// The minimum rasterization rates to apply to sections of a layer in the render target.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RasterizationRateLayerDescriptor */

// Initializes the layer map with an empty grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/init(sampleCount:)
func NewRasterizationRateLayerDescriptorWithSampleCount(sampleCount objc.IObject /* cross-framework: MTLSize */) RasterizationRateLayerDescriptor {
	instance := getRasterizationRateLayerDescriptorClass().Alloc()
	rv := objc.Send[RasterizationRateLayerDescriptor](instance.ID, objc.Sel("initWithSampleCount:"), sampleCount)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRasterizationRateLayerDescriptorWithSampleCount */


// Initializes the layer map with the provided grid size and rasterization rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/initWithSampleCount:horizontal:vertical:
func NewRasterizationRateLayerDescriptorWithSampleCountHorizontalVertical(sampleCount objc.IObject /* cross-framework: MTLSize */, horizontal objectivec.IObject, vertical objectivec.IObject) RasterizationRateLayerDescriptor {
	instance := getRasterizationRateLayerDescriptorClass().Alloc()
	rv := objc.Send[RasterizationRateLayerDescriptor](instance.ID, objc.Sel("initWithSampleCount:horizontal:vertical:"), sampleCount, horizontal, vertical)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRasterizationRateLayerDescriptorWithSampleCountHorizontalVertical */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RasterizationRateLayerDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RasterizationRateLayerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RasterizationRateLayerDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RasterizationRateLayerDescriptor */

// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/horizontal
func (r_ RasterizationRateLayerDescriptor) Horizontal() IMTLRasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("horizontal"))
	return rv
}/* debug [instance_properties/getter]: horizontal */


// A pointer to the storage for the layer map’s horizontal rasterization rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/horizontalSampleStorage
func (r_ RasterizationRateLayerDescriptor) HorizontalSampleStorage() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("horizontalSampleStorage"))
	return rv
}/* debug [instance_properties/getter]: horizontalSampleStorage */


// The maximum number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/maxSampleCount
func (r_ RasterizationRateLayerDescriptor) MaxSampleCount() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("maxSampleCount"))
	return rv
}/* debug [instance_properties/getter]: maxSampleCount */


// The number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/sampleCount
func (r_ RasterizationRateLayerDescriptor) SampleCount() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/vertical
func (r_ RasterizationRateLayerDescriptor) Vertical() IMTLRasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("vertical"))
	return rv
}/* debug [instance_properties/getter]: vertical */


// A pointer to the storage for the layer map’s vertical rasterization rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateLayerDescriptor/verticalSampleStorage
func (r_ RasterizationRateLayerDescriptor) VerticalSampleStorage() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("verticalSampleStorage"))
	return rv
}/* debug [instance_properties/getter]: verticalSampleStorage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRasterizationRateLayerDescriptor */


