// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTextureAtlas] class.
var sKTextureAtlasClass = _SKTextureAtlasClass{objc.GetClass("SKTextureAtlas")}

type _SKTextureAtlasClass struct {
	class objc.Class
}

// An interface definition for the [SKTextureAtlas] class.
type ISKTextureAtlas interface {
	objectivec.IObject
	PreloadWithCompletionHandler(completionHandler unsafe.Pointer)
	TextureNamed(name string) unsafe.Pointer
}

// A collection of textures optimized for storage and drawing performance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas

type SKTextureAtlas struct {
	objectivec.Object
}

// SKTextureAtlasFrom constructs a [SKTextureAtlas] from an unsafe.Pointer.
//
// A collection of textures optimized for storage and drawing performance.
func SKTextureAtlasFrom(ptr unsafe.Pointer) SKTextureAtlas {
	return SKTextureAtlas{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKTextureAtlasClass) Alloc() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKTextureAtlasClass) New() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTextureAtlas) Init() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTextureAtlas) Autorelease() SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTextureAtlas creates a new SKTextureAtlas instance.
func NewSKTextureAtlas() SKTextureAtlas {
	return sKTextureAtlasClass.New()
}


// Creates a texture atlas from data stored in the app bundle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/init(named:)
func NewAtlasNamed(name string) SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](objc.ID(sKTextureAtlasClass.class), objc.Sel("atlasNamed:"), name)
	rv.Autorelease()
	return rv
}
// Creates a texture atlas from a set of image files. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/init(dictionary:)
func NewAtlasWithDictionary(properties unsafe.Pointer) SKTextureAtlas {
	rv := objc.Send[SKTextureAtlas](objc.ID(sKTextureAtlasClass.class), objc.Sel("atlasWithDictionary:"), properties)
	rv.Autorelease()
	return rv
}


// Creates a texture atlas from a set of image files. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/init(dictionary:)
func (sc _SKTextureAtlasClass) AtlasWithDictionary(properties unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("atlasWithDictionary:"), properties)
	return rv
}
// Creates a texture atlas from data stored in the app bundle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/init(named:)
func (sc _SKTextureAtlasClass) AtlasNamed(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("atlasNamed:"), name)
	return rv
}
// Loads the textures of multiple atlas objects into memory, calling a completion handler after the task completes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/preloadTextureAtlases(_:withCompletionHandler:)
func (sc _SKTextureAtlasClass) PreloadTextureAtlasesWithCompletionHandler(textureAtlases unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("preloadTextureAtlases:withCompletionHandler:"), textureAtlases, completionHandler)
}
// Loads the textures of multiple atlases into memory, calling a completion handler after the task completes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/preloadTextureAtlasesNamed(_:withCompletionHandler:)
func (sc _SKTextureAtlasClass) PreloadTextureAtlasesNamedWithCompletionHandler(atlasNames unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("preloadTextureAtlasesNamed:withCompletionHandler:"), atlasNames, completionHandler)
}
// Loads an atlas object’s textures into memory, calling a completion handler after the task completes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/preload(completionHandler:)
func (s_ SKTextureAtlas) PreloadWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("preloadWithCompletionHandler:"), completionHandler)
}
// Creates a texture from data stored in the texture atlas. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTextureAtlas/textureNamed(_:)
func (s_ SKTextureAtlas) TextureNamed(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("textureNamed:"), name)
	return rv
}

