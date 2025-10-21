// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTimedInvokeRequestParams] class.
var (
	MTRTestClusterClusterTimedInvokeRequestParamsClass     _MTRTestClusterClusterTimedInvokeRequestParamsClass
	MTRTestClusterClusterTimedInvokeRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTimedInvokeRequestParamsClass() _MTRTestClusterClusterTimedInvokeRequestParamsClass {
	MTRTestClusterClusterTimedInvokeRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTimedInvokeRequestParamsClass = _MTRTestClusterClusterTimedInvokeRequestParamsClass{objc.GetClass("MTRTestClusterClusterTimedInvokeRequestParams")}
	})
	return MTRTestClusterClusterTimedInvokeRequestParamsClass
}

type _MTRTestClusterClusterTimedInvokeRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTimedInvokeRequestParams] class.
type IMTRTestClusterClusterTimedInvokeRequestParams interface {
	IMTRUnitTestingClusterTimedInvokeRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTimedInvokeRequestParams
type MTRTestClusterClusterTimedInvokeRequestParams struct {
	MTRUnitTestingClusterTimedInvokeRequestParams
}

// MTRTestClusterClusterTimedInvokeRequestParamsFrom constructs a [MTRTestClusterClusterTimedInvokeRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTimedInvokeRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTimedInvokeRequestParams {
	return MTRTestClusterClusterTimedInvokeRequestParams{
		MTRUnitTestingClusterTimedInvokeRequestParams: MTRUnitTestingClusterTimedInvokeRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTimedInvokeRequestParamsClass) Alloc() MTRTestClusterClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRTestClusterClusterTimedInvokeRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTimedInvokeRequestParamsClass) New() MTRTestClusterClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRTestClusterClusterTimedInvokeRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTimedInvokeRequestParams) Init() MTRTestClusterClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRTestClusterClusterTimedInvokeRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTimedInvokeRequestParams) Autorelease() MTRTestClusterClusterTimedInvokeRequestParams {
	rv := objc.Send[MTRTestClusterClusterTimedInvokeRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTimedInvokeRequestParams creates a new MTRTestClusterClusterTimedInvokeRequestParams instance.
func NewMTRTestClusterClusterTimedInvokeRequestParams() MTRTestClusterClusterTimedInvokeRequestParams {
	return getMTRTestClusterClusterTimedInvokeRequestParamsClass().New()
}




