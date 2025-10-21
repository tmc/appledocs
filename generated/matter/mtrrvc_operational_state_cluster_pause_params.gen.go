// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterPauseParams] class.
var (
	MTRRVCOperationalStateClusterPauseParamsClass     _MTRRVCOperationalStateClusterPauseParamsClass
	MTRRVCOperationalStateClusterPauseParamsClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterPauseParamsClass() _MTRRVCOperationalStateClusterPauseParamsClass {
	MTRRVCOperationalStateClusterPauseParamsClassOnce.Do(func() {
		MTRRVCOperationalStateClusterPauseParamsClass = _MTRRVCOperationalStateClusterPauseParamsClass{objc.GetClass("MTRRVCOperationalStateClusterPauseParams")}
	})
	return MTRRVCOperationalStateClusterPauseParamsClass
}

type _MTRRVCOperationalStateClusterPauseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterPauseParams] class.
type IMTRRVCOperationalStateClusterPauseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterPauseParams
type MTRRVCOperationalStateClusterPauseParams struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterPauseParamsFrom constructs a [MTRRVCOperationalStateClusterPauseParams] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterPauseParamsFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterPauseParams {
	return MTRRVCOperationalStateClusterPauseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterPauseParamsClass) Alloc() MTRRVCOperationalStateClusterPauseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterPauseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterPauseParamsClass) New() MTRRVCOperationalStateClusterPauseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterPauseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterPauseParams) Init() MTRRVCOperationalStateClusterPauseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterPauseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterPauseParams) Autorelease() MTRRVCOperationalStateClusterPauseParams {
	rv := objc.Send[MTRRVCOperationalStateClusterPauseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterPauseParams creates a new MTRRVCOperationalStateClusterPauseParams instance.
func NewMTRRVCOperationalStateClusterPauseParams() MTRRVCOperationalStateClusterPauseParams {
	return getMTRRVCOperationalStateClusterPauseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusterpauseparams/serversideprocessingtimeout
func (m_ MTRRVCOperationalStateClusterPauseParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusterpauseparams/serversideprocessingtimeout
func (m_ MTRRVCOperationalStateClusterPauseParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusterpauseparams/timedinvoketimeoutms
func (m_ MTRRVCOperationalStateClusterPauseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusterpauseparams/timedinvoketimeoutms
func (m_ MTRRVCOperationalStateClusterPauseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



