// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestStructArrayArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass     _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass
	MTRTestClusterClusterTestStructArrayArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestStructArrayArgumentRequestParamsClass() _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass {
	MTRTestClusterClusterTestStructArrayArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass = _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestStructArrayArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestStructArrayArgumentRequestParams] class.
type IMTRTestClusterClusterTestStructArrayArgumentRequestParams interface {
	IMTRUnitTestingClusterTestStructArrayArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestStructArrayArgumentRequestParams
type MTRTestClusterClusterTestStructArrayArgumentRequestParams struct {
	MTRUnitTestingClusterTestStructArrayArgumentRequestParams
}

// MTRTestClusterClusterTestStructArrayArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestStructArrayArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestStructArrayArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	return MTRTestClusterClusterTestStructArrayArgumentRequestParams{
		MTRUnitTestingClusterTestStructArrayArgumentRequestParams: MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestStructArrayArgumentRequestParamsClass) New() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Init() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestStructArrayArgumentRequestParams) Autorelease() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestStructArrayArgumentRequestParams creates a new MTRTestClusterClusterTestStructArrayArgumentRequestParams instance.
func NewMTRTestClusterClusterTestStructArrayArgumentRequestParams() MTRTestClusterClusterTestStructArrayArgumentRequestParams {
	return getMTRTestClusterClusterTestStructArrayArgumentRequestParamsClass().New()
}




