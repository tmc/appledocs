// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MatrixCopyDescriptor] class.
var (
	MatrixCopyDescriptorClass     _MatrixCopyDescriptorClass
	MatrixCopyDescriptorClassOnce sync.Once
)

func getMatrixCopyDescriptorClass() _MatrixCopyDescriptorClass {
	MatrixCopyDescriptorClassOnce.Do(func() {
		MatrixCopyDescriptorClass = _MatrixCopyDescriptorClass{objc.GetClass("MPSMatrixCopyDescriptor")}
	})
	return MatrixCopyDescriptorClass
}

type _MatrixCopyDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MatrixCopyDescriptor] class.
type IMatrixCopyDescriptor interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	SetCopyOperationAt()
	SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets(index uint, sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets)


}





// Alloc allocates a new instance without initialization.
func (mc _MatrixCopyDescriptorClass) Alloc() MatrixCopyDescriptor {
	rv := objc.Send[MatrixCopyDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixCopyDescriptorClass) New() MatrixCopyDescriptor {
	rv := objc.Send[MatrixCopyDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixCopyDescriptor) Init() MatrixCopyDescriptor {
	rv := objc.Send[MatrixCopyDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixCopyDescriptor) Autorelease() MatrixCopyDescriptor {
	rv := objc.Send[MatrixCopyDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixCopyDescriptor creates a new MatrixCopyDescriptor instance.
func NewMatrixCopyDescriptor() MatrixCopyDescriptor {
	return getMatrixCopyDescriptorClass().New()
}





// A description of multiple matrix copy operations.


// A description of multiple matrix copy operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyDescriptor
type MatrixCopyDescriptor struct {
	objectivec.Object
}

// MatrixCopyDescriptorFrom constructs a [MatrixCopyDescriptor] from an unsafe.Pointer.
//
// A description of multiple matrix copy operations.
func MatrixCopyDescriptorFrom(ptr unsafe.Pointer) MatrixCopyDescriptor {
	return MatrixCopyDescriptor{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915324-initwithdevice
func NewMatrixCopyDescriptorWithDeviceCount(device unsafe.Pointer, count uint) MatrixCopyDescriptor {
	instance := getMatrixCopyDescriptorClass().Alloc()
	rv := objc.Send[MatrixCopyDescriptor](instance.ID, objc.Sel("initWithDevice:count:"), device, count)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915344-initwithsourcematrices
func NewMatrixCopyDescriptorWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices unsafe.Pointer, destinationMatrices unsafe.Pointer, offsets IVector, byteOffset uint) MatrixCopyDescriptor {
	instance := getMatrixCopyDescriptorClass().Alloc()
	rv := objc.Send[MatrixCopyDescriptor](instance.ID, objc.Sel("initWithSourceMatrices:destinationMatrices:offsetVector:offset:"), sourceMatrices, destinationMatrices, offsets, byteOffset)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915333-descriptorwithsourcematrix
func (mc _MatrixCopyDescriptorClass) DescriptorWithSourceMatrixDestinationMatrixOffsets(sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("descriptorWithSourceMatrix:destinationMatrix:offsets:"), sourceMatrix, destinationMatrix, offsets)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915331-setcopyoperationat
func (m_ MatrixCopyDescriptor) SetCopyOperationAt() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCopyOperationAt"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915331-setcopyoperationatindex
func (m_ MatrixCopyDescriptor) SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets(index uint, sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCopyOperationAtIndex:sourceMatrix:destinationMatrix:offsets:"), index, sourceMatrix, destinationMatrix, offsets)
}












