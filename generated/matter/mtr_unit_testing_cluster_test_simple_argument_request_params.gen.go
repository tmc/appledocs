// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestSimpleArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass     _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass
	MTRUnitTestingClusterTestSimpleArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSimpleArgumentRequestParamsClass() _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass {
	MTRUnitTestingClusterTestSimpleArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass = _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestSimpleArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestSimpleArgumentRequestParams] class.
type IMTRUnitTestingClusterTestSimpleArgumentRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentRequestParams
type MTRUnitTestingClusterTestSimpleArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSimpleArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestSimpleArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSimpleArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	return MTRUnitTestingClusterTestSimpleArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestSimpleArgumentRequestParamsClass) New() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) Init() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSimpleArgumentRequestParams creates a new MTRUnitTestingClusterTestSimpleArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestSimpleArgumentRequestParams() MTRUnitTestingClusterTestSimpleArgumentRequestParams {
	return getMTRUnitTestingClusterTestSimpleArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



