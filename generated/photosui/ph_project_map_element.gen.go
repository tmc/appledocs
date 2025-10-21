// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHProjectMapElement] class.
var (
	PHProjectMapElementClass     _PHProjectMapElementClass
	PHProjectMapElementClassOnce sync.Once
)

func getPHProjectMapElementClass() _PHProjectMapElementClass {
	PHProjectMapElementClassOnce.Do(func() {
		PHProjectMapElementClass = _PHProjectMapElementClass{objc.GetClass("PHProjectMapElement")}
	})
	return PHProjectMapElementClass
}

type _PHProjectMapElementClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectMapElement] class.
type IPHProjectMapElement interface {
	IPHProjectElement
}

// An element that represents a map within project section content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement
type PHProjectMapElement struct {
	PHProjectElement
}

// PHProjectMapElementFrom constructs a [PHProjectMapElement] from an unsafe.Pointer.
//
// An element that represents a map within project section content.
func PHProjectMapElementFrom(ptr unsafe.Pointer) PHProjectMapElement {
	return PHProjectMapElement{
		PHProjectElement: PHProjectElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectMapElementClass) Alloc() PHProjectMapElement {
	rv := objc.Send[PHProjectMapElement](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectMapElementClass) New() PHProjectMapElement {
	rv := objc.Send[PHProjectMapElement](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectMapElement) Init() PHProjectMapElement {
	rv := objc.Send[PHProjectMapElement](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectMapElement) Autorelease() PHProjectMapElement {
	rv := objc.Send[PHProjectMapElement](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectMapElement creates a new PHProjectMapElement instance.
func NewPHProjectMapElement() PHProjectMapElement {
	return getPHProjectMapElementClass().New()
}


// The altitude of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement/altitude
func (p_ PHProjectMapElement) Altitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("altitude"))
	return rv
}

// An array of optional annotations attached to the map.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement/annotations
func (p_ PHProjectMapElement) Annotations() []objc.ID {
	rv := objc.Send[[]objc.ID](p_.ID, objc.Sel("annotations"))
	return rv
}

// The coordinates of the center of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement/centerCoordinate
func (p_ PHProjectMapElement) CenterCoordinate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("centerCoordinate"))
	return rv
}

// The heading of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement/heading
func (p_ PHProjectMapElement) Heading() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("heading"))
	return rv
}

// The type of map in the project.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement/mapType
func (p_ PHProjectMapElement) MapType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mapType"))
	return rv
}

// The pitch of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectMapElement/pitch
func (p_ PHProjectMapElement) Pitch() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("pitch"))
	return rv
}



