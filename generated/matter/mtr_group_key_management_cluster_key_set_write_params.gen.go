// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetWriteParams] class.
var (
	MTRGroupKeyManagementClusterKeySetWriteParamsClass     _MTRGroupKeyManagementClusterKeySetWriteParamsClass
	MTRGroupKeyManagementClusterKeySetWriteParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetWriteParamsClass() _MTRGroupKeyManagementClusterKeySetWriteParamsClass {
	MTRGroupKeyManagementClusterKeySetWriteParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetWriteParamsClass = _MTRGroupKeyManagementClusterKeySetWriteParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetWriteParams")}
	})
	return MTRGroupKeyManagementClusterKeySetWriteParamsClass
}

type _MTRGroupKeyManagementClusterKeySetWriteParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetWriteParams] class.
type IMTRGroupKeyManagementClusterKeySetWriteParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetWriteParams
type MTRGroupKeyManagementClusterKeySetWriteParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetWriteParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetWriteParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetWriteParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetWriteParams {
	return MTRGroupKeyManagementClusterKeySetWriteParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetWriteParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetWriteParamsClass) New() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) Init() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) Autorelease() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetWriteParams creates a new MTRGroupKeyManagementClusterKeySetWriteParams instance.
func NewMTRGroupKeyManagementClusterKeySetWriteParams() MTRGroupKeyManagementClusterKeySetWriteParams {
	return getMTRGroupKeyManagementClusterKeySetWriteParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetwriteparams/groupkeyset
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) GroupKeySet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupKeySet"))
	return rv
}


// SetGroupKeySet sets the value of the groupKeySet property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetwriteparams/groupkeyset
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) SetGroupKeySet(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySet:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetwriteparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetwriteparams/serversideprocessingtimeout
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetwriteparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetwriteparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



