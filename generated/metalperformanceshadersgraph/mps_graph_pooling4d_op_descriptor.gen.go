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
	// properties:
	CeilMode() bool
	SetCeilMode(value bool)
	DilationRates() objc.IObject /* cross-framework: NSNumber */
	SetDilationRates(value objc.IObject /* cross-framework: NSNumber */)
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	KernelSizes() objc.IObject /* cross-framework: NSNumber */
	SetKernelSizes(value objc.IObject /* cross-framework: NSNumber */)
	PaddingStyle() GraphPaddingStyle
	SetPaddingStyle(value GraphPaddingStyle)
	PaddingValues() objc.IObject /* cross-framework: NSNumber */
	SetPaddingValues(value objc.IObject /* cross-framework: NSNumber */)
	ReturnIndicesDataType() DataType /* not a class type */
	SetReturnIndicesDataType(value DataType /* not a class type */)
	ReturnIndicesMode() GraphPoolingReturnIndicesMode
	SetReturnIndicesMode(value GraphPoolingReturnIndicesMode)
	Strides() objc.IObject /* cross-framework: NSNumber */
	SetStrides(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

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



// Affects how MPSGraph computes the output size: if set to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/ceilmode
func (g_ GraphPooling4DOpDescriptor) CeilMode() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("ceilMode"))
	return rv
}


// Affects how MPSGraph computes the output size: if set to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/ceilmode
func (g_ GraphPooling4DOpDescriptor) SetCeilMode(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCeilMode:"), value)
}


// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/dilationrates
func (g_ GraphPooling4DOpDescriptor) DilationRates() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("dilationRates"))
	return rv
}


// Defines dilation rates for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/dilationrates
func (g_ GraphPooling4DOpDescriptor) SetDilationRates(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDilationRates:"), value)
}


// Defines a mode for average pooling, where samples outside the input tensor count as
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/includezeropadtoaverage
func (g_ GraphPooling4DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("includeZeroPadToAverage"))
	return rv
}


// Defines a mode for average pooling, where samples outside the input tensor count as
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/includezeropadtoaverage
func (g_ GraphPooling4DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIncludeZeroPadToAverage:"), value)
}


// Defines the pooling window size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/kernelsizes
func (g_ GraphPooling4DOpDescriptor) KernelSizes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("kernelSizes"))
	return rv
}


// Defines the pooling window size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/kernelsizes
func (g_ GraphPooling4DOpDescriptor) SetKernelSizes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKernelSizes:"), value)
}


// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/paddingstyle
func (g_ GraphPooling4DOpDescriptor) PaddingStyle() GraphPaddingStyle {
	rv := objc.Send[GraphPaddingStyle](g_.ID, objc.Sel("paddingStyle"))
	return rv
}


// Defines what kind of padding graph applies to the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/paddingstyle
func (g_ GraphPooling4DOpDescriptor) SetPaddingStyle(value GraphPaddingStyle) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingStyle:"), value)
}


// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/paddingvalues
func (g_ GraphPooling4DOpDescriptor) PaddingValues() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("paddingValues"))
	return rv
}


// Defines padding values for spatial dimensions which must be eight numbers, two for each spatial dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/paddingvalues
func (g_ GraphPooling4DOpDescriptor) SetPaddingValues(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingValues:"), value)
}


// Defines the data type for returned indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesdatatype
func (g_ GraphPooling4DOpDescriptor) ReturnIndicesDataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("returnIndicesDataType"))
	return rv
}


// Defines the data type for returned indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesdatatype
func (g_ GraphPooling4DOpDescriptor) SetReturnIndicesDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesDataType:"), value)
}


// Defines the mode for returned indices of maximum values within each pooling window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesmode
func (g_ GraphPooling4DOpDescriptor) ReturnIndicesMode() GraphPoolingReturnIndicesMode {
	rv := objc.Send[GraphPoolingReturnIndicesMode](g_.ID, objc.Sel("returnIndicesMode"))
	return rv
}


// Defines the mode for returned indices of maximum values within each pooling window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/returnindicesmode
func (g_ GraphPooling4DOpDescriptor) SetReturnIndicesMode(value GraphPoolingReturnIndicesMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReturnIndicesMode:"), value)
}


// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/strides
func (g_ GraphPooling4DOpDescriptor) Strides() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](g_.ID, objc.Sel("strides"))
	return rv
}


// Defines strides for spatial dimensions. Must be four numbers, one for each spatial dimension, fastest running index last.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/strides
func (g_ GraphPooling4DOpDescriptor) SetStrides(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStrides:"), value)
}



