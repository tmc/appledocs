// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterSimpleStructResponseParams] class.
var (
	MTRTestClusterClusterSimpleStructResponseParamsClass     _MTRTestClusterClusterSimpleStructResponseParamsClass
	MTRTestClusterClusterSimpleStructResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterSimpleStructResponseParamsClass() _MTRTestClusterClusterSimpleStructResponseParamsClass {
	MTRTestClusterClusterSimpleStructResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterSimpleStructResponseParamsClass = _MTRTestClusterClusterSimpleStructResponseParamsClass{objc.GetClass("MTRTestClusterClusterSimpleStructResponseParams")}
	})
	return MTRTestClusterClusterSimpleStructResponseParamsClass
}

type _MTRTestClusterClusterSimpleStructResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterSimpleStructResponseParams] class.
type IMTRTestClusterClusterSimpleStructResponseParams interface {
	IMTRUnitTestingClusterSimpleStructResponseParams
	Arg1() MTRUnitTestingClusterSimpleStruct
	SetArg1(value IMTRUnitTestingClusterSimpleStruct)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterSimpleStructResponseParams
type MTRTestClusterClusterSimpleStructResponseParams struct {
	MTRUnitTestingClusterSimpleStructResponseParams
}

// MTRTestClusterClusterSimpleStructResponseParamsFrom constructs a [MTRTestClusterClusterSimpleStructResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterSimpleStructResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterSimpleStructResponseParams {
	return MTRTestClusterClusterSimpleStructResponseParams{
		MTRUnitTestingClusterSimpleStructResponseParams: MTRUnitTestingClusterSimpleStructResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterSimpleStructResponseParamsClass) Alloc() MTRTestClusterClusterSimpleStructResponseParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterSimpleStructResponseParamsClass) New() MTRTestClusterClusterSimpleStructResponseParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterSimpleStructResponseParams) Init() MTRTestClusterClusterSimpleStructResponseParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterSimpleStructResponseParams) Autorelease() MTRTestClusterClusterSimpleStructResponseParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterSimpleStructResponseParams creates a new MTRTestClusterClusterSimpleStructResponseParams instance.
func NewMTRTestClusterClusterSimpleStructResponseParams() MTRTestClusterClusterSimpleStructResponseParams {
	return getMTRTestClusterClusterSimpleStructResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestructresponseparams/arg1
func (m_ MTRTestClusterClusterSimpleStructResponseParams) Arg1() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestructresponseparams/arg1
func (m_ MTRTestClusterClusterSimpleStructResponseParams) SetArg1(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestructresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterSimpleStructResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestructresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterSimpleStructResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



