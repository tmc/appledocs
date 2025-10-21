// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Point3D] class.
var (
	Point3DClass     _Point3DClass
	Point3DClassOnce sync.Once
)

func getPoint3DClass() _Point3DClass {
	Point3DClassOnce.Do(func() {
		Point3DClass = _Point3DClass{objc.GetClass("VNPoint3D")}
	})
	return Point3DClass
}

type _Point3DClass struct {
	class objc.Class
}

// An interface definition for the [Point3D] class.
type IPoint3D interface {
	objectivec.IObject
}

// An object that represents a 3D point in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint3D
type Point3D struct {
	objectivec.Object
}

// Point3DFrom constructs a [Point3D] from an unsafe.Pointer.
//
// An object that represents a 3D point in an image.
func Point3DFrom(ptr unsafe.Pointer) Point3D {
	return Point3D{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _Point3DClass) Alloc() Point3D {
	rv := objc.Send[Point3D](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _Point3DClass) New() Point3D {
	rv := objc.Send[Point3D](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Point3D) Init() Point3D {
	rv := objc.Send[Point3D](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Point3D) Autorelease() Point3D {
	rv := objc.Send[Point3D](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPoint3D creates a new Point3D instance.
func NewPoint3D() Point3D {
	return getPoint3DClass().New()
}


// The three-dimensional position.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint3d/position
func (p_ Point3D) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The three-dimensional position.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint3d/position
func (p_ Point3D) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPosition:"), value)
}



