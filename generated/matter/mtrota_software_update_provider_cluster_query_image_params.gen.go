// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterQueryImageParams] class.
var (
	MTROTASoftwareUpdateProviderClusterQueryImageParamsClass     _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass
	MTROTASoftwareUpdateProviderClusterQueryImageParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterQueryImageParamsClass() _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass {
	MTROTASoftwareUpdateProviderClusterQueryImageParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterQueryImageParamsClass = _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterQueryImageParams")}
	})
	return MTROTASoftwareUpdateProviderClusterQueryImageParamsClass
}

type _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterQueryImageParams] class.
type IMTROTASoftwareUpdateProviderClusterQueryImageParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b
type MTROTASoftwareUpdateProviderClusterQueryImageParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterQueryImageParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageParams {
	return MTROTASoftwareUpdateProviderClusterQueryImageParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass) New() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Init() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Autorelease() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterQueryImageParams creates a new MTROTASoftwareUpdateProviderClusterQueryImageParams instance.
func NewMTROTASoftwareUpdateProviderClusterQueryImageParams() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	return getMTROTASoftwareUpdateProviderClusterQueryImageParamsClass().New()
}




