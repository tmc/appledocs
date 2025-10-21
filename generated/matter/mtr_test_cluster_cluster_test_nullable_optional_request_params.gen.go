// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestNullableOptionalRequestParams] class.
var (
	MTRTestClusterClusterTestNullableOptionalRequestParamsClass     _MTRTestClusterClusterTestNullableOptionalRequestParamsClass
	MTRTestClusterClusterTestNullableOptionalRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestNullableOptionalRequestParamsClass() _MTRTestClusterClusterTestNullableOptionalRequestParamsClass {
	MTRTestClusterClusterTestNullableOptionalRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestNullableOptionalRequestParamsClass = _MTRTestClusterClusterTestNullableOptionalRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestNullableOptionalRequestParams")}
	})
	return MTRTestClusterClusterTestNullableOptionalRequestParamsClass
}

type _MTRTestClusterClusterTestNullableOptionalRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestNullableOptionalRequestParams] class.
type IMTRTestClusterClusterTestNullableOptionalRequestParams interface {
	IMTRUnitTestingClusterTestNullableOptionalRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestNullableOptionalRequestParams
type MTRTestClusterClusterTestNullableOptionalRequestParams struct {
	MTRUnitTestingClusterTestNullableOptionalRequestParams
}

// MTRTestClusterClusterTestNullableOptionalRequestParamsFrom constructs a [MTRTestClusterClusterTestNullableOptionalRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestNullableOptionalRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestNullableOptionalRequestParams {
	return MTRTestClusterClusterTestNullableOptionalRequestParams{
		MTRUnitTestingClusterTestNullableOptionalRequestParams: MTRUnitTestingClusterTestNullableOptionalRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestNullableOptionalRequestParamsClass) Alloc() MTRTestClusterClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestNullableOptionalRequestParamsClass) New() MTRTestClusterClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) Init() MTRTestClusterClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) Autorelease() MTRTestClusterClusterTestNullableOptionalRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNullableOptionalRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestNullableOptionalRequestParams creates a new MTRTestClusterClusterTestNullableOptionalRequestParams instance.
func NewMTRTestClusterClusterTestNullableOptionalRequestParams() MTRTestClusterClusterTestNullableOptionalRequestParams {
	return getMTRTestClusterClusterTestNullableOptionalRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalrequestparams/arg1
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalrequestparams/arg1
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnullableoptionalrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNullableOptionalRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



