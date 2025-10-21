// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccountLoginClusterLogoutParams] class.
var (
	MTRAccountLoginClusterLogoutParamsClass     _MTRAccountLoginClusterLogoutParamsClass
	MTRAccountLoginClusterLogoutParamsClassOnce sync.Once
)

func getMTRAccountLoginClusterLogoutParamsClass() _MTRAccountLoginClusterLogoutParamsClass {
	MTRAccountLoginClusterLogoutParamsClassOnce.Do(func() {
		MTRAccountLoginClusterLogoutParamsClass = _MTRAccountLoginClusterLogoutParamsClass{objc.GetClass("MTRAccountLoginClusterLogoutParams")}
	})
	return MTRAccountLoginClusterLogoutParamsClass
}

type _MTRAccountLoginClusterLogoutParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterLogoutParams] class.
type IMTRAccountLoginClusterLogoutParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLogoutParams
type MTRAccountLoginClusterLogoutParams struct {
	objectivec.Object
}

// MTRAccountLoginClusterLogoutParamsFrom constructs a [MTRAccountLoginClusterLogoutParams] from an unsafe.Pointer.
func MTRAccountLoginClusterLogoutParamsFrom(ptr unsafe.Pointer) MTRAccountLoginClusterLogoutParams {
	return MTRAccountLoginClusterLogoutParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterLogoutParamsClass) Alloc() MTRAccountLoginClusterLogoutParams {
	rv := objc.Send[MTRAccountLoginClusterLogoutParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterLogoutParamsClass) New() MTRAccountLoginClusterLogoutParams {
	rv := objc.Send[MTRAccountLoginClusterLogoutParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterLogoutParams) Init() MTRAccountLoginClusterLogoutParams {
	rv := objc.Send[MTRAccountLoginClusterLogoutParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterLogoutParams) Autorelease() MTRAccountLoginClusterLogoutParams {
	rv := objc.Send[MTRAccountLoginClusterLogoutParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterLogoutParams creates a new MTRAccountLoginClusterLogoutParams instance.
func NewMTRAccountLoginClusterLogoutParams() MTRAccountLoginClusterLogoutParams {
	return getMTRAccountLoginClusterLogoutParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/node
func (m_ MTRAccountLoginClusterLogoutParams) Node() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("node"))
	return rv
}


// SetNode sets the value of the node property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/node
func (m_ MTRAccountLoginClusterLogoutParams) SetNode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLogoutParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLogoutParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLogoutParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLogoutParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



