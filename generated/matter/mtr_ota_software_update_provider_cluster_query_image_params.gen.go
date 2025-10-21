// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterQueryImageParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass     _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass
	MTROtaSoftwareUpdateProviderClusterQueryImageParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterQueryImageParamsClass() _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass {
	MTROtaSoftwareUpdateProviderClusterQueryImageParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass = _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterQueryImageParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterQueryImageParams] class.
type IMTROtaSoftwareUpdateProviderClusterQueryImageParams interface {
	IMTROTASoftwareUpdateProviderClusterQueryImageParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv
type MTROtaSoftwareUpdateProviderClusterQueryImageParams struct {
	MTROTASoftwareUpdateProviderClusterQueryImageParams
}

// MTROtaSoftwareUpdateProviderClusterQueryImageParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterQueryImageParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterQueryImageParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	return MTROtaSoftwareUpdateProviderClusterQueryImageParams{
		MTROTASoftwareUpdateProviderClusterQueryImageParams: MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass) New() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Init() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Autorelease() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterQueryImageParams creates a new MTROtaSoftwareUpdateProviderClusterQueryImageParams instance.
func NewMTROtaSoftwareUpdateProviderClusterQueryImageParams() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	return getMTROtaSoftwareUpdateProviderClusterQueryImageParamsClass().New()
}




