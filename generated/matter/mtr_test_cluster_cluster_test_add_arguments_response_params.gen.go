// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestAddArgumentsResponseParams] class.
var (
	MTRTestClusterClusterTestAddArgumentsResponseParamsClass     _MTRTestClusterClusterTestAddArgumentsResponseParamsClass
	MTRTestClusterClusterTestAddArgumentsResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestAddArgumentsResponseParamsClass() _MTRTestClusterClusterTestAddArgumentsResponseParamsClass {
	MTRTestClusterClusterTestAddArgumentsResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestAddArgumentsResponseParamsClass = _MTRTestClusterClusterTestAddArgumentsResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestAddArgumentsResponseParams")}
	})
	return MTRTestClusterClusterTestAddArgumentsResponseParamsClass
}

type _MTRTestClusterClusterTestAddArgumentsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestAddArgumentsResponseParams] class.
type IMTRTestClusterClusterTestAddArgumentsResponseParams interface {
	IMTRUnitTestingClusterTestAddArgumentsResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestAddArgumentsResponseParams
type MTRTestClusterClusterTestAddArgumentsResponseParams struct {
	MTRUnitTestingClusterTestAddArgumentsResponseParams
}

// MTRTestClusterClusterTestAddArgumentsResponseParamsFrom constructs a [MTRTestClusterClusterTestAddArgumentsResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestAddArgumentsResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestAddArgumentsResponseParams {
	return MTRTestClusterClusterTestAddArgumentsResponseParams{
		MTRUnitTestingClusterTestAddArgumentsResponseParams: MTRUnitTestingClusterTestAddArgumentsResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestAddArgumentsResponseParamsClass) Alloc() MTRTestClusterClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestAddArgumentsResponseParamsClass) New() MTRTestClusterClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestAddArgumentsResponseParams) Init() MTRTestClusterClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestAddArgumentsResponseParams) Autorelease() MTRTestClusterClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestAddArgumentsResponseParams creates a new MTRTestClusterClusterTestAddArgumentsResponseParams instance.
func NewMTRTestClusterClusterTestAddArgumentsResponseParams() MTRTestClusterClusterTestAddArgumentsResponseParams {
	return getMTRTestClusterClusterTestAddArgumentsResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestaddargumentsresponseparams/returnvalue
func (m_ MTRTestClusterClusterTestAddArgumentsResponseParams) ReturnValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("returnValue"))
	return rv
}


// SetReturnValue sets the value of the returnValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestaddargumentsresponseparams/returnvalue
func (m_ MTRTestClusterClusterTestAddArgumentsResponseParams) SetReturnValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestaddargumentsresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestAddArgumentsResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestaddargumentsresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestAddArgumentsResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



