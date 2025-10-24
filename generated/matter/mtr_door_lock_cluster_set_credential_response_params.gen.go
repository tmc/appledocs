// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetCredentialResponseParams] class.
var (
	MTRDoorLockClusterSetCredentialResponseParamsClass     _MTRDoorLockClusterSetCredentialResponseParamsClass
	MTRDoorLockClusterSetCredentialResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetCredentialResponseParamsClass() _MTRDoorLockClusterSetCredentialResponseParamsClass {
	MTRDoorLockClusterSetCredentialResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetCredentialResponseParamsClass = _MTRDoorLockClusterSetCredentialResponseParamsClass{objc.GetClass("MTRDoorLockClusterSetCredentialResponseParams")}
	})
	return MTRDoorLockClusterSetCredentialResponseParamsClass
}

type _MTRDoorLockClusterSetCredentialResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetCredentialResponseParams] class.
type IMTRDoorLockClusterSetCredentialResponseParams interface {
	objectivec.IObject
	// properties:
	NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */
	SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialResponseParams
type MTRDoorLockClusterSetCredentialResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetCredentialResponseParamsFrom constructs a [MTRDoorLockClusterSetCredentialResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetCredentialResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetCredentialResponseParams {
	return MTRDoorLockClusterSetCredentialResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetCredentialResponseParamsClass) Alloc() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetCredentialResponseParamsClass) New() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetCredentialResponseParams) Init() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetCredentialResponseParams) Autorelease() MTRDoorLockClusterSetCredentialResponseParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetCredentialResponseParams creates a new MTRDoorLockClusterSetCredentialResponseParams instance.
func NewMTRDoorLockClusterSetCredentialResponseParams() MTRDoorLockClusterSetCredentialResponseParams {
	return getMTRDoorLockClusterSetCredentialResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/nextcredentialindex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextCredentialIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/nextcredentialindex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextCredentialIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/status
func (m_ MTRDoorLockClusterSetCredentialResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/status
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetCredentialResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/userindex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustersetcredentialresponseparams/userindex
func (m_ MTRDoorLockClusterSetCredentialResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



