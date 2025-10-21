// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] class.
var (
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass     _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass() _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass {
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass = _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterQueryImageResponseParams")}
	})
	return MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass
}

type _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] class.
type IMTROTASoftwareUpdateProviderClusterQueryImageResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt
type MTROTASoftwareUpdateProviderClusterQueryImageResponseParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	return MTROTASoftwareUpdateProviderClusterQueryImageResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass) New() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Init() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Autorelease() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterQueryImageResponseParams creates a new MTROTASoftwareUpdateProviderClusterQueryImageResponseParams instance.
func NewMTROTASoftwareUpdateProviderClusterQueryImageResponseParams() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	return getMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass().New()
}




