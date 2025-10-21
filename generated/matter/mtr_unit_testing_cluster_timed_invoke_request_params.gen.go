// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTimedInvokeRequestParams] class.
var (
	MTRUnitTestingClusterTimedInvokeRequestParamsClass     _MTRUnitTestingClusterTimedInvokeRequestParamsClass
	MTRUnitTestingClusterTimedInvokeRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTimedInvokeRequestParamsClass() _MTRUnitTestingClusterTimedInvokeRequestParamsClass {
	MTRUnitTestingClusterTimedInvokeRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTimedInvokeRequestParamsClass = _MTRUnitTestingClusterTimedInvokeRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTimedInvokeRequestParams")}
	})
	return MTRUnitTestingClusterTimedInvokeRequestParamsClass
}

type _MTRUnitTestingClusterTimedInvokeRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTimedInvokeRequestParams] class.
type IMTRUnitTestingClusterTimedInvokeRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTimedInvokeRequestParams
type MTRUnitTestingClusterTimedInvokeRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTimedInvokeRequestParamsFrom constructs a [MTRUnitTestingClusterTimedInvokeRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTimedInvokeRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTimedInvokeRequestParams {
	return MTRUnitTestingClusterTimedInvokeRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTimedInvokeRequestParamsClass) Alloc() MTRUnitTestingClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTimedInvokeRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTimedInvokeRequestParamsClass) New() MTRUnitTestingClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTimedInvokeRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTimedInvokeRequestParams) Init() MTRUnitTestingClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTimedInvokeRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTimedInvokeRequestParams) Autorelease() MTRUnitTestingClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTimedInvokeRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTimedInvokeRequestParams creates a new MTRUnitTestingClusterTimedInvokeRequestParams instance.
func NewMTRUnitTestingClusterTimedInvokeRequestParams() MTRUnitTestingClusterTimedInvokeRequestParams {
	return getMTRUnitTestingClusterTimedInvokeRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertimedinvokerequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTimedInvokeRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertimedinvokerequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTimedInvokeRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertimedinvokerequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTimedInvokeRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertimedinvokerequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTimedInvokeRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



