// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSTemporaryMatrix */


/* debug [class_header]: Header for MPSTemporaryMatrix */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TemporaryMatrix */
// An interface definition for the [TemporaryMatrix] class.
type ITemporaryMatrix interface {
	IMatrix
	
/* debug [class_interface_properties]: Properties for TemporaryMatrix */
	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TemporaryMatrix */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TemporaryMatrix */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TemporaryMatrix */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TemporaryMatrix *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TemporaryMatrix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867073-prefetchstorage
func (tc _TemporaryMatrixClass) PrefetchStorage() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorage"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrefetchStorage) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867073-prefetchstoragewithcommandbuffer
func (tc _TemporaryMatrixClass) PrefetchStorageWithCommandBufferMatrixDescriptorList(commandBuffer unsafe.Pointer, descriptorList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorageWithCommandBuffer:matrixDescriptorList:"), commandBuffer, descriptorList)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrefetchStorageWithCommandBufferMatrixDescriptorList) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867180-temporarymatrixwithcommandbuffer
func (tc _TemporaryMatrixClass) TemporaryMatrixWithCommandBufferMatrixDescriptor(commandBuffer unsafe.Pointer, matrixDescriptor IMatrixDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryMatrixWithCommandBuffer:matrixDescriptor:"), commandBuffer, matrixDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryMatrixWithCommandBufferMatrixDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TemporaryMatrix */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TemporaryMatrix */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TemporaryMatrix */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867151-readcount
func (t_ TemporaryMatrix) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("readCount"))
	return rv
}/* debug [instance_properties/getter]: readCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporarymatrix/2867151-readcount
func (t_ TemporaryMatrix) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}/* debug [instance_properties/setter]: readCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSTemporaryMatrix */



