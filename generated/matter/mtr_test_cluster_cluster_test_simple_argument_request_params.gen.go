// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestSimpleArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestSimpleArgumentRequestParamsClass     _MTRTestClusterClusterTestSimpleArgumentRequestParamsClass
	MTRTestClusterClusterTestSimpleArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestSimpleArgumentRequestParamsClass() _MTRTestClusterClusterTestSimpleArgumentRequestParamsClass {
	MTRTestClusterClusterTestSimpleArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestSimpleArgumentRequestParamsClass = _MTRTestClusterClusterTestSimpleArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestSimpleArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestSimpleArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestSimpleArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestSimpleArgumentRequestParams] class.
type IMTRTestClusterClusterTestSimpleArgumentRequestParams interface {
	IMTRUnitTestingClusterTestSimpleArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestSimpleArgumentRequestParams
type MTRTestClusterClusterTestSimpleArgumentRequestParams struct {
	MTRUnitTestingClusterTestSimpleArgumentRequestParams
}

// MTRTestClusterClusterTestSimpleArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestSimpleArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestSimpleArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestSimpleArgumentRequestParams {
	return MTRTestClusterClusterTestSimpleArgumentRequestParams{
		MTRUnitTestingClusterTestSimpleArgumentRequestParams: MTRUnitTestingClusterTestSimpleArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestSimpleArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestSimpleArgumentRequestParamsClass) New() MTRTestClusterClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestSimpleArgumentRequestParams) Init() MTRTestClusterClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestSimpleArgumentRequestParams) Autorelease() MTRTestClusterClusterTestSimpleArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestSimpleArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestSimpleArgumentRequestParams creates a new MTRTestClusterClusterTestSimpleArgumentRequestParams instance.
func NewMTRTestClusterClusterTestSimpleArgumentRequestParams() MTRTestClusterClusterTestSimpleArgumentRequestParams {
	return getMTRTestClusterClusterTestSimpleArgumentRequestParamsClass().New()
}




