// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphPooling4DOpDescriptor */


/* debug [class_header]: Header for MPSGraphPooling4DOpDescriptor */
// The class instance for the [GraphPooling4DOpDescriptor] class.
var (
	GraphPooling4DOpDescriptorClass     _GraphPooling4DOpDescriptorClass
	GraphPooling4DOpDescriptorClassOnce sync.Once
)

func getGraphPooling4DOpDescriptorClass() _GraphPooling4DOpDescriptorClass {
	GraphPooling4DOpDescriptorClassOnce.Do(func() {
		GraphPooling4DOpDescriptorClass = _GraphPooling4DOpDescriptorClass{objc.GetClass("MPSGraphPooling4DOpDescriptor")}
	})
	return GraphPooling4DOpDescriptorClass
}

type _GraphPooling4DOpDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphPooling4DOpDescriptor */
// An interface definition for the [GraphPooling4DOpDescriptor] class.
type IGraphPooling4DOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphPooling4DOpDescriptor */
	// properties:
	CeilMode() bool
	SetCeilMode(value bool)
	DilationRates() []objc.IObject /* cross-framework: Number */
	SetDilationRates(value []objc.IObject /* cross-framework: Number */)
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	KernelSizes() []objc.IObject /* cross-framework: Number */
	SetKernelSizes(value []objc.IObject /* cross-framework: Number */)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingValues() []objc.IObject /* cross-framework: Number */
	SetPaddingValues(value []objc.IObject /* cross-framework: Number */)
	ReturnIndicesDataType() DataType /* not a class type */
	SetReturnIndicesDataType(value DataType /* not a class type */)
	ReturnIndicesMode() GraphPoolingReturnIndicesMode
	SetReturnIndicesMode(value GraphPoolingReturnIndicesMode)
	Strides() []objc.IObject /* cross-framework: Number */
	SetStrides(value []objc.IObject /* cross-framework: Number */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphPooling4DOpDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphPooling4DOpDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphPooling4DOpDescriptorClass) Alloc() GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphPooling4DOpDescriptorClass) New() GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphPooling4DOpDescriptor) Init() GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphPooling4DOpDescriptor) Autorelease() GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphPooling4DOpDescriptor creates a new GraphPooling4DOpDescriptor instance.
func NewGraphPooling4DOpDescriptor() GraphPooling4DOpDescriptor {
	return getGraphPooling4DOpDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphPooling4DOpDescriptor */
// The class that defines the parameters for a 4D pooling operation.
//
// Use this descriptor with the following methods:


// The class that defines the parameters for a 4D pooling operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor
type GraphPooling4DOpDescriptor struct {
	GraphObject
}

// GraphPooling4DOpDescriptorFrom constructs a [GraphPooling4DOpDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a 4D pooling operation.
func GraphPooling4DOpDescriptorFrom(ptr unsafe.Pointer) GraphPooling4DOpDescriptor {
	return GraphPooling4DOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphPooling4DOpDescriptor */

// Creates a 4D pooling descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/init(kernelSizes:paddingStyle:)
func NewGraphPooling4DOpDescriptorWithKernelSizesPaddingStyle(kernelSizes []objc.IObject /* cross-framework: Number */, paddingStyle GraphPaddingStyle) GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](objc.ID(getGraphPooling4DOpDescriptorClass().class), objc.Sel("descriptorWithKernelSizes:paddingStyle:"), kernelSizes, paddingStyle)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphPooling4DOpDescriptorWithKernelSizesPaddingStyle */


// Creates a 4D pooling descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/init(kernelSizes:strides:dilationRates:paddingValues:paddingStyle:)
func NewGraphPooling4DOpDescriptorWithKernelSizesStridesDilationRatesPaddingValuesPaddingStyle(kernelSizes []objc.IObject /* cross-framework: Number */, strides []objc.IObject /* cross-framework: Number */, dilationRates []objc.IObject /* cross-framework: Number */, paddingValues []objc.IObject /* cross-framework: Number */, paddingStyle GraphPaddingStyle) GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](objc.ID(getGraphPooling4DOpDescriptorClass().class), objc.Sel("descriptorWithKernelSizes:strides:dilationRates:paddingValues:paddingStyle:"), kernelSizes, strides, dilationRates, paddingValues, paddingStyle)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphPooling4DOpDescriptorWithKernelSizesStridesDilationRatesPaddingValuesPaddingStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphPooling4DOpDescriptor */

// Creates a 4D pooling descriptor with default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/init(kernelSizes:paddingStyle:)
func (gc _GraphPooling4DOpDescriptorClass) DescriptorWithKernelSizesPaddingStyle(kernelSizes []objc.IObject /* cross-framework: Number */, paddingStyle GraphPaddingStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelSizes:paddingStyle:"), kernelSizes, paddingStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelSizesPaddingStyle) */


// Creates a 4D pooling descriptor with given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/init(kernelSizes:strides:dilationRates:paddingValues:paddingStyle:)
func (gc _GraphPooling4DOpDescriptorClass) DescriptorWithKernelSizesStridesDilationRatesPaddingValuesPaddingStyle(kernelSizes []objc.IObject /* cross-framework: Number */, strides []objc.IObject /* cross-framework: Number */, dilationRates []objc.IObject /* cross-framework: Number */, paddingValues []objc.IObject /* cross-framework: Number */, paddingStyle GraphPaddingStyle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithKernelSizes:strides:dilationRates:paddingValues:paddingStyle:"), kernelSizes, strides, dilationRates, paddingValues, paddingStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithKernelSizesStridesDilationRatesPaddingValuesPaddingStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphPooling4DOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphPooling4DOpDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphPooling4DOpDescriptor */

// Affects how MPSGraph computes the output size: if set to then output size is computed by rounding up instead of down when dividing input size by stride.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/ceilMode
func (g_ GraphPooling4DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("ceilMode"))
	return rv
}/* debug [instance_properties/getter]: ceilMode */


// Affects how MPSGraph computes the output size: if set to then output size is computed by rounding up instead of down when dividing input size by stride.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/ceilMode
func (g_ GraphPooling4DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCeilMode:"), value)
}/* debug [instance_properties/setter]: ceilMode */


// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/dilationRates
func (g_ GraphPooling4DOpDescriptor) DilationRates() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("dilationRates"))
	return rv
}/* debug [instance_properties/getter]: dilationRates */


// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/dilationRates
func (g_ GraphPooling4DOpDescriptor) SetDilationRates(value []objc.IObject /* cross-framework: Number */) {
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


// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/includeZeroPadToAverage
func (g_ GraphPooling4DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}/* debug [instance_properties/getter]: includeZeroPadToAverage */


// Defines a mode for average pooling, where samples outside the input tensor count as zeroes in the average computation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/includeZeroPadToAverage
func (g_ GraphPooling4DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}/* debug [instance_properties/setter]: includeZeroPadToAverage */


// Defines the pooling window size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/kernelSizes
func (g_ GraphPooling4DOpDescriptor) KernelSizes() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("kernelSizes"))
	return rv
}/* debug [instance_properties/getter]: kernelSizes */


// Defines the pooling window size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/kernelSizes
func (g_ GraphPooling4DOpDescriptor) SetKernelSizes(value []objc.IObject /* cross-framework: Number */) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelSizes:"), nsArray)
}/* debug [instance_properties/setter]: kernelSizes */


// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingStyle
func (g_ GraphPooling4DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}/* debug [instance_properties/getter]: paddingStyle */


// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingStyle
func (g_ GraphPooling4DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}/* debug [instance_properties/setter]: paddingStyle */


// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingValues
func (g_ GraphPooling4DOpDescriptor) PaddingValues() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("paddingValues"))
	return rv
}/* debug [instance_properties/getter]: paddingValues */


// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingValues
func (g_ GraphPooling4DOpDescriptor) SetPaddingValues(value []objc.IObject /* cross-framework: Number */) {
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


// Defines the data type for returned indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/returnIndicesDataType
func (g_ GraphPooling4DOpDescriptor) ReturnIndicesDataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("returnIndicesDataType"))
	return rv
}/* debug [instance_properties/getter]: returnIndicesDataType */


// Defines the data type for returned indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/returnIndicesDataType
func (g_ GraphPooling4DOpDescriptor) SetReturnIndicesDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesDataType:"), value)
}/* debug [instance_properties/setter]: returnIndicesDataType */


// Defines the mode for returned indices of maximum values within each pooling window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/returnIndicesMode
func (g_ GraphPooling4DOpDescriptor) ReturnIndicesMode() GraphPoolingReturnIndicesMode {
	rv := objc.Send[GraphPoolingReturnIndicesMode](g_.ID, objc.Sel("returnIndicesMode"))
	return rv
}/* debug [instance_properties/getter]: returnIndicesMode */


// Defines the mode for returned indices of maximum values within each pooling window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/returnIndicesMode
func (g_ GraphPooling4DOpDescriptor) SetReturnIndicesMode(value GraphPoolingReturnIndicesMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesMode:"), value)
}/* debug [instance_properties/setter]: returnIndicesMode */


// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/strides
func (g_ GraphPooling4DOpDescriptor) Strides() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/strides
func (g_ GraphPooling4DOpDescriptor) SetStrides(value []objc.IObject /* cross-framework: Number */) {
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


/* debug [class.gen.go]: End class MPSGraphPooling4DOpDescriptor */


