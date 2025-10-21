// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MotionKeyframeData] class.
var (
	MotionKeyframeDataClass     _MotionKeyframeDataClass
	MotionKeyframeDataClassOnce sync.Once
)

func getMotionKeyframeDataClass() _MotionKeyframeDataClass {
	MotionKeyframeDataClassOnce.Do(func() {
		MotionKeyframeDataClass = _MotionKeyframeDataClass{objc.GetClass("MTLMotionKeyframeData")}
	})
	return MotionKeyframeDataClass
}

type _MotionKeyframeDataClass struct {
	class objc.Class
}

// An interface definition for the [MotionKeyframeData] class.
type IMotionKeyframeData interface {
	objectivec.IObject
}

// Geometry data for a specific keyframe to use in a moving instance.
//
// An instance describes the location of geometry data for a keyframe. The exact type of data can vary, depending on which kind of motion descriptor you create. For an instance, the buffer data is a list of bounding boxes. For an , the buffer data is a list of vertices.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData
type MotionKeyframeData struct {
	objectivec.Object
}

// MotionKeyframeDataFrom constructs a [MotionKeyframeData] from an unsafe.Pointer.
//
// Geometry data for a specific keyframe to use in a moving instance.
func MotionKeyframeDataFrom(ptr unsafe.Pointer) MotionKeyframeData {
	return MotionKeyframeData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MotionKeyframeDataClass) Alloc() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MotionKeyframeDataClass) New() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MotionKeyframeData) Init() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MotionKeyframeData) Autorelease() MotionKeyframeData {
	rv := objc.Send[MotionKeyframeData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMotionKeyframeData creates a new MotionKeyframeData instance.
func NewMotionKeyframeData() MotionKeyframeData {
	return getMotionKeyframeDataClass().New()
}




