// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestSimpleArgumentResponseParams] class.
var (
	MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass     _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass
	MTRUnitTestingClusterTestSimpleArgumentResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSimpleArgumentResponseParamsClass() _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass {
	MTRUnitTestingClusterTestSimpleArgumentResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass = _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestSimpleArgumentResponseParams")}
	})
	return MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass
}

type _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestSimpleArgumentResponseParams] class.
type IMTRUnitTestingClusterTestSimpleArgumentResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSimpleArgumentResponseParams
type MTRUnitTestingClusterTestSimpleArgumentResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSimpleArgumentResponseParamsFrom constructs a [MTRUnitTestingClusterTestSimpleArgumentResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSimpleArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	return MTRUnitTestingClusterTestSimpleArgumentResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass) Alloc() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestSimpleArgumentResponseParamsClass) New() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) Init() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) Autorelease() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSimpleArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSimpleArgumentResponseParams creates a new MTRUnitTestingClusterTestSimpleArgumentResponseParams instance.
func NewMTRUnitTestingClusterTestSimpleArgumentResponseParams() MTRUnitTestingClusterTestSimpleArgumentResponseParams {
	return getMTRUnitTestingClusterTestSimpleArgumentResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) ReturnValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("returnValue"))
	return rv
}


// SetReturnValue sets the value of the returnValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) SetReturnValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



