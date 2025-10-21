// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestAddArgumentsParams] class.
var (
	MTRUnitTestingClusterTestAddArgumentsParamsClass     _MTRUnitTestingClusterTestAddArgumentsParamsClass
	MTRUnitTestingClusterTestAddArgumentsParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestAddArgumentsParamsClass() _MTRUnitTestingClusterTestAddArgumentsParamsClass {
	MTRUnitTestingClusterTestAddArgumentsParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestAddArgumentsParamsClass = _MTRUnitTestingClusterTestAddArgumentsParamsClass{objc.GetClass("MTRUnitTestingClusterTestAddArgumentsParams")}
	})
	return MTRUnitTestingClusterTestAddArgumentsParamsClass
}

type _MTRUnitTestingClusterTestAddArgumentsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestAddArgumentsParams] class.
type IMTRUnitTestingClusterTestAddArgumentsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestAddArgumentsParams
type MTRUnitTestingClusterTestAddArgumentsParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestAddArgumentsParamsFrom constructs a [MTRUnitTestingClusterTestAddArgumentsParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestAddArgumentsParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestAddArgumentsParams {
	return MTRUnitTestingClusterTestAddArgumentsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestAddArgumentsParamsClass) Alloc() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestAddArgumentsParamsClass) New() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) Init() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestAddArgumentsParams) Autorelease() MTRUnitTestingClusterTestAddArgumentsParams {
	rv := objc.Send[MTRUnitTestingClusterTestAddArgumentsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestAddArgumentsParams creates a new MTRUnitTestingClusterTestAddArgumentsParams instance.
func NewMTRUnitTestingClusterTestAddArgumentsParams() MTRUnitTestingClusterTestAddArgumentsParams {
	return getMTRUnitTestingClusterTestAddArgumentsParamsClass().New()
}




