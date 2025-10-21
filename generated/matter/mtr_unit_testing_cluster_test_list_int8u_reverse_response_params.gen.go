// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestListInt8UReverseResponseParams] class.
var (
	MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass     _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass
	MTRUnitTestingClusterTestListInt8UReverseResponseParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestListInt8UReverseResponseParamsClass() _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass {
	MTRUnitTestingClusterTestListInt8UReverseResponseParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass = _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass{objc.GetClass("MTRUnitTestingClusterTestListInt8UReverseResponseParams")}
	})
	return MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass
}

type _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestListInt8UReverseResponseParams] class.
type IMTRUnitTestingClusterTestListInt8UReverseResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestListInt8UReverseResponseParams
type MTRUnitTestingClusterTestListInt8UReverseResponseParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestListInt8UReverseResponseParamsFrom constructs a [MTRUnitTestingClusterTestListInt8UReverseResponseParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestListInt8UReverseResponseParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	return MTRUnitTestingClusterTestListInt8UReverseResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass) Alloc() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestListInt8UReverseResponseParamsClass) New() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) Init() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestListInt8UReverseResponseParams) Autorelease() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	rv := objc.Send[MTRUnitTestingClusterTestListInt8UReverseResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestListInt8UReverseResponseParams creates a new MTRUnitTestingClusterTestListInt8UReverseResponseParams instance.
func NewMTRUnitTestingClusterTestListInt8UReverseResponseParams() MTRUnitTestingClusterTestListInt8UReverseResponseParams {
	return getMTRUnitTestingClusterTestListInt8UReverseResponseParamsClass().New()
}




