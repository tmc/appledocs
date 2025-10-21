// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnOffClusterOnWithRecallGlobalSceneParams] class.
var (
	MTROnOffClusterOnWithRecallGlobalSceneParamsClass     _MTROnOffClusterOnWithRecallGlobalSceneParamsClass
	MTROnOffClusterOnWithRecallGlobalSceneParamsClassOnce sync.Once
)

func getMTROnOffClusterOnWithRecallGlobalSceneParamsClass() _MTROnOffClusterOnWithRecallGlobalSceneParamsClass {
	MTROnOffClusterOnWithRecallGlobalSceneParamsClassOnce.Do(func() {
		MTROnOffClusterOnWithRecallGlobalSceneParamsClass = _MTROnOffClusterOnWithRecallGlobalSceneParamsClass{objc.GetClass("MTROnOffClusterOnWithRecallGlobalSceneParams")}
	})
	return MTROnOffClusterOnWithRecallGlobalSceneParamsClass
}

type _MTROnOffClusterOnWithRecallGlobalSceneParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterOnWithRecallGlobalSceneParams] class.
type IMTROnOffClusterOnWithRecallGlobalSceneParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnWithRecallGlobalSceneParams
type MTROnOffClusterOnWithRecallGlobalSceneParams struct {
	objectivec.Object
}

// MTROnOffClusterOnWithRecallGlobalSceneParamsFrom constructs a [MTROnOffClusterOnWithRecallGlobalSceneParams] from an unsafe.Pointer.
func MTROnOffClusterOnWithRecallGlobalSceneParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOnWithRecallGlobalSceneParams {
	return MTROnOffClusterOnWithRecallGlobalSceneParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOnWithRecallGlobalSceneParamsClass) Alloc() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterOnWithRecallGlobalSceneParamsClass) New() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) Init() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOnWithRecallGlobalSceneParams) Autorelease() MTROnOffClusterOnWithRecallGlobalSceneParams {
	rv := objc.Send[MTROnOffClusterOnWithRecallGlobalSceneParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOnWithRecallGlobalSceneParams creates a new MTROnOffClusterOnWithRecallGlobalSceneParams instance.
func NewMTROnOffClusterOnWithRecallGlobalSceneParams() MTROnOffClusterOnWithRecallGlobalSceneParams {
	return getMTROnOffClusterOnWithRecallGlobalSceneParamsClass().New()
}




