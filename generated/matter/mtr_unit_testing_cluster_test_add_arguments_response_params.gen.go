// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestAddArgumentsResponseParams] class.
var (
	MTRUnitTestingClusterTestAddArgumentsResponseParamsClass     _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass
	MTRUnitTestingClusterTestAddArgumentsResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestAddArgumentsResponseParamsClass() _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass {
	MTRUnitTestingClusterTestAddArgumentsResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestAddArgumentsResponseParamsClass = _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestAddArgumentsResponseParams")}
	})
	return MTRUnitTestingClusterTestAddArgumentsResponseParamsClass
}

type _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestAddArgumentsResponseParams] class.
type IMTRUnitTestingClusterTestAddArgumentsResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsResponseParams
type MTRUnitTestingClusterTestAddArgumentsResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestAddArgumentsResponseParamsFrom constructs a [MTRUnitTestingClusterTestAddArgumentsResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestAddArgumentsResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestAddArgumentsResponseParams {
	return MTRUnitTestingClusterTestAddArgumentsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass) Alloc() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestAddArgumentsResponseParamsClass) New() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) Init() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) Autorelease() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestAddArgumentsResponseParams creates a new MTRUnitTestingClusterTestAddArgumentsResponseParams instance.
func NewMTRUnitTestingClusterTestAddArgumentsResponseParams() MTRUnitTestingClusterTestAddArgumentsResponseParams {
	return getMTRUnitTestingClusterTestAddArgumentsResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestaddargumentsresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) ReturnValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("returnValue"))
	return rv
}


// SetReturnValue sets the value of the returnValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestaddargumentsresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) SetReturnValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestaddargumentsresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestaddargumentsresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestAddArgumentsResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



