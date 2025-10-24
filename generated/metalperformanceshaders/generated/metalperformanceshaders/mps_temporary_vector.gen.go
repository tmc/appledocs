// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSTemporaryVector */


/* debug [class_header]: Header for MPSTemporaryVector */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TemporaryVector */
// An interface definition for the [TemporaryVector] class.
type ITemporaryVector interface {
	IVector
	
/* debug [class_interface_properties]: Properties for TemporaryVector */
	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TemporaryVector */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TemporaryVector */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TemporaryVector */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TemporaryVector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TemporaryVector */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935544-prefetchstorage
func (tc _TemporaryVectorClass) PrefetchStorage() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorage"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrefetchStorage) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935544-prefetchstoragewithcommandbuffer
func (tc _TemporaryVectorClass) PrefetchStorageWithCommandBufferDescriptorList(commandBuffer unsafe.Pointer, descriptorList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorageWithCommandBuffer:descriptorList:"), commandBuffer, descriptorList)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrefetchStorageWithCommandBufferDescriptorList) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935550-temporaryvectorwithcommandbuffer
func (tc _TemporaryVectorClass) TemporaryVectorWithCommandBufferDescriptor(commandBuffer unsafe.Pointer, descriptor IVectorDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryVectorWithCommandBuffer:descriptor:"), commandBuffer, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryVectorWithCommandBufferDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TemporaryVector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TemporaryVector */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TemporaryVector */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935547-readcount
func (t_ TemporaryVector) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("readCount"))
	return rv
}/* debug [instance_properties/getter]: readCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryvector/2935547-readcount
func (t_ TemporaryVector) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}/* debug [instance_properties/setter]: readCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSTemporaryVector */



