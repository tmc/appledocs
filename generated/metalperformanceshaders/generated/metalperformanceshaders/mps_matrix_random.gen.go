// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixRandom */


/* debug [class_header]: Header for MPSMatrixRandom */
// The class instance for the [MatrixRandom] class.
var (
	MatrixRandomClass     _MatrixRandomClass
	MatrixRandomClassOnce sync.Once
)

func getMatrixRandomClass() _MatrixRandomClass {
	MatrixRandomClassOnce.Do(func() {
		MatrixRandomClass = _MatrixRandomClass{objc.GetClass("MPSMatrixRandom")}
	})
	return MatrixRandomClass
}

type _MatrixRandomClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixRandom */
// An interface definition for the [MatrixRandom] class.
type IMatrixRandom interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for MatrixRandom */
	// properties:
	BatchSize() objectivec.IObject
	SetBatchSize(value objectivec.IObject)
	BatchStart() objectivec.IObject
	SetBatchStart(value objectivec.IObject)
	DestinationDataType() DataType get /* not a class type */
	SetDestinationDataType(value DataType get /* not a class type */)
	DistributionType() MatrixRandomDistribution get /* not a class type */
	SetDistributionType(value MatrixRandomDistribution get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixRandom */
	// methods:
	Encode()
	EncodeToCommandBufferDestinationVector(commandBuffer unsafe.Pointer, destinationVector IVector)
	EncodeToCommandBufferDestinationMatrix(commandBuffer unsafe.Pointer, destinationMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixRandom */
// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomClass) Alloc() MatrixRandom {
	rv := objc.Send[MatrixRandom](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixRandomClass) New() MatrixRandom {
	rv := objc.Send[MatrixRandom](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixRandom) Init() MatrixRandom {
	rv := objc.Send[MatrixRandom](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixRandom) Autorelease() MatrixRandom {
	rv := objc.Send[MatrixRandom](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixRandom creates a new MatrixRandom instance.
func NewMatrixRandom() MatrixRandom {
	return getMatrixRandomClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixRandom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandom
type MatrixRandom struct {
	Kernel
}

// MatrixRandomFrom constructs a [MatrixRandom] from an unsafe.Pointer.
func MatrixRandomFrom(ptr unsafe.Pointer) MatrixRandom {
	return MatrixRandom{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixRandom *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixRandom */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixRandom */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixRandom */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242851-encode
func (m_ MatrixRandom) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242851-encodetocommandbuffer
func (m_ MatrixRandom) EncodeToCommandBufferDestinationVector(commandBuffer unsafe.Pointer, destinationVector IVector) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:destinationVector:"), commandBuffer, destinationVector)
}/* debug [instance_methods/method]: EncodeToCommandBufferDestinationVector */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3325839-encodetocommandbuffer
func (m_ MatrixRandom) EncodeToCommandBufferDestinationMatrix(commandBuffer unsafe.Pointer, destinationMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:destinationMatrix:"), commandBuffer, destinationMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferDestinationMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixRandom */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242847-batchsize
func (m_ MatrixRandom) BatchSize() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchSize"))
	return rv
}/* debug [instance_properties/getter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242847-batchsize
func (m_ MatrixRandom) SetBatchSize(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchSize:"), value)
}/* debug [instance_properties/setter]: batchSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242848-batchstart
func (m_ MatrixRandom) BatchStart() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("batchStart"))
	return rv
}/* debug [instance_properties/getter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242848-batchstart
func (m_ MatrixRandom) SetBatchStart(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBatchStart:"), value)
}/* debug [instance_properties/setter]: batchStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242849-destinationdatatype
func (m_ MatrixRandom) DestinationDataType() DataType get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("destinationDataType"))
	return rv
}/* debug [instance_properties/getter]: destinationDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242849-destinationdatatype
func (m_ MatrixRandom) SetDestinationDataType(value DataType get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationDataType:"), value)
}/* debug [instance_properties/setter]: destinationDataType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242850-distributiontype
func (m_ MatrixRandom) DistributionType() MatrixRandomDistribution get /* not a class type */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("distributionType"))
	return rv
}/* debug [instance_properties/getter]: distributionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandom/3242850-distributiontype
func (m_ MatrixRandom) SetDistributionType(value MatrixRandomDistribution get /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDistributionType:"), value)
}/* debug [instance_properties/setter]: distributionType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixRandom */



