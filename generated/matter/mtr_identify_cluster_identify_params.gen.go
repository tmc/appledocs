// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRIdentifyClusterIdentifyParams] class.
var (
	MTRIdentifyClusterIdentifyParamsClass     _MTRIdentifyClusterIdentifyParamsClass
	MTRIdentifyClusterIdentifyParamsClassOnce sync.Once
)

func getMTRIdentifyClusterIdentifyParamsClass() _MTRIdentifyClusterIdentifyParamsClass {
	MTRIdentifyClusterIdentifyParamsClassOnce.Do(func() {
		MTRIdentifyClusterIdentifyParamsClass = _MTRIdentifyClusterIdentifyParamsClass{objc.GetClass("MTRIdentifyClusterIdentifyParams")}
	})
	return MTRIdentifyClusterIdentifyParamsClass
}

type _MTRIdentifyClusterIdentifyParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRIdentifyClusterIdentifyParams] class.
type IMTRIdentifyClusterIdentifyParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRIdentifyClusterIdentifyParams
type MTRIdentifyClusterIdentifyParams struct {
	objectivec.Object
}

// MTRIdentifyClusterIdentifyParamsFrom constructs a [MTRIdentifyClusterIdentifyParams] from an unsafe.Pointer.
func MTRIdentifyClusterIdentifyParamsFrom(ptr unsafe.Pointer) MTRIdentifyClusterIdentifyParams {
	return MTRIdentifyClusterIdentifyParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRIdentifyClusterIdentifyParamsClass) Alloc() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRIdentifyClusterIdentifyParamsClass) New() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRIdentifyClusterIdentifyParams) Init() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRIdentifyClusterIdentifyParams) Autorelease() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRIdentifyClusterIdentifyParams creates a new MTRIdentifyClusterIdentifyParams instance.
func NewMTRIdentifyClusterIdentifyParams() MTRIdentifyClusterIdentifyParams {
	return getMTRIdentifyClusterIdentifyParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/identifytime
func (m_ MTRIdentifyClusterIdentifyParams) IdentifyTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("identifyTime"))
	return rv
}


// SetIdentifyTime sets the value of the identifyTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/identifytime
func (m_ MTRIdentifyClusterIdentifyParams) SetIdentifyTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifyTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/serversideprocessingtimeout
func (m_ MTRIdentifyClusterIdentifyParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/serversideprocessingtimeout
func (m_ MTRIdentifyClusterIdentifyParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/timedinvoketimeoutms
func (m_ MTRIdentifyClusterIdentifyParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/timedinvoketimeoutms
func (m_ MTRIdentifyClusterIdentifyParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



