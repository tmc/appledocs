// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRasterizationRateMapDescriptor */


/* debug [class_header]: Header for MTLRasterizationRateMapDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RasterizationRateMapDescriptor */
// An interface definition for the [RasterizationRateMapDescriptor] class.
type IRasterizationRateMapDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RasterizationRateMapDescriptor */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	LayerCount() uint
	Layers() IMTLRasterizationRateLayerArray
	ScreenSize() objc.IObject /* cross-framework: MTLSize */
	SetScreenSize(value objc.IObject /* cross-framework: MTLSize */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RasterizationRateMapDescriptor */
	// methods:
	LayerAtIndex(layerIndex uint) IRasterizationRateLayerDescriptor
	SetLayerAtIndex(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RasterizationRateMapDescriptor */
// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateMapDescriptorClass) Alloc() RasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RasterizationRateMapDescriptor */
// An object that you use to configure new rasterization rate maps.
//
// To create a new rate map, first create an instance and set its property values. Then, create a new rasterization rate-map by calling an instance’s method. When creating a rate map, Metal copies into it property values from the descriptor. You can reuse a descrptor by modifying its property values, which doesn’t affect the other rate-map instances that already exist.


// An object that you use to configure new rasterization rate maps.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RasterizationRateMapDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RasterizationRateMapDescriptor */

// Creates a rate map descriptor with a given size and identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSize(screenSize objc.IObject /* cross-framework: MTLSize */) IRasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](objc.ID(rc.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:"), screenSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RasterizationRateMapDescriptorWithScreenSize) */


// Creates a rate map descriptor with a single rate layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:layer:
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayer(screenSize objc.IObject /* cross-framework: MTLSize */, layer IMTLRasterizationRateLayerDescriptor) IRasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](objc.ID(rc.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layer:"), screenSize, layer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RasterizationRateMapDescriptorWithScreenSizeLayer) */


// Creates a rate map descriptor with a set of layer descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:
func (rc _RasterizationRateMapDescriptorClass) RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers(screenSize objc.IObject /* cross-framework: MTLSize */, layerCount uint, layers objectivec.IObject) IRasterizationRateMapDescriptor {
	rv := objc.Send[RasterizationRateMapDescriptor](objc.ID(rc.class), objc.Sel("rasterizationRateMapDescriptorWithScreenSize:layerCount:layers:"), screenSize, layerCount, layers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RasterizationRateMapDescriptorWithScreenSizeLayerCountLayers) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RasterizationRateMapDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RasterizationRateMapDescriptor */

// Returns the layer description for a layer in the rate map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layer(at:)
func (r_ RasterizationRateMapDescriptor) LayerAtIndex(layerIndex uint) IRasterizationRateLayerDescriptor {
	rv := objc.Send[RasterizationRateLayerDescriptor](r_.ID, objc.Sel("layerAtIndex:"), layerIndex)
	return rv
}/* debug [instance_methods/method]: LayerAtIndex */


// Sets a configuration for a layer rate map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/setLayer(_:at:)
func (r_ RasterizationRateMapDescriptor) SetLayerAtIndex(layer IMTLRasterizationRateLayerDescriptor, layerIndex uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLayer:atIndex:"), layer, layerIndex)
}/* debug [instance_methods/method]: SetLayerAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RasterizationRateMapDescriptor */

// A string used to identify the rate map you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/label
func (r_ RasterizationRateMapDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string used to identify the rate map you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/label
func (r_ RasterizationRateMapDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The number of layers in the rate map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layerCount
func (r_ RasterizationRateMapDescriptor) LayerCount() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("layerCount"))
	return rv
}/* debug [instance_properties/getter]: layerCount */


// The rasterization rates for one or more layers in the rate map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/layers
func (r_ RasterizationRateMapDescriptor) Layers() IMTLRasterizationRateLayerArray {
	rv := objc.Send[RasterizationRateLayerArray](r_.ID, objc.Sel("layers"))
	return rv
}/* debug [instance_properties/getter]: layers */


// The size of the viewport coordinate system, in logical pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/screenSize
func (r_ RasterizationRateMapDescriptor) ScreenSize() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("screenSize"))
	return rv
}/* debug [instance_properties/getter]: screenSize */


// The size of the viewport coordinate system, in logical pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateMapDescriptor/screenSize
func (r_ RasterizationRateMapDescriptor) SetScreenSize(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScreenSize:"), value)
}/* debug [instance_properties/setter]: screenSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRasterizationRateMapDescriptor */



