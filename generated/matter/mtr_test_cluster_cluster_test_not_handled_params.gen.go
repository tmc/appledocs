// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestNotHandledParams] class.
var (
	MTRTestClusterClusterTestNotHandledParamsClass     _MTRTestClusterClusterTestNotHandledParamsClass
	MTRTestClusterClusterTestNotHandledParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestNotHandledParamsClass() _MTRTestClusterClusterTestNotHandledParamsClass {
	MTRTestClusterClusterTestNotHandledParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestNotHandledParamsClass = _MTRTestClusterClusterTestNotHandledParamsClass{objc.GetClass("MTRTestClusterClusterTestNotHandledParams")}
	})
	return MTRTestClusterClusterTestNotHandledParamsClass
}

type _MTRTestClusterClusterTestNotHandledParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestNotHandledParams] class.
type IMTRTestClusterClusterTestNotHandledParams interface {
	IMTRUnitTestingClusterTestNotHandledParams
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestNotHandledParams
type MTRTestClusterClusterTestNotHandledParams struct {
	MTRUnitTestingClusterTestNotHandledParams
}

// MTRTestClusterClusterTestNotHandledParamsFrom constructs a [MTRTestClusterClusterTestNotHandledParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestNotHandledParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestNotHandledParams {
	return MTRTestClusterClusterTestNotHandledParams{
		MTRUnitTestingClusterTestNotHandledParams: MTRUnitTestingClusterTestNotHandledParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestNotHandledParamsClass) Alloc() MTRTestClusterClusterTestNotHandledParams {
	rv := objc.Send[MTRTestClusterClusterTestNotHandledParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestNotHandledParamsClass) New() MTRTestClusterClusterTestNotHandledParams {
	rv := objc.Send[MTRTestClusterClusterTestNotHandledParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestNotHandledParams) Init() MTRTestClusterClusterTestNotHandledParams {
	rv := objc.Send[MTRTestClusterClusterTestNotHandledParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestNotHandledParams) Autorelease() MTRTestClusterClusterTestNotHandledParams {
	rv := objc.Send[MTRTestClusterClusterTestNotHandledParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestNotHandledParams creates a new MTRTestClusterClusterTestNotHandledParams instance.
func NewMTRTestClusterClusterTestNotHandledParams() MTRTestClusterClusterTestNotHandledParams {
	return getMTRTestClusterClusterTestNotHandledParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnothandledparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNotHandledParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnothandledparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestNotHandledParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnothandledparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNotHandledParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestnothandledparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestNotHandledParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



