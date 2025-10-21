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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/normalDistributionDescriptor(withMean:standardDeviation:minimum:maximum:)
func (mc _MatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviationMinimumMaximum(mean unsafe.Pointer, standardDeviation unsafe.Pointer, minimum unsafe.Pointer, maximum unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:minimum:maximum:"), mean, standardDeviation, minimum, maximum)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/mean
func (m_ MatrixRandomDistributionDescriptor) Mean() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mean"))
	return rv
}


// SetMean sets the value of the mean property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/mean
func (m_ MatrixRandomDistributionDescriptor) SetMean(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMean:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/distributiontype
func (m_ MatrixRandomDistributionDescriptor) DistributionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("distributionType"))
	return rv
}


// SetDistributionType sets the value of the distributionType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/distributiontype
func (m_ MatrixRandomDistributionDescriptor) SetDistributionType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDistributionType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/maximum
func (m_ MatrixRandomDistributionDescriptor) Maximum() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maximum"))
	return rv
}


// SetMaximum sets the value of the maximum property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/maximum
func (m_ MatrixRandomDistributionDescriptor) SetMaximum(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximum:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/minimum
func (m_ MatrixRandomDistributionDescriptor) Minimum() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minimum"))
	return rv
}


// SetMinimum sets the value of the minimum property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/minimum
func (m_ MatrixRandomDistributionDescriptor) SetMinimum(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimum:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/standardDeviation
func (m_ MatrixRandomDistributionDescriptor) StandardDeviation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("standardDeviation"))
	return rv
}


// SetStandardDeviation sets the value of the standardDeviation property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor/standardDeviation
func (m_ MatrixRandomDistributionDescriptor) SetStandardDeviation(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStandardDeviation:"), value)
}



