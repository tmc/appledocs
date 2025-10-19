// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTileSet] class.
var (
	sKTileSetClass     _SKTileSetClass
	sKTileSetClassOnce sync.Once
)

func getSKTileSetClass() _SKTileSetClass {
	sKTileSetClassOnce.Do(func() {
		sKTileSetClass = _SKTileSetClass{objc.GetClass("SKTileSet")}
	})
	return sKTileSetClass
}

type _SKTileSetClass struct {
	class objc.Class
}

// An interface definition for the [SKTileSet] class.
type ISKTileSet interface {
	objectivec.IObject
}

// A container for related tile groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet
type SKTileSet struct {
	objectivec.Object
}

// SKTileSetFrom constructs a [SKTileSet] from an unsafe.Pointer.
//
// A container for related tile groups.
func SKTileSetFrom(ptr unsafe.Pointer) SKTileSet {
	return SKTileSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTileSetClass) Alloc() SKTileSet {
	rv := objc.Send[SKTileSet](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTileSetClass) New() SKTileSet {
	rv := objc.Send[SKTileSet](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTileSet) Init() SKTileSet {
	rv := objc.Send[SKTileSet](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTileSet) Autorelease() SKTileSet {
	rv := objc.Send[SKTileSet](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileSet creates a new SKTileSet instance.
func NewSKTileSet() SKTileSet {
	return getSKTileSetClass().New()
}


// Initializes a tile set from a URL to an archived .sks file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(from:)
func NewSKTileSetFromURL(url unsafe.Pointer) SKTileSet {
	rv := objc.Send[SKTileSet](objc.ID(getSKTileSetClass().class), objc.Sel("tileSetFromURL:"), url)
	return rv
}
// Initializes a tile set by searching the app bundle for an archived file by name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(named:)
func NewSKTileSetNamed(name string) SKTileSet {
	rv := objc.Send[SKTileSet](objc.ID(getSKTileSetClass().class), objc.Sel("tileSetNamed:"), objc.String(name))
	return rv
}
// Initializes a new tile set with an array of tile groups and rectangular grid layout. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(tileGroups:)
func NewSKTileSetWithTileGroups(tileGroups unsafe.Pointer) SKTileSet {
	instance := getSKTileSetClass().Alloc()
	rv := objc.Send[SKTileSet](instance.ID, objc.Sel("initWithTileGroups:"), tileGroups)
	rv.Autorelease()
	return rv
}
// Initializes a new tile set with an array of tile groups and specified layout. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(tileGroups:tileSetType:)
func NewSKTileSetWithTileGroupsTileSetType(tileGroups unsafe.Pointer, tileSetType unsafe.Pointer) SKTileSet {
	instance := getSKTileSetClass().Alloc()
	rv := objc.Send[SKTileSet](instance.ID, objc.Sel("initWithTileGroups:tileSetType:"), tileGroups, tileSetType)
	rv.Autorelease()
	return rv
}


// Initializes a tile set from a URL to an archived .sks file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(from:)
func (sc _SKTileSetClass) TileSetFromURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileSetFromURL:"), url)
	return rv
}
// Initializes a tile set by searching the app bundle for an archived file by name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(named:)
func (sc _SKTileSetClass) TileSetNamed(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileSetNamed:"), objc.String(name))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/tileSetWithTileGroups:
func (sc _SKTileSetClass) TileSetWithTileGroups(tileGroups unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileSetWithTileGroups:"), tileGroups)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileSet/tileSetWithTileGroups:tileSetType:
func (sc _SKTileSetClass) TileSetWithTileGroupsTileSetType(tileGroups unsafe.Pointer, tileSetType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileSetWithTileGroups:tileSetType:"), tileGroups, tileSetType)
	return rv
}

