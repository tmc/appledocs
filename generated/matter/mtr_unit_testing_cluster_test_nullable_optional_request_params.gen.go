// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestNullableOptionalRequestParams] class.
var (
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClass     _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNullableOptionalRequestParamsClass() _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass {
	MTRUnitTestingClusterTestNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNullableOptionalRequestParamsClass = _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNullableOptionalRequestParams")}
	})
	return MTRUnitTestingClusterTestNullableOptionalRequestParamsClass
}

type _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNullableOptionalRequestParams] class.
type IMTRUnitTestingClusterTestNullableOptionalRequestParams interface {
	objectivec.IObject
	Arg1() foundation.Number
	SetArg1(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNullableOptionalRequestParams
type MTRUnitTestingClusterTestNullableOptionalRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom constructs a [MTRUnitTestingClusterTestNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNullableOptionalRequestParams {
	return MTRUnitTestingClusterTestNullableOptionalRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass) Alloc() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNullableOptionalRequestParamsClass) New() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Init() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Autorelease() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNullableOptionalRequestParams creates a new MTRUnitTestingClusterTestNullableOptionalRequestParams instance.
func NewMTRUnitTestingClusterTestNullableOptionalRequestParams() MTRUnitTestingClusterTestNullableOptionalRequestParams {
	return getMTRUnitTestingClusterTestNullableOptionalRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



