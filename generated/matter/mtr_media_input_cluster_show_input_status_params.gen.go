// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaInputClusterShowInputStatusParams] class.
var (
	MTRMediaInputClusterShowInputStatusParamsClass     _MTRMediaInputClusterShowInputStatusParamsClass
	MTRMediaInputClusterShowInputStatusParamsClassOnce sync.Once
)

func getMTRMediaInputClusterShowInputStatusParamsClass() _MTRMediaInputClusterShowInputStatusParamsClass {
	MTRMediaInputClusterShowInputStatusParamsClassOnce.Do(func() {
		MTRMediaInputClusterShowInputStatusParamsClass = _MTRMediaInputClusterShowInputStatusParamsClass{objc.GetClass("MTRMediaInputClusterShowInputStatusParams")}
	})
	return MTRMediaInputClusterShowInputStatusParamsClass
}

type _MTRMediaInputClusterShowInputStatusParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterShowInputStatusParams] class.
type IMTRMediaInputClusterShowInputStatusParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterShowInputStatusParams
type MTRMediaInputClusterShowInputStatusParams struct {
	objectivec.Object
}

// MTRMediaInputClusterShowInputStatusParamsFrom constructs a [MTRMediaInputClusterShowInputStatusParams] from an unsafe.Pointer.
func MTRMediaInputClusterShowInputStatusParamsFrom(ptr unsafe.Pointer) MTRMediaInputClusterShowInputStatusParams {
	return MTRMediaInputClusterShowInputStatusParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterShowInputStatusParamsClass) Alloc() MTRMediaInputClusterShowInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterShowInputStatusParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterShowInputStatusParamsClass) New() MTRMediaInputClusterShowInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterShowInputStatusParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterShowInputStatusParams) Init() MTRMediaInputClusterShowInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterShowInputStatusParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterShowInputStatusParams) Autorelease() MTRMediaInputClusterShowInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterShowInputStatusParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterShowInputStatusParams creates a new MTRMediaInputClusterShowInputStatusParams instance.
func NewMTRMediaInputClusterShowInputStatusParams() MTRMediaInputClusterShowInputStatusParams {
	return getMTRMediaInputClusterShowInputStatusParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclustershowinputstatusparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterShowInputStatusParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclustershowinputstatusparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterShowInputStatusParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclustershowinputstatusparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterShowInputStatusParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclustershowinputstatusparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterShowInputStatusParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



