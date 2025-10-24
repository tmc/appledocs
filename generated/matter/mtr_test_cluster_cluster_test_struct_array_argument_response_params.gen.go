// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestStructArrayArgumentResponseParams] class.
var (
	MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass     _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass
	MTRTestClusterClusterTestStructArrayArgumentResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestStructArrayArgumentResponseParamsClass() _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass {
	MTRTestClusterClusterTestStructArrayArgumentResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass = _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestStructArrayArgumentResponseParams")}
	})
	return MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass
}

type _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestStructArrayArgumentResponseParams] class.
type IMTRTestClusterClusterTestStructArrayArgumentResponseParams interface {
	IMTRUnitTestingClusterTestStructArrayArgumentResponseParams
	// properties:
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	Arg2() unsafe.Pointer
	SetArg2(value unsafe.Pointer)
	Arg3() unsafe.Pointer
	SetArg3(value unsafe.Pointer)
	Arg4() unsafe.Pointer
	SetArg4(value unsafe.Pointer)
	Arg5() objc.IObject /* cross-framework: NSNumber */
	SetArg5(value objc.IObject /* cross-framework: NSNumber */)
	Arg6() objc.IObject /* cross-framework: NSNumber */
	SetArg6(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestStructArrayArgumentResponseParams
type MTRTestClusterClusterTestStructArrayArgumentResponseParams struct {
	MTRUnitTestingClusterTestStructArrayArgumentResponseParams
}

// MTRTestClusterClusterTestStructArrayArgumentResponseParamsFrom constructs a [MTRTestClusterClusterTestStructArrayArgumentResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestStructArrayArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	return MTRTestClusterClusterTestStructArrayArgumentResponseParams{
		MTRUnitTestingClusterTestStructArrayArgumentResponseParams: MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass) Alloc() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestStructArrayArgumentResponseParamsClass) New() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Init() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Autorelease() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestStructArrayArgumentResponseParams creates a new MTRTestClusterClusterTestStructArrayArgumentResponseParams instance.
func NewMTRTestClusterClusterTestStructArrayArgumentResponseParams() MTRTestClusterClusterTestStructArrayArgumentResponseParams {
	return getMTRTestClusterClusterTestStructArrayArgumentResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg1
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg1
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg2
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Arg2() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg2"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg2
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetArg2(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg3
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Arg3() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg3"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg3
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetArg3(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg4
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Arg4() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg4"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg4
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetArg4(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg5
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Arg5() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg5"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg5
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetArg5(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg6
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) Arg6() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("arg6"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/arg6
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetArg6(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterteststructarrayargumentresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestStructArrayArgumentResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



