// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestListInt8UReverseRequestParams] class.
var (
	MTRTestClusterClusterTestListInt8UReverseRequestParamsClass     _MTRTestClusterClusterTestListInt8UReverseRequestParamsClass
	MTRTestClusterClusterTestListInt8UReverseRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestListInt8UReverseRequestParamsClass() _MTRTestClusterClusterTestListInt8UReverseRequestParamsClass {
	MTRTestClusterClusterTestListInt8UReverseRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestListInt8UReverseRequestParamsClass = _MTRTestClusterClusterTestListInt8UReverseRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestListInt8UReverseRequestParams")}
	})
	return MTRTestClusterClusterTestListInt8UReverseRequestParamsClass
}

type _MTRTestClusterClusterTestListInt8UReverseRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestListInt8UReverseRequestParams] class.
type IMTRTestClusterClusterTestListInt8UReverseRequestParams interface {
	IMTRUnitTestingClusterTestListInt8UReverseRequestParams
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestListInt8UReverseRequestParams
type MTRTestClusterClusterTestListInt8UReverseRequestParams struct {
	MTRUnitTestingClusterTestListInt8UReverseRequestParams
}

// MTRTestClusterClusterTestListInt8UReverseRequestParamsFrom constructs a [MTRTestClusterClusterTestListInt8UReverseRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestListInt8UReverseRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestListInt8UReverseRequestParams {
	return MTRTestClusterClusterTestListInt8UReverseRequestParams{
		MTRUnitTestingClusterTestListInt8UReverseRequestParams: MTRUnitTestingClusterTestListInt8UReverseRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestListInt8UReverseRequestParamsClass) Alloc() MTRTestClusterClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestListInt8UReverseRequestParamsClass) New() MTRTestClusterClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) Init() MTRTestClusterClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) Autorelease() MTRTestClusterClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestListInt8UReverseRequestParams creates a new MTRTestClusterClusterTestListInt8UReverseRequestParams instance.
func NewMTRTestClusterClusterTestListInt8UReverseRequestParams() MTRTestClusterClusterTestListInt8UReverseRequestParams {
	return getMTRTestClusterClusterTestListInt8UReverseRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverserequestparams/arg1
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverserequestparams/arg1
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverserequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverserequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverserequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverserequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListInt8UReverseRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



