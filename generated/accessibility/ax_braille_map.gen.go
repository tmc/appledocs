// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXBrailleMap] class.
var (
	AXBrailleMapClass     _AXBrailleMapClass
	AXBrailleMapClassOnce sync.Once
)

func getAXBrailleMapClass() _AXBrailleMapClass {
	AXBrailleMapClassOnce.Do(func() {
		AXBrailleMapClass = _AXBrailleMapClass{objc.GetClass("AXBrailleMap")}
	})
	return AXBrailleMapClass
}

type _AXBrailleMapClass struct {
	class objc.Class
}

// An interface definition for the [AXBrailleMap] class.
type IAXBrailleMap interface {
	objectivec.IObject
	HeightAtPoint(point coregraphics.CGPoint) float32
	PresentImage(image coregraphics.CGImageRef)
	SetHeightAtPoint(status float32, point coregraphics.CGPoint)
	Dimensions() coregraphics.CGSize
}

// A representation of a two-dimensional braille display.
//
// A braille map object represents a two-dimensional braille display that’s connected to the current Apple device. By specifying the dot patterns in the braille map, you can change the content the user experiences. To render the data from the braille map to the display, implement .


// A representation of a two-dimensional braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap
type AXBrailleMap struct {
	objectivec.Object
}

// AXBrailleMapFrom constructs a [AXBrailleMap] from an unsafe.Pointer.
//
// A representation of a two-dimensional braille display.
func AXBrailleMapFrom(ptr unsafe.Pointer) AXBrailleMap {
	return AXBrailleMap{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXBrailleMapClass) Alloc() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXBrailleMapClass) New() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXBrailleMap) Init() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXBrailleMap) Autorelease() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleMap creates a new AXBrailleMap instance.
func NewAXBrailleMap() AXBrailleMap {
	return getAXBrailleMapClass().New()
}



// Retrieves the height of an individual pin on the braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/height(at:)
func (a_ AXBrailleMap) HeightAtPoint(point coregraphics.CGPoint) float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("heightAtPoint:"), point)
	return rv
}


// Converts the data from the image you specify into the braille map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/present(_:)
func (a_ AXBrailleMap) PresentImage(image coregraphics.CGImageRef) {
	objc.Send[objc.ID](a_.ID, objc.Sel("presentImage:"), image)
}


// Sets the height of an individual pin on the braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/setHeight(_:at:)
func (a_ AXBrailleMap) SetHeightAtPoint(status float32, point coregraphics.CGPoint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHeight:atPoint:"), status, point)
}


// The number of pins in each dimension of the braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/dimensions
func (a_ AXBrailleMap) Dimensions() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("dimensions"))
	return rv
}



