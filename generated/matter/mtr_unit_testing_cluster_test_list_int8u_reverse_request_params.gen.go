// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestListInt8UReverseRequestParams] class.
var (
	MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass     _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass
	MTRUnitTestingClusterTestListInt8UReverseRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListInt8UReverseRequestParamsClass() _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass {
	MTRUnitTestingClusterTestListInt8UReverseRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass = _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListInt8UReverseRequestParams")}
	})
	return MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass
}

type _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestListInt8UReverseRequestParams] class.
type IMTRUnitTestingClusterTestListInt8UReverseRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseRequestParams
type MTRUnitTestingClusterTestListInt8UReverseRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListInt8UReverseRequestParamsFrom constructs a [MTRUnitTestingClusterTestListInt8UReverseRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListInt8UReverseRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	return MTRUnitTestingClusterTestListInt8UReverseRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass) Alloc() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestListInt8UReverseRequestParamsClass) New() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) Init() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) Autorelease() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListInt8UReverseRequestParams creates a new MTRUnitTestingClusterTestListInt8UReverseRequestParams instance.
func NewMTRUnitTestingClusterTestListInt8UReverseRequestParams() MTRUnitTestingClusterTestListInt8UReverseRequestParams {
	return getMTRUnitTestingClusterTestListInt8UReverseRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistint8ureverserequestparams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistint8ureverserequestparams/arg1
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistint8ureverserequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistint8ureverserequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistint8ureverserequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestlistint8ureverserequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestListInt8UReverseRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



