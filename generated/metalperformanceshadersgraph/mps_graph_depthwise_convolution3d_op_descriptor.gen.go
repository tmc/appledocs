// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPSGraphDepthwiseConvolution3DOpDescriptor */


/* debug [class_header]: Header for MPSGraphDepthwiseConvolution3DOpDescriptor */
// The class instance for the [GraphDepthwiseConvolution3DOpDescriptor] class.
var (
	GraphDepthwiseConvolution3DOpDescriptorClass     _GraphDepthwiseConvolution3DOpDescriptorClass
	GraphDepthwiseConvolution3DOpDescriptorClassOnce sync.Once
)

func getGraphDepthwiseConvolution3DOpDescriptorClass() _GraphDepthwiseConvolution3DOpDescriptorClass {
	GraphDepthwiseConvolution3DOpDescriptorClassOnce.Do(func() {
		GraphDepthwiseConvolution3DOpDescriptorClass = _GraphDepthwiseConvolution3DOpDescriptorClass{objc.GetClass("MPSGraphDepthwiseConvolution3DOpDescriptor")}
	})
	return GraphDepthwiseConvolution3DOpDescriptorClass
}

type _GraphDepthwiseConvolution3DOpDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphDepthwiseConvolution3DOpDescriptor */
// An interface definition for the [GraphDepthwiseConvolution3DOpDescriptor] class.
type IGraphDepthwiseConvolution3DOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphDepthwiseConvolution3DOpDescriptor */
	// properties:
	ChannelDimensionIndex() int
	SetChannelDimensionIndex(value int)
	DilationRates() []foundation.Number
	SetDilationRates(value []foundation.Number)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingValues() []foundation.Number
	SetPaddingValues(value []foundation.Number)
	Strides() []foundation.Number
	SetStrides(value []foundation.Number)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphDepthwiseConvolution3DOpDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphDepthwiseConvolution3DOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) Alloc() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) New() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Init() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Autorelease() GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphDepthwiseConvolution3DOpDescriptor creates a new GraphDepthwiseConvolution3DOpDescriptor instance.
func NewGraphDepthwiseConvolution3DOpDescriptor() GraphDepthwiseConvolution3DOpDescriptor {
	return getGraphDepthwiseConvolution3DOpDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphDepthwiseConvolution3DOpDescriptor */
// The class that defines the parameters for a 3D-depthwise convolution operation.
//
// A defines constant parameters for 3D depthwise convolutions. Use this class with , and methods.


// The class that defines the parameters for a 3D-depthwise convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor
type GraphDepthwiseConvolution3DOpDescriptor struct {
	GraphObject
}

// GraphDepthwiseConvolution3DOpDescriptorFrom constructs a [GraphDepthwiseConvolution3DOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a 3D-depthwise convolution operation.
func GraphDepthwiseConvolution3DOpDescriptorFrom(ptr unsafe.Pointer) GraphDepthwiseConvolution3DOpDescriptor {
	return GraphDepthwiseConvolution3DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphDepthwiseConvolution3DOpDescriptor */

// Creates a 3D depthwise convolution descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(paddingStyle:)
func NewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle(paddingStyle GraphPaddingStyle) GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(getGraphDepthwiseConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphDepthwiseConvolution3DOpDescriptorWithPaddingStyle */


// Creates a 3D depthwise convolution descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(strides:dilationRates:paddingValues:paddingStyle:)
func NewGraphDepthwiseConvolution3DOpDescriptorWithStridesDilationRatesPaddingValuesPaddingStyle(strides []foundation.Number, dilationRates []foundation.Number, paddingValues []foundation.Number, paddingStyle GraphPaddingStyle) GraphDepthwiseConvolution3DOpDescriptor {
	rv := objc.Send[GraphDepthwiseConvolution3DOpDescriptor](objc.ID(getGraphDepthwiseConvolution3DOpDescriptorClass().class), objc.Sel("descriptorWithStrides:dilationRates:paddingValues:paddingStyle:"), strides, dilationRates, paddingValues, paddingStyle)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphDepthwiseConvolution3DOpDescriptorWithStridesDilationRatesPaddingValuesPaddingStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphDepthwiseConvolution3DOpDescriptor */

// Creates a 3D depthwise convolution descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(paddingStyle:)
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) DescriptorWithPaddingStyle(paddingStyle GraphPaddingStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithPaddingStyle) */


// Creates a 3D depthwise convolution descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/init(strides:dilationRates:paddingValues:paddingStyle:)
func (gc _GraphDepthwiseConvolution3DOpDescriptorClass) DescriptorWithStridesDilationRatesPaddingValuesPaddingStyle(strides []foundation.Number, dilationRates []foundation.Number, paddingValues []foundation.Number, paddingStyle GraphPaddingStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithStrides:dilationRates:paddingValues:paddingStyle:"), strides, dilationRates, paddingValues, paddingStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithStridesDilationRatesPaddingValuesPaddingStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphDepthwiseConvolution3DOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphDepthwiseConvolution3DOpDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphDepthwiseConvolution3DOpDescriptor */

// The axis that contains the channels in the input and the weights, within the 4D tile of the last dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/channelDimensionIndex
func (g_ GraphDepthwiseConvolution3DOpDescriptor) ChannelDimensionIndex() int {
	rv := objc.Send[int](g_.ID, objc.Sel("channelDimensionIndex"))
	return rv
}/* debug [instance_properties/getter]: channelDimensionIndex */


// The axis that contains the channels in the input and the weights, within the 4D tile of the last dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/channelDimensionIndex
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetChannelDimensionIndex(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setChannelDimensionIndex:"), value)
}/* debug [instance_properties/setter]: channelDimensionIndex */


// The dilation rates for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/dilationRates
func (g_ GraphDepthwiseConvolution3DOpDescriptor) DilationRates() []foundation.Number {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("dilationRates"))
	return rv
}/* debug [instance_properties/getter]: dilationRates */


// The dilation rates for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/dilationRates
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetDilationRates(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRates:"), nsArray)
}/* debug [instance_properties/setter]: dilationRates */


// The padding style for the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/paddingStyle
func (g_ GraphDepthwiseConvolution3DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}/* debug [instance_properties/getter]: paddingStyle */


// The padding style for the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/paddingStyle
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}/* debug [instance_properties/setter]: paddingStyle */


// The padding values for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/paddingValues
func (g_ GraphDepthwiseConvolution3DOpDescriptor) PaddingValues() []foundation.Number {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("paddingValues"))
	return rv
}/* debug [instance_properties/getter]: paddingValues */


// The padding values for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/paddingValues
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetPaddingValues(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingValues:"), nsArray)
}/* debug [instance_properties/setter]: paddingValues */


// The strides for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/strides
func (g_ GraphDepthwiseConvolution3DOpDescriptor) Strides() []foundation.Number {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// The strides for spatial dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphDepthwiseConvolution3DOpDescriptor/strides
func (g_ GraphDepthwiseConvolution3DOpDescriptor) SetStrides(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), nsArray)
}/* debug [instance_properties/setter]: strides */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphDepthwiseConvolution3DOpDescriptor */


