// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEmitTestEventResponseParams] class.
var (
	MTRUnitTestingClusterTestEmitTestEventResponseParamsClass     _MTRUnitTestingClusterTestEmitTestEventResponseParamsClass
	MTRUnitTestingClusterTestEmitTestEventResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEmitTestEventResponseParamsClass() _MTRUnitTestingClusterTestEmitTestEventResponseParamsClass {
	MTRUnitTestingClusterTestEmitTestEventResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEmitTestEventResponseParamsClass = _MTRUnitTestingClusterTestEmitTestEventResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestEmitTestEventResponseParams")}
	})
	return MTRUnitTestingClusterTestEmitTestEventResponseParamsClass
}

type _MTRUnitTestingClusterTestEmitTestEventResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEmitTestEventResponseParams] class.
type IMTRUnitTestingClusterTestEmitTestEventResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEmitTestEventResponseParams
type MTRUnitTestingClusterTestEmitTestEventResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEmitTestEventResponseParamsFrom constructs a [MTRUnitTestingClusterTestEmitTestEventResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEmitTestEventResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEmitTestEventResponseParams {
	return MTRUnitTestingClusterTestEmitTestEventResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEmitTestEventResponseParamsClass) Alloc() MTRUnitTestingClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEmitTestEventResponseParamsClass) New() MTRUnitTestingClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEmitTestEventResponseParams) Init() MTRUnitTestingClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEmitTestEventResponseParams) Autorelease() MTRUnitTestingClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEmitTestEventResponseParams creates a new MTRUnitTestingClusterTestEmitTestEventResponseParams instance.
func NewMTRUnitTestingClusterTestEmitTestEventResponseParams() MTRUnitTestingClusterTestEmitTestEventResponseParams {
	return getMTRUnitTestingClusterTestEmitTestEventResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestEventResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestEventResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventresponseparams/value
func (m_ MTRUnitTestingClusterTestEmitTestEventResponseParams) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittesteventresponseparams/value
func (m_ MTRUnitTestingClusterTestEmitTestEventResponseParams) SetValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



