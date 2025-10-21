// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAudioOutputClusterSelectOutputParams] class.
var (
	MTRAudioOutputClusterSelectOutputParamsClass     _MTRAudioOutputClusterSelectOutputParamsClass
	MTRAudioOutputClusterSelectOutputParamsClassOnce sync.Once
)

func getMTRAudioOutputClusterSelectOutputParamsClass() _MTRAudioOutputClusterSelectOutputParamsClass {
	MTRAudioOutputClusterSelectOutputParamsClassOnce.Do(func() {
		MTRAudioOutputClusterSelectOutputParamsClass = _MTRAudioOutputClusterSelectOutputParamsClass{objc.GetClass("MTRAudioOutputClusterSelectOutputParams")}
	})
	return MTRAudioOutputClusterSelectOutputParamsClass
}

type _MTRAudioOutputClusterSelectOutputParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAudioOutputClusterSelectOutputParams] class.
type IMTRAudioOutputClusterSelectOutputParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAudioOutputClusterSelectOutputParams
type MTRAudioOutputClusterSelectOutputParams struct {
	objectivec.Object
}

// MTRAudioOutputClusterSelectOutputParamsFrom constructs a [MTRAudioOutputClusterSelectOutputParams] from an unsafe.Pointer.
func MTRAudioOutputClusterSelectOutputParamsFrom(ptr unsafe.Pointer) MTRAudioOutputClusterSelectOutputParams {
	return MTRAudioOutputClusterSelectOutputParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAudioOutputClusterSelectOutputParamsClass) Alloc() MTRAudioOutputClusterSelectOutputParams {
	rv := objc.Send[MTRAudioOutputClusterSelectOutputParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAudioOutputClusterSelectOutputParamsClass) New() MTRAudioOutputClusterSelectOutputParams {
	rv := objc.Send[MTRAudioOutputClusterSelectOutputParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAudioOutputClusterSelectOutputParams) Init() MTRAudioOutputClusterSelectOutputParams {
	rv := objc.Send[MTRAudioOutputClusterSelectOutputParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAudioOutputClusterSelectOutputParams) Autorelease() MTRAudioOutputClusterSelectOutputParams {
	rv := objc.Send[MTRAudioOutputClusterSelectOutputParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAudioOutputClusterSelectOutputParams creates a new MTRAudioOutputClusterSelectOutputParams instance.
func NewMTRAudioOutputClusterSelectOutputParams() MTRAudioOutputClusterSelectOutputParams {
	return getMTRAudioOutputClusterSelectOutputParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterselectoutputparams/index
func (m_ MTRAudioOutputClusterSelectOutputParams) Index() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("index"))
	return rv
}


// SetIndex sets the value of the index property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterselectoutputparams/index
func (m_ MTRAudioOutputClusterSelectOutputParams) SetIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterselectoutputparams/serversideprocessingtimeout
func (m_ MTRAudioOutputClusterSelectOutputParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterselectoutputparams/serversideprocessingtimeout
func (m_ MTRAudioOutputClusterSelectOutputParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterselectoutputparams/timedinvoketimeoutms
func (m_ MTRAudioOutputClusterSelectOutputParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusterselectoutputparams/timedinvoketimeoutms
func (m_ MTRAudioOutputClusterSelectOutputParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



