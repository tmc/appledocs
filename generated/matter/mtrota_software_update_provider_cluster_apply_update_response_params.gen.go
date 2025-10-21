// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
var (
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass     _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass() _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass {
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass = _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams")}
	})
	return MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
}

type _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
type IMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9
type MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) New() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Init() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Autorelease() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams creates a new MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams instance.
func NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().New()
}




