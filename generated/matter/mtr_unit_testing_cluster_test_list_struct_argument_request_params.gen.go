// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestListStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestListStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestListStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestListStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestListStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestListStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestListStructArgumentRequestParams interface {
	objectivec.IObject
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListStructArgumentRequestParams
type MTRUnitTestingClusterTestListStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestListStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListStructArgumentRequestParams {
	return MTRUnitTestingClusterTestListStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestListStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) Init() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListStructArgumentRequestParams creates a new MTRUnitTestingClusterTestListStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestListStructArgumentRequestParams() MTRUnitTestingClusterTestListStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestListStructArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestliststructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestListStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



