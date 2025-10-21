// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterSimpleStructEchoRequestParams] class.
var (
	MTRUnitTestingClusterSimpleStructEchoRequestParamsClass     _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass
	MTRUnitTestingClusterSimpleStructEchoRequestParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructEchoRequestParamsClass() _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass {
	MTRUnitTestingClusterSimpleStructEchoRequestParamsClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructEchoRequestParamsClass = _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass{objc.GetClass("MTRUnitTestingClusterSimpleStructEchoRequestParams")}
	})
	return MTRUnitTestingClusterSimpleStructEchoRequestParamsClass
}

type _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterSimpleStructEchoRequestParams] class.
type IMTRUnitTestingClusterSimpleStructEchoRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStructEchoRequestParams
type MTRUnitTestingClusterSimpleStructEchoRequestParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructEchoRequestParamsFrom constructs a [MTRUnitTestingClusterSimpleStructEchoRequestParams] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructEchoRequestParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStructEchoRequestParams {
	return MTRUnitTestingClusterSimpleStructEchoRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass) Alloc() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterSimpleStructEchoRequestParamsClass) New() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) Init() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStructEchoRequestParams) Autorelease() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	rv := objc.Send[MTRUnitTestingClusterSimpleStructEchoRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStructEchoRequestParams creates a new MTRUnitTestingClusterSimpleStructEchoRequestParams instance.
func NewMTRUnitTestingClusterSimpleStructEchoRequestParams() MTRUnitTestingClusterSimpleStructEchoRequestParams {
	return getMTRUnitTestingClusterSimpleStructEchoRequestParamsClass().New()
}




