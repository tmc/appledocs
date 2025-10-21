// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestEmitTestEventRequestParams] class.
var (
	MTRTestClusterClusterTestEmitTestEventRequestParamsClass     _MTRTestClusterClusterTestEmitTestEventRequestParamsClass
	MTRTestClusterClusterTestEmitTestEventRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEmitTestEventRequestParamsClass() _MTRTestClusterClusterTestEmitTestEventRequestParamsClass {
	MTRTestClusterClusterTestEmitTestEventRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEmitTestEventRequestParamsClass = _MTRTestClusterClusterTestEmitTestEventRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestEmitTestEventRequestParams")}
	})
	return MTRTestClusterClusterTestEmitTestEventRequestParamsClass
}

type _MTRTestClusterClusterTestEmitTestEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEmitTestEventRequestParams] class.
type IMTRTestClusterClusterTestEmitTestEventRequestParams interface {
	IMTRUnitTestingClusterTestEmitTestEventRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEmitTestEventRequestParams
type MTRTestClusterClusterTestEmitTestEventRequestParams struct {
	MTRUnitTestingClusterTestEmitTestEventRequestParams
}

// MTRTestClusterClusterTestEmitTestEventRequestParamsFrom constructs a [MTRTestClusterClusterTestEmitTestEventRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEmitTestEventRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEmitTestEventRequestParams {
	return MTRTestClusterClusterTestEmitTestEventRequestParams{
		MTRUnitTestingClusterTestEmitTestEventRequestParams: MTRUnitTestingClusterTestEmitTestEventRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEmitTestEventRequestParamsClass) Alloc() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEmitTestEventRequestParamsClass) New() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Init() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEmitTestEventRequestParams) Autorelease() MTRTestClusterClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEmitTestEventRequestParams creates a new MTRTestClusterClusterTestEmitTestEventRequestParams instance.
func NewMTRTestClusterClusterTestEmitTestEventRequestParams() MTRTestClusterClusterTestEmitTestEventRequestParams {
	return getMTRTestClusterClusterTestEmitTestEventRequestParamsClass().New()
}




