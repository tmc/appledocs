// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HumanBodyRecognizedPoint3D] class.
var (
	HumanBodyRecognizedPoint3DClass     _HumanBodyRecognizedPoint3DClass
	HumanBodyRecognizedPoint3DClassOnce sync.Once
)

func getHumanBodyRecognizedPoint3DClass() _HumanBodyRecognizedPoint3DClass {
	HumanBodyRecognizedPoint3DClassOnce.Do(func() {
		HumanBodyRecognizedPoint3DClass = _HumanBodyRecognizedPoint3DClass{objc.GetClass("VNHumanBodyRecognizedPoint3D")}
	})
	return HumanBodyRecognizedPoint3DClass
}

type _HumanBodyRecognizedPoint3DClass struct {
	class objc.Class
}

// An interface definition for the [HumanBodyRecognizedPoint3D] class.
type IHumanBodyRecognizedPoint3D interface {
	IRecognizedPoint3D
}

// A recognized 3D point that includes a parent joint.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D
type HumanBodyRecognizedPoint3D struct {
	RecognizedPoint3D
}

// HumanBodyRecognizedPoint3DFrom constructs a [HumanBodyRecognizedPoint3D] from an unsafe.Pointer.
//
// A recognized 3D point that includes a parent joint.
func HumanBodyRecognizedPoint3DFrom(ptr unsafe.Pointer) HumanBodyRecognizedPoint3D {
	return HumanBodyRecognizedPoint3D{
		RecognizedPoint3D: RecognizedPoint3DFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HumanBodyRecognizedPoint3DClass) Alloc() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HumanBodyRecognizedPoint3DClass) New() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanBodyRecognizedPoint3D) Init() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanBodyRecognizedPoint3D) Autorelease() HumanBodyRecognizedPoint3D {
	rv := objc.Send[HumanBodyRecognizedPoint3D](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanBodyRecognizedPoint3D creates a new HumanBodyRecognizedPoint3D instance.
func NewHumanBodyRecognizedPoint3D() HumanBodyRecognizedPoint3D {
	return getHumanBodyRecognizedPoint3DClass().New()
}


// The three-dimensional position.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyrecognizedpoint3d/localposition
func (h_ HumanBodyRecognizedPoint3D) LocalPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("localPosition"))
	return rv
}


// SetLocalPosition sets the value of the localPosition property.
// The three-dimensional position.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyrecognizedpoint3d/localposition
func (h_ HumanBodyRecognizedPoint3D) SetLocalPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLocalPosition:"), value)
}

// The parent joint in the observation.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyrecognizedpoint3d/parentjoint
func (h_ HumanBodyRecognizedPoint3D) ParentJoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("parentJoint"))
	return rv
}


// SetParentJoint sets the value of the parentJoint property.
// The parent joint in the observation.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhumanbodyrecognizedpoint3d/parentjoint
func (h_ HumanBodyRecognizedPoint3D) SetParentJoint(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setParentJoint:"), value)
}



