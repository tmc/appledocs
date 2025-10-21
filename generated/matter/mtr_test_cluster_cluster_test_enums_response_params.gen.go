// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestEnumsResponseParams] class.
var (
	MTRTestClusterClusterTestEnumsResponseParamsClass     _MTRTestClusterClusterTestEnumsResponseParamsClass
	MTRTestClusterClusterTestEnumsResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEnumsResponseParamsClass() _MTRTestClusterClusterTestEnumsResponseParamsClass {
	MTRTestClusterClusterTestEnumsResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEnumsResponseParamsClass = _MTRTestClusterClusterTestEnumsResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestEnumsResponseParams")}
	})
	return MTRTestClusterClusterTestEnumsResponseParamsClass
}

type _MTRTestClusterClusterTestEnumsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEnumsResponseParams] class.
type IMTRTestClusterClusterTestEnumsResponseParams interface {
	IMTRUnitTestingClusterTestEnumsResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEnumsResponseParams
type MTRTestClusterClusterTestEnumsResponseParams struct {
	MTRUnitTestingClusterTestEnumsResponseParams
}

// MTRTestClusterClusterTestEnumsResponseParamsFrom constructs a [MTRTestClusterClusterTestEnumsResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEnumsResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEnumsResponseParams {
	return MTRTestClusterClusterTestEnumsResponseParams{
		MTRUnitTestingClusterTestEnumsResponseParams: MTRUnitTestingClusterTestEnumsResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEnumsResponseParamsClass) Alloc() MTRTestClusterClusterTestEnumsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEnumsResponseParamsClass) New() MTRTestClusterClusterTestEnumsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEnumsResponseParams) Init() MTRTestClusterClusterTestEnumsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEnumsResponseParams) Autorelease() MTRTestClusterClusterTestEnumsResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEnumsResponseParams creates a new MTRTestClusterClusterTestEnumsResponseParams instance.
func NewMTRTestClusterClusterTestEnumsResponseParams() MTRTestClusterClusterTestEnumsResponseParams {
	return getMTRTestClusterClusterTestEnumsResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsresponseparams/arg1
func (m_ MTRTestClusterClusterTestEnumsResponseParams) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsresponseparams/arg1
func (m_ MTRTestClusterClusterTestEnumsResponseParams) SetArg1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsresponseparams/arg2
func (m_ MTRTestClusterClusterTestEnumsResponseParams) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsresponseparams/arg2
func (m_ MTRTestClusterClusterTestEnumsResponseParams) SetArg2(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEnumsResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestenumsresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEnumsResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



