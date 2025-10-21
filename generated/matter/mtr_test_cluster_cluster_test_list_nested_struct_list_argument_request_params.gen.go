// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestListNestedStructListArgumentRequestParams] class.
var (
	MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass     _MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass
	MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass() _MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass {
	MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass = _MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestListNestedStructListArgumentRequestParams")}
	})
	return MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass
}

type _MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestListNestedStructListArgumentRequestParams] class.
type IMTRTestClusterClusterTestListNestedStructListArgumentRequestParams interface {
	IMTRUnitTestingClusterTestListNestedStructListArgumentRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestListNestedStructListArgumentRequestParams
type MTRTestClusterClusterTestListNestedStructListArgumentRequestParams struct {
	MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams
}

// MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsFrom constructs a [MTRTestClusterClusterTestListNestedStructListArgumentRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestListNestedStructListArgumentRequestParams {
	return MTRTestClusterClusterTestListNestedStructListArgumentRequestParams{
		MTRUnitTestingClusterTestListNestedStructListArgumentRequestParams: MTRUnitTestingClusterTestListNestedStructListArgumentRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass) Alloc() MTRTestClusterClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass) New() MTRTestClusterClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListNestedStructListArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestListNestedStructListArgumentRequestParams) Init() MTRTestClusterClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListNestedStructListArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestListNestedStructListArgumentRequestParams) Autorelease() MTRTestClusterClusterTestListNestedStructListArgumentRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestListNestedStructListArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestListNestedStructListArgumentRequestParams creates a new MTRTestClusterClusterTestListNestedStructListArgumentRequestParams instance.
func NewMTRTestClusterClusterTestListNestedStructListArgumentRequestParams() MTRTestClusterClusterTestListNestedStructListArgumentRequestParams {
	return getMTRTestClusterClusterTestListNestedStructListArgumentRequestParamsClass().New()
}




