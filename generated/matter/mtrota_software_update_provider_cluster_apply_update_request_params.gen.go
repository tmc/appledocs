// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
var (
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass     _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass() _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass {
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass = _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams")}
	})
	return MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
}

type _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
type IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams-1mlcr
type MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) New() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) Init() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams) Autorelease() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams creates a new MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams instance.
func NewMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams() MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return getMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsClass().New()
}




