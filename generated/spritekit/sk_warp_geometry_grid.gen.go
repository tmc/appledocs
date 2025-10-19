// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKWarpGeometryGrid] class.
var (
	sKWarpGeometryGridClass     _SKWarpGeometryGridClass
	sKWarpGeometryGridClassOnce sync.Once
)

func getSKWarpGeometryGridClass() _SKWarpGeometryGridClass {
	sKWarpGeometryGridClassOnce.Do(func() {
		sKWarpGeometryGridClass = _SKWarpGeometryGridClass{objc.GetClass("SKWarpGeometryGrid")}
	})
	return sKWarpGeometryGridClass
}

type _SKWarpGeometryGridClass struct {
	class objc.Class
}

// An interface definition for the [SKWarpGeometryGrid] class.
type ISKWarpGeometryGrid interface {
	ISKWarpGeometry
	DestPositionAtIndex(index int) unsafe.Pointer
	GridByReplacingDestPositions(destPositions unsafe.Pointer) unsafe.Pointer
	GridByReplacingSourcePositions(sourcePositions unsafe.Pointer) unsafe.Pointer
	SourcePositionAtIndex(index int) unsafe.Pointer
}

// A definition for a grid-based deformation of nodes that conform to .
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid
type SKWarpGeometryGrid struct {
	SKWarpGeometry
}

// SKWarpGeometryGridFrom constructs a [SKWarpGeometryGrid] from an unsafe.Pointer.
//
// A definition for a grid-based deformation of nodes that conform to .
func SKWarpGeometryGridFrom(ptr unsafe.Pointer) SKWarpGeometryGrid {
	return SKWarpGeometryGrid{
		SKWarpGeometry: SKWarpGeometryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKWarpGeometryGridClass) Alloc() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKWarpGeometryGridClass) New() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKWarpGeometryGrid) Init() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKWarpGeometryGrid) Autorelease() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKWarpGeometryGrid creates a new SKWarpGeometryGrid instance.
func NewSKWarpGeometryGrid() SKWarpGeometryGrid {
	return getSKWarpGeometryGridClass().New()
}


// Tells you when to intialize a grid that was loaded from an archive.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/init(coder:)
func NewSKWarpGeometryGridWithCoder(aDecoder unsafe.Pointer) SKWarpGeometryGrid {
	instance := getSKWarpGeometryGridClass().Alloc()
	rv := objc.Send[SKWarpGeometryGrid](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}
// Creates a warp geometry grid of a specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/init(columns:rows:)
func NewSKWarpGeometryGridWithColumnsRows(cols int, rows int) SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](objc.ID(getSKWarpGeometryGridClass().class), objc.Sel("gridWithColumns:rows:"), cols, rows)
	return rv
}
// Creates a warp geometry grid of a specific size and warp translation, in pointers to point arrays.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/initWithColumns:rows:sourcePositions:destPositions:
func NewSKWarpGeometryGridWithColumnsRowsSourcePositionsDestPositions(cols int, rows int, sourcePositions unsafe.Pointer, destPositions unsafe.Pointer) SKWarpGeometryGrid {
	instance := getSKWarpGeometryGridClass().Alloc()
	rv := objc.Send[SKWarpGeometryGrid](instance.ID, objc.Sel("initWithColumns:rows:sourcePositions:destPositions:"), cols, rows, sourcePositions, destPositions)
	rv.Autorelease()
	return rv
}


// Initializes a new empty grid.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/grid
func (sc _SKWarpGeometryGridClass) Grid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("grid"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/gridWithColumns:rows:sourcePositions:destPositions:
func (sc _SKWarpGeometryGridClass) GridWithColumnsRowsSourcePositionsDestPositions(cols int, rows int, sourcePositions unsafe.Pointer, destPositions unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("gridWithColumns:rows:sourcePositions:destPositions:"), cols, rows, sourcePositions, destPositions)
	return rv
}
// Creates a warp geometry grid of a specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/init(columns:rows:)
func (sc _SKWarpGeometryGridClass) GridWithColumnsRows(cols int, rows int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("gridWithColumns:rows:"), cols, rows)
	return rv
}
// Returns the destination position of a vertex.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/destPosition(at:)
func (s_ SKWarpGeometryGrid) DestPositionAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("destPositionAtIndex:"), index)
	return rv
}
// Returns a copy of the receiver with the destination positions replaced by a specified array.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/gridByReplacingDestPositions:
func (s_ SKWarpGeometryGrid) GridByReplacingDestPositions(destPositions unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("gridByReplacingDestPositions:"), destPositions)
	return rv
}
// Returns a copy of the receiver with the source positions replaced by a specified array.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/gridByReplacingSourcePositions:
func (s_ SKWarpGeometryGrid) GridByReplacingSourcePositions(sourcePositions unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("gridByReplacingSourcePositions:"), sourcePositions)
	return rv
}
// Returns the source position of a vertex.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/sourcePosition(at:)
func (s_ SKWarpGeometryGrid) SourcePositionAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sourcePositionAtIndex:"), index)
	return rv
}

