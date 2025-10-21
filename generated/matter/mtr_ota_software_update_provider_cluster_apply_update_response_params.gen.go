// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass     _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass() _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass {
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass = _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
type IMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams interface {
	IMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als
type MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams struct {
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams
}

// MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams{
		MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams: MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) New() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Init() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Autorelease() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams creates a new MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams instance.
func NewMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return getMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().New()
}




