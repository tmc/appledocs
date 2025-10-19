// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderInfo] class.
var (
	renderInfoClass     _RenderInfoClass
	renderInfoClassOnce sync.Once
)

func getRenderInfoClass() _RenderInfoClass {
	renderInfoClassOnce.Do(func() {
		renderInfoClass = _RenderInfoClass{objc.GetClass("CIRenderInfo")}
	})
	return renderInfoClass
}

type _RenderInfoClass struct {
	class objc.Class
}

// An interface definition for the [RenderInfo] class.
type IRenderInfo interface {
	objectivec.IObject
}

// An encapsulation of a render task’s timing, passes, and pixels processed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderInfo
type RenderInfo struct {
	objectivec.Object
}

// RenderInfoFrom constructs a [RenderInfo] from an unsafe.Pointer.
//
// An encapsulation of a render task’s timing, passes, and pixels processed.
func RenderInfoFrom(ptr unsafe.Pointer) RenderInfo {
	return RenderInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RenderInfoClass) Alloc() RenderInfo {
	rv := objc.Send[RenderInfo](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RenderInfoClass) New() RenderInfo {
	rv := objc.Send[RenderInfo](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderInfo) Init() RenderInfo {
	rv := objc.Send[RenderInfo](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderInfo) Autorelease() RenderInfo {
	rv := objc.Send[RenderInfo](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderInfo creates a new RenderInfo instance.
func NewRenderInfo() RenderInfo {
	return getRenderInfoClass().New()
}




