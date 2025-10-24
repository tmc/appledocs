// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixBinaryKernel] class.
var (
	MatrixBinaryKernelClass     _MatrixBinaryKernelClass
	MatrixBinaryKernelClassOnce sync.Once
)

func getMatrixBinaryKernelClass() _MatrixBinaryKernelClass {
	MatrixBinaryKernelClassOnce.Do(func() {
		MatrixBinaryKernelClass = _MatrixBinaryKernelClass{objc.GetClass("MPSMatrixBinaryKernel")}
	})
	return MatrixBinaryKernelClass
}

type _MatrixBinaryKernelClass struct {
	class objc.Class
}





// An interface definition for the [MatrixBinaryKernel] class.
type IMatrixBinaryKernel interface {
	IKernel
	

	// properties:
	BatchSize() objectivec.IObject
	SetBatchSize(value objectivec.IObject)
	SecondarySourceMatrixOrigin() Origin get set /* not a class type */
	SetSecondarySourceMatrixOrigin(value Origin get set /* not a class type */)
	BatchStart() objectivec.IObject
	SetBatchStart(value objectivec.IObject)
	PrimarySourceMatrixOrigin() Origin get set /* not a class type */
	SetPrimarySourceMatrixOrigin(value Origin get set /* not a class type */)
	ResultMatrixOrigin() Origin get set /* not a class type */
	SetResultMatrixOrigin(value Origin get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixBinaryKernelClass) Alloc() MatrixBinaryKernel {
	rv := objc.Send[MatrixBinaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixBinaryKernelClass) New() MatrixBinaryKernel {
	rv := objc.Send[MatrixBinaryKernel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixBinaryKernel) Init() MatrixBinaryKernel {
	rv := objc.Send[MatrixBinaryKernel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixBinaryKernel) Autorelease() MatrixBinaryKernel {
	rv := objc.Send[MatrixBinaryKernel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixBinaryKernel creates a new MatrixBinaryKernel instance.
func NewMatrixBinaryKernel() MatrixBinaryKernel {
	return getMatrixBinaryKernelClass().New()
}





// A kernel that consumes two matrices and produces one matrix.


// A kernel that consumes two matrices and produces one matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixBinaryKernel
type MatrixBinaryKernel struct {
	Kernel
}

// MatrixBinaryKernelFrom constructs a [MatrixBinaryKernel] from an unsafe.Pointer.
//
// A kernel that consumes two matrices and produces one matrix.
func MatrixBinaryKernelFrom(ptr unsafe.Pointer) MatrixBinaryKernel {
	return MatrixBinaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867089-batchsize
func (m_ MatrixBinaryKernel) BatchSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867089-batchsize
func (m_ MatrixBinaryKernel) SetBatchSize(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867096-secondarysourcematrixorigin
func (m_ MatrixBinaryKernel) SecondarySourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("secondarySourceMatrixOrigin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867096-secondarysourcematrixorigin
func (m_ MatrixBinaryKernel) SetSecondarySourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecondarySourceMatrixOrigin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867152-batchstart
func (m_ MatrixBinaryKernel) BatchStart() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchStart"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867152-batchstart
func (m_ MatrixBinaryKernel) SetBatchStart(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchStart:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867182-primarysourcematrixorigin
func (m_ MatrixBinaryKernel) PrimarySourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("primarySourceMatrixOrigin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867182-primarysourcematrixorigin
func (m_ MatrixBinaryKernel) SetPrimarySourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimarySourceMatrixOrigin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867193-resultmatrixorigin
func (m_ MatrixBinaryKernel) ResultMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultMatrixOrigin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867193-resultmatrixorigin
func (m_ MatrixBinaryKernel) SetResultMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultMatrixOrigin:"), value)
}








