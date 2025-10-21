// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnOffClusterOnWithTimedOffParams] class.
var (
	MTROnOffClusterOnWithTimedOffParamsClass     _MTROnOffClusterOnWithTimedOffParamsClass
	MTROnOffClusterOnWithTimedOffParamsClassOnce sync.Once
)

func getMTROnOffClusterOnWithTimedOffParamsClass() _MTROnOffClusterOnWithTimedOffParamsClass {
	MTROnOffClusterOnWithTimedOffParamsClassOnce.Do(func() {
		MTROnOffClusterOnWithTimedOffParamsClass = _MTROnOffClusterOnWithTimedOffParamsClass{objc.GetClass("MTROnOffClusterOnWithTimedOffParams")}
	})
	return MTROnOffClusterOnWithTimedOffParamsClass
}

type _MTROnOffClusterOnWithTimedOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterOnWithTimedOffParams] class.
type IMTROnOffClusterOnWithTimedOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithTimedOffParams
type MTROnOffClusterOnWithTimedOffParams struct {
	objectivec.Object
}

// MTROnOffClusterOnWithTimedOffParamsFrom constructs a [MTROnOffClusterOnWithTimedOffParams] from an unsafe.Pointer.
func MTROnOffClusterOnWithTimedOffParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOnWithTimedOffParams {
	return MTROnOffClusterOnWithTimedOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOnWithTimedOffParamsClass) Alloc() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterOnWithTimedOffParamsClass) New() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOnWithTimedOffParams) Init() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOnWithTimedOffParams) Autorelease() MTROnOffClusterOnWithTimedOffParams {
	rv := objc.Send[MTROnOffClusterOnWithTimedOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOnWithTimedOffParams creates a new MTROnOffClusterOnWithTimedOffParams instance.
func NewMTROnOffClusterOnWithTimedOffParams() MTROnOffClusterOnWithTimedOffParams {
	return getMTROnOffClusterOnWithTimedOffParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/offwaittime
func (m_ MTROnOffClusterOnWithTimedOffParams) OffWaitTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offWaitTime"))
	return rv
}


// SetOffWaitTime sets the value of the offWaitTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/offwaittime
func (m_ MTROnOffClusterOnWithTimedOffParams) SetOffWaitTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffWaitTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/onoffcontrol
func (m_ MTROnOffClusterOnWithTimedOffParams) OnOffControl() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("onOffControl"))
	return rv
}


// SetOnOffControl sets the value of the onOffControl property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/onoffcontrol
func (m_ MTROnOffClusterOnWithTimedOffParams) SetOnOffControl(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnOffControl:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/ontime
func (m_ MTROnOffClusterOnWithTimedOffParams) OnTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("onTime"))
	return rv
}


// SetOnTime sets the value of the onTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/ontime
func (m_ MTROnOffClusterOnWithTimedOffParams) SetOnTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOnTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOnWithTimedOffParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOnWithTimedOffParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOnWithTimedOffParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronwithtimedoffparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOnWithTimedOffParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



