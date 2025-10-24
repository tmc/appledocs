// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterUnitLocalization] class.
var (
	MTRBaseClusterUnitLocalizationClass     _MTRBaseClusterUnitLocalizationClass
	MTRBaseClusterUnitLocalizationClassOnce sync.Once
)

func getMTRBaseClusterUnitLocalizationClass() _MTRBaseClusterUnitLocalizationClass {
	MTRBaseClusterUnitLocalizationClassOnce.Do(func() {
		MTRBaseClusterUnitLocalizationClass = _MTRBaseClusterUnitLocalizationClass{objc.GetClass("MTRBaseClusterUnitLocalization")}
	})
	return MTRBaseClusterUnitLocalizationClass
}

type _MTRBaseClusterUnitLocalizationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterUnitLocalization] class.
type IMTRBaseClusterUnitLocalization interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterUnitLocalization
type MTRBaseClusterUnitLocalization struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterUnitLocalizationFrom constructs a [MTRBaseClusterUnitLocalization] from an unsafe.Pointer.
func MTRBaseClusterUnitLocalizationFrom(ptr unsafe.Pointer) MTRBaseClusterUnitLocalization {
	return MTRBaseClusterUnitLocalization{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterUnitLocalizationClass) Alloc() MTRBaseClusterUnitLocalization {
	rv := objc.Send[MTRBaseClusterUnitLocalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterUnitLocalizationClass) New() MTRBaseClusterUnitLocalization {
	rv := objc.Send[MTRBaseClusterUnitLocalization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterUnitLocalization) Init() MTRBaseClusterUnitLocalization {
	rv := objc.Send[MTRBaseClusterUnitLocalization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterUnitLocalization) Autorelease() MTRBaseClusterUnitLocalization {
	rv := objc.Send[MTRBaseClusterUnitLocalization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterUnitLocalization creates a new MTRBaseClusterUnitLocalization instance.
func NewMTRBaseClusterUnitLocalization() MTRBaseClusterUnitLocalization {
	return getMTRBaseClusterUnitLocalizationClass().New()
}




