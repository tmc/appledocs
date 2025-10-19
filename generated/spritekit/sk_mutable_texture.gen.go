// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKMutableTexture] class.
var sKMutableTextureClass = _SKMutableTextureClass{objc.GetClass("SKMutableTexture")}

type _SKMutableTextureClass struct {
	class objc.Class
}

// An interface definition for the [SKMutableTexture] class.
type ISKMutableTexture interface {
	ISKTexture
	ModifyPixelDataWithBlock(block unsafe.Pointer)
}

// A texture whose contents can be dynamically updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture

type SKMutableTexture struct {
	SKTexture
}

// SKMutableTextureFrom constructs a [SKMutableTexture] from an unsafe.Pointer.
//
// A texture whose contents can be dynamically updated.
func SKMutableTextureFrom(ptr unsafe.Pointer) SKMutableTexture {
	return SKMutableTexture{
		SKTexture: SKTextureFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKMutableTextureClass) Alloc() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKMutableTextureClass) New() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKMutableTexture) Init() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKMutableTexture) Autorelease() SKMutableTexture {
	rv := objc.Send[SKMutableTexture](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKMutableTexture creates a new SKMutableTexture instance.
func NewSKMutableTexture() SKMutableTexture {
	return sKMutableTextureClass.New()
}


// Initializes an empty texture with a specific size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/init(size:)
func NewSKMutableTextureWithSize(size unsafe.Pointer) SKMutableTexture {
	instance := sKMutableTextureClass.Alloc()
	rv := objc.Send[SKMutableTexture](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}
// Initializes an empty texture with a specific size and format. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/init(size:pixelFormat:)
func NewSKMutableTextureWithSizePixelFormat(size unsafe.Pointer, format int) SKMutableTexture {
	instance := sKMutableTextureClass.Alloc()
	rv := objc.Send[SKMutableTexture](instance.ID, objc.Sel("initWithSize:pixelFormat:"), size, format)
	rv.Autorelease()
	return rv
}


// Creates an empty texture with a specific size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/mutableTextureWithSize:
func (sc _SKMutableTextureClass) MutableTextureWithSize(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("mutableTextureWithSize:"), size)
	return rv
}
// Modifies the contents of a mutable texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKMutableTexture/modifyPixelData(_:)
func (s_ SKMutableTexture) ModifyPixelDataWithBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("modifyPixelDataWithBlock:"), block)
}

