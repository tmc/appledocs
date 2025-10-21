// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRContentLauncherClusterDimension] class.
var (
	MTRContentLauncherClusterDimensionClass     _MTRContentLauncherClusterDimensionClass
	MTRContentLauncherClusterDimensionClassOnce sync.Once
)

func getMTRContentLauncherClusterDimensionClass() _MTRContentLauncherClusterDimensionClass {
	MTRContentLauncherClusterDimensionClassOnce.Do(func() {
		MTRContentLauncherClusterDimensionClass = _MTRContentLauncherClusterDimensionClass{objc.GetClass("MTRContentLauncherClusterDimension")}
	})
	return MTRContentLauncherClusterDimensionClass
}

type _MTRContentLauncherClusterDimensionClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterDimension] class.
type IMTRContentLauncherClusterDimension interface {
	IMTRContentLauncherClusterDimensionStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimension
type MTRContentLauncherClusterDimension struct {
	MTRContentLauncherClusterDimensionStruct
}

// MTRContentLauncherClusterDimensionFrom constructs a [MTRContentLauncherClusterDimension] from an unsafe.Pointer.
func MTRContentLauncherClusterDimensionFrom(ptr unsafe.Pointer) MTRContentLauncherClusterDimension {
	return MTRContentLauncherClusterDimension{
		MTRContentLauncherClusterDimensionStruct: MTRContentLauncherClusterDimensionStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterDimensionClass) Alloc() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterDimensionClass) New() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterDimension) Init() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterDimension) Autorelease() MTRContentLauncherClusterDimension {
	rv := objc.Send[MTRContentLauncherClusterDimension](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterDimension creates a new MTRContentLauncherClusterDimension instance.
func NewMTRContentLauncherClusterDimension() MTRContentLauncherClusterDimension {
	return getMTRContentLauncherClusterDimensionClass().New()
}




