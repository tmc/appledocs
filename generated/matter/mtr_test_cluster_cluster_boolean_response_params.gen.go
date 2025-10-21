// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterBooleanResponseParams] class.
var (
	MTRTestClusterClusterBooleanResponseParamsClass     _MTRTestClusterClusterBooleanResponseParamsClass
	MTRTestClusterClusterBooleanResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterBooleanResponseParamsClass() _MTRTestClusterClusterBooleanResponseParamsClass {
	MTRTestClusterClusterBooleanResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterBooleanResponseParamsClass = _MTRTestClusterClusterBooleanResponseParamsClass{objc.GetClass("MTRTestClusterClusterBooleanResponseParams")}
	})
	return MTRTestClusterClusterBooleanResponseParamsClass
}

type _MTRTestClusterClusterBooleanResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterBooleanResponseParams] class.
type IMTRTestClusterClusterBooleanResponseParams interface {
	IMTRUnitTestingClusterBooleanResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterBooleanResponseParams
type MTRTestClusterClusterBooleanResponseParams struct {
	MTRUnitTestingClusterBooleanResponseParams
}

// MTRTestClusterClusterBooleanResponseParamsFrom constructs a [MTRTestClusterClusterBooleanResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterBooleanResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterBooleanResponseParams {
	return MTRTestClusterClusterBooleanResponseParams{
		MTRUnitTestingClusterBooleanResponseParams: MTRUnitTestingClusterBooleanResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterBooleanResponseParamsClass) Alloc() MTRTestClusterClusterBooleanResponseParams {
	rv := objc.Send[MTRTestClusterClusterBooleanResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterBooleanResponseParamsClass) New() MTRTestClusterClusterBooleanResponseParams {
	rv := objc.Send[MTRTestClusterClusterBooleanResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterBooleanResponseParams) Init() MTRTestClusterClusterBooleanResponseParams {
	rv := objc.Send[MTRTestClusterClusterBooleanResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterBooleanResponseParams) Autorelease() MTRTestClusterClusterBooleanResponseParams {
	rv := objc.Send[MTRTestClusterClusterBooleanResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterBooleanResponseParams creates a new MTRTestClusterClusterBooleanResponseParams instance.
func NewMTRTestClusterClusterBooleanResponseParams() MTRTestClusterClusterBooleanResponseParams {
	return getMTRTestClusterClusterBooleanResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterbooleanresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterBooleanResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterbooleanresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterBooleanResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterbooleanresponseparams/value
func (m_ MTRTestClusterClusterBooleanResponseParams) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterbooleanresponseparams/value
func (m_ MTRTestClusterClusterBooleanResponseParams) SetValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



