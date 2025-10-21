// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestEmitTestEventRequestParams] class.
var (
	MTRUnitTestingClusterTestEmitTestEventRequestParamsClass     _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass
	MTRUnitTestingClusterTestEmitTestEventRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEmitTestEventRequestParamsClass() _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass {
	MTRUnitTestingClusterTestEmitTestEventRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEmitTestEventRequestParamsClass = _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestEmitTestEventRequestParams")}
	})
	return MTRUnitTestingClusterTestEmitTestEventRequestParamsClass
}

type _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEmitTestEventRequestParams] class.
type IMTRUnitTestingClusterTestEmitTestEventRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEmitTestEventRequestParams
type MTRUnitTestingClusterTestEmitTestEventRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEmitTestEventRequestParamsFrom constructs a [MTRUnitTestingClusterTestEmitTestEventRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEmitTestEventRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEmitTestEventRequestParams {
	return MTRUnitTestingClusterTestEmitTestEventRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass) Alloc() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEmitTestEventRequestParamsClass) New() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Init() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEmitTestEventRequestParams) Autorelease() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestEmitTestEventRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEmitTestEventRequestParams creates a new MTRUnitTestingClusterTestEmitTestEventRequestParams instance.
func NewMTRUnitTestingClusterTestEmitTestEventRequestParams() MTRUnitTestingClusterTestEmitTestEventRequestParams {
	return getMTRUnitTestingClusterTestEmitTestEventRequestParamsClass().New()
}




