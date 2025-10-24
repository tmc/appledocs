// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSTemporaryNDArray */


/* debug [class_header]: Header for MPSTemporaryNDArray */
// The class instance for the [TemporaryNDArray] class.
var (
	TemporaryNDArrayClass     _TemporaryNDArrayClass
	TemporaryNDArrayClassOnce sync.Once
)

func getTemporaryNDArrayClass() _TemporaryNDArrayClass {
	TemporaryNDArrayClassOnce.Do(func() {
		TemporaryNDArrayClass = _TemporaryNDArrayClass{objc.GetClass("MPSTemporaryNDArray")}
	})
	return TemporaryNDArrayClass
}

type _TemporaryNDArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TemporaryNDArray */
// An interface definition for the [TemporaryNDArray] class.
type ITemporaryNDArray interface {
	INDArray
	
/* debug [class_interface_properties]: Properties for TemporaryNDArray */
	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TemporaryNDArray */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TemporaryNDArray */
// Alloc allocates a new instance without initialization.
func (tc _TemporaryNDArrayClass) Alloc() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TemporaryNDArrayClass) New() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporaryNDArray) Init() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporaryNDArray) Autorelease() TemporaryNDArray {
	rv := objc.Send[TemporaryNDArray](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporaryNDArray creates a new TemporaryNDArray instance.
func NewTemporaryNDArray() TemporaryNDArray {
	return getTemporaryNDArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TemporaryNDArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryNDArray
type TemporaryNDArray struct {
	NDArray
}

// TemporaryNDArrayFrom constructs a [TemporaryNDArray] from an unsafe.Pointer.
func TemporaryNDArrayFrom(ptr unsafe.Pointer) TemporaryNDArray {
	return TemporaryNDArray{
		NDArray: NDArrayFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TemporaryNDArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TemporaryNDArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3114075-temporaryndarraywithcommandbuffe
func (tc _TemporaryNDArrayClass) TemporaryNDArrayWithCommandBufferDescriptor(commandBuffer unsafe.Pointer, descriptor INDArrayDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryNDArrayWithCommandBuffer:descriptor:"), commandBuffer, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryNDArrayWithCommandBufferDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3131732-defaultallocator
func (tc _TemporaryNDArrayClass) DefaultAllocator() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("defaultAllocator"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultAllocator) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TemporaryNDArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TemporaryNDArray */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TemporaryNDArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3114074-readcount
func (t_ TemporaryNDArray) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("readCount"))
	return rv
}/* debug [instance_properties/getter]: readCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryndarray/3114074-readcount
func (t_ TemporaryNDArray) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}/* debug [instance_properties/setter]: readCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSTemporaryNDArray */



