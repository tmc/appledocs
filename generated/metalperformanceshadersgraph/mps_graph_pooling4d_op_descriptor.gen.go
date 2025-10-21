// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [GraphPooling4DOpDescriptor] class.
type IGraphPooling4DOpDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a 4D pooling operation.
//
// Use this descriptor with the following methods:
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphPooling4DOpDescriptorClass) Alloc() GraphPooling4DOpDescriptor {
	rv := objc.Send[GraphPooling4DOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/paddingvalues
func (g_ GraphPooling4DOpDescriptor) PaddingValues() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("paddingValues"))
	return rv
}


// SetPaddingValues sets the value of the paddingValues property.
// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/paddingvalues
func (g_ GraphPooling4DOpDescriptor) SetPaddingValues(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingValues:"), value)
}

// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/strides
func (g_ GraphPooling4DOpDescriptor) Strides() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("strides"))
	return rv
}


// SetStrides sets the value of the strides property.
// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/strides
func (g_ GraphPooling4DOpDescriptor) SetStrides(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), value)
}

// Defines the pooling window size.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/kernelsizes
func (g_ GraphPooling4DOpDescriptor) KernelSizes() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("kernelSizes"))
	return rv
}


// SetKernelSizes sets the value of the kernelSizes property.
// Defines the pooling window size.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/kernelsizes
func (g_ GraphPooling4DOpDescriptor) SetKernelSizes(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelSizes:"), value)
}

// Defines a mode for average pooling, where samples outside the input tensor count as
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/includezeropadtoaverage
func (g_ GraphPooling4DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}


// SetIncludeZeroPadToAverage sets the value of the includeZeroPadToAverage property.
// Defines a mode for average pooling, where samples outside the input tensor count as

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/includezeropadtoaverage
func (g_ GraphPooling4DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}

// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/dilationrates
func (g_ GraphPooling4DOpDescriptor) DilationRates() foundation.Number {
	rv := objc.Send[foundation.Number](g_.ID, objc.Sel("dilationRates"))
	return rv
}


// SetDilationRates sets the value of the dilationRates property.
// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/dilationrates
func (g_ GraphPooling4DOpDescriptor) SetDilationRates(value foundation.Number) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRates:"), value)
}

// Defines the mode for returned indices of maximum values within each pooling window.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesmode
func (g_ GraphPooling4DOpDescriptor) ReturnIndicesMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("returnIndicesMode"))
	return rv
}


// SetReturnIndicesMode sets the value of the returnIndicesMode property.
// Defines the mode for returned indices of maximum values within each pooling window.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesmode
func (g_ GraphPooling4DOpDescriptor) SetReturnIndicesMode(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesMode:"), value)
}

// Defines the data type for returned indices.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesdatatype
func (g_ GraphPooling4DOpDescriptor) ReturnIndicesDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("returnIndicesDataType"))
	return rv
}


// SetReturnIndicesDataType sets the value of the returnIndicesDataType property.
// Defines the data type for returned indices.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesdatatype
func (g_ GraphPooling4DOpDescriptor) SetReturnIndicesDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesDataType:"), value)
}

// Affects how MPSGraph computes the output size: if set to then output size is computed by rounding up instead of down when dividing input size by stride.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/ceilMode
func (g_ GraphPooling4DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("ceilMode"))
	return rv
}


// SetCeilMode sets the value of the ceilMode property.
// Affects how MPSGraph computes the output size: if set to then output size is computed by rounding up instead of down when dividing input size by stride.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/ceilMode
func (g_ GraphPooling4DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCeilMode:"), value)
}

// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingStyle
func (g_ GraphPooling4DOpDescriptor) PaddingStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// SetPaddingStyle sets the value of the paddingStyle property.
// Defines what kind of padding graph applies to the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphPooling4DOpDescriptor/paddingStyle
func (g_ GraphPooling4DOpDescriptor) SetPaddingStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}



