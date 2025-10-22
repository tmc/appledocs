// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestListInt8UArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass     _MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass
	MTRTestClusterClusterTestListInt8UArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestListInt8UArgumentRequestParamsClass() _MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass {
	MTRTestClusterClusterTestListInt8UArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass = _MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestListInt8UArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestListInt8UArgumentRequestParams] class.
type IMTRTestClusterClusterTestListInt8UArgumentRequestParams interface {
	IMTRUnitTestingClusterTestListInt8UArgumentRequestParams
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestListInt8UArgumentRequestParams
type MTRTestClusterClusterTestListInt8UArgumentRequestParams struct {
	MTRUnitTestingClusterTestListInt8UArgumentRequestParams
}

// MTRTestClusterClusterTestListInt8UArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestListInt8UArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestListInt8UArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestListInt8UArgumentRequestParams {
	return MTRTestClusterClusterTestListInt8UArgumentRequestParams{
		MTRUnitTestingClusterTestListInt8UArgumentRequestParams: MTRUnitTestingClusterTestListInt8UArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestListInt8UArgumentRequestParamsClass) New() MTRTestClusterClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) Init() MTRTestClusterClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) Autorelease() MTRTestClusterClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestListInt8UArgumentRequestParams creates a new MTRTestClusterClusterTestListInt8UArgumentRequestParams instance.
func NewMTRTestClusterClusterTestListInt8UArgumentRequestParams() MTRTestClusterClusterTestListInt8UArgumentRequestParams {
	return getMTRTestClusterClusterTestListInt8UArgumentRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8uargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8uargumentrequestparams/arg1
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8uargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8uargumentrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8uargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8uargumentrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListInt8UArgumentRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



