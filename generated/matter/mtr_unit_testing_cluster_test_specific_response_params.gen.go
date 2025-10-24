// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestSpecificResponseParams] class.
var (
	MTRUnitTestingClusterTestSpecificResponseParamsClass     _MTRUnitTestingClusterTestSpecificResponseParamsClass
	MTRUnitTestingClusterTestSpecificResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestSpecificResponseParamsClass() _MTRUnitTestingClusterTestSpecificResponseParamsClass {
	MTRUnitTestingClusterTestSpecificResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestSpecificResponseParamsClass = _MTRUnitTestingClusterTestSpecificResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestSpecificResponseParams")}
	})
	return MTRUnitTestingClusterTestSpecificResponseParamsClass
}

type _MTRUnitTestingClusterTestSpecificResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestSpecificResponseParams] class.
type IMTRUnitTestingClusterTestSpecificResponseParams interface {
	objectivec.IObject
	// properties:
	ReturnValue() objc.IObject /* cross-framework: NSNumber */
	SetReturnValue(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestSpecificResponseParams
type MTRUnitTestingClusterTestSpecificResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestSpecificResponseParamsFrom constructs a [MTRUnitTestingClusterTestSpecificResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestSpecificResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestSpecificResponseParams {
	return MTRUnitTestingClusterTestSpecificResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestSpecificResponseParamsClass) Alloc() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestSpecificResponseParamsClass) New() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) Init() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) Autorelease() MTRUnitTestingClusterTestSpecificResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestSpecificResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestSpecificResponseParams creates a new MTRUnitTestingClusterTestSpecificResponseParams instance.
func NewMTRUnitTestingClusterTestSpecificResponseParams() MTRUnitTestingClusterTestSpecificResponseParams {
	return getMTRUnitTestingClusterTestSpecificResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) ReturnValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("returnValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificresponseparams/returnvalue
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) SetReturnValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReturnValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestspecificresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestSpecificResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



