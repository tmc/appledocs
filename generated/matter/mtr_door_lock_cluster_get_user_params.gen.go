// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetUserParams] class.
var (
	MTRDoorLockClusterGetUserParamsClass     _MTRDoorLockClusterGetUserParamsClass
	MTRDoorLockClusterGetUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetUserParamsClass() _MTRDoorLockClusterGetUserParamsClass {
	MTRDoorLockClusterGetUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetUserParamsClass = _MTRDoorLockClusterGetUserParamsClass{objc.GetClass("MTRDoorLockClusterGetUserParams")}
	})
	return MTRDoorLockClusterGetUserParamsClass
}

type _MTRDoorLockClusterGetUserParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetUserParams] class.
type IMTRDoorLockClusterGetUserParams interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserParams
type MTRDoorLockClusterGetUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetUserParamsFrom constructs a [MTRDoorLockClusterGetUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetUserParams {
	return MTRDoorLockClusterGetUserParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetUserParamsClass) Alloc() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetUserParamsClass) New() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetUserParams) Init() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetUserParams) Autorelease() MTRDoorLockClusterGetUserParams {
	rv := objc.Send[MTRDoorLockClusterGetUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetUserParams creates a new MTRDoorLockClusterGetUserParams instance.
func NewMTRDoorLockClusterGetUserParams() MTRDoorLockClusterGetUserParams {
	return getMTRDoorLockClusterGetUserParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetUserParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/serversideprocessingtimeout
func (m_ MTRDoorLockClusterGetUserParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetUserParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/userindex
func (m_ MTRDoorLockClusterGetUserParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetuserparams/userindex
func (m_ MTRDoorLockClusterGetUserParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



