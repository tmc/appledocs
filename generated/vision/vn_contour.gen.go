// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Contour] class.
var (
	ContourClass     _ContourClass
	ContourClassOnce sync.Once
)

func getContourClass() _ContourClass {
	ContourClassOnce.Do(func() {
		ContourClass = _ContourClass{objc.GetClass("VNContour")}
	})
	return ContourClass
}

type _ContourClass struct {
	class objc.Class
}

// An interface definition for the [Contour] class.
type IContour interface {
	objectivec.IObject
}

// A class that represents a detected contour in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour
type Contour struct {
	objectivec.Object
}

// ContourFrom constructs a [Contour] from an unsafe.Pointer.
//
// A class that represents a detected contour in an image.
func ContourFrom(ptr unsafe.Pointer) Contour {
	return Contour{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContourClass) Alloc() Contour {
	rv := objc.Send[Contour](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContourClass) New() Contour {
	rv := objc.Send[Contour](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Contour) Init() Contour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Contour) Autorelease() Contour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContour creates a new Contour instance.
func NewContour() Contour {
	return getContourClass().New()
}




