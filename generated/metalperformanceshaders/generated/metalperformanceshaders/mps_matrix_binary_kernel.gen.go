// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixBinaryKernel */


/* debug [class_header]: Header for MPSMatrixBinaryKernel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixBinaryKernel */
// An interface definition for the [MatrixBinaryKernel] class.
type IMatrixBinaryKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for MatrixBinaryKernel */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixBinaryKernel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixBinaryKernel */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixBinaryKernel */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixBinaryKernel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixBinaryKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixBinaryKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixBinaryKernel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixBinaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867089-batchsize
func (m_ MatrixBinaryKernel) BatchSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchSize"))
	return rv
}/* debug [instance_properties/getter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867089-batchsize
func (m_ MatrixBinaryKernel) SetBatchSize(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchSize:"), value)
}/* debug [instance_properties/setter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867096-secondarysourcematrixorigin
func (m_ MatrixBinaryKernel) SecondarySourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("secondarySourceMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: secondarySourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867096-secondarysourcematrixorigin
func (m_ MatrixBinaryKernel) SetSecondarySourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSecondarySourceMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: secondarySourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867152-batchstart
func (m_ MatrixBinaryKernel) BatchStart() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchStart"))
	return rv
}/* debug [instance_properties/getter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867152-batchstart
func (m_ MatrixBinaryKernel) SetBatchStart(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchStart:"), value)
}/* debug [instance_properties/setter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867182-primarysourcematrixorigin
func (m_ MatrixBinaryKernel) PrimarySourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("primarySourceMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: primarySourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867182-primarysourcematrixorigin
func (m_ MatrixBinaryKernel) SetPrimarySourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimarySourceMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: primarySourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867193-resultmatrixorigin
func (m_ MatrixBinaryKernel) ResultMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: resultMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867193-resultmatrixorigin
func (m_ MatrixBinaryKernel) SetResultMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: resultMatrixOrigin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixBinaryKernel */



