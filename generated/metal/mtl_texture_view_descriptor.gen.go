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
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor
type TextureViewDescriptor struct {
	objectivec.Object
}

// TextureViewDescriptorFrom constructs a [TextureViewDescriptor] from an unsafe.Pointer.
func TextureViewDescriptorFrom(ptr unsafe.Pointer) TextureViewDescriptor {
	return TextureViewDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextureViewDescriptorClass) Alloc() TextureViewDescriptor {
	rv := objc.Send[TextureViewDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t_ TextureViewDescriptor) TextureType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("textureType"))
	return rv
}


// SetTextureType sets the value of the textureType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureViewDescriptor/textureType
func (t_ TextureViewDescriptor) SetTextureType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextureType:"), value)
}


