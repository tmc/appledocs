// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass     _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass() _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass {
	MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass = _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams] class.
type IMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams
type MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	return MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass) New() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) Init() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams creates a new MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams() MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams {
	return getMTRUnitTestingClusterTestSimpleOptionalArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleoptionalargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleoptionalargumentrequestparams/arg1
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) SetArg1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleoptionalargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleoptionalargumentrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleoptionalargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleoptionalargumentrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestSimpleOptionalArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



