// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestParams] class.
var (
	MTRUnitTestingClusterTestParamsClass     _MTRUnitTestingClusterTestParamsClass
	MTRUnitTestingClusterTestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestParamsClass() _MTRUnitTestingClusterTestParamsClass {
	MTRUnitTestingClusterTestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestParamsClass = _MTRUnitTestingClusterTestParamsClass{objc.GetClass("MTRUnitTestingClusterTestParams")}
	})
	return MTRUnitTestingClusterTestParamsClass
}

type _MTRUnitTestingClusterTestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestParams] class.
type IMTRUnitTestingClusterTestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestParams
type MTRUnitTestingClusterTestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestParamsFrom constructs a [MTRUnitTestingClusterTestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestParams {
	return MTRUnitTestingClusterTestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestParamsClass) Alloc() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestParamsClass) New() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestParams) Init() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestParams) Autorelease() MTRUnitTestingClusterTestParams {
	rv := objc.Send[MTRUnitTestingClusterTestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestParams creates a new MTRUnitTestingClusterTestParams instance.
func NewMTRUnitTestingClusterTestParams() MTRUnitTestingClusterTestParams {
	return getMTRUnitTestingClusterTestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



