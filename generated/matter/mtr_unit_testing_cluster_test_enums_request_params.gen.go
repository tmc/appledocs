// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEnumsRequestParams] class.
var (
	MTRUnitTestingClusterTestEnumsRequestParamsClass     _MTRUnitTestingClusterTestEnumsRequestParamsClass
	MTRUnitTestingClusterTestEnumsRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEnumsRequestParamsClass() _MTRUnitTestingClusterTestEnumsRequestParamsClass {
	MTRUnitTestingClusterTestEnumsRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEnumsRequestParamsClass = _MTRUnitTestingClusterTestEnumsRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestEnumsRequestParams")}
	})
	return MTRUnitTestingClusterTestEnumsRequestParamsClass
}

type _MTRUnitTestingClusterTestEnumsRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEnumsRequestParams] class.
type IMTRUnitTestingClusterTestEnumsRequestParams interface {
	objectivec.IObject
	// properties:
	Arg1() objc.IObject /* cross-framework: NSNumber */
	SetArg1(value objc.IObject /* cross-framework: NSNumber */)
	Arg2() objc.IObject /* cross-framework: NSNumber */
	SetArg2(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEnumsRequestParams
type MTRUnitTestingClusterTestEnumsRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEnumsRequestParamsFrom constructs a [MTRUnitTestingClusterTestEnumsRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEnumsRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEnumsRequestParams {
	return MTRUnitTestingClusterTestEnumsRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEnumsRequestParamsClass) Alloc() MTRUnitTestingClusterTestEnumsRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEnumsRequestParamsClass) New() MTRUnitTestingClusterTestEnumsRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) Init() MTRUnitTestingClusterTestEnumsRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) Autorelease() MTRUnitTestingClusterTestEnumsRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEnumsRequestParams creates a new MTRUnitTestingClusterTestEnumsRequestParams instance.
func NewMTRUnitTestingClusterTestEnumsRequestParams() MTRUnitTestingClusterTestEnumsRequestParams {
	return getMTRUnitTestingClusterTestEnumsRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) Arg1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetArg1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg2
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) Arg2() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg2"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg2
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetArg2(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



