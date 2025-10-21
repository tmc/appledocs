// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterSimpleStructEchoRequestParams] class.
var (
	MTRTestClusterClusterSimpleStructEchoRequestParamsClass     _MTRTestClusterClusterSimpleStructEchoRequestParamsClass
	MTRTestClusterClusterSimpleStructEchoRequestParamsClassOnce sync.Once
)

func getMTRTestClusterClusterSimpleStructEchoRequestParamsClass() _MTRTestClusterClusterSimpleStructEchoRequestParamsClass {
	MTRTestClusterClusterSimpleStructEchoRequestParamsClassOnce.Do(func() {
		MTRTestClusterClusterSimpleStructEchoRequestParamsClass = _MTRTestClusterClusterSimpleStructEchoRequestParamsClass{objc.GetClass("MTRTestClusterClusterSimpleStructEchoRequestParams")}
	})
	return MTRTestClusterClusterSimpleStructEchoRequestParamsClass
}

type _MTRTestClusterClusterSimpleStructEchoRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterSimpleStructEchoRequestParams] class.
type IMTRTestClusterClusterSimpleStructEchoRequestParams interface {
	IMTRUnitTestingClusterSimpleStructEchoRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterSimpleStructEchoRequestParams
type MTRTestClusterClusterSimpleStructEchoRequestParams struct {
	MTRUnitTestingClusterSimpleStructEchoRequestParams
}

// MTRTestClusterClusterSimpleStructEchoRequestParamsFrom constructs a [MTRTestClusterClusterSimpleStructEchoRequestParams] from an unsafe.Pointer.
func MTRTestClusterClusterSimpleStructEchoRequestParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterSimpleStructEchoRequestParams {
	return MTRTestClusterClusterSimpleStructEchoRequestParams{
		MTRUnitTestingClusterSimpleStructEchoRequestParams: MTRUnitTestingClusterSimpleStructEchoRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterSimpleStructEchoRequestParamsClass) Alloc() MTRTestClusterClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructEchoRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterSimpleStructEchoRequestParamsClass) New() MTRTestClusterClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructEchoRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterSimpleStructEchoRequestParams) Init() MTRTestClusterClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructEchoRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterSimpleStructEchoRequestParams) Autorelease() MTRTestClusterClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRTestClusterClusterSimpleStructEchoRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterSimpleStructEchoRequestParams creates a new MTRTestClusterClusterSimpleStructEchoRequestParams instance.
func NewMTRTestClusterClusterSimpleStructEchoRequestParams() MTRTestClusterClusterSimpleStructEchoRequestParams {
	return getMTRTestClusterClusterSimpleStructEchoRequestParamsClass().New()
}




