// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SliderCell] class.
var (
	sliderCellClass     _SliderCellClass
	sliderCellClassOnce sync.Once
)

func getSliderCellClass() _SliderCellClass {
	sliderCellClassOnce.Do(func() {
		sliderCellClass = _SliderCellClass{objc.GetClass("NSSliderCell")}
	})
	return sliderCellClass
}

type _SliderCellClass struct {
	class objc.Class
}

// An interface definition for the [SliderCell] class.
type ISliderCell interface {
	IActionCell
}

// The appearance and behavior of an object.
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




