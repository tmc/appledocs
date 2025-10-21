// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestStructArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestStructArgumentRequestParamsClass     _MTRTestClusterClusterTestStructArgumentRequestParamsClass
	MTRTestClusterClusterTestStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestStructArgumentRequestParamsClass() _MTRTestClusterClusterTestStructArgumentRequestParamsClass {
	MTRTestClusterClusterTestStructArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestStructArgumentRequestParamsClass = _MTRTestClusterClusterTestStructArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestStructArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestStructArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestStructArgumentRequestParams] class.
type IMTRTestClusterClusterTestStructArgumentRequestParams interface {
	IMTRUnitTestingClusterTestStructArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestStructArgumentRequestParams
type MTRTestClusterClusterTestStructArgumentRequestParams struct {
	MTRUnitTestingClusterTestStructArgumentRequestParams
}

// MTRTestClusterClusterTestStructArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestStructArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestStructArgumentRequestParams {
	return MTRTestClusterClusterTestStructArgumentRequestParams{
		MTRUnitTestingClusterTestStructArgumentRequestParams: MTRUnitTestingClusterTestStructArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestStructArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestStructArgumentRequestParamsClass) New() MTRTestClusterClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) Init() MTRTestClusterClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) Autorelease() MTRTestClusterClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestStructArgumentRequestParams creates a new MTRTestClusterClusterTestStructArgumentRequestParams instance.
func NewMTRTestClusterClusterTestStructArgumentRequestParams() MTRTestClusterClusterTestStructArgumentRequestParams {
	return getMTRTestClusterClusterTestStructArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestStructArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



