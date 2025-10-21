// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestAddArgumentsParams] class.
var (
	MTRTestClusterClusterTestAddArgumentsParamsClass     _MTRTestClusterClusterTestAddArgumentsParamsClass
	MTRTestClusterClusterTestAddArgumentsParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestAddArgumentsParamsClass() _MTRTestClusterClusterTestAddArgumentsParamsClass {
	MTRTestClusterClusterTestAddArgumentsParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestAddArgumentsParamsClass = _MTRTestClusterClusterTestAddArgumentsParamsClass{objc.GetClass("MTRTestClusterClusterTestAddArgumentsParams")}
	})
	return MTRTestClusterClusterTestAddArgumentsParamsClass
}

type _MTRTestClusterClusterTestAddArgumentsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestAddArgumentsParams] class.
type IMTRTestClusterClusterTestAddArgumentsParams interface {
	IMTRUnitTestingClusterTestAddArgumentsParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestAddArgumentsParams
type MTRTestClusterClusterTestAddArgumentsParams struct {
	MTRUnitTestingClusterTestAddArgumentsParams
}

// MTRTestClusterClusterTestAddArgumentsParamsFrom constructs a [MTRTestClusterClusterTestAddArgumentsParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestAddArgumentsParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestAddArgumentsParams {
	return MTRTestClusterClusterTestAddArgumentsParams{
		MTRUnitTestingClusterTestAddArgumentsParams: MTRUnitTestingClusterTestAddArgumentsParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestAddArgumentsParamsClass) Alloc() MTRTestClusterClusterTestAddArgumentsParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestAddArgumentsParamsClass) New() MTRTestClusterClusterTestAddArgumentsParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestAddArgumentsParams) Init() MTRTestClusterClusterTestAddArgumentsParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestAddArgumentsParams) Autorelease() MTRTestClusterClusterTestAddArgumentsParams {
	rv := objc.Send[MTRTestClusterClusterTestAddArgumentsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestAddArgumentsParams creates a new MTRTestClusterClusterTestAddArgumentsParams instance.
func NewMTRTestClusterClusterTestAddArgumentsParams() MTRTestClusterClusterTestAddArgumentsParams {
	return getMTRTestClusterClusterTestAddArgumentsParamsClass().New()
}




