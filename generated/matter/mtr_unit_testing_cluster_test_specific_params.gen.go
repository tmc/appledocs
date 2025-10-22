// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestSpecificParams] class.
var (
	MTRUnitTestingClusterTestSpecificParamsClass     _MTRUnitTestingClusterTestSpecificParamsClass
	MTRUnitTestingClusterTestSpecificParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSpecificParamsClass() _MTRUnitTestingClusterTestSpecificParamsClass {
	MTRUnitTestingClusterTestSpecificParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSpecificParamsClass = _MTRUnitTestingClusterTestSpecificParamsClass{objc.GetClass("MTRUnitTestingClusterTestSpecificParams")}
	})
	return MTRUnitTestingClusterTestSpecificParamsClass
}

type _MTRUnitTestingClusterTestSpecificParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestSpecificParams] class.
type IMTRUnitTestingClusterTestSpecificParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificParams
type MTRUnitTestingClusterTestSpecificParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSpecificParamsFrom constructs a [MTRUnitTestingClusterTestSpecificParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSpecificParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSpecificParams {
	return MTRUnitTestingClusterTestSpecificParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSpecificParamsClass) Alloc() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestSpecificParamsClass) New() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSpecificParams) Init() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSpecificParams) Autorelease() MTRUnitTestingClusterTestSpecificParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSpecificParams creates a new MTRUnitTestingClusterTestSpecificParams instance.
func NewMTRUnitTestingClusterTestSpecificParams() MTRUnitTestingClusterTestSpecificParams {
	return getMTRUnitTestingClusterTestSpecificParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestSpecificParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestSpecificParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSpecificParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSpecificParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



