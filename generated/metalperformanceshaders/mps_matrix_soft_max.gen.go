// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixSoftMax */


/* debug [class_header]: Header for MPSMatrixSoftMax */
// The class instance for the [MatrixSoftMax] class.
var (
	MatrixSoftMaxClass     _MatrixSoftMaxClass
	MatrixSoftMaxClassOnce sync.Once
)

func getMatrixSoftMaxClass() _MatrixSoftMaxClass {
	MatrixSoftMaxClassOnce.Do(func() {
		MatrixSoftMaxClass = _MatrixSoftMaxClass{objc.GetClass("MPSMatrixSoftMax")}
	})
	return MatrixSoftMaxClass
}

type _MatrixSoftMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixSoftMax */
// An interface definition for the [MatrixSoftMax] class.
type IMatrixSoftMax interface {
	IMatrixUnaryKernel
	
/* debug [class_interface_properties]: Properties for MatrixSoftMax */
	// properties:
	SourceColumns() objectivec.IObject
	SetSourceColumns(value objectivec.IObject)
	SourceRows() objectivec.IObject
	SetSourceRows(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixSoftMax */
	// methods:
	Encode()
	EncodeToCommandBufferInputMatrixResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, resultMatrix IMatrix)
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixSoftMax */
// Alloc allocates a new instance without initialization.
func (mc _MatrixSoftMaxClass) Alloc() MatrixSoftMax {
	rv := objc.Send[MatrixSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixSoftMaxClass) New() MatrixSoftMax {
	rv := objc.Send[MatrixSoftMax](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixSoftMax) Init() MatrixSoftMax {
	rv := objc.Send[MatrixSoftMax](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixSoftMax) Autorelease() MatrixSoftMax {
	rv := objc.Send[MatrixSoftMax](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixSoftMax creates a new MatrixSoftMax instance.
func NewMatrixSoftMax() MatrixSoftMax {
	return getMatrixSoftMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixSoftMax */
// A softmax kernel that operates on matrices.


// A softmax kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixSoftMax
type MatrixSoftMax struct {
	MatrixUnaryKernel
}

// MatrixSoftMaxFrom constructs a [MatrixSoftMax] from an unsafe.Pointer.
//
// A softmax kernel that operates on matrices.
func MatrixSoftMaxFrom(ptr unsafe.Pointer) MatrixSoftMax {
	return MatrixSoftMax{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixSoftMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935565-initwithcoder
func NewMatrixSoftMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixSoftMax {
	instance := getMatrixSoftMaxClass().Alloc()
	rv := objc.Send[MatrixSoftMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixSoftMaxWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935562-initwithdevice
func NewMatrixSoftMaxWithDevice(device unsafe.Pointer) MatrixSoftMax {
	instance := getMatrixSoftMaxClass().Alloc()
	rv := objc.Send[MatrixSoftMax](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixSoftMaxWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixSoftMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixSoftMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixSoftMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935563-encode
func (m_ MatrixSoftMax) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935563-encodetocommandbuffer
func (m_ MatrixSoftMax) EncodeToCommandBufferInputMatrixResultMatrix(commandBuffer unsafe.Pointer, inputMatrix IMatrix, resultMatrix IMatrix) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:inputMatrix:resultMatrix:"), commandBuffer, inputMatrix, resultMatrix)
}/* debug [instance_methods/method]: EncodeToCommandBufferInputMatrixResultMatrix */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935566-copywithzone
func (m_ MatrixSoftMax) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixSoftMax */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935560-sourcecolumns
func (m_ MatrixSoftMax) SourceColumns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceColumns"))
	return rv
}/* debug [instance_properties/getter]: sourceColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935560-sourcecolumns
func (m_ MatrixSoftMax) SetSourceColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceColumns:"), value)
}/* debug [instance_properties/setter]: sourceColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935561-sourcerows
func (m_ MatrixSoftMax) SourceRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourceRows"))
	return rv
}/* debug [instance_properties/getter]: sourceRows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935561-sourcerows
func (m_ MatrixSoftMax) SetSourceRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceRows:"), value)
}/* debug [instance_properties/setter]: sourceRows */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixSoftMax */


