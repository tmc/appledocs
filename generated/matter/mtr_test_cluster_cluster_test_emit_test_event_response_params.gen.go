// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestEmitTestEventResponseParams] class.
var (
	MTRTestClusterClusterTestEmitTestEventResponseParamsClass     _MTRTestClusterClusterTestEmitTestEventResponseParamsClass
	MTRTestClusterClusterTestEmitTestEventResponseParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEmitTestEventResponseParamsClass() _MTRTestClusterClusterTestEmitTestEventResponseParamsClass {
	MTRTestClusterClusterTestEmitTestEventResponseParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEmitTestEventResponseParamsClass = _MTRTestClusterClusterTestEmitTestEventResponseParamsClass{objc.GetClass("MTRTestClusterClusterTestEmitTestEventResponseParams")}
	})
	return MTRTestClusterClusterTestEmitTestEventResponseParamsClass
}

type _MTRTestClusterClusterTestEmitTestEventResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEmitTestEventResponseParams] class.
type IMTRTestClusterClusterTestEmitTestEventResponseParams interface {
	IMTRUnitTestingClusterTestEmitTestEventResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEmitTestEventResponseParams
type MTRTestClusterClusterTestEmitTestEventResponseParams struct {
	MTRUnitTestingClusterTestEmitTestEventResponseParams
}

// MTRTestClusterClusterTestEmitTestEventResponseParamsFrom constructs a [MTRTestClusterClusterTestEmitTestEventResponseParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEmitTestEventResponseParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEmitTestEventResponseParams {
	return MTRTestClusterClusterTestEmitTestEventResponseParams{
		MTRUnitTestingClusterTestEmitTestEventResponseParams: MTRUnitTestingClusterTestEmitTestEventResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEmitTestEventResponseParamsClass) Alloc() MTRTestClusterClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEmitTestEventResponseParamsClass) New() MTRTestClusterClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEmitTestEventResponseParams) Init() MTRTestClusterClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEmitTestEventResponseParams) Autorelease() MTRTestClusterClusterTestEmitTestEventResponseParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEmitTestEventResponseParams creates a new MTRTestClusterClusterTestEmitTestEventResponseParams instance.
func NewMTRTestClusterClusterTestEmitTestEventResponseParams() MTRTestClusterClusterTestEmitTestEventResponseParams {
	return getMTRTestClusterClusterTestEmitTestEventResponseParamsClass().New()
}




