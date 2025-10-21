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
}

// A class that describes the random operation.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphRandomOpDescriptorClass) Alloc() GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The data type of the generated result values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// The data type of the generated result values.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}

// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) Distribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("distribution"))
	return rv
}


// SetDistribution sets the value of the distribution property.
// The type of distribution to draw samples from. See MPSGraphRandomDistribution.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) SetDistribution(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDistribution:"), value)
}

// The upper range of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/max
func (g_ GraphRandomOpDescriptor) Max() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("max"))
	return rv
}


// SetMax sets the value of the max property.
// The upper range of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/max
func (g_ GraphRandomOpDescriptor) SetMax(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMax:"), value)
}

// The upper range of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/maxinteger
func (g_ GraphRandomOpDescriptor) MaxInteger() int {
	rv := objc.Send[int](g_.ID, objc.Sel("maxInteger"))
	return rv
}


// SetMaxInteger sets the value of the maxInteger property.
// The upper range of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/maxinteger
func (g_ GraphRandomOpDescriptor) SetMaxInteger(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMaxInteger:"), value)
}

// The mean of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/mean
func (g_ GraphRandomOpDescriptor) Mean() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mean"))
	return rv
}


// SetMean sets the value of the mean property.
// The mean of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/mean
func (g_ GraphRandomOpDescriptor) SetMean(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMean:"), value)
}

// The lower range of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/min
func (g_ GraphRandomOpDescriptor) Min() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("min"))
	return rv
}


// SetMin sets the value of the min property.
// The lower range of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/min
func (g_ GraphRandomOpDescriptor) SetMin(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMin:"), value)
}

// The lower range of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/mininteger
func (g_ GraphRandomOpDescriptor) MinInteger() int {
	rv := objc.Send[int](g_.ID, objc.Sel("minInteger"))
	return rv
}


// SetMinInteger sets the value of the minInteger property.
// The lower range of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/mininteger
func (g_ GraphRandomOpDescriptor) SetMinInteger(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMinInteger:"), value)
}

// The sampling method of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/samplingmethod
func (g_ GraphRandomOpDescriptor) SamplingMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("samplingMethod"))
	return rv
}


// SetSamplingMethod sets the value of the samplingMethod property.
// The sampling method of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/samplingmethod
func (g_ GraphRandomOpDescriptor) SetSamplingMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSamplingMethod:"), value)
}

// The standard deviation of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/standarddeviation
func (g_ GraphRandomOpDescriptor) StandardDeviation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("standardDeviation"))
	return rv
}


// SetStandardDeviation sets the value of the standardDeviation property.
// The standard deviation of the distribution.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphrandomopdescriptor/standarddeviation
func (g_ GraphRandomOpDescriptor) SetStandardDeviation(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStandardDeviation:"), value)
}



