// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitVolume] class.
var (
	UnitVolumeClass     _UnitVolumeClass
	UnitVolumeClassOnce sync.Once
)

func getUnitVolumeClass() _UnitVolumeClass {
	UnitVolumeClassOnce.Do(func() {
		UnitVolumeClass = _UnitVolumeClass{objc.GetClass("NSUnitVolume")}
	})
	return UnitVolumeClass
}

type _UnitVolumeClass struct {
	class objc.Class
}

// An interface definition for the [UnitVolume] class.
type IUnitVolume interface {
	IDimension
}

// A unit of measure for volume.
//
// You typically use instances of to represent specific quantities of volume using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume
type UnitVolume struct {
	Dimension
}

// UnitVolumeFrom constructs a [UnitVolume] from an unsafe.Pointer.
//
// A unit of measure for volume.
func UnitVolumeFrom(ptr unsafe.Pointer) UnitVolume {
	return UnitVolume{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitVolumeClass) Alloc() UnitVolume {
	rv := objc.Send[UnitVolume](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitVolumeClass) New() UnitVolume {
	rv := objc.Send[UnitVolume](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitVolume) Init() UnitVolume {
	rv := objc.Send[UnitVolume](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitVolume) Autorelease() UnitVolume {
	rv := objc.Send[UnitVolume](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitVolume creates a new UnitVolume instance.
func NewUnitVolume() UnitVolume {
	return getUnitVolumeClass().New()
}


// The imperial gallons unit of volume.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialGallons
func (uc _UnitVolumeClass) ImperialGallons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialGallons"))
	return rv
}
// The liters unit of volume.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/liters
func (uc _UnitVolumeClass) Liters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("liters"))
	return rv
}
// The imperial gallons unit of volume.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialGallons
func (u_ UnitVolume) ImperialGallons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialGallons"))
	return rv
}

// The liters unit of volume.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/liters
func (u_ UnitVolume) Liters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("liters"))
	return rv
}



