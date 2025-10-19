// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTileGroup] class.
var (
	sKTileGroupClass     _SKTileGroupClass
	sKTileGroupClassOnce sync.Once
)

func getSKTileGroupClass() _SKTileGroupClass {
	sKTileGroupClassOnce.Do(func() {
		sKTileGroupClass = _SKTileGroupClass{objc.GetClass("SKTileGroup")}
	})
	return sKTileGroupClass
}

type _SKTileGroupClass struct {
	class objc.Class
}

// An interface definition for the [SKTileGroup] class.
type ISKTileGroup interface {
	objectivec.IObject
}

// A set of tiles that collectively define one type of terrain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroup
type SKTileGroup struct {
	objectivec.Object
}

// SKTileGroupFrom constructs a [SKTileGroup] from an unsafe.Pointer.
//
// A set of tiles that collectively define one type of terrain.
func SKTileGroupFrom(ptr unsafe.Pointer) SKTileGroup {
	return SKTileGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTileGroupClass) Alloc() SKTileGroup {
	rv := objc.Send[SKTileGroup](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTileGroupClass) New() SKTileGroup {
	rv := objc.Send[SKTileGroup](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTileGroup) Init() SKTileGroup {
	rv := objc.Send[SKTileGroup](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTileGroup) Autorelease() SKTileGroup {
	rv := objc.Send[SKTileGroup](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileGroup creates a new SKTileGroup instance.
func NewSKTileGroup() SKTileGroup {
	return getSKTileGroupClass().New()
}


// Creates and initializes a tile group with the specified tile group rules. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/init(rules:)
func NewSKTileGroupWithRules(rules unsafe.Pointer) SKTileGroup {
	instance := getSKTileGroupClass().Alloc()
	rv := objc.Send[SKTileGroup](instance.ID, objc.Sel("initWithRules:"), rules)
	rv.Autorelease()
	return rv
}
// Creates and initializes a simple tile group with a single tile definition. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/init(tileDefinition:)
func NewSKTileGroupWithTileDefinition(tileDefinition unsafe.Pointer) SKTileGroup {
	instance := getSKTileGroupClass().Alloc()
	rv := objc.Send[SKTileGroup](instance.ID, objc.Sel("initWithTileDefinition:"), tileDefinition)
	rv.Autorelease()
	return rv
}


// Creates an empty tile that erases the existing tile at that location on a tile map. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/empty()
func (sc _SKTileGroupClass) EmptyTileGroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("emptyTileGroup"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/tileGroupWithRules:
func (sc _SKTileGroupClass) TileGroupWithRules(rules unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileGroupWithRules:"), rules)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/tileGroupWithTileDefinition:
func (sc _SKTileGroupClass) TileGroupWithTileDefinition(tileDefinition unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileGroupWithTileDefinition:"), tileDefinition)
	return rv
}

