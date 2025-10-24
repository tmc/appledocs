// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixUnaryKernel */


/* debug [class_header]: Header for MPSMatrixUnaryKernel */
// The class instance for the [MatrixUnaryKernel] class.
var (
	MatrixUnaryKernelClass     _MatrixUnaryKernelClass
	MatrixUnaryKernelClassOnce sync.Once
)

func getMatrixUnaryKernelClass() _MatrixUnaryKernelClass {
	MatrixUnaryKernelClassOnce.Do(func() {
		MatrixUnaryKernelClass = _MatrixUnaryKernelClass{objc.GetClass("MPSMatrixUnaryKernel")}
	})
	return MatrixUnaryKernelClass
}

type _MatrixUnaryKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixUnaryKernel */
// An interface definition for the [MatrixUnaryKernel] class.
type IMatrixUnaryKernel interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for MatrixUnaryKernel */
	// properties:
	BatchStart() objectivec.IObject
	SetBatchStart(value objectivec.IObject)
	SourceMatrixOrigin() Origin get set /* not a class type */
	SetSourceMatrixOrigin(value Origin get set /* not a class type */)
	BatchSize() objectivec.IObject
	SetBatchSize(value objectivec.IObject)
	ResultMatrixOrigin() Origin get set /* not a class type */
	SetResultMatrixOrigin(value Origin get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixUnaryKernel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixUnaryKernel */
// Alloc allocates a new instance without initialization.
func (mc _MatrixUnaryKernelClass) Alloc() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixUnaryKernelClass) New() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixUnaryKernel) Init() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixUnaryKernel) Autorelease() MatrixUnaryKernel {
	rv := objc.Send[MatrixUnaryKernel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixUnaryKernel creates a new MatrixUnaryKernel instance.
func NewMatrixUnaryKernel() MatrixUnaryKernel {
	return getMatrixUnaryKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixUnaryKernel */
// A kernel that consumes one matrix and produces one matrix.


// A kernel that consumes one matrix and produces one matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixUnaryKernel
type MatrixUnaryKernel struct {
	Kernel
}

// MatrixUnaryKernelFrom constructs a [MatrixUnaryKernel] from an unsafe.Pointer.
//
// A kernel that consumes one matrix and produces one matrix.
func MatrixUnaryKernelFrom(ptr unsafe.Pointer) MatrixUnaryKernel {
	return MatrixUnaryKernel{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixUnaryKernel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixUnaryKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixUnaryKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixUnaryKernel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixUnaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2866990-batchstart
func (m_ MatrixUnaryKernel) BatchStart() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchStart"))
	return rv
}/* debug [instance_properties/getter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2866990-batchstart
func (m_ MatrixUnaryKernel) SetBatchStart(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchStart:"), value)
}/* debug [instance_properties/setter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867053-sourcematrixorigin
func (m_ MatrixUnaryKernel) SourceMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("sourceMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: sourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867053-sourcematrixorigin
func (m_ MatrixUnaryKernel) SetSourceMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: sourceMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867118-batchsize
func (m_ MatrixUnaryKernel) BatchSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchSize"))
	return rv
}/* debug [instance_properties/getter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867118-batchsize
func (m_ MatrixUnaryKernel) SetBatchSize(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchSize:"), value)
}/* debug [instance_properties/setter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867150-resultmatrixorigin
func (m_ MatrixUnaryKernel) ResultMatrixOrigin() Origin get set /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("resultMatrixOrigin"))
	return rv
}/* debug [instance_properties/getter]: resultMatrixOrigin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867150-resultmatrixorigin
func (m_ MatrixUnaryKernel) SetResultMatrixOrigin(value Origin get set /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResultMatrixOrigin:"), value)
}/* debug [instance_properties/setter]: resultMatrixOrigin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixUnaryKernel */



