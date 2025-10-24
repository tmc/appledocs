// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TemporaryVector] class.
var (
	TemporaryVectorClass     _TemporaryVectorClass
	TemporaryVectorClassOnce sync.Once
)

func getTemporaryVectorClass() _TemporaryVectorClass {
	TemporaryVectorClassOnce.Do(func() {
		TemporaryVectorClass = _TemporaryVectorClass{objc.GetClass("MPSTemporaryVector")}
	})
	return TemporaryVectorClass
}

type _TemporaryVectorClass struct {
	class objc.Class
}





// An interface definition for the [TemporaryVector] class.
type ITemporaryVector interface {
	IVector
	

	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TemporaryVectorClass) Alloc() TemporaryVector {
	rv := objc.Send[TemporaryVector](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TemporaryVectorClass) New() TemporaryVector {
	rv := objc.Send[TemporaryVector](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporaryVector) Init() TemporaryVector {
	rv := objc.Send[TemporaryVector](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporaryVector) Autorelease() TemporaryVector {
	rv := objc.Send[TemporaryVector](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporaryVector creates a new TemporaryVector instance.
func NewTemporaryVector() TemporaryVector {
	return getTemporaryVectorClass().New()
}





// A vector allocated on GPU private memory.


// A vector allocated on GPU private memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryVector
type TemporaryVector struct {
	Vector
}

// TemporaryVectorFrom constructs a [TemporaryVector] from an unsafe.Pointer.
//
// A vector allocated on GPU private memory.
func TemporaryVectorFrom(ptr unsafe.Pointer) TemporaryVector {
	return TemporaryVector{
		Vector: VectorFrom(ptr),
	}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935544-prefetchstorage
func (tc _TemporaryVectorClass) PrefetchStorage() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorage"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935544-prefetchstoragewithcommandbuffer
func (tc _TemporaryVectorClass) PrefetchStorageWithCommandBufferDescriptorList(commandBuffer unsafe.Pointer, descriptorList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorageWithCommandBuffer:descriptorList:"), commandBuffer, descriptorList)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935550-temporaryvectorwithcommandbuffer
func (tc _TemporaryVectorClass) TemporaryVectorWithCommandBufferDescriptor(commandBuffer unsafe.Pointer, descriptor IVectorDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryVectorWithCommandBuffer:descriptor:"), commandBuffer, descriptor)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935547-readcount
func (t_ TemporaryVector) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("readCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935547-readcount
func (t_ TemporaryVector) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}








