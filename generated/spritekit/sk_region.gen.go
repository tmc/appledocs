// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [SKRegion] class.
var (
	sKRegionClass     _SKRegionClass
	sKRegionClassOnce sync.Once
)

func getSKRegionClass() _SKRegionClass {
	sKRegionClassOnce.Do(func() {
		sKRegionClass = _SKRegionClass{objc.GetClass("SKRegion")}
	})
	return sKRegionClass
}

type _SKRegionClass struct {
	class objc.Class
}

// An interface definition for the [SKRegion] class.
type ISKRegion interface {
	objectivec.IObject
	RegionByDifferenceFromRegion(region unsafe.Pointer) unsafe.Pointer
	RegionByIntersectionWithRegion(region unsafe.Pointer) unsafe.Pointer
	RegionByUnionWithRegion(region unsafe.Pointer) unsafe.Pointer
	ContainsPoint(point unsafe.Pointer) bool
	InverseRegion() unsafe.Pointer
}

// The definition of an arbitrary area.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion
type SKRegion struct {
	objectivec.Object
}

// SKRegionFrom constructs a [SKRegion] from an unsafe.Pointer.
//
// The definition of an arbitrary area.
func SKRegionFrom(ptr unsafe.Pointer) SKRegion {
	return SKRegion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKRegionClass) Alloc() SKRegion {
	rv := objc.Send[SKRegion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKRegionClass) New() SKRegion {
	rv := objc.Send[SKRegion](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKRegion) Init() SKRegion {
	rv := objc.Send[SKRegion](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKRegion) Autorelease() SKRegion {
	rv := objc.Send[SKRegion](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKRegion creates a new SKRegion instance.
func NewSKRegion() SKRegion {
	return getSKRegionClass().New()
}


// Initializes a new region with a rectangular area.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(size:)
func NewSKRegionWithSize(size unsafe.Pointer) SKRegion {
	instance := getSKRegionClass().Alloc()
	rv := objc.Send[SKRegion](instance.ID, objc.Sel("initWithSize:"), size)
	rv.Autorelease()
	return rv
}
// Initializes a new region using a Core Graphics path.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(path:)
func NewSKRegionWithPath(path coregraphics.CGPathRef) SKRegion {
	instance := getSKRegionClass().Alloc()
	rv := objc.Send[SKRegion](instance.ID, objc.Sel("initWithPath:"), path)
	rv.Autorelease()
	return rv
}
// Initializes a new region with a circular area.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/init(radius:)
func NewSKRegionWithRadius(radius float32) SKRegion {
	instance := getSKRegionClass().Alloc()
	rv := objc.Send[SKRegion](instance.ID, objc.Sel("initWithRadius:"), radius)
	rv.Autorelease()
	return rv
}


// Returns a region that defines a region that includes all points.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/infinite()
func (sc _SKRegionClass) InfiniteRegion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("infiniteRegion"))
	return rv
}
// Returns a new region created by subtracting the contents of another region from this region.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/byDifference(from:)
func (s_ SKRegion) RegionByDifferenceFromRegion(region unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("regionByDifferenceFromRegion:"), region)
	return rv
}
// Returns a new region created by intersecting the contents of this region with another region.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/byIntersection(with:)
func (s_ SKRegion) RegionByIntersectionWithRegion(region unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("regionByIntersectionWithRegion:"), region)
	return rv
}
// Returns a new region created by combining the contents of this region with another region.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/byUnion(with:)
func (s_ SKRegion) RegionByUnionWithRegion(region unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("regionByUnionWithRegion:"), region)
	return rv
}
// Returns a Boolean value that indicates whether a particular point is contained in the region.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/contains(_:)
func (s_ SKRegion) ContainsPoint(point unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containsPoint:"), point)
	return rv
}
// Returns a new region that is the mathematical inverse of an existing region.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRegion/inverse()
func (s_ SKRegion) InverseRegion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("inverseRegion"))
	return rv
}

