// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams] class.
var (
	MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass     _MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass
	MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass() _MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass {
	MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass = _MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams")}
	})
	return MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass
}

type _MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams] class.
type IMTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams interface {
	objectivec.IObject
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams
type MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsFrom constructs a [MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams {
	return MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass) Alloc() MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass) New() MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams) Init() MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams) Autorelease() MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams creates a new MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams instance.
func NewMTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams() MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams {
	return getMTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventresponseparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventresponseparams/value
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestemittestfabricscopedeventresponseparams/value
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventResponseParams) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}
