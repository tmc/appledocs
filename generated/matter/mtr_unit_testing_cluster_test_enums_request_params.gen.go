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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetArg1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg2
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/arg2
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetArg2(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEnumsRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



