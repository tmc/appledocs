// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [TextureViewDescriptor] class.
type ITextureViewDescriptor interface {
	objectivec.IObject
	

	// properties:
	LevelRange() foundation.Range
	SetLevelRange(value foundation.Range)
	PixelFormat() PixelFormat
	SetPixelFormat(value PixelFormat)
	SliceRange() foundation.Range
	SetSliceRange(value foundation.Range)
	Swizzle() MTLTextureSwizzleChannels
	SetSwizzle(value MTLTextureSwizzleChannels)
	TextureType() TextureType
	SetTextureType(value TextureType)


	

	// methods:


}





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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor
type TextureViewDescriptor struct {
	objectivec.Object
}

// TextureViewDescriptorFrom constructs a [TextureViewDescriptor] from an unsafe.Pointer.
func TextureViewDescriptorFrom(ptr unsafe.Pointer) TextureViewDescriptor {
	return TextureViewDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/levelRange-7e7f3
func (t_ TextureViewDescriptor) LevelRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("levelRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/levelRange-7e7f3
func (t_ TextureViewDescriptor) SetLevelRange(value foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLevelRange:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/pixelFormat
func (t_ TextureViewDescriptor) PixelFormat() PixelFormat {
	rv := objc.Send[PixelFormat](t_.ID, objc.Sel("pixelFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/pixelFormat
func (t_ TextureViewDescriptor) SetPixelFormat(value PixelFormat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPixelFormat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/sliceRange-3cs9b
func (t_ TextureViewDescriptor) SliceRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("sliceRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/sliceRange-3cs9b
func (t_ TextureViewDescriptor) SetSliceRange(value foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSliceRange:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/swizzle
func (t_ TextureViewDescriptor) Swizzle() MTLTextureSwizzleChannels {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("swizzle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/swizzle
func (t_ TextureViewDescriptor) SetSwizzle(value MTLTextureSwizzleChannels) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSwizzle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t_ TextureViewDescriptor) TextureType() TextureType {
	rv := objc.Send[TextureType](t_.ID, objc.Sel("textureType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t_ TextureViewDescriptor) SetTextureType(value TextureType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextureType:"), value)
}








