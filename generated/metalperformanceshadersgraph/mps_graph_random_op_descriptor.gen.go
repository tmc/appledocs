// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphRandomOpDescriptor */


/* debug [class_header]: Header for MPSGraphRandomOpDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphRandomOpDescriptor */
// An interface definition for the [GraphRandomOpDescriptor] class.
type IGraphRandomOpDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphRandomOpDescriptor */
	// properties:
	DataType() objc.IObject /* cross-framework: DataType */
	SetDataType(value objc.IObject /* cross-framework: DataType */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphRandomOpDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphRandomOpDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphRandomOpDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphRandomOpDescriptor */

// Class method to initialize a distribution descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/init(distribution:dataType:)
func NewGraphRandomOpDescriptorWithDistributionDataType(distribution GraphRandomDistribution, dataType objc.IObject /* cross-framework: DataType */) GraphRandomOpDescriptor {
	rv := objc.Send[GraphRandomOpDescriptor](objc.ID(getGraphRandomOpDescriptorClass().class), objc.Sel("descriptorWithDistribution:dataType:"), distribution, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewGraphRandomOpDescriptorWithDistributionDataType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphRandomOpDescriptor */

// Class method to initialize a distribution descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/init(distribution:dataType:)
func (gc _GraphRandomOpDescriptorClass) DescriptorWithDistributionDataType(distribution GraphRandomDistribution, dataType objc.IObject /* cross-framework: DataType */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptorWithDistribution:dataType:"), distribution, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithDistributionDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphRandomOpDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphRandomOpDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphRandomOpDescriptor */

// The data type of the generated result values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) DataType() objc.IObject /* cross-framework: DataType */ {
	rv := objc.Send[metalperformanceshaders.DataType](g_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the generated result values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/dataType
func (g_ GraphRandomOpDescriptor) SetDataType(value objc.IObject /* cross-framework: DataType */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) Distribution() GraphRandomDistribution {
	rv := objc.Send[GraphRandomDistribution](g_.ID, objc.Sel("distribution"))
	return rv
}/* debug [instance_properties/getter]: distribution */


// The type of distribution to draw samples from. See MPSGraphRandomDistribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/distribution
func (g_ GraphRandomOpDescriptor) SetDistribution(value GraphRandomDistribution) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDistribution:"), value)
}/* debug [instance_properties/setter]: distribution */


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/max
func (g_ GraphRandomOpDescriptor) Max() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("max"))
	return rv
}/* debug [instance_properties/getter]: max */


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/max
func (g_ GraphRandomOpDescriptor) SetMax(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMax:"), value)
}/* debug [instance_properties/setter]: max */


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/maxInteger
func (g_ GraphRandomOpDescriptor) MaxInteger() int {
	rv := objc.Send[int](g_.ID, objc.Sel("maxInteger"))
	return rv
}/* debug [instance_properties/getter]: maxInteger */


// The upper range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/maxInteger
func (g_ GraphRandomOpDescriptor) SetMaxInteger(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMaxInteger:"), value)
}/* debug [instance_properties/setter]: maxInteger */


// The mean of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/mean
func (g_ GraphRandomOpDescriptor) Mean() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("mean"))
	return rv
}/* debug [instance_properties/getter]: mean */


// The mean of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/mean
func (g_ GraphRandomOpDescriptor) SetMean(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMean:"), value)
}/* debug [instance_properties/setter]: mean */


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/min
func (g_ GraphRandomOpDescriptor) Min() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("min"))
	return rv
}/* debug [instance_properties/getter]: min */


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/min
func (g_ GraphRandomOpDescriptor) SetMin(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMin:"), value)
}/* debug [instance_properties/setter]: min */


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/minInteger
func (g_ GraphRandomOpDescriptor) MinInteger() int {
	rv := objc.Send[int](g_.ID, objc.Sel("minInteger"))
	return rv
}/* debug [instance_properties/getter]: minInteger */


// The lower range of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/minInteger
func (g_ GraphRandomOpDescriptor) SetMinInteger(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMinInteger:"), value)
}/* debug [instance_properties/setter]: minInteger */


// The sampling method of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/samplingMethod
func (g_ GraphRandomOpDescriptor) SamplingMethod() GraphRandomNormalSamplingMethod {
	rv := objc.Send[GraphRandomNormalSamplingMethod](g_.ID, objc.Sel("samplingMethod"))
	return rv
}/* debug [instance_properties/getter]: samplingMethod */


// The sampling method of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/samplingMethod
func (g_ GraphRandomOpDescriptor) SetSamplingMethod(value GraphRandomNormalSamplingMethod) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSamplingMethod:"), value)
}/* debug [instance_properties/setter]: samplingMethod */


// The standard deviation of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/standardDeviation
func (g_ GraphRandomOpDescriptor) StandardDeviation() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("standardDeviation"))
	return rv
}/* debug [instance_properties/getter]: standardDeviation */


// The standard deviation of the distribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphRandomOpDescriptor/standardDeviation
func (g_ GraphRandomOpDescriptor) SetStandardDeviation(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStandardDeviation:"), value)
}/* debug [instance_properties/setter]: standardDeviation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphRandomOpDescriptor */


