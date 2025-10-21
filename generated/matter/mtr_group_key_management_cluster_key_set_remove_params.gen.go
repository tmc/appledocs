// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetRemoveParams] class.
var (
	MTRGroupKeyManagementClusterKeySetRemoveParamsClass     _MTRGroupKeyManagementClusterKeySetRemoveParamsClass
	MTRGroupKeyManagementClusterKeySetRemoveParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetRemoveParamsClass() _MTRGroupKeyManagementClusterKeySetRemoveParamsClass {
	MTRGroupKeyManagementClusterKeySetRemoveParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetRemoveParamsClass = _MTRGroupKeyManagementClusterKeySetRemoveParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetRemoveParams")}
	})
	return MTRGroupKeyManagementClusterKeySetRemoveParamsClass
}

type _MTRGroupKeyManagementClusterKeySetRemoveParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetRemoveParams] class.
type IMTRGroupKeyManagementClusterKeySetRemoveParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetRemoveParams
type MTRGroupKeyManagementClusterKeySetRemoveParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetRemoveParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetRemoveParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetRemoveParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetRemoveParams {
	return MTRGroupKeyManagementClusterKeySetRemoveParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetRemoveParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetRemoveParamsClass) New() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) Init() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) Autorelease() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetRemoveParams creates a new MTRGroupKeyManagementClusterKeySetRemoveParams instance.
func NewMTRGroupKeyManagementClusterKeySetRemoveParams() MTRGroupKeyManagementClusterKeySetRemoveParams {
	return getMTRGroupKeyManagementClusterKeySetRemoveParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetremoveparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetremoveparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetremoveparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetremoveparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetremoveparams/groupkeysetid
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) GroupKeySetID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupKeySetID"))
	return rv
}


// SetGroupKeySetID sets the value of the groupKeySetID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetremoveparams/groupkeysetid
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) SetGroupKeySetID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySetID:"), value)
}



