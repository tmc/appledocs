// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [SliderCell] class.
var (
	SliderCellClass     _SliderCellClass
	SliderCellClassOnce sync.Once
)

func getSliderCellClass() _SliderCellClass {
	SliderCellClassOnce.Do(func() {
		SliderCellClass = _SliderCellClass{objc.GetClass("NSSliderCell")}
	})
	return SliderCellClass
}

type _SliderCellClass struct {
	class objc.Class
}

// An interface definition for the [SliderCell] class.
type ISliderCell interface {
	IActionCell
	KnobRectFlipped(flipped bool) coregraphics.CGRect
}

// The appearance and behavior of an object.
//
// You can customize an to a certain degree, using its properties. If this doesn’t give you sufficient flexibility, you can create a subclass. In that subclass, you can override any of the following methods: , , , and .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell
type SliderCell struct {
	ActionCell
}

// SliderCellFrom constructs a [SliderCell] from an unsafe.Pointer.
//
// The appearance and behavior of an object.
func SliderCellFrom(ptr unsafe.Pointer) SliderCell {
	return SliderCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SliderCellClass) Alloc() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SliderCellClass) New() SliderCell {
	rv := objc.Send[SliderCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderCell) Init() SliderCell {
	rv := objc.Send[SliderCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderCell) Autorelease() SliderCell {
	rv := objc.Send[SliderCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderCell creates a new SliderCell instance.
func NewSliderCell() SliderCell {
	return getSliderCellClass().New()
}


// Returns the rectangle in which the slider knob is drawn.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/knobRect(flipped:)
func (s_ SliderCell) KnobRectFlipped(flipped bool) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("knobRectFlipped:"), flipped)
	return rv
}

// The slider type, either linear or circular.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s_ SliderCell) SliderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sliderType"))
	return rv
}


// SetSliderType sets the value of the sliderType property.
// The slider type, either linear or circular.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderCell/sliderType
func (s_ SliderCell) SetSliderType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSliderType:"), value)
}



