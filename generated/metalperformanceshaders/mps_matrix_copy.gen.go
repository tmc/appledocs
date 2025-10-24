// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixCopy */


/* debug [class_header]: Header for MPSMatrixCopy */
// The class instance for the [MatrixCopy] class.
var (
	MatrixCopyClass     _MatrixCopyClass
	MatrixCopyClassOnce sync.Once
)

func getMatrixCopyClass() _MatrixCopyClass {
	MatrixCopyClassOnce.Do(func() {
		MatrixCopyClass = _MatrixCopyClass{objc.GetClass("MPSMatrixCopy")}
	})
	return MatrixCopyClass
}

type _MatrixCopyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixCopy */
// An interface definition for the [MatrixCopy] class.
type IMatrixCopy interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for MatrixCopy */
	// properties:
	CopyColumns() objectivec.IObject
	SetCopyColumns(value objectivec.IObject)
	DestinationsAreTransposed() objectivec.IObject
	SetDestinationsAreTransposed(value objectivec.IObject)
	SourcesAreTransposed() objectivec.IObject
	SetSourcesAreTransposed(value objectivec.IObject)
	CopyRows() objectivec.IObject
	SetCopyRows(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixCopy */
	// methods:
	Encode()
	EncodeToCommandBufferCopyDescriptor(commandBuffer unsafe.Pointer, copyDescriptor IMatrixCopyDescriptor)
	EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset(commandBuffer unsafe.Pointer, copyDescriptor IMatrixCopyDescriptor, rowPermuteIndices IVector, rowPermuteOffset uint, columnPermuteIndices IVector, columnPermuteOffset uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixCopy */
// Alloc allocates a new instance without initialization.
func (mc _MatrixCopyClass) Alloc() MatrixCopy {
	rv := objc.Send[MatrixCopy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixCopyClass) New() MatrixCopy {
	rv := objc.Send[MatrixCopy](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixCopy) Init() MatrixCopy {
	rv := objc.Send[MatrixCopy](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixCopy) Autorelease() MatrixCopy {
	rv := objc.Send[MatrixCopy](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixCopy creates a new MatrixCopy instance.
func NewMatrixCopy() MatrixCopy {
	return getMatrixCopyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixCopy */
// A class that can perform multiple matrix copy operations.


// A class that can perform multiple matrix copy operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopy
type MatrixCopy struct {
	Kernel
}

// MatrixCopyFrom constructs a [MatrixCopy] from an unsafe.Pointer.
//
// A class that can perform multiple matrix copy operations.
func MatrixCopyFrom(ptr unsafe.Pointer) MatrixCopy {
	return MatrixCopy{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixCopy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915334-initwithcoder
func NewMatrixCopyWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) MatrixCopy {
	instance := getMatrixCopyClass().Alloc()
	rv := objc.Send[MatrixCopy](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixCopyWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915345-initwithdevice
func NewMatrixCopyWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed(device unsafe.Pointer, copyRows uint, copyColumns uint, sourcesAreTransposed bool, destinationsAreTransposed bool) MatrixCopy {
	instance := getMatrixCopyClass().Alloc()
	rv := objc.Send[MatrixCopy](instance.ID, objc.Sel("initWithDevice:copyRows:copyColumns:sourcesAreTransposed:destinationsAreTransposed:"), device, copyRows, copyColumns, sourcesAreTransposed, destinationsAreTransposed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixCopyWithDeviceCopyRowsCopyColumnsSourcesAreTransposedDestinationsAreTransposed */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixCopy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixCopy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixCopy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915341-encode
func (m_ MatrixCopy) Encode() {
	objc.Send[objc.ID](m_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915341-encodetocommandbuffer
func (m_ MatrixCopy) EncodeToCommandBufferCopyDescriptor(commandBuffer unsafe.Pointer, copyDescriptor IMatrixCopyDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:copyDescriptor:"), commandBuffer, copyDescriptor)
}/* debug [instance_methods/method]: EncodeToCommandBufferCopyDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2935558-encodetocommandbuffer
func (m_ MatrixCopy) EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset(commandBuffer unsafe.Pointer, copyDescriptor IMatrixCopyDescriptor, rowPermuteIndices IVector, rowPermuteOffset uint, columnPermuteIndices IVector, columnPermuteOffset uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("encodeToCommandBuffer:copyDescriptor:rowPermuteIndices:rowPermuteOffset:columnPermuteIndices:columnPermuteOffset:"), commandBuffer, copyDescriptor, rowPermuteIndices, rowPermuteOffset, columnPermuteIndices, columnPermuteOffset)
}/* debug [instance_methods/method]: EncodeToCommandBufferCopyDescriptorRowPermuteIndicesRowPermuteOffsetColumnPermuteIndicesColumnPermuteOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixCopy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915325-copycolumns
func (m_ MatrixCopy) CopyColumns() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyColumns"))
	return rv
}/* debug [instance_properties/getter]: copyColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915325-copycolumns
func (m_ MatrixCopy) SetCopyColumns(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCopyColumns:"), value)
}/* debug [instance_properties/setter]: copyColumns */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915326-destinationsaretransposed
func (m_ MatrixCopy) DestinationsAreTransposed() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("destinationsAreTransposed"))
	return rv
}/* debug [instance_properties/getter]: destinationsAreTransposed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915326-destinationsaretransposed
func (m_ MatrixCopy) SetDestinationsAreTransposed(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDestinationsAreTransposed:"), value)
}/* debug [instance_properties/setter]: destinationsAreTransposed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915340-sourcesaretransposed
func (m_ MatrixCopy) SourcesAreTransposed() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sourcesAreTransposed"))
	return rv
}/* debug [instance_properties/getter]: sourcesAreTransposed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915340-sourcesaretransposed
func (m_ MatrixCopy) SetSourcesAreTransposed(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourcesAreTransposed:"), value)
}/* debug [instance_properties/setter]: sourcesAreTransposed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915342-copyrows
func (m_ MatrixCopy) CopyRows() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("copyRows"))
	return rv
}/* debug [instance_properties/getter]: copyRows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopy/2915342-copyrows
func (m_ MatrixCopy) SetCopyRows(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCopyRows:"), value)
}/* debug [instance_properties/setter]: copyRows */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixCopy */


