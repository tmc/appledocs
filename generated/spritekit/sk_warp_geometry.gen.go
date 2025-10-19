// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKWarpGeometry] class.
var (
	sKWarpGeometryClass     _SKWarpGeometryClass
	sKWarpGeometryClassOnce sync.Once
)

func getSKWarpGeometryClass() _SKWarpGeometryClass {
	sKWarpGeometryClassOnce.Do(func() {
		sKWarpGeometryClass = _SKWarpGeometryClass{objc.GetClass("SKWarpGeometry")}
	})
	return sKWarpGeometryClass
}

type _SKWarpGeometryClass struct {
	class objc.Class
}

// An interface definition for the [SKWarpGeometry] class.
type ISKWarpGeometry interface {
	objectivec.IObject
}

// A definition for a deformation of nodes that conform to . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometry
type SKWarpGeometry struct {
	objectivec.Object
}

// SKWarpGeometryFrom constructs a [SKWarpGeometry] from an unsafe.Pointer.
//
// A definition for a deformation of nodes that conform to .
func SKWarpGeometryFrom(ptr unsafe.Pointer) SKWarpGeometry {
	return SKWarpGeometry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKWarpGeometryClass) Alloc() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKWarpGeometryClass) New() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKWarpGeometry) Init() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKWarpGeometry) Autorelease() SKWarpGeometry {
	rv := objc.Send[SKWarpGeometry](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKWarpGeometry creates a new SKWarpGeometry instance.
func NewSKWarpGeometry() SKWarpGeometry {
	return getSKWarpGeometryClass().New()
}




