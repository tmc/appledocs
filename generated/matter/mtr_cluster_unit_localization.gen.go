// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterUnitLocalization] class.
var (
	MTRClusterUnitLocalizationClass     _MTRClusterUnitLocalizationClass
	MTRClusterUnitLocalizationClassOnce sync.Once
)

func getMTRClusterUnitLocalizationClass() _MTRClusterUnitLocalizationClass {
	MTRClusterUnitLocalizationClassOnce.Do(func() {
		MTRClusterUnitLocalizationClass = _MTRClusterUnitLocalizationClass{objc.GetClass("MTRClusterUnitLocalization")}
	})
	return MTRClusterUnitLocalizationClass
}

type _MTRClusterUnitLocalizationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterUnitLocalization] class.
type IMTRClusterUnitLocalization interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterUnitLocalization
type MTRClusterUnitLocalization struct {
	MTRGenericCluster
}

// MTRClusterUnitLocalizationFrom constructs a [MTRClusterUnitLocalization] from an unsafe.Pointer.
func MTRClusterUnitLocalizationFrom(ptr unsafe.Pointer) MTRClusterUnitLocalization {
	return MTRClusterUnitLocalization{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterUnitLocalizationClass) Alloc() MTRClusterUnitLocalization {
	rv := objc.Send[MTRClusterUnitLocalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterUnitLocalizationClass) New() MTRClusterUnitLocalization {
	rv := objc.Send[MTRClusterUnitLocalization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterUnitLocalization) Init() MTRClusterUnitLocalization {
	rv := objc.Send[MTRClusterUnitLocalization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterUnitLocalization) Autorelease() MTRClusterUnitLocalization {
	rv := objc.Send[MTRClusterUnitLocalization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterUnitLocalization creates a new MTRClusterUnitLocalization instance.
func NewMTRClusterUnitLocalization() MTRClusterUnitLocalization {
	return getMTRClusterUnitLocalizationClass().New()
}




