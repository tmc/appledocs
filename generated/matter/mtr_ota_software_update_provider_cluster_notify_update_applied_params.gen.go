// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass     _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass() _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass {
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass = _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
type IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams interface {
	IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv
type MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams struct {
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams
}

// MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams{
		MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams: MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) New() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Init() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Autorelease() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams creates a new MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams instance.
func NewMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return getMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass().New()
}




