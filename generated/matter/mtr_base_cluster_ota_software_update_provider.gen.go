// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOtaSoftwareUpdateProvider] class.
var (
	MTRBaseClusterOtaSoftwareUpdateProviderClass     _MTRBaseClusterOtaSoftwareUpdateProviderClass
	MTRBaseClusterOtaSoftwareUpdateProviderClassOnce sync.Once
)

func getMTRBaseClusterOtaSoftwareUpdateProviderClass() _MTRBaseClusterOtaSoftwareUpdateProviderClass {
	MTRBaseClusterOtaSoftwareUpdateProviderClassOnce.Do(func() {
		MTRBaseClusterOtaSoftwareUpdateProviderClass = _MTRBaseClusterOtaSoftwareUpdateProviderClass{objc.GetClass("MTRBaseClusterOtaSoftwareUpdateProvider")}
	})
	return MTRBaseClusterOtaSoftwareUpdateProviderClass
}

type _MTRBaseClusterOtaSoftwareUpdateProviderClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOtaSoftwareUpdateProvider] class.
type IMTRBaseClusterOtaSoftwareUpdateProvider interface {
	IMTRBaseClusterOTASoftwareUpdateProvider
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOtaSoftwareUpdateProvider-2vync
type MTRBaseClusterOtaSoftwareUpdateProvider struct {
	MTRBaseClusterOTASoftwareUpdateProvider
}

// MTRBaseClusterOtaSoftwareUpdateProviderFrom constructs a [MTRBaseClusterOtaSoftwareUpdateProvider] from an unsafe.Pointer.
func MTRBaseClusterOtaSoftwareUpdateProviderFrom(ptr unsafe.Pointer) MTRBaseClusterOtaSoftwareUpdateProvider {
	return MTRBaseClusterOtaSoftwareUpdateProvider{
		MTRBaseClusterOTASoftwareUpdateProvider: MTRBaseClusterOTASoftwareUpdateProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOtaSoftwareUpdateProviderClass) Alloc() MTRBaseClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOtaSoftwareUpdateProviderClass) New() MTRBaseClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateProvider](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOtaSoftwareUpdateProvider) Init() MTRBaseClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateProvider](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOtaSoftwareUpdateProvider) Autorelease() MTRBaseClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOtaSoftwareUpdateProvider](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOtaSoftwareUpdateProvider creates a new MTRBaseClusterOtaSoftwareUpdateProvider instance.
func NewMTRBaseClusterOtaSoftwareUpdateProvider() MTRBaseClusterOtaSoftwareUpdateProvider {
	return getMTRBaseClusterOtaSoftwareUpdateProviderClass().New()
}




