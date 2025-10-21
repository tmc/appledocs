// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTargetNavigatorClusterNavigateTargetParams] class.
var (
	MTRTargetNavigatorClusterNavigateTargetParamsClass     _MTRTargetNavigatorClusterNavigateTargetParamsClass
	MTRTargetNavigatorClusterNavigateTargetParamsClassOnce sync.Once
)

func getMTRTargetNavigatorClusterNavigateTargetParamsClass() _MTRTargetNavigatorClusterNavigateTargetParamsClass {
	MTRTargetNavigatorClusterNavigateTargetParamsClassOnce.Do(func() {
		MTRTargetNavigatorClusterNavigateTargetParamsClass = _MTRTargetNavigatorClusterNavigateTargetParamsClass{objc.GetClass("MTRTargetNavigatorClusterNavigateTargetParams")}
	})
	return MTRTargetNavigatorClusterNavigateTargetParamsClass
}

type _MTRTargetNavigatorClusterNavigateTargetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTargetNavigatorClusterNavigateTargetParams] class.
type IMTRTargetNavigatorClusterNavigateTargetParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterNavigateTargetParams
type MTRTargetNavigatorClusterNavigateTargetParams struct {
	objectivec.Object
}

// MTRTargetNavigatorClusterNavigateTargetParamsFrom constructs a [MTRTargetNavigatorClusterNavigateTargetParams] from an unsafe.Pointer.
func MTRTargetNavigatorClusterNavigateTargetParamsFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterNavigateTargetParams {
	return MTRTargetNavigatorClusterNavigateTargetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterNavigateTargetParamsClass) Alloc() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTargetNavigatorClusterNavigateTargetParamsClass) New() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Init() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Autorelease() MTRTargetNavigatorClusterNavigateTargetParams {
	rv := objc.Send[MTRTargetNavigatorClusterNavigateTargetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterNavigateTargetParams creates a new MTRTargetNavigatorClusterNavigateTargetParams instance.
func NewMTRTargetNavigatorClusterNavigateTargetParams() MTRTargetNavigatorClusterNavigateTargetParams {
	return getMTRTargetNavigatorClusterNavigateTargetParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/data
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Data() string {
	rv := objc.Send[string](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/data
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetData(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/serversideprocessingtimeout
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/serversideprocessingtimeout
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/target
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) Target() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/target
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetTarget(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTarget:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/timedinvoketimeoutms
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclusternavigatetargetparams/timedinvoketimeoutms
func (m_ MTRTargetNavigatorClusterNavigateTargetParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



