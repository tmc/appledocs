// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestSpecificResponseParams] class.
var (
	MTRTestClusterClusterTestSpecificResponseParamsClass     _MTRTestClusterClusterTestSpecificResponseParamsClass
	MTRTestClusterClusterTestSpecificResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestSpecificResponseParamsClass() _MTRTestClusterClusterTestSpecificResponseParamsClass {
	MTRTestClusterClusterTestSpecificResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestSpecificResponseParamsClass = _MTRTestClusterClusterTestSpecificResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestSpecificResponseParams")}
	})
	return MTRTestClusterClusterTestSpecificResponseParamsClass
}

type _MTRTestClusterClusterTestSpecificResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestSpecificResponseParams] class.
type IMTRTestClusterClusterTestSpecificResponseParams interface {
	IMTRUnitTestingClusterTestSpecificResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestSpecificResponseParams
type MTRTestClusterClusterTestSpecificResponseParams struct {
	MTRUnitTestingClusterTestSpecificResponseParams
}

// MTRTestClusterClusterTestSpecificResponseParamsFrom constructs a [MTRTestClusterClusterTestSpecificResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestSpecificResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestSpecificResponseParams {
	return MTRTestClusterClusterTestSpecificResponseParams{
		MTRUnitTestingClusterTestSpecificResponseParams: MTRUnitTestingClusterTestSpecificResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestSpecificResponseParamsClass) Alloc() MTRTestClusterClusterTestSpecificResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestSpecificResponseParamsClass) New() MTRTestClusterClusterTestSpecificResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestSpecificResponseParams) Init() MTRTestClusterClusterTestSpecificResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestSpecificResponseParams) Autorelease() MTRTestClusterClusterTestSpecificResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestSpecificResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestSpecificResponseParams creates a new MTRTestClusterClusterTestSpecificResponseParams instance.
func NewMTRTestClusterClusterTestSpecificResponseParams() MTRTestClusterClusterTestSpecificResponseParams {
	return getMTRTestClusterClusterTestSpecificResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestSpecificResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestSpecificResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificresponseparams/returnvalue
func (m_ MTRTestClusterClusterTestSpecificResponseParams) ReturnValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("returnValue"))
	return rv
}


// SetReturnValue sets the value of the returnValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestspecificresponseparams/returnvalue
func (m_ MTRTestClusterClusterTestSpecificResponseParams) SetReturnValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}



