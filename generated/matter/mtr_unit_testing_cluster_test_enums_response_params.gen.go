// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEnumsResponseParams] class.
var (
	MTRUnitTestingClusterTestEnumsResponseParamsClass     _MTRUnitTestingClusterTestEnumsResponseParamsClass
	MTRUnitTestingClusterTestEnumsResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEnumsResponseParamsClass() _MTRUnitTestingClusterTestEnumsResponseParamsClass {
	MTRUnitTestingClusterTestEnumsResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEnumsResponseParamsClass = _MTRUnitTestingClusterTestEnumsResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestEnumsResponseParams")}
	})
	return MTRUnitTestingClusterTestEnumsResponseParamsClass
}

type _MTRUnitTestingClusterTestEnumsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEnumsResponseParams] class.
type IMTRUnitTestingClusterTestEnumsResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEnumsResponseParams
type MTRUnitTestingClusterTestEnumsResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEnumsResponseParamsFrom constructs a [MTRUnitTestingClusterTestEnumsResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEnumsResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEnumsResponseParams {
	return MTRUnitTestingClusterTestEnumsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEnumsResponseParamsClass) Alloc() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEnumsResponseParamsClass) New() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) Init() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) Autorelease() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEnumsResponseParams creates a new MTRUnitTestingClusterTestEnumsResponseParams instance.
func NewMTRUnitTestingClusterTestEnumsResponseParams() MTRUnitTestingClusterTestEnumsResponseParams {
	return getMTRUnitTestingClusterTestEnumsResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsresponseparams/arg1
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsresponseparams/arg1
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsresponseparams/arg2
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsresponseparams/arg2
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) SetArg2(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestenumsresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



