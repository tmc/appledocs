// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TemporaryMatrix] class.
var (
	TemporaryMatrixClass     _TemporaryMatrixClass
	TemporaryMatrixClassOnce sync.Once
)

func getTemporaryMatrixClass() _TemporaryMatrixClass {
	TemporaryMatrixClassOnce.Do(func() {
		TemporaryMatrixClass = _TemporaryMatrixClass{objc.GetClass("MPSTemporaryMatrix")}
	})
	return TemporaryMatrixClass
}

type _TemporaryMatrixClass struct {
	class objc.Class
}





// An interface definition for the [TemporaryMatrix] class.
type ITemporaryMatrix interface {
	IMatrix
	

	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TemporaryMatrixClass) Alloc() TemporaryMatrix {
	rv := objc.Send[TemporaryMatrix](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TemporaryMatrixClass) New() TemporaryMatrix {
	rv := objc.Send[TemporaryMatrix](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporaryMatrix) Init() TemporaryMatrix {
	rv := objc.Send[TemporaryMatrix](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporaryMatrix) Autorelease() TemporaryMatrix {
	rv := objc.Send[TemporaryMatrix](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporaryMatrix creates a new TemporaryMatrix instance.
func NewTemporaryMatrix() TemporaryMatrix {
	return getTemporaryMatrixClass().New()
}





// A matrix allocated on GPU private memory.


// A matrix allocated on GPU private memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryMatrix
type TemporaryMatrix struct {
	Matrix
}

// TemporaryMatrixFrom constructs a [TemporaryMatrix] from an unsafe.Pointer.
//
// A matrix allocated on GPU private memory.
func TemporaryMatrixFrom(ptr unsafe.Pointer) TemporaryMatrix {
	return TemporaryMatrix{
		Matrix: MatrixFrom(ptr),
	}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867073-prefetchstorage
func (tc _TemporaryMatrixClass) PrefetchStorage() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorage"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867073-prefetchstoragewithcommandbuffer
func (tc _TemporaryMatrixClass) PrefetchStorageWithCommandBufferMatrixDescriptorList(commandBuffer unsafe.Pointer, descriptorList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorageWithCommandBuffer:matrixDescriptorList:"), commandBuffer, descriptorList)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867180-temporarymatrixwithcommandbuffer
func (tc _TemporaryMatrixClass) TemporaryMatrixWithCommandBufferMatrixDescriptor(commandBuffer unsafe.Pointer, matrixDescriptor IMatrixDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryMatrixWithCommandBuffer:matrixDescriptor:"), commandBuffer, matrixDescriptor)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867151-readcount
func (t_ TemporaryMatrix) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("readCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867151-readcount
func (t_ TemporaryMatrix) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}








