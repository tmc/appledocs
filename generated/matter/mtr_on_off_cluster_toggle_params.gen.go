// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnOffClusterToggleParams] class.
var (
	MTROnOffClusterToggleParamsClass     _MTROnOffClusterToggleParamsClass
	MTROnOffClusterToggleParamsClassOnce sync.Once
)

func getMTROnOffClusterToggleParamsClass() _MTROnOffClusterToggleParamsClass {
	MTROnOffClusterToggleParamsClassOnce.Do(func() {
		MTROnOffClusterToggleParamsClass = _MTROnOffClusterToggleParamsClass{objc.GetClass("MTROnOffClusterToggleParams")}
	})
	return MTROnOffClusterToggleParamsClass
}

type _MTROnOffClusterToggleParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterToggleParams] class.
type IMTROnOffClusterToggleParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterToggleParams
type MTROnOffClusterToggleParams struct {
	objectivec.Object
}

// MTROnOffClusterToggleParamsFrom constructs a [MTROnOffClusterToggleParams] from an unsafe.Pointer.
func MTROnOffClusterToggleParamsFrom(ptr unsafe.Pointer) MTROnOffClusterToggleParams {
	return MTROnOffClusterToggleParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterToggleParamsClass) Alloc() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterToggleParamsClass) New() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterToggleParams) Init() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterToggleParams) Autorelease() MTROnOffClusterToggleParams {
	rv := objc.Send[MTROnOffClusterToggleParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterToggleParams creates a new MTROnOffClusterToggleParams instance.
func NewMTROnOffClusterToggleParams() MTROnOffClusterToggleParams {
	return getMTROnOffClusterToggleParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclustertoggleparams/serversideprocessingtimeout
func (m_ MTROnOffClusterToggleParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclustertoggleparams/serversideprocessingtimeout
func (m_ MTROnOffClusterToggleParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclustertoggleparams/timedinvoketimeoutms
func (m_ MTROnOffClusterToggleParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclustertoggleparams/timedinvoketimeoutms
func (m_ MTROnOffClusterToggleParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



