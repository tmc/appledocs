// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestEmitTestEventRequestParams] class.
var (
	MTRTestClusterClusterTestEmitTestEventRequestParamsClass     _MTRTestClusterClusterTestEmitTestEventRequestParamsClass
	MTRTestClusterClusterTestEmitTestEventRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEmitTestEventRequestParamsClass() _MTRTestClusterClusterTestEmitTestEventRequestParamsClass {
	MTRTestClusterClusterTestEmitTestEventRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEmitTestEventRequestParamsClass = _MTRTestClusterClusterTestEmitTestEventRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestEmitTestEventRequestParams")}
	})
	return MTRTestClusterClusterTestEmitTestEventRequestParamsClass
}

type _MTRTestClusterClusterTestEmitTestEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEmitTestEventRequestParams] class.
type IMTRTestClusterClusterTestEmitTestEventRequestParams interface {
	IMTRUnitTestingClusterTestEmitTestEventRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEmitTestEventRequestParams
type MTRTestClusterClusterTestEmitTestEventRequestParams struct {
	MTRUnitTestingClusterTestEmitTestEventRequestParams
}

// MTRTestClusterClusterTestEmitTestEventRequestParamsFrom constructs a [MTRTestClusterClusterTestEmitTestEventRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEmitTestEventRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEmitTestEventRequestParams {
	return MTRTestClusterClusterTestEmitTestEventRequestParams{
		MTRUnitTestingClusterTestEmitTestEventRequestParams: MTRUnitTestingClusterTestEmitTestEventRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEmitTestEventRequestParamsClass) Alloc() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEmitTestEventRequestParamsClass) New() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Init() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Autorelease() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEmitTestEventRequestParams creates a new MTRTestClusterClusterTestEmitTestEventRequestParams instance.
func NewMTRTestClusterClusterTestEmitTestEventRequestParams() MTRTestClusterClusterTestEmitTestEventRequestParams {
	return getMTRTestClusterClusterTestEmitTestEventRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/arg1
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/arg1
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) SetArg1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/arg2
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/arg2
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) SetArg2(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/arg3
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Arg3() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg3"))
	return rv
}


// SetArg3 sets the value of the arg3 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/arg3
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) SetArg3(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittesteventrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



