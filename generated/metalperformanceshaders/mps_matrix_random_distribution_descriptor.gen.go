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
	DistributionType() MatrixRandomDistribution get set /* not a class type */
	SetDistributionType(value MatrixRandomDistribution get set /* not a class type */)
	Maximum() objectivec.IObject
	SetMaximum(value objectivec.IObject)
	Mean() objectivec.IObject
	SetMean(value objectivec.IObject)
	Minimum() objectivec.IObject
	SetMinimum(value objectivec.IObject)
	StandardDeviation() objectivec.IObject
	SetStandardDeviation(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomDistributionDescriptorClass) Alloc() MatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomDistributionDescriptor
type MatrixRandomDistributionDescriptor struct {
	objectivec.Object
}

// MatrixRandomDistributionDescriptorFrom constructs a [MatrixRandomDistributionDescriptor] from an unsafe.Pointer.
func MatrixRandomDistributionDescriptorFrom(ptr unsafe.Pointer) MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242856-default
func (mc _MatrixRandomDistributionDescriptorClass) `default`() {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("`default`"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242856-defaultdistributiondescriptor
func (mc _MatrixRandomDistributionDescriptorClass) DefaultDistributionDescriptor() IMatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("defaultDistributionDescriptor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242862-uniformdistributiondescriptor
func (mc _MatrixRandomDistributionDescriptorClass) UniformDistributionDescriptor() {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("uniformDistributionDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242862-uniformdistributiondescriptorwit
func (mc _MatrixRandomDistributionDescriptorClass) UniformDistributionDescriptorWithMinimumMaximum(minimum float32, maximum float32) IMatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("uniformDistributionDescriptorWithMinimum:maximum:"), minimum, maximum)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3547979-normaldistributiondescriptor
func (mc _MatrixRandomDistributionDescriptorClass) NormalDistributionDescriptor() {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("normalDistributionDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3547979-normaldistributiondescriptorwith
func (mc _MatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviation(mean float32, standardDeviation float32) IMatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:"), mean, standardDeviation)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3547980-normaldistributiondescriptorwith
func (mc _MatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviationMinimumMaximum(mean float32, standardDeviation float32, minimum float32, maximum float32) IMatrixRandomDistributionDescriptor {
	rv := objc.Send[MatrixRandomDistributionDescriptor](objc.ID(mc.class), objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:minimum:maximum:"), mean, standardDeviation, minimum, maximum)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242857-distributiontype
func (m_ MatrixRandomDistributionDescriptor) DistributionType() MatrixRandomDistribution get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("distributionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242857-distributiontype
func (m_ MatrixRandomDistributionDescriptor) SetDistributionType(value MatrixRandomDistribution get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDistributionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242858-maximum
func (m_ MatrixRandomDistributionDescriptor) Maximum() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("maximum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242858-maximum
func (m_ MatrixRandomDistributionDescriptor) SetMaximum(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242859-mean
func (m_ MatrixRandomDistributionDescriptor) Mean() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("mean"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242859-mean
func (m_ MatrixRandomDistributionDescriptor) SetMean(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMean:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242860-minimum
func (m_ MatrixRandomDistributionDescriptor) Minimum() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("minimum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242860-minimum
func (m_ MatrixRandomDistributionDescriptor) SetMinimum(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242861-standarddeviation
func (m_ MatrixRandomDistributionDescriptor) StandardDeviation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("standardDeviation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242861-standarddeviation
func (m_ MatrixRandomDistributionDescriptor) SetStandardDeviation(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStandardDeviation:"), value)
}








