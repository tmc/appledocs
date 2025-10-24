// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSSVGFDefaultTextureAllocator */


/* debug [class_header]: Header for MPSSVGFDefaultTextureAllocator */
// The class instance for the [SVGFDefaultTextureAllocator] class.
var (
	SVGFDefaultTextureAllocatorClass     _SVGFDefaultTextureAllocatorClass
	SVGFDefaultTextureAllocatorClassOnce sync.Once
)

func getSVGFDefaultTextureAllocatorClass() _SVGFDefaultTextureAllocatorClass {
	SVGFDefaultTextureAllocatorClassOnce.Do(func() {
		SVGFDefaultTextureAllocatorClass = _SVGFDefaultTextureAllocatorClass{objc.GetClass("MPSSVGFDefaultTextureAllocator")}
	})
	return SVGFDefaultTextureAllocatorClass
}

type _SVGFDefaultTextureAllocatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SVGFDefaultTextureAllocator */
// An interface definition for the [SVGFDefaultTextureAllocator] class.
type ISVGFDefaultTextureAllocator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SVGFDefaultTextureAllocator */
	// properties:
	AllocatedTextureCount() objectivec.IObject
	SetAllocatedTextureCount(value objectivec.IObject)
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SVGFDefaultTextureAllocator */
	// methods:
	Reset()
	`return`()
	ReturnTexture(texture unsafe.Pointer)
	Texture()
	TextureWithPixelFormatWidthHeight(pixelFormat PixelFormat /* not a class type */, width uint, height uint) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SVGFDefaultTextureAllocator */
// Alloc allocates a new instance without initialization.
func (sc _SVGFDefaultTextureAllocatorClass) Alloc() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SVGFDefaultTextureAllocatorClass) New() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SVGFDefaultTextureAllocator) Init() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SVGFDefaultTextureAllocator) Autorelease() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSVGFDefaultTextureAllocator creates a new SVGFDefaultTextureAllocator instance.
func NewSVGFDefaultTextureAllocator() SVGFDefaultTextureAllocator {
	return getSVGFDefaultTextureAllocatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SVGFDefaultTextureAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator
type SVGFDefaultTextureAllocator struct {
	objectivec.Object
}

// SVGFDefaultTextureAllocatorFrom constructs a [SVGFDefaultTextureAllocator] from an unsafe.Pointer.
func SVGFDefaultTextureAllocatorFrom(ptr unsafe.Pointer) SVGFDefaultTextureAllocator {
	return SVGFDefaultTextureAllocator{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SVGFDefaultTextureAllocator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242897-initwithdevice
func NewSVGFDefaultTextureAllocatorWithDevice(device unsafe.Pointer) SVGFDefaultTextureAllocator {
	instance := getSVGFDefaultTextureAllocatorClass().Alloc()
	rv := objc.Send[SVGFDefaultTextureAllocator](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSVGFDefaultTextureAllocatorWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SVGFDefaultTextureAllocator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SVGFDefaultTextureAllocator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SVGFDefaultTextureAllocator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242898-reset
func (s_ SVGFDefaultTextureAllocator) Reset() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242899-return
func (s_ SVGFDefaultTextureAllocator) `return`() {
	objc.Send[objc.ID](s_.ID, objc.Sel("`return`"))
}/* debug [instance_methods/method]: `return` */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242899-returntexture
func (s_ SVGFDefaultTextureAllocator) ReturnTexture(texture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("returnTexture:"), texture)
}/* debug [instance_methods/method]: ReturnTexture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242900-texture
func (s_ SVGFDefaultTextureAllocator) Texture() {
	objc.Send[objc.ID](s_.ID, objc.Sel("texture"))
}/* debug [instance_methods/method]: Texture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242900-texturewithpixelformat
func (s_ SVGFDefaultTextureAllocator) TextureWithPixelFormatWidthHeight(pixelFormat PixelFormat /* not a class type */, width uint, height uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("textureWithPixelFormat:width:height:"), pixelFormat, width, height)
	return rv
}/* debug [instance_methods/method]: TextureWithPixelFormatWidthHeight */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SVGFDefaultTextureAllocator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242895-allocatedtexturecount
func (s_ SVGFDefaultTextureAllocator) AllocatedTextureCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("allocatedTextureCount"))
	return rv
}/* debug [instance_properties/getter]: allocatedTextureCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242895-allocatedtexturecount
func (s_ SVGFDefaultTextureAllocator) SetAllocatedTextureCount(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllocatedTextureCount:"), value)
}/* debug [instance_properties/setter]: allocatedTextureCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242896-device
func (s_ SVGFDefaultTextureAllocator) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242896-device
func (s_ SVGFDefaultTextureAllocator) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSSVGFDefaultTextureAllocator */


