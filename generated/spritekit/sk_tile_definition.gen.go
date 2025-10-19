// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTileDefinition] class.
var (
	sKTileDefinitionClass     _SKTileDefinitionClass
	sKTileDefinitionClassOnce sync.Once
)

func getSKTileDefinitionClass() _SKTileDefinitionClass {
	sKTileDefinitionClassOnce.Do(func() {
		sKTileDefinitionClass = _SKTileDefinitionClass{objc.GetClass("SKTileDefinition")}
	})
	return sKTileDefinitionClass
}

type _SKTileDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [SKTileDefinition] class.
type ISKTileDefinition interface {
	objectivec.IObject
}

// A single tile that can be repeated in a tile map. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition
type SKTileDefinition struct {
	objectivec.Object
}

// SKTileDefinitionFrom constructs a [SKTileDefinition] from an unsafe.Pointer.
//
// A single tile that can be repeated in a tile map.
func SKTileDefinitionFrom(ptr unsafe.Pointer) SKTileDefinition {
	return SKTileDefinition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTileDefinitionClass) Alloc() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTileDefinitionClass) New() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTileDefinition) Init() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTileDefinition) Autorelease() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileDefinition creates a new SKTileDefinition instance.
func NewSKTileDefinition() SKTileDefinition {
	return getSKTileDefinitionClass().New()
}


// Initializes a new tile definition with a single texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:)
func NewSKTileDefinitionWithTexture(texture unsafe.Pointer) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[SKTileDefinition](instance.ID, objc.Sel("initWithTexture:"), texture)
	rv.Autorelease()
	return rv
}
// Initializes a new tile definition with a single texture and separate normal texture for simulating 3D lighting. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:normalTexture:size:)
func NewSKTileDefinitionWithTextureNormalTextureSize(texture unsafe.Pointer, normalTexture unsafe.Pointer, size unsafe.Pointer) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[SKTileDefinition](instance.ID, objc.Sel("initWithTexture:normalTexture:size:"), texture, normalTexture, size)
	rv.Autorelease()
	return rv
}
// Initializes a new tile definition of a specified size with a single texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:size:)
func NewSKTileDefinitionWithTextureSize(texture unsafe.Pointer, size unsafe.Pointer) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[SKTileDefinition](instance.ID, objc.Sel("initWithTexture:size:"), texture, size)
	rv.Autorelease()
	return rv
}
// Initializes a new tile definition with arrays of textures and normal textures for animation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(textures:normalTextures:size:timePerFrame:)
func NewSKTileDefinitionWithTexturesNormalTexturesSizeTimePerFrame(textures unsafe.Pointer, normalTextures unsafe.Pointer, size unsafe.Pointer, timePerFrame float64) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[SKTileDefinition](instance.ID, objc.Sel("initWithTextures:normalTextures:size:timePerFrame:"), textures, normalTextures, size, timePerFrame)
	rv.Autorelease()
	return rv
}
// Initializes a new tile definition with an array of textures for animation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(textures:size:timePerFrame:)
func NewSKTileDefinitionWithTexturesSizeTimePerFrame(textures unsafe.Pointer, size unsafe.Pointer, timePerFrame float64) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[SKTileDefinition](instance.ID, objc.Sel("initWithTextures:size:timePerFrame:"), textures, size, timePerFrame)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/tileDefinitionWithTexture:
func (sc _SKTileDefinitionClass) TileDefinitionWithTexture(texture unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileDefinitionWithTexture:"), texture)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/tileDefinitionWithTexture:normalTexture:size:
func (sc _SKTileDefinitionClass) TileDefinitionWithTextureNormalTextureSize(texture unsafe.Pointer, normalTexture unsafe.Pointer, size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileDefinitionWithTexture:normalTexture:size:"), texture, normalTexture, size)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/tileDefinitionWithTexture:size:
func (sc _SKTileDefinitionClass) TileDefinitionWithTextureSize(texture unsafe.Pointer, size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileDefinitionWithTexture:size:"), texture, size)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/tileDefinitionWithTextures:normalTextures:size:timePerFrame:
func (sc _SKTileDefinitionClass) TileDefinitionWithTexturesNormalTexturesSizeTimePerFrame(textures unsafe.Pointer, normalTextures unsafe.Pointer, size unsafe.Pointer, timePerFrame float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileDefinitionWithTextures:normalTextures:size:timePerFrame:"), textures, normalTextures, size, timePerFrame)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/tileDefinitionWithTextures:size:timePerFrame:
func (sc _SKTileDefinitionClass) TileDefinitionWithTexturesSizeTimePerFrame(textures unsafe.Pointer, size unsafe.Pointer, timePerFrame float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileDefinitionWithTextures:size:timePerFrame:"), textures, size, timePerFrame)
	return rv
}

