// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNDArrayStridedSlice */


/* debug [class_header]: Header for MPSNDArrayStridedSlice */
// The class instance for the [NDArrayStridedSlice] class.
var (
	NDArrayStridedSliceClass     _NDArrayStridedSliceClass
	NDArrayStridedSliceClassOnce sync.Once
)

func getNDArrayStridedSliceClass() _NDArrayStridedSliceClass {
	NDArrayStridedSliceClassOnce.Do(func() {
		NDArrayStridedSliceClass = _NDArrayStridedSliceClass{objc.GetClass("MPSNDArrayStridedSlice")}
	})
	return NDArrayStridedSliceClass
}

type _NDArrayStridedSliceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayStridedSlice */
// An interface definition for the [NDArrayStridedSlice] class.
type INDArrayStridedSlice interface {
	INDArrayUnaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayStridedSlice */
	// properties:
	Strides() NDArrayOffsets get set /* not a class type */
	SetStrides(value NDArrayOffsets get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayStridedSlice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayStridedSlice */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayStridedSliceClass) Alloc() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayStridedSliceClass) New() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayStridedSlice) Init() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayStridedSlice) Autorelease() NDArrayStridedSlice {
	rv := objc.Send[NDArrayStridedSlice](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayStridedSlice creates a new NDArrayStridedSlice instance.
func NewNDArrayStridedSlice() NDArrayStridedSlice {
	return getNDArrayStridedSliceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayStridedSlice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSlice
type NDArrayStridedSlice struct {
	NDArrayUnaryKernel
}

// NDArrayStridedSliceFrom constructs a [NDArrayStridedSlice] from an unsafe.Pointer.
func NDArrayStridedSliceFrom(ptr unsafe.Pointer) NDArrayStridedSlice {
	return NDArrayStridedSlice{
		NDArrayUnaryKernel: NDArrayUnaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayStridedSlice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayStridedSlice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayStridedSlice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayStridedSlice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayStridedSlice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraystridedslice/3143546-strides
func (n_ NDArrayStridedSlice) Strides() NDArrayOffsets get set /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("strides"))
	return rv
}/* debug [instance_properties/getter]: strides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraystridedslice/3143546-strides
func (n_ NDArrayStridedSlice) SetStrides(value NDArrayOffsets get set /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setStrides:"), value)
}/* debug [instance_properties/setter]: strides */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayStridedSlice */



