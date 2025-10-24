// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NDArrayMultiaryBase] class.
var (
	NDArrayMultiaryBaseClass     _NDArrayMultiaryBaseClass
	NDArrayMultiaryBaseClassOnce sync.Once
)

func getNDArrayMultiaryBaseClass() _NDArrayMultiaryBaseClass {
	NDArrayMultiaryBaseClassOnce.Do(func() {
		NDArrayMultiaryBaseClass = _NDArrayMultiaryBaseClass{objc.GetClass("MPSNDArrayMultiaryBase")}
	})
	return NDArrayMultiaryBaseClass
}

type _NDArrayMultiaryBaseClass struct {
	class objc.Class
}

// An interface definition for the [NDArrayMultiaryBase] class.
type INDArrayMultiaryBase interface {
	IKernel
	// properties:
	DestinationArrayAllocator() NDArrayAllocator /* not a class type */
	SetDestinationArrayAllocator(value NDArrayAllocator /* not a class type */)
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device objectivec.IObject) unsafe.Pointer
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase
type NDArrayMultiaryBase struct {
	Kernel
}

// NDArrayMultiaryBaseFrom constructs a [NDArrayMultiaryBase] from an unsafe.Pointer.
func NDArrayMultiaryBaseFrom(ptr unsafe.Pointer) NDArrayMultiaryBase {
	return NDArrayMultiaryBase{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryBaseClass) Alloc() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NDArrayMultiaryBaseClass) New() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMultiaryBase) Init() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMultiaryBase) Autorelease() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMultiaryBase creates a new NDArrayMultiaryBase instance.
func NewNDArrayMultiaryBase() NDArrayMultiaryBase {
	return getNDArrayMultiaryBaseClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/copy(with:device:)
func (n_ NDArrayMultiaryBase) CopyWithZoneDevice(zone Zone /* not a class type */, device objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/destinationarrayallocator
func (n_ NDArrayMultiaryBase) DestinationArrayAllocator() NDArrayAllocator /* not a class type */ {
	rv := objc.Send[NDArrayAllocator](n_.ID, objc.Sel("destinationArrayAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/destinationarrayallocator
func (n_ NDArrayMultiaryBase) SetDestinationArrayAllocator(value NDArrayAllocator /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationArrayAllocator:"), value)
}



