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
	// properties:
	ReturnValue() objc.IObject /* cross-framework: NSNumber */
	SetReturnValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) ReturnValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("returnValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) SetReturnValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestsimpleargumentresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSimpleArgumentResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



