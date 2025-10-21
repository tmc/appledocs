// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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




