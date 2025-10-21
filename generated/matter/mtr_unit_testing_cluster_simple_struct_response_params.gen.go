// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterSimpleStructResponseParams] class.
var (
	MTRUnitTestingClusterSimpleStructResponseParamsClass     _MTRUnitTestingClusterSimpleStructResponseParamsClass
	MTRUnitTestingClusterSimpleStructResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructResponseParamsClass() _MTRUnitTestingClusterSimpleStructResponseParamsClass {
	MTRUnitTestingClusterSimpleStructResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructResponseParamsClass = _MTRUnitTestingClusterSimpleStructResponseParamsClass{objc.GetClass("MTRUnitTestingClusterSimpleStructResponseParams")}
	})
	return MTRUnitTestingClusterSimpleStructResponseParamsClass
}

type _MTRUnitTestingClusterSimpleStructResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterSimpleStructResponseParams] class.
type IMTRUnitTestingClusterSimpleStructResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructResponseParams
type MTRUnitTestingClusterSimpleStructResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructResponseParamsFrom constructs a [MTRUnitTestingClusterSimpleStructResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStructResponseParams {
	return MTRUnitTestingClusterSimpleStructResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructResponseParamsClass) Alloc() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterSimpleStructResponseParamsClass) New() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) Init() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) Autorelease() MTRUnitTestingClusterSimpleStructResponseParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStructResponseParams creates a new MTRUnitTestingClusterSimpleStructResponseParams instance.
func NewMTRUnitTestingClusterSimpleStructResponseParams() MTRUnitTestingClusterSimpleStructResponseParams {
	return getMTRUnitTestingClusterSimpleStructResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestructresponseparams/arg1
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestructresponseparams/arg1
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestructresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestructresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterSimpleStructResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



