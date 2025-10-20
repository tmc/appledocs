// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TemporaryImage] class.
var (
	TemporaryImageClass     _TemporaryImageClass
	TemporaryImageClassOnce sync.Once
)

func getTemporaryImageClass() _TemporaryImageClass {
	TemporaryImageClassOnce.Do(func() {
		TemporaryImageClass = _TemporaryImageClass{objc.GetClass("MPSTemporaryImage")}
	})
	return TemporaryImageClass
}

type _TemporaryImageClass struct {
	class objc.Class
}

// An interface definition for the [TemporaryImage] class.
type ITemporaryImage interface {
	IImage
}

// A texture for use in convolutional neural networks that stores transient data to be used and discarded promptly.
//
// objects can provide a profound reduction in the aggregate texture memory and associated CPU-side allocation cost in your app. Metal Performance Shaders achieves this by automatically identifying objects that do not overlap in time over the course of a object’s lifetime and can therefore reuse the same memory. objects leverage an internal cache of preallocated reusable memory to hold pixel data to avoid typical memory allocation performance penalties common to ordinary and objects. To avoid data corruption due to aliasing, objects impose some important restrictions: The underlying texture storage mode is . You cannot, for example, use the or methods with them. Temporary images are strictly read and written by the GPU. The temporary image may be used only on a single object. This limits the chronology to a single linear time stream. The property must be managed correctly. Temporary images must also adhere to the general pixel format restrictions for objects. Since temporary images can only be used with a single command buffer, and can not be used off the GPU, they generally should not be kept around past the completion of their associated command buffer. The lifetime of a temporary image is typically expected to be extremely short, perhaps spanning only a few lines of code. To keep the lifetime of the underlying texture allocation as short as possible, the texture is not allocated until the first time the object is used by an object or until the first time the property is read. The property serves to limit the lifetime of the texture on deallocation. You may use the property with the methods of an subclass, if and the texture conforms to the requirements of the given kernel. In such cases, the property is not modified, since the enclosing object is not available. There is no locking mechanism provided to prevent a object returned from the property from becoming invalid when the value of the property reaches 0. objects can otherwise be used wherever objects are used.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporaryImage
type TemporaryImage struct {
	Image
}

// TemporaryImageFrom constructs a [TemporaryImage] from an unsafe.Pointer.
//
// A texture for use in convolutional neural networks that stores transient data to be used and discarded promptly.
func TemporaryImageFrom(ptr unsafe.Pointer) TemporaryImage {
	return TemporaryImage{
		Image: ImageFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TemporaryImageClass) Alloc() TemporaryImage {
	rv := objc.Send[TemporaryImage](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TemporaryImageClass) New() TemporaryImage {
	rv := objc.Send[TemporaryImage](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporaryImage) Init() TemporaryImage {
	rv := objc.Send[TemporaryImage](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporaryImage) Autorelease() TemporaryImage {
	rv := objc.Send[TemporaryImage](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporaryImage creates a new TemporaryImage instance.
func NewTemporaryImage() TemporaryImage {
	return getTemporaryImageClass().New()
}




