// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass     _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass() _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass {
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass = _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArrayArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] class.
type IMTRUnitTestingClusterTestStructArrayArgumentRequestParams interface {
	objectivec.IObject
	// properties:
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	Arg2() unsafe.Pointer
	SetArg2(value unsafe.Pointer)
	Arg3() unsafe.Pointer
	SetArg3(value unsafe.Pointer)
	Arg4() unsafe.Pointer
	SetArg4(value unsafe.Pointer)
	Arg5() objc.IObject /* cross-framework: NSNumber */
	SetArg5(value objc.IObject /* cross-framework: NSNumber */)
	Arg6() objc.IObject /* cross-framework: NSNumber */
	SetArg6(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams
type MTRUnitTestingClusterTestStructArrayArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	return MTRUnitTestingClusterTestStructArrayArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass) New() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Init() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArrayArgumentRequestParams creates a new MTRUnitTestingClusterTestStructArrayArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestStructArrayArgumentRequestParams() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	return getMTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg2() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg2"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg2(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg3() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg3"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg3(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg4() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg4"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg4(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg5() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg5"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg5(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Arg6() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg6"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetArg6(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



