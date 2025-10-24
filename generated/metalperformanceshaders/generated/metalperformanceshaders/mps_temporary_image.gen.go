// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSTemporaryImage */


/* debug [class_header]: Header for MPSTemporaryImage */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TemporaryImage */
// An interface definition for the [TemporaryImage] class.
type ITemporaryImage interface {
	IImage
	
/* debug [class_interface_properties]: Properties for TemporaryImage */
	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)
	Texture() Texture /* not a class type */
	SetTexture(value Texture /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TemporaryImage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TemporaryImage */
// Alloc allocates a new instance without initialization.
func (tc _TemporaryImageClass) Alloc() TemporaryImage {
	rv := objc.Send[TemporaryImage](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TemporaryImage */
// A texture for use in convolutional neural networks that stores transient data to be used and discarded promptly.
//
// objects can provide a profound reduction in the aggregate texture memory and associated CPU-side allocation cost in your app. Metal Performance Shaders achieves this by automatically identifying objects that do not overlap in time over the course of a object’s lifetime and can therefore reuse the same memory. objects leverage an internal cache of preallocated reusable memory to hold pixel data to avoid typical memory allocation performance penalties common to ordinary and objects. To avoid data corruption due to aliasing, objects impose some important restrictions: The underlying texture storage mode is . You cannot, for example, use the or methods with them. Temporary images are strictly read and written by the GPU. The temporary image may be used only on a single object. This limits the chronology to a single linear time stream. The property must be managed correctly. Temporary images must also adhere to the general pixel format restrictions for objects. Since temporary images can only be used with a single command buffer, and can not be used off the GPU, they generally should not be kept around past the completion of their associated command buffer. The lifetime of a temporary image is typically expected to be extremely short, perhaps spanning only a few lines of code. To keep the lifetime of the underlying texture allocation as short as possible, the texture is not allocated until the first time the object is used by an object or until the first time the property is read. The property serves to limit the lifetime of the texture on deallocation. You may use the property with the methods of an subclass, if and the texture conforms to the requirements of the given kernel. In such cases, the property is not modified, since the enclosing object is not available. There is no locking mechanism provided to prevent a object returned from the property from becoming invalid when the value of the property reaches 0. objects can otherwise be used wherever objects are used.


// A texture for use in convolutional neural networks that stores transient data to be used and discarded promptly.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TemporaryImage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TemporaryImage */

// Low-level interface for creating a temporary image using a texture descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097543-temporaryimagewithcommandbuffer
func (tc _TemporaryImageClass) TemporaryImageWithCommandBufferTextureDescriptor(commandBuffer unsafe.Pointer, textureDescriptor TextureDescriptor /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryImageWithCommandBuffer:textureDescriptor:"), commandBuffer, textureDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryImageWithCommandBufferTextureDescriptor) */


// A method that helps the framework decide which allocations to make ahead of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097544-prefetchstorage
func (tc _TemporaryImageClass) PrefetchStorage() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorage"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrefetchStorage) */


// A method that helps the framework decide which allocations to make ahead of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097544-prefetchstoragewithcommandbuffer
func (tc _TemporaryImageClass) PrefetchStorageWithCommandBufferImageDescriptorList(commandBuffer unsafe.Pointer, descriptorList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("prefetchStorageWithCommandBuffer:imageDescriptorList:"), commandBuffer, descriptorList)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrefetchStorageWithCommandBufferImageDescriptorList) */


// Initializes a temporary image for use on a command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097545-temporaryimagewithcommandbuffer
func (tc _TemporaryImageClass) TemporaryImageWithCommandBufferImageDescriptor(commandBuffer unsafe.Pointer, imageDescriptor IImageDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryImageWithCommandBuffer:imageDescriptor:"), commandBuffer, imageDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryImageWithCommandBufferImageDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2867130-defaultallocator
func (tc _TemporaryImageClass) DefaultAllocator() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("defaultAllocator"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultAllocator) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2942489-temporaryimagewithcommandbuffer
func (tc _TemporaryImageClass) TemporaryImageWithCommandBufferTextureDescriptorFeatureChannels(commandBuffer unsafe.Pointer, textureDescriptor TextureDescriptor /* not a class type */, featureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("temporaryImageWithCommandBuffer:textureDescriptor:featureChannels:"), commandBuffer, textureDescriptor, featureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TemporaryImageWithCommandBufferTextureDescriptorFeatureChannels) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TemporaryImage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TemporaryImage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TemporaryImage */

// The number of times a temporary image may be read by a CNN kernel before its contents become undefined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097546-readcount
func (t_ TemporaryImage) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("readCount"))
	return rv
}/* debug [instance_properties/getter]: readCount */


// The number of times a temporary image may be read by a CNN kernel before its contents become undefined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097546-readcount
func (t_ TemporaryImage) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadCount:"), value)
}/* debug [instance_properties/setter]: readCount */


// The underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/texture
func (t_ TemporaryImage) Texture() Texture /* not a class type */ {
	rv := objc.Send[Texture](t_.ID, objc.Sel("texture"))
	return rv
}/* debug [instance_properties/getter]: texture */


// The underlying texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/texture
func (t_ TemporaryImage) SetTexture(value Texture /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTexture:"), value)
}/* debug [instance_properties/setter]: texture */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSTemporaryImage */



