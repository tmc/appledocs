// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestListStructArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestListStructArgumentRequestParamsClass     _MTRTestClusterClusterTestListStructArgumentRequestParamsClass
	MTRTestClusterClusterTestListStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestListStructArgumentRequestParamsClass() _MTRTestClusterClusterTestListStructArgumentRequestParamsClass {
	MTRTestClusterClusterTestListStructArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestListStructArgumentRequestParamsClass = _MTRTestClusterClusterTestListStructArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestListStructArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestListStructArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestListStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestListStructArgumentRequestParams] class.
type IMTRTestClusterClusterTestListStructArgumentRequestParams interface {
	IMTRUnitTestingClusterTestListStructArgumentRequestParams
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestListStructArgumentRequestParams
type MTRTestClusterClusterTestListStructArgumentRequestParams struct {
	MTRUnitTestingClusterTestListStructArgumentRequestParams
}

// MTRTestClusterClusterTestListStructArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestListStructArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestListStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestListStructArgumentRequestParams {
	return MTRTestClusterClusterTestListStructArgumentRequestParams{
		MTRUnitTestingClusterTestListStructArgumentRequestParams: MTRUnitTestingClusterTestListStructArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestListStructArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestListStructArgumentRequestParamsClass) New() MTRTestClusterClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) Init() MTRTestClusterClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) Autorelease() MTRTestClusterClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestListStructArgumentRequestParams creates a new MTRTestClusterClusterTestListStructArgumentRequestParams instance.
func NewMTRTestClusterClusterTestListStructArgumentRequestParams() MTRTestClusterClusterTestListStructArgumentRequestParams {
	return getMTRTestClusterClusterTestListStructArgumentRequestParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestliststructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
