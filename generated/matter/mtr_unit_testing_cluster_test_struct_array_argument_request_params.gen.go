// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass     _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass() _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass {
	MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass = _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArrayArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] class.
type IMTRUnitTestingClusterTestStructArrayArgumentRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentRequestParams
type MTRUnitTestingClusterTestStructArrayArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestStructArrayArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArrayArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	return MTRUnitTestingClusterTestStructArrayArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass) New() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Init() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArrayArgumentRequestParams creates a new MTRUnitTestingClusterTestStructArrayArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestStructArrayArgumentRequestParams() MTRUnitTestingClusterTestStructArrayArgumentRequestParams {
	return getMTRUnitTestingClusterTestStructArrayArgumentRequestParamsClass().New()
}




