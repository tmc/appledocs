// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestListInt8UReverseResponseParams] class.
var (
	MTRTestClusterClusterTestListInt8UReverseResponseParamsClass     _MTRTestClusterClusterTestListInt8UReverseResponseParamsClass
	MTRTestClusterClusterTestListInt8UReverseResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestListInt8UReverseResponseParamsClass() _MTRTestClusterClusterTestListInt8UReverseResponseParamsClass {
	MTRTestClusterClusterTestListInt8UReverseResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestListInt8UReverseResponseParamsClass = _MTRTestClusterClusterTestListInt8UReverseResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestListInt8UReverseResponseParams")}
	})
	return MTRTestClusterClusterTestListInt8UReverseResponseParamsClass
}

type _MTRTestClusterClusterTestListInt8UReverseResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestListInt8UReverseResponseParams] class.
type IMTRTestClusterClusterTestListInt8UReverseResponseParams interface {
	IMTRUnitTestingClusterTestListInt8UReverseResponseParams
	// properties:
	Arg1() unsafe.Pointer
	SetArg1(value unsafe.Pointer)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestListInt8UReverseResponseParams
type MTRTestClusterClusterTestListInt8UReverseResponseParams struct {
	MTRUnitTestingClusterTestListInt8UReverseResponseParams
}

// MTRTestClusterClusterTestListInt8UReverseResponseParamsFrom constructs a [MTRTestClusterClusterTestListInt8UReverseResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestListInt8UReverseResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestListInt8UReverseResponseParams {
	return MTRTestClusterClusterTestListInt8UReverseResponseParams{
		MTRUnitTestingClusterTestListInt8UReverseResponseParams: MTRUnitTestingClusterTestListInt8UReverseResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestListInt8UReverseResponseParamsClass) Alloc() MTRTestClusterClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestListInt8UReverseResponseParamsClass) New() MTRTestClusterClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestListInt8UReverseResponseParams) Init() MTRTestClusterClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestListInt8UReverseResponseParams) Autorelease() MTRTestClusterClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestListInt8UReverseResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestListInt8UReverseResponseParams creates a new MTRTestClusterClusterTestListInt8UReverseResponseParams instance.
func NewMTRTestClusterClusterTestListInt8UReverseResponseParams() MTRTestClusterClusterTestListInt8UReverseResponseParams {
	return getMTRTestClusterClusterTestListInt8UReverseResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverseresponseparams/arg1
func (m_ MTRTestClusterClusterTestListInt8UReverseResponseParams) Arg1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverseresponseparams/arg1
func (m_ MTRTestClusterClusterTestListInt8UReverseResponseParams) SetArg1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverseresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListInt8UReverseResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestlistint8ureverseresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestListInt8UReverseResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



