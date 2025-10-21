// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] class.
var (
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass     _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass() _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass {
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass = _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArrayArgumentResponseParams")}
	})
	return MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass
}

type _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] class.
type IMTRUnitTestingClusterTestStructArrayArgumentResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams
type MTRUnitTestingClusterTestStructArrayArgumentResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom constructs a [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	return MTRUnitTestingClusterTestStructArrayArgumentResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass) Alloc() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass) New() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Init() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Autorelease() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArrayArgumentResponseParams creates a new MTRUnitTestingClusterTestStructArrayArgumentResponseParams instance.
func NewMTRUnitTestingClusterTestStructArrayArgumentResponseParams() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	return getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg2() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg2
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg2(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg6() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg6"))
	return rv
}


// SetArg6 sets the value of the arg6 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg6
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg6(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg5() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg5"))
	return rv
}


// SetArg5 sets the value of the arg5 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg5
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg5(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg4() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg4"))
	return rv
}


// SetArg4 sets the value of the arg4 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg4
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg4(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg1
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Arg3() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg3"))
	return rv
}


// SetArg3 sets the value of the arg3 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterteststructarrayargumentresponseparams/arg3
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) SetArg3(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}



