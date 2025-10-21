// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestNestedStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestNestedStructArgumentRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNestedStructArgumentRequestParams
type MTRUnitTestingClusterTestNestedStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestNestedStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNestedStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	return MTRUnitTestingClusterTestNestedStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Init() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestNestedStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNestedStructArgumentRequestParams creates a new MTRUnitTestingClusterTestNestedStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestNestedStructArgumentRequestParams() MTRUnitTestingClusterTestNestedStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestNestedStructArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnestedstructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNestedStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



