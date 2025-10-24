// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestNestedStructListArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass     _MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass
	MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass() _MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass {
	MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass = _MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestNestedStructListArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestNestedStructListArgumentRequestParams] class.
type IMTRTestClusterClusterTestNestedStructListArgumentRequestParams interface {
	IMTRUnitTestingClusterTestNestedStructListArgumentRequestParams
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestNestedStructListArgumentRequestParams
type MTRTestClusterClusterTestNestedStructListArgumentRequestParams struct {
	MTRUnitTestingClusterTestNestedStructListArgumentRequestParams
}

// MTRTestClusterClusterTestNestedStructListArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestNestedStructListArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestNestedStructListArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestNestedStructListArgumentRequestParams {
	return MTRTestClusterClusterTestNestedStructListArgumentRequestParams{
		MTRUnitTestingClusterTestNestedStructListArgumentRequestParams: MTRUnitTestingClusterTestNestedStructListArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass) New() MTRTestClusterClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) Init() MTRTestClusterClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructListArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) Autorelease() MTRTestClusterClusterTestNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructListArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestNestedStructListArgumentRequestParams creates a new MTRTestClusterClusterTestNestedStructListArgumentRequestParams instance.
func NewMTRTestClusterClusterTestNestedStructListArgumentRequestParams() MTRTestClusterClusterTestNestedStructListArgumentRequestParams {
	return getMTRTestClusterClusterTestNestedStructListArgumentRequestParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) Arg1() IMTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("arg1"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterNestedStructList) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
