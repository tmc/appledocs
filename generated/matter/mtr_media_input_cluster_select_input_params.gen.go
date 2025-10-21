// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaInputClusterSelectInputParams] class.
var (
	MTRMediaInputClusterSelectInputParamsClass     _MTRMediaInputClusterSelectInputParamsClass
	MTRMediaInputClusterSelectInputParamsClassOnce sync.Once
)

func getMTRMediaInputClusterSelectInputParamsClass() _MTRMediaInputClusterSelectInputParamsClass {
	MTRMediaInputClusterSelectInputParamsClassOnce.Do(func() {
		MTRMediaInputClusterSelectInputParamsClass = _MTRMediaInputClusterSelectInputParamsClass{objc.GetClass("MTRMediaInputClusterSelectInputParams")}
	})
	return MTRMediaInputClusterSelectInputParamsClass
}

type _MTRMediaInputClusterSelectInputParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterSelectInputParams] class.
type IMTRMediaInputClusterSelectInputParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterSelectInputParams
type MTRMediaInputClusterSelectInputParams struct {
	objectivec.Object
}

// MTRMediaInputClusterSelectInputParamsFrom constructs a [MTRMediaInputClusterSelectInputParams] from an unsafe.Pointer.
func MTRMediaInputClusterSelectInputParamsFrom(ptr unsafe.Pointer) MTRMediaInputClusterSelectInputParams {
	return MTRMediaInputClusterSelectInputParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterSelectInputParamsClass) Alloc() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterSelectInputParamsClass) New() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterSelectInputParams) Init() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterSelectInputParams) Autorelease() MTRMediaInputClusterSelectInputParams {
	rv := objc.Send[MTRMediaInputClusterSelectInputParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterSelectInputParams creates a new MTRMediaInputClusterSelectInputParams instance.
func NewMTRMediaInputClusterSelectInputParams() MTRMediaInputClusterSelectInputParams {
	return getMTRMediaInputClusterSelectInputParamsClass().New()
}




