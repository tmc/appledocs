// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterDimensionStruct] class.
var (
	MTRContentLauncherClusterDimensionStructClass     _MTRContentLauncherClusterDimensionStructClass
	MTRContentLauncherClusterDimensionStructClassOnce sync.Once
)

func getMTRContentLauncherClusterDimensionStructClass() _MTRContentLauncherClusterDimensionStructClass {
	MTRContentLauncherClusterDimensionStructClassOnce.Do(func() {
		MTRContentLauncherClusterDimensionStructClass = _MTRContentLauncherClusterDimensionStructClass{objc.GetClass("MTRContentLauncherClusterDimensionStruct")}
	})
	return MTRContentLauncherClusterDimensionStructClass
}

type _MTRContentLauncherClusterDimensionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterDimensionStruct] class.
type IMTRContentLauncherClusterDimensionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterDimensionStruct
type MTRContentLauncherClusterDimensionStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterDimensionStructFrom constructs a [MTRContentLauncherClusterDimensionStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterDimensionStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterDimensionStruct {
	return MTRContentLauncherClusterDimensionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterDimensionStructClass) Alloc() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterDimensionStructClass) New() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterDimensionStruct) Init() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterDimensionStruct) Autorelease() MTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterDimensionStruct creates a new MTRContentLauncherClusterDimensionStruct instance.
func NewMTRContentLauncherClusterDimensionStruct() MTRContentLauncherClusterDimensionStruct {
	return getMTRContentLauncherClusterDimensionStructClass().New()
}




