// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestListInt8UArgumentRequestParams] class.
var (
	MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass     _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass
	MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass() _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass {
	MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass = _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass{objc.GetClass("MTRUnitTestingClusterTestListInt8UArgumentRequestParams")}
	})
	return MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass
}

type _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestListInt8UArgumentRequestParams] class.
type IMTRUnitTestingClusterTestListInt8UArgumentRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UArgumentRequestParams
type MTRUnitTestingClusterTestListInt8UArgumentRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListInt8UArgumentRequestParamsFrom constructs a [MTRUnitTestingClusterTestListInt8UArgumentRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListInt8UArgumentRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	return MTRUnitTestingClusterTestListInt8UArgumentRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass) Alloc() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass) New() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) Init() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListInt8UArgumentRequestParams) Autorelease() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UArgumentRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListInt8UArgumentRequestParams creates a new MTRUnitTestingClusterTestListInt8UArgumentRequestParams instance.
func NewMTRUnitTestingClusterTestListInt8UArgumentRequestParams() MTRUnitTestingClusterTestListInt8UArgumentRequestParams {
	return getMTRUnitTestingClusterTestListInt8UArgumentRequestParamsClass().New()
}




