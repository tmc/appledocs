// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [Slice] class.
type ISlice interface {
	ICNNKernel
	

	// properties:


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnslice/2942403-initwithcoder
func NewSliceWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) Slice {
	instance := getSliceClass().Alloc()
	rv := objc.Send[Slice](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnslice/2942401-initwithdevice
func NewSliceWithDevice(device unsafe.Pointer) Slice {
	instance := getSliceClass().Alloc()
	rv := objc.Send[Slice](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























