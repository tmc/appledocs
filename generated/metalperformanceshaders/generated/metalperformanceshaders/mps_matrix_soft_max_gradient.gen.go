// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixSoftMaxGradient */


/* debug [class_header]: Header for MPSMatrixSoftMaxGradient */
// The class instance for the [MatrixSoftMaxGradient] class.
var (
	MatrixSoftMaxGradientClass     _MatrixSoftMaxGradientClass
	MatrixSoftMaxGradientClassOnce sync.Once
)

func getMatrixSoftMaxGradientClass() _MatrixSoftMaxGradientClass {
	MatrixSoftMaxGradientClassOnce.Do(func() {
		MatrixSoftMaxGradientClass = _MatrixSoftMaxGradientClass{objc.GetClass("MPSMatrixSoftMaxGradient")}
	})
	return MatrixSoftMaxGradientClass
}

type _MatrixSoftMaxGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixSoftMaxGradient */
// An interface definition for the [MatrixSoftMaxGradient] class.
type IMatrixSoftMaxGradient interface {
	IMatrixBinaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixSoftMaxGradient */
	// properties:
	SourceColumns() objectivec.IObject
	SetSourceColumns(value objectivec.IObject)
	SourceRows() objectivec.IObject
	SetSourceRows(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixSoftMaxGradient */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	Encode()
	EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixSoftMaxGradient */
// Alloc allocates a new instance without initialization.
func (mc _MatrixSoftMaxGradientClass) Alloc() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSoftMaxGradientClass) New() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSoftMaxGradient) Init() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSoftMaxGradient) Autorelease() MatrixSoftMaxGradient {
	rv := objc.Send[MatrixSoftMaxGradient](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSoftMaxGradient creates a new MatrixSoftMaxGradient instance.
func NewMatrixSoftMaxGradient() MatrixSoftMaxGradient {
	return getMatrixSoftMaxGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixSoftMaxGradient */
// A gradient softmax kernel that operates on matrices.


// A gradient softmax kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMaxGradient
type MatrixSoftMaxGradient struct {
	MatrixBinaryKernel
}

// MatrixSoftMaxGradientFrom constructs a [MatrixSoftMaxGradient] from an unsafe.Pointer.
//
// A gradient softmax kernel that operates on matrices.
func MatrixSoftMaxGradientFrom(ptr unsafe.Pointer) MatrixSoftMaxGradient {
	return MatrixSoftMaxGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixSoftMaxGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966653-initwithcoder
func NewMatrixSoftMaxGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) MatrixSoftMaxGradient {
	instance := getMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[MatrixSoftMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixSoftMaxGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966654-initwithdevice
func NewMatrixSoftMaxGradientWithDevice(device unsafe.Pointer) MatrixSoftMaxGradient {
	instance := getMatrixSoftMaxGradientClass().Alloc()
	rv := objc.Send[MatrixSoftMaxGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixSoftMaxGradientWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixSoftMaxGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixSoftMaxGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixSoftMaxGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966651-copywithzone
func (m_ MatrixSoftMaxGradient) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966652-encode
func (m_ MatrixSoftMaxGradient) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966652-encodetocommandbuffer
func (m_ MatrixSoftMaxGradient) EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer unsafe.Pointer, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:gradientMatrix:forwardOutputMatrix:resultMatrix:"), commandBuffer, gradientMatrix, forwardOutputMatrix, resultMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixSoftMaxGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966655-sourcecolumns
func (m_ MatrixSoftMaxGradient) SourceColumns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceColumns"))
	return rv
}/* debug [instance_properties/getter]: sourceColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966655-sourcecolumns
func (m_ MatrixSoftMaxGradient) SetSourceColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceColumns:"), value)
}/* debug [instance_properties/setter]: sourceColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966656-sourcerows
func (m_ MatrixSoftMaxGradient) SourceRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceRows"))
	return rv
}/* debug [instance_properties/getter]: sourceRows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966656-sourcerows
func (m_ MatrixSoftMaxGradient) SetSourceRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRows:"), value)
}/* debug [instance_properties/setter]: sourceRows */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixSoftMaxGradient */


