// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestStructArrayArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass     _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass
	MTRTestClusterClusterTestStructArrayArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestStructArrayArgumentRequestParamsClass() _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass {
	MTRTestClusterClusterTestStructArrayArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass = _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestStructArrayArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestStructArrayArgumentRequestParams] class.
type IMTRTestClusterClusterTestStructArrayArgumentRequestParams interface {
	IMTRUnitTestingClusterTestStructArrayArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestStructArrayArgumentRequestParams
type MTRTestClusterClusterTestStructArrayArgumentRequestParams struct {
	MTRUnitTestingClusterTestStructArrayArgumentRequestParams
}

// MTRTestClusterClusterTestStructArrayArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestStructArrayArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestStructArrayArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	return MTRTestClusterClusterTestStructArrayArgumentRequestParams{
		MTRUnitTestingClusterTestStructArrayArgumentRequestParams: MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass) New() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Init() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Autorelease() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestStructArrayArgumentRequestParams creates a new MTRTestClusterClusterTestStructArrayArgumentRequestParams instance.
func NewMTRTestClusterClusterTestStructArrayArgumentRequestParams() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	return getMTRTestClusterClusterTestStructArrayArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg2
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Arg2() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg2
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetArg2(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg3
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Arg3() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg3"))
	return rv
}


// SetArg3 sets the value of the arg3 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg3
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetArg3(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg4
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Arg4() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg4"))
	return rv
}


// SetArg4 sets the value of the arg4 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg4
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetArg4(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg5
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Arg5() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg5"))
	return rv
}


// SetArg5 sets the value of the arg5 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg5
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetArg5(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg6
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Arg6() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg6"))
	return rv
}


// SetArg6 sets the value of the arg6 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/arg6
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetArg6(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



