// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTexture] class.
var sKTextureClass = _SKTextureClass{objc.GetClass("SKTexture")}

type _SKTextureClass struct {
	class objc.Class
}

// An interface definition for the [SKTexture] class.
type ISKTexture interface {
	objectivec.IObject
	CGImage() unsafe.Pointer
	PreloadWithCompletionHandler(completionHandler unsafe.Pointer)
	Size() unsafe.Pointer
	TextureRect() unsafe.Pointer
}

// An image, decoded on the GPU, that can be used to render various SpriteKit objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture

type SKTexture struct {
	objectivec.Object
}

// SKTextureFrom constructs a [SKTexture] from an unsafe.Pointer.
//
// An image, decoded on the GPU, that can be used to render various SpriteKit objects.
func SKTextureFrom(ptr unsafe.Pointer) SKTexture {
	return SKTexture{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKTextureClass) Alloc() SKTexture {
	rv := objc.Send[SKTexture](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKTextureClass) New() SKTexture {
	rv := objc.Send[SKTexture](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTexture) Init() SKTexture {
	rv := objc.Send[SKTexture](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTexture) Autorelease() SKTexture {
	rv := objc.Send[SKTexture](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTexture creates a new SKTexture instance.
func NewSKTexture() SKTexture {
	return sKTextureClass.New()
}


// Creates a texture from the specified noise map. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/init(noiseMap:)
func NewTextureWithNoiseMap(noiseMap unsafe.Pointer) SKTexture {
	rv := objc.Send[SKTexture](objc.ID(sKTextureClass.class), objc.Sel("textureWithNoiseMap:"), noiseMap)
	rv.Autorelease()
	return rv
}


// Creates a texture from the specified noise map. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/init(noiseMap:)
func (sc _SKTextureClass) TextureWithNoiseMap(noiseMap unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("textureWithNoiseMap:"), noiseMap)
	return rv
}
// Load the data of multiple textures into memory. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/preload(_:withCompletionHandler:)
func (sc _SKTextureClass) PreloadTexturesWithCompletionHandler(textures unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("preloadTextures:withCompletionHandler:"), textures, completionHandler)
}
// Returns the texture’s image data as a Quartz 2D image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/cgImage()
func (s_ SKTexture) CGImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("CGImage"))
	return rv
}
// Load texture data into memory, calling a completion handler after the task completes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/preload(completionHandler:)
func (s_ SKTexture) PreloadWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("preloadWithCompletionHandler:"), completionHandler)
}
// Gets the size of the texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/size()
func (s_ SKTexture) Size() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("size"))
	return rv
}
// Gets a rectangle that defines the portion of the texture used to render its image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTexture/textureRect()
func (s_ SKTexture) TextureRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("textureRect"))
	return rv
}

