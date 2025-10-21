// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestStructArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestStructArgumentRequestParamsClass     _MTRUnitTestingClusterTestStructArgumentRequestParamsClass
	MTRUnitTestingClusterTestStructArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArgumentRequestParamsClass() _MTRUnitTestingClusterTestStructArgumentRequestParamsClass {
	MTRUnitTestingClusterTestStructArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArgumentRequestParamsClass = _MTRUnitTestingClusterTestStructArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestStructArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestStructArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestStructArgumentRequestParams] class.
type IMTRUnitTestingClusterTestStructArgumentRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArgumentRequestParams
type MTRUnitTestingClusterTestStructArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestStructArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArgumentRequestParams {
	return MTRUnitTestingClusterTestStructArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestStructArgumentRequestParamsClass) New() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) Init() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestStructArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArgumentRequestParams creates a new MTRUnitTestingClusterTestStructArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestStructArgumentRequestParams() MTRUnitTestingClusterTestStructArgumentRequestParams {
	return getMTRUnitTestingClusterTestStructArgumentRequestParamsClass().New()
}




