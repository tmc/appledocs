// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass     _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass
	MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass() _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass {
	MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass = _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] class.
type IMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams interface {
	IMTROTASoftwareUpdateProviderClusterQueryImageResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao
type MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams struct {
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParams
}

// MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	return MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams{
		MTROTASoftwareUpdateProviderClusterQueryImageResponseParams: MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass) New() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Init() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Autorelease() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams creates a new MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams instance.
func NewMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	return getMTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass().New()
}




