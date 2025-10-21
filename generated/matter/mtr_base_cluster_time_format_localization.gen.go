// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterTimeFormatLocalization] class.
var (
	MTRBaseClusterTimeFormatLocalizationClass     _MTRBaseClusterTimeFormatLocalizationClass
	MTRBaseClusterTimeFormatLocalizationClassOnce sync.Once
)

func getMTRBaseClusterTimeFormatLocalizationClass() _MTRBaseClusterTimeFormatLocalizationClass {
	MTRBaseClusterTimeFormatLocalizationClassOnce.Do(func() {
		MTRBaseClusterTimeFormatLocalizationClass = _MTRBaseClusterTimeFormatLocalizationClass{objc.GetClass("MTRBaseClusterTimeFormatLocalization")}
	})
	return MTRBaseClusterTimeFormatLocalizationClass
}

type _MTRBaseClusterTimeFormatLocalizationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTimeFormatLocalization] class.
type IMTRBaseClusterTimeFormatLocalization interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeFormatLocalization
type MTRBaseClusterTimeFormatLocalization struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTimeFormatLocalizationFrom constructs a [MTRBaseClusterTimeFormatLocalization] from an unsafe.Pointer.
func MTRBaseClusterTimeFormatLocalizationFrom(ptr unsafe.Pointer) MTRBaseClusterTimeFormatLocalization {
	return MTRBaseClusterTimeFormatLocalization{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTimeFormatLocalizationClass) Alloc() MTRBaseClusterTimeFormatLocalization {
	rv := objc.Send[MTRBaseClusterTimeFormatLocalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTimeFormatLocalizationClass) New() MTRBaseClusterTimeFormatLocalization {
	rv := objc.Send[MTRBaseClusterTimeFormatLocalization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTimeFormatLocalization) Init() MTRBaseClusterTimeFormatLocalization {
	rv := objc.Send[MTRBaseClusterTimeFormatLocalization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTimeFormatLocalization) Autorelease() MTRBaseClusterTimeFormatLocalization {
	rv := objc.Send[MTRBaseClusterTimeFormatLocalization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTimeFormatLocalization creates a new MTRBaseClusterTimeFormatLocalization instance.
func NewMTRBaseClusterTimeFormatLocalization() MTRBaseClusterTimeFormatLocalization {
	return getMTRBaseClusterTimeFormatLocalizationClass().New()
}




