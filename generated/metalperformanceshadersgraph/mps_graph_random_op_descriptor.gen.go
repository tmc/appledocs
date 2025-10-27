// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GraphRandomOpDescriptor] class.
var (
	GraphRandomOpDescriptorClass     _GraphRandomOpDescriptorClass
	GraphRandomOpDescriptorClassOnce sync.Once
)

func getGraphRandomOpDescriptorClass() _GraphRandomOpDescriptorClass {
	GraphRandomOpDescriptorClassOnce.Do(func() {
		GraphRandomOpDescriptorClass = _GraphRandomOpDescriptorClass{objc.GetClass("MPSGraphRandomOpDescriptor")}
	})
	return GraphRandomOpDescriptorClass
}

type _GraphRandomOpDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [GraphRandomOpDescriptor] class.
type IGraphRandomOpDescriptor interface {
	IGraphObject
	

	// properties:
	DataType() DataType /* not a class type */
	SetDataType(value DataType /* not a class type */)
	Distribution() GraphRandomDistribution
	SetDistribution(value GraphRandomDistribution)
	Max() float32
	SetMax(value float32)
	MaxInteger() int
	SetMaxInteger(value int)
	Mean() float32
	SetMean(value float32)
	Min() float32
	SetMin(value float32)
	MinInteger() int
	SetMinInteger(value int)
	SamplingMethod() GraphRandomNormalSamplingMethod
	SetSamplingMethod(value GraphRandomNormalSamplingMethod)
	StandardDeviation() float32
	SetStandardDeviation(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GraphRandomOpDescriptorClass) Alloc() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphRandomOpDescriptorClass) New() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphRandomOpDescriptor) Init() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphRandomOpDescriptor) Autorelease() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphRandomOpDescriptor creates a new GraphRandomOpDescriptor instance.
func NewGraphRandomOpDescriptor() GraphRandomOpDescriptor {
	return getGraphRandomOpDescriptorClass().New()
}





// A class that describes the random operation.


// A class that describes the random operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor
type GraphRandomOpDescriptor struct {
	GraphObject
}

// GraphRandomOpDescriptorFrom constructs a [GraphRandomOpDescriptor] from an unsafe.Pointer.
//
// A class that describes the random operation.
func GraphRandomOpDescriptorFrom(ptr unsafe.Pointer) GraphRandomOpDescriptor {
	return GraphRandomOpDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}






// Class method to initialize a distribution descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/init(distribution:dataType:)
func NewGraphRandomOpDescriptorWithDistributionDataType(distribution GraphRandomDistribution, dataType DataType /* not a class type */) GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(getGraphRandomOpDescriptorClass().class), objc.Sel("descriptorWithDistribution:dataType:"), distribution, dataType)
	return rv
}







// Class method to initialize a distribution descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/init(distribution:dataType:)
func (gc _GraphRandomOpDescriptorClass) DescriptorWithDistributionDataType(distribution GraphRandomDistribution, dataType DataType /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithDistribution:dataType:"), distribution, dataType)
	return rv
}

















// The data type of the generated result values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) DataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](g_.ID, objc.Sel("dataType"))
	return rv
}


// The data type of the generated result values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) SetDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}


// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) Distribution() GraphRandomDistribution {
	rv := objc.Send[GraphRandomDistribution](g_.ID, objc.Sel("distribution"))
	return rv
}


// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) SetDistribution(value GraphRandomDistribution) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDistribution:"), value)
}


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/max
func (g_ GraphRandomOpDescriptor) Max() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("max"))
	return rv
}


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/max
func (g_ GraphRandomOpDescriptor) SetMax(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMax:"), value)
}


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/maxInteger
func (g_ GraphRandomOpDescriptor) MaxInteger() int {
	rv := objc.Send[int](g_.ID, objc.Sel("maxInteger"))
	return rv
}


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/maxInteger
func (g_ GraphRandomOpDescriptor) SetMaxInteger(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMaxInteger:"), value)
}


// The mean of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/mean
func (g_ GraphRandomOpDescriptor) Mean() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("mean"))
	return rv
}


// The mean of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/mean
func (g_ GraphRandomOpDescriptor) SetMean(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMean:"), value)
}


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/min
func (g_ GraphRandomOpDescriptor) Min() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("min"))
	return rv
}


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/min
func (g_ GraphRandomOpDescriptor) SetMin(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMin:"), value)
}


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/minInteger
func (g_ GraphRandomOpDescriptor) MinInteger() int {
	rv := objc.Send[int](g_.ID, objc.Sel("minInteger"))
	return rv
}


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/minInteger
func (g_ GraphRandomOpDescriptor) SetMinInteger(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMinInteger:"), value)
}


// The sampling method of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/samplingMethod
func (g_ GraphRandomOpDescriptor) SamplingMethod() GraphRandomNormalSamplingMethod {
	rv := objc.Send[GraphRandomNormalSamplingMethod](g_.ID, objc.Sel("samplingMethod"))
	return rv
}


// The sampling method of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/samplingMethod
func (g_ GraphRandomOpDescriptor) SetSamplingMethod(value GraphRandomNormalSamplingMethod) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSamplingMethod:"), value)
}


// The standard deviation of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/standardDeviation
func (g_ GraphRandomOpDescriptor) StandardDeviation() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("standardDeviation"))
	return rv
}


// The standard deviation of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/standardDeviation
func (g_ GraphRandomOpDescriptor) SetStandardDeviation(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStandardDeviation:"), value)
}







