// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestEnumsRequestParams] class.
var (
	MTRTestClusterClusterTestEnumsRequestParamsClass     _MTRTestClusterClusterTestEnumsRequestParamsClass
	MTRTestClusterClusterTestEnumsRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestEnumsRequestParamsClass() _MTRTestClusterClusterTestEnumsRequestParamsClass {
	MTRTestClusterClusterTestEnumsRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestEnumsRequestParamsClass = _MTRTestClusterClusterTestEnumsRequestParamsClass{objc.GetClass("MTRTestClusterClusterTestEnumsRequestParams")}
	})
	return MTRTestClusterClusterTestEnumsRequestParamsClass
}

type _MTRTestClusterClusterTestEnumsRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEnumsRequestParams] class.
type IMTRTestClusterClusterTestEnumsRequestParams interface {
	IMTRUnitTestingClusterTestEnumsRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEnumsRequestParams
type MTRTestClusterClusterTestEnumsRequestParams struct {
	MTRUnitTestingClusterTestEnumsRequestParams
}

// MTRTestClusterClusterTestEnumsRequestParamsFrom constructs a [MTRTestClusterClusterTestEnumsRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestEnumsRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEnumsRequestParams {
	return MTRTestClusterClusterTestEnumsRequestParams{
		MTRUnitTestingClusterTestEnumsRequestParams: MTRUnitTestingClusterTestEnumsRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEnumsRequestParamsClass) Alloc() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEnumsRequestParamsClass) New() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEnumsRequestParams) Init() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEnumsRequestParams) Autorelease() MTRTestClusterClusterTestEnumsRequestParams {
	rv := objc.Send[MTRTestClusterClusterTestEnumsRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEnumsRequestParams creates a new MTRTestClusterClusterTestEnumsRequestParams instance.
func NewMTRTestClusterClusterTestEnumsRequestParams() MTRTestClusterClusterTestEnumsRequestParams {
	return getMTRTestClusterClusterTestEnumsRequestParamsClass().New()
}




