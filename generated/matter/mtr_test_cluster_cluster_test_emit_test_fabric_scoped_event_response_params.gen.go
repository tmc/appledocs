// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams] class.
var (
	MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass     _MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass
	MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass() _MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass {
	MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass = _MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams")}
	})
	return MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass
}

type _MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams] class.
type IMTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams interface {
	IMTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams
type MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams struct {
	MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams
}

// MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsFrom constructs a [MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams {
	return MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams{
		MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams: MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass) Alloc() MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass) New() MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams) Init() MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams) Autorelease() MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams creates a new MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams instance.
func NewMTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams() MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams {
	return getMTRTestClusterClusterTestEmitTestFabricScopedEventResponseParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventresponseparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventresponseparams/value
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestemittestfabricscopedeventresponseparams/value
func (m_ MTRTestClusterClusterTestEmitTestFabricScopedEventResponseParams) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



