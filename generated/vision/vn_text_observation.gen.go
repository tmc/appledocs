// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextObservation] class.
var (
	TextObservationClass     _TextObservationClass
	TextObservationClassOnce sync.Once
)

func getTextObservationClass() _TextObservationClass {
	TextObservationClassOnce.Do(func() {
		TextObservationClass = _TextObservationClass{objc.GetClass("VNTextObservation")}
	})
	return TextObservationClass
}

type _TextObservationClass struct {
	class objc.Class
}

// An interface definition for the [TextObservation] class.
type ITextObservation interface {
	IRectangleObservation
}

// Information about regions of text that an image-analysis request detects.
//
// This type of observation results from a . It expresses the location of each detected character by its bounding box.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTextObservation
type TextObservation struct {
	RectangleObservation
}

// TextObservationFrom constructs a [TextObservation] from an unsafe.Pointer.
//
// Information about regions of text that an image-analysis request detects.
func TextObservationFrom(ptr unsafe.Pointer) TextObservation {
	return TextObservation{
		RectangleObservation: RectangleObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextObservationClass) Alloc() TextObservation {
	rv := objc.Send[TextObservation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextObservationClass) New() TextObservation {
	rv := objc.Send[TextObservation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextObservation) Init() TextObservation {
	rv := objc.Send[TextObservation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextObservation) Autorelease() TextObservation {
	rv := objc.Send[TextObservation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextObservation creates a new TextObservation instance.
func NewTextObservation() TextObservation {
	return getTextObservationClass().New()
}


// An array of detected individual character bounding boxes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTextObservation/characterBoxes
func (t_ TextObservation) CharacterBoxes() []RectangleObservation {
	rv := objc.Send[[]RectangleObservation](t_.ID, objc.Sel("characterBoxes"))
	return rv
}



