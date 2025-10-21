// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams] class.
var (
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass     _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass() _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass {
	MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass = _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams")}
	})
	return MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass
}

type _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams] class.
type IMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams
type MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsFrom constructs a [MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	return MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass) Alloc() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass) New() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) Init() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams) Autorelease() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams creates a new MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams instance.
func NewMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams() MTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParams {
	return getMTRUnitTestingClusterTestEmitTestFabricScopedEventRequestParamsClass().New()
}




