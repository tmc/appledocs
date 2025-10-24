// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEmitTestEventRequestParams] class.
var (
	MTRUnitTestingClusterTestEmitTestEventRequestParamsClass     _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass
	MTRUnitTestingClusterTestEmitTestEventRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEmitTestEventRequestParamsClass() _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass {
	MTRUnitTestingClusterTestEmitTestEventRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEmitTestEventRequestParamsClass = _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestEmitTestEventRequestParams")}
	})
	return MTRUnitTestingClusterTestEmitTestEventRequestParamsClass
}

type _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEmitTestEventRequestParams] class.
type IMTRUnitTestingClusterTestEmitTestEventRequestParams interface {
	objectivec.IObject
	// properties:
	Arg1() objc.IObject /* cross-framework: NSNumber */
	SetArg1(value objc.IObject /* cross-framework: NSNumber */)
	Arg2() objc.IObject /* cross-framework: NSNumber */
	SetArg2(value objc.IObject /* cross-framework: NSNumber */)
	Arg3() objc.IObject /* cross-framework: NSNumber */
	SetArg3(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEmitTestEventRequestParams
type MTRUnitTestingClusterTestEmitTestEventRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEmitTestEventRequestParamsFrom constructs a [MTRUnitTestingClusterTestEmitTestEventRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEmitTestEventRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEmitTestEventRequestParams {
	return MTRUnitTestingClusterTestEmitTestEventRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass) Alloc() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass) New() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Init() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Autorelease() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEmitTestEventRequestParams creates a new MTRUnitTestingClusterTestEmitTestEventRequestParams instance.
func NewMTRUnitTestingClusterTestEmitTestEventRequestParams() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	return getMTRUnitTestingClusterTestEmitTestEventRequestParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Arg1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg1"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) SetArg1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/arg2
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Arg2() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg2"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/arg2
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) SetArg2(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/arg3
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Arg3() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg3"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/arg3
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) SetArg3(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
