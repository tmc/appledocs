// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKShape] class.
var (
	MKShapeClass     _MKShapeClass
	MKShapeClassOnce sync.Once
)

func getMKShapeClass() _MKShapeClass {
	MKShapeClassOnce.Do(func() {
		MKShapeClass = _MKShapeClass{objc.GetClass("MKShape")}
	})
	return MKShapeClass
}

type _MKShapeClass struct {
	class objc.Class
}

// An interface definition for the [MKShape] class.
type IMKShape interface {
	objectivec.IObject
	// properties:
	Subtitle() string
	SetSubtitle(value string)
	Title() string
	SetTitle(value string)
	// methods:
}

// An abstract class that defines the basic properties for all shape-based overlay objects.
//
// You can’t instantiate this class directly; use a subclass instead. Subclasses are responsible for defining the geometry of the shape and providing an appropriate value for the coordinate property they inherit from the protocol.


// An abstract class that defines the basic properties for all shape-based overlay objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKShape
type MKShape struct {
	objectivec.Object
}

// MKShapeFrom constructs a [MKShape] from an unsafe.Pointer.
//
// An abstract class that defines the basic properties for all shape-based overlay objects.
func MKShapeFrom(ptr unsafe.Pointer) MKShape {
	return MKShape{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKShapeClass) Alloc() MKShape {
	rv := objc.Send[MKShape](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKShapeClass) New() MKShape {
	rv := objc.Send[MKShape](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKShape) Init() MKShape {
	rv := objc.Send[MKShape](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKShape) Autorelease() MKShape {
	rv := objc.Send[MKShape](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKShape creates a new MKShape instance.
func NewMKShape() MKShape {
	return getMKShapeClass().New()
}



// The subtitle of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkshape/subtitle
func (m_ MKShape) Subtitle() string {
	rv := objc.Send[string](m_.ID, objc.Sel("subtitle"))
	return rv
}


// The subtitle of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkshape/subtitle
func (m_ MKShape) SetSubtitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}


// The title of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkshape/title
func (m_ MKShape) Title() string {
	rv := objc.Send[string](m_.ID, objc.Sel("title"))
	return rv
}


// The title of the shape annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkshape/title
func (m_ MKShape) SetTitle(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), objc.String(value))
}



