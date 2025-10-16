
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Slider] class.
var SliderClass _SliderClass

func init() {
	SliderClass = _SliderClass{objc.GetClass("NSSlider")}
}

type _SliderClass struct {
	objc.Class
}

// An interface definition for the [Slider] class.
type ISlider interface {
	ID() objc.ID
	SetTitleFont(fontObj unsafe.Pointer)
}

type Slider struct {
	id objc.ID
}

func SliderFrom(ptr unsafe.Pointer) Slider {
	return Slider{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Slider) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SliderClass) Alloc() Slider {
	rv := objc.Send[Slider](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SliderClass) New() Slider {
	rv := objc.Send[Slider](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSlider creates and returns a new initialized instance.
func NewSlider() Slider {
	return SliderClass.New()
}

// Init initializes the instance.
func (s_ Slider) Init() Slider {
	rv := objc.Send[Slider](s_.ID(), selInit)
	return rv
}
// Sets the font used to draw the slider’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSlider/setTitleFont:
func (s_ Slider) SetTitleFont(fontObj unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTitleFont:"), fontObj)
}
// The maximum value the slider can send to its target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("maxValue"))
	return rv
}
// SetMaxValue sets the value of the maxValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSlider/maxValue
func (s_ Slider) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMaxValue:"), value)
}
// The color of the filled portion of the slider track, in appearances that support it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) TrackFillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("trackFillColor"))
	return rv
}
// SetTrackFillColor sets the value of the trackFillColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSlider/trackFillColor
func (s_ Slider) SetTrackFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTrackFillColor:"), value)
}
