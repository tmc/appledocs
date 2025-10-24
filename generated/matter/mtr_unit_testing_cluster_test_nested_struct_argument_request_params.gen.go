// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNestedStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestNestedStructArgumentRequestParams interface {
	objectivec.IObject
	// properties:
	Arg1() IMTRUnitTestingClusterNestedStruct
	SetArg1(value IMTRUnitTestingClusterNestedStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams
type MTRUnitTestingClusterTestNestedStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	return MTRUnitTestingClusterTestNestedStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Init() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNestedStructArgumentRequestParams creates a new MTRUnitTestingClusterTestNestedStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestNestedStructArgumentRequestParams() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Arg1() IMTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterNestedStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



