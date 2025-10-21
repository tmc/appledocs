// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] class.
var (
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass     _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass() _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass {
	MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass = _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestStructArrayArgumentResponseParams")}
	})
	return MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass
}

type _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] class.
type IMTRUnitTestingClusterTestStructArrayArgumentResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestStructArrayArgumentResponseParams
type MTRUnitTestingClusterTestStructArrayArgumentResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom constructs a [MTRUnitTestingClusterTestStructArrayArgumentResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestStructArrayArgumentResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	return MTRUnitTestingClusterTestStructArrayArgumentResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass) Alloc() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass) New() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Init() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestStructArrayArgumentResponseParams) Autorelease() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestStructArrayArgumentResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestStructArrayArgumentResponseParams creates a new MTRUnitTestingClusterTestStructArrayArgumentResponseParams instance.
func NewMTRUnitTestingClusterTestStructArrayArgumentResponseParams() MTRUnitTestingClusterTestStructArrayArgumentResponseParams {
	return getMTRUnitTestingClusterTestStructArrayArgumentResponseParamsClass().New()
}




