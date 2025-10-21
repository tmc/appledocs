// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetReadAllIndicesParams] class.
var (
	MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass     _MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass
	MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass() _MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass {
	MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass = _MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetReadAllIndicesParams")}
	})
	return MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass
}

type _MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetReadAllIndicesParams] class.
type IMTRGroupKeyManagementClusterKeySetReadAllIndicesParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetReadAllIndicesParams
type MTRGroupKeyManagementClusterKeySetReadAllIndicesParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetReadAllIndicesParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetReadAllIndicesParams {
	return MTRGroupKeyManagementClusterKeySetReadAllIndicesParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetReadAllIndicesParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass) New() MTRGroupKeyManagementClusterKeySetReadAllIndicesParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) Init() MTRGroupKeyManagementClusterKeySetReadAllIndicesParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) Autorelease() MTRGroupKeyManagementClusterKeySetReadAllIndicesParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetReadAllIndicesParams creates a new MTRGroupKeyManagementClusterKeySetReadAllIndicesParams instance.
func NewMTRGroupKeyManagementClusterKeySetReadAllIndicesParams() MTRGroupKeyManagementClusterKeySetReadAllIndicesParams {
	return getMTRGroupKeyManagementClusterKeySetReadAllIndicesParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadallindicesparams/groupkeysetids
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) GroupKeySetIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupKeySetIDs"))
	return rv
}


// SetGroupKeySetIDs sets the value of the groupKeySetIDs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadallindicesparams/groupkeysetids
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) SetGroupKeySetIDs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySetIDs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadallindicesparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadallindicesparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadallindicesparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadallindicesparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



