// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RecognizedPoint3D] class.
var (
	RecognizedPoint3DClass     _RecognizedPoint3DClass
	RecognizedPoint3DClassOnce sync.Once
)

func getRecognizedPoint3DClass() _RecognizedPoint3DClass {
	RecognizedPoint3DClassOnce.Do(func() {
		RecognizedPoint3DClass = _RecognizedPoint3DClass{objc.GetClass("VNRecognizedPoint3D")}
	})
	return RecognizedPoint3DClass
}

type _RecognizedPoint3DClass struct {
	class objc.Class
}

// An interface definition for the [RecognizedPoint3D] class.
type IRecognizedPoint3D interface {
	IPoint3D
}

// A 3D point that includes an identifier to the point.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D
type RecognizedPoint3D struct {
	Point3D
}

// RecognizedPoint3DFrom constructs a [RecognizedPoint3D] from an unsafe.Pointer.
//
// A 3D point that includes an identifier to the point.
func RecognizedPoint3DFrom(ptr unsafe.Pointer) RecognizedPoint3D {
	return RecognizedPoint3D{
		Point3D: Point3DFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizedPoint3DClass) Alloc() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizedPoint3DClass) New() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedPoint3D) Init() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedPoint3D) Autorelease() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedPoint3D creates a new RecognizedPoint3D instance.
func NewRecognizedPoint3D() RecognizedPoint3D {
	return getRecognizedPoint3DClass().New()
}


// The identifier that provides context about what kind of point the request recognizes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D/identifier
func (r_ RecognizedPoint3D) Identifier() RecognizedPointKey {
	rv := objc.Send[RecognizedPointKey](r_.ID, objc.Sel("identifier"))
	return rv
}



