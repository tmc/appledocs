// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterTimeFormatLocalization] class.
var (
	MTRClusterTimeFormatLocalizationClass     _MTRClusterTimeFormatLocalizationClass
	MTRClusterTimeFormatLocalizationClassOnce sync.Once
)

func getMTRClusterTimeFormatLocalizationClass() _MTRClusterTimeFormatLocalizationClass {
	MTRClusterTimeFormatLocalizationClassOnce.Do(func() {
		MTRClusterTimeFormatLocalizationClass = _MTRClusterTimeFormatLocalizationClass{objc.GetClass("MTRClusterTimeFormatLocalization")}
	})
	return MTRClusterTimeFormatLocalizationClass
}

type _MTRClusterTimeFormatLocalizationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTimeFormatLocalization] class.
type IMTRClusterTimeFormatLocalization interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTimeFormatLocalization
type MTRClusterTimeFormatLocalization struct {
	MTRGenericCluster
}

// MTRClusterTimeFormatLocalizationFrom constructs a [MTRClusterTimeFormatLocalization] from an unsafe.Pointer.
func MTRClusterTimeFormatLocalizationFrom(ptr unsafe.Pointer) MTRClusterTimeFormatLocalization {
	return MTRClusterTimeFormatLocalization{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTimeFormatLocalizationClass) Alloc() MTRClusterTimeFormatLocalization {
	rv := objc.Send[MTRClusterTimeFormatLocalization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTimeFormatLocalizationClass) New() MTRClusterTimeFormatLocalization {
	rv := objc.Send[MTRClusterTimeFormatLocalization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTimeFormatLocalization) Init() MTRClusterTimeFormatLocalization {
	rv := objc.Send[MTRClusterTimeFormatLocalization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTimeFormatLocalization) Autorelease() MTRClusterTimeFormatLocalization {
	rv := objc.Send[MTRClusterTimeFormatLocalization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTimeFormatLocalization creates a new MTRClusterTimeFormatLocalization instance.
func NewMTRClusterTimeFormatLocalization() MTRClusterTimeFormatLocalization {
	return getMTRClusterTimeFormatLocalizationClass().New()
}
