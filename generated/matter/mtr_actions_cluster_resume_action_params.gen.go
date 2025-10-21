// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterResumeActionParams] class.
var (
	MTRActionsClusterResumeActionParamsClass     _MTRActionsClusterResumeActionParamsClass
	MTRActionsClusterResumeActionParamsClassOnce sync.Once
)

func getMTRActionsClusterResumeActionParamsClass() _MTRActionsClusterResumeActionParamsClass {
	MTRActionsClusterResumeActionParamsClassOnce.Do(func() {
		MTRActionsClusterResumeActionParamsClass = _MTRActionsClusterResumeActionParamsClass{objc.GetClass("MTRActionsClusterResumeActionParams")}
	})
	return MTRActionsClusterResumeActionParamsClass
}

type _MTRActionsClusterResumeActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterResumeActionParams] class.
type IMTRActionsClusterResumeActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterResumeActionParams
type MTRActionsClusterResumeActionParams struct {
	objectivec.Object
}

// MTRActionsClusterResumeActionParamsFrom constructs a [MTRActionsClusterResumeActionParams] from an unsafe.Pointer.
func MTRActionsClusterResumeActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterResumeActionParams {
	return MTRActionsClusterResumeActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterResumeActionParamsClass) Alloc() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterResumeActionParamsClass) New() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterResumeActionParams) Init() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterResumeActionParams) Autorelease() MTRActionsClusterResumeActionParams {
	rv := objc.Send[MTRActionsClusterResumeActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterResumeActionParams creates a new MTRActionsClusterResumeActionParams instance.
func NewMTRActionsClusterResumeActionParams() MTRActionsClusterResumeActionParams {
	return getMTRActionsClusterResumeActionParamsClass().New()
}




