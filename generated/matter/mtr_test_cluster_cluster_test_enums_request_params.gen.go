// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestEnumsRequestParams] class.
var (
	MTRTestClusterClusterTestEnumsRequestParamsClass     _MTRTestClusterClusterTestEnumsRequestParamsClass
	MTRTestClusterClusterTestEnumsRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEnumsRequestParamsClass() _MTRTestClusterClusterTestEnumsRequestParamsClass {
	MTRTestClusterClusterTestEnumsRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEnumsRequestParamsClass = _MTRTestClusterClusterTestEnumsRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestEnumsRequestParams")}
	})
	return MTRTestClusterClusterTestEnumsRequestParamsClass
}

type _MTRTestClusterClusterTestEnumsRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEnumsRequestParams] class.
type IMTRTestClusterClusterTestEnumsRequestParams interface {
	IMTRUnitTestingClusterTestEnumsRequestParams
	Arg1() foundation.Number
	SetArg1(value foundation.INumber)
	Arg2() foundation.Number
	SetArg2(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEnumsRequestParams
type MTRTestClusterClusterTestEnumsRequestParams struct {
	MTRUnitTestingClusterTestEnumsRequestParams
}

// MTRTestClusterClusterTestEnumsRequestParamsFrom constructs a [MTRTestClusterClusterTestEnumsRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEnumsRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEnumsRequestParams {
	return MTRTestClusterClusterTestEnumsRequestParams{
		MTRUnitTestingClusterTestEnumsRequestParams: MTRUnitTestingClusterTestEnumsRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEnumsRequestParamsClass) Alloc() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEnumsRequestParamsClass) New() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEnumsRequestParams) Init() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEnumsRequestParams) Autorelease() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEnumsRequestParams creates a new MTRTestClusterClusterTestEnumsRequestParams instance.
func NewMTRTestClusterClusterTestEnumsRequestParams() MTRTestClusterClusterTestEnumsRequestParams {
	return getMTRTestClusterClusterTestEnumsRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/arg1
func (m_ MTRTestClusterClusterTestEnumsRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/arg1
func (m_ MTRTestClusterClusterTestEnumsRequestParams) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/arg2
func (m_ MTRTestClusterClusterTestEnumsRequestParams) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/arg2
func (m_ MTRTestClusterClusterTestEnumsRequestParams) SetArg2(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestEnumsRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestEnumsRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEnumsRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEnumsRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



