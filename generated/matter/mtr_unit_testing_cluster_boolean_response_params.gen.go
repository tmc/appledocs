// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterBooleanResponseParams] class.
var (
	MTRUnitTestingClusterBooleanResponseParamsClass     _MTRUnitTestingClusterBooleanResponseParamsClass
	MTRUnitTestingClusterBooleanResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterBooleanResponseParamsClass() _MTRUnitTestingClusterBooleanResponseParamsClass {
	MTRUnitTestingClusterBooleanResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterBooleanResponseParamsClass = _MTRUnitTestingClusterBooleanResponseParamsClass{objc.GetClass("MTRUnitTestingClusterBooleanResponseParams")}
	})
	return MTRUnitTestingClusterBooleanResponseParamsClass
}

type _MTRUnitTestingClusterBooleanResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterBooleanResponseParams] class.
type IMTRUnitTestingClusterBooleanResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterBooleanResponseParams
type MTRUnitTestingClusterBooleanResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterBooleanResponseParamsFrom constructs a [MTRUnitTestingClusterBooleanResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterBooleanResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterBooleanResponseParams {
	return MTRUnitTestingClusterBooleanResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterBooleanResponseParamsClass) Alloc() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterBooleanResponseParamsClass) New() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterBooleanResponseParams) Init() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterBooleanResponseParams) Autorelease() MTRUnitTestingClusterBooleanResponseParams {
	rv := objc.Send[MTRUnitTestingClusterBooleanResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterBooleanResponseParams creates a new MTRUnitTestingClusterBooleanResponseParams instance.
func NewMTRUnitTestingClusterBooleanResponseParams() MTRUnitTestingClusterBooleanResponseParams {
	return getMTRUnitTestingClusterBooleanResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterbooleanresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterBooleanResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterbooleanresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterBooleanResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterbooleanresponseparams/value
func (m_ MTRUnitTestingClusterBooleanResponseParams) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterbooleanresponseparams/value
func (m_ MTRUnitTestingClusterBooleanResponseParams) SetValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



