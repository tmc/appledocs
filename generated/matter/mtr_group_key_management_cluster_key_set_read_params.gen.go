// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetReadParams] class.
var (
	MTRGroupKeyManagementClusterKeySetReadParamsClass     _MTRGroupKeyManagementClusterKeySetReadParamsClass
	MTRGroupKeyManagementClusterKeySetReadParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetReadParamsClass() _MTRGroupKeyManagementClusterKeySetReadParamsClass {
	MTRGroupKeyManagementClusterKeySetReadParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetReadParamsClass = _MTRGroupKeyManagementClusterKeySetReadParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetReadParams")}
	})
	return MTRGroupKeyManagementClusterKeySetReadParamsClass
}

type _MTRGroupKeyManagementClusterKeySetReadParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetReadParams] class.
type IMTRGroupKeyManagementClusterKeySetReadParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetReadParams
type MTRGroupKeyManagementClusterKeySetReadParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetReadParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetReadParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetReadParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetReadParams {
	return MTRGroupKeyManagementClusterKeySetReadParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetReadParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetReadParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetReadParamsClass) New() MTRGroupKeyManagementClusterKeySetReadParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) Init() MTRGroupKeyManagementClusterKeySetReadParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) Autorelease() MTRGroupKeyManagementClusterKeySetReadParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetReadParams creates a new MTRGroupKeyManagementClusterKeySetReadParams instance.
func NewMTRGroupKeyManagementClusterKeySetReadParams() MTRGroupKeyManagementClusterKeySetReadParams {
	return getMTRGroupKeyManagementClusterKeySetReadParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadparams/groupkeysetid
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) GroupKeySetID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupKeySetID"))
	return rv
}


// SetGroupKeySetID sets the value of the groupKeySetID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadparams/groupkeysetid
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) SetGroupKeySetID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySetID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetReadParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



