// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNSlice */


/* debug [class_header]: Header for MPSNNSlice */
// The class instance for the [Slice] class.
var (
	SliceClass     _SliceClass
	SliceClassOnce sync.Once
)

func getSliceClass() _SliceClass {
	SliceClassOnce.Do(func() {
		SliceClass = _SliceClass{objc.GetClass("MPSNNSlice")}
	})
	return SliceClass
}

type _SliceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Slice */
// An interface definition for the [Slice] class.
type ISlice interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for Slice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Slice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Slice */
// Alloc allocates a new instance without initialization.
func (sc _SliceClass) Alloc() Slice {
	rv := objc.Send[Slice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SliceClass) New() Slice {
	rv := objc.Send[Slice](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Slice) Init() Slice {
	rv := objc.Send[Slice](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Slice) Autorelease() Slice {
	rv := objc.Send[Slice](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSlice creates a new Slice instance.
func NewSlice() Slice {
	return getSliceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Slice */
// A kernel that extracts a slice from an image.


// A kernel that extracts a slice from an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSlice
type Slice struct {
	CNNKernel
}

// SliceFrom constructs a [Slice] from an unsafe.Pointer.
//
// A kernel that extracts a slice from an image.
func SliceFrom(ptr unsafe.Pointer) Slice {
	return Slice{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Slice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnslice/2942403-initwithcoder
func NewSliceWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) Slice {
	instance := getSliceClass().Alloc()
	rv := objc.Send[Slice](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSliceWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnslice/2942401-initwithdevice
func NewSliceWithDevice(device unsafe.Pointer) Slice {
	instance := getSliceClass().Alloc()
	rv := objc.Send[Slice](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSliceWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Slice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Slice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Slice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Slice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNSlice */


