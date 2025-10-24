// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetReadResponseParams] class.
var (
	MTRGroupKeyManagementClusterKeySetReadResponseParamsClass     _MTRGroupKeyManagementClusterKeySetReadResponseParamsClass
	MTRGroupKeyManagementClusterKeySetReadResponseParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetReadResponseParamsClass() _MTRGroupKeyManagementClusterKeySetReadResponseParamsClass {
	MTRGroupKeyManagementClusterKeySetReadResponseParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetReadResponseParamsClass = _MTRGroupKeyManagementClusterKeySetReadResponseParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetReadResponseParams")}
	})
	return MTRGroupKeyManagementClusterKeySetReadResponseParamsClass
}

type _MTRGroupKeyManagementClusterKeySetReadResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetReadResponseParams] class.
type IMTRGroupKeyManagementClusterKeySetReadResponseParams interface {
	objectivec.IObject
	// properties:
	GroupKeySet() IMTRGroupKeyManagementClusterGroupKeySetStruct
	SetGroupKeySet(value IMTRGroupKeyManagementClusterGroupKeySetStruct)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetReadResponseParams
type MTRGroupKeyManagementClusterKeySetReadResponseParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetReadResponseParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetReadResponseParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetReadResponseParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetReadResponseParams {
	return MTRGroupKeyManagementClusterKeySetReadResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetReadResponseParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetReadResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetReadResponseParamsClass) New() MTRGroupKeyManagementClusterKeySetReadResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetReadResponseParams) Init() MTRGroupKeyManagementClusterKeySetReadResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetReadResponseParams) Autorelease() MTRGroupKeyManagementClusterKeySetReadResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetReadResponseParams creates a new MTRGroupKeyManagementClusterKeySetReadResponseParams instance.
func NewMTRGroupKeyManagementClusterKeySetReadResponseParams() MTRGroupKeyManagementClusterKeySetReadResponseParams {
	return getMTRGroupKeyManagementClusterKeySetReadResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadresponseparams/groupkeyset
func (m_ MTRGroupKeyManagementClusterKeySetReadResponseParams) GroupKeySet() IMTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](m_.ID, objc.Sel("groupKeySet"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadresponseparams/groupkeyset
func (m_ MTRGroupKeyManagementClusterKeySetReadResponseParams) SetGroupKeySet(value IMTRGroupKeyManagementClusterGroupKeySetStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySet:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadresponseparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetReadResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclusterkeysetreadresponseparams/timedinvoketimeoutms
func (m_ MTRGroupKeyManagementClusterKeySetReadResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



