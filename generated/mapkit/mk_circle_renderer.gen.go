// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKCircleRenderer] class.
var (
	MKCircleRendererClass     _MKCircleRendererClass
	MKCircleRendererClassOnce sync.Once
)

func getMKCircleRendererClass() _MKCircleRendererClass {
	MKCircleRendererClassOnce.Do(func() {
		MKCircleRendererClass = _MKCircleRendererClass{objc.GetClass("MKCircleRenderer")}
	})
	return MKCircleRendererClass
}

type _MKCircleRendererClass struct {
	class objc.Class
}

// An interface definition for the [MKCircleRenderer] class.
type IMKCircleRenderer interface {
	IMKOverlayPathRenderer
}

// The visual representation of a circular overlay.
//
// This renderer fills and strokes the circular region that the overlay object represents. You can change the color and other drawing attributes of the circle by modifying the properties it inherits from the main class. You typically use this class as-is and don’t subclass it. You create an instance of this class in your map view delegate’s method.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer
type MKCircleRenderer struct {
	MKOverlayPathRenderer
}

// MKCircleRendererFrom constructs a [MKCircleRenderer] from an unsafe.Pointer.
//
// The visual representation of a circular overlay.
func MKCircleRendererFrom(ptr unsafe.Pointer) MKCircleRenderer {
	return MKCircleRenderer{
		MKOverlayPathRenderer: MKOverlayPathRendererFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKCircleRendererClass) Alloc() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKCircleRendererClass) New() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKCircleRenderer) Init() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKCircleRenderer) Autorelease() MKCircleRenderer {
	rv := objc.Send[MKCircleRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKCircleRenderer creates a new MKCircleRenderer instance.
func NewMKCircleRenderer() MKCircleRenderer {
	return getMKCircleRendererClass().New()
}




// Creates a new overlay view using the specified circle overlay object.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/init(circle:)
func NewMKCircleRendererWithCircle(circle unsafe.Pointer) MKCircleRenderer {
	instance := getMKCircleRendererClass().Alloc()
	rv := objc.Send[MKCircleRenderer](instance.ID, objc.Sel("initWithCircle:"), circle)
	rv.Autorelease()
	return rv
}


// The circle overlay object that contains the information for drawing the overlay.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/circle
func (m_ MKCircleRenderer) Circle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("circle"))
	return rv
}

// The unit distance along the circle where the stroke ends.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeEnd
func (m_ MKCircleRenderer) StrokeEnd() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeEnd"))
	return rv
}


// SetStrokeEnd sets the value of the strokeEnd property.
// The unit distance along the circle where the stroke ends.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeEnd
func (m_ MKCircleRenderer) SetStrokeEnd(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeEnd:"), value)
}

// The unit distance along the circle where the stroke starts.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeStart
func (m_ MKCircleRenderer) StrokeStart() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("strokeStart"))
	return rv
}


// SetStrokeStart sets the value of the strokeStart property.
// The unit distance along the circle where the stroke starts.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleRenderer/strokeStart
func (m_ MKCircleRenderer) SetStrokeStart(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeStart:"), value)
}


