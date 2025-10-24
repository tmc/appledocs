// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass     _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass() _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass {
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass = _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams] class.
type IMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams interface {
	objectivec.IObject
	// properties:
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams
type MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	return MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass) New() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) Init() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams creates a new MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams() MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams {
	return getMTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistnestedstructlistargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistnestedstructlistargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



