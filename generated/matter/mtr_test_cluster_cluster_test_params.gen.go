// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestParams] class.
var (
	MTRTestClusterClusterTestParamsClass     _MTRTestClusterClusterTestParamsClass
	MTRTestClusterClusterTestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestParamsClass() _MTRTestClusterClusterTestParamsClass {
	MTRTestClusterClusterTestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestParamsClass = _MTRTestClusterClusterTestParamsClass{objc.GetClass("MTRTestClusterClusterTestParams")}
	})
	return MTRTestClusterClusterTestParamsClass
}

type _MTRTestClusterClusterTestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestParams] class.
type IMTRTestClusterClusterTestParams interface {
	IMTRUnitTestingClusterTestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestParams
type MTRTestClusterClusterTestParams struct {
	MTRUnitTestingClusterTestParams
}

// MTRTestClusterClusterTestParamsFrom constructs a [MTRTestClusterClusterTestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestParams {
	return MTRTestClusterClusterTestParams{
		MTRUnitTestingClusterTestParams: MTRUnitTestingClusterTestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestParamsClass) Alloc() MTRTestClusterClusterTestParams {
	rv := objc.Send[MTRTestClusterClusterTestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestParamsClass) New() MTRTestClusterClusterTestParams {
	rv := objc.Send[MTRTestClusterClusterTestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestParams) Init() MTRTestClusterClusterTestParams {
	rv := objc.Send[MTRTestClusterClusterTestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestParams) Autorelease() MTRTestClusterClusterTestParams {
	rv := objc.Send[MTRTestClusterClusterTestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestParams creates a new MTRTestClusterClusterTestParams instance.
func NewMTRTestClusterClusterTestParams() MTRTestClusterClusterTestParams {
	return getMTRTestClusterClusterTestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}



