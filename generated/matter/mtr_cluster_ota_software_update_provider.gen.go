// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOtaSoftwareUpdateProvider] class.
var (
	MTRClusterOtaSoftwareUpdateProviderClass     _MTRClusterOtaSoftwareUpdateProviderClass
	MTRClusterOtaSoftwareUpdateProviderClassOnce sync.Once
)

func getMTRClusterOtaSoftwareUpdateProviderClass() _MTRClusterOtaSoftwareUpdateProviderClass {
	MTRClusterOtaSoftwareUpdateProviderClassOnce.Do(func() {
		MTRClusterOtaSoftwareUpdateProviderClass = _MTRClusterOtaSoftwareUpdateProviderClass{objc.GetClass("MTRClusterOtaSoftwareUpdateProvider")}
	})
	return MTRClusterOtaSoftwareUpdateProviderClass
}

type _MTRClusterOtaSoftwareUpdateProviderClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOtaSoftwareUpdateProvider] class.
type IMTRClusterOtaSoftwareUpdateProvider interface {
	IMTRClusterOTASoftwareUpdateProvider
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOtaSoftwareUpdateProvider-4p8us
type MTRClusterOtaSoftwareUpdateProvider struct {
	MTRClusterOTASoftwareUpdateProvider
}

// MTRClusterOtaSoftwareUpdateProviderFrom constructs a [MTRClusterOtaSoftwareUpdateProvider] from an unsafe.Pointer.
func MTRClusterOtaSoftwareUpdateProviderFrom(ptr unsafe.Pointer) MTRClusterOtaSoftwareUpdateProvider {
	return MTRClusterOtaSoftwareUpdateProvider{
		MTRClusterOTASoftwareUpdateProvider: MTRClusterOTASoftwareUpdateProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOtaSoftwareUpdateProviderClass) Alloc() MTRClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOtaSoftwareUpdateProviderClass) New() MTRClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateProvider](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOtaSoftwareUpdateProvider) Init() MTRClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateProvider](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOtaSoftwareUpdateProvider) Autorelease() MTRClusterOtaSoftwareUpdateProvider {
	rv := objc.Send[MTRClusterOtaSoftwareUpdateProvider](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOtaSoftwareUpdateProvider creates a new MTRClusterOtaSoftwareUpdateProvider instance.
func NewMTRClusterOtaSoftwareUpdateProvider() MTRClusterOtaSoftwareUpdateProvider {
	return getMTRClusterOtaSoftwareUpdateProviderClass().New()
}
