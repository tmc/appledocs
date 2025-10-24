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
	// properties:
	Node() objc.IObject /* cross-framework: NSNumber */
	SetNode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/node
func (m_ MTRAccountLoginClusterLogoutParams) Node() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("node"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/node
func (m_ MTRAccountLoginClusterLogoutParams) SetNode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLogoutParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/serversideprocessingtimeout
func (m_ MTRAccountLoginClusterLogoutParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLogoutParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccountloginclusterlogoutparams/timedinvoketimeoutms
func (m_ MTRAccountLoginClusterLogoutParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



