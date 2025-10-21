// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEnumsResponseParams] class.
var (
	MTRUnitTestingClusterTestEnumsResponseParamsClass     _MTRUnitTestingClusterTestEnumsResponseParamsClass
	MTRUnitTestingClusterTestEnumsResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEnumsResponseParamsClass() _MTRUnitTestingClusterTestEnumsResponseParamsClass {
	MTRUnitTestingClusterTestEnumsResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestEnumsResponseParamsClass = _MTRUnitTestingClusterTestEnumsResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestEnumsResponseParams")}
	})
	return MTRUnitTestingClusterTestEnumsResponseParamsClass
}

type _MTRUnitTestingClusterTestEnumsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEnumsResponseParams] class.
type IMTRUnitTestingClusterTestEnumsResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEnumsResponseParams
type MTRUnitTestingClusterTestEnumsResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEnumsResponseParamsFrom constructs a [MTRUnitTestingClusterTestEnumsResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEnumsResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEnumsResponseParams {
	return MTRUnitTestingClusterTestEnumsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEnumsResponseParamsClass) Alloc() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEnumsResponseParamsClass) New() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) Init() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEnumsResponseParams) Autorelease() MTRUnitTestingClusterTestEnumsResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestEnumsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEnumsResponseParams creates a new MTRUnitTestingClusterTestEnumsResponseParams instance.
func NewMTRUnitTestingClusterTestEnumsResponseParams() MTRUnitTestingClusterTestEnumsResponseParams {
	return getMTRUnitTestingClusterTestEnumsResponseParamsClass().New()
}




