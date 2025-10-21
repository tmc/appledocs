// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructlistargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNestedStructListArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



