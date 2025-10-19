// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKTileMapNode] class.
var (
	sKTileMapNodeClass     _SKTileMapNodeClass
	sKTileMapNodeClassOnce sync.Once
)

func getSKTileMapNodeClass() _SKTileMapNodeClass {
	sKTileMapNodeClassOnce.Do(func() {
		sKTileMapNodeClass = _SKTileMapNodeClass{objc.GetClass("SKTileMapNode")}
	})
	return sKTileMapNodeClass
}

type _SKTileMapNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKTileMapNode] class.
type ISKTileMapNode interface {
	ISKNode
	CenterOfTileAtColumnRow(column uint, row uint) unsafe.Pointer
	FillWithTileGroup(tileGroup unsafe.Pointer)
	SetTileGroupAndTileDefinitionForColumnRow(tileGroup unsafe.Pointer, tileDefinition unsafe.Pointer, column uint, row uint)
	SetValueForAttributeNamed(value unsafe.Pointer, key string)
	TileColumnIndexFromPosition(position unsafe.Pointer) uint
	TileDefinitionAtColumnRow(column uint, row uint) unsafe.Pointer
	TileGroupAtColumnRow(column uint, row uint) unsafe.Pointer
	TileRowIndexFromPosition(position unsafe.Pointer) uint
	ValueForAttributeNamed(key string) unsafe.Pointer
}

// A two-dimensional array of images.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode
type SKTileMapNode struct {
	SKNode
}

// SKTileMapNodeFrom constructs a [SKTileMapNode] from an unsafe.Pointer.
//
// A two-dimensional array of images.
func SKTileMapNodeFrom(ptr unsafe.Pointer) SKTileMapNode {
	return SKTileMapNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTileMapNodeClass) Alloc() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTileMapNodeClass) New() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTileMapNode) Init() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTileMapNode) Autorelease() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileMapNode creates a new SKTileMapNode instance.
func NewSKTileMapNode() SKTileMapNode {
	return getSKTileMapNodeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/centerOfTile(atColumn:row:)
func (s_ SKTileMapNode) CenterOfTileAtColumnRow(column uint, row uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("centerOfTileAtColumn:row:"), column, row)
	return rv
}
// When creating a tile map node programmatically, this function performs a fill operation with the specified tile group.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/fill(with:)
func (s_ SKTileMapNode) FillWithTileGroup(tileGroup unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("fillWithTileGroup:"), tileGroup)
}
// Set the tile group and tile definition at the specified tile index.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/setTileGroup(_:andTileDefinition:forColumn:row:)
func (s_ SKTileMapNode) SetTileGroupAndTileDefinitionForColumnRow(tileGroup unsafe.Pointer, tileDefinition unsafe.Pointer, column uint, row uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTileGroup:andTileDefinition:forColumn:row:"), tileGroup, tileDefinition, column, row)
}
// Sets an attribute value for an attached shader.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/setValue(_:forAttribute:)
func (s_ SKTileMapNode) SetValueForAttributeNamed(value unsafe.Pointer, key string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileColumnIndex(fromPosition:)
func (s_ SKTileMapNode) TileColumnIndexFromPosition(position unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("tileColumnIndexFromPosition:"), position)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileDefinition(atColumn:row:)
func (s_ SKTileMapNode) TileDefinitionAtColumnRow(column uint, row uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("tileDefinitionAtColumn:row:"), column, row)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileGroup(atColumn:row:)
func (s_ SKTileMapNode) TileGroupAtColumnRow(column uint, row uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("tileGroupAtColumn:row:"), column, row)
	return rv
}
// Returns the tile map node object’s tile row index for the specified position in points.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileRowIndex(fromPosition:)
func (s_ SKTileMapNode) TileRowIndexFromPosition(position unsafe.Pointer) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("tileRowIndexFromPosition:"), position)
	return rv
}
// The value of a shader attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/value(forAttributeNamed:)
func (s_ SKTileMapNode) ValueForAttributeNamed(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("valueForAttributeNamed:"), objc.String(key))
	return rv
}


