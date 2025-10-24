// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLTextureViewDescriptor */


/* debug [class_header]: Header for MTLTextureViewDescriptor */
// The class instance for the [TextureViewDescriptor] class.
var (
	TextureViewDescriptorClass     _TextureViewDescriptorClass
	TextureViewDescriptorClassOnce sync.Once
)

func getTextureViewDescriptorClass() _TextureViewDescriptorClass {
	TextureViewDescriptorClassOnce.Do(func() {
		TextureViewDescriptorClass = _TextureViewDescriptorClass{objc.GetClass("MTLTextureViewDescriptor")}
	})
	return TextureViewDescriptorClass
}

type _TextureViewDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextureViewDescriptor */
// An interface definition for the [TextureViewDescriptor] class.
type ITextureViewDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextureViewDescriptor */
	// properties:
	LevelRange() corefoundation.Range
	SetLevelRange(value corefoundation.Range)
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
	SliceRange() corefoundation.Range
	SetSliceRange(value corefoundation.Range)
	Swizzle() objc.IObject /* cross-framework: MTLTextureSwizzleChannels */
	SetSwizzle(value objc.IObject /* cross-framework: MTLTextureSwizzleChannels */)
	TextureType() TextureType
	SetTextureType(value TextureType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextureViewDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextureViewDescriptor */
// Alloc allocates a new instance without initialization.
func (tc _TextureViewDescriptorClass) Alloc() TextureViewDescriptor {
	rv := objc.Send[TextureViewDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextureViewDescriptorClass) New() TextureViewDescriptor {
	rv := objc.Send[TextureViewDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextureViewDescriptor) Init() TextureViewDescriptor {
	rv := objc.Send[TextureViewDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextureViewDescriptor) Autorelease() TextureViewDescriptor {
	rv := objc.Send[TextureViewDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextureViewDescriptor creates a new TextureViewDescriptor instance.
func NewTextureViewDescriptor() TextureViewDescriptor {
	return getTextureViewDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextureViewDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor
type TextureViewDescriptor struct {
	objectivec.Object
}

// TextureViewDescriptorFrom constructs a [TextureViewDescriptor] from an unsafe.Pointer.
func TextureViewDescriptorFrom(ptr unsafe.Pointer) TextureViewDescriptor {
	return TextureViewDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextureViewDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextureViewDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextureViewDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextureViewDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextureViewDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/levelRange-7e7f3
func (t_ TextureViewDescriptor) LevelRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("levelRange"))
	return rv
}/* debug [instance_properties/getter]: levelRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/levelRange-7e7f3
func (t_ TextureViewDescriptor) SetLevelRange(value corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLevelRange:"), value)
}/* debug [instance_properties/setter]: levelRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/pixelFormat
func (t_ TextureViewDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](t_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/pixelFormat
func (t_ TextureViewDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/sliceRange-3cs9b
func (t_ TextureViewDescriptor) SliceRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("sliceRange"))
	return rv
}/* debug [instance_properties/getter]: sliceRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/sliceRange-3cs9b
func (t_ TextureViewDescriptor) SetSliceRange(value corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSliceRange:"), value)
}/* debug [instance_properties/setter]: sliceRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/swizzle
func (t_ TextureViewDescriptor) Swizzle() objc.IObject /* cross-framework: MTLTextureSwizzleChannels */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("swizzle"))
	return rv
}/* debug [instance_properties/getter]: swizzle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/swizzle
func (t_ TextureViewDescriptor) SetSwizzle(value objc.IObject /* cross-framework: MTLTextureSwizzleChannels */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSwizzle:"), value)
}/* debug [instance_properties/setter]: swizzle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t_ TextureViewDescriptor) TextureType() TextureType {
	rv := objc.Send[TextureType](t_.ID, objc.Sel("textureType"))
	return rv
}/* debug [instance_properties/getter]: textureType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t_ TextureViewDescriptor) SetTextureType(value TextureType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextureType:"), value)
}/* debug [instance_properties/setter]: textureType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTextureViewDescriptor */



