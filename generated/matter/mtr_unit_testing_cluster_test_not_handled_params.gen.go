// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestNotHandledParams] class.
var (
	MTRUnitTestingClusterTestNotHandledParamsClass     _MTRUnitTestingClusterTestNotHandledParamsClass
	MTRUnitTestingClusterTestNotHandledParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestNotHandledParamsClass() _MTRUnitTestingClusterTestNotHandledParamsClass {
	MTRUnitTestingClusterTestNotHandledParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestNotHandledParamsClass = _MTRUnitTestingClusterTestNotHandledParamsClass{objc.GetClass("MTRUnitTestingClusterTestNotHandledParams")}
	})
	return MTRUnitTestingClusterTestNotHandledParamsClass
}

type _MTRUnitTestingClusterTestNotHandledParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestNotHandledParams] class.
type IMTRUnitTestingClusterTestNotHandledParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestNotHandledParams
type MTRUnitTestingClusterTestNotHandledParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestNotHandledParamsFrom constructs a [MTRUnitTestingClusterTestNotHandledParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestNotHandledParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestNotHandledParams {
	return MTRUnitTestingClusterTestNotHandledParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestNotHandledParamsClass) Alloc() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestNotHandledParamsClass) New() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestNotHandledParams) Init() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestNotHandledParams) Autorelease() MTRUnitTestingClusterTestNotHandledParams {
	rv := objc.Send[MTRUnitTestingClusterTestNotHandledParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestNotHandledParams creates a new MTRUnitTestingClusterTestNotHandledParams instance.
func NewMTRUnitTestingClusterTestNotHandledParams() MTRUnitTestingClusterTestNotHandledParams {
	return getMTRUnitTestingClusterTestNotHandledParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnothandledparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNotHandledParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnothandledparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestNotHandledParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnothandledparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNotHandledParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestnothandledparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestNotHandledParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



