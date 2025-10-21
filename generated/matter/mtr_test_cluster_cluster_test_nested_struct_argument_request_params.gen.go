// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestNestedStructArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass     _MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass
	MTRTestClusterClusterTestNestedStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestNestedStructArgumentRequestParamsClass() _MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass {
	MTRTestClusterClusterTestNestedStructArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass = _MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestNestedStructArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestNestedStructArgumentRequestParams] class.
type IMTRTestClusterClusterTestNestedStructArgumentRequestParams interface {
	IMTRUnitTestingClusterTestNestedStructArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestNestedStructArgumentRequestParams
type MTRTestClusterClusterTestNestedStructArgumentRequestParams struct {
	MTRUnitTestingClusterTestNestedStructArgumentRequestParams
}

// MTRTestClusterClusterTestNestedStructArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestNestedStructArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestNestedStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestNestedStructArgumentRequestParams {
	return MTRTestClusterClusterTestNestedStructArgumentRequestParams{
		MTRUnitTestingClusterTestNestedStructArgumentRequestParams: MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestNestedStructArgumentRequestParamsClass) New() MTRTestClusterClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) Init() MTRTestClusterClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) Autorelease() MTRTestClusterClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestNestedStructArgumentRequestParams creates a new MTRTestClusterClusterTestNestedStructArgumentRequestParams instance.
func NewMTRTestClusterClusterTestNestedStructArgumentRequestParams() MTRTestClusterClusterTestNestedStructArgumentRequestParams {
	return getMTRTestClusterClusterTestNestedStructArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) Arg1() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) SetArg1(value IMTRUnitTestingClusterNestedStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnestedstructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNestedStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



