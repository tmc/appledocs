// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterClearUserParams] class.
var (
	MTRDoorLockClusterClearUserParamsClass     _MTRDoorLockClusterClearUserParamsClass
	MTRDoorLockClusterClearUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearUserParamsClass() _MTRDoorLockClusterClearUserParamsClass {
	MTRDoorLockClusterClearUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearUserParamsClass = _MTRDoorLockClusterClearUserParamsClass{objc.GetClass("MTRDoorLockClusterClearUserParams")}
	})
	return MTRDoorLockClusterClearUserParamsClass
}

type _MTRDoorLockClusterClearUserParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearUserParams] class.
type IMTRDoorLockClusterClearUserParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearUserParams
type MTRDoorLockClusterClearUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearUserParamsFrom constructs a [MTRDoorLockClusterClearUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearUserParams {
	return MTRDoorLockClusterClearUserParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearUserParamsClass) Alloc() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearUserParamsClass) New() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearUserParams) Init() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearUserParams) Autorelease() MTRDoorLockClusterClearUserParams {
	rv := objc.Send[MTRDoorLockClusterClearUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearUserParams creates a new MTRDoorLockClusterClearUserParams instance.
func NewMTRDoorLockClusterClearUserParams() MTRDoorLockClusterClearUserParams {
	return getMTRDoorLockClusterClearUserParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearUserParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterClearUserParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearUserParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterClearUserParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/userindex
func (m_ MTRDoorLockClusterClearUserParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterclearuserparams/userindex
func (m_ MTRDoorLockClusterClearUserParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



