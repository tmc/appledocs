// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams] class.
var (
	MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass     _MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass
	MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass() _MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass {
	MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass = _MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams")}
	})
	return MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass
}

type _MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams] class.
type IMTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams interface {
	IMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams
type MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams struct {
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams
}

// MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsFrom constructs a [MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams {
	return MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams{
		MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams: MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass) Alloc() MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass) New() MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) Init() MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) Autorelease() MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams creates a new MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams instance.
func NewMTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams() MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams {
	return getMTRTestClusterClusterTestEmitTestFabricScopedEventRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventrequestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventrequestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventrequestparams/arg1
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventrequestparams/arg1
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventRequestParams) SetArg1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}



