// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams] class.
var (
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass     _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass() _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass {
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass = _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams")}
	})
	return MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass
}

type _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams] class.
type IMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams interface {
	objectivec.IObject
	Arg1() foundation.Number
	SetArg1(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams
type MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsFrom constructs a [MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	return MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass) Alloc() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass) New() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) Init() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) Autorelease() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams creates a new MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams instance.
func NewMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	return getMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventrequestparams/arg1
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventrequestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventrequestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



