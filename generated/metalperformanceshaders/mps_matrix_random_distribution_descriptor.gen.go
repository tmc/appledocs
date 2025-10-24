// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MatrixRandomDistributionDescriptor] class.
var (
	MatrixRandomDistributionDescriptorClass     _MatrixRandomDistributionDescriptorClass
	MatrixRandomDistributionDescriptorClassOnce sync.Once
)

func getMatrixRandomDistributionDescriptorClass() _MatrixRandomDistributionDescriptorClass {
	MatrixRandomDistributionDescriptorClassOnce.Do(func() {
		MatrixRandomDistributionDescriptorClass = _MatrixRandomDistributionDescriptorClass{objc.GetClass("MPSMatrixRandomDistributionDescriptor")}
	})
	return MatrixRandomDistributionDescriptorClass
}

type _MatrixRandomDistributionDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MatrixRandomDistributionDescriptor] class.
type IMatrixRandomDistributionDescriptor interface {
	objectivec.IObject
	// properties:
	Minimum() float32
	SetMinimum(value float32)
	StandardDeviation() float32
	SetStandardDeviation(value float32)
	DistributionType() MatrixRandomDistribution /* not a class type */
	SetDistributionType(value MatrixRandomDistribution /* not a class type */)
	Maximum() float32
	SetMaximum(value float32)
	Mean() float32
	SetMean(value float32)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor
type MatrixRandomDistributionDescriptor struct {
	objectivec.Object
}

// MatrixRandomDistributionDescriptorFrom constructs a [MatrixRandomDistributionDescriptor] from an unsafe.Pointer.
func MatrixRandomDistributionDescriptorFrom(ptr unsafe.Pointer) MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomDistributionDescriptorClass) Alloc() MatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatrixRandomDistributionDescriptorClass) New() MatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixRandomDistributionDescriptor) Init() MatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixRandomDistributionDescriptor) Autorelease() MatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixRandomDistributionDescriptor creates a new MatrixRandomDistributionDescriptor instance.
func NewMatrixRandomDistributionDescriptor() MatrixRandomDistributionDescriptor {
	return getMatrixRandomDistributionDescriptorClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/normalDistributionDescriptor(withMean:standardDeviation:minimum:maximum:)
func (mc _MatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviationMinimumMaximum(mean float32, standardDeviation float32, minimum float32, maximum float32) IMatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:minimum:maximum:"), mean, standardDeviation, minimum, maximum)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/minimum
func (m_ MatrixRandomDistributionDescriptor) Minimum() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("minimum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/minimum
func (m_ MatrixRandomDistributionDescriptor) SetMinimum(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/standardDeviation
func (m_ MatrixRandomDistributionDescriptor) StandardDeviation() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("standardDeviation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/standardDeviation
func (m_ MatrixRandomDistributionDescriptor) SetStandardDeviation(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStandardDeviation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/distributiontype
func (m_ MatrixRandomDistributionDescriptor) DistributionType() MatrixRandomDistribution /* not a class type */ {
	rv := objc.Send[MatrixRandomDistribution](m_.ID, objc.Sel("distributionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/distributiontype
func (m_ MatrixRandomDistributionDescriptor) SetDistributionType(value MatrixRandomDistribution /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDistributionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/maximum
func (m_ MatrixRandomDistributionDescriptor) Maximum() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("maximum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/maximum
func (m_ MatrixRandomDistributionDescriptor) SetMaximum(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/mean
func (m_ MatrixRandomDistributionDescriptor) Mean() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("mean"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/mean
func (m_ MatrixRandomDistributionDescriptor) SetMean(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMean:"), value)
}



