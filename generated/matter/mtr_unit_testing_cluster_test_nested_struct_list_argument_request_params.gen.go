// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestNestedStructListArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass     _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass() _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass {
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass = _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNestedStructListArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNestedStructListArgumentRequestParams] class.
type IMTRUnitTestingClusterTestNestedStructListArgumentRequestParams interface {
	objectivec.IObject
	// properties:
	Arg1() IMTRUnitTestingClusterNestedStructList
	SetArg1(value IMTRUnitTestingClusterNestedStructList)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructListArgumentRequestParams
type MTRUnitTestingClusterTestNestedStructListArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestNestedStructListArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	return MTRUnitTestingClusterTestNestedStructListArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass) New() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) Init() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructListArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNestedStructListArgumentRequestParams creates a new MTRUnitTestingClusterTestNestedStructListArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestNestedStructListArgumentRequestParams() MTRUnitTestingClusterTestNestedStructListArgumentRequestParams {
	return getMTRUnitTestingClusterTestNestedStructListArgumentRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructlistargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) Arg1() IMTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructlistargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterNestedStructList) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNestedStructListArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



